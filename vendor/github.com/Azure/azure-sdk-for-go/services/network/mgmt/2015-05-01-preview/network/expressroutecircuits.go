package network

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
	"net/http"
)

// ExpressRouteCircuitsClient is the network Client
type ExpressRouteCircuitsClient struct {
	BaseClient
}

// NewExpressRouteCircuitsClient creates an instance of the ExpressRouteCircuitsClient client.
func NewExpressRouteCircuitsClient(subscriptionID string) ExpressRouteCircuitsClient {
	return NewExpressRouteCircuitsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteCircuitsClientWithBaseURI creates an instance of the ExpressRouteCircuitsClient client.
func NewExpressRouteCircuitsClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteCircuitsClient {
	return ExpressRouteCircuitsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put ExpressRouteCircuit operation creates/updates a ExpressRouteCircuit
//
// resourceGroupName is the name of the resource group. circuitName is the name of the circuit. parameters is
// parameters supplied to the create/delete ExpressRouteCircuit operation
func (client ExpressRouteCircuitsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit) (result ExpressRouteCircuitsCreateOrUpdateFuture, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, circuitName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteCircuitsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) CreateOrUpdateSender(req *http.Request) (future ExpressRouteCircuitsCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) CreateOrUpdateResponder(resp *http.Response) (result ExpressRouteCircuit, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the delete ExpressRouteCircuit operation deletes the specified ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route Circuit.
func (client ExpressRouteCircuitsClient) Delete(ctx context.Context, resourceGroupName string, circuitName string) (result ExpressRouteCircuitsDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, circuitName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExpressRouteCircuitsClient) DeletePreparer(ctx context.Context, resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) DeleteSender(req *http.Request) (future ExpressRouteCircuitsDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get ExpressRouteCircuit operation retreives information about the specified ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the circuit.
func (client ExpressRouteCircuitsClient) Get(ctx context.Context, resourceGroupName string, circuitName string) (result ExpressRouteCircuit, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, circuitName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExpressRouteCircuitsClient) GetPreparer(ctx context.Context, resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) GetResponder(resp *http.Response) (result ExpressRouteCircuit, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List ExpressRouteCircuit opertion retrieves all the ExpressRouteCircuits in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client ExpressRouteCircuitsClient) List(ctx context.Context, resourceGroupName string) (result ExpressRouteCircuitListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.erclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "List", resp, "Failure sending request")
		return
	}

	result.erclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteCircuitsClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListResponder(resp *http.Response) (result ExpressRouteCircuitListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsClient) listNextResults(lastResults ExpressRouteCircuitListResult) (result ExpressRouteCircuitListResult, err error) {
	req, err := lastResults.expressRouteCircuitListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteCircuitsClient) ListComplete(ctx context.Context, resourceGroupName string) (result ExpressRouteCircuitListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// ListAll the List ExpressRouteCircuit opertion retrieves all the ExpressRouteCircuits in a subscription.
func (client ExpressRouteCircuitsClient) ListAll(ctx context.Context) (result ExpressRouteCircuitListResultPage, err error) {
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.erclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.erclr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client ExpressRouteCircuitsClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListAllResponder(resp *http.Response) (result ExpressRouteCircuitListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsClient) listAllNextResults(lastResults ExpressRouteCircuitListResult) (result ExpressRouteCircuitListResult, err error) {
	req, err := lastResults.expressRouteCircuitListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteCircuitsClient) ListAllComplete(ctx context.Context) (result ExpressRouteCircuitListResultIterator, err error) {
	result.page, err = client.ListAll(ctx)
	return
}

// ListArpTable the ListArpTable from ExpressRouteCircuit opertion retrieves the currently advertised arp table
// associated with the ExpressRouteCircuits in a resource group.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the circuit.
func (client ExpressRouteCircuitsClient) ListArpTable(ctx context.Context, resourceGroupName string, circuitName string) (result ExpressRouteCircuitsArpTableListResultPage, err error) {
	result.fn = client.listArpTableNextResults
	req, err := client.ListArpTablePreparer(ctx, resourceGroupName, circuitName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListArpTable", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListArpTableSender(req)
	if err != nil {
		result.ercatlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListArpTable", resp, "Failure sending request")
		return
	}

	result.ercatlr, err = client.ListArpTableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListArpTable", resp, "Failure responding to request")
	}

	return
}

// ListArpTablePreparer prepares the ListArpTable request.
func (client ExpressRouteCircuitsClient) ListArpTablePreparer(ctx context.Context, resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}arpTable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListArpTableSender sends the ListArpTable request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListArpTableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListArpTableResponder handles the response to the ListArpTable request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListArpTableResponder(resp *http.Response) (result ExpressRouteCircuitsArpTableListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listArpTableNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsClient) listArpTableNextResults(lastResults ExpressRouteCircuitsArpTableListResult) (result ExpressRouteCircuitsArpTableListResult, err error) {
	req, err := lastResults.expressRouteCircuitsArpTableListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listArpTableNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListArpTableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listArpTableNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListArpTableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listArpTableNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListArpTableComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteCircuitsClient) ListArpTableComplete(ctx context.Context, resourceGroupName string, circuitName string) (result ExpressRouteCircuitsArpTableListResultIterator, err error) {
	result.page, err = client.ListArpTable(ctx, resourceGroupName, circuitName)
	return
}

// ListRoutesTable the ListRoutesTable from ExpressRouteCircuit opertion retrieves the currently advertised routes
// table associated with the ExpressRouteCircuits in a resource group.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the circuit.
func (client ExpressRouteCircuitsClient) ListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string) (result ExpressRouteCircuitsRoutesTableListResultPage, err error) {
	result.fn = client.listRoutesTableNextResults
	req, err := client.ListRoutesTablePreparer(ctx, resourceGroupName, circuitName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListRoutesTable", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRoutesTableSender(req)
	if err != nil {
		result.ercrtlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListRoutesTable", resp, "Failure sending request")
		return
	}

	result.ercrtlr, err = client.ListRoutesTableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListRoutesTable", resp, "Failure responding to request")
	}

	return
}

// ListRoutesTablePreparer prepares the ListRoutesTable request.
func (client ExpressRouteCircuitsClient) ListRoutesTablePreparer(ctx context.Context, resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}routesTable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRoutesTableSender sends the ListRoutesTable request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListRoutesTableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListRoutesTableResponder handles the response to the ListRoutesTable request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListRoutesTableResponder(resp *http.Response) (result ExpressRouteCircuitsRoutesTableListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listRoutesTableNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsClient) listRoutesTableNextResults(lastResults ExpressRouteCircuitsRoutesTableListResult) (result ExpressRouteCircuitsRoutesTableListResult, err error) {
	req, err := lastResults.expressRouteCircuitsRoutesTableListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listRoutesTableNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListRoutesTableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listRoutesTableNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListRoutesTableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listRoutesTableNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListRoutesTableComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteCircuitsClient) ListRoutesTableComplete(ctx context.Context, resourceGroupName string, circuitName string) (result ExpressRouteCircuitsRoutesTableListResultIterator, err error) {
	result.page, err = client.ListRoutesTable(ctx, resourceGroupName, circuitName)
	return
}

// ListStats the Liststats ExpressRouteCircuit opertion retrieves all the stats from a ExpressRouteCircuits in a
// resource group.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the loadBalancer.
func (client ExpressRouteCircuitsClient) ListStats(ctx context.Context, resourceGroupName string, circuitName string) (result ExpressRouteCircuitsStatsListResultPage, err error) {
	result.fn = client.listStatsNextResults
	req, err := client.ListStatsPreparer(ctx, resourceGroupName, circuitName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListStats", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListStatsSender(req)
	if err != nil {
		result.ercslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListStats", resp, "Failure sending request")
		return
	}

	result.ercslr, err = client.ListStatsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "ListStats", resp, "Failure responding to request")
	}

	return
}

// ListStatsPreparer prepares the ListStats request.
func (client ExpressRouteCircuitsClient) ListStatsPreparer(ctx context.Context, resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}stats", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListStatsSender sends the ListStats request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListStatsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListStatsResponder handles the response to the ListStats request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListStatsResponder(resp *http.Response) (result ExpressRouteCircuitsStatsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listStatsNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsClient) listStatsNextResults(lastResults ExpressRouteCircuitsStatsListResult) (result ExpressRouteCircuitsStatsListResult, err error) {
	req, err := lastResults.expressRouteCircuitsStatsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listStatsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListStatsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listStatsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListStatsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitsClient", "listStatsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListStatsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteCircuitsClient) ListStatsComplete(ctx context.Context, resourceGroupName string, circuitName string) (result ExpressRouteCircuitsStatsListResultIterator, err error) {
	result.page, err = client.ListStats(ctx, resourceGroupName, circuitName)
	return
}
