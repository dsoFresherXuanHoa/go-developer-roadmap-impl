FROM golang:1.20

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /gin-restful-gorm
CMD [ "./gin-restful-gorm" ]