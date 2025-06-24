# TecAlliance Link (Go)

[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)](https://go.dev/)
[![Go Report Card](https://goreportcard.com/badge/github.com/cooler-SAI/tecalliance-link)](https://goreportcard.com/report/github.com/cooler-SAI/tecalliance-link)
[![Build Status](https://img.shields.io/github/actions/workflow/status/cooler-SAI/tecalliance-link/go.yml?branch=main)](https://github.com/cooler-SAI/tecalliance-link/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go client library for interacting with the [TecAlliance Order Manager API](https://www.tecalliance.net/en/products-and-solutions/order-manager/). This library aims to provide a simple, robust, and idiomatic Go interface for accessing TecAlliance services.

---

## Features

*   **Type-Safe:** Provides Go structs for API requests and responses.
*   **Authentication:** Handles API key authentication seamlessly.
*   **Article Search:** Methods for searching and retrieving article data.
*   **Order Management:** (Planned) Functionality for creating and managing orders.
*   **Error Handling:** Clear and structured error handling for API responses.

## Getting Started

### Prerequisites

*   Go 1.22 or later.
*   A valid TecAlliance API Key.

### Installation

To add the library to your project, use `go get`:


### Configuration

The client needs to be configured with your API key. It's recommended to use environment variables to keep your credentials secure.

Export your API key in your terminal:


The client will automatically pick up this environment variable.

## Usage

Here is a basic example of how to initialize the client and make a request.

## Contributing

## Contributing

Contributions are welcome! If you'd like to help improve this library, please feel free to:

1. **Fork the repository**
2. **Create a new branch** (`git checkout -b feature/your-feature-name`)
3. **Make your changes**
4. **Commit your changes** (`git commit -m 'feat: Add some amazing feature'`)
5. **Push to the branch** (`git push origin feature/your-feature-name`)
6. **Open a new Pull Request**

Please make sure your code follows the project's coding standards and includes tests for new features.

## License

This project is licensed under the MIT License - see the LICENSE file for details.