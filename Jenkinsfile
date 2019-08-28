pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Setup') {
            steps {
                dir("/.cache/go-build") {
                    sh 'pwd'
                }
                dir("src") {
                    sh 'pwd'
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
