package llongdocker

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	docker "github.com/fsouza/go-dockerclient"
)

const imageConfigKey string = "uk.co.oneiota-config"

// ImageConfig - which is on the image label
type ImageConfig struct {
	AppName        string          `json:"applicationName"`
	AppDescription string          `json:"applicationDescription"`
	HostPort       int             `json:"hostPort"`
	ContainerPort  int             `json:"containerPort"`
	Dependencies   map[string]bool `json:"dependencies"`
	Image          string
	Env            string
}

// LlongDockerClient - client to do all the things
type LlongDockerClient struct {
	dockerCli      *docker.Client
	ecrSvs         *ecr.ECR
	dockerAuth     docker.AuthConfiguration
	localEndpoint  string
	remoteEndpoint string
}

// NewLlongDockerClient - returns a new instance of a LlongDockerClient
func NewLlongDockerClient(awsRegion string, localEndpoint string, remoteEnpoint string, awsUsername string, awsPassword string) *LlongDockerClient {
	ecrSvs := ecr.New(session.New(), &aws.Config{Region: aws.String(awsRegion)})
	var err error
	dockerCli, err := docker.NewClient(localEndpoint)
	if err != nil {
		panic(err)
	}
	dockerAuth := docker.AuthConfiguration{
		Username:      awsUsername,
		Password:      awsPassword,
		ServerAddress: remoteEnpoint,
	}
	return &LlongDockerClient{
		dockerCli:  dockerCli,
		ecrSvs:     ecrSvs,
		dockerAuth: dockerAuth,
	}
}

// Docker -
type Docker interface {
	GetRepoImages(repoName string) *ecr.ListImagesOutput
	GetImageConfig(image string) *ImageConfig
	getImageLabels(image string, tag string) map[string]string
}

// GetRepoImages - wrapper for aws ecs to get repo images.
func (llong *LlongDockerClient) GetRepoImages(repoName string) (*ecr.ListImagesOutput, error) {
	listImagesImp := ecr.ListImagesInput{
		RepositoryName: &repoName,
	}
	resp, err := llong.ecrSvs.ListImages(&listImagesImp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// GetImageConfig - returns image config from the label
func (llong *LlongDockerClient) GetImageConfig(image string, tag string) (ImageConfig, error) {
	var retErr error
	retIC := ImageConfig{}
	tags, err := llong.getImageLabels(image, tag)
	if err != nil {
		return retIC, err
	}
	if cStr, ok := tags[imageConfigKey]; ok {
		cByte := []byte(cStr)
		retErr := json.Unmarshal(cByte, &retIC)
		if retErr != nil {
			return retIC, retErr
		}
	}
	return retIC, retErr
}

func (llong *LlongDockerClient) getImageLabels(image string, tag string) (map[string]string, error) {
	returnMap := make(map[string]string)
	fullImagePath := image + ":" + tag
	pullOpts := docker.PullImageOptions{
		Repository: image,
		Tag:        tag,
	}
	err := llong.dockerCli.PullImage(pullOpts, llong.dockerAuth)
	if err != nil {
		return returnMap, err
	}
	imageMeta, err := llong.dockerCli.InspectImage(fullImagePath)
	if err != nil {
		return returnMap, err
	}
	go llong.dockerCli.RemoveImageExtended(fullImagePath, docker.RemoveImageOptions{Force: true})
	return imageMeta.Config.Labels, err
}
