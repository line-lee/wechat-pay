// Copyright 2021 Tencent Inc. All rights reserved.
//
// H5支付
//
// H5支付API
//
// API version: 1.2.3

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package h5

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/line-lee/wechat-pay/core"
	"github.com/line-lee/wechat-pay/core/consts"
	"github.com/line-lee/wechat-pay/services"
	"github.com/line-lee/wechat-pay/services/payments"
)

type H5ApiService services.Service

// CloseOrder 关闭订单
//
// 以下情况需要调用关单接口：
// 1. 商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；
// 2. 系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
func (a *H5ApiService) CloseOrder(ctx context.Context, req CloseOrderRequest) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.OutTradeNo == nil {
		return nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in CloseOrderRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/pay/transactions/out-trade-no/{out_trade_no}/close"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"out_trade_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutTradeNo, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &CloseRequest{
		Mchid: req.Mchid,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return result, err
	}

	return result, nil
}

// Prepay H5支付下单
//
// 商户系统先调用该接口在微信支付服务后台生成预支付交易单，返回正确的预支付交易会话标识后再按Native、JSAPI、APP等不同场景生成交易串调起支付。
func (a *H5ApiService) Prepay(ctx context.Context, req PrepayRequest) (resp *PrepayResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/pay/transactions/h5"
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

	// Extract PrepayResponse from Http Response
	resp = new(PrepayResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryOrderById 微信支付订单号查询订单
//
// 商户可以通过查询订单接口主动查询订单状态
func (a *H5ApiService) QueryOrderById(ctx context.Context, req QueryOrderByIdRequest) (resp *payments.Transaction, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.TransactionId == nil {
		return nil, nil, fmt.Errorf("field `TransactionId` is required and must be specified in QueryOrderByIdRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/pay/transactions/id/{transaction_id}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"transaction_id"+"}", neturl.PathEscape(core.ParameterToString(*req.TransactionId, "")), -1)

	// Make sure All Required Params are properly set
	if req.Mchid == nil {
		return nil, nil, fmt.Errorf("field `Mchid` is required and must be specified in QueryOrderByIdRequest")
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

	// Extract payments.Transaction from Http Response
	resp = new(payments.Transaction)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryOrderByOutTradeNo 商户订单号查询订单
//
// 商户可以通过查询订单接口主动查询订单状态
func (a *H5ApiService) QueryOrderByOutTradeNo(ctx context.Context, req QueryOrderByOutTradeNoRequest) (resp *payments.Transaction, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.OutTradeNo == nil {
		return nil, nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in QueryOrderByOutTradeNoRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/pay/transactions/out-trade-no/{out_trade_no}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"out_trade_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutTradeNo, "")), -1)

	// Make sure All Required Params are properly set
	if req.Mchid == nil {
		return nil, nil, fmt.Errorf("field `Mchid` is required and must be specified in QueryOrderByOutTradeNoRequest")
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

	// Extract payments.Transaction from Http Response
	resp = new(payments.Transaction)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
