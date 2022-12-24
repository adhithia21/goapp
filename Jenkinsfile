pipeline {
    agent {
        label 'docker'
    }
    environment {
        DISCORD_NOTIFICATION = credentials('discord-alert-development')
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
                sh 'cat "$GCP_SERVICE_ACCOUNT" | docker login -u _json_key --password-stdin https://asia.gcr.io'
                sh 'docker push asia.gcr.io/studidevops-369306/goapp:${BUILD_NUMBER}'
            }
            post {
                success {
                    discordSend description: "Pushed image asia.gcr.io/studidevops-369306/goapp:${BUILD_NUMBER}", footer: "Build", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "${DISCORD_NOTIFICATION}"
                }
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
                KUBE_CONFIG = credentials('kubernetes-demo-jenkins')
            }
            steps {
                echo 'deploy with helm'
                sh 'helm repo add adhithia-charts https://adhithia21.github.io/helm-charts/charts'
                sh 'helm upgrade --kubeconfig "$KUBE_CONFIG" --install goapp adhithia-charts/application'
            }
        }
    }
}