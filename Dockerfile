FROM golang:latest
RUN mkdir /code
WORKDIR /code
COPY . /code/
COPY go.mod go.sum ./
RUN go mod tidy