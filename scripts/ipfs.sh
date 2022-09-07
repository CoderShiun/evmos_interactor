#!/bin/sh

docker pull ipfs/go-ipfs

mkdir -p /data/ipfs/ipfs_staging
mkdir -p /data/ipfs/ipfs_data
export ipfs_staging=/data/ipfs/ipfs_staging
export ipfs_data=/data/ipfs/ipfs_data

docker run -d --name ipfs_host \
		   -v $ipfs_staging:/export \
		   -v $ipfs_data:/data/ipfs \
		   -p 4001:4001 -p 127.0.0.1:8080:8080 \
		   -p 127.0.0.1:5001:5001 ipfs/go-ipfs:latest