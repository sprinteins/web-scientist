pipeline {
    agent { docker { image 'golang' } }
    withEnv(['GOPATH=${PWD}/server']) {
        stages {
            stage('Setup') {
                steps {
                    sh 'go get'
                }
            }
            stage('Test') {
                steps {
                    sh 'go version'
                }
            }
        }
    }
}
