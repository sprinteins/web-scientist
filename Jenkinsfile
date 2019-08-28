pipeline {
    agent { docker { image 'golang' } }
    stages {
        dir("src")
            stage('Setup') {
                steps {
                    sh 'go get'
                }
            }
            dir("server") {
                stage('Test') {
                    steps {
                        sh 'go test'
                    }
                }
            }
        }
    }
}
