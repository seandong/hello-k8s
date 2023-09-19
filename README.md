[k8s tutorial](https://k8s-tutorials.pages.dev/)

---

## startup

#### startup minikube

```Shell
minikube start

# delete old service if start failed

# minikube delete
# minikube start

kubectl delete deployment,service,pod --all
```

#### build and push docker image `hellok8s`

```Shell
docker build . -t seandong/hellok8s:v5
docker push seandong/hellok8s:v5
```


#### startup deployments and services

```Shell
kubectl apply -f nginx.yaml
kubectl apply -f hellok8s.yaml


# kubectl get pods
# kubectl get service
```

#### startup ingress

```Shell
minikube addons enable ingress

kubectl apply -f ingress.yaml

# kubectl get ingress
```

#### start minikube url service

```Shell
minikube service list
minikube service ingress-nginx-controller -n ingress-nginx --url
```

#### test

```Shell
curl http://127.0.0.1:56770/hello # proxy to hello-k8s
curl http://127.0.0.1:56770       # proxy to nginx
```



## namespace to isolate environments

```Shell
kubectl apply -f namespaces.yaml
kubectl get configmap --all-namespaces

kubectl apply -f deployment.yaml -n dev
kubectl get pods -n dev
```

## configmap

```Shell
# dev
kubectl apply -f hellok8s-config-dev.yaml -n dev
kubectl apply -f pod-hellok8s.yaml -n dev
kubectl port-forward hellok8s-pod 3000:3000 -n dev
curl http://localhost:3000

# test
kubectl apply -f hellok8s-config-dev.yaml -n test
kubectl apply -f pod-hellok8s.yaml -n test
kubectl port-forward hellok8s-pod 3000:3000 -n test
curl http://localhost:3000
```