FROM golang:1.13.3-alpine3.10
ENV APP_DIR $GOPATH/src/hello_world
WORKDIR $APP_DIR
COPY . .
RUN cd $APP_DIR/app && \
    go build -o app && \
    chmod u+x app
RUN apk add --no-cache bash
# ENV ARGS=""
# CMD ./app $ARGS
# ENTRYPOINT ["hello2"]
CMD [ "sh" ]