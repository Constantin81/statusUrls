FROM golang:1.16-alpine AS golang
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD [ "/app/main" ]


