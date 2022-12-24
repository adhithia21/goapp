pipeline {
    agent {
        label 'docker'
    }
    stages {
        stage('Build') {
            steps {
                echo 'Build app'
                sh 'docker build -t asia.gcr.io/studidevops-369306/goapp:${BUILD_NUMBER} .'
            }
        }
        stage('Push GCR'){
            steps {
                echo "Push docker image to gcr"
                sh 'docker push asia.gcr.io/studidevops-369306/goapp:${BUILD_NUMBER}'
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