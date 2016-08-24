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
		"https://597304777786.dkr.ecr.eu-west-1.amazonaws.com",
		"AWS",
		"AQECAHh+dS+BlNu0NxnXwowbILs115yjd+LNAZhBLZsunOxk3AAAAvEwggLtBgkqhkiG9w0BBwagggLeMIIC2gIBADCCAtMGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMd90K7Qe01TJN2HZwAgEQgIICpIcbw6UQlHWwwxB/ibJV2t+rAuCfWLL41sRKdfjrpHTC+/YVkMA3679bCqsg2XPE73i6offn5DLW5/KjljySpFvq0jK/8n0su9xSki1NV4D6OBhpBCY4wMQXc+qOXKVj1/5BYJoo1TDE8hxVf946fkAQctXfUEN9vdnxreknPGaNC1UlIlLGZ2xhPiZ0t7cjiXzk1wpwsocf39SSHHF7LYbrEkMFhEXBsTqCP7zx2uo9AFonYMFA6fqmWELyansPVvXytRHF3g7iMY/8++uWLMWdsq9I/Waproosphzl4fUSmOobzomp09x1eOfZvdyI4WRKGFFsHG4T4mWB1vSbJIDVl0lr64LcwjauE0IlcXgtdPvv5lMAjz+tqdZ6MnqQkZMObXLx8xnzyHRjnTx3SwmBTfHatFVeK9RC3pW9DTIYfStVTgZOdbDgZeuadkWD8fJqW+hH3XtCkOnBnDaeyOTnWnnS+ANWfLaUy2tWFRmDnlcOOkB2riNA70hYkmj+aXRJF52pi6HEDMUqp+zhCSGfew82FSKc9nFDAX2pyod1M9gxuwJC6FReyiimyHH7jT+qiVSvntPMBABWZWuZBbH7/iIZKPIsiHth//7bS71pMwim5HjJLPB+WFHuC1/xEbwkQmzmbpJWFmWDyST6H6VuCHUxes/mZKoSycWTtDqPwTyyU54xvSiIX/zT3XI8AXxdtDp5afbkADhwy9950DO7P1K5XxBKGPYzB18TLu7dytHBDp8HCt5YaxAVTNqR7VJZjUGyt3uXCCunJMO/Ktkq8xN53u+L8X3JiBsh+G90DqS7A5DEEcpTlkztT/OYCERYVblLFZYOFpBARSaea4NXHLjUT5wHlniAZ4AO76u6M4hV4bNJYDydyKbM6isIazXuYIE=",
	)
}
