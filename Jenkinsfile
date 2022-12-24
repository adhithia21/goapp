pipeline {
    agent {
        label 'docker'
    }
    environment {
        GCP_SSH_KEY = credentials('gcp-ssh-private-key')
        DISCORD_NOTIFICATION = credentials('discord-alert-development')
        GCP_SERVICE_ACCOUNT = credentials('gcp_service_account')
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
        stage ('Activate GCP Account'){
            steps {
                echo 'active gcp account'
                sh 'ssh -o StrictHostKeyChecking=no -i "$GCP_SSH_KEY" trainer@34.101.80.191 "rm -rf ~/gcp-service-account.json"'
                sh 'scp -o StrictHostKeyChecking=no -i "$GCP_SSH_KEY" "$GCP_SERVICE_ACCOUNT" trainer@34.101.80.191:~/gcp-service-account.json'
                // sh 'ssh -o StrictHostKeyChecking=no -i "$GCP_SSH_KEY" trainer@34.101.80.191 "gcloud auth activate-service-account $(cat gcp-service-account.json | jq -r .client_email) --key-file=gcp-service-account.json"'
                // sh 'ssh -o StrictHostKeyChecking=no -i "$GCP_SSH_KEY" trainer@34.101.80.191 "gcloud auth list"'
                // sh 'ssh -o StrictHostKeyChecking=no -i "$GCP_SSH_KEY" trainer@34.101.80.191 "gcloud container clusters get-credentials demo-jenkins --zone australia-southeast1-b --project studidevops-369306"'
            }
        }
        stage('Deploy') {
            steps {
                echo 'deploy with helm'
                sh 'helm repo add adhithia-charts https://adhithia21.github.io/helm-charts/charts'
                sh 'helm upgrade --install goapp adhithia-charts/application'
            }
        }
    }
    post {
        always {
            echo 'clean'
            cleanWs()
        }
    }
}