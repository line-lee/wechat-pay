// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付分停车服务
//
// 微信支付分停车服务 扣费API
//
// API version: 1.2.1

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package wexinpayscoreparking_test

import (
	"context"
	"log"
	"time"

	"github.com/line-lee/wechat-pay/core"
	"github.com/line-lee/wechat-pay/core/option"
	"github.com/line-lee/wechat-pay/services/wexinpayscoreparking"
	"github.com/line-lee/wechat-pay/utils"
)

func ExampleParkingsApiService_CreateParking() {
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

	svc := wexinpayscoreparking.ParkingsApiService{Client: client}
	resp, result, err := svc.CreateParking(ctx,
		wexinpayscoreparking.CreateParkingRequest{
			SubMchid:     core.String("1900000109"),
			OutParkingNo: core.String("1231243"),
			PlateNumber:  core.String("粤B888888"),
			PlateColor:   wexinpayscoreparking.PLATECOLOR_BLUE.Ptr(),
			NotifyUrl:    core.String("https://yoursite.com/wxpay.html"),
			StartTime:    core.Time(time.Now()),
			ParkingName:  core.String("欢乐海岸停车场"),
			FreeDuration: core.Int64(3600),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateParking err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
