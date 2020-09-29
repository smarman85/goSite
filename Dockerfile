FROM smarman/alpine-base

RUN apk update && \
    apk add go \
        git \
        gcc \
        libc-dev \
        dumb-init && \
    go get -u github.com/gorilla/mux

COPY ./ /srv

WORKDIR /srv

#CMD ["/bin/sh"]
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["go", "run", "main.go"]
