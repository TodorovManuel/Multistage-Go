FROM golang:alpine as builder
WORKDIR /opt/app
COPY . ./
RUN go build

FROM alpine
WORKDIR /opt/app
COPY --from=builder /opt/app/grupo9 /opt/app/grupo9
COPY --from=builder /opt/app/templates /opt/app/templates

EXPOSE 3000
CMD ["/opt/app/grupo9"]
