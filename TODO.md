- Add `POST /pay` endpoint in both `payment-gateway` and `payment-service`
    - `payment-gateway` simply accepts a JSON payment request and forwards it to `payment-service`

- Add Tests - Unit Tests, Integration Tests
    - For Request Handling

- Handle OS Signals - SIGTERM? What else? SIGKILL? SIGHUP?
    - Shutdown the process only if all the HTTP requests are done processing!
        - Also, ignore readiness and health check HTTP requests when checking if all requests are done processing

- Run the containers as non root user
    - Make changes in the container image - in the `Dockerfile` and also ensure that the Kubernetes deployment and local Docker deployment run the containers with non root user

- Add CI for building and publishing container images to a container registry
    - For example, use GitHub Actions and GitHub Container Image Registry

- Release versioned container images. Currently it's always just `latest` image tag

- Add 2 separate endpoints for liveness and readiness - `livez` and `ready` and remove `healthz`

- Consider adding a `Makefile` for simplifying local build and other local operations to be done for local development

- Add authentication and authorization for all the requests coming to the services

- In Kubernetes - Expose the `payment-gateway` through an `Ingress` resource or using `HTTPRoute` resource (Gateway API), so that clients external to the Kubernetes cluster can access the `payment-gateway`. Also, add instructions and resources (Helm Charts or YAML files) on how to deploy a ingress controller / gateway / load balancer like nginx, kong or similar in the local Kubernetes cluster to test the routing from external clients

- Add instructions and resources (Helm Charts or YAML files) on how to deploy a Prometheus or Prometheus-like monitoring system in the local Kubernetes cluster and also configure the Prometheus server to monitor the `payment-gateway` and `payment-processor` service - for example using custom resources like `ServiceMonitor`. Custom resources like `Prometheus` can also be used to install Prometheus. Operators like Prometheus operator help with managing such custom resources

- Emit metrics around HTTP server in both the services. Like HTTP request count, count of different HTTP response codes, HTTP response times. This will help with monitoring the services at the HTTP level

- Add instructions and resources (Helm Charts or YAML files) on how to configure Prometheus to emit alerts when some condition is true - using alert rules. For example, service down. This can be done using custom resources like `PrometheusRule` that Prometheus operator provides

- Add instructions and resources (Helm Charts or YAML files) on how to deploy and configure Alertmanager to manage the alerts that Prometheus emits. This can be done using `Alertmanager` custom resource to deploy Alertmanager and to configure it using `AlertmanagerConfig` custom resource. Both of these custom resources are provided by Prometheus operator

- Autoscaling
    - Use the Horizontal Pod Autoscaler configuration to scale based on basic metrics like CPU, Memory
    - In the future, consider what other metrics / custom metrics can be used to scale the services - like number of HTTP requests etc

- Use the following for better security -
    - Rate Limiting Systems - this is important in a production environment - especially in public networks (Internet) but even in private networks where clients can send way too many requests than what the services might be able to handle. So, it's better to use rate limiting to ensure fair usage policy (FUP) across all clients

    - Blocking IP if noticing unexpected behaviour / attacker behaviour for example - accessing API multiple times without Authentication or proper Authorization

    - Make the communication between `payment-gateway` and `payment-service` secure by using mTLS - Mutual TLS (Transport Layer Security)

    - For mTLS communication - distribute the certificates to all the entities - client and server, using a secure channel - say, use some secret management system like Hashicorp Vault

    - Expose only the payment gateway and not the payment service

    - Whitelist / Allow list trusted IPs / IP range for clients accessing the payment gateway

    - Expose payment gateway in a private network if all clients can access from a private network. Say, do peering, VPN etc to connect two private networks. If client is in public network and will access with a public IP, then we need to expose payment gateway in the public network (Internet) and also take care of securing the publicly exposed service - using Rate Limiting Systems, Allow listing IPs, Blocking IP address / user for unexpected behaviour / attacker behaviour, Using Web Application Firewalls (WAFs) in general

    - Ensure that only payment gateway can access payment service and that no other entity can access payment service

    - Run all containers in all pods securely. Secrets should not be leaked. No exec access
