# choose a base golang image
FROM golang:1.13.0

# add maintainer label for LTS
LABEL maintainer="Niraj Georgian <nirajgeorgian@gmail.com>"

# define system module and update them
RUN apt-get update
RUN go version

# setup local project and it's dependenciew
ENV PROJECT github.com/nirajgeorgian/job
WORKDIR /go/src/$PROJECT

# copy makefile to execute make commands
COPY Makefile ./

RUN make setup-dep
RUN dep version

# install dependencies
COPY Gopkg.* ./
RUN dep ensure -vendor-only -v

COPY . .

# build the application and install it
RUN make build
RUN go install /go/src/github.com/nirajgeorgian/job

# test job version
RUN job version

# expose and run the main application
EXPOSE 3000
CMD ["job", "serve"]
