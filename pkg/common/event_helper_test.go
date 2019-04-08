/*
Copyright 2019 The Unity Scheduler Authors

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

package common

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestAllocateTaskEventArgs(t *testing.T) {
	alloc := NewAllocateTaskEvent("UID-0001", "node-0001")
	args := alloc.getArgs()

	assert.Equal(t, len(args), 2)
	assert.Equal(t, fmt.Sprint(args[0]), "UID-0001")
	assert.Equal(t, fmt.Sprint(args[1]), "node-0001")
}

func TestGetAllocateTaskEventArgs(t *testing.T) {
	alloc := NewAllocateTaskEvent("UID-0001", "node-0001")
	args := alloc.getArgs()
	assert.Equal(t, len(args), 2)
	assert.Equal(t, fmt.Sprint(args[0]), "UID-0001")
	assert.Equal(t, fmt.Sprint(args[1]), "node-0001")

	out := make([]string, 2)
	GetEventArgsAsStrings(out, args)
	assert.Equal(t, out[0], "UID-0001")
	assert.Equal(t, out[1], "node-0001")

	out = make([]string, 0)
	err := GetEventArgsAsStrings(out, args)
	assert.Assert(t, err != nil)

	out = make([]string, 5)
	err = GetEventArgsAsStrings(out, args)
	assert.Assert(t, err != nil)

	err = GetEventArgsAsStrings(nil, args)
	assert.Assert(t, err != nil)
}
