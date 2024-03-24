# Finance Tracker Expense Service

The Expense Service is a crucial microservice within the Finance Tracker application, responsible for managing expense transactions. Designed with Go, it leverages gRPC for inter-service communication and SQLite3 for data persistence, ensuring a lightweight yet robust solution for tracking expenses. This microservice works in tandem with the Finance Tracker Gateway and other services, following a microservice architecture pattern to provide scalability and modularity.

## Overview

The Expense Service handles the creation and retrieval of expense records. It defines and implements the necessary gRPC interfaces and relies on Protocol Buffers for efficient data serialization. Designed to be containerized, it seamlessly integrates with Docker and Docker Compose deployments, facilitating easy setup and scalability within the Finance Tracker ecosystem.

## Getting Started

To run and develop the Expense Service as part of the Finance Tracker system, follow these steps:

### Prerequisites

- Docker and Docker Compose if running within containers.
- Go (version 1.22 or newer) for local development.
- Access to the Finance Tracker Proto Files and Gateway service repositories.

### Setup

1. **Clone the Repository**:
   Begin by cloning this repository to your machine. Ensure you also have the Gateway and Proto Files repositories cloned if you're setting up the entire system.

   ```bash
   git clone https://github.com/almaraz333/finance-tracker-expense-service.git
   ```

2. **Install Dependencies**:
   Navigate into the cloned repository directory and install the necessary Go packages and dependencies.

   ```bash
   go mod tidy
   ```

3. **Running Locally**:
   To run the Expense Service locally, execute the following command within the repository directory:

   ```bash
   go run main.go
   ```

### Integration with Gateway

To integrate this service with the Finance Tracker Gateway, ensure:

- The gRPC server address of the Expense Service is correctly configured in the Gateway service.
- The Expense Service Proto Files are up to date with the [Protocol Buffers Repository](https://github.com/almaraz333/finance-tracker-proto-files).

## Dockerization

This service is designed to be run in a Docker container. A `Dockerfile` is included in the repository, and the service can be deployed using Docker Compose in conjunction with the Gateway and any other services. For details on how to set this up, refer to the Finance Tracker Gateway repository's Docker Compose instructions.

## API Reference

The Expense Service defines the following gRPC methods:

- `CreateExpense`: Creates a new expense record.
- `ListExpenses`: Lists all stored expense records.

Refer to the proto files in the [Protocol Buffers Repository](https://github.com/almaraz333/finance-tracker-proto-files) for detailed API specifications.

## Contributing

Contributions to the Expense Service are welcome. To contribute:

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -am 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

This service is part of the Finance Tracker application, built using open-source technologies. Special thanks to the contributors and maintainers of Go, gRPC, Protocol Buffers, and SQLite3 for their incredible tools and libraries.
