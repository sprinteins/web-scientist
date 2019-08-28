pipeline {
    agent { docker { image 'golang' } }
    stages {
        withEnv(['GOPATH=${PWD}/server']) {
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
