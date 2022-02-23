FROM golang:1.13.15

# ENV GO111MODULE=on

WORKDIR /app
COPY . .

# swag
# RUN wget https://github.com/swaggo/swag/releases/download/v1.6.3/swag_1.6.3_Linux_x86_64.tar.gz -O /tmp/swag_1.6.3_Linux_x86_64.tar.gz
# RUN cd /tmp && \
#     tar zxvf /tmp/swag_1.6.3_Linux_x86_64.tar.gz
# RUN cp /tmp/swag /go/bin/
# RUN swag init

RUN go mod download
RUN go build

# goose - DB migration tool
RUN go get github.com/pressly/goose/cmd/goose

# air - live reload utility
RUN go get -u github.com/cosmtrek/air@v1.27.8

EXPOSE 3000