FROM golang:1.19

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN cd app/backend/cmd/main/ && CGO_ENABLED=0 go build -o ./app/backend/cmd/main

EXPOSE 8080

RUN chmod a+x ./app/backend/cmd/main
RUN chmod a+x ./app/backend/cmd/main/main.go

CMD ["go", "run", "./app/backend/cmd/main/main.go"]