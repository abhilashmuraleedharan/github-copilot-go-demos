# CDR Transformation Microservice

A Go-based microservice for ingesting, enriching, transforming, filtering, and exporting Call Detail Records (CDR).

## Project Structure

```
cdr-microservice/
├── main.go                 # Application entry point
├── config/                 # Configuration management
│   ├── config.go          # Configuration structs and loader
│   └── config.yaml        # Configuration file
├── ingestion/             # CDR ingestion service
│   └── service.go         # HTTP server for receiving CDRs
├── enrichment/            # Data enrichment service
│   └── service.go         # Lookup and cache logic
├── transformation/        # Data transformation service
│   └── service.go         # Field transformation rules
├── filtering/             # Data filtering service
│   └── service.go         # Condition-based filtering
└── exporting/             # Data export service
    └── service.go         # Export to various destinations
```

## Features

- **Ingestion**: HTTP endpoint for receiving CDR data with buffering
- **Enrichment**: Lookup-based data enrichment with caching
- **Transformation**: Configurable field transformations (masking, formatting, conversion)
- **Filtering**: Rule-based record filtering
- **Exporting**: Batch export to HTTP, Kafka, or file destinations

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Configuration file at `config/config.yaml`

### Installation

```bash
cd cdr-microservice
go mod download
```

### Running

```bash
go run main.go
```

### Configuration

Edit `config/config.yaml` to customize:
- Ingestion port and buffering
- Enrichment sources and caching
- Transformation rules
- Filtering conditions
- Export destinations

### API Endpoints

- `POST /cdr` - Submit CDR records
- `GET /health` - Health check endpoint

## License

MIT
