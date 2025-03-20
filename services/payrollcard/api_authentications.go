// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微工卡接口文档
//
// 服务商通过本API文档提供的接口，查询商户和微工卡的授权关系、生成预授权的token口令、核身预下单、核身结果的查询等。
//
// API version: 1.5.2

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package payrollcard

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/line-lee/wechat-pay/core"
	"github.com/line-lee/wechat-pay/core/consts"
	"github.com/line-lee/wechat-pay/services"
)

type AuthenticationsApiService services.Service

// GetAuthentication 获取核身结果
//
// 按商户拉起核身时预下单的单号获取该次微工卡核身的结果
func (a *AuthenticationsApiService) GetAuthentication(ctx context.Context, req GetAuthenticationRequest) (resp *AuthenticationEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.AuthenticateNumber == nil {
		return nil, nil, fmt.Errorf("field `AuthenticateNumber` is required and must be specified in GetAuthenticationRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/payroll-card/authentications/{authenticate_number}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"authenticate_number"+"}", neturl.PathEscape(core.ParameterToString(*req.AuthenticateNumber, "")), -1)

	// Make sure All Required Params are properly set
	if req.SubMchid == nil {
		return nil, nil, fmt.Errorf("field `SubMchid` is required and must be specified in GetAuthenticationRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("sub_mchid", core.ParameterToString(*req.SubMchid, ""))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract AuthenticationEntity from Http Response
	resp = new(AuthenticationEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ListAuthentications 查询核身记录
//
// 查询指定用户指定日期微工卡核身记录，查询结果仅展示核身成功的记录
func (a *AuthenticationsApiService) ListAuthentications(ctx context.Context, req ListAuthenticationsRequest) (resp *ListAuthenticationsResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/payroll-card/authentications"
	// Make sure All Required Params are properly set
	if req.Openid == nil {
		return nil, nil, fmt.Errorf("field `Openid` is required and must be specified in ListAuthenticationsRequest")
	}
	if req.SubMchid == nil {
		return nil, nil, fmt.Errorf("field `SubMchid` is required and must be specified in ListAuthenticationsRequest")
	}
	if req.AuthenticateDate == nil {
		return nil, nil, fmt.Errorf("field `AuthenticateDate` is required and must be specified in ListAuthenticationsRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("openid", core.ParameterToString(*req.Openid, ""))
	if req.Appid != nil {
		localVarQueryParams.Add("appid", core.ParameterToString(*req.Appid, ""))
	}
	if req.SubAppid != nil {
		localVarQueryParams.Add("sub_appid", core.ParameterToString(*req.SubAppid, ""))
	}
	localVarQueryParams.Add("sub_mchid", core.ParameterToString(*req.SubMchid, ""))
	localVarQueryParams.Add("authenticate_date", core.ParameterToString(*req.AuthenticateDate, ""))
	if req.AuthenticateState != nil {
		localVarQueryParams.Add("authenticate_state", core.ParameterToString(*req.AuthenticateState, ""))
	}
	if req.Offset != nil {
		localVarQueryParams.Add("offset", core.ParameterToString(*req.Offset, ""))
	}
	if req.Limit != nil {
		localVarQueryParams.Add("limit", core.ParameterToString(*req.Limit, ""))
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract ListAuthenticationsResponse from Http Response
	resp = new(ListAuthenticationsResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// PreOrderAuthentication 微工卡核身预下单
//
// 服务商在拉起微工卡前端服务给用户做微工卡核身前，需要先调用本接口预下单，下单成功后才能进行核身
func (a *AuthenticationsApiService) PreOrderAuthentication(ctx context.Context, req PreOrderAuthenticationRequest) (resp *PreOrderAuthenticationResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/payroll-card/authentications/pre-order"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PreOrderAuthenticationResponse from Http Response
	resp = new(PreOrderAuthenticationResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// PreOrderAuthenticationWithAuth 微工卡核身预下单（流程中完成授权）
//
// 本接口适用于用户需同步完成服务开通、授权及身份核验的场景。在拉起微工卡前端服务为用户核身前，需调用本接口预下单，下单成功后才能进行核身。如此时用户未开通微工卡服务或未完成对商户的授权，则先完成开通、授权，同步完成身份核验，并提供可信的核验结果。
func (a *AuthenticationsApiService) PreOrderAuthenticationWithAuth(ctx context.Context, req PreOrderAuthenticationWithAuthRequest) (resp *PreOrderAuthenticationWithAuthResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// 对请求中敏感字段进行加密
	encReq := req.Clone()
	encryptCertificate, err := a.Client.EncryptRequest(ctx, encReq)
	if err != nil {
		return nil, nil, fmt.Errorf("encrypt request failed: %v", err)
	}

	if encryptCertificate != "" {
		localVarHeaderParams.Set(consts.WechatPaySerial, encryptCertificate)
	}
	req = *encReq

	localVarPath := consts.WechatPayAPIServer + "/v3/payroll-card/authentications/pre-order-with-auth"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PreOrderAuthenticationWithAuthResponse from Http Response
	resp = new(PreOrderAuthenticationWithAuthResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
