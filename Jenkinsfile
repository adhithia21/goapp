pipeline {
    agent any 
    stages {
        stage('Build') {
            steps {
                echo 'Build app'
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