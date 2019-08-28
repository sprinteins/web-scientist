pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Test') {
            steps {
                sh 'apt-get install sudo'
                sh 'cd src/server && sudo go test'
            }
        }
    }
}
