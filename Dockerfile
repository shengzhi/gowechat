FROM google/golang
MAINTAINER jong "gongshengzhi@gmail.com"

# Build app
RUN mkdir -p /gopath/app/src
ENV GOPATH /gopath/app
COPY ./gowechat /gopath/app/src

# RUN go get github.com/shengzhi/gowechat
RUN go install gowechat

EXPOSE 80
CMD ["/gopath/app/bin/gowechat"]
