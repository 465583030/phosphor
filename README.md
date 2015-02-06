![](docs/logo.png)

Phosphor is a distributed tracing system, in the same vein as Twitter's Zipkin and Google's Dapper.

It is comprised of a few simple components:

 - Phosphor Client, used to send traces from applications
 - Phosphor Daemon, collects traces and forwards onto the main server
 - Phosphor Server, stores traces and aggregated trace information
 - Phosphor UI, view trace and debug information about your infrastructure

![Phosphor Architecture](docs/phosphor/outline.png)

## Dependencies

 - NSQ is currently used as the delivery transport for traces between PhosphorD and the Phosphor Server
