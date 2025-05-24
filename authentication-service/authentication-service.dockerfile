FROM alpine:3.18

RUN mkdir /app

COPY authApp /app

CMD ["/app/authApp"]