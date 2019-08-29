pipeline {
    agent { docker { image 'golang' } }
    environment {
        XDG_CACHE_HOME = "${WORKSPACE}/tmp"
    }
    stages {
        stage('Initialize') {
            steps {
                dir("src") {
                    sh 'go mod download'
                }
            }
        }
        stage('Test') {
            steps {
                dir("src/server") {
                    sh 'go test'
                }
            }
        }
    }
}
