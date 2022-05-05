# -----------
# Build Stage
# -----------
FROM artifactory.pegadaian.co.id:8084/golang:1.17 as builder

# add ssl certificate
ADD data/ssl_certificate.crt /usr/local/share/ca-certificates/ssl_certificate.crt
RUN chmod 644 /usr/local/share/ca-certificates/ssl_certificate.crt && update-ca-certificates

# copy netrc configuration
COPY config/.netrc /root/.netrc
RUN chmod 600 /root/.netrc

# Install some dependencies needed to build the project
RUN apt install git make 

# Here we copy the rest of the source code
RUN mkdir /usr/src/app/
WORKDIR /usr/src/app

# COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .
COPY Makefile .
COPY .env .
COPY .git ./.git
COPY script ./script
COPY config ./config

# copy source code
COPY cmd ./cmd
COPY internal ./internal

# Compiling...
RUN make setup && make release

# ------------
# Deploy Stage
# ------------
FROM artifactory.pegadaian.co.id:8084/alpine:3.15

ARG ARG_PORT=3000

WORKDIR /usr/src/app

RUN apk add --no-cache tzdata ca-certificates

COPY --from=builder /usr/src/app/bin/release /usr/src/app
COPY migration /usr/src/app/migration

EXPOSE ${ARG_PORT}

ENTRYPOINT ["./go-boiler-plate"]
