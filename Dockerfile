FROM golang:1.15 as build

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o figApi

FROM debian:buster

WORKDIR /app
COPY --from=build /app/figApi /app 

CMD ["./figApi"]
