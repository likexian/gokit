/*
 * Copyright 2012-2023 Li Kexian
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

package xmail

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"io/ioutil"
	"net/smtp"
	"path/filepath"
	"strings"
	"time"

	"github.com/likexian/gokit/xhash"
)

// attachment storing mail attachment
type attachment struct {
	name    string
	content []byte
}

// auth storing mail auth
type auth struct {
	tls    bool
	server string
	host   string
	auth   smtp.Auth
}

// Message storing mail message
type Message struct {
	from        string
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	contentType string
	attachments map[string]*attachment
	auth        *auth
}

// Version returns package version
func Version() string {
	return "0.8.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// New returns a new xmail
func New(server, username, password string, tls bool) (m *Message) {
	m = &Message{
		from:        username,
		to:          []string{username},
		cc:          []string{},
		bcc:         []string{},
		subject:     "Mailer Test",
		body:        "Hello World. This is xmail via github.com/likexian/gokit/xmail.",
		contentType: "text/plain",
		attachments: map[string]*attachment{},
	}

	servers := strings.Split(server, ":")
	m.auth = &auth{
		tls:    tls,
		server: server,
		host:   servers[0],
		auth:   smtp.PlainAuth("", username, password, servers[0]),
	}

	return
}

// From set mail from
func (m *Message) From(s string) error {
	m.from = s
	return nil
}

// To set mail to
func (m *Message) To(s ...string) error {
	m.to = s
	return nil
}

// Cc set mail cc
func (m *Message) Cc(s ...string) error {
	m.cc = s
	return nil
}

// BCc set mail bcc
func (m *Message) BCc(s ...string) error {
	m.bcc = s
	return nil
}

// ContentType set mail content type
func (m *Message) ContentType(t string) error {
	m.contentType = t
	return nil
}

// Content set mail content
func (m *Message) Content(subject, body string) error {
	m.subject = subject
	m.body = body
	return nil
}

// Attach add a attachment
func (m *Message) Attach(fname string) (err error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return
	}

	_, name := filepath.Split(fname)
	m.attachments[name] = &attachment{
		name:    name,
		content: data,
	}

	return
}

// Send do the sending
func (m *Message) Send() (err error) {
	if !m.auth.tls {
		return smtp.SendMail(m.auth.server, m.auth.auth, m.from, m.innerTo(), m.innerBody())
	}

	return m.tlsSendMail()
}

// tlsSendMail send using tls
func (m *Message) tlsSendMail() (err error) {
	conn, err := tls.Dial("tcp", m.auth.server, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, m.auth.host)
	if err != nil {
		return
	}
	defer client.Close()

	if ok, _ := client.Extension("AUTH"); ok {
		err = client.Auth(m.auth.auth)
		if err != nil {
			return
		}
	}

	err = client.Mail(m.from)
	if err != nil {
		return
	}

	for _, v := range m.innerTo() {
		err = client.Rcpt(v)
		if err != nil {
			return
		}
	}

	fd, err := client.Data()
	if err != nil {
		return
	}

	_, err = fd.Write(m.innerBody())
	if err != nil {
		return
	}

	err = fd.Close()
	if err != nil {
		return
	}

	err = client.Quit()

	return
}

// innerTo returns mail receipt
func (m *Message) innerTo() (to []string) {
	to = m.to
	for _, v := range m.cc {
		if v == "" {
			continue
		}
		to = append(to, v)
	}

	for _, v := range m.bcc {
		if v == "" {
			continue
		}
		to = append(to, v)
	}

	return
}

// innerBody returns mail body
func (m *Message) innerBody() (body []byte) {
	now := time.Now()
	date := now.Format(time.RFC822)
	buf := bytes.NewBuffer(nil)
	boundary := "----=_NextPart_" + xhash.Md5(now.UnixNano()).Hex()[:16]

	buf.WriteString("Date: " + date + "\r\n")
	buf.WriteString("From: " + m.from + "\r\n")
	buf.WriteString("To: " + strings.Join(m.to, ",") + "\r\n")
	if len(m.cc) > 0 {
		buf.WriteString("Cc: " + strings.Join(m.cc, ",") + "\r\n")
	}
	buf.WriteString("Subject: " + m.subject + "\r\n")
	buf.WriteString("X-Priority: 3\r\n")
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString("X-Mailer: github.com/likexian/gokit/xmail\r\n")
	buf.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\r\n\r\n")

	buf.WriteString("This is a multi-part message in MIME format.\r\n")
	buf.WriteString("\r\n--" + boundary + "\r\n")
	buf.WriteString("Content-Type: " + m.contentType + "; charset=utf-8\r\n\r\n")
	buf.WriteString(m.body)
	buf.WriteString("\r\n")

	if len(m.attachments) > 0 {
		for _, attachment := range m.attachments {
			buf.WriteString("\r\n--" + boundary + "\r\n")
			buf.WriteString("Content-Type: application/octet-stream\r\n")
			buf.WriteString("Content-Transfer-Encoding: base64\r\n")
			buf.WriteString("Content-ID: <" + attachment.name + ">\r\n")
			buf.WriteString("Content-Disposition: attachment; filename=\"" + attachment.name + "\"\r\n\r\n")

			data := make([]byte, base64.StdEncoding.EncodedLen(len(attachment.content)))
			base64.StdEncoding.Encode(data, attachment.content)
			buf.Write(data)
			buf.WriteString("\r\n")
		}
	}

	buf.WriteString("\r\n--" + boundary + "--\r\n")
	body = buf.Bytes()

	return
}
