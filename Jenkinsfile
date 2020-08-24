pipeline {
    agent any
    tools {
        go 'go-1.15'
    }

    environment {
        GO11MODULE = 'on'
    } 
    stages {
        stage('Compile') {
            steps {
                sh 'go version'
            }
        }
    }
}
