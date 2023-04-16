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
               sh 'cd userhandle'
               sh 'docker build -t . navinshrinivas/userhandle'

               sh 'cd products'
               sh 'docker build -t . navinshrinivas/products'

               sh 'cd orders'
               sh 'docker build -t . navinshrinivas/orders'
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
