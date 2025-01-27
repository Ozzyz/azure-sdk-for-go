package visualstudio

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
	"net/http"
)

// ProjectsClient is the use these APIs to manage Visual Studio Team Services resources through the Azure Resource
// Manager. All task operations conform to the HTTP/1.1 protocol specification and each operation returns an
// x-ms-request-id header that can be used to obtain information about the request. You must make sure that requests
// made to these resources are secure. For more information, see https://docs.microsoft.com/en-us/rest/api/index.
type ProjectsClient struct {
	BaseClient
}

// NewProjectsClient creates an instance of the ProjectsClient client.
func NewProjectsClient(subscriptionID string) ProjectsClient {
	return NewProjectsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProjectsClientWithBaseURI creates an instance of the ProjectsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProjectsClientWithBaseURI(baseURI string, subscriptionID string) ProjectsClient {
	return ProjectsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a Team Services project in the collection with the specified name. 'VersionControlOption' and
// 'ProcessTemplateId' must be specified in the resource properties. Valid values for VersionControlOption: Git, Tfvc.
// Valid values for ProcessTemplateId: 6B724908-EF14-45CF-84F8-768B5384DA45, ADCC42AB-9882-485E-A3ED-7678F01F66BC,
// 27450541-8E31-4150-9947-DC59F998FC01 (these IDs correspond to Scrum, Agile, and CMMI process templates).
// Parameters:
// body - the request data.
// resourceGroupName - name of the resource group within the Azure subscription.
// rootResourceName - name of the Team Services account.
// resourceName - name of the Team Services project.
// validating - this parameter is ignored and should be set to an empty string.
func (client ProjectsClient) Create(ctx context.Context, body ProjectResource, resourceGroupName string, rootResourceName string, resourceName string, validating string) (result ProjectsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, body, resourceGroupName, rootResourceName, resourceName, validating)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ProjectsClient) CreatePreparer(ctx context.Context, body ProjectResource, resourceGroupName string, rootResourceName string, resourceName string, validating string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"rootResourceName":  autorest.Encode("path", rootResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(validating) > 0 {
		queryParameters["validating"] = autorest.Encode("query", validating)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{rootResourceName}/project/{resourceName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) CreateSender(req *http.Request) (future ProjectsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ProjectsClient) (pr ProjectResource, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "visualstudio.ProjectsCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("visualstudio.ProjectsCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		pr.Response.Response, err = future.GetResult(sender)
		if pr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "visualstudio.ProjectsCreateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && pr.Response.Response.StatusCode != http.StatusNoContent {
			pr, err = client.CreateResponder(pr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "visualstudio.ProjectsCreateFuture", "Result", pr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ProjectsClient) CreateResponder(resp *http.Response) (result ProjectResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the details of a Team Services project resource.
// Parameters:
// resourceGroupName - name of the resource group within the Azure subscription.
// rootResourceName - name of the Team Services account.
// resourceName - name of the Team Services project.
func (client ProjectsClient) Get(ctx context.Context, resourceGroupName string, rootResourceName string, resourceName string) (result ProjectResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, rootResourceName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProjectsClient) GetPreparer(ctx context.Context, resourceGroupName string, rootResourceName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"rootResourceName":  autorest.Encode("path", rootResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{rootResourceName}/project/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProjectsClient) GetResponder(resp *http.Response) (result ProjectResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetJobStatus gets the status of the project resource creation job.
// Parameters:
// resourceGroupName - name of the resource group within the Azure subscription.
// rootResourceName - name of the Team Services account.
// resourceName - name of the Team Services project.
// subContainerName - this parameter should be set to the resourceName.
// operation - the operation type. The only supported value is 'put'.
// jobID - the job identifier.
func (client ProjectsClient) GetJobStatus(ctx context.Context, resourceGroupName string, rootResourceName string, resourceName string, subContainerName string, operation string, jobID *uuid.UUID) (result ProjectResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.GetJobStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetJobStatusPreparer(ctx, resourceGroupName, rootResourceName, resourceName, subContainerName, operation, jobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "GetJobStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetJobStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "GetJobStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetJobStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "GetJobStatus", resp, "Failure responding to request")
		return
	}

	return
}

// GetJobStatusPreparer prepares the GetJobStatus request.
func (client ProjectsClient) GetJobStatusPreparer(ctx context.Context, resourceGroupName string, rootResourceName string, resourceName string, subContainerName string, operation string, jobID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"rootResourceName":  autorest.Encode("path", rootResourceName),
		"subContainerName":  autorest.Encode("path", subContainerName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"operation":   autorest.Encode("query", operation),
	}
	if jobID != nil {
		queryParameters["jobId"] = autorest.Encode("query", *jobID)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{rootResourceName}/project/{resourceName}/subContainers/{subContainerName}/status", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetJobStatusSender sends the GetJobStatus request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) GetJobStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetJobStatusResponder handles the response to the GetJobStatus request. The method always
// closes the http.Response Body.
func (client ProjectsClient) GetJobStatusResponder(resp *http.Response) (result ProjectResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets all Visual Studio Team Services project resources created in the specified Team Services
// account.
// Parameters:
// resourceGroupName - name of the resource group within the Azure subscription.
// rootResourceName - name of the Team Services account.
func (client ProjectsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, rootResourceName string) (result ProjectResourceListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, rootResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ProjectsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, rootResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"rootResourceName":  autorest.Encode("path", rootResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{rootResourceName}/project", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ProjectsClient) ListByResourceGroupResponder(resp *http.Response) (result ProjectResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates the tags of the specified Team Services project.
// Parameters:
// resourceGroupName - name of the resource group within the Azure subscription.
// body - the request data.
// rootResourceName - name of the Team Services account.
// resourceName - name of the Team Services project.
func (client ProjectsClient) Update(ctx context.Context, resourceGroupName string, body ProjectResource, rootResourceName string, resourceName string) (result ProjectResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, body, rootResourceName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ProjectsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, body ProjectResource, rootResourceName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"rootResourceName":  autorest.Encode("path", rootResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{rootResourceName}/project/{resourceName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ProjectsClient) UpdateResponder(resp *http.Response) (result ProjectResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
