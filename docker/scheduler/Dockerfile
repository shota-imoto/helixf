FROM golang:1.19 as dev
ENV HELIXF_ROOT=/go/src/app
COPY docker/conf/air.toml $HELIXF_ROOT/air.toml
COPY src/scheduler $HELIXF_ROOT/src/scheduler
COPY lib $HELIXF_ROOT/lib
WORKDIR $HELIXF_ROOT/src/scheduler
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", "../../air.toml"]
# CMD ["go", "run", "main.go"]
# CMD ["sleep", "7200"]
