FROM docker.io/golang:1.15.6

MAINTAINER Jay <zhanchenjin.2008@163.com>


COPY beego_test /var/app/beego_test

WORKDIR /var/app/beego_test

RUN bash cicd/depend/depend.sh

ENTRYPOINT ["go", "run", "beego_test"]
