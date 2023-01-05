FROM golang
COPY getcpufreq.go .
RUN go build getcpufreq.go

FROM scratch
COPY --from=0 /go/getcpufreq .
CMD ["./getcpufreq"]