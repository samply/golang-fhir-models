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

// ActionRelationshipType is documented here http://hl7.org/fhir/ValueSet/action-relationship-type
type ActionRelationshipType int

const (
	ActionRelationshipTypeBeforeStart ActionRelationshipType = iota
	ActionRelationshipTypeBefore
	ActionRelationshipTypeBeforeEnd
	ActionRelationshipTypeConcurrentWithStart
	ActionRelationshipTypeConcurrent
	ActionRelationshipTypeConcurrentWithEnd
	ActionRelationshipTypeAfterStart
	ActionRelationshipTypeAfter
	ActionRelationshipTypeAfterEnd
)

func (code ActionRelationshipType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionRelationshipType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "before-start":
		*code = ActionRelationshipTypeBeforeStart
	case "before":
		*code = ActionRelationshipTypeBefore
	case "before-end":
		*code = ActionRelationshipTypeBeforeEnd
	case "concurrent-with-start":
		*code = ActionRelationshipTypeConcurrentWithStart
	case "concurrent":
		*code = ActionRelationshipTypeConcurrent
	case "concurrent-with-end":
		*code = ActionRelationshipTypeConcurrentWithEnd
	case "after-start":
		*code = ActionRelationshipTypeAfterStart
	case "after":
		*code = ActionRelationshipTypeAfter
	case "after-end":
		*code = ActionRelationshipTypeAfterEnd
	default:
		return fmt.Errorf("unknown ActionRelationshipType code `%s`", s)
	}
	return nil
}
func (code ActionRelationshipType) String() string {
	return code.Code()
}
func (code ActionRelationshipType) Code() string {
	switch code {
	case ActionRelationshipTypeBeforeStart:
		return "before-start"
	case ActionRelationshipTypeBefore:
		return "before"
	case ActionRelationshipTypeBeforeEnd:
		return "before-end"
	case ActionRelationshipTypeConcurrentWithStart:
		return "concurrent-with-start"
	case ActionRelationshipTypeConcurrent:
		return "concurrent"
	case ActionRelationshipTypeConcurrentWithEnd:
		return "concurrent-with-end"
	case ActionRelationshipTypeAfterStart:
		return "after-start"
	case ActionRelationshipTypeAfter:
		return "after"
	case ActionRelationshipTypeAfterEnd:
		return "after-end"
	}
	return "<unknown>"
}
func (code ActionRelationshipType) Display() string {
	switch code {
	case ActionRelationshipTypeBeforeStart:
		return "Before Start"
	case ActionRelationshipTypeBefore:
		return "Before"
	case ActionRelationshipTypeBeforeEnd:
		return "Before End"
	case ActionRelationshipTypeConcurrentWithStart:
		return "Concurrent With Start"
	case ActionRelationshipTypeConcurrent:
		return "Concurrent"
	case ActionRelationshipTypeConcurrentWithEnd:
		return "Concurrent With End"
	case ActionRelationshipTypeAfterStart:
		return "After Start"
	case ActionRelationshipTypeAfter:
		return "After"
	case ActionRelationshipTypeAfterEnd:
		return "After End"
	}
	return "<unknown>"
}
func (code ActionRelationshipType) Definition() string {
	switch code {
	case ActionRelationshipTypeBeforeStart:
		return "The action must be performed before the start of the related action."
	case ActionRelationshipTypeBefore:
		return "The action must be performed before the related action."
	case ActionRelationshipTypeBeforeEnd:
		return "The action must be performed before the end of the related action."
	case ActionRelationshipTypeConcurrentWithStart:
		return "The action must be performed concurrent with the start of the related action."
	case ActionRelationshipTypeConcurrent:
		return "The action must be performed concurrent with the related action."
	case ActionRelationshipTypeConcurrentWithEnd:
		return "The action must be performed concurrent with the end of the related action."
	case ActionRelationshipTypeAfterStart:
		return "The action must be performed after the start of the related action."
	case ActionRelationshipTypeAfter:
		return "The action must be performed after the related action."
	case ActionRelationshipTypeAfterEnd:
		return "The action must be performed after the end of the related action."
	}
	return "<unknown>"
}
