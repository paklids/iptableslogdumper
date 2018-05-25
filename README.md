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

### How do I use this?

Check the variables set within the `docker-compose.yml` and if they fit your needs then run:

```docker-compose build```

```docker-compose up -d```

and when the run is complete you can tear it all down:

```docker-compose down```

A few of the variables are:

      LogsPerSecond - The number of logs that you want to generate per second
      TotalLogs - The number of logs you want to generate for this run
      MyProgramName - If you want to customize the name of the application running
      StartDelayInSeconds - The delay before the log generator start creating log files
      (allows the pipeline some time to initiate)

### Who has helped with this so far?

Flowroute ( https://www.flowroute.com/ )