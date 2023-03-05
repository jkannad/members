FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY cmd/web/*.go ./
RUN go build -o /docker-spaa-members
EXPOSE 8080
CMD ["/docker-spaa-members"]