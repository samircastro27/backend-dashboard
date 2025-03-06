## Architecture

## Introduction

This project is designed with a **clean architecture** philosophy, adopting a **hexagonal** approach in each module within the `cmd/` directory. In addition, **Dapr** is implemented as an abstraction layer for inter-service communication, following best practices for distributed application implementation and standardization in **monorepos**. 

Dapr, with its modular approach, integrates seamlessly with this framework and offers support for patterns such as pub/sub, bindings, and stateful storage, which enhances the flexibility and scalability of the system.

## Architecture Philosophy

The project is organized following **Clean Architecture** principles, clearly separating responsibilities into layers:  

1. **Domain**: Defines the entities and logic of the pure domain.
2. **Application**: Contains use cases and application logic, communicating with other layers through **ports**.
3. **Infrastructure**: Handles the interaction with frameworks, external libraries and the underlying infrastructure.

### Integration with Dapr

**Dapr** (Distributed Application Runtime) is the core of the architecture. This runtime simplifies the construction of microservices by means of:
- Pub/Sub for asynchronous messaging.
- State Management for distributed state management.
- Input/Output Bindings for integration with external systems.
- Native observability with support for tracing and metrics.

The project is organized as a **monorepo**, allowing code and resources to be shared among multiple services, simplifying their development and deployment.

### Directory `cmd/`

Each subdirectory under `cmd/` represents a standalone service that follows the hex approach. For example:
- `cmd/api/`: Primary HTTP service that exposes endpoints and manages middleware.
- cmd/deletionworker/`: Worker specialized in resource deletion following specific strategies.
- cmd/events/`: Service in charge of event generation and management.
- cmd/metrics/`: Service for the collection and display of metrics.

### Directory `pkg/`

This directory contains utilities and shared libraries, such as:
- `dapr/`: Client for interacting with Dapr bindings.
- `logger/`: Centralized logging library.
- Prometheus/`: Integration with Prometheus for metrics.
- `teamswebhook/`: Client for Microsoft Teams notifications.

## Main features

1. **Modularity**: Each module has a clear separation of responsibilities.
2. **Compatibility with Dapr**: Use of Pub/Sub, State Management and Bindings as key elements.
3. **Extensive Testing**: Detailed unit tests in each layer of the system.
4. **Monorepo**: Facilitates collaboration, code reusability and deployment.

# Development and Contributions
Contribution Principles

- Follow clean architecture conventions in each module.
- Implement unit tests for any new functionality.
- Clearly document public interfaces and functions.

# Test creation and structuring

Tests are organized in a way that is consistent with the structure of the project modules. This allows specific tests to be performed for both the use cases and services of each module, ensuring that the separation of responsibilities defined in the architecture is respected. Each module has its own set of unit and integration tests, aligned with the logic and functionalities of the corresponding domain. This organization facilitates the traceability of errors and promotes a more efficient and modular development.

In addition, the test structure follows Clean Architecture principles, ensuring that:

- Unit tests focus on pure logic within the domain and use cases, decoupled from the infrastructure.
- Integration tests validate the interaction between adapters, infrastructure and use cases.

This ensures complete system coverage, from individual components to their integration as a whole.