pipeline{
    agent { label 'kubepod'}

    stages {
        stage("checkout git"){
            step{
                git url "https://github.com/letran3691/echo_gorm.git"
            }
        }
        stage("Deployment"){
            script{
                kubernetesDeploy(configs:"deploy_go_pvc.yml", kubeconfiId: "mykubeconfig")
            }
        }

    }

}