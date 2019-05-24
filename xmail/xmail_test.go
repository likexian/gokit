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

package xmail

import (
	"github.com/likexian/gokit/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestMailer(t *testing.T) {
	// Set the smtp info
	m := New("smtp.likexian.com:25", "i@likexian.com", "8Bd0a7681333214")

	// Set email from
	m.From("i@likexian.com")

	// Set send to
	m.To("i@likexian.com")

	// Set send cc
	m.Cc("cc@likexian.com")

	// Set send bcc
	m.BCc("bcc@likexian.com")

	// set mail content type
	m.ContentType("text/html")

	// Set mail subject
	m.Content("Mailer Test", "xmail via github.com/likexian/gokit/xmail.<br /><img src=\"cid:xmail_test.jpg\" />")

	// Add attachment
	err := m.Attach("xmail_test.jpg")
	assert.Nil(t, err)

	// Add attachment
	err = m.Attach("not-exists.jpg")
	assert.NotNil(t, err)

	err = m.Send()
	// The smtp auth info is fake, sending will never success.
	// Change below line to
	// assert.Nil(t, err)
	// If specify the valid smtp auth info
	assert.NotNil(t, err)
}
