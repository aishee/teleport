/*
Copyright 2020 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pam

import (
    "bytes"
    "encoding/json"
    "fmt"
    "testing"

    "github.com/gravitational/teleport/lib/utils"

    "gopkg.in/check.v1"
)

type Suite struct{}

var _ = fmt.Printf
var _ = check.Suite(&Suite{})

func TestPAM(t *testing.T) { check.TestingT(t) }

func (s *Suite) SetUpSuite(c *check.C) {
    utils.InitLoggerForTests()
}
func (s *Suite) TearDownSuite(c *check.C) {}
func (s *Suite) SetUpTest(c *check.C)     {}
func (s *Suite) TearDownTest(c *check.C)  {}

func (s *Suite) TestEcho(c *check.C) {
    var buf bytes.Buffer

    if !BuildHasPAM() || !SystemHasPAM() {
        c.Skip("Skipping test: PAM support not enabled.")
    }

    pamContext, err := Open(&Config{
        Enabled:     true,
        ServiceName: "teleport-session-echo-ruser",
        LoginContext: &LoginContextV1{
            Version:  1,
            Username: "foo",
            Login:    "bar",
            Roles:    []string{"baz", "qux"},
        },
        Stdin:  &discardReader{},
        Stdout: &buf,
        Stderr: &buf,
    })
    defer pamContext.Close()
    c.Assert(err, check.IsNil)

    var context LoginContextV1
    err = json.Unmarshal(buf.Bytes(), &context)
    c.Assert(err, check.IsNil)

    c.Assert(context.Username, check.Equals, "foo")
    c.Assert(context.Login, check.Equals, "bar")
    c.Assert(context.Roles, check.DeepEquals, []string{"baz", "qux"})
}

func (s *Suite) TestEnvironment(c *check.C) {
    var buf bytes.Buffer

    if !BuildHasPAM() || !SystemHasPAM() {
        c.Skip("Skipping test: PAM support not enabled.")
    }

    pamContext, err := Open(&Config{
        Enabled:      true,
        ServiceName:  "teleport-session-environment",
        LoginContext: &LoginContextV1{},
        Stdin:        &discardReader{},
        Stdout:       &buf,
        Stderr:       &buf,
    })
    defer pamContext.Close()
    c.Assert(err, check.IsNil)

    c.Assert(pamContext.Environment(), check.HasLen, 1)
    c.Assert(pamContext.Environment()[0], check.Equals, "foo=bar")
}

type discardReader struct {
}

func (d *discardReader) Read(p []byte) (int, error) {
    return len(p), nil
}
