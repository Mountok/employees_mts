FROM golang:latest

COPY ./ ./
RUN go build -o emp_api cmd/main.go
CMD [ "./emp_api" ]
