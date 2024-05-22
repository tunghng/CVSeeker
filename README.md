## CVSeeker
CVSeeker is designed to streamline the talent acquisition process, enabling teams to efficiently find and engage candidates within their talent pool
- [Slides](./backend/statics/PresentationSlides.pdf)

## Deployment Instructions

To deploy the CVSeeker application using Docker Compose, follow these steps:

1. **Ensure Docker and Docker Compose are installed on your system.**

2. **Clone the repository:**
    ```sh
    git clone https://github.com/tunghng/CVSeeker.git
    cd CVSeeker
    ```

3. **Build and start the containers:**
    ```sh
    docker-compose up --build
    ```

4. **Access the application:**
   Open your browser and navigate to `http://localhost:5173` to use the CVSeeker application.

5. **Stop the containers:**
    ```sh
    docker-compose down
    ```

Ensure that any environment variables required by the application are correctly configured in the `.env` file before starting the containers.
