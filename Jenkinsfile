pipeline {
    agent any
    tools {
        go 'go-1.1.5'
    }

    environment {
        GO11MODULE = 'on'
    } 
    stages {
        stage('build') {
            steps {
                sh 'go build'
            }
        }
    }
}
