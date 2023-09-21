package mails

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkhoa145/go-training/config"
)

func TestSendEmailWithGmail(t *testing.T) {
	sender := NewGmailSender(config.LoadConfig().EmailSenderName, config.LoadConfig().EmailSenderAddress, config.LoadConfig().EmailSenderPassword)
	subject := "a test email"
	content := `
	<h1>Hello World</h1>
	<p>This is a test message from <a href="http://google.com">Link</a></p>
	`
	to := []string{"khoavodang1451997@gmail.com"}
	// attachFiles := []string{"../../README.md"}
	err := sender.SendMail(subject, content, to, nil, nil, nil)
	require.NoError(t, err)
}
