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

// CoverageEligibilityRequest is documented here http://hl7.org/fhir/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequest struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes               `bson:"status" json:"status"`
	Priority          *CodeableConcept                           `bson:"priority,omitempty" json:"priority,omitempty"`
	Purpose           []EligibilityRequestPurpose                `bson:"purpose" json:"purpose"`
	Patient           Reference                                  `bson:"patient" json:"patient"`
	Created           string                                     `bson:"created" json:"created"`
	Enterer           *Reference                                 `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Provider          *Reference                                 `bson:"provider,omitempty" json:"provider,omitempty"`
	Insurer           Reference                                  `bson:"insurer" json:"insurer"`
	Facility          *Reference                                 `bson:"facility,omitempty" json:"facility,omitempty"`
	SupportingInfo    []CoverageEligibilityRequestSupportingInfo `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Insurance         []CoverageEligibilityRequestInsurance      `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Item              []CoverageEligibilityRequestItem           `bson:"item,omitempty" json:"item,omitempty"`
}
type CoverageEligibilityRequestSupportingInfo struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int         `bson:"sequence" json:"sequence"`
	Information       Reference   `bson:"information" json:"information"`
	AppliesToAll      *bool       `bson:"appliesToAll,omitempty" json:"appliesToAll,omitempty"`
}
type CoverageEligibilityRequestInsurance struct {
	Id                  *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Focal               *bool       `bson:"focal,omitempty" json:"focal,omitempty"`
	Coverage            Reference   `bson:"coverage" json:"coverage"`
	BusinessArrangement *string     `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
}
type CoverageEligibilityRequestItem struct {
	Id                     *string                                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension                               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SupportingInfoSequence []int                                     `bson:"supportingInfoSequence,omitempty" json:"supportingInfoSequence,omitempty"`
	Category               *CodeableConcept                          `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService       *CodeableConcept                          `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	Modifier               []CodeableConcept                         `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Provider               *Reference                                `bson:"provider,omitempty" json:"provider,omitempty"`
	Quantity               *Quantity                                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice              *Money                                    `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Facility               *Reference                                `bson:"facility,omitempty" json:"facility,omitempty"`
	Diagnosis              []CoverageEligibilityRequestItemDiagnosis `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Detail                 []Reference                               `bson:"detail,omitempty" json:"detail,omitempty"`
}
type CoverageEligibilityRequestItemDiagnosis struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type OtherCoverageEligibilityRequest CoverageEligibilityRequest

// MarshalJSON marshals the given CoverageEligibilityRequest as JSON into a byte slice
func (r CoverageEligibilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverageEligibilityRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverageEligibilityRequest: OtherCoverageEligibilityRequest(r),
		ResourceType:                    "CoverageEligibilityRequest",
	})
}

// UnmarshalCoverageEligibilityRequest unmarshals a CoverageEligibilityRequest.
func UnmarshalCoverageEligibilityRequest(b []byte) (CoverageEligibilityRequest, error) {
	var coverageEligibilityRequest CoverageEligibilityRequest
	if err := json.Unmarshal(b, &coverageEligibilityRequest); err != nil {
		return coverageEligibilityRequest, err
	}
	return coverageEligibilityRequest, nil
}
