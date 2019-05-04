/*
 * Go module for sending email
 * https://www.likexian.com/
 *
 * Copyright 2015-2018, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package mailer


import (
    "fmt"
    "time"
    "strings"
    "bytes"
    "encoding/base64"
    "crypto/md5"
    "path/filepath"
    "io/ioutil"
    "net/smtp"
)


type Attachment struct {
    Name        string
    Content     []byte
}

type Auth struct {
    Server      string
    Auth        smtp.Auth
}

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


func Version() string {
    return "0.1.0"
}


func Author() string {
    return "[Li Kexian](https://www.likexian.com/)"
}


func License() string {
    return "Apache License, Version 2.0"
}


func New(server, username, password string, ishtml bool) (m *Message) {
    m = &Message{
        From: username,
        To: []string{username},
        Subject: "Mailer Test",
        Body: "Hello World. This is mailer via github.com/likexian/mailer-go.",
        Attachments: make(map[string]*Attachment),
    }

    if ishtml {
        m.ContentType = "text/html"
    } else {
        m.ContentType = "text/plain"
    }

    servers := strings.Split(server, ":")
    m.Auth = &Auth {
        Server: server,
        Auth: smtp.PlainAuth("", username, password, servers[0]),
    }

    return
}


func (m *Message) Attach(fname string) (err error) {
    data, err := ioutil.ReadFile(fname)
    if err != nil {
        return
    }

    _, name := filepath.Split(fname)
    m.Attachments[name] = &Attachment {
        Name: name,
        Content: data,
    }

    return
}


func (m *Message) Send() (err error) {
    return smtp.SendMail(m.Auth.Server, m.Auth.Auth, m.From, m.innerTo(), m.innerBody())
}


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
    buf.WriteString("X-Mailer: github.com/likexian/mailer-go\r\n")
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
