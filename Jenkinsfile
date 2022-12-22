pipeline {
    agent any 
    stages {
        stage('Build') {
            steps {
                echo 'Build app'
                sleep(5)
                echo 'Test Trigger'
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