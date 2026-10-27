pipeline {
    agent any

    environment {
        DOCKERHUB_REGISTRY = "https://index.docker.io/v1/"
    }

    stages {
        stage('Build') {
            steps {
                script {
                    docker.withRegistry("${env.DOCKERHUB_REGISTRY}", "dockerhub") {
                        def COMMIT_HASH = sh(
                            script: 'git rev-parse --short HEAD',
                            returnStdout: true
                        ).trim()

                        env.IMAGE_TAG = "${COMMIT_HASH}-${env.BUILD_NUMBER}"

                        def imageName = "tarof429/mail-client:${env.IMAGE_TAG}"

                        def myImage = docker.build(imageName, "-f docker/Dockerfile .")

                        myImage.push()
                    }
                
                }
            }
        }
    }
}
