FROM golang:1.14.3
ENV GO111MODULE "on"
ENV GOPROXY "https://goproxy.cn"
WORKDIR /library_service_v2
COPY . .
RUN make
EXPOSE 8080
CMD ["./main", "-c", "conf/config.yaml"]
