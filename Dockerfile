FROM alpine
ENV APP_BIN /usr/local/bin/boogeyman-rest
COPY ./bin/boogeyman-rest-linux-64 ${APP_BIN}
ENTRYPOINT "$APP_BIN"
