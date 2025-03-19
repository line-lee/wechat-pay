// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付营销系统开放API
//
// 新增立减金api
//
// API version: 3.4.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package cashcoupons

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"

	"github.com/wechat-pay/core"
	"github.com/wechat-pay/core/consts"
	"github.com/wechat-pay/services"
)

type CallBackUrlApiService services.Service

// QueryCallback 查询代金券消息通知地址
//
// 通过此接口查询营销事件通知的回调URL地址
//
// 接口频率：2000QPS
//
// 前置条件：开通营销事件推送产品权限
func (a *CallBackUrlApiService) QueryCallback(ctx context.Context, req QueryCallbackRequest) (resp *Callback, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/callbacks"
	// Make sure All Required Params are properly set
	if req.Mchid == nil {
		return nil, nil, fmt.Errorf("field `Mchid` is required and must be specified in QueryCallbackRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("mchid", core.ParameterToString(*req.Mchid, ""))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract Callback from Http Response
	resp = new(Callback)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// SetCallback 设置代金券消息通知地址
//
// 通过此接口设置营销事件通知的回调URL地址
//
// 接口频率：1000/min
//
// 前置条件：开通营销事件推送产品权限
//
// 可接收营销相关的事件通知:核销通知
//
// 注意：
//
// 1. 该接口只能创建商户号进行设置
func (a *CallBackUrlApiService) SetCallback(ctx context.Context, req SetCallbackRequest) (resp *SetCallbackResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/callbacks"
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

	// Extract SetCallbackResponse from Http Response
	resp = new(SetCallbackResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
