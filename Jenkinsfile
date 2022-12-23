pipeline {
    agent {
     labels 'docker && helm'   
    }
    stages {
        stage('Build') {
            steps {
                echo 'Build app'
                sh 'docker build -t asia.gcr.io/studidevops-369306/goapp:${BUILD_NUMBER}' .
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
                echo 'Deploy' 
                sleep(5)
            }
        }
    }
}
