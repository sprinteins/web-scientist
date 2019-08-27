pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Test') {
            steps {
                sh 'sudo cd src/server && sudo go test'
            }
        }
    }
}
