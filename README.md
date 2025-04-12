# Date & Time App

## Overview
This is a simple web application built using GoLang that displays the current date and time. The app is containerized using Docker and deployed to Kubernetes with two replicas, exposed to the internet using a LoadBalancer service.

## Features
- Displays the current date and time in `YYYY-MM-DD HH:MM:SS` format.
- Built with GoLang for simplicity and performance.
- Containerized using Docker.
- Deployed on Kubernetes with two replicas.

## Prerequisites

- Docker
- Kubernetes (Minikube or any cloud provider's Kubernetes service)
- Git
- GoLang (for development purposes)

## Project Setup

### 1. Clone the Repository

Clone the repository to your local machine:
 git clone https://github.com/your-username/date-time-app.git
 cd date-time-app

 ### 2. Docker Setup
-Build the Docker Image
To build the Docker image for the Go application, run the following command:
 docker build -t date-time-app .

-Run the Docker Container
To run the Docker container locally and expose the app on port 8080, execute:
docker run -p 8080:8080 date-time-app

### 3. Kubernetes Setup
-Deployment
First, create the Kubernetes deployment and expose the application with the following kubectl commands:
kubectl apply -f deployment.yaml

-Expose the Service
To expose the application to the internet, you need to expose the service with a LoadBalancer. Run the following command:
kubectl expose deployment date-time-app --type=LoadBalancer --name=date-time-app-service

-Verify the Service
Check the external IP of the service:
kubectl get svc


###Docker Hub
-The Docker image for the app is also available on DockerHub. You can pull the image from there and run it without building it locally:

docker pull dhanusankar66/date-time-app:latest
Then run the container using:
docker run -p 8080:8080 your-docker-username/date-time-app:latest

###License
This project is open source and available under the MIT License.

Just replace `https://github.com/your-username/date-time-app.git` with the actual URL of your GitHub repository.

Let me know if you need further assistance!






