pipeline {
    agent any

    environment {
        REPO_URL = 'https://github.com/spaciery/mindmirror.git'
        IMAGE_NAME = 'mindmirror-go-app'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'master', url: "${REPO_URL}"
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Build the Docker image
                    docker.build("${IMAGE_NAME}")
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    // Run tests inside the Docker container
                    docker.image("${IMAGE_NAME}").inside {
                        sh 'go test ./...'
                    }
                }
            }
        }

        stage('Run Application') {
            steps {
                script {
                    // Run the application in a Docker container
                    docker.image("${IMAGE_NAME}").run('-p 8081:8081')
                }
            }
        }
    }
}
