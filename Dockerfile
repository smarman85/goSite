FROM smarman/alpine-base

ENV GOVERSION 1.13.3
ENV GOAPP goSite

RUN apk update && \
    apk add --no-cache \
        git \
        gcc \
        libc-dev \
        dumb-init && \
    wget -O go.tgz https://dl.google.com/go/go${GOVERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go.tgz

ENV GOPATH /go

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
    mkdir -p ${GOPATH}/src ${GOPATH}/bin

COPY ./ ${GOPATH}/src/${GOAPP}

WORKDIR ${GOPATH}/src/${GOAPP}

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/usr/local/go/bin/go", "run", "main.go"]
