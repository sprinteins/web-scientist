pipeline {
    agent { docker { image 'golang' } }
    environment {
        XDG_CACHE_HOME = "tmp"
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
