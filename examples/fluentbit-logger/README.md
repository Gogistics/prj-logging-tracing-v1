# FluentBit

```sh
# send CPU metrics to fluentd which is brought up in the exmaple, fluentd-logger.
$ docker run -d \
    --name atai-fluentbit \
    --network atai_logging_tracing \
    --ip "190.10.0.110" \
    fluent/fluent-bit:1.8 \
    /fluent-bit/bin/fluent-bit -i cpu -t atai_fluent_bit -o forward://190.10.0.100:24224
# ...
# 2021-11-30T23:43:38+00:00	atai_fluent_bit	{"cpu_p":0.5,"user_p":0.0,"system_p":0.5,"cpu0.p_cpu":2.0,"cpu0.p_user":0.0,"cpu0.p_system":2.0,"cpu1.p_cpu":1.0,"cpu1.p_user":1.0,"cpu1.p_system":0.0,"cpu2.p_cpu":0.0,"cpu2.p_user":0.0,"cpu2.p_system":0.0,"cpu3.p_cpu":0.0,"cpu3.p_user":0.0,"cpu3.p_system":0.0,"cpu4.p_cpu":0.0,"cpu4.p_user":0.0,"cpu4.p_system":0.0,"cpu5.p_cpu":1.0,"cpu5.p_user":0.0,"cpu5.p_system":1.0}
# 2021-11-30T23:43:39+00:00	atai_fluent_bit	{"cpu_p":0.3333333333333333,"user_p":0.16666666666666666,"system_p":0.16666666666666666,"cpu0.p_cpu":1.0,"cpu0.p_user":1.0,"cpu0.p_system":0.0,"cpu1.p_cpu":1.0,"cpu1.p_user":0.0,"cpu1.p_system":1.0,"cpu2.p_cpu":0.0,"cpu2.p_user":0.0,"cpu2.p_system":0.0,"cpu3.p_cpu":0.0,"cpu3.p_user":0.0,"cpu3.p_system":0.0,"cpu4.p_cpu":0.0,"cpu4.p_user":0.0,"cpu4.p_system":0.0,"cpu5.p_cpu":1.0,"cpu5.p_user":0.0,"cpu5.p_system":1.0}
# 2021-11-30T23:43:40+00:00	atai_fluent_bit	{"cpu_p":0.5,"user_p":0.0,"system_p":0.5,"cpu0.p_cpu":1.0,"cpu0.p_user":0.0,"cpu0.p_system":1.0,"cpu1.p_cpu":0.0,"cpu1.p_user":0.0,"cpu1.p_system":0.0,"cpu2.p_cpu":0.0,"cpu2.p_user":0.0,"cpu2.p_system":0.0,"cpu3.p_cpu":1.0,"cpu3.p_user":0.0,"cpu3.p_system":1.0,"cpu4.p_cpu":1.0,"cpu4.p_user":0.0,"cpu4.p_system":1.0,"cpu5.p_cpu":0.0,"cpu5.p_user":0.0,"cpu5.p_system":0.0}
# 2021-11-30T23:43:41+00:00	atai_fluent_bit	{"cpu_p":0.6666666666666666,"user_p":0.16666666666666666,"system_p":0.5,"cpu0.p_cpu":0.0,"cpu0.p_user":0.0,"cpu0.p_system":0.0,"cpu1.p_cpu":1.0,"cpu1.p_user":0.0,"cpu1.p_system":1.0,"cpu2.p_cpu":1.0,"cpu2.p_user":0.0,"cpu2.p_system":1.0,"cpu3.p_cpu":1.0,"cpu3.p_user":1.0,"cpu3.p_system":0.0,"cpu4.p_cpu":0.0,"cpu4.p_user":0.0,"cpu4.p_system":0.0,"cpu5.p_cpu":2.0,"cpu5.p_user":1.0,"cpu5.p_system":1.0}
# 2021-11-30T23:43:42+00:00	atai_fluent_bit	{"cpu_p":0.16666666666666666,"user_p":0.0,"system_p":0.16666666666666666,"cpu0.p_cpu":1.0,"cpu0.p_user":0.0,"cpu0.p_system":1.0,"cpu1.p_cpu":0.0,"cpu1.p_user":0.0,"cpu1.p_system":0.0,"cpu2.p_cpu":0.0,"cpu2.p_user":0.0,"cpu2.p_system":0.0,"cpu3.p_cpu":0.0,"cpu3.p_user":0.0,"cpu3.p_system":0.0,"cpu4.p_cpu":0.0,"cpu4.p_user":0.0,"cpu4.p_system":0.0,"cpu5.p_cpu":1.0,"cpu5.p_user":0.0,"cpu5.p_system":1.0}
# ...
```

Ref:
- [FluentBit Hands On! 101](https://docs.fluentbit.io/manual/stream-processing/getting-started/hands-on)
- [fluent/fluent-bit](https://hub.docker.com/r/fluent/fluent-bit/)
- [Forward protocol of Fluentd & FluentBit](https://github.com/fluent/fluent-bit-docs/blob/master/pipeline/outputs/forward.md)
