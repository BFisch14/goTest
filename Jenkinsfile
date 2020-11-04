pipeline {
    agent any
    tools {
        go 'go-1.15'
    }

    environment {
        GO11MODULE = 'on'
        GOROOT = '${root}'
    }

    stages {
        stage('Compile') {
            steps {
                echo 'Hello world'
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                sh 'go test'
            }
        }
    }
}
