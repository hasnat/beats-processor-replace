ARG FILEBEAT_VERSION=6.4.2

FROM golang:1.10.3 as builder
#ENV GO111MODULE on
ARG FILEBEAT_VERSION

WORKDIR /go/src/github.com/hasnat/beats-processor-replace

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep init

#RUN go fmt ./...
COPY . .
RUN dep ensure -add github.com/elastic/beats@${FILEBEAT_VERSION}

# golang:1.11.1
#RUN go get github.com/elastic/beats@v${FILEBEAT_VERSION}
#RUN go mod init github.com/hasnat/beats-processor-replace
#RUN go mod edit -require github.com/elastic/beats@v${FILEBEAT_VERSION}


RUN go test
RUN go build -buildmode=plugin -o processor-replace-linux.so



FROM docker.elastic.co/beats/filebeat:${FILEBEAT_VERSION}
USER root
RUN mkdir -p /usr/local/plugins/
COPY --from=builder /go/src/github.com/hasnat/beats-processor-replace/processor-replace-linux.so /usr/local/plugins/
RUN chown -R filebeat:filebeat /usr/local/plugins/
USER filebeat

CMD ["/bin/sh", "-c", "/usr/local/bin/docker-entrypoint -e --plugin /usr/local/plugins/processor-replace-linux.so"]
