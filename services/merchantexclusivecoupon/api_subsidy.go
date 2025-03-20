// Copyright 2021 Tencent Inc. All rights reserved.
//
// 营销商家券对外API
//
// No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
// API version: 0.0.11

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package merchantexclusivecoupon

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

type SubsidyApiService services.Service

// PayReceiptInfo 查询商家券营销补差付款单详情
//
// 查询商家券营销补差付款单详情
func (a *SubsidyApiService) PayReceiptInfo(ctx context.Context, req PayReceiptInfoRequest) (resp *SubsidyPayReceipt, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.SubsidyReceiptId == nil {
		return nil, nil, fmt.Errorf("field `SubsidyReceiptId` is required and must be specified in PayReceiptInfoRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/subsidy/pay-receipts/{subsidy_receipt_id}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"subsidy_receipt_id"+"}", neturl.PathEscape(core.ParameterToString(*req.SubsidyReceiptId, "")), -1)

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

	// Extract SubsidyPayReceipt from Http Response
	resp = new(SubsidyPayReceipt)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// PayReceiptList 查询商家券营销补差付款单列表
//
// 查询商家券营销补差付款单列表
func (a *SubsidyApiService) PayReceiptList(ctx context.Context, req PayReceiptListRequest) (resp *SubsidyPayReceiptListResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/subsidy/pay-receipts"
	// Make sure All Required Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in PayReceiptListRequest")
	}
	if req.CouponCode == nil {
		return nil, nil, fmt.Errorf("field `CouponCode` is required and must be specified in PayReceiptListRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("stock_id", core.ParameterToString(*req.StockId, ""))
	localVarQueryParams.Add("coupon_code", core.ParameterToString(*req.CouponCode, ""))
	if req.OutSubsidyNo != nil {
		localVarQueryParams.Add("out_subsidy_no", core.ParameterToString(*req.OutSubsidyNo, ""))
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

	// Extract SubsidyPayReceiptListResponse from Http Response
	resp = new(SubsidyPayReceiptListResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ReturnReceiptInfo 查询商家券营销补差回退单详情
//
// 查询商家券营销补差回退单详情
func (a *SubsidyApiService) ReturnReceiptInfo(ctx context.Context, req ReturnReceiptInfoRequest) (resp *SubsidyReturnReceipt, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.SubsidyReturnReceiptId == nil {
		return nil, nil, fmt.Errorf("field `SubsidyReturnReceiptId` is required and must be specified in ReturnReceiptInfoRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/subsidy/return-receipts/{subsidy_return_receipt_id}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"subsidy_return_receipt_id"+"}", neturl.PathEscape(core.ParameterToString(*req.SubsidyReturnReceiptId, "")), -1)

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

	// Extract SubsidyReturnReceipt from Http Response
	resp = new(SubsidyReturnReceipt)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// SubsidyPay 商家券营销补差付款
//
// 适用接口场景：服务商/商圈给核销了商家券的商户做营销资金补差；
// 接口适用对象：普通服务商、普通直连商户、渠道商；
// 前置条件：商家必须核销了商家券且发起了微信支付收款；
// 是否支持幂等：是；
func (a *SubsidyApiService) SubsidyPay(ctx context.Context, req SubsidyPayRequest) (resp *SubsidyPayReceipt, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/subsidy/pay-receipts"
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

	// Extract SubsidyPayReceipt from Http Response
	resp = new(SubsidyPayReceipt)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// SubsidyReturn 商家券营销补差回退
//
// 适用接口场景：服务商通过该接口可回退补差款；
// 接口适用对象：普通服务商、普通直连商户、渠道商；
// 前置条件：进行补差的微信支付订单发起了退款，且回退金额不得超过补差金额；
// 是否支持幂等：是；
func (a *SubsidyApiService) SubsidyReturn(ctx context.Context, req SubsidyReturnRequest) (resp *SubsidyReturnReceipt, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/subsidy/return-receipts"
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

	// Extract SubsidyReturnReceipt from Http Response
	resp = new(SubsidyReturnReceipt)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
