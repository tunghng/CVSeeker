# CVSeeker Backend Services D

## Table of Contents
- [1. Introduction](#1-introduction)
- [2. Data Processing Service](#2-data-processing-service)
- [3. Search Service](#3-search-service)
- [4. Chatbot Service](#4-chatbot-service)
- [5. API Endpoints](#5-api-endpoints)
- [6. Implementation Specifics](#6-implementation-specifics)

## 1. Introduction
CVSeeker is designed to streamline the talent acquisition process, enabling teams to efficiently find and engage candidates within their talent pool. This documentation outlines the backend services that support this application.

**[Insert System Architecture Diagram Placeholder]**

## 2. Data Processing Service
### Workflow
When a resume is uploaded, the data processing service initiates a background job to handle the file:
1. **File Storage:** The resume is stored in AWS S3.
2. **Data Parsing:** The full text of the resume is extracted and formatted using OpenAI's GPT into a predefined JSON structure.
3. **Vector Embedding:** The text is also sent to a Hugging Face model to be converted into vector format.
4. **Indexing:** The JSON data, vector array, and S3 link are indexed in Elasticsearch.
5. **Notification:** A WebSocket sends real-time notifications to the client about the status of the upload.

### Data Structure Example
```json
{
    "summary": "Provide a concise professional summary based on the resume.",
    "skills": ["List of skills"],
    "basic_info": {
        "full_name": "Invented Full Name",
        "university": "Generated University Name",
        "education_level": "BS",
        "majors": ["List of Majors", "GPA: 3.5"]
    },
    "work_experience": [{
        "job_title": "Title",
        "company": "Company Name",
        "location": "Location",
        "duration": "Duration",
        "job_summary": "Job responsibilities and achievements"
    }],
    "project_experience": [{
        "project_name": "Project Name",
        "project_description": "Project details including technologies used"
    }],
    "award": [{"award_name": "Award Name"}]
}
```

**[Insert Data Flow Diagram Placeholder]**

## 3. Search Service
The search service allows users to perform hybrid searches combining keyword and semantic approaches:
1. **Query Input:** Users input a search query with a semantic threshold.
2. **Vectorization:** The query is vectorized using the same Hugging Face model.
3. **Matching:**
    - Step 1: Use ANN to narrow down potential matches.
    - Step 2: Compute cosine similarity with all resumes in the Elasticsearch index.

Results are presented in the search interface, ranked by match quality.

**[Insert Search Workflow Diagram Placeholder]**

## 4. Chatbot Service
Users can interact directly with selected resumes through a chat interface powered by OpenAI's Assistant API:
1. **Session Management:** Users start chat sessions with selected resumes. Each session creates a new thread, and all candidate information is loaded into this thread.
2. **Interaction:** User messages are processed by the Assistant API, with responses streamed back to the frontend via WebSocket.
3. **Session Continuity:** Users can revisit previous threads to continue interactions and review associated resumes.

**[Insert Chatbot Service Workflow Diagram Placeholder]**

## 5. API Endpoints
### Data Processing Service
#### Upload Resumes
- **POST** `/cvseeker/resumes/upload`
   - **Description:** Processes uploaded resume files and associated metadata as JSON.
   - **Request Body:** `{ "resumeData": [array of resume files and metadata] }` (required)
   - **Responses:** `200 OK`, `400 Bad Request`, `401 Unauthorized`, `500 Internal Server Error`

### Search Service
#### Perform Hybrid Search
- **POST** `/cvseeker/resumes/search`
   - **Description:** Executes a search combining keyword and vector-based queries.
   - **Request Body:** `{ "query": "search keywords", "knnBoost": 0.5 }` (required)
   - **Responses:** `200 OK`, `400 Bad Request`, `401 Unauthorized`, `500 Internal Server Error`

### Chatbot Service
#### Start Chat Session
- **POST** `/cvseeker/resumes/thread/start`
   - **Description:** Starts a new chat session using specified document IDs.
   - **Request Body:** `{ "ids": "comma-separated document IDs", "threadName": "optional name" }` (required)
   - **Responses:** `200 OK`, `400 Bad Request`, `500 Internal Server Error`

#### Get Messages from a Thread
- **GET** `/cvseeker/resumes/thread/{threadId}/messages`
   - **Description:** Retrieves messages from a specified chat thread.
   - **Path Parameters:** `threadId` (required)
   - **Query Parameters:** `limit`, `after`, `before` for pagination
   - **Responses:** `200 OK`, `400 Bad Request`, `401 Unauthorized`, `404 Not Found`, `500 Internal Server Error`

#### Send Message to a Thread
- **POST** `/cvseeker/resumes/thread/{threadId}/send`
   - **Description:** Sends a message to the specified thread.
   - **Path Parameters:** `threadId` (required)
   - **Request Body:** `{ "message": "text" }` (required)
   - **Responses:** `200 OK`, `400 Bad Request`, `500 Internal Server Error`

## 6. Implementation Specifics
(Provide details on database connections, security considerations, performance metrics, error handling, etc.)

**[Insert Additional Diagrams or Technical Specification Placeholders as Needed]**
