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

// AdverseEventActuality is documented here http://hl7.org/fhir/ValueSet/adverse-event-actuality
type AdverseEventActuality int

const (
	AdverseEventActualityActual AdverseEventActuality = iota
	AdverseEventActualityPotential
)

func (code AdverseEventActuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AdverseEventActuality) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "actual":
		*code = AdverseEventActualityActual
	case "potential":
		*code = AdverseEventActualityPotential
	default:
		return fmt.Errorf("unknown AdverseEventActuality code `%s`", s)
	}
	return nil
}
func (code AdverseEventActuality) String() string {
	return code.Code()
}
func (code AdverseEventActuality) Code() string {
	switch code {
	case AdverseEventActualityActual:
		return "actual"
	case AdverseEventActualityPotential:
		return "potential"
	}
	return "<unknown>"
}
func (code AdverseEventActuality) Display() string {
	switch code {
	case AdverseEventActualityActual:
		return "Adverse Event"
	case AdverseEventActualityPotential:
		return "Potential Adverse Event"
	}
	return "<unknown>"
}
func (code AdverseEventActuality) Definition() string {
	switch code {
	case AdverseEventActualityActual:
		return "The adverse event actually happened regardless of whether anyone was affected or harmed."
	case AdverseEventActualityPotential:
		return "A potential adverse event."
	}
	return "<unknown>"
}
