FROM golang:1.15 AS build
ARG GIT_COMMIT_HASH
RUN echo "Based on commit: $GIT_COMMIT_HASH"
COPY . /go/src/anomaly_detector
WORKDIR /go/src/anomaly_detector
RUN make build_docker hash="$GIT_COMMIT_HASH"

# 2 шаг - создание контейнера для запуска
FROM debian:buster-slim
WORKDIR /app
COPY --from=build /go/src/anomaly_detector/bin/* /app/
COPY --from=build /go/src/anomaly_detector/configs/* /app/configs/
COPY --from=build /go/src/anomaly_detector/scripts/* /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
RUN chmod +x /app/*
EXPOSE 31115/tcp
#CMD /app/gathering
ENTRYPOINT ["/app/gathering", "-config=/app/configs/common_docker.yml"]