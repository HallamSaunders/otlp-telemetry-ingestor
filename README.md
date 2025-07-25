## Use Case
This is a simple program to unmarshal `.jsonl` files containing [https://opentelemetry.io/docs/specs/otlp/](OTLP)-formatted telemetry data outputted by the [https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/protocol/file-exporter.md](OpenTelemetry Collector's File Exporter) into flat SQLite files.

This may be useful in those situations where a program emitting (potentially sensitive) telemetry data resides on an airgapped (or otherwise unable to communicate this telemetry data externally) machine. The telemetry can be written to a file rather than streamed over gRPC/HTTP as is usually done.
In such situations, the telemetry within the file may need to be read and analysed at a later date. There is currently no good way to do this. Using this tool, the file can be serialized into a standardised database format to make it more usable.
Once in SQLite format, you could (for example) use something like the [https://grafana.com/grafana/plugins/frser-sqlite-datasource/](Grafana SQLite Plugin) to view telemetry data on a dashboard.
