FROM golang
COPY getcpufreq.go .
RUN go build getcpufreq.go

FROM scratch
COPY --from=0 /go/getcpufreq .
### Containers should NOT run as root as a good practice
### Refer to the following for more details
### https://docs.docker.com/engine/security/#linux-kernel-capabilities
### https://docs.bitnami.com/tutorials/why-non-root-containers-are-important-for-security
USER 10001
CMD ["./getcpufreq"]
