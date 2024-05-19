# Crawler Service

## Introduction

The Crawler Service is designed to fetch information of a person on LinkedIn. The retrieved data is returned in JSON format and includes fields such as about, experiences, education, skills, certifications, and awards. The return format can be customized to meet the requirements of the CVSeeker Server.

## Features
Input a LinkedIn URL and receive a JSON profile as output.

## Project Structure

This project utilizes Django, a Python web framework, and follows the typical Django project structure:

### /crawler

- `__init__.py`: An empty file indicating that this directory is a Python package.
- `settings.py`: Configuration settings for the Django project.
- `urls.py`: URL declarations for the Django project, serving as a "table of contents."
- `asgi.py`: Entry-point for ASGI-compatible web servers.
- `wsgi.py`: Entry-point for WSGI-compatible web servers.

### /main

- `migrations/`: Manages and updates the database structure automatically.
- `filterdata/`: Inclunding modules help to filter response from web services
- `templates/`: Contains HTML files for the user interface.
- `admin.py`: Registers providers and their management models for Django's admin interface.
- `apps.py`: Configuration specific to this Django application.
- `models.py`: Defines providers and their management models representing objects in the database.
- `tests.py`: Contains test cases for verifying the correctness of the application.
- `views.py`: Defines functions for handling HTTP requests and returning responses.

### Other Files

- `manage.py`: Command-line utility for interacting with the Django project.
- `db.sqlite3`: SQLite database used by the project.
- `requirements.txt`: Requirement packages need to be installed for projects

## Getting Started

Follow these steps to set up and run the Crawler Service on your local machine:

### Prerequisites

- Python 3.x installed on your system
- pip package manager

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/tunghng/CVSeeker-server.git
    ```

2. Navigate to the project directory:

    ```bash
    cd CVSeeker/crawler
    ```

3. Create a virtual environment (optional but recommended):

    ```bash
    python3 -m venv venv
    ```

4. Activate the virtual environment:

    - On Windows:

    ```bash
    venv\Scripts\activate
    ```

    - On macOS and Linux:

    ```bash
    source venv/bin/activate
    ```

5. Install the required dependencies using pip:

    ```bash
    pip install -r requirements.txt
    ```


### Usage

1. Start the Django development server:

    ```bash
    python manage.py runserver
    ```

2. Access the Crawler Service at `http://localhost:8000/` in your web browser.
## API Endpoint

### Get Full Text from LinkedIn URLs

Endpoint: `http://127.0.0.1:8000/api/getfulltext/`

This endpoint retrieves the full text content and link profile from the provided LinkedIn URLs.

#### Parameters

- `list_url`: A string containing a list of LinkedIn URLs, separated by commas.

#### Example

```json
{
  "list_url": "https://www.linkedin.com/in/person1, https://www.linkedin.com/in/person2, https://www.linkedin.com/in/person3"
}
```

## Note

Ensure that the virtual environment is activated whenever you work on the project to avoid conflicts with system-wide packages.
