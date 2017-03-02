package llongdocker

import "testing"

func TestCanGetImageConfig(t *testing.T) {
	llong := llongClient()
	imageConfig, err := llong.GetImageConfig("597304777786.dkr.ecr.eu-west-1.amazonaws.com/docker-php", "latest")
	if err != nil {
		t.Errorf("Error getting image config. " + err.Error())
	}
	if imageConfig.AppName == "" {
		t.Errorf("Image config App name is empty")
	}
}

func TestCanGetAListOfImages(t *testing.T) {
	llong := llongClient()
	imageHistory, err := llong.GetRepoImages("docker-php")
	if err != nil {
		t.Errorf("Error getting image history. " + err.Error())
	}
	if len(imageHistory.ImageIds) == 0 {
		t.Errorf("Error getting image history. Array is empty")
	}
}

func llongClient() *LlongDockerClient {
	return NewLlongDockerClient(
		"eu-west-1",
		"unix:///var/run/docker.sock",
		"",
		"AWS",
		""
	)
}
