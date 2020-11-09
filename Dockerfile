FROM golang:1.10-alpine as build

WORKDIR /root/k8s-device-plugin-scheduler
COPY . .

RUN go build -o /go/bin/k8s-device-plugin-scheduler cmd/*.go

FROM alpine

COPY --from=build /go/bin/k8s-device-plugin-scheduler /usr/bin/k8s-device-plugin-scheduler

CMD ["k8s-device-plugin-scheduler"]
