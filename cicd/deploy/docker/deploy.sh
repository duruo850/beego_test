#!/bin/bash
cd /var/app/beegi_test
docker build -t beegi_test -f /var/app/beegi_test/cicd/deploy/docker/Dockerfile .