#! /bin/sh

# Make sure the script fails fast.
set -e
set -u
set -x

PROTO_DIR=groupcachepb

protoc --proto_path=$PROTO_DIR \
  --go_opt=paths=source_relative \
  --go_out=$PROTO_DIR \
  groupcache.proto

protoc --proto_path=example \
  --go_opt=paths=source_relative \
  --go_out=example \
  example.proto

protoc --proto_path=testpb \
  --go_opt=paths=source_relative \
  --go_out=testpb \
  test.proto

