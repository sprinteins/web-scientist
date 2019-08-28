pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Setup') {
            steps {
                sh 'cd src'
                sh 'pwd'
            }
        }
        stage('Test') {
            steps {
                sh 'go version'
            }
        }
    }
}
