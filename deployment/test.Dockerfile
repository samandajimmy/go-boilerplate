FROM golang:1.17

# add golang env
ENV GOPRIVATE="github.com/samandajimmy/*"
ENV GOPROXY=""
ENV GONOSUMDB=""

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

# Compiling...
RUN make setup && make configure && make configure-ginkgo

# COPY the source code
COPY . .

# Run test
CMD [ "make", "testing" ]
