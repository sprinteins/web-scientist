pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Setup') {
            steps {
                dir("test") {
                    sh 'pwd'
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
