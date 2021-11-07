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

// PlanDefinition is documented here http://hl7.org/fhir/StructureDefinition/PlanDefinition
type PlanDefinition struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string                `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                `bson:"title,omitempty" json:"title,omitempty"`
	Subtitle          *string                `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	Type              *CodeableConcept       `bson:"type,omitempty" json:"type,omitempty"`
	Status            PublicationStatus      `bson:"status" json:"status"`
	Experimental      *bool                  `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail        `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext         `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept      `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage             *string                `bson:"usage,omitempty" json:"usage,omitempty"`
	Copyright         *string                `bson:"copyright,omitempty" json:"copyright,omitempty"`
	ApprovalDate      *string                `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate    *string                `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Topic             []CodeableConcept      `bson:"topic,omitempty" json:"topic,omitempty"`
	Author            []ContactDetail        `bson:"author,omitempty" json:"author,omitempty"`
	Editor            []ContactDetail        `bson:"editor,omitempty" json:"editor,omitempty"`
	Reviewer          []ContactDetail        `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	Endorser          []ContactDetail        `bson:"endorser,omitempty" json:"endorser,omitempty"`
	RelatedArtifact   []RelatedArtifact      `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Library           []string               `bson:"library,omitempty" json:"library,omitempty"`
	Goal              []PlanDefinitionGoal   `bson:"goal,omitempty" json:"goal,omitempty"`
	Action            []PlanDefinitionAction `bson:"action,omitempty" json:"action,omitempty"`
}
type PlanDefinitionGoal struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          *CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Description       CodeableConcept            `bson:"description" json:"description"`
	Priority          *CodeableConcept           `bson:"priority,omitempty" json:"priority,omitempty"`
	Start             *CodeableConcept           `bson:"start,omitempty" json:"start,omitempty"`
	Addresses         []CodeableConcept          `bson:"addresses,omitempty" json:"addresses,omitempty"`
	Documentation     []RelatedArtifact          `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Target            []PlanDefinitionGoalTarget `bson:"target,omitempty" json:"target,omitempty"`
}
type PlanDefinitionGoalTarget struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Measure           *CodeableConcept `bson:"measure,omitempty" json:"measure,omitempty"`
	Due               *Duration        `bson:"due,omitempty" json:"due,omitempty"`
}
type PlanDefinitionAction struct {
	Id                  *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Prefix              *string                             `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Title               *string                             `bson:"title,omitempty" json:"title,omitempty"`
	Description         *string                             `bson:"description,omitempty" json:"description,omitempty"`
	TextEquivalent      *string                             `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	Priority            *RequestPriority                    `bson:"priority,omitempty" json:"priority,omitempty"`
	Code                []CodeableConcept                   `bson:"code,omitempty" json:"code,omitempty"`
	Reason              []CodeableConcept                   `bson:"reason,omitempty" json:"reason,omitempty"`
	Documentation       []RelatedArtifact                   `bson:"documentation,omitempty" json:"documentation,omitempty"`
	GoalId              []string                            `bson:"goalId,omitempty" json:"goalId,omitempty"`
	Trigger             []TriggerDefinition                 `bson:"trigger,omitempty" json:"trigger,omitempty"`
	Condition           []PlanDefinitionActionCondition     `bson:"condition,omitempty" json:"condition,omitempty"`
	Input               []DataRequirement                   `bson:"input,omitempty" json:"input,omitempty"`
	Output              []DataRequirement                   `bson:"output,omitempty" json:"output,omitempty"`
	RelatedAction       []PlanDefinitionActionRelatedAction `bson:"relatedAction,omitempty" json:"relatedAction,omitempty"`
	Participant         []PlanDefinitionActionParticipant   `bson:"participant,omitempty" json:"participant,omitempty"`
	Type                *CodeableConcept                    `bson:"type,omitempty" json:"type,omitempty"`
	GroupingBehavior    *ActionGroupingBehavior             `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty"`
	SelectionBehavior   *ActionSelectionBehavior            `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty"`
	RequiredBehavior    *ActionRequiredBehavior             `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty"`
	PrecheckBehavior    *ActionPrecheckBehavior             `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty"`
	CardinalityBehavior *ActionCardinalityBehavior          `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty"`
	Transform           *string                             `bson:"transform,omitempty" json:"transform,omitempty"`
	DynamicValue        []PlanDefinitionActionDynamicValue  `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty"`
	Action              []PlanDefinitionAction              `bson:"action,omitempty" json:"action,omitempty"`
}
type PlanDefinitionActionCondition struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Kind              ActionConditionKind `bson:"kind" json:"kind"`
	Expression        *Expression         `bson:"expression,omitempty" json:"expression,omitempty"`
}
type PlanDefinitionActionRelatedAction struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ActionId          string                 `bson:"actionId" json:"actionId"`
	Relationship      ActionRelationshipType `bson:"relationship" json:"relationship"`
}
type PlanDefinitionActionParticipant struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              ActionParticipantType `bson:"type" json:"type"`
	Role              *CodeableConcept      `bson:"role,omitempty" json:"role,omitempty"`
}
type PlanDefinitionActionDynamicValue struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Path              *string     `bson:"path,omitempty" json:"path,omitempty"`
	Expression        *Expression `bson:"expression,omitempty" json:"expression,omitempty"`
}
type OtherPlanDefinition PlanDefinition

// MarshalJSON marshals the given PlanDefinition as JSON into a byte slice
func (r PlanDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPlanDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherPlanDefinition: OtherPlanDefinition(r),
		ResourceType:        "PlanDefinition",
	})
}

// UnmarshalPlanDefinition unmarshals a PlanDefinition.
func UnmarshalPlanDefinition(b []byte) (PlanDefinition, error) {
	var planDefinition PlanDefinition
	if err := json.Unmarshal(b, &planDefinition); err != nil {
		return planDefinition, err
	}
	return planDefinition, nil
}
