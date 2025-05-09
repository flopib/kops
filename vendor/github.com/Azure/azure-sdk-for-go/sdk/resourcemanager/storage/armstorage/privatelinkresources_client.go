// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// ListByStorageAccount - Gets the private link resources that need to be created for a storage account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - options - PrivateLinkResourcesClientListByStorageAccountOptions contains the optional parameters for the PrivateLinkResourcesClient.ListByStorageAccount
//     method.
func (client *PrivateLinkResourcesClient) ListByStorageAccount(ctx context.Context, resourceGroupName string, accountName string, options *PrivateLinkResourcesClientListByStorageAccountOptions) (PrivateLinkResourcesClientListByStorageAccountResponse, error) {
	var err error
	const operationName = "PrivateLinkResourcesClient.ListByStorageAccount"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listByStorageAccountCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return PrivateLinkResourcesClientListByStorageAccountResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkResourcesClientListByStorageAccountResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateLinkResourcesClientListByStorageAccountResponse{}, err
	}
	resp, err := client.listByStorageAccountHandleResponse(httpResp)
	return resp, err
}

// listByStorageAccountCreateRequest creates the ListByStorageAccount request.
func (client *PrivateLinkResourcesClient) listByStorageAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, _ *PrivateLinkResourcesClientListByStorageAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateLinkResources"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByStorageAccountHandleResponse handles the ListByStorageAccount response.
func (client *PrivateLinkResourcesClient) listByStorageAccountHandleResponse(resp *http.Response) (PrivateLinkResourcesClientListByStorageAccountResponse, error) {
	result := PrivateLinkResourcesClientListByStorageAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesClientListByStorageAccountResponse{}, err
	}
	return result, nil
}
