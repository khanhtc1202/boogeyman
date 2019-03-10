FROM alpine
COPY ./bin/boogeyman-rest-linux-64 /usr/local/bin/boogeyman-rest
ENTRYPOINT ["/usr/local/bin/boogeyman-rest"]
