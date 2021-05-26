# Use a Knative Service to Read Secret and Config Map Data

Create a Knative service that can be requested to read data in a secret or a
config map.

For the procedures here, I don't use [Kubernetes
client-go](https://github.com/kubernetes/client-go) to access secrets or config
maps. Instead, I am mounting the secret or config map to a pod through Knative
and accessing that through the pod's volumes.

There are some advantages to not using client-go:

- minimize dependencies on third party packages
- minimize load on the Kubernetes API
- it is faster to access data in a mounted volume compared to sending an HTTP
  request to a server

This procedure assumes the following have been installed:

- Kubernetes and its `kubectl` client
- Knative

I use [Kind](https://kind.sigs.k8s.io/) for my local Kubernetes development
cluster. I used [KonK](https://github.com/csantanapr/knative-kind) to help
install Knative.

## Procedure

1. Build with `ko`

    `ko` was created to help Knative developers build images and binaries
    without the need for Dockerfiles nor Makefiles. See more information in
    [this Knative
    blog](https://knative.dev/blog/2018/12/18/ko-fast-kubernetes-microservice-development-in-go/).

    Follow setup instructions at https://github.com/google/ko.

    After installing ko, set the destination for images with an environment variable.

    ``` bash
    KO_DOCKER_REPO=my-dockerhub-user
    ```

    Run the following command from the `kn-go-echo` directory. It will build and
    push the image to your local Docker daemon. To have `ko` publish and use a
    container image registry, remove the `--local` flag.

    ```bash
    ko publish --local --tags 0.1 --base-import-paths .
    ```

1. Create secret

    ```bash
    kubectl create secret generic secret-msg --from-file=messages/secret.toml
    ```

1. Create config map

    ```bash
    kubectl create configmap cm-msg --from-file=messages/configmap.toml
    ```

1. Deploy read-secret-and-cm Knative service

    ```bash
    kubectl apply -f knative.yaml
    ```

    Note: Currently, the kn client does not support mounting secrets or config
    maps into pods. That is why we're using kubectl and Kubernetes manifests to
    create the Knative service and trigger.

1. Get the URL to the read-secret-and-cm service

    ```bash
    NAME                 URL                                                  LATEST                     AGE     CONDITIONS   READY   REASON
    read-secret-and-cm   http://read-secret-and-cm.default.127.0.0.1.nip.io   read-secret-and-cm-00001   17h     3 OK / 3     True
    ```

1. Go to the following endpoints from a browser to see the messages

    - http://read-secret-and-cm.default.127.0.0.1.nip.io
    - http://read-secret-and-cm.default.127.0.0.1.nip.io/cm
    - http://read-secret-and-cm.default.127.0.0.1.nip.io/secret
