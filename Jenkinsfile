pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Setup') {
            steps {
                dir("/.cache/go-build") {}
                dir("src") {
                    sh 'go get'
                }
            }
        }
        stage('Test') {
            steps {
                dir("server") {
                    sh 'go test'
                }
            }
        }
    }
}
