pipeline {
    agent any 
    stages {
        stage('Build') {
            agent {
                docker {
                    image 'golang:alpine3.16'
                    label 'docker'
                }
            }
            steps {
                echo 'Build app'
                sh '''#/bin/sh
                    go version
                    GOOS=linux GOARCH=386 go build -o golang-sample-linux386 main.go
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