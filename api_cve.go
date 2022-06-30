/*
Future Vuls Public API

Future Vuls Public API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// CveApiService CveApi service
type CveApiService service

type ApiCveGetCveDetailRequest struct {
	ctx context.Context
	ApiService *CveApiService
	cveID string
	authorization *string
}

// api key auth
func (r ApiCveGetCveDetailRequest) Authorization(authorization string) ApiCveGetCveDetailRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCveGetCveDetailRequest) Execute() (*CveGetCveDetailResponseBody, *http.Response, error) {
	return r.ApiService.CveGetCveDetailExecute(r)
}

/*
CveGetCveDetail getCveDetail cve

cve detail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cveID Cve ID
 @return ApiCveGetCveDetailRequest
*/
func (a *CveApiService) CveGetCveDetail(ctx context.Context, cveID string) ApiCveGetCveDetailRequest {
	return ApiCveGetCveDetailRequest{
		ApiService: a,
		ctx: ctx,
		cveID: cveID,
	}
}

// Execute executes the request
//  @return CveGetCveDetailResponseBody
func (a *CveApiService) CveGetCveDetailExecute(r ApiCveGetCveDetailRequest) (*CveGetCveDetailResponseBody, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CveGetCveDetailResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CveApiService.CveGetCveDetail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cve/{cveID}"
	localVarPath = strings.Replace(localVarPath, "{"+"cveID"+"}", url.PathEscape(parameterToString(r.cveID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml", "application/gob"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header_Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCveGetCveListRequest struct {
	ctx context.Context
	ApiService *CveApiService
	page *int32
	limit *int32
	offset *int32
	filterServerID *int32
	filterRoleID *int32
	filterPkgID *int32
	filterCpeID *int32
	authorization *string
}

// Page Number
func (r ApiCveGetCveListRequest) Page(page int32) ApiCveGetCveListRequest {
	r.page = &page
	return r
}

// Limit
func (r ApiCveGetCveListRequest) Limit(limit int32) ApiCveGetCveListRequest {
	r.limit = &limit
	return r
}

// Offset
func (r ApiCveGetCveListRequest) Offset(offset int32) ApiCveGetCveListRequest {
	r.offset = &offset
	return r
}

// ServerID filter
func (r ApiCveGetCveListRequest) FilterServerID(filterServerID int32) ApiCveGetCveListRequest {
	r.filterServerID = &filterServerID
	return r
}

// ServerRoleID filter
func (r ApiCveGetCveListRequest) FilterRoleID(filterRoleID int32) ApiCveGetCveListRequest {
	r.filterRoleID = &filterRoleID
	return r
}

// PackageID filter
func (r ApiCveGetCveListRequest) FilterPkgID(filterPkgID int32) ApiCveGetCveListRequest {
	r.filterPkgID = &filterPkgID
	return r
}

// CpeID filter
func (r ApiCveGetCveListRequest) FilterCpeID(filterCpeID int32) ApiCveGetCveListRequest {
	r.filterCpeID = &filterCpeID
	return r
}

// api key auth
func (r ApiCveGetCveListRequest) Authorization(authorization string) ApiCveGetCveListRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCveGetCveListRequest) Execute() (*CveGetCveListResponseBody, *http.Response, error) {
	return r.ApiService.CveGetCveListExecute(r)
}

/*
CveGetCveList getCveList cve

cve list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCveGetCveListRequest
*/
func (a *CveApiService) CveGetCveList(ctx context.Context) ApiCveGetCveListRequest {
	return ApiCveGetCveListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CveGetCveListResponseBody
func (a *CveApiService) CveGetCveListExecute(r ApiCveGetCveListRequest) (*CveGetCveListResponseBody, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CveGetCveListResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CveApiService.CveGetCveList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cves"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.filterServerID != nil {
		localVarQueryParams.Add("filterServerID", parameterToString(*r.filterServerID, ""))
	}
	if r.filterRoleID != nil {
		localVarQueryParams.Add("filterRoleID", parameterToString(*r.filterRoleID, ""))
	}
	if r.filterPkgID != nil {
		localVarQueryParams.Add("filterPkgID", parameterToString(*r.filterPkgID, ""))
	}
	if r.filterCpeID != nil {
		localVarQueryParams.Add("filterCpeID", parameterToString(*r.filterCpeID, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml", "application/gob"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header_Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
