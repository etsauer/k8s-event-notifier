# Send notification when a namespace has exceeded its quota

The following is an example showing how k8s-notify can be used to alert when a namespace has reached its quota.

Before we begin, make sure that k8s-notify is running, [either locally or in your Kubernetes cluster](/README.adoc).

First, create a new namespace to represent the app namespace, and apply a quota.

```
kubectl create namespace quota-alert
kubectl apply -f https://raw.githubusercontent.com/redhat-cop/openshift-toolkit/v.10/quota-management/files/quota-small.yml -n quota-alert
```

Now, we'll deploy an app to the namespace.
```
kubectl run resource-consumer --image=gcr.io/kubernetes-e2e-test-images/resource-consumer:1.4 --expose --service-overrides='{ "spec": { "type": "ClusterIP" } }' --port 8080 --requests='cpu=100m,memory=256Mi' --limits='cpu=100m,memory=256Mi' -n quota-alert
```

Now modify and apply one of the [sample notifiers](/examples/notifiers/) via `kubectl apply -f sample_notifier_<type>.yaml -n quota-alert`

Let's subscribe to events about exceeded quotas. For that, we'll apply the `quota_exceeded_eventsub.yaml` EventSubscription.

```
kubectl apply -f quota_exceeded_eventsub.yaml -n quota-alert
```

Finally, let's generate the alert. To do this, we'll attempt to scale our sample app to a larger footprint than our Quota will allow.

```
kubectl scale deployment/resource-consumer --replicas=15 -n quota-alert
```

If everything was done right, you should have received an message via the Notifier you configured.

## Other notes and useful commands

```
kubectl autoscale deployment resource-consumer --min=1 --max=10 --cpu-percent=70
curl --data "millicores=2300&durationSec=30" http://resource-consumer-k8s-notify.apps.d2.casl.rht-labs.com/ConsumeCPU
```
