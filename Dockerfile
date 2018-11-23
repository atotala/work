FROM alpine:3.8 

RUN apk --no-cache add ca-certificates

WORKDIR /opt/dellemc

COPY app .
RUN chmod +x /opt/dellemc/*

CMD ["./app"]
