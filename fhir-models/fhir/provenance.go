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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Provenance is documented here http://hl7.org/fhir/StructureDefinition/Provenance
type Provenance struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Target            []Reference        `bson:"target" json:"target"`
	Recorded          string             `bson:"recorded" json:"recorded"`
	Policy            []string           `bson:"policy,omitempty" json:"policy,omitempty"`
	Location          *Reference         `bson:"location,omitempty" json:"location,omitempty"`
	Reason            []CodeableConcept  `bson:"reason,omitempty" json:"reason,omitempty"`
	Activity          *CodeableConcept   `bson:"activity,omitempty" json:"activity,omitempty"`
	Agent             []ProvenanceAgent  `bson:"agent" json:"agent"`
	Entity            []ProvenanceEntity `bson:"entity,omitempty" json:"entity,omitempty"`
	Signature         []Signature        `bson:"signature,omitempty" json:"signature,omitempty"`
}
type ProvenanceAgent struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Role              []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Who               Reference         `bson:"who" json:"who"`
	OnBehalfOf        *Reference        `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type ProvenanceEntity struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              ProvenanceEntityRole `bson:"role" json:"role"`
	What              Reference            `bson:"what" json:"what"`
	Agent             []ProvenanceAgent    `bson:"agent,omitempty" json:"agent,omitempty"`
}
type OtherProvenance Provenance

// MarshalJSON marshals the given Provenance as JSON into a byte slice
func (r Provenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProvenance
		ResourceType string `json:"resourceType"`
	}{
		OtherProvenance: OtherProvenance(r),
		ResourceType:    "Provenance",
	})
}

// UnmarshalProvenance unmarshals a Provenance.
func UnmarshalProvenance(b []byte) (Provenance, error) {
	var provenance Provenance
	if err := json.Unmarshal(b, &provenance); err != nil {
		return provenance, err
	}
	return provenance, nil
}
