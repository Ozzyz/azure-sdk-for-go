package search

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

// PrivateLinkResourcesClient is the client that can be used to manage Azure Cognitive Search services and API keys.
type PrivateLinkResourcesClient struct {
	BaseClient
}

// NewPrivateLinkResourcesClient creates an instance of the PrivateLinkResourcesClient client.
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return NewPrivateLinkResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateLinkResourcesClientWithBaseURI creates an instance of the PrivateLinkResourcesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return PrivateLinkResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListSupported gets a list of all supported private link resource types for the given service.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service associated with the specified resource
// group.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client PrivateLinkResourcesClient) ListSupported(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result PrivateLinkResourcesResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourcesClient.ListSupported")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListSupportedPreparer(ctx, resourceGroupName, searchServiceName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.PrivateLinkResourcesClient", "ListSupported", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSupportedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.PrivateLinkResourcesClient", "ListSupported", resp, "Failure sending request")
		return
	}

	result, err = client.ListSupportedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.PrivateLinkResourcesClient", "ListSupported", resp, "Failure responding to request")
		return
	}

	return
}

// ListSupportedPreparer prepares the ListSupported request.
func (client PrivateLinkResourcesClient) ListSupportedPreparer(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/privateLinkResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSupportedSender sends the ListSupported request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkResourcesClient) ListSupportedSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListSupportedResponder handles the response to the ListSupported request. The method always
// closes the http.Response Body.
func (client PrivateLinkResourcesClient) ListSupportedResponder(resp *http.Response) (result PrivateLinkResourcesResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
