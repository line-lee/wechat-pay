// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付分账API
//
// 微信支付分账API
//
// API version: 0.0.3

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package profitsharing

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"

	"github.com/wechat-pay/core"
	"github.com/wechat-pay/core/consts"
	"github.com/wechat-pay/services"
)

type ReceiversApiService services.Service

// AddReceiver 添加分账接收方API
//
// 商户发起添加分账接收方请求，建立分账接收方列表。后续可通过发起分账请求，将分账方商户结算后的资金，分到该分账接收方
//
//
// 错误码列表
// 名称|状态码|描述示例|原因|解决方案
// |-|-|-|-|-|
// SYSTEM_ERROR|500|系统错误|系统超时|系统异常，请使用相同参数稍后重新调用
// PARAM_ERROR|400|参数格式不正确|请求参数不符合参数格式|请使用正确的参数重新调用
// INVALID_REQUEST|400|无效请求|请确认分账接收方是否存在|请确认分账接收方是否存在
// NO_AUTH|403|商户无权限|未开通分账权限|请开通商户号分账权限
// RATELIMIT_EXCEED|429|添加接收方频率过高|接口有频率限制|请降低频率后重试
func (a *ReceiversApiService) AddReceiver(ctx context.Context, req AddReceiverRequest) (resp *AddReceiverResponse, result *core.APIResult, err error) {
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

	localVarPath := consts.WechatPayAPIServer + "/v3/profitsharing/receivers/add"
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

	// Extract AddReceiverResponse from Http Response
	resp = new(AddReceiverResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	// 对应答中隐私字段进行解密
	err = a.Client.DecryptResponse(ctx, resp)
	if err != nil {
		return resp, result, err
	}

	return resp, result, nil
}

// DeleteReceiver 删除分账接收方API
//
// 商户发起删除分账接收方请求。删除后，不支持将分账方商户结算后的资金，分到该分账接收方
//
//
// 错误码列表
// 名称|状态码|描述示例|原因|解决方案
// |-|-|-|-|-|
// SYSTEM_ERROR|500|系统错误|系统超时|系统异常，请使用相同参数稍后重新调用
// PARAM_ERROR|400|参数格式不正确|请求参数不符合参数格式|请使用正确的参数重新调用
// INVALID_REQUEST|400|无效请求|请确认分账接收方是否存在|请确认分账接收方是否存在
// NO_AUTH|403|商户无权限|未开通分账权限|请开通商户号分账权限
// RATELIMIT_EXCEED|429|删除接收方频率过高|接口有频率限制|请降低频率后重试
func (a *ReceiversApiService) DeleteReceiver(ctx context.Context, req DeleteReceiverRequest) (resp *DeleteReceiverResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/profitsharing/receivers/delete"
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

	// Extract DeleteReceiverResponse from Http Response
	resp = new(DeleteReceiverResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
