pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Setup') {
            steps {
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
