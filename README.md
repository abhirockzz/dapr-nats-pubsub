## Dapr and NATS pubsub

[Dapr version 0.2.0](https://github.com/dapr/dapr/blob/master/docs/release_notes/v0.2.0.md) comes with a bunch of [new components added to the runtime](https://github.com/dapr/dapr/blob/master/docs/release_notes/v0.2.0.md#dapr-runtime). One such component includes pubsub capability with [NATS](https://nats.io/) which is a [Go based](https://golang.org/) open source messaging system for cloud native applications, IoT messaging, and microservices architectures. This blog will provide a step-by-step walk through of how to use it. 

We will deploy `Dapr` on `Kubernetes` (`minikube`) and use the NATS server at `demo.nats.io:4222` for demonstration purposes.

![](https://thepracticaldev.s3.amazonaws.com/i/mml7r8815pr7u7bpi4y7.jpg)