pipeline {
    agent any

    environment {
        GOROOT = '/usr/local/go'          // 指定Go安装目录
        GOPATH = "${WORKSPACE}/go"          // 将GOPATH设置为当前工作区下的go文件夹
        PATH = "${GOROOT}/bin:${GOPATH}/bin:${env.PATH}"
    }

    stages {
        stage('Checkout') {
            steps {
                // 从Git仓库中拉取代码
                git 'https://github.com/wanghanfeng/go-demo.git'
            }
        }

        stage('Install Dependencies') {
            steps {
                // 安装依赖项
                sh 'go mod download'
            }
        }

        stage('Build') {
            steps {
                // 编译项目
                sh 'go build -o myapp main.go'
            }
        }

        // stage('Test') {
        //     steps {
        //         // 运行测试
        //         sh 'go test ./...'
        //     }
        // }

        stage('Package') {
            steps {
                // 打包应用
                sh 'tar -czf myapp.tar.gz myapp'
            }
        }

        stage('Archive Artifacts') {
            steps {
                // 存档生成的构件（可选）
                archiveArtifacts artifacts: 'myapp.tar.gz', followSymlinks: false
            }
        }
    }

    post {
        always {
            // 清理工作区
            cleanWs()
        }
        success {
            echo 'Build succeeded!'
        }
        failure {
            echo 'Build failed.'
        }
    }
}
