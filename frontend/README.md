# CVSeeker Frontend

This is the frontend of the CVSeeker project. It is a web application that allows users to search for job candidates based on their resumes. The frontend is built using ReactJS.

## Features

## Project Structure

### /design
- Contains design wireframes, mockups, UI/UX designs for the application.

### /public
- logo.png: The logo used for the application.
- logo.ico: The favicon for the application.

### /src
- **/assets**: Contains static assets like images, fonts, and other resources used in the application.
- **/components**: Contains reusable components that can be used across different parts of the application.
- **/contexts**: Contains context providers and consumers for managing global state in the application.
- **/pages**: Contains the main pages of the application, each representing a different view or route.
- **/services**: Contains service files that handle API calls and other data-related operations.
- **/styles**: Contains global styles and CSS files used throughout the application.
- **App.jsx**: The main component that serves as the entry point for the React application.
- **main.jsx**: The main file that renders the React application into the DOM.
- **routes.js**: Contains the routing configuration for the application.

### Configuration and Miscellaneous
- **.gitignore**: Specifies intentionally untracked files to ignore in Git.
- **index.html**: The main HTML file that serves as the entry point for the React application.
- **package.json**: Contains metadata and dependencies for the project.
- **tailwind.config.js**: Configuration file for Tailwind CSS.
- **README.md**: Contains information about the project and instructions for setting up the development environment.


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

These steps assume that you have **Node.js** downloaded, here is how to check the version

```bash
node -v
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
npm install
```

3. Run the development server:

```bash
npm run dev
```

4. Open your browser at `http://localhost:5173` to view the application.