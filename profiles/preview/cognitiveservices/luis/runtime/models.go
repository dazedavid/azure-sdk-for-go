// +build go1.9

// Copyright 2019 Microsoft Corporation
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

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package runtime

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/luis/runtime"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type APIError = original.APIError
type CompositeChildModel = original.CompositeChildModel
type CompositeEntityModel = original.CompositeEntityModel
type EntityModel = original.EntityModel
type EntityWithResolution = original.EntityWithResolution
type EntityWithScore = original.EntityWithScore
type IntentModel = original.IntentModel
type LuisResult = original.LuisResult
type Sentiment = original.Sentiment
type PredictionClient = original.PredictionClient

func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func NewPredictionClient() PredictionClient {
	return original.NewPredictionClient()
}
func NewPredictionClientWithBaseURI(baseURI string) PredictionClient {
	return original.NewPredictionClientWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
