# Fluentd

```sh
# bring up fluentd container
$ docker run -d \
    --name atai-fluentd \
    --network atai_logging_tracing \
    --ip "190.10.0.100" \
    --log-opt mode=non-blocking \
    --log-opt max-buffer-size=5m \
    --log-opt max-size=100m \
    --log-opt max-file=5 \
    fluent/fluentd:edge


# build Golang app by Bazel
$ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 \
    //examples/fluentd-logger:v0.0.0
$ bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 \
    //examples/fluentd-logger:v0.0.0

# connect to fluentd server at application level
$ docker run -it --rm \
    --name atai-fluentd-logging-v1 \
    --network atai_logging_tracing \
    --ip "190.10.0.50" \
    --log-opt mode=non-blocking \
    --log-opt max-buffer-size=5m \
    --log-opt max-size=100m \
    --log-opt max-file=5 \
    alantai/prj-logging-tracing-v1/examples/fluentd-logger:v0.0.0

# connect to fluentd server at container level
$ docker run -it --rm \
    --name atai-fluentd-logging-v1 \
    --network atai_logging_tracing \
    --ip "190.10.0.50" \
    --log-driver=fluentd \
    --log-opt fluentd-address=190.10.0.100:24224 \
    alantai/prj-logging-tracing-v1/examples/fluentd-logger:v0.0.0

# for the fluentd container based on alpine distribution, the logs are stored under fluentd/log
$ docker exec -it atai-fluentd sh

$ ls fluentd/log/
# data.b5d207acfb1cbcb9bb1e20094fb9318e5.log       data.b5d207acfb1cbcb9bb1e20094fb9318e5.log.meta  data.log

$ cat fluentd/log/data.b5d207acfb1cbcb9bb1e20094fb9318e5.log
# ...
# 2021-11-30T20:48:12+00:00	d920f9916ee9	{"container_id":"d920f9916ee958a45812f7ea99fbbfe6fd7df1f797a1b7b2aae2f701ee756d31","container_name":"/atai-fluentd-logging-v1","source":"stdout","log":"Start running fluent
# d logging example...\r"}
# 2021-11-30T20:48:12+00:00	d920f9916ee9	{"log":"2021/11/30 20:48:12 <=== 0-th log ===>\r","container_id":"d920f9916ee958a45812f7ea99fbbfe6fd7df1f797a1b7b2aae2f701ee756d31","container_name":"/atai-fluentd-logging-
# v1","source":"stdout"}
# 2021-11-30T20:48:12+00:00	d920f9916ee9	{"log":"2021/11/30 20:48:12 Successfully wrote log to Fluentd\r","container_id":"d920f9916ee958a45812f7ea99fbbfe6fd7df1f797a1b7b2aae2f701ee756d31","container_name":"/atai-f
# luentd-logging-v1","source":"stdout"}
# 2021-11-30T20:48:13+00:00	d920f9916ee9	{"container_id":"d920f9916ee958a45812f7ea99fbbfe6fd7df1f797a1b7b2aae2f701ee756d31","container_name":"/atai-fluentd-logging-v1","source":"stdout","log":"2021/11/30 20:48:13 
# <=== 1-th log ===>\r"}
# 2021-11-30T20:48:13+00:00	d920f9916ee9	{"container_name":"/atai-fluentd-logging-v1","source":"stdout","log":"2021/11/30 20:48:13 Successfully wrote log to Fluentd\r","container_id":"d920f9916ee958a45812f7ea99fbb
# fe6fd7df1f797a1b7b2aae2f701ee756d31"}
# 2021-11-30T20:48:14+00:00	d920f9916ee9	{"container_name":"/atai-fluentd-logging-v1","source":"stdout","log":"2021/11/30 20:48:14 <=== 2-th log ===>\r","container_id":"d920f9916ee958a45812f7ea99fbbfe6fd7df1f797a1
# b7b2aae2f701ee756d31"}
# 2021-11-30T20:48:14+00:00	d920f9916ee9	{"container_id":"d920f9916ee958a45812f7ea99fbbfe6fd7df1f797a1b7b2aae2f701ee756d31","container_name":"/atai-fluentd-logging-v1","source":"stdout","log":"2021/11/30 20:48:14 
# Successfully wrote log to Fluentd\r"}
# 2021-11-30T20:48:14+00:00	d920f9916ee9	{"container_id":"d920f9916ee958a45812f7ea99fbbfe6fd7df1f797a1b7b2aae2f701ee756d31","container_name":"/atai-fluentd-logging-v1","source":"stdout","log":"[2021-11-30T20:48:14
# Z] Unable to send logs to fluentd, reconnecting...\r"}
# 2021-11-30T20:48:15+00:00	d920f9916ee9	{"container_id":"d920f9916ee958a45812f7ea99fbbfe6fd7df1f797a1b7b2aae2f701ee756d31","container_name":"/atai-fluentd-logging-v1","source":"stdout","log":"2021/11/30 20:48:15 
# <=== 3-th log ===>\r"}
# ...
```

Ref:
- [Fluent logger Golang](https://github.com/fluent/fluent-logger-golang)
- [Fluentd logging driver](https://docs.docker.com/config/containers/logging/fluentd/)
- [Config: Transport Section](https://docs.fluentd.org/configuration/transport-section)
- [Docker fluent/fluentd](https://hub.docker.com/r/fluent/fluentd/)
- [Fluentd Trouble Shooting](https://docs.fluentd.org/deployment/trouble-shooting)
- [Forward protocol of Fluentd & FluentBit](https://github.com/fluent/fluent-bit-docs/blob/master/pipeline/outputs/forward.md)

Issues:
- [Issue: ForceStopAsyncSend option for graceful stop in async mode](https://github.com/fluent/fluent-logger-golang/pull/77)