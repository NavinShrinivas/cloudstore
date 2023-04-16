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
               sh 'docker login -u navinshrinivas -p kakana071129'
               sh 'docker build ./userhandle/ -t  navinshrinivas/cloudstore_userhandle'
               sh 'docker push navinshrinivas/cloudstore_userhandle'
               sh 'docker build ./products/ -t navinshrinivas/cloudstore_product'
               sh 'docker push navinshrinivas/cloudstore_product'
               sh 'docker build ./orders/ -t navinshrinivas/cloudstore_orders'
               sh 'docker push navinshrinivas/cloudstore_orders'
               sh 'docker build ./cloudstore_site/ -t navinshrinivas/cloudstore_site'
               sh 'docker push navinshrinivas/cloudstore_site'
            }
        }

        stage("deploy"){
            steps{
            //Assuming minikube to be running
               sh 'kubectl apply -f ./userhandle/user_deployment.yaml'
               sh 'kubectl rollout restart deployment userhandle'

               sh 'kubectl apply -f ./products/products_deployment.yaml'
               sh 'kubectl rollout restart deployment products'

               sh 'kubectl apply -f ./orders/order_deployment.yaml'
               sh 'kubectl rollout restart deployment orders'

               sh 'kubectl apply -f ./cloudstore_site/website_deployment.yaml'
               sh 'kubectl rollout restart deployment website'

               sh 'kubectl apply -f cloudstore_ingress.yaml'
            }
        }
    }
    post{
      failure{
         echo "Pipeline failed"
      }
    }
}
