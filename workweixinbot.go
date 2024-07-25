package wwxb

import (
	"context"
	"io"
	"net/http"
	"time"
)

type Message interface {
	Json() io.Reader
}

type WorkWeiXinBot struct {
	webHookUrl string
}

func NewBot(hook string) *WorkWeiXinBot {
	return &WorkWeiXinBot{webHookUrl: hook}
}

// Send
// 静默发送
func (w *WorkWeiXinBot) Send(msg Message) error {
	req, err := http.NewRequest("POST", w.webHookUrl, msg.Json())
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req.Header.Set("Content-Type", "application/json;charset-UTF-8")
	client := &http.Client{}
	_, err = client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	return nil
}
