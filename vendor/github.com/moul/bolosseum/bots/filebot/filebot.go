package filebot

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/moul/bolosseum/bots"
)

type FileBot struct {
	path string
	name string
}

func (b *FileBot) Name() string {
	return b.name
}

func (b *FileBot) SetName(name string) {
	b.name = name
}

func (b *FileBot) Path() string {
	return b.path
}

func (b *FileBot) SendMessage(msg bots.QuestionMessage) (*bots.ReplyMessage, error) {
	// marshal json
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	// execute script
	cmd := exec.Command(b.path, string(data))
	cmd.Stderr = os.Stderr
	body, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// parse json
	var response bots.ReplyMessage
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (b *FileBot) Start() error {
	return nil
}

func NewBot(path string) (*FileBot, error) {
	return &FileBot{
		path: path,
	}, nil
}
