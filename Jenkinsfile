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
               bash'''#!/bin/bash 
                  cd ./products
                  pwd
                  ls
                  docker build . -t navinshrinivas/products
               '''


               sh 'cd ../orders'
               sh 'docker build . -t navinshrinivas/orders'
            }
        }

        stage("deploy"){
            steps{
            //Assuming minikube to be running
               sh 'cd userhandle'
               sh 'kubectl apply -f user_deployment.yaml'
               sh 'kubectl rollout restart deployment userhandle'

               sh 'cd ../products'
               sh 'kubectl apply -f products_deployment.yaml'
               sh 'kubectl rollout restart deployment products'

               sh 'cd ..'
               sh 'kubectl apply -f cloudstore_ingress.yaml'
            }
        }
        stage("test"){
            steps{
               sh './hello'
            }
        }
    }
    post{
      failure{
         echo "Pipline failed"
      }
    }
}