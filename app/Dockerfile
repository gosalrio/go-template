FROM golang:1.19-alpine as development

RUN apk update
RUN apk upgrade
RUN apk add build-base make git openssh

#set working space directory
WORKDIR /opt

#Copy files
COPY . .

WORKDIR /opt/app

#Install all dependencies
RUN go mod download

#Install air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

# Expose web-app API port
EXPOSE 7790

#Start App
CMD /opt/app/bin/air