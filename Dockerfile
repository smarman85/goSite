FROM golang:1.16.7-alpine3.14

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
    mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
    mkdir -p ${GOPATH}/src ${GOPATH}/bin

COPY . ${GOPATH}/src/${GOAPP}/

RUN cd ${GOPATH}/src/${GOAPP} && go build -o $GOAPP
# RUN CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o goSite .
USER ${USER}
WORKDIR ${GOPATH}/src/${GOAPP}
# CMD ["tail" "/dev/null"]
ENTRYPOINT ["/usr/bin/dumb-init", "--", "/go/src/goSite/goSite"]
#ENTRYPOINT ["./init/app"]



##### Final image: (WORK IN PROGRESS - copy-from builder copies files but they aren't found)
#FROM smarman/alpine-base
##COPY --from=builder /usr/bin/dumb-init /usr/bin/dumb-init
##COPY --from=builder /go/src/goSite /app/goSite
##RUN echo "Final" && pwd && ls -al 
##RUN ls -al /app/goSite | grep goSite || echo "MISSING" && pwd
##ENTRYPOINT ["/usr/bin/dumb-init", "--", "/app/goSite/goSite"]

#CMD ["/usr/local/go/bin/go", "run", "main.go"]
