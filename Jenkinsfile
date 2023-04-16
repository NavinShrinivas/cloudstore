pipeline {
    agent any
    stages {
        stage('cloning') {
            steps {
               git branch : "main",
               url : "https://github.com/NavinShrinivas/PES2UG20CS237_Jenkins"
            }
        }
        stage("build"){
            steps{
               sh 'rm -r hello_exec'
               sh 'make'
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
