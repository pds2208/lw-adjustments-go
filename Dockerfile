FROM golang:alpine3.10 as builder

LABEL stage=builder
WORKDIR /app

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN go build -gcflags "all=-N -l" -a -installsuffix -o /server .

# Final stage
FROM alpine:3.7

# Allow delve to run on Alpine based containers.
RUN apk add --no-cache libc6-compat

WORKDIR /

COPY --from=builder /server /
COPY --from=builder /app/config/config.lw.toml /app/config/

ENV CONFIG development
# to talk to postgres running on localhost
ENV DB_SERVER host.docker.internal

CMD [ "/server"]