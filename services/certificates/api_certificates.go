// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付平台证书下载服务
//
// 为了确保在定期更换平台证书时，不影响商户使用微信支付的各种功能，微信支付提供API接口供商户下载最新的平台证书。 商户可使用该接口实现平台证书的平滑切换。
//
// API version: 1.0.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package certificates

import (
	"context"
	nethttp "net/http"
	neturl "net/url"

	"github.com/wechat-pay/core"
	"github.com/wechat-pay/core/consts"
	"github.com/wechat-pay/services"
)

type CertificatesApiService services.Service

// DownloadCertificates 获取平台证书列表
//
// 获取商户当前可用的平台证书列表。微信支付提供该接口，帮助商户后台系统实现平台证书的平滑更换。
func (a *CertificatesApiService) DownloadCertificates(ctx context.Context) (resp *DownloadCertificatesResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/certificates"
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

	// Extract DownloadCertificatesResponse from Http Response
	resp = new(DownloadCertificatesResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
