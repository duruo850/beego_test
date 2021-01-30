#!/bin/bash
cd /var/app
docker build -t beegi_test -f /var/app/beego_test/cicd/deploy/docker/Dockerfile .