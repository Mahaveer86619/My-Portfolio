# Animated Portfolio Website

This project is a one-page animated portfolio website built using Go, MongoDB, Tailwind CSS, and htmx. It showcases professional experience, education, skills, and projects with a dynamic and interactive design.

## Project Overview

The portfolio is built with:

-   **Go:** For the backend server and API.
-   **Tailwind CSS:** For styling and layout, generated via CLI.
-   **MongoDB:** For storing and retrieving resume data.
-   **htmx:** For dynamic behaviors and animations.

The application provides a resume API and several HTML routes.

## Available Routes

-   `/`: Landing page with animated intro.
-   `/api/resume`: JSON data for resume.
-   `/projects`: List of projects.
-   `/experience`: Professional experience timeline.
-   `/education`: Education details.
-   `/skills`: Skills grid.
-   `/contacts`: Contact information.

## Resume Data Format

The resume data is stored in MongoDB and served as JSON from the `/api/resume` endpoint. The structure is defined in the `resume.go` file.

## Landing Page Animation

The landing page features a full-screen background and a circular profile picture. Upon loading, after a 1.5-second delay, the background zooms out, and widgets slide into view from off-screen positions.

## How to Use

### Prerequisites

Make sure you have the following installed:

-   Go 1.22+
-   Node.js and npm (for Tailwind CSS)
-   Docker and Docker Compose

### Building and Running

1.  **Build CSS:**
