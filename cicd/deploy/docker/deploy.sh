#!/bin/bash
cd /var/app
docker build -t beegi_test:1.0.0 -f /var/app/beego_test/cicd/deploy/docker/Dockerfile .