# FOSS Asia Summit 2023
## Session Name: Deploy to Kubernetes from GitHub Container Registry
<br/>

[Slides Link](https://github.com/chamodshehanka/fossasia-summit-2023/blob/main/docs/Deploy%20to%20Kubernetes%20from%20GHCR.pdf).

## About

This Dockerfile starts with an official Go runtime image, sets the working directory to /app, copies the contents of the current directory (where the Go application code is located) into the container at /app, builds the Go application using the go build command, and sets the CMD to run the main executable that was just built.

The EXPOSE instruction indicates that the container will listen on port 8080, but you will still need to use the -p option when running the container to map the container's port 8080 to a port on your host machine. For example, to run this container and map port 8080 to port 8080 on your host machine, you can use the following command:

```bash
docker run -p 8080:8080 <image-name>
```

Replace `<image-name>` with the name you gave to the Docker image when you built it (e.g., my-go-app).



## Setting up Kubernetes Cluster and Deploying the Application
1. Create a Google Kubernetes Engine cluster(GKE) using the Google Cloud Platform Console.
2. Open Google Cloud Shell.
3. Connect to the cluster
4. Create a Kubernetes Deployment file using the following command:

```bash
kubectl create deployment fossasia-2023 --image=ghcr.io/chamodshehanka/fossasia-summit-2023:latest --port=8080 -n dev
```

To expose the deployment as a service, run the following command:

```bash
kubectl expose deployment fossasia-2023 --type=LoadBalancer --port=8080 --target-port=8080 -n dev
```

Or 

```bash
kubectl expose deployment/fossasia-2023 -n dev --type=LoadBalancer
```

### Check the status of the service

```bash

kubectl get service -n dev

```

### Curl the service

```bash

curl http://<EXTERNAL-IP>:8080

```
