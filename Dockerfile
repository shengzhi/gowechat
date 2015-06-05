FROM google/golang
MAINTAINER jong "gongshengzhi@gmail.com"

# Build app
WORKDIR /gopath/app/src
ENV GOPATH /gopath/app
ADD . /gopath/app/src

# RUN go get github.com/shengzhi/gowechat
RUN go install gowechat

EXPOSE 80
CMD ["/gopath/app/bin/gowechat"]
