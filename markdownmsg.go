package wwxb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

const (
	ModeH1 = iota + 1
	ModeH2
	ModeH3
	ModeH4
	ModeH5
	ModeH6
	ModeBold
	ModeQuote
	ModeInfo
	ModeComment
	ModeWarning
	ModeText
)

type MakeDown struct {
	Content string `json:"content"`
}

type MarkdownBody struct {
	MsgType  string   `json:"msgtype"`
	MakeDown MakeDown `json:"markdown"`
}

type MdMsg struct {
	contents []string
}

func NewMdMsg() *MdMsg {
	return &MdMsg{
		contents: make([]string, 0),
	}
}

func (m *MdMsg) Json() io.Reader {
	body := MarkdownBody{
		MsgType:  "markdown",
		MakeDown: MakeDown{Content: strings.Join(m.contents, "")},
	}

	buf, _ := json.Marshal(body)
	fmt.Println(string(buf))
	return bytes.NewReader(buf)
}

func (m *MdMsg) Add(field string, mode int, enter bool) {
	var msg string
	switch mode {
	case ModeH1:
		msg = fmt.Sprintf("# %s", field)
	case ModeH2:
		msg = fmt.Sprintf("## %s", field)
	case ModeH3:
		msg = fmt.Sprintf("### %s", field)
	case ModeH4:
		msg = fmt.Sprintf("#### %s", field)
	case ModeH5:
		msg = fmt.Sprintf("##### %s", field)
	case ModeH6:
		msg = fmt.Sprintf("###### %s", field)
	case ModeBold:
		msg = fmt.Sprintf("**%s**", field)
	case ModeQuote:
		msg = fmt.Sprintf("> %s", field)
	case ModeInfo:
		msg = fmt.Sprintf("<font color=\"info\">%s</font>", field)
	case ModeComment:
		msg = fmt.Sprintf("<font color=\"comment\">%s</font>", field)
	case ModeWarning:
		msg = fmt.Sprintf("<font color=\"warning\">%s</font>", field)
	case ModeText:
		msg = field
	default:
		msg = field
	}
	if enter {
		msg = "\n" + msg
		m.contents = append(m.contents, msg)
	} else {
		m.contents = append(m.contents, msg)
	}
}

func H1(field string) string {
	return fmt.Sprintf("# %s", field)
}

func H2(field string) string {
	return fmt.Sprintf("## %s", field)
}

func H3(field string) string {
	return fmt.Sprintf("### %s", field)
}

func H4(field string) string {
	return fmt.Sprintf("#### %s", field)
}

func H5(field string) string {
	return fmt.Sprintf("##### %s", field)
}

func H6(field string) string {
	return fmt.Sprintf("###### %s", field)
}

func Bold(field string) string {
	return fmt.Sprintf("**%s**", field)
}

func Quote(field string) string {
	return fmt.Sprintf("> %s", field)
}

func Info(field string) string {
	return fmt.Sprintf("<font color=\"info\">%s</font>", field)
}

func Comment(field string) string {
	return fmt.Sprintf("<font color=\"comment\">%s</font>", field)
}

func Warning(field string) string {
	return fmt.Sprintf("<font color=\"warning\">%s</font>", field)
}
