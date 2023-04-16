sudo systemctl start docker
minikube stop 
minikube delete
minikube start
minikube addons enable ingress
sudo systemctl start jenkins
kubectl proxy --port=8080

