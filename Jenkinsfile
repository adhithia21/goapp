pipeline {
    agent any 
    stages {
        stage('Build') {
            agent {
                label 'docker'
            }
            steps {
                echo 'Build app'
                sh 'docker build -t asia.gcr.io/studidevops-369306/goapp:${BUILD_NUMBER}'
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
                echo 'deploy'
            }
        }
    }
}