kubectl apply -f ./userhandle/user_deployment.yaml
kubectl rollout restart deployment userhandle
kubectl apply -f ./products/products_deployment.yaml
kubectl rollout restart deployment products
kubectl apply -f ./orders/order_deployment.yaml
kubectl rollout restart deployment orders
kubectl apply -f ./cloudstore_site/website_deployment.yaml
kubectl rollout restart deployment website
kubectl apply -f cloudstore_ingress.yaml
