pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Test') {
            steps {
                sh 'cd src/server && go test'
            }
        }
    }
}
