# syntax=docker/dockerfile:1
# https://docs.docker.com/language/golang/build-images/
# docker build --tag nyanpasu-go-backend .
# docker run -d -p 8080:8080 --env-file ./.env --name happy_nyanpasu nyanpasu-go-backend
# docker stop happy_nyanpasu

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /nyanpasu-go-backend

EXPOSE 8080

CMD [ "/nyanpasu-go-backend" ]