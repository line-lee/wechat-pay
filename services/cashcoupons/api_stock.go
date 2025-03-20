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
	"strings"

	"github.com/line-lee/wechat-pay/core"
	"github.com/line-lee/wechat-pay/core/consts"
	"github.com/line-lee/wechat-pay/services"
)

type StockApiService services.Service

// CreateCouponStock 创建代金券批次
//
// 商户通过这个接口创建代金券批次
//
// 前置条件：商户开通了代金券产品权限
//
// 是否支持幂等：是
func (a *StockApiService) CreateCouponStock(ctx context.Context, req CreateCouponStockRequest) (resp *CreateCouponStockResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/coupon-stocks"
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

	// Extract CreateCouponStockResponse from Http Response
	resp = new(CreateCouponStockResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ListAvailableMerchants 查询代金券可用商户
//
// 查使用代金券业务的商户询代金券可用商户
//
// 接口频率：1000/min
//
// 前置条件：非制券方查询主批次信息以获取发券商户列表，调用方只能查询自己制的批次、自己可发的批次
//
// 是否支持幂等：是
func (a *StockApiService) ListAvailableMerchants(ctx context.Context, req ListAvailableMerchantsRequest) (resp *AvailableMerchantCollection, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in ListAvailableMerchantsRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/stocks/{stock_id}/merchants"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

	// Make sure All Required Params are properly set
	if req.Offset == nil {
		return nil, nil, fmt.Errorf("field `Offset` is required and must be specified in ListAvailableMerchantsRequest")
	}
	if req.Limit == nil {
		return nil, nil, fmt.Errorf("field `Limit` is required and must be specified in ListAvailableMerchantsRequest")
	}
	if req.StockCreatorMchid == nil {
		return nil, nil, fmt.Errorf("field `StockCreatorMchid` is required and must be specified in ListAvailableMerchantsRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("offset", core.ParameterToString(*req.Offset, ""))
	localVarQueryParams.Add("limit", core.ParameterToString(*req.Limit, ""))
	localVarQueryParams.Add("stock_creator_mchid", core.ParameterToString(*req.StockCreatorMchid, ""))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract AvailableMerchantCollection from Http Response
	resp = new(AvailableMerchantCollection)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ListAvailableSingleitems 查询可核销商品编码
//
// 查询可核销商品编码
//
// 接口频率：600/min
//
// 前置条件：调用方只能查询自己制的批次、自己可发的批次的可核销商品编码
//
// 是否支持幂等：是
func (a *StockApiService) ListAvailableSingleitems(ctx context.Context, req ListAvailableSingleitemsRequest) (resp *AvailableSingleitemCollection, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in ListAvailableSingleitemsRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/stocks/{stock_id}/items"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

	// Make sure All Required Params are properly set
	if req.Offset == nil {
		return nil, nil, fmt.Errorf("field `Offset` is required and must be specified in ListAvailableSingleitemsRequest")
	}
	if req.Limit == nil {
		return nil, nil, fmt.Errorf("field `Limit` is required and must be specified in ListAvailableSingleitemsRequest")
	}
	if req.StockCreatorMchid == nil {
		return nil, nil, fmt.Errorf("field `StockCreatorMchid` is required and must be specified in ListAvailableSingleitemsRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("offset", core.ParameterToString(*req.Offset, ""))
	localVarQueryParams.Add("limit", core.ParameterToString(*req.Limit, ""))
	localVarQueryParams.Add("stock_creator_mchid", core.ParameterToString(*req.StockCreatorMchid, ""))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract AvailableSingleitemCollection from Http Response
	resp = new(AvailableSingleitemCollection)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ListStocks 条件查询批次列表
//
// 商户通过这个接口可以查询创建的批次列表
//
// 接口频率：1000/s
//
// 前置条件：商户开通了代金券产品权限，并创建过代金券
func (a *StockApiService) ListStocks(ctx context.Context, req ListStocksRequest) (resp *StockCollection, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/stocks"
	// Make sure All Required Params are properly set
	if req.Offset == nil {
		return nil, nil, fmt.Errorf("field `Offset` is required and must be specified in ListStocksRequest")
	}
	if req.Limit == nil {
		return nil, nil, fmt.Errorf("field `Limit` is required and must be specified in ListStocksRequest")
	}
	if req.StockCreatorMchid == nil {
		return nil, nil, fmt.Errorf("field `StockCreatorMchid` is required and must be specified in ListStocksRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("offset", core.ParameterToString(*req.Offset, ""))
	localVarQueryParams.Add("limit", core.ParameterToString(*req.Limit, ""))
	localVarQueryParams.Add("stock_creator_mchid", core.ParameterToString(*req.StockCreatorMchid, ""))
	if req.CreateStartTime != nil {
		localVarQueryParams.Add("create_start_time", core.ParameterToString(*req.CreateStartTime, ""))
	}
	if req.CreateEndTime != nil {
		localVarQueryParams.Add("create_end_time", core.ParameterToString(*req.CreateEndTime, ""))
	}
	if req.Status != nil {
		localVarQueryParams.Add("status", core.ParameterToString(*req.Status, ""))
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

	// Extract StockCollection from Http Response
	resp = new(StockCollection)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// PauseStock 暂停批次
//
// 通过此接口暂停代金券批次
//
// 接口频率：1000/min
//
// 前置条件：成功激活代金券批次
//
// 注意事项：
//
// 1. 暂停后，该代金券批次暂停发放，用户无法通过任何渠道再领取该批次的券
func (a *StockApiService) PauseStock(ctx context.Context, req PauseStockRequest) (resp *PauseStockResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in PauseStockRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/stocks/{stock_id}/pause"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &PauseStockBody{
		StockCreatorMchid: req.StockCreatorMchid,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PauseStockResponse from Http Response
	resp = new(PauseStockResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryStock 查询批次详情
//
// 查询代金券批次详情信息
//
// 接口频率：500qps
//
// 前置条件：已创建代金券批次
//
// 注意：
//
// 1. 该接口支持批次创建商户号与批次发放商户调用
func (a *StockApiService) QueryStock(ctx context.Context, req QueryStockRequest) (resp *Stock, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in QueryStockRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/stocks/{stock_id}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

	// Make sure All Required Params are properly set
	if req.StockCreatorMchid == nil {
		return nil, nil, fmt.Errorf("field `StockCreatorMchid` is required and must be specified in QueryStockRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("stock_creator_mchid", core.ParameterToString(*req.StockCreatorMchid, ""))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract Stock from Http Response
	resp = new(Stock)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// RefundFlow 下载批次退款明细
//
// 通过此接口获取指定批次的退款明细数据，包括订单号、单品信息、银行流水号等，用于对账/数据分析
//
// 接口频率：1000/min
//
// 前置条件：批次过期10小时
//
// 注意：
//
// 1. 账单文件的下载地址的有效时间为30s。
// 2. 强烈建议商户将实际账单文件的哈希值和之前从接口获取到的哈希值进行比对，以确认数据的完整性。
// 3. 该接口响应的信息请求头中不包含微信接口响应的签名值，因此需要跳过验签的流程
// 4. 该接口需要活动结束后次日10点才可以下载
// 5. 该接口只能使用创建活动方进行调用”
func (a *StockApiService) RefundFlow(ctx context.Context, req RefundFlowRequest) (resp *RefundFlowResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in RefundFlowRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/stocks/{stock_id}/refund-flow"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

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

	// Extract RefundFlowResponse from Http Response
	resp = new(RefundFlowResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// RestartStock 重启批次
//
// 商户通过这个接口重新激活已暂停的代金券批次
//
// 接口频率：1000/min
//
// 前置条件：商户开通了代金券产品权限,且批次处于暂停状态
func (a *StockApiService) RestartStock(ctx context.Context, req RestartStockRequest) (resp *RestartStockResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in RestartStockRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/stocks/{stock_id}/restart"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &RestartStockBody{
		StockCreatorMchid: req.StockCreatorMchid,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract RestartStockResponse from Http Response
	resp = new(RestartStockResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// StartStock 激活开启批次
//
// 通过此接口激活代金券批次
//
// 前置条件：成功创建代金券批次
//
// 注意事项：
//
// 1. 如果是预充值代金券，激活时会从商户账户余额中锁定本批次的营销资金
func (a *StockApiService) StartStock(ctx context.Context, req StartStockRequest) (resp *StartStockResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in StartStockRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/stocks/{stock_id}/start"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &StartStockBody{
		StockCreatorMchid: req.StockCreatorMchid,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract StartStockResponse from Http Response
	resp = new(StartStockResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// StopStock 终止批次
//
// 商户通过这个接口终止未激活的代金券批次
//
// 前置条件：商户开通了代金券产品权限,且批次未激活
func (a *StockApiService) StopStock(ctx context.Context, req StopStockRequest) (resp *StopStockResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in StopStockRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/stocks/{stock_id}/stop"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &StopStockBody{
		StockCreatorMchid: req.StockCreatorMchid,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract StopStockResponse from Http Response
	resp = new(StopStockResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// UseFlow 下载批次核销明细
//
// 通过此接口获取指定批次的核销明细数据，包括订单号、单品信息、银行流水号等，用于对账/数据分析
//
// 接口频率：1000/min
//
// 前置条件：批次过期10小时
//
// 注意：
//
// 1. 账单文件的下载地址的有效时间为30s。
// 2. 强烈建议商户将实际账单文件的哈希值和之前从接口获取到的哈希值进行比对，以确认数据的完整性。
// 3. 该接口响应的信息请求头中不包含微信接口响应的签名值，因此需要跳过验签的流程。
// 4. 该接口需要活动结束后次日10点才可以下载。
// 5. 该接口只能使用创建活动方进行调用。
func (a *StockApiService) UseFlow(ctx context.Context, req UseFlowRequest) (resp *UseFlowResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in UseFlowRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/stocks/{stock_id}/use-flow"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

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

	// Extract UseFlowResponse from Http Response
	resp = new(UseFlowResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
