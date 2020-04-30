pipeline {
  agent {
    kubernetes {
        label 'go-pipeline-pod'
        yamlFile 'podTemplate/go-pipeline-pod.yaml'
        idleMinutes 120
    }
  }
  stages {
    stage('Build') {
        steps {
            container('golang'){
                sh 'go build'
            }
        }
    }
    stage('Unit Tests') {
        steps {
            container('golang'){
                sh 'go test ./... -run Unit'
            }
        }
    }
    stage('Docker Build') {
      steps {
        container('docker'){
            sh "docker build -t partnership-public-images.jfrog.io/goci-example:latest ."
        }
      }
    }
    stage('Docker Push to Repo') {
      steps {
        container('docker'){
            script {
              docker.withRegistry( 'https://partnership-public-images.jfrog.io', 'gociexamplerepo' ) {
                sh "docker push partnership-public-images.jfrog.io/goci-example:latest"
              }
           }
        }
      }
    }
    stage('Publish Build Info') {
      environment {
        JFROG_CLI_OFFER_CONFIG = false
      }
      steps {
        container('jfrog-cli-go'){
            withCredentials([usernamePassword(credentialsId: 'gociexamplerepo', passwordVariable: 'APIKEY', usernameVariable: 'USER')]) {
                sh "jfrog rt bce $JOB_NAME $BUILD_NUMBER"
                sh "jfrog rt bag $JOB_NAME $BUILD_NUMBER"
                sh "jfrog rt bad $JOB_NAME $BUILD_NUMBER \"go.*\""
                sh "jfrog rt bp --build-url=https://jenkins.openshiftk8s.com/ --url=https://partnership.jfrog.io/artifactory --user=$USER --apikey=$APIKEY $JOB_NAME $BUILD_NUMBER"
            }
        }
      }
    }
  }
  post {
      success {
        script {
           sh "curl -v -XPOST -H \"authorization: Basic amVmZmY6amZyMGdqM25rMW5z\" \"https://partnership-pipelines-api.jfrog.io/v1/projectIntegrations/17/hook\" -d '{\"buildName\":\"$JOB_NAME\",\"buildNumber\":\"$BUILD_NUMBER\",\"buildInfoResourceName\":\"jenkinsBuildInfo\"}' -H \"Content-Type: application/json\""
        }
      }
  }
}