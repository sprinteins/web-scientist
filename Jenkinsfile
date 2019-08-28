pipeline {
    agent { docker { image 'golang' } }
    environment {
        GOPATH = "${PWD}"
        XDG_CACHE_HOME = "${GOPATH}/tmp"
    }
    stages {
        stage('Initialize') {
            steps {
                dir("src") {
                    sh 'go get'
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
