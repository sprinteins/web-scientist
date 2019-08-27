pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Installation') {
            steps {
                sh 'cd src'
                sh 'go get'
            }
        }
        stage('Test') {
            steps {
                sh 'cd server'
                sh 'go test'
            }
        }
    }
}
