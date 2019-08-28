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
                sh 'ls ${XDG_CACHE_HOME}/go-build/70'
                dir("src/server") {
                    sh 'go test'
                }
            }
        }
    }
}
