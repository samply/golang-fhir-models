// Copyright 2019 - 2021 The Samply Community
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

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ListMode is documented here http://hl7.org/fhir/ValueSet/list-mode
type ListMode int

const (
	ListModeWorking ListMode = iota
	ListModeSnapshot
	ListModeChanges
)

func (code ListMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ListMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "working":
		*code = ListModeWorking
	case "snapshot":
		*code = ListModeSnapshot
	case "changes":
		*code = ListModeChanges
	default:
		return fmt.Errorf("unknown ListMode code `%s`", s)
	}
	return nil
}
func (code ListMode) String() string {
	return code.Code()
}
func (code ListMode) Code() string {
	switch code {
	case ListModeWorking:
		return "working"
	case ListModeSnapshot:
		return "snapshot"
	case ListModeChanges:
		return "changes"
	}
	return "<unknown>"
}
func (code ListMode) Display() string {
	switch code {
	case ListModeWorking:
		return "Working List"
	case ListModeSnapshot:
		return "Snapshot List"
	case ListModeChanges:
		return "Change List"
	}
	return "<unknown>"
}
func (code ListMode) Definition() string {
	switch code {
	case ListModeWorking:
		return "This list is the master list, maintained in an ongoing fashion with regular updates as the real world list it is tracking changes."
	case ListModeSnapshot:
		return "This list was prepared as a snapshot. It should not be assumed to be current."
	case ListModeChanges:
		return "A point-in-time list that shows what changes have been made or recommended.  E.g. a discharge medication list showing what was added and removed during an encounter."
	}
	return "<unknown>"
}
