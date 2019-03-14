FROM alpine
COPY ./bin/boogeyman-rest-linux-64 /usr/local/bin/boogeyman-rest
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true
ENTRYPOINT ["/usr/local/bin/boogeyman-rest"]
