# stage 1: Build
FROM golang:1.24.5-alpine3.21  AS builder


WORKDIR /fakeAPI

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /fakeAPI/app .

#stage 2: Runtime
FROM alpine:3.21

RUN apk add --no-cache ca-certificates

WORKDIR /root/
ENV PORT=3000
COPY --from=builder /fakeAPI/app .

EXPOSE 3000
CMD [ "./app" ]