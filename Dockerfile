FROM google/golang
MAINTAINER jong "gongshengzhi@gmail.com"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

RUN go get github.com/shengzhi/gowechat
RUN go install github.com/shengzhi/gowechat

EXPOSE 80
CMD ["/gopath/app/bin/gowechat"]
