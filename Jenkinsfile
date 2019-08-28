pipeline {
    agent { docker { image 'golang' } }
    environment {
        XDG_CACHE_HOME = "tmp"
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
                    sh 'ls ${XDG_CACHE_HOME}'
                    sh 'go test'
                }
            }
        }
    }
}
