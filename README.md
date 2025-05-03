# Animated Portfolio Website

This project is a one-page animated portfolio website designed to showcase professional experience, education, skills, and projects with a dynamic and interactive design. It is built using a combination of modern technologies including Go, MongoDB, Tailwind CSS, and htmx. This document provides instructions for both developers looking to contribute to the project and regular users aiming to run and explore the application.

## Features

*   **Dynamic Content:** The website pulls data from a MongoDB database, allowing for easy updates to the displayed content.
*   **Animated Intro:** Features a full-screen background with a circular profile picture and a engaging load animation.
*   **Responsive Design:** Built with Tailwind CSS, ensuring a seamless experience across various devices.
*   **Interactive Elements:** Uses htmx to provide a smooth and interactive experience.
*   **API Support:** Includes a RESTful API endpoint to serve resume data in JSON format.
*	**Modular Design:** Each section (projects, experience, education, skills, contacts) is a dedicated route with it's own design.

## Technology Stack
The portfolio is built with:

-   **Go:** For the backend server and API.
-   **Tailwind CSS:** For styling and layout, generated via CLI.
-   **MongoDB:** For storing and retrieving resume data.
-   **htmx:** For dynamic behaviors and animations.

The application provides a resume API endpoint and several HTML routes.

### Available Routes

-   `/`: Landing page with animated intro.
-   `/api/resume`: JSON data for resume.
-   `/projects`: List of projects.
-   `/experience`: Professional experience timeline.
-   `/education`: Education details.
-   `/skills`: Skills grid.
-   `/contacts`: Contact information.

### Resume Data Format

The resume data is stored in MongoDB and served as JSON from the `/api/resume` endpoint. The structure is defined in the `resume.go` file.

## Prerequisites
Before you begin, ensure you have the following installed on your system:

-   Go 1.22+
-   Node.js and npm (for Tailwind CSS)
-   Docker and Docker Compose

## Getting Started

This section outlines the steps for both developers and regular users to get the application up and running.

### For Regular Users

If you just want to view the portfolio website, follow these steps:

1.  **Clone the Repository:**


1.  **Build CSS:**
