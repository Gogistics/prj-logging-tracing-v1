# Logging & Telemetry (WIP)

Options:
* fluentd
    * send to fluentd directly from applications
    * collect logs via Docker logging driver
* OpenTelemetry
    * collect logs/metrics directly from applications
* influxdb


Ref:
- [How to collect, standardize, and centralize Golang logs](https://www.datadoghq.com/blog/go-logging/)
- [Fluentd logging driver](https://docs.docker.com/config/containers/logging/fluentd/)
- [Log and trace with InfluxDB](https://docs.influxdata.com/influxdb/v1.8/administration/logs/)
- [OpenTelemetry Collector](https://github.com/open-telemetry/opentelemetry-go/blob/main/example/otel-collector/main.go)
- [OpenTelemetry NewExporter](https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp#example-NewExporter)
- [Understanding Envoy Proxy HTTP Access Logs](https://blog.getambassador.io/understanding-envoy-proxy-and-ambassador-http-access-logs-fee7802a2ec5)
- [Engarde : Parse envoy and istio-proxy access logs like a champ](https://medium.com/@nitishmalhotra/engarde-parse-envoy-and-istio-proxy-logs-like-a-champ-faec31c563e7)
- [11 Open Source Log Collectors for Centralized Logging](https://geekflare.com/open-source-centralized-logging/)
- [Experiment: Introducing the FluentBit exporter for OpenTelemetry](https://medium.com/opentelemetry/introducing-the-fluentbit-exporter-for-opentelemetry-574ec133b4b4)
- [A simple golang web server with basic logging, tracing, health check, and graceful shutdown](https://gist.github.com/enricofoltran/10b4a980cd07cb02836f70a4ab3e72d7)



## Examples
1. Set up Goalng module for all Golang applicationss
```sh
$ go mod init github.com/Gogistics/prj-logging-tracing-v1
```

2. Set up build environment with Bazel
2-1. Create .bazelversion, .bazelrc, and utils/scripts/bazel_stamp.sh

2-2. Create WORKSPACE, BUILD.bazel, and deps.bzl at root directory.


### Example 1. Application based on OpenTelemetery examples

### Example 2. Web application
* Build web app by Caddy.
* Use fluentd as the log collector
* Export data/logs by OpenTelemetry

Ref:
- [OpenTelemetry: Getting Started with Golang](https://opentelemetry.io/docs/go/getting-started/)
- [Caddy server](https://github.com/caddyserver/caddy)