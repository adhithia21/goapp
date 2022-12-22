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
                    GOCACHE=/tmp/ GOOS=linux GOARCH=386 go build -o goapp main.go
                '''
                stash includes: 'goapp', name: 'GOAPP_ARTIFACT'
            }
        }
        stage('Test') {
            steps {
                echo 'Test app' 
                sleep(5)
            }
        }
        stage('Deploy') {
            environment {
                GCP_SSH_KEY = credentials('gcp-ssh-private-key')
                DISCORD_NOTIFICATION = credentials('discord-alert-development')
            }
            steps {
                echo 'Deploy app'
                unstash name: 'GOAPP_ARTIFACT'
                sh 'scp -o StrictHostKeyChecking=no -i "$GCP_SSH_KEY" goapp trainer@34.123.135.41:~/goapp'
            }
        }
    }
}