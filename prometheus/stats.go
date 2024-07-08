package prometheus

import "github.com/gh-chao/groupcache"

type GroupStatistics struct {
	group *groupcache.Group
}

// New creates a new Group.
func New(group *groupcache.Group) *GroupStatistics {
	return &GroupStatistics{group: group}
}

// Name returns the group's name
func (g *GroupStatistics) Name() string {
	return g.group.Name()
}

// Gets represents any Get request, including from peers
func (g *GroupStatistics) Gets() int64 {
	return g.group.Stats.Gets.Get()
}

// CacheHits represents either cache was good
func (g *GroupStatistics) CacheHits() int64 {
	return g.group.Stats.CacheHits.Get()
}

// GetFromPeersLatencyLower represents slowest duration to request value from peers
func (g *GroupStatistics) GetFromPeersLatencyLower() int64 {
	return g.group.Stats.GetFromPeersLatencyLower.Get()
}

// PeerLoads represents either remote load or remote cache hit (not an error)
func (g *GroupStatistics) PeerLoads() int64 {
	return g.group.Stats.PeerLoads.Get()
}

// PeerErrors represents a count of errors from peers
func (g *GroupStatistics) PeerErrors() int64 {
	return g.group.Stats.PeerErrors.Get()
}

// Loads represents (gets - cacheHits)
func (g *GroupStatistics) Loads() int64 {
	return g.group.Stats.Loads.Get()
}

// LoadsDeduped represents after singleflight
func (g *GroupStatistics) LoadsDeduped() int64 {
	return g.group.Stats.LoadsDeduped.Get()
}

// LocalLoads represents total good local loads
func (g *GroupStatistics) LocalLoads() int64 {
	return g.group.Stats.LocalLoads.Get()
}

// LocalLoadErrs represents total bad local loads
func (g *GroupStatistics) LocalLoadErrs() int64 {
	return g.group.Stats.LocalLoadErrs.Get()
}

// ServerRequests represents gets that came over the network from peers
func (g *GroupStatistics) ServerRequests() int64 {
	return g.group.Stats.ServerRequests.Get()
}

// MainCacheItems represents number of items in the main cache
func (g *GroupStatistics) MainCacheItems() int64 {
	return g.group.CacheStats(groupcache.MainCache).Items
}

// MainCacheBytes represents number of bytes in the main cache
func (g *GroupStatistics) MainCacheBytes() int64 {
	return g.group.CacheStats(groupcache.MainCache).Bytes
}

// MainCacheGets represents number of get requests in the main cache
func (g *GroupStatistics) MainCacheGets() int64 {
	return g.group.CacheStats(groupcache.MainCache).Gets
}

// MainCacheHits represents number of hit in the main cache
func (g *GroupStatistics) MainCacheHits() int64 {
	return g.group.CacheStats(groupcache.MainCache).Hits
}

// MainCacheEvictions represents number of evictions in the main cache
func (g *GroupStatistics) MainCacheEvictions() int64 {
	return g.group.CacheStats(groupcache.MainCache).Evictions
}

// HotCacheItems represents number of items in the main cache
func (g *GroupStatistics) HotCacheItems() int64 {
	return g.group.CacheStats(groupcache.HotCache).Items
}

// HotCacheBytes represents number of bytes in the hot cache
func (g *GroupStatistics) HotCacheBytes() int64 {
	return g.group.CacheStats(groupcache.HotCache).Bytes
}

// HotCacheGets represents number of get requests in the hot cache
func (g *GroupStatistics) HotCacheGets() int64 {
	return g.group.CacheStats(groupcache.HotCache).Gets
}

// HotCacheHits represents number of hit in the hot cache
func (g *GroupStatistics) HotCacheHits() int64 {
	return g.group.CacheStats(groupcache.HotCache).Hits
}

// HotCacheEvictions represents number of evictions in the hot cache
func (g *GroupStatistics) HotCacheEvictions() int64 {
	return g.group.CacheStats(groupcache.HotCache).Evictions
}
