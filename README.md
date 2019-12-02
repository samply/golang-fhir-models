# Golang FHIR Models

This repository contains a FHIR® R4 models for Go. The models consist of Go structs for each resource and data type. The structs are suitable for JSON encoding/decoding. 

## Features

* resources implement the [Marshaler][1] interface
* unmarshal functions are provided for every resource
* enums are provided for every ValueSet used in a [required binding][2], has a computer friendly name and refers only to one CodeSystem

## Usage

In your project, import `github.com/samply/golang-fhir-models/fhir-models/fhir` and you are done.

## TODOs

* [Support Polymorphic Data Elements](https://github.com/samply/golang-fhir-models/issues/1)
* [Support ValueSets Referring to Multiple CodeSystems](https://github.com/samply/golang-fhir-models/issues/2)

## Develop

This repository contains two Go modules, the generated models itself and the generator. Both modules use `go generate` to generate the FHIR models. For `go generate` to work, you have to install the generator first. To do that, run `go install` in the `fhir-models-gen` directory. After that, you can regenerate the FHIR Models under `fhir-models` and the subset of FHIR models under `fhir-models-gen`.

## License

Copyright 2019 The Samply Development Community

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

[1]: <https://golang.org/pkg/encoding/json/#Marshaler>
[2]: <https://www.hl7.org/fhir/terminologies.html#strength>
