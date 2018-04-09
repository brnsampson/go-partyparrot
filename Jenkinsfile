#!/usr/bin/env groovy

// this will start an executor on a Jenkins agent with the docker label
pipeline {
    agent any
    // Setup tools
    tools {
        go 'go 1.10.1'
    }
    // Setup variables
    environment {
        // application name will be used in a few places so create a variable and use string interpolation to use it where needed
        applicationName = "go-partyparrot" // TODO replace this with repository name
        // a basic build number so that when we build and push to Artifactory we will not overwrite our previous builds
        // buildNumber = "0.1.${env.BUILD_NUMBER}"
        // a build number based on git tag so that when we build and push to s3 we will not overwrite our previous builds
        buildVersion = "${env.TAG.substring(1)}"
        // Path we will mount the project to for the Docker container
        GOPATH = "${env.WORKSPACE}"
        // You will need the credential plugin for this pipeline. You'll also need to create a matching global credential, of course.
        APIKEY = credentials("octopus-deploy-api-key")
        // This is just what my build server happens to be named currently. To be fixed...
        OCTOSERVER = "DESKTOP-ENP9NLA"
    }
    
    stages {
        // Clean out the workspace.
        stage("Clean workspace") {
            steps {
                deleteDir()
            }
        }

        // Checkout the code from Github, stages allow Jenkins to visualize the different sections of your build steps in the UI
        stage('Checkout from GitHub') {
            // No special needs here, if your projects relys on submodules the checkout step would need to be different
            steps {
                checkout scm
            }
        }

        // Start a build using golang:1.8.0 by mounting the current directory to the goPath we specified earlier
        stage("Build") {
            steps {
                // build the Linux x64 binary
                sh "go get github.com/aws/aws-lambda-go/events"
                sh "go get github.com/aws/aws-lambda-go/lambda"
                sh "go get github.com/brnsampson/go-partyparrot/partyparrot"
                sh "GOOS=linux GOARCH=amd64 go build -o build/${env.buildVersion}/${env.applicationName}.${env.buildVersion} ./pplambda"
                sh "chmod +x build/${env.buildVersion}/${env.applicationName}.${env.buildVersion}"
            }
        }
    
        stage("Zip") {
            steps {
                dir("build/${env.buildVersion}/") {
                    zip zipFile: "${env.applicationName}.${env.buildVersion}.zip", glob: "${env.applicationName}.${env.buildVersion}"
                }
            }
        }

        stage("Push to S3") {
            steps {
                s3Upload consoleLogLevel: 'INFO', dontWaitForConcurrentBuildCompletion: false, entries: [[bucket: 'whobe-deploy', pathStyleAccessEnabled: true, path: "${env.applicationName}.${env.buildVersion}.zip", excludedFile: '', flatten: false, gzipFiles: false, keepForever: false, managedArtifacts: false, noUploadOnFailure: true, selectedRegion: 'us-west-2', showDirectlyInBrowser: false, sourceFile: "build/${env.buildVersion}/${env.applicationName}.${env.buildVersion}.zip", storageClass: 'STANDARD', uploadFromSlave: false, useServerSideEncryption: false]], pluginFailureResultConstraint: 'FAILURE', profileName: 'Jenkins', userMetadata: []
            }
        }
    }
}