FROM docker.io/alpine:latest

WORKDIR /data
COPY asciinema-viewer /data
COPY conf /data/conf
COPY static /data/static
COPY views /data/views
COPY upload /data/upload

EXPOSE 8080

ENTRYPOINT ["/data/asciinema-viewer"]