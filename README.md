# Animated Portfolio Website: Dynamic and Interactive Resume

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

This project is a dynamic and interactive single-page portfolio website, designed to showcase professional experience, education, skills, and projects. Built using a modern technology stack, including Go, MongoDB, Tailwind CSS, and htmx, this portfolio offers an engaging and interactive user experience. This document provides comprehensive instructions for both developers and end-users to set up, run, and use the application.

## Table of Contents

1.  [Features](#features)
2.  [Technology Stack](#technology-stack)
3.  [Routes](#routes)
    *   [Web Page Routes](#web-page-routes)
    *   [API Routes](#api-routes)
4.  [Prerequisites](#prerequisites)
5.  [Getting Started](#getting-started)
    *   [For End-Users](#for-end-users)
    *   [For Developers](#for-developers)
6.  [API Usage](#api-usage)
7.  [Contributing](#contributing)
8.  [License](#license)

## Features <a name="features"></a>

*   **Dynamic Content:** The website pulls data from a MongoDB database, allowing for easy updates to the displayed content.
*   **Animated Intro:** Features a full-screen background with a circular profile picture and a engaging load animation.
*   **Responsive Design:** Built with Tailwind CSS, ensuring a seamless experience across various devices.
*   **Interactive Elements:** Uses htmx to provide a smooth and interactive experience.
*   **API Support:** Includes RESTful API endpoints to manage resume data.
*   **Modular Design:** Each section (projects, experience, education, skills, contacts) is a dedicated route with its own design.
* **Easy Setup**: Uses docker for easy setup.

## Technology Stack <a name="technology-stack"></a>

The portfolio is built with:

-   **Go:** For the backend server and API.
-   **Tailwind CSS:** For styling and layout, generated via CLI.
- **htmx:** for dynamic behavior.
-   **MongoDB:** For storing and retrieving resume data.
-   **htmx:** For dynamic behaviors and animations.

## Routes
The application provides both web page routes and API routes.

### Web Page Routes <a name="web-page-routes"></a>

These routes serve the HTML content for the website:

*   `/`: Landing page with animated introduction.
*   `/projects`: Displays the list of projects.
*   `/experience`: Displays the professional experience timeline.
*   `/education`: Displays education details.
*   `/skills`: Displays the skills grid.
*   `/contacts`: Displays contact information.

### API Routes <a name="api-routes"></a>

These routes are used for interacting with the resume data:

*   `GET /api/resume`: Retrieves the complete resume data in JSON format.
    *   **Response Example:**



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
