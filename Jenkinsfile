pipeline {

  agent any

  stages {

    stage('Checkout Source') {
      steps {
        git 'https://github.com/letran3691/echo_gorm.git'
      }
    }
    stage('Deploy App') {
      steps {
        script {
          kubernetesDeploy(configs: "deploy_go_pvc.yml", kubeconfigId: "mykubeconfig")
        }
      }
    }

  }

}