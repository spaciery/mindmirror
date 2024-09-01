pipeline {
    agent any
    
    environment {
        REPO_URL = 'https://github.com/spaciery/mindmirror.git'
        IMAGE_NAME = 'mindmirror-go-app'
    }
    
    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'master', url: "${env.REPO_URL}"
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    sh "docker build -t ${env.IMAGE_NAME} ."
                }
            }
        }
        
      /*  stage('Run Tests') {
            steps {
                script {
                    sh "docker run --rm ${env.IMAGE_NAME} go test ./..."
                }
            }
        } */
        
        stage('Run Application') {
            steps {
                script {
                    sh "docker run -d -p 8081:8081 ${env.IMAGE_NAME}"
                }
            }
        }
    }
    
    post {
        always {
            sh "docker stop \$(docker ps -q --filter ancestor=${env.IMAGE_NAME}) || true"
            sh "docker rm \$(docker ps -aq --filter ancestor=${env.IMAGE_NAME}) || true"
        }
    }
}
