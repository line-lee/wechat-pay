// Copyright 2021 Tencent Inc. All rights reserved.
//
// 服务商批量转账API
//
// No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
// API version: 0.0.2

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package partnertransferbatch_test

import (
	"context"
	"log"

	"github.com/line-lee/wechat-pay/core"
	"github.com/line-lee/wechat-pay/core/option"
	"github.com/line-lee/wechat-pay/services/partnertransferbatch"
	"github.com/line-lee/wechat-pay/utils"
)

func ExampleTransferBatchApiService_GetTransferBatchByNo() {
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

	svc := partnertransferbatch.TransferBatchApiService{Client: client}
	resp, result, err := svc.GetTransferBatchByNo(ctx,
		partnertransferbatch.GetTransferBatchByNoRequest{
			BatchId:         core.String("1030000071100999991182020050700019480001"),
			NeedQueryDetail: core.Bool(true),
			Offset:          core.Int64(0),
			Limit:           core.Int64(20),
			DetailStatus:    core.String("FAIL"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetTransferBatchByNo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleTransferBatchApiService_GetTransferBatchByOutNo() {
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

	svc := partnertransferbatch.TransferBatchApiService{Client: client}
	resp, result, err := svc.GetTransferBatchByOutNo(ctx,
		partnertransferbatch.GetTransferBatchByOutNoRequest{
			OutBatchNo:      core.String("plfk2020042013"),
			NeedQueryDetail: core.Bool(true),
			Offset:          core.Int64(0),
			Limit:           core.Int64(20),
			DetailStatus:    core.String("FAIL"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetTransferBatchByOutNo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleTransferBatchApiService_InitiateTransferBatch() {
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

	svc := partnertransferbatch.TransferBatchApiService{Client: client}
	resp, result, err := svc.InitiateTransferBatch(ctx,
		partnertransferbatch.InitiateTransferBatchRequest{
			SubMchid:          core.String("1900000109"),
			SubAppid:          core.String("wxf636efh567hg4356"),
			AuthorizationType: partnertransferbatch.AUTHTYPE_INFORMATION_AUTHORIZATION_TYPE.Ptr(),
			OutBatchNo:        core.String("plfk2020042013"),
			BatchName:         core.String("2019年1月深圳分部报销单"),
			BatchRemark:       core.String("2019年1月深圳分部报销单"),
			TotalAmount:       core.Int64(4000000),
			TotalNum:          core.Int64(200),
			TransferDetailList: []partnertransferbatch.TransferDetailInput{partnertransferbatch.TransferDetailInput{
				Openid:         core.String("o-MYE42l80oelYMDE34nYD456Xoy"),
				OutDetailNo:    core.String("x23zy545Bd5436"),
				TransferAmount: core.Int64(200000),
				TransferRemark: core.String("2020年4月报销"),
				UserIdCard:     core.String("UserIdCard_example"),
				UserName:       core.String("UserName_example"),
			}},
			SpAppid:         core.String("wxf636efh567hg4388"),
			TransferPurpose: partnertransferbatch.TRANSFERUSETYPE_GOODSPAYMENT.Ptr(),
			TransferScene:   partnertransferbatch.TRANSFERSCENE_ORDINARY_TRANSFER.Ptr(),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call InitiateTransferBatch err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
