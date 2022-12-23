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
        stage('Push') {
            environment {
                GCP_SERVICE_ACCOUNT = credentials('gcp_service_account')
            }
            steps {
                echo 'Build app'
                sh 'cat "$GCP_SERVICE_ACCOUNT" | docker login -u _json_key --password-stdin https://asia.gcr.io'
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
            environment {
                KUBE_CONFIG = credentials('kubernetes-config')
            }
            steps {
                echo 'Deploy'
                sh 'helm repo add adhithia-charts https://adhithia21.github.io/helm-charts/charts'
                sh 'helm repo update adhithia-charts'
                sh 'helm search repo adhithia-charts/'
                sh 'helm upgrade --kubeconfig "$KUBE_CONFIG" --install goapp adhithia-charts/application --set replicas=1 --set image=asia.gcr.io/studidevops-369306/goapp:${BUILD_NUMBER}'
            }
        }
    }
}
