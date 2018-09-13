FROM docker.io/alpine:latest

WORKDIR /data
COPY asciinema-viewer /data
COPY conf /data/conf
COPY static /data/static
COPY views /data/views
COPY upload /data/upload
RUN chmod +x /data/asciinema-viewer

EXPOSE 8080

ENTRYPOINT ["/data/asciinema-viewer"]