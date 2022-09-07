FROM golang:1.19 as dev
ENV HELIXF_ROOT=/go/src/app
COPY docker/conf/air.toml $HELIXF_ROOT/air.toml
COPY src/scheduler/go.mod src/scheduler/go.sum $HELIXF_ROOT/src/scheduler/
COPY lib/go.mod lib/go.sum $HELIXF_ROOT/lib/
WORKDIR $HELIXF_ROOT/src/scheduler
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", "../../air.toml"]
# CMD ["sleep", "7200"]
