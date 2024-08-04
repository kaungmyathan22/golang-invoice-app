# Golang Invoice App

## Overview

The Golang Invoice App is a simple and efficient tool for managing and generating invoices. Built with Go, this application allows users to create, view, and manage invoices seamlessly.

## Features

- **Create Invoices**: Generate new invoices with itemized details.
- **View Invoices**: View and search through a list of created invoices.
- **Manage Invoices**: Edit and delete invoices as needed.
- **Data Persistence**: Store invoice data using a chosen database.

## Installation

To get started with the Golang Invoice App, follow these steps:

1. **Clone the Repository**:

    ```bash
    git clone https://github.com/kaungmyathan22/golang-invoice-app.git
    cd golang-invoice-app
    ```

2. **Install Dependencies**:

    Ensure you have Go installed. If not, download and install Go from the [official Go website](https://golang.org/dl/).

    Run the following command to download the required dependencies:

    ```bash
    go mod tidy
    ```

3. **Build the Application**:

    ```bash
    go build -o invoice-app
    ```

4. **Run the Application**:

    ```bash
    ./invoice-app
    ```

    The application will start running on `localhost:8080` by default. You can change the port by setting the `PORT` environment variable.

## Configuration

Configuration options can be set using environment variables or a configuration file.

### Environment Variables

- `PORT`: Port on which the server will run (default: `8080`).
- `DATABASE_URL`: URL for the database connection.

### Configuration File

You can also use a configuration file (`config.json`) for advanced settings.

## Usage

Once the application is running, you can access the following endpoints:

- **Create Invoice**: `POST /invoices` – Create a new invoice.
- **View Invoices**: `GET /invoices` – List all invoices.
- **View Invoice**: `GET /invoices/{id}` – View a specific invoice by ID.
- **Update Invoice**: `PUT /invoices/{id}` – Update a specific invoice by ID.
- **Delete Invoice**: `DELETE /invoices/{id}` – Delete a specific invoice by ID.

## Testing

To run the tests for the application, use:

```bash
go test ./...
```

## Contributing

If you'd like to contribute to the project, please follow these steps:

1. **Fork the repository**.
2. **Create a new branch**:

    ```bash
    git checkout -b feature-branch
    ```

3. **Commit your changes**:

    ```bash
    git commit -am 'Add new feature'
    ```

4. **Push to the branch**:

    ```bash
    git push origin feature-branch
    ```

5. **Create a new Pull Request**.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- **Go programming language**
- **Go Gin Framework** - For building the web server
