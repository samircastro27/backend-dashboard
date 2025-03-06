# CCP Platform

This service groups key functionalities related to **cp**, such as **Background Services** (e.g. `DeletionWorker`):
- **Background services** (e.g. `DeletionWorker`).
- **Metrics**.
- **Logs**.
- **Events**.

## Installation and Configuration

### Prerequisites

Make sure you have installed:
- **Docker** and **Docker Compose**.
- **Go 1.20+**.
- **Dapr CLI**, you can follow the [official installation guide](https://docs.dapr.io/getting-started/).

### Environment Variables Configuration

Configure the necessary environment variables based on the `.env.example` file.
All these variables can be found in the Loop documentation and there follow the indications of the cp-platform table.

### Pre-commit

Install pre-commit
```bash
pip3 install pre-commit
```

then install the hooks
```bash
pre-commit install
pre-commit install --hook-type commit-msg
```

With this, the pre-commit would be ready for validations.

### Start Service

#### Step by Step
1. **Start Dapr**:
   Once the Dapr CLI is installed, run:
```bash
  dapr init --slim
```
This will set up a lightweight Dapr environment.

2. **Run the service:**
Start the service using the Dapr configuration file:
> remember to have connection to postgres and rabbit.

***Options***
- docker-compose
- port-forward to the service

**Using the Makefile**
You can simplify the process by running:
```bash
  make run-server
```
This command will raise the entire environment, provided that the external tools (such as Dapr) are correctly configured.

## Test
For test execution it is
- `make test`
- `go test -v -test.short -p=1 -run “^*__Unit” ./test/...`

## Architecture

For more details about the architecture, see the [architecture document](docs/architecture.md).
