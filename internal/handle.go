package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"openAi/log"
	"strings"
)

// 这里填写 api key
const Authorization = "Bearer <拿key替换我>"
const OpenAiApiAddr = "https://api.openai.com/v1/completions"

func Handle() {
	var prompt string
	logs.InfoF("请提问：")
	//控制台的输出
	fmt.Scan(&prompt)
	bf := ReqData(prompt)
	respData := RespType{}
	err := json.Unmarshal(bf.Bytes(), &respData)
	if err != nil {
		logs.Error(fmt.Sprintf("resp data json unmarshl err:%s", err))
		return
	}
	if len(respData.Choices) <= 0 {
		return
	}
	logs.AI(fmt.Sprintf("AI:"))
	for _, data := range respData.Choices {
		logs.AI(fmt.Sprintf("【%s】", data.Text))
	}
}

// 请求类型
type ReqType struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
}

type RespType struct {
	Id      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Choices []Choices `json:"choices"`
	Usage   Usage     `json:"usage"`
}

type Choices struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func ReqData(prompt string) (bf *bytes.Buffer) {

	reqData := ReqType{
		Model:            "text-davinci-003",
		Prompt:           prompt,
		Temperature:      0.3,
		MaxTokens:        500,
		TopP:             1.0,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.0,
	}
	bytesData, err := json.Marshal(reqData)
	if err != nil {
		logs.Error(fmt.Sprintf(err.Error()))
		return
	}
	reader := strings.NewReader(string(bytesData))
	req, err := http.NewRequest(http.MethodPost, OpenAiApiAddr, reader)

	if err != nil {
		logs.Error(fmt.Sprintf(err.Error()))
		return
	}
	req.Header.Add("Authorization", Authorization)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		logs.Error(fmt.Sprintf(err.Error()))
		return
	}

	bf = new(bytes.Buffer)
	defer resp.Body.Close()
	_, err = io.Copy(bf, resp.Body)
	if err != nil {
		logs.Info(fmt.Sprintf(err.Error()))
		return
	}
	return
}
