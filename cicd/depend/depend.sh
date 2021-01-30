#!/bin/bash
apt-get update
apt-get install -y vim mysql-client
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
go build