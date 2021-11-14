# JFrog Pipelines and Jenkins Example

This is a repo that is used for the blog [_Jenkins and JFrog Pipelines: Working Together to Release Your Software_](https://jfrog.com/blog/jenkins-and-jfrog-pipelines-ci-cd/). The blog shows how to use Jenkins and JFrog Pipelines together to orchestrate your code from build to release.

For this example, we have built a Go REST application that we want to build, run unit tests and then a Docker push to a staging repository. Jenkins will handle this part. Next, JFrog Pipelines will deploy your Docker Go application from the staging repository to a Kubernetes cluster. We will use Google Kubernetes Engine (GKE). Additionally, we will use Artifactory as our Docker registry. This makes it easy to promote the build to release without pushing the same build to another release registry. The Jenkins pipeline is defined in the [Jenkinsfile](./Jenkinsfile) and the JFrog Pipeline is defined in the [pipeline.yml](./pipeline.yml).

![Diagram-all](https://user-images.githubusercontent.com/6440106/80759359-09d39e80-8aec-11ea-9beb-0334a81b2e50.png)

## Jenkins

![Jenkins](https://user-images.githubusercontent.com/6440106/80759523-4c957680-8aec-11ea-84d5-f29a9828cd1e.png)

## JFrog Pipelines

![JFrog Pipelines](https://user-images.githubusercontent.com/6440106/80759614-6b940880-8aec-11ea-9a9b-d1447b4dad81.png)
