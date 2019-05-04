/*
 * Go module for sending email
 * https://www.likexian.com/
 *
 * Copyright 2015-2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package mailer


import (
    "fmt"
    "testing"
)


func TestMailer(t *testing.T) {
    // Set the smtp info
    m := New("smtp.likexian.com:25", "i@likexian.com", "8Bd0a7681333214", true)

    // Set email from
    m.From = "i@likexian.com"

    // Set send to
    m.To = []string{"i@likexian.com"}

    // Set mail subject
    m.Subject = "Mailer Test"

    // Set mail body
    m.Body = "Hello World. This is mailer via github.com/likexian/mailer-go.<br /><img src=\"cid:mailer_test.jpg\" />"

    // Add attachment
    err := m.Attach("mailer_test.jpg")
    if err != nil {
        t.Error(err)
    }

    if _, ok := m.Attachments["mailer_test.jpg"]; !ok {
        t.Error("Fail to add attachment")
    }

    // The smtp info is fake, sending will never success.
    err = m.Send()
    if (err != nil) {
        // t.Error(err)
        fmt.Println("There is something WRONG:")
        fmt.Println(err)
        fmt.Println("")
    }
}
