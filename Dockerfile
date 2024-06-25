FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /ecommerce-app

EXPOSE 8080

CMD [ "/ecommerce-app" ]