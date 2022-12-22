pipeline {
    agent any 
    stages {
        stage('Build') {
            agent {
                docker {
                    image 'golang:alpine3.16'
                }
            }
            steps {
                echo 'Build app'
                sh '''#/bin/sh
                    go version
                '''
            }
        }
        stage('Test') {
            steps {
                echo 'Test app' 
                sleep(5)
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploy app' 
                sleep(5)
            }
        }
    }
}