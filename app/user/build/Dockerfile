#build stage
FROM golang:1.18-alpine AS builder
WORKDIR /app

COPY . . 

RUN go mod download

EXPOSE 8080

RUN GOOS=linux go build -o main.app ./app/user/cmd 


#final stage

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata
COPY --from=builder app/main.app /app/main.app

EXPOSE 8080

ENTRYPOINT [ "/app/main.app" ]


