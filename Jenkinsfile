pipeline {
    agent any

    environment {
        GOROOT = '/usr/local/go'          // 指定Go安装目录
        GOPATH = "${WORKSPACE}/go"          // 将GOPATH设置为当前工作区下的go文件夹
        PATH = "${GOROOT}/bin:${GOPATH}/bin:${env.PATH}"
        http_proxy = "http://192.168.31.192:7890"
        https_proxy = "http://192.168.31.192:7890"
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

        stage('Deploy to Server') {
            steps {
                // 使用 SSH 将二进制文件传输到远程服务器
                sshPublisher(
                    publishers: [
                        sshPublisherDesc(
                            configName: 'dev-machine-config',
                            transfers: [
                                sshTransfer(
                                    sourceFiles: 'myapp.tar.gz',
                                    remoteDirectory: './',
                                    execCommand: ''' 
                                        cd /root/user-bin && 
                                        tar -xzf myapp.tar.gz && 
                                        systemctl restart myapp.service
                                    '''  // 解压并重启服务
                                )
                            ]
                        )
                    ]
                )
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
