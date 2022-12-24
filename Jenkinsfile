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
            environment {
                GCP_SERVICE_ACCOUNT = credentials('gcp_service_account')
            }
            steps {
                echo "Push docker image to gcr"
                sh 'cat $GCP_SERVICE_ACCOUNT | docker login -u _json_key --password-stdin https://sia.gcr.io'
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