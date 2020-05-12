FROM smarman/alpine-base

RUN apk update && \
    apk add go \
        git \
        gcc \
        libc-dev && \
    go get -u github.com/gorilla/mux

COPY ./ /srv

WORKDIR /srv

CMD ["/bin/sh"]
#ENTRYPOINT ["go", "run", "main.go"]
