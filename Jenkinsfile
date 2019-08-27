pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Installation') {
            steps {
                sh 'cd src && go get'
            }
        }
        stage('Test') {
            steps {
                sh 'cd src/server && go test'
            }
        }
    }
}
