FROM golang:1.13.3-alpine3.10
ENV APP_DIR ${GOPATH}/src/app
WORKDIR ${APP_DIR}
COPY . .
RUN apk add --no-cache bash
RUN go build -o app && \
    chmod u+x app
CMD ./app