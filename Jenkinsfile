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
            dir("server") {
                steps {
                    sh 'go test'
                }
            }
        }
    }
}
