// Copyright 2021 Tencent Inc. All rights reserved.
//
// 爱心餐对外API
//
// 微信支付爱心餐公益计划旨在面向深圳市的市政一线环卫工人提供每周一餐的1分钱用餐公益服务。在受助端，微信支付联动上千家餐饮门店关爱特殊人群，通过微信支付数字化能力将人群身份认证与公益福利领用全流程线上化，实现公益福利精准到人。在捐赠端，微信支付发挥连接优势与平台能力，结合用户就餐场景通过爱心餐一块捐插件让用户可在点餐时顺手捐1元，带动更多社会力量致谢城市美容师。
//
// API version: 0.0.4

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package lovefeast

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

type OrdersApiService services.Service

// GetByUser 查询用户捐赠单详情
//
// 接口介绍：商户根据商户订单号与用户标识查询捐赠单详情
// 适用对象：服务商、直连商户
func (a *OrdersApiService) GetByUser(ctx context.Context, req GetByUserRequest) (resp *OrdersEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.Openid == nil {
		return nil, nil, fmt.Errorf("field `Openid` is required and must be specified in GetByUserRequest")
	}
	if req.OutTradeNo == nil {
		return nil, nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in GetByUserRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/lovefeast/users/{openid}/orders/out-trade-no/{out_trade_no}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"openid"+"}", neturl.PathEscape(core.ParameterToString(*req.Openid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"out_trade_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutTradeNo, "")), -1)

	// Make sure All Required Params are properly set

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract OrdersEntity from Http Response
	resp = new(OrdersEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ListByUser 查询用户捐赠单列表
//
// 接口介绍：商户可根据品牌ID与用户标识查询捐赠单列表
// 适用对象：服务商、直连商户
func (a *OrdersApiService) ListByUser(ctx context.Context, req ListByUserRequest) (resp *OrdersListByUserResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.Openid == nil {
		return nil, nil, fmt.Errorf("field `Openid` is required and must be specified in ListByUserRequest")
	}
	if req.BrandId == nil {
		return nil, nil, fmt.Errorf("field `BrandId` is required and must be specified in ListByUserRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/lovefeast/users/{openid}/orders/brand-id/{brand_id}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"openid"+"}", neturl.PathEscape(core.ParameterToString(*req.Openid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"brand_id"+"}", neturl.PathEscape(core.ParameterToString(*req.BrandId, "")), -1)

	// Make sure All Required Params are properly set
	if req.Limit == nil {
		return nil, nil, fmt.Errorf("field `Limit` is required and must be specified in ListByUserRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("limit", core.ParameterToString(*req.Limit, ""))
	if req.Offset != nil {
		localVarQueryParams.Add("offset", core.ParameterToString(*req.Offset, ""))
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

	// Extract OrdersListByUserResponse from Http Response
	resp = new(OrdersListByUserResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
