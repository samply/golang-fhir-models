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

// AllergyIntolerance is documented here http://hl7.org/fhir/StructureDefinition/AllergyIntolerance
type AllergyIntolerance struct {
	Id                 *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ClinicalStatus     *CodeableConcept               `bson:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	VerificationStatus *CodeableConcept               `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	Type               *AllergyIntoleranceType        `bson:"type,omitempty" json:"type,omitempty"`
	Category           []AllergyIntoleranceCategory   `bson:"category,omitempty" json:"category,omitempty"`
	Criticality        *AllergyIntoleranceCriticality `bson:"criticality,omitempty" json:"criticality,omitempty"`
	Code               *CodeableConcept               `bson:"code,omitempty" json:"code,omitempty"`
	Patient            Reference                      `bson:"patient" json:"patient"`
	Encounter          *Reference                     `bson:"encounter,omitempty" json:"encounter,omitempty"`
	RecordedDate       *string                        `bson:"recordedDate,omitempty" json:"recordedDate,omitempty"`
	Recorder           *Reference                     `bson:"recorder,omitempty" json:"recorder,omitempty"`
	Asserter           *Reference                     `bson:"asserter,omitempty" json:"asserter,omitempty"`
	LastOccurrence     *string                        `bson:"lastOccurrence,omitempty" json:"lastOccurrence,omitempty"`
	Note               []Annotation                   `bson:"note,omitempty" json:"note,omitempty"`
	Reaction           []AllergyIntoleranceReaction   `bson:"reaction,omitempty" json:"reaction,omitempty"`
}
type AllergyIntoleranceReaction struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Substance         *CodeableConcept            `bson:"substance,omitempty" json:"substance,omitempty"`
	Manifestation     []CodeableConcept           `bson:"manifestation" json:"manifestation"`
	Description       *string                     `bson:"description,omitempty" json:"description,omitempty"`
	Onset             *string                     `bson:"onset,omitempty" json:"onset,omitempty"`
	Severity          *AllergyIntoleranceSeverity `bson:"severity,omitempty" json:"severity,omitempty"`
	ExposureRoute     *CodeableConcept            `bson:"exposureRoute,omitempty" json:"exposureRoute,omitempty"`
	Note              []Annotation                `bson:"note,omitempty" json:"note,omitempty"`
}
type OtherAllergyIntolerance AllergyIntolerance

// MarshalJSON marshals the given AllergyIntolerance as JSON into a byte slice
func (r AllergyIntolerance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAllergyIntolerance
		ResourceType string `json:"resourceType"`
	}{
		OtherAllergyIntolerance: OtherAllergyIntolerance(r),
		ResourceType:            "AllergyIntolerance",
	})
}

// UnmarshalAllergyIntolerance unmarshals a AllergyIntolerance.
func UnmarshalAllergyIntolerance(b []byte) (AllergyIntolerance, error) {
	var allergyIntolerance AllergyIntolerance
	if err := json.Unmarshal(b, &allergyIntolerance); err != nil {
		return allergyIntolerance, err
	}
	return allergyIntolerance, nil
}
