// Copyright 2019 The Samply Development Community
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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Extension is documented here http://hl7.org/fhir/StructureDefinition/Extension
type Extension struct {
	Id          *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	Url         string       `bson:"url,omitempty" json:"url,omitempty"`
	ValueString interface{}  `bson:"valueString,omitempty" json:"valueString,omitempty"` // can be string or encrypted field
	ValueCode   *ValueCode   `bson:"_valueCode,omitempty" json:"_valueCode,omitempty"`
	ValuePeriod *ValuePeriod `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
}

type ValueCode struct {
	Value  string `bson:"value,omitempty" json:"value,omitempty"`
	Code   string `bson:"code,omitempty" json:"code,omitempty"`
	System string `bson:"system,omitempty" json:"system,omitempty"`
}

type ValuePeriod struct {
	Start string `bson:"start,omitempty" json:"start,omitempty"`
	End   string `bson:"end,omitempty" json:"end,omitempty"`
}
