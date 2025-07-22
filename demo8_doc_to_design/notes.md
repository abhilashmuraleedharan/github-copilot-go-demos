# Copilot Modes Demo: Use Cases, Stages, and Example Prompts

This guide will help you showcase GitHub Copilot’s **Ask**, **Edit**, and **Agent** modes in a practical, stepwise demo using your CDR Transformation Microservice specification (`specs.md`).  
For each stage, you’ll find the recommended mode and sample prompts.

---

## Stage 1: Extracting Requirements (Ask Mode)

**Mode:** Ask  
**Goal:** Show Copilot as a smart assistant for understanding and clarifying specs.

**Demo Prompts:**
- `Extract all functional and non-functional requirements from specs.md file. Organize them into two separate lists.`
- `Summarize the key business and technical use cases described in specs.md file.`

---

## Stage 2: Designing Architecture (Ask Mode)

**Mode:** Ask  
**Goal:** Show Copilot brainstorming design and architecture from natural language requirements.

**Demo Prompts:**
- `Propose a modular architecture for the CDR Transformation Microservice based on specs.md file. List the main components and their responsibilities.`
- `Suggest a data flow diagram description for the system outlined in specs.md file.`

---

## Stage 3: Project Scaffolding (Edit Mode)

**Mode:** Edit  
**Goal:** Show Copilot generating actual project structure and code files.

**Demo Prompts:**
- In the root of your repo, prompt:  
  `Generate Go project scaffolding for the CDR Transformation Microservice, with subfolders for ingestion, enrichment, transformation, filtering, exporting, and config, based on specs.md file.`
- In the config package or folder:  
  `Add a Go struct to represent the microservice configuration, including rule selection, enrichment sources, and output formats.`

---

## Stage 4: Advanced Automation (Agent Mode)

**Mode:** Agent  
**Goal:** Demonstrate Copilot’s ability to automate multi-step or cross-file tasks, such as wiring up modules, creating tests, or opening PRs.

**Demo Prompts:**
- `Implement the CDR Enricher module, wire it into the service main, and create a test file for it.`
- `Add a REST endpoint for audit logs, generate OpenAPI docs, and open a pull request.`

---