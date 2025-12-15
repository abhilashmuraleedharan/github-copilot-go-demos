# Demo 10 – Plan Mode Prompt for 5G CDR Pipeline Architecture

## Primary Prompt (/plan)

```
/plan

I want to build a Go microservice called `cdr-pipeline` for processing
5G CDRs. It should ingest Voice, Data, SMS, and Event CDRs from
Kafka, HTTP, and SFTP. It must support enrichment using subscriber DB,
location metadata, and service-type inference.

The service must apply mapping/transformation rules based on CDR type
and output processed CDRs into a Kafka billing topic and archival storage.

Peak throughput: 150K CDRs/sec.
Pipeline must support pluggable stages and config-driven rule selection.

Please produce a detailed execution plan with:

- Architecture proposal
- Module breakdown
- Sequence of implementation tasks
- Concurrency and backpressure strategy
- Testing approach
- CI/CD + Dockerfile + Helm outline
- Observability plan (metrics, logs, tracing)

Also suggest a suitable Go folder hierarchy.
```

## Follow-up Prompt 1
```
/follow-up
Expand the concurrency model. Explain batching, worker pools,
channel buffering, backpressure, and shutdown behavior.
```

## Follow-up Prompt 2
```
/follow-up
Generate interface stubs for ingestion, enrichment,
transformation, and export stages.
```

## Follow-up Prompt 3
```
/follow-up
Propose a load-testing and replay-testing approach for CDR streams.
```
