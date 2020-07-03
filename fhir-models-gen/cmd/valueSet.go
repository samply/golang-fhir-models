package cmd

import (
	"errors"
	"fmt"
	"github.com/dave/jennifer/jen"
	"fhir-models-gen/fhir"
	"strings"
)

func generateValueSet(resources ResourceMap, valueSet fhir.ValueSet) (*jen.File, error) {
	if valueSet.Name == nil {
		return nil, errors.New("ValueSet without name")
	}

	if valueSet.Compose == nil || len(valueSet.Compose.Include) == 0 {
		return nil, fmt.Errorf("the ValueSet `%s` doens't include any CodeSystems", *valueSet.Name)
	}

	if len(valueSet.Compose.Include) > 1 {
		return nil, fmt.Errorf("the ValueSet `%s` includes more than one CodeSystem", *valueSet.Name)
	}

	url := canonical(valueSet.Compose.Include[0])
	if url == "" {
		return nil, fmt.Errorf("the ValueSet `%s` doens't include any CodeSystems", *valueSet.Name)
	}

	bytes := resources["CodeSystem"][url]
	if bytes == nil {
		return nil, fmt.Errorf("missing CodeSystem with canonical URL `%s` in ValueSet `%s`", url, *valueSet.Name)
	}

	codeSystem, err := fhir.UnmarshalCodeSystem(bytes)
	if err != nil {
		return nil, err
	}

	if len(codeSystem.Concept) == 0 {
		return nil, fmt.Errorf("the CodeSystem with canonical URL `%s` has no codes", url)
	}

	fmt.Printf("Generate Go sources for ValueSet: %s\n", *valueSet.Name)
	file := jen.NewFile("fhir")
	appendLicenseComment(file)
	appendGeneratorComment(file)

	// type
	file.Commentf("%s is documented here %s", *valueSet.Name, *valueSet.Url)
	file.Type().Id(*valueSet.Name).Int()
	file.Const().DefsFunc(constsRoot(*valueSet.Name, codeSystem.Concept))

	// MarshalJSON function
	file.Func().
		Params(jen.Id("code").Id(*valueSet.Name)).
		Id("MarshalJSON").
		Params().
		Params(jen.Op("[]").Byte(), jen.Error()).
		Block(
			jen.Return(jen.Qual("encoding/json", "Marshal").Call(jen.Id("code").Op(".").Id("Code").Call())),
		)

	// UnmarshalJSON function
	file.Func().
		Params(jen.Id("code").Op("*").Id(*valueSet.Name)).
		Id("UnmarshalJSON").
		Params(jen.Id("json").Op("[]").Byte()).
		Error().
		Block(
			jen.Id("s").Op(":=").Qual("strings", "Trim").Call(jen.String().Call(jen.Id("json")), jen.Lit(`"`)),
			jen.Switch(jen.Id("s")).BlockFunc(unmarshalRoot(*valueSet.Name, codeSystem.Concept)),
			jen.Return(jen.Nil()),
		)

	// String function
	file.Func().
		Params(jen.Id("code").Id(*valueSet.Name)).
		Id("String").
		Params().
		String().
		Block(
			jen.Return(jen.Id("code").Op(".").Id("Code").Call()),
		)

	// Code function
	file.Func().
		Params(jen.Id("code").Id(*valueSet.Name)).
		Id("Code").
		Params().
		String().
		Block(
			jen.Switch(jen.Id("code")).BlockFunc(codes(*valueSet.Name, codeSystem.Concept)),
			jen.Return(jen.Lit("<unknown>")),
		)

	// Display function
	file.Func().
		Params(jen.Id("code").Id(*valueSet.Name)).
		Id("Display").
		Params().
		String().
		Block(
			jen.Switch(jen.Id("code")).BlockFunc(displays(*valueSet.Name, codeSystem.Concept)),
			jen.Return(jen.Lit("<unknown>")),
		)

	// Definition function
	file.Func().
		Params(jen.Id("code").Id(*valueSet.Name)).
		Id("Definition").
		Params().
		String().
		Block(
			jen.Switch(jen.Id("code")).BlockFunc(definitions(*valueSet.Name, codeSystem.Concept)),
			jen.Return(jen.Lit("<unknown>")),
		)

	return file, nil
}

func canonical(include fhir.ValueSetComposeInclude) string {
	if system := include.System; system != nil {
		if version := include.Version; version != nil {
			return *system + "|" + *version
		}
		return *system
	}
	return ""
}

func constsRoot(valueSetName string, concepts []fhir.CodeSystemConcept) func(*jen.Group) {
	return func(group *jen.Group) {
		group.Id(codeIdentifier(valueSetName, concepts[0].Code)).Id(valueSetName).Op("=").Iota()
		if len(concepts[0].Concept) > 0 {
			consts(valueSetName, concepts[0].Concept)(group)
		}
		consts(valueSetName, concepts[1:])(group)
	}
}

func consts(valueSetName string, concepts []fhir.CodeSystemConcept) func(*jen.Group) {
	return func(group *jen.Group) {
		for _, concept := range concepts {
			group.Id(codeIdentifier(valueSetName, concept.Code))
			if len(concept.Concept) > 0 {
				consts(valueSetName, concept.Concept)(group)
			}
		}
	}
}

func codeIdentifier(valueSetName, s string) string {
	switch s {
	case "=":
		return valueSetName + "Equals"
	case "!=":
		return valueSetName + "NotEquals"
	case ">":
		return valueSetName + "GreaterThan"
	case "<":
		return valueSetName + "LessThan"
	case ">=":
		return valueSetName + "GreaterOrEquals"
	case "<=":
		return valueSetName + "LessOrEquals"
	default:
		return valueSetName + strings.ReplaceAll(strings.ReplaceAll(strings.Title(s), "-", ""), ".", "_")
	}
}

func unmarshalRoot(valueSetName string, concepts []fhir.CodeSystemConcept) func(group *jen.Group) {
	return func(group *jen.Group) {
		unmarshal(valueSetName, concepts)(group)
		group.Default().Block(
			jen.Return(jen.Qual("fmt", "Errorf").Call(jen.Lit("unknown "+valueSetName+" code `%s`"), jen.Id("s"))),
		)
	}
}

func unmarshal(valueSetName string, concepts []fhir.CodeSystemConcept) func(group *jen.Group) {
	return func(group *jen.Group) {
		for _, concept := range concepts {
			group.Case(jen.Lit(concept.Code)).Block(jen.Op("*").Id("code").Op("=").Id(codeIdentifier(valueSetName, concept.Code)))
			if len(concept.Concept) > 0 {
				unmarshal(valueSetName, concept.Concept)(group)
			}
		}
	}
}

func codes(valueSetName string, concepts []fhir.CodeSystemConcept) func(group *jen.Group) {
	return func(group *jen.Group) {
		for _, concept := range concepts {
			group.Case(jen.Id(codeIdentifier(valueSetName, concept.Code))).Block(jen.Return(jen.Lit(concept.Code)))
			if len(concept.Concept) > 0 {
				displays(valueSetName, concept.Concept)(group)
			}
		}
	}
}

func displays(valueSetName string, concepts []fhir.CodeSystemConcept) func(group *jen.Group) {
	return func(group *jen.Group) {
		for _, concept := range concepts {
			if concept.Display != nil {
				group.Case(jen.Id(codeIdentifier(valueSetName, concept.Code))).Block(jen.Return(jen.Lit(*concept.Display)))
			}
			if len(concept.Concept) > 0 {
				displays(valueSetName, concept.Concept)(group)
			}
		}
	}
}

func definitions(valueSetName string, concepts []fhir.CodeSystemConcept) func(group *jen.Group) {
	return func(group *jen.Group) {
		for _, concept := range concepts {
			if concept.Definition != nil {
				group.Case(jen.Id(codeIdentifier(valueSetName, concept.Code))).Block(jen.Return(jen.Lit(*concept.Definition)))
			}
			if len(concept.Concept) > 0 {
				definitions(valueSetName, concept.Concept)(group)
			}
		}
	}
}
