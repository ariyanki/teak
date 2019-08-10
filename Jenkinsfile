pipeline {
    agent any
    environment {
        TEAK_STAGING_HOST=credentials('TEAK_STAGING_HOST')
        TEAK_STAGING_USER=credentials('TEAK_STAGING_USER')
        SONAR_GITHUB_TOKEN=credentials('SONAR_GITHUB_TOKEN')
        SONAR_URL=credentials('SONAR_URL')
        
        DOCKER_STAGING_USERNAME=credentials('DOCKER_STAGING_USERNAME')
        DOCKER_STAGING_PASSWORD=credentials('DOCKER_STAGING_PASSWORD')
        
        CONFIG_STAGING_FILE=credentials('CONFIG_STAGING_FILE')
        KUBESTAGING=credentials('KUBESTAGING')

    }
    stages {
        stage('Prepare') {
            steps {
                script {
                    def root = tool name: 'Go 1.12', type: 'go'

                    withEnv(["GOPATH=${env.WORKSPACE}/go", "GOROOT=${root}", "GOBIN=${root}/bin", "PATH+GO=${root}/bin", "TEAK_APP_PATH=${env.WORKSPACE}/go/src/teak", "TEAK_ENV=testing"]) {
                        sh 'go get -u github.com/golang/dep/cmd/dep'
                        sh "mkdir -p ${env.WORKSPACE}/go/src/teak"
                        sh 'cp -r $(ls --ignore=go) ${TEAK_APP_PATH}'
                        sh 'cd ${TEAK_APP_PATH} && cp cfg/example.json cfg/testing.json'
                        sh 'cd ${TEAK_APP_PATH} && dep ensure'
                        sh 'cd ${TEAK_APP_PATH} && dep ensure -update'
                    }
                }
            }
        }
        stage('Test') {
            steps {
                script {
                    def root = tool name: 'Go 1.12', type: 'go'

                    withEnv(["GOPATH=${env.WORKSPACE}/go", "GOROOT=${root}", "GOBIN=${root}/bin", "PATH+GO=${root}/bin", "TEAK_APP_PATH=${env.WORKSPACE}/go/src/teak", "TEAK_ENV=testing"]) {
                        sh 'cd ${TEAK_APP_PATH} && go test ./... -coverprofile=coverage.out'
                    }
                }
            }
        }
        stage('Sonar Scanner') {
            steps {
                script {
                    def root = tool name: 'Go 1.12', type: 'go'

                    withEnv(["GOPATH=${env.WORKSPACE}/go", "GOROOT=${root}", "GOBIN=${root}/bin", "PATH+GO=${root}/bin", "TEAK_APP_PATH=${env.WORKSPACE}/go/src/teak", "TEAK_ENV=testing"]) {
                        sh 'cd ${TEAK_APP_PATH} && sonar-scanner -Dsonar.host.url=$SONAR_URL -Dsonar.login=$SONAR_GITHUB_TOKEN'
                    }
                }
            }
        }
        stage('Deploy Staging') {
            when {
                beforeAgent true
                branch 'qa'
            }
            parallel {
                stage('Staging Server') {
                    steps {
                        script {
                            def root = tool name: 'Go 1.12', type: 'go'

                            withEnv(["GOPATH=${env.WORKSPACE}/go", "GOROOT=${root}", "GOBIN=${root}/bin", "PATH+GO=${root}/bin", "TEAK_APP_PATH=${env.WORKSPACE}/go/src/teak", "TEAK_ENV=testing"]) {
                                sh 'cd ${TEAK_APP_PATH} && go build -o teak_new'
                                sh 'ssh-keyscan -H $TEAK_STAGING_HOST >> ~/.ssh/known_hosts'
                                sh 'cd ${TEAK_APP_PATH} && scp swagger/teak.yaml $TEAK_STAGING_USER@$TEAK_STAGING_HOST:/ariyanki/teak/docs/http'
                                sh 'cd ${TEAK_APP_PATH} && scp teak_new $TEAK_STAGING_USER@$TEAK_STAGING_HOST:/ariyanki/teak'
                                sh 'cd ${TEAK_APP_PATH} && ssh -v $TEAK_STAGING_USER@$TEAK_STAGING_HOST \'bash -s\' < script/deploy_staging.sh'
                            }
                        }
                    }
                }
                stage('Staging k8s') {
                    steps {
                        script {
                            def root = tool name: 'Go 1.12', type: 'go'

                            withEnv(["GOPATH=${env.WORKSPACE}/go", "GOROOT=${root}", "GOBIN=${root}/bin", "PATH+GO=${root}/bin", "TEAK_APP_PATH=${env.WORKSPACE}/go/src/teak", "TEAK_ENV=testing"]) {
                                sh 'cd ${TEAK_APP_PATH} && cp $CONFIG_STAGING_FILE cfg/temp.json'
                                sh 'cd ${TEAK_APP_PATH} && env GOOS=linux go build'
                                sh 'cd ${TEAK_APP_PATH} && docker build -t teak:staging-$BUILD_NUMBER .'
                                sh 'cd ${TEAK_APP_PATH} && echo "$DOCKER_STAGING_PASSWORD" | docker login harbor.ariyanki.dev -u "$DOCKER_STAGING_USERNAME" --password-stdin'
                                sh 'cd ${TEAK_APP_PATH} && docker tag teak:staging-$BUILD_NUMBER harbor.ariyanki.dev/teak/teak:staging-$BUILD_NUMBER'
                                sh 'cd ${TEAK_APP_PATH} && docker push harbor.ariyanki.dev/teak/teak:staging-$BUILD_NUMBER'

                                // sh 'cd ${TEAK_APP_PATH} && kubectl --namespace=orca --kubeconfig $KUBESTAGING set image deployment/teak-apps teak-apps=harbor.ariyanki.dev/teak/teak:$TEAK_BUILD_NUMBER'
                            }
                        }
                    }
                }
            }
        }
        stage('Deploy Production') {
            when {
                beforeAgent true
                branch 'master'
            }
            steps {
                script {
                    def root = tool name: 'Go 1.12', type: 'go'

                    withEnv(["GOPATH=${env.WORKSPACE}/go", "GOROOT=${root}", "GOBIN=${root}/bin", "PATH+GO=${root}/bin", "TEAK_APP_PATH=${env.WORKSPACE}/go/src/teak", "TEAK_ENV=testing"]) {
                        sh 'cd ${TEAK_APP_PATH} && go build -o teak_new'
                        sh 'echo production-script-here'
                    }
                }
            }
        }
        stage('Release Binary') {
            when {
                beforeAgent true
                buildingTag()
                tag pattern: "^[0-9.]+\$", comparator: "REGEXP"
            }
            steps {
                script {
                    def root = tool name: 'Go 1.12', type: 'go'

                    withEnv(["GOPATH=${env.WORKSPACE}/go", "GOROOT=${root}", "GOBIN=${root}/bin", "PATH+GO=${root}/bin", "TEAK_APP_PATH=${env.WORKSPACE}/go/src/teak", "TEAK_ENV=testing"]) {
                        // Build Linux amd64
                        sh 'cd ${TEAK_APP_PATH} && mkdir -p "teak-${TAG_NAME}-linux-amd64"/cfg'
                        sh 'cd ${TEAK_APP_PATH} && env GOOS=linux GOARCH=amd64 go build -o teak-${TAG_NAME}-linux-amd64/teak'
                        sh 'cd ${TEAK_APP_PATH} && cp cfg/example.json teak-${TAG_NAME}-linux-amd64/cfg/config.json'
                        sh 'cd ${TEAK_APP_PATH} && tar -czvf teak-${TAG_NAME}-linux-amd64.tar.gz teak-${TAG_NAME}-linux-amd64'
                        sh 'cd ${TEAK_APP_PATH} && scp teak-${TAG_NAME}-linux-amd64.tar.gz $TEAK_STAGING_USER@$TEAK_STAGING_HOST:/ariyanki/teak/download'
                    }
                }
            }
        }
    }
}
