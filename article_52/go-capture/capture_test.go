/*
Copyright © 2020 Pavel Tisnovsky

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

package capture_test

import (
	"fmt"
	"github.com/tisnik/go-capture"
	"os"
	"testing"
)

// TestNoOutput checks if empty standard output is captured properly
func TestNoOutput(t *testing.T) {
	captured, err := capture.StandardOutput(func() {
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "" {
		t.Fatal("Standard should be empty")
	}
}

// TestEmptyOutput checks if empty standard output is captured properly
func TestEmptyOutput(t *testing.T) {
	captured, err := capture.StandardOutput(func() {
		fmt.Print("")
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "" {
		t.Fatal("Standard should be empty")
	}
}

// TestOutputWithoutNewlines checks if standard output created by fmt.Print is captured properly
func TestOutputWithoutNewlines(t *testing.T) {
	captured, err := capture.StandardOutput(func() {
		fmt.Print("Hello!")
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "Hello!" {
		t.Fatal("Incorrect output has been captured:", captured)
	}
}

// TestOutputWithNewlines checks if standard output created by fmt.Println is captured properly
func TestOutputWithNewlines(t *testing.T) {
	captured, err := capture.StandardOutput(func() {
		fmt.Println("Hello!")
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "Hello!\n" {
		t.Fatal("Incorrect output has been captured:", captured)
	}
}

// TestOutputToStdErr checks whether output to stderr is captured or not
func TestOutputToStdErr(t *testing.T) {
	captured, err := capture.StandardOutput(func() {
		fmt.Fprint(os.Stderr, "Hello!")
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "" {
		t.Fatal("Incorrect output has been captured:", captured)
	}
}
