package migrateapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/migrate/mgmt/2018-09-01-preview/migrate"
	"github.com/Azure/go-autorest/autorest"
)

// ErrorsClientAPI contains the set of methods on the ErrorsClient type.
type ErrorsClientAPI interface {
	DeleteError(ctx context.Context, resourceGroupName string, migrateProjectName string, errorName string) (result autorest.Response, err error)
	EnumerateErrors(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string) (result migrate.ErrorCollection, err error)
	GetError(ctx context.Context, resourceGroupName string, migrateProjectName string, errorName string) (result migrate.Error, err error)
}

var _ ErrorsClientAPI = (*migrate.ErrorsClient)(nil)

// MachinesClientAPI contains the set of methods on the MachinesClient type.
type MachinesClientAPI interface {
	EnumerateMachines(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string) (result migrate.MachineCollection, err error)
	GetMachine(ctx context.Context, resourceGroupName string, migrateProjectName string, machineName string) (result migrate.Machine, err error)
}

var _ MachinesClientAPI = (*migrate.MachinesClient)(nil)

// ProjectsClientAPI contains the set of methods on the ProjectsClient type.
type ProjectsClientAPI interface {
	DeleteMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string) (result autorest.Response, err error)
	GetMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string) (result migrate.Project, err error)
	PatchMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string, body migrate.Project) (result migrate.Project, err error)
	PutMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string, body migrate.Project) (result migrate.Project, err error)
	RefreshMigrateProjectSummary(ctx context.Context, resourceGroupName string, migrateProjectName string) (result migrate.RefreshSummaryResult, err error)
	RegisterTool(ctx context.Context, resourceGroupName string, migrateProjectName string, input migrate.RegisterToolInput) (result migrate.RegistrationResult, err error)
}

var _ ProjectsClientAPI = (*migrate.ProjectsClient)(nil)

// SolutionsClientAPI contains the set of methods on the SolutionsClient type.
type SolutionsClientAPI interface {
	DeleteSolution(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string) (result autorest.Response, err error)
	EnumerateSolutions(ctx context.Context, resourceGroupName string, migrateProjectName string) (result migrate.SolutionsCollection, err error)
	GetConfig(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string) (result migrate.SolutionConfig, err error)
	GetSolution(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string) (result migrate.Solution, err error)
	PatchSolution(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, solutionInput migrate.Solution) (result migrate.Solution, err error)
	PutSolution(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, solutionInput migrate.Solution) (result migrate.Solution, err error)
}

var _ SolutionsClientAPI = (*migrate.SolutionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result migrate.OperationResultList, err error)
}

var _ OperationsClientAPI = (*migrate.OperationsClient)(nil)