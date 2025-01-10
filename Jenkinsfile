pipeline {
    agent {
        node {
            label 'new-agent'
        }
    }

    stages {
        stage('Continuous integration') {
            steps {
                git branch: 'main', url: 'https://github.com/fredericEducentre/jenkins_project_go.git'
            }
        }
        stage('controle de qualité') {
            steps {
                 sh '''
               sonar-scanner \
                    -Dsonar.projectKey=projetgo \
                    -Dsonar.sources=. \
                    -Dsonar.host.url=http://192.168.1.5:9000 \
                    -Dsonar.token=sqp_b09150c7c7e83946b182edde6505c79c95f4f3e6
                     '''
            }
        }
        stage('build') {
            steps {
                 sh "docker build . -t current_time_go"

            }
        }        
            
        
    }
}
