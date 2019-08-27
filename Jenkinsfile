pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Test') {
            steps {
                sh 'cd src/server'
                sh 'go test'
            }
        }
    }
}
