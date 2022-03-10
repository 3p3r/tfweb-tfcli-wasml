// borrowed from https://github.com/coreos/go-systemd/blob/da560eaf4d2e74df33e0138791f3c76f8ef60bf3/journal/journal.go
// and https://github.com/coreos/go-systemd/blob/da560eaf4d2e74df33e0138791f3c76f8ef60bf3/journal/journal_unix.go
//
// replace Send with Println

// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package journal provides write bindings to the local systemd journal.
// It is implemented in pure Go and connects to the journal directly over its
// unix socket.
//
// To read from the journal, see the "sdjournal" package, which wraps the
// sd-journal a C API.
//
// http://www.freedesktop.org/software/systemd/man/systemd-journald.service.html
package journal

import (
	"fmt"
)

// Priority of a journal message
type Priority int

const (
	PriEmerg Priority = iota
	PriAlert
	PriCrit
	PriErr
	PriWarning
	PriNotice
	PriInfo
	PriDebug
)

// Print prints a message to the local systemd journal using Send().
func Print(priority Priority, format string, a ...interface{}) error {
	fmt.Println(fmt.Sprintf(format, a...))
	return nil
}

func Enabled() bool {
	return true
}

func Send(message string, priority Priority, vars map[string]string) error {
	Print(priority, "%v", vars)
	return nil
}
