#!/bin/bash


build() {
    echo "build..."
    cd /var/app
    sudo docker build -t beegi_test:1.0.0 -f /var/app/beego_test/cicd/deploy/docker/Dockerfile .
}

start() {
    echo "start..."
    sudo docker run --name beege_test -it --network default_network -d --restart=always -p 0.0.0.0:8900:8900 beegi_test:1.0.0
}


cmd=$1
case $cmd in
build)
        build
;;
start)
        start
;;
esac
exit 0