pipeline {
    agent { docker { image 'golang' } }
    environment {
        XDG_CACHE_HOME = "tmp"
    }
    stages {
        stage('Initialize') {
            steps {
                dir("src") {
                    sh 'go mod download'
                    sh 'go get'
                    sh 'go build'
                }
            }
        }
        stage('Test') {
            steps {
                sh 'env'
                dir("src/server") {
                    sh 'go test'
                }
            }
        }
    }
}
