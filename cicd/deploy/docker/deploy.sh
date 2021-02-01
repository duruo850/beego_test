#!/bin/bash
cd /var/app
docker build -t beegi_test:1.0.0 -f /var/app/beego_test/cicd/deploy/docker/Dockerfile .

# sudo docker run -it --network default_network -p 0.0.0.0:8900:8900 beegi_test:1.0.0 /bin/bash