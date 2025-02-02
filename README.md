# Project Setup Guide

## Prerequisites

Before you begin, make sure you have the following installed on your system:

- [Docker](https://www.docker.com/) (latest version)

## Steps to Execute the Project

1. **Clone the Repository**
   ```bash
   git clone <https://github.com/adrinant/portfolio-tracker.git>
   cd <portfolio-tracker>
   ```

2. **Build and Start the Services**
   Run the following command to build and start the project:
   ```bash
   docker-compose up --build
   ```

3. **Access the Application**
   Once the services are up, you can access the application (replace `<port>` with the actual port if applicable):
   ```
   http://localhost:<port>
   ```
   the default port is 5173

## Notes
- Make sure Docker is running before executing the commands.