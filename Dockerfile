FROM golang:1.22 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./

COPY .env ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weather-api .

FROM scratch

WORKDIR /app

COPY --from=build /app/weather-api .

COPY --from=build /app/.env ./

ENTRYPOINT ["./weather-api"]
