//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.7.6, generator: @autorest/go@4.0.0-preview.42)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type CommitOrRollbackBackupClient struct {
	pl runtime.Pipeline
}

// NewCommitOrRollbackBackupClient creates a new instance of CommitOrRollbackBackupClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewCommitOrRollbackBackupClient(pl runtime.Pipeline) *CommitOrRollbackBackupClient {
	client := &CommitOrRollbackBackupClient{
		pl: pl,
	}
	return client
}

// Cancel - Cancel the operation. Poll the LRO to get the final status.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique Id of this LRO operation
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - CommitOrRollbackBackupClientCancelOptions contains the optional parameters for the CommitOrRollbackBackupClient.Cancel
// method.
func (client *CommitOrRollbackBackupClient) Cancel(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters CancelRequest, options *CommitOrRollbackBackupClientCancelOptions) (CommitOrRollbackBackupClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return CommitOrRollbackBackupClientCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CommitOrRollbackBackupClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CommitOrRollbackBackupClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return client.cancelHandleResponse(resp)
}

// cancelCreateRequest creates the Cancel request.
func (client *CommitOrRollbackBackupClient) cancelCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters CancelRequest, options *CommitOrRollbackBackupClientCancelOptions) (*policy.Request, error) {
	urlPath := "/plugin/commitOrRollbackBackupOperations/{operationId}:cancel"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// cancelHandleResponse handles the Cancel response.
func (client *CommitOrRollbackBackupClient) cancelHandleResponse(resp *http.Response) (CommitOrRollbackBackupClientCancelResponse, error) {
	result := CommitOrRollbackBackupClientCancelResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return CommitOrRollbackBackupClientCancelResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommitOrRollbackBackupResponse); err != nil {
		return CommitOrRollbackBackupClientCancelResponse{}, err
	}
	return result, nil
}

// Get - Gets the status of a Backup LRO.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique Id of this LRO operation
// subscriptionID - SubscriptionId of the resource
// resourceID - Unique id of the resource
// taskID - Unique id of the current task
// xmsClientRequestID - Correlation request Id for tracking a particular request.
// options - CommitOrRollbackBackupClientGetOptions contains the optional parameters for the CommitOrRollbackBackupClient.Get
// method.
func (client *CommitOrRollbackBackupClient) Get(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, options *CommitOrRollbackBackupClientGetOptions) (CommitOrRollbackBackupClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, options)
	if err != nil {
		return CommitOrRollbackBackupClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CommitOrRollbackBackupClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CommitOrRollbackBackupClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CommitOrRollbackBackupClient) getCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, options *CommitOrRollbackBackupClientGetOptions) (*policy.Request, error) {
	urlPath := "/plugin/commitOrRollbackBackupOperations/{operationId}"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CommitOrRollbackBackupClient) getHandleResponse(resp *http.Response) (CommitOrRollbackBackupClientGetResponse, error) {
	result := CommitOrRollbackBackupClientGetResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return CommitOrRollbackBackupClientGetResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommitOrRollbackBackupResponse); err != nil {
		return CommitOrRollbackBackupClientGetResponse{}, err
	}
	return result, nil
}

// RefreshTokens - Refresh tokens for a given operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique Id of this LRO operation
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - CommitOrRollbackBackupClientRefreshTokensOptions contains the optional parameters for the CommitOrRollbackBackupClient.RefreshTokens
// method.
func (client *CommitOrRollbackBackupClient) RefreshTokens(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters RefreshTokensRequest, options *CommitOrRollbackBackupClientRefreshTokensOptions) (CommitOrRollbackBackupClientRefreshTokensResponse, error) {
	req, err := client.refreshTokensCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return CommitOrRollbackBackupClientRefreshTokensResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CommitOrRollbackBackupClientRefreshTokensResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CommitOrRollbackBackupClientRefreshTokensResponse{}, runtime.NewResponseError(resp)
	}
	return client.refreshTokensHandleResponse(resp)
}

// refreshTokensCreateRequest creates the RefreshTokens request.
func (client *CommitOrRollbackBackupClient) refreshTokensCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters RefreshTokensRequest, options *CommitOrRollbackBackupClientRefreshTokensOptions) (*policy.Request, error) {
	urlPath := "/plugin/commitOrRollbackBackupOperations/{operationId}:refreshTokens"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// refreshTokensHandleResponse handles the RefreshTokens response.
func (client *CommitOrRollbackBackupClient) refreshTokensHandleResponse(resp *http.Response) (CommitOrRollbackBackupClientRefreshTokensResponse, error) {
	result := CommitOrRollbackBackupClientRefreshTokensResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return CommitOrRollbackBackupClientRefreshTokensResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommitOrRollbackBackupResponse); err != nil {
		return CommitOrRollbackBackupClientRefreshTokensResponse{}, err
	}
	return result, nil
}

