[![Deploy to ECR](https://github.com/tbriot/go-webapp-hello-world/actions/workflows/gh_worflow.yml/badge.svg?branch=main)](https://github.com/tbriot/go-webapp-hello-world/actions/workflows/gh_worflow.yml)

# Docker images size comparison
![docker image size screenshot](./images/docker-image-size-comparison.png)

Reasons smaller is better:
- image storage costs $$$
- image download/upload cost $$$ (network traffic)
- faster image download = faster container startup
- less stuff installed = less cpu and memory usage
- less stuff installed = less vulnerabilities
- less stuff installed = potential less licence issues

Drawbacks:
- less tool to log in the container and debugging
- maybe some missing dependencies

References:
- [Containers Are Not VMs! Which Base Container Docker Images Should We Use?](https://www.youtube.com/watch?v=82ZCJw9poxM)
- [Google "Distroless" images](https://github.com/GoogleContainerTools/distroless)
- [Create the smallest and secured golang docker image based on scratch](https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324)
