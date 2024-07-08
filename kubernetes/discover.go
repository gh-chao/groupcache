package kubernetes

import (
	"context"
	"fmt"
	"github.com/gh-chao/groupcache"
	"github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"net"
	"reflect"
)

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

type Discover struct {
	client kubernetes.Interface

	log logrus.FieldLogger

	namespace string
	selector  string
	port      string

	poll *groupcache.HTTPPool
}

func NewDiscover(pool *groupcache.HTTPPool, client kubernetes.Interface, namespace string, selector string, port string, log logrus.FieldLogger) *Discover {
	discover := &Discover{
		poll:      pool,
		namespace: namespace,
		selector:  selector,
		port:      port,
		client:    client,
		log:       log,
	}
	return discover
}

func (d *Discover) WatchPod(ctx context.Context) error {
	listWatch := &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			options.LabelSelector = d.selector
			return d.client.CoreV1().Pods(d.namespace).List(ctx, options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			options.LabelSelector = d.selector
			return d.client.CoreV1().Pods(d.namespace).Watch(ctx, options)
		},
	}

	informer := cache.NewSharedIndexInformer(
		listWatch,
		&apiv1.Pod{},
		0,
		cache.Indexers{},
	)

	handleEvent := func(obj interface{}) {
		key, _ := cache.MetaNamespaceKeyFunc(obj)
		d.log.Debugf("Queue Pod '%s'", key)
		d.updatePeersFromPods(informer.GetStore().List())
	}

	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    handleEvent,
		UpdateFunc: func(oldObj, newObj interface{}) { handleEvent(newObj) },
		DeleteFunc: handleEvent,
	})
	if err != nil {
		return err
	}

	go informer.Run(ctx.Done())

	if !cache.WaitForCacheSync(ctx.Done(), informer.HasSynced) {
		return fmt.Errorf("timed out waiting for caches to sync")
	}

	return nil
}

func (d *Discover) WatchEndpoint(ctx context.Context) error {
	listWatch := &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			options.LabelSelector = d.selector
			return d.client.CoreV1().Endpoints(d.namespace).List(ctx, options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			options.LabelSelector = d.selector
			return d.client.CoreV1().Endpoints(d.namespace).Watch(ctx, options)
		},
	}

	informer := cache.NewSharedIndexInformer(
		listWatch,
		&apiv1.Endpoints{},
		0,
		cache.Indexers{},
	)

	handleEvent := func(obj interface{}) {
		key, _ := cache.MetaNamespaceKeyFunc(obj)
		d.log.Debugf("Queue '%s'", key)
		d.updatePeersFromEndpoints(informer.GetStore().List())
	}

	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    handleEvent,
		UpdateFunc: func(oldObj, newObj interface{}) { handleEvent(newObj) },
		DeleteFunc: handleEvent,
	})
	if err != nil {
		return err
	}

	go informer.Run(ctx.Done())

	if !cache.WaitForCacheSync(ctx.Done(), informer.HasSynced) {
		return fmt.Errorf("timed out waiting for caches to sync")
	}

	return nil
}

func (d *Discover) updatePeersFromPods(objs []any) {
	d.log.Debug("Fetching peer list from pods API")
	var peers []string
main:
	for _, obj := range objs {
		pod, ok := obj.(*apiv1.Pod)
		if !ok {
			d.log.Errorf("expected type v1.Endpoints got '%s' instead", reflect.TypeOf(obj).String())
		}

		peer := fmt.Sprintf("http://%s:%s", pod.Status.PodIP, d.port)

		// if containers are not ready or not running then skip this peer
		for _, status := range pod.Status.ContainerStatuses {
			if !status.Ready || status.State.Running == nil {
				d.log.Debugf("Skipping peer because it's not ready or not running: %+v\n", peer)
				continue main
			}
		}

		d.log.Debugf("Peer: %+v\n", peer)
		peers = append(peers, peer)
	}
	d.poll.Set(peers...)
}

func (d *Discover) updatePeersFromEndpoints(objs []any) {
	d.log.Debug("Fetching peer list from endpoints API")
	var peers []string
	for _, obj := range objs {
		endpoint, ok := obj.(*apiv1.Endpoints)
		if !ok {
			d.log.Errorf("expected type v1.Endpoints got '%s' instead", reflect.TypeOf(obj).String())
		}

		for _, s := range endpoint.Subsets {
			for _, addr := range s.Addresses {
				peers = append(peers, fmt.Sprintf("%s:%s", addr.IP, d.port))
				d.log.Debugf("Peer: %+v\n", fmt.Sprintf("http://%s:%s", addr.IP, d.port))
			}
		}
	}
	d.poll.Set(peers...)
}
