# Knative Serving Networking Openshift

This is an implementation of Knative's `Ingress` resource, specific to Openshift needs. This has been "forked" from upstream's `networking-istio` code and as such, the `pkg/reconciler/ingress` package is a nearly identical copy of upstream's code. The goal is to keep this code up-to-speed with upstream advancements and enhance it where necessary to accomodate for Openshift's needs.

## Keeping on track with upstream

To keep on track with upstream's code base we can cherry-pick changes to the `pkg/reconciler/ingress` folder upstream as we see fit. To see just commits that are local to that folder, use `git log pkg/reconciler/ingress` or navigate to https://github.com/knative/serving/commits/master/pkg/reconciler/ingress.

To cherry-pick a commit, use the `hack/cherry-pick.sh $COMMIT_ID` script as it'll automate a few tasks for you.

Remember to add a **[CP]** prefix to each cherry-picked commit before submitting it.

## Versioning scheme

The versioning scheme consists of the Knative Serving the code should be compatible with, followed by the Openshift Serverless version it'll be shipped with and an increasing number. For example:

- Knative Serving version: **v0.9.0**
- Openshift Serverless version: **1.2.0**
- Current build number: **01**

Makes **v0.9.0-1.2.0-01**

## Building and releasing a new image

To build a new image, use the `hack/build-image.sh` script. It wraps `go build` and `docker build` in a way that makes it look like an image build via operator-sdk. Push the image via `docker push` to quay.io to "release" it.

```bash
$ ./hack/build-image quay.io/openshift-knative/knative-networking-openshift:v0.9.0-1.2.0-01
$ docker push quay.io/openshift-knative/knative-networking-openshift:v0.9.0-1.2.0-01
```