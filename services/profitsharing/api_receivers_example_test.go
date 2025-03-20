// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付分账API
//
// 微信支付分账API
//
// API version: 0.0.3

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package profitsharing_test

import (
	"context"
	"log"

	"github.com/line-lee/wechat-pay/core"
	"github.com/line-lee/wechat-pay/core/option"
	"github.com/line-lee/wechat-pay/services/profitsharing"
	"github.com/line-lee/wechat-pay/utils"
)

func ExampleReceiversApiService_AddReceiver() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := profitsharing.ReceiversApiService{Client: client}
	resp, result, err := svc.AddReceiver(ctx,
		profitsharing.AddReceiverRequest{
			Account:        core.String("86693852"),
			Appid:          core.String("wx8888888888888888"),
			CustomRelation: core.String("代理商"),
			Name:           core.String("hu89ohu89ohu89o"),
			RelationType:   profitsharing.RECEIVERRELATIONTYPE_SERVICE_PROVIDER.Ptr(),
			SubAppid:       core.String("wx8888888888888889"),
			SubMchid:       core.String("1900000109"),
			Type:           profitsharing.RECEIVERTYPE_MERCHANT_ID.Ptr(),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call AddReceiver err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleReceiversApiService_DeleteReceiver() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := profitsharing.ReceiversApiService{Client: client}
	resp, result, err := svc.DeleteReceiver(ctx,
		profitsharing.DeleteReceiverRequest{
			Account:  core.String("86693852"),
			Appid:    core.String("wx8888888888888888"),
			SubAppid: core.String("wx8888888888888889"),
			SubMchid: core.String("1900000109"),
			Type:     profitsharing.RECEIVERTYPE_MERCHANT_ID.Ptr(),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call DeleteReceiver err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
