FROM golang:1.7

RUN mkdir -p /go/src/beego-blog
WORKDIR /go/src/beego-blog
ENV GOPATH /go

# Install beego and the bee dev tool
RUN go get github.com/astaxie/beego && go get github.com/beego/bee && go get github.com/astaxie/beedb
RUN go get github.com/ziutek/mymysql/godrv
# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["bee", "run"]
