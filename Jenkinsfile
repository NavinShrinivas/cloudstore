pipeline {
    agent any
    stages {
        stage('cloning') {
            steps {
               git branch : "main",
               url : "git@github.com:NavinShrinivas/cloudstore.git"
            }
        }
        stage("build"){
            steps{
               sh 'docker login'
               sh 'docker build ./userhandle/ -t  navinshrinivas/cloudstore_userhandle'
               sh `docker push navinshrinivas/userhandle`
               sh 'docker build ./products/ -t navinshrinivas/products'
               sh `docker push navinshrinivas/products`
               sh 'docker build ./orders/ -t navinshrinivas/orders'
               sh `docker push navinshrinivas/orders`
            }
        }

        stage("deploy"){
            steps{
            //Assuming minikube to be running
               sh 'kubectl config set-cluster minikube'
               sh 'kubectl apply -f ./userhandle/user_deployment.yaml'
               sh 'kubectl rollout restart deployment userhandle'

               sh 'kubectl apply -f ./products/products_deployment.yaml'
               sh 'kubectl rollout restart deployment products'

               sh 'kubectl apply -f ./orders/order_deployment.yaml'
               sh 'kubectl rollout restart deployment orders'

               sh 'kubectl apply -f cloudstore_ingress.yaml'
            }
        }
    }
    post{
      failure{
         echo "Pipline failed"
      }
    }
}
