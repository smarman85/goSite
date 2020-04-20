FROM smarman/alpine-base

RUN apk update && \
    apk add go \
        git \
        gcc \
        libc-dev && \
    go get -u github.com/gorilla/mux

COPY ./ /srv

CMD ["/bin/sh"]
