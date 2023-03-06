FROM golang:1.19-alpine

RUN mkdir /ALTA-Dashboard-BE

WORKDIR /ALTA-Dashboard-BE

COPY ./ /ALTA-Dashboard-BE

RUN go mod tidy

RUN go build -o alta-dashboard-be

CMD ["./alta-dashboard-be"]