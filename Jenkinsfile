pipeline {
    agent { docker { image 'golang' } }
    environment {
        XDG_CACHE_HOME = "${GOPATH}/tmp/"
    }
    stages {
        stage('Test') {
            steps {
                dir("src/server") {
                    sh 'go test'
                }
            }
        }
    }
}
