pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Setup') {
            steps {
                sh 'pwd'
                dir("/.cache/go-build")
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
