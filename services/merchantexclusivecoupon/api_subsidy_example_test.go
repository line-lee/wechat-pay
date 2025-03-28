// Copyright 2021 Tencent Inc. All rights reserved.
//
// 营销商家券对外API
//
// No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
// API version: 0.0.11

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package merchantexclusivecoupon_test

import (
	"context"
	"log"

	"github.com/line-lee/wechat-pay/core"
	"github.com/line-lee/wechat-pay/core/option"
	"github.com/line-lee/wechat-pay/services/merchantexclusivecoupon"
	"github.com/line-lee/wechat-pay/utils"
)

func ExampleSubsidyApiService_PayReceiptInfo() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := merchantexclusivecoupon.SubsidyApiService{Client: client}
	resp, result, err := svc.PayReceiptInfo(ctx,
		merchantexclusivecoupon.PayReceiptInfoRequest{
			SubsidyReceiptId: core.String("1120200119165100000000000001"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call PayReceiptInfo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleSubsidyApiService_PayReceiptList() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := merchantexclusivecoupon.SubsidyApiService{Client: client}
	resp, result, err := svc.PayReceiptList(ctx,
		merchantexclusivecoupon.PayReceiptListRequest{
			StockId:      core.String("128888000000001"),
			CouponCode:   core.String("ABCD12345678"),
			OutSubsidyNo: core.String("subsidy-abcd-12345678"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call PayReceiptList err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleSubsidyApiService_ReturnReceiptInfo() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := merchantexclusivecoupon.SubsidyApiService{Client: client}
	resp, result, err := svc.ReturnReceiptInfo(ctx,
		merchantexclusivecoupon.ReturnReceiptInfoRequest{
			SubsidyReturnReceiptId: core.String("2120200119165100000000000001"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ReturnReceiptInfo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleSubsidyApiService_SubsidyPay() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := merchantexclusivecoupon.SubsidyApiService{Client: client}
	resp, result, err := svc.SubsidyPay(ctx,
		merchantexclusivecoupon.SubsidyPayRequest{
			StockId:       core.String("128888000000001"),
			CouponCode:    core.String("ABCD12345678"),
			TransactionId: core.String("4200000913202101152566792388"),
			PayerMerchant: core.String("1900000001"),
			PayeeMerchant: core.String("1900000002"),
			Amount:        core.Int64(100),
			Description:   core.String("20210115DESCRIPTION"),
			OutSubsidyNo:  core.String("subsidy-abcd-12345678"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call SubsidyPay err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleSubsidyApiService_SubsidyReturn() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := merchantexclusivecoupon.SubsidyApiService{Client: client}
	resp, result, err := svc.SubsidyReturn(ctx,
		merchantexclusivecoupon.SubsidyReturnRequest{
			StockId:            core.String("128888000000001"),
			CouponCode:         core.String("ABCD12345678"),
			TransactionId:      core.String("4200000913202101152566792388"),
			RefundId:           core.String("50100506732021010105138718375"),
			PayerMerchant:      core.String("1900000001"),
			PayeeMerchant:      core.String("1900000002"),
			Amount:             core.Int64(100),
			Description:        core.String("20210115DESCRIPTION"),
			OutSubsidyReturnNo: core.String("subsidy-abcd-12345678"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call SubsidyReturn err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
