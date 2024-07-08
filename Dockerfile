FROM golang:1.22.5-alpine as build-stage

WORKDIR /app

COPY . /app
RUN go mod download

COPY . /app 

RUN CGO_ENABLED=0 GOOS=linux go build -o /interview-service

FROM gcr.io/distroless/static-debian12

COPY --from=build-stage /dining-table-service .

EXPOSE 8080

CMD ["/dining-table-service"]