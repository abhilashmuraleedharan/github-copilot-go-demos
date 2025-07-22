# CDR Transformation Microservice Specification

## Overview
The CDR Transformation Service is a stateless microservice responsible for ingesting, enriching, transforming, and routing Call Data Records (CDRs) from multiple telecom sources, as part of a 5G Mediation Layer. The service must be highly extensible and performant, supporting both modern (5G) and legacy (3G/4G) feeds.

## Functional Requirements

- **Multi-Source Ingestion**  
  - Consume CDRs from multiple sources, including:
    - Kafka topics (for real-time feeds)
    - SFTP or local flat files (for batch feeds)
  - Supported CDR types: Voice, Data, SMS, Event, Roaming

- **CDR Enrichment**
  - Enrich CDRs with:
    - Subscriber details (resolve IMSI → MSISDN, subscriber profile)
    - Equipment metadata (resolve IMEI → device type)
    - Location information (cell ID → geo-coordinates, region)
    - Service plan lookup (rate plan, account type)
  - Support caching to reduce lookup latency

- **Transformation & Mapping**
  - Apply configurable transformation/mapping rules:
    - Normalize field names/types based on downstream system
    - Convert timestamps to UTC
    - Mask or redact sensitive fields (e.g., IMSI, MSISDN)
    - Add calculated fields (session duration, data volume buckets)
    - Support versioned rule sets per destination

- **Filtering**
  - Filter CDRs based on:
    - Event type (e.g., only "call start" events)
    - Timestamp windows (e.g., last 24h, or configurable)
    - Service ID or custom tags

- **Export/Publishing**
  - Produce transformed CDRs in one or more output formats:
    - CSV, JSON, Avro (configurable per destination)
  - Publish to:
    - Kafka topics (for real-time billing or analytics)
    - SFTP directories (for offline/legacy billing systems)
    - Optionally, HTTP endpoints (for REST-based integrations)
  - Support delivery acknowledgements and retry on failure

- **Audit & Dry-Run Modes**
  - Dry-run: process CDRs without exporting, log all transformations
  - Audit mode: generate transformation logs with before/after snapshots of each CDR

## Non-Functional Requirements

- **Performance**
  - Minimum throughput: 100,000 CDRs/sec
  - Maximum end-to-end latency: < 500ms per CDR (excluding downstream delays)

- **Scalability & Availability**
  - Horizontal scaling via container orchestration (Kubernetes)
  - Zero-downtime deployments supported

- **Extensibility**
  - Pluggable architecture:
    - Each major step (ingestion, enrichment, transformation, export) must be an interface with at least one default implementation
    - Support for third-party enrichment plugins
    - New mapping rules, export formats, and filters can be added at runtime via configuration reload

- **Configuration**
  - All transformation and routing logic must be config-driven (YAML/JSON)
  - Config changes reloadable without service restart

- **Resilience & Error Handling**
  - Automatic retries for transient ingestion/export failures
  - Dead-letter queue support for unprocessable CDRs
  - Centralized structured logging and metrics (Prometheus, OpenTelemetry)

- **Security**
  - All data at rest and in transit must be encrypted (TLS, SFTP)
  - Role-based access for configuration changes and audit logs

- **Backward Compatibility**
  - Must handle legacy 3G/4G CDR schemas as well as 5G

- **Compliance**
  - Support for GDPR redaction rules (PII masking on export)
  - Full audit trail for all transformations

## Out of Scope
- Rating or charging logic (handled downstream)
- Direct integration with external billing systems

## Example Use Cases

1. **Real-Time 5G Data Feed:**  
   Ingests 5G data session CDRs from Kafka, enriches with subscriber and location, masks MSISDN, and exports Avro records to downstream analytics via Kafka.

2. **Legacy Voice Batch Processing:**  
   Pulls 3G/4G voice CDRs as nightly batch files from SFTP, enriches, applies custom mapping per legacy schema, and exports CSV files for legacy billing.

3. **Audit Scenario:**  
   Runs in dry-run mode against a test dataset, logging before/after transformation records for all enrichment and mapping steps, without publishing to any destination.