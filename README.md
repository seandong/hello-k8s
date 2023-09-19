[k8s tutorial](https://k8s-tutorials.pages.dev/)

---

## startup
0. startup minikube

```bash
minikube start

# delete old service if start failed

# minikube delete
# minikube start
```


1. startup deployments and services

```bash
kubectl apply -f nginx.yaml
kubectl apply -f hellok8s.yaml


# kubectl get pods
# kubectl get service
```

2. startup ingress

```bash
kubectl apply -f ingress.yaml

# kubectl get ingress
```

3. start minikube url service

```bash
minikube service list
minikube service ingress-nginx-controller -n ingress-nginx --url
```

4. test

```bash
curl http://127.0.0.1:56770/hello # proxy to hello-k8s
curl http://127.0.0.1:56770       # proxy to nginx
```



