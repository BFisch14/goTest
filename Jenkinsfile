pipeline {
    agent any
    tools {
        go 'go-1.15'
    }

    environment {
        GO11MODULE = 'on'
        GOROOT = '${root}'
        PATH+GO = '${root}/bin'
    }

    

    stages {
        stage('Compile') {
            steps {
                sh 'go version'
            }
        }
    }
}
