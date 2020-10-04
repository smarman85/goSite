FROM smarman/alpine-base AS builder

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

RUN cd ${GOPATH}/src/${GOAPP} && /usr/local/go/bin/go build -o $GOAPP
RUN echo "HELLO" && ls -al /${GOPATH}/src/${GOAPP}/

WORKDIR /${GOPATH}/src/${GOAPP}/
ENTRYPOINT ["/usr/bin/dumb-init", "--", "./goSite"]



##### Final image: (WORK IN PROGRESS - copy-from builder copies files but they aren't found)
##FROM smarman/alpine-base
##RUN apk add --no-cache dumb-init
###COPY --from=builder /usr/bin/dumb-init /usr/bin/dumb-init
##COPY --from=builder /go/src/goSite/ /app/goSite/
##COPY --from=builder /go/src/goSite/ /app/goSite/
##RUN echo "Final" && pwd && ls -al 
##RUN ls -al /app/goSite | grep goSite || echo "MISSING" && pwd
##ENTRYPOINT ["/usr/bin/dumb-init", "--", "/app/goSite/goSite"]

#CMD ["/usr/local/go/bin/go", "run", "main.go"]
