# -----------
# Build Stage
# -----------
FROM artifactory.pegadaian.co.id:8084/golang:1.17 as builder

# add golang env
ENV GOPRIVATE="repo.pegadaian.co.id,artifactory.pegadaian.co.id/repository/go-group-01"
ENV GOPROXY="https://artifactory.pegadaian.co.id/repository/go-group-01/"
ENV GONOSUMDB="github.com/*,golang.org/*,gopkg.in/*,gitlab.com/*,cloud.google.com/*,go.*,google.golang.org/*,gotest.*,honnef.co/*,mellium.im/*"

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
COPY .git ./.git
COPY script ./script
COPY config ./config

# download package dependencies
RUN make setup && make configure && make configure-swag

# copy source code
COPY cmd ./cmd
COPY internal ./internal
COPY docs ./docs
COPY migration ./migration

# Compiling...
RUN make release-dev

# ------------
# Deploy Stage
# ------------
FROM artifactory.pegadaian.co.id:8084/alpine:3.15

ARG ARG_PORT=3000

WORKDIR /usr/src/app

RUN apk add --no-cache tzdata ca-certificates

COPY --from=builder /usr/src/app/bin/debug /usr/src/app
COPY --from=builder /usr/src/app/docs /usr/src/app/docs
COPY --from=builder /usr/src/app/migration /usr/src/app/migration

EXPOSE ${ARG_PORT}

ENTRYPOINT ["./go-boiler-plate", "-dir=/usr/src/app"]
