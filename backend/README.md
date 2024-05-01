# CVSeeker Server

CVSeeker Server is the backend server component for the CVSeeker application, designed to manage and facilitate operations related to job seekers' curriculum vitae (CV) storage, processing, and retrieval. This server also provides functionalities for matching job seekers with potential employers based on criteria specified by both parties.

## Features

- **CV Management**: Allows users to upload, update, and retrieve CVs in various formats.
- **Matching Algorithm**: Utilizes advanced algorithms to match job seekers with potential employers based on multiple factors like skills, experience, and job preferences.
- **RESTful API Endpoints**: Provides a set of API endpoints to facilitate communication between the server and client-side applications or third-party integrations.

## Project Structure

The project is structured to facilitate a modular and scalable architecture. Each component has a specific role, ensuring separation of concerns and ease of maintenance. Here's a detailed overview of the repository's structure:

### /cmd
- **/CVSeeker**: This directory contains the main application entry point.
    - **/internal**:
        -  **/cfg**: Manages all configuration-related activities for the CVSeeker application
        -  **/handlers**: Responsible for processing incoming HTTP requests.
        -  **/providers**: Ensuring that dependencies are managed efficiently across the application using dependency injection.
        -  **/services**: Include the business logic needed to interact with resources that the application depends on.
    - **/pkg**: Holds common utilities for reusability.
    - **/statics**: Configuration files
    - **main.go**: Primary executable for the CVSeeker server

### /internal
- **/dtos**: Defines objects that carry data between processes, typically used to aggregate the data and send it to clients.
- **/errors**: Contains error handling logic and error definitions for the application, centralizing error management in one location.
- **/ginLogger**: Custom logging functionality.
- **/ginMiddleware**: Middleware components for the Gin framework.
- **/ginServer**: Setup and configuration code for the Gin HTTP server.
- **/handlers**: Methods for handling response, errors in handlers.
- **/meta**: Metadata utilities or definitions that are used across various parts of the application.
- **/models**: Domain models representing the core business objects within the application.
- **/repositories**: Abstractions over the data layer, providing a collection of methods for accessing and manipulating database records.

### /pkg
- Contains library code that can be used across different projects or within different parts of this application.
    - **/api**: General components or utilities related to API operations.
    - **/aws**: Utilities and helpers for interacting with Amazon Web Services.
    - **/cfg**: Configuration-related code that helps in managing application settings.
    - **/db**: Database interaction utilities and helpers, potentially abstracting some ORM functionalities or database connections.
    - **/elasticsearch**: Components specifically for interacting with Elasticsearch, providing search capabilities.
    - **/gpt**: Utilities or components to interact with OpenAI's GPT models.
    - **/huggingface**: Integrations or utilities for interacting with Hugging Face APIs or models.
    - **/logger**: Centralized logging utilities that can be used throughout the application for consistent logging.
    - **/utils**: Miscellaneous utilities that provide generic functionalities used by various parts of the application.

### Configuration and Miscellaneous
- **.env**: Stores environment-specific variables that configure various aspects of the application, such as database connections, API keys, and operational parameters.
- **/go.mod**: Defines the module's dependencies and manages versions, ensuring consistent builds by locking down specific versions of external packages.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

These steps assume that you have **go 1.21** downloaded, here is how to check the version

```bash
go version 
```

### Installing

A step-by-step guide to setting up a development environment:

1. Clone the repository:
   ```bash
   git clone https://github.com/tunghng/CVSeeker-server.git
   cd CVSeeker
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables:
   Copy the `.env.example` file to `.env` and modify it with your settings.

4. Run the server:
   ```bash
   go build CVSeeker/cmd/CVSeeker
   ```

