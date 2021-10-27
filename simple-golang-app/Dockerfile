FROM golang:1.15-alpine as builder

RUN go env -w GOPROXY="https://repo.snapp.tech/repository/goproxy/"
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /app/main


#################
FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/main
ENTRYPOINT ["/app/main"]
