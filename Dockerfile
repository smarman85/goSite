FROM smarman/alpine-base AS builder

ENV GOVERSION 1.13.3
ENV GOAPP goSite
ENV GOPATH /go
ENV USER gosite

RUN apk update && \
    apk add --no-cache \
        git \
        gcc \
        libc-dev \
        dumb-init && \
    adduser -D ${USER} && \
    wget -O go.tgz https://dl.google.com/go/go${GOVERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go.tgz && \
    mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
    mkdir -p ${GOPATH}/src ${GOPATH}/bin

COPY . ${GOPATH}/src/${GOAPP}/

# AMD64
RUN cd ${GOPATH}/src/${GOAPP} && /usr/local/go/bin/go build -o $GOAPP
# RASPBERRY PI
#RUN cd ${GOPATH}/src/${GOAPP} && env GOOS=linux GOARCH=arm GOARM=5 /usr/local/go/bin/go build -o $GOAPP
#RUN echo "HELLO" && ls -al /${GOPATH}/src/${GOAPP}/

USER ${USER}
WORKDIR /${GOPATH}/src/${GOAPP}/
#ENTRYPOINT ["/usr/bin/dumb-init", "--", "./goSite"]
ENTRYPOINT ["./init/app"]



##### Final image: (WORK IN PROGRESS - copy-from builder copies files but they aren't found)
#FROM smarman/alpine-base
##COPY --from=builder /usr/bin/dumb-init /usr/bin/dumb-init
##COPY --from=builder /go/src/goSite /app/goSite
##RUN echo "Final" && pwd && ls -al 
##RUN ls -al /app/goSite | grep goSite || echo "MISSING" && pwd
##ENTRYPOINT ["/usr/bin/dumb-init", "--", "/app/goSite/goSite"]

#CMD ["/usr/local/go/bin/go", "run", "main.go"]
