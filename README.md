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
- less tool for logging in the container and debugging
- maybe some missing dependencies

References:
- [Containers Are Not VMs! Which Base Container Docker Images Should We Use?](https://www.youtube.com/watch?v=82ZCJw9poxM)
