/*
 * Go module for sending email
 * https://www.likexian.com/
 *
 * Copyright 2015, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */


package mailer


import (
    "fmt"
    "testing"
    "github.com/bmizerany/assert"
)


func TestMailer(t *testing.T) {
    // Set the smtp info
    m := New("smtp.likexian.com:25", "i@likexian.com", "8Bd0a7681333214", true)
    assert.Equal(t, "i@likexian.com", m.From)

    // Set email from
    m.From = "i@likexian.com"
    assert.Equal(t, "i@likexian.com", m.From)

    // Set send to
    m.To = []string{"i@likexian.com"}
    assert.Equal(t, 1, len(m.To))

    // Set mail subject
    m.Subject = "Mailer Test"
    assert.NotEqual(t, "", m.Body)

    // Set mail body
    m.Body = "Hello World. This is mailer via github.com/likexian/mailer-go.<br /><img src=\"cid:mailer_test.jpg\" />"
    assert.NotEqual(t, "", m.Body)

    // Add attachment
    err := m.Attach("mailer_test.jpg")
    assert.Equal(t, nil, err)
    if _, ok := m.Attachments["mailer_test.jpg"]; !ok {
        assert.Equal(t, true, ok)
    }

    // The smtp info is fake, sending will never success.
    // If you changed smtp info, change assert to Equal nil
    err = m.Send()
    assert.NotEqual(t, nil, err)
    if (err != nil) {
        fmt.Println("There is something WRONG:")
        fmt.Println(err)
        fmt.Println("")
    }
}
