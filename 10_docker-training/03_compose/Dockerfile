FROM golang:1.13.3-alpine3.10
ENV APP_DIR /go/src/app
RUN mkdir -p ${APP_DIR}
WORKDIR ${APP_DIR}
COPY . .
RUN apk add --no-cache bash
RUN go build -o app && \
    chmod u+x app
ENTRYPOINT ["tail", "-f", "/dev/null"]