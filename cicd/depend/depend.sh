#!/bin/bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
go build