/*
 * Copyright 2012-2019 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * A toolkit for Golang development
 * https://www.likexian.com/
 */

package xmailer

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"path/filepath"
	"strings"
	"time"
)

// Attachment storing mail attachment
type Attachment struct {
	Name    string
	Content []byte
}

// Auth storing mail auth
type Auth struct {
	Server string
	Auth   smtp.Auth
}

// Message storing mail message
type Message struct {
	From        string
	To          []string
	Cc          []string
	Bcc         []string
	Subject     string
	Body        string
	ContentType string
	Attachments map[string]*Attachment
	Auth        *Auth
}

// Version returns package version
func Version() string {
	return "0.4.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// New returns a new mailer
func New(server, username, password string, ishtml bool) (m *Message) {
	m = &Message{
		From:        username,
		To:          []string{username},
		Subject:     "Mailer Test",
		Body:        "Hello World. This is mailer via github.com/likexian/gokit/xmailer.",
		Attachments: make(map[string]*Attachment),
	}

	if ishtml {
		m.ContentType = "text/html"
	} else {
		m.ContentType = "text/plain"
	}

	servers := strings.Split(server, ":")
	m.Auth = &Auth{
		Server: server,
		Auth:   smtp.PlainAuth("", username, password, servers[0]),
	}

	return
}

// Attach add a attachment
func (m *Message) Attach(fname string) (err error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return
	}

	_, name := filepath.Split(fname)
	m.Attachments[name] = &Attachment{
		Name:    name,
		Content: data,
	}

	return
}

// Send do the sending
func (m *Message) Send() (err error) {
	return smtp.SendMail(m.Auth.Server, m.Auth.Auth, m.From, m.innerTo(), m.innerBody())
}

// innerTo returns mail receipt
func (m *Message) innerTo() (to []string) {
	to = m.To
	for _, v := range m.Cc {
		to = append(to, v)
	}

	for _, v := range m.Bcc {
		to = append(to, v)
	}

	return
}

// innerBody returns mail body
func (m *Message) innerBody() (body []byte) {
	now := time.Now()
	date := now.Format(time.RFC822)
	buf := bytes.NewBuffer(nil)
	boundary := "----=_NextPart_" + fmt.Sprintf("%x", md5.Sum([]byte(date)))[:16]

	buf.WriteString("Date: " + date + "\r\n")
	buf.WriteString("From: " + m.From + "\r\n")
	buf.WriteString("To: " + strings.Join(m.To, ",") + "\r\n")
	if len(m.Cc) > 0 {
		buf.WriteString("Cc: " + strings.Join(m.Cc, ",") + "\r\n")
	}
	buf.WriteString("Subject: " + m.Subject + "\r\n")
	buf.WriteString("X-Priority: 3\r\n")
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString("X-Mailer: github.com/likexian/gokit/xmailer\r\n")
	buf.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\r\n\r\n")

	buf.WriteString("This is a multi-part message in MIME format.\r\n")
	buf.WriteString("\r\n--" + boundary + "\r\n")
	buf.WriteString("Content-Type: " + m.ContentType + "; charset=utf-8\r\n\r\n")
	buf.WriteString(m.Body)
	buf.WriteString("\r\n")

	if len(m.Attachments) > 0 {
		for _, attachment := range m.Attachments {
			buf.WriteString("\r\n--" + boundary + "\r\n")
			buf.WriteString("Content-Type: application/octet-stream\r\n")
			buf.WriteString("Content-Transfer-Encoding: base64\r\n")
			buf.WriteString("Content-ID: <" + attachment.Name + ">\r\n")
			buf.WriteString("Content-Disposition: attachment; filename=\"" + attachment.Name + "\"\r\n\r\n")

			data := make([]byte, base64.StdEncoding.EncodedLen(len(attachment.Content)))
			base64.StdEncoding.Encode(data, attachment.Content)
			buf.Write(data)
			buf.WriteString("\r\n")
		}
	}

	buf.WriteString("\r\n--" + boundary + "--\r\n")
	body = buf.Bytes()

	return
}