## IPTables Log Dumper pipeline

### What Does this do?

The GoLang snippet, running within a docker container, generates fake iptables logs 
at a rate that you specify.

The docker-compose pipeline looks like this:

log_generator(tcp_syslog) --> logstash_pre --> redis_queue --> logstash_post(processor) --> output

The redis queue allows for bursts in traffic so that processing is not a blocking event 
in the pipeline

### Why build this?

This allows anyone to see how many events their logstash rules can process per second and
benchmark those results.

In my example, logstash is also doing some basic geoip lookups, which could be useful
in shaping firewall decisions.