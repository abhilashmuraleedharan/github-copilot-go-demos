# Demo 10 – Architecting a 5G CDR Pipeline Using Copilot Plan Mode

This demo showcases how to use GitHub Copilot’s **Plan Mode** to design a
telecom-grade CDR processing pipeline before writing any code.

## Demo Highlights

- Use `/plan` to generate a full architecture + implementation blueprint.
- Pipeline covers:
  - Kafka, HTTP, and SFTP ingestion
  - Subscriber & location enrichment
  - CDR-type–specific rule transformation
  - Kafka + archive exporter
- Copilot generates:
  - Proposed folder hierarchy
  - Modular breakdown
  - Concurrency and worker-pool strategy
  - Testing & benchmarking approach
  - Deployment strategy (Docker + Helm)
  - CI/CD recommendations
  - Observability plan (metrics, logging, tracing)

