FROM centos
RUN yum install -y  wget
MAINTAINER JeffJiang 
RUN yum install -y  gcc
RUN yum install -y  go
ENV GOROOT /usr/lib/golang
ENV PATH=$PATH:/usr/lib/golang/bin
RUN mkdir -p /root/gopath
RUN mkdir -p /root/gopath/src
RUN mkdir -p /root/gopath/pkg
RUN mkdir -p /root/gopath/bin
ENV GOPATH /root/gopath
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jackistom/gin_date/controllers
COPY data/* /root/gopath/src/
WORKDIR /root/gopath/src
RUN go run  main.go
CMD /root/gopath/src/main.go


