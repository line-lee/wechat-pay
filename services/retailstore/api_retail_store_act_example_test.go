// Copyright 2021 Tencent Inc. All rights reserved.
//
// 营销加价购对外API
//
// 指定服务商可通过该接口报名加价购活动、查询某个区域内的加价购活动列表、锁定加价活动购资格以及解锁加价购活动资格。
//
// API version: 1.4.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package retailstore_test

import (
	"context"
	"log"

	"github.com/line-lee/wechat-pay/core"
	"github.com/line-lee/wechat-pay/core/option"
	"github.com/line-lee/wechat-pay/services/retailstore"
	"github.com/line-lee/wechat-pay/utils"
)

func ExampleRetailStoreActApiService_AddRepresentative() {
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

	svc := retailstore.RetailStoreActApiService{Client: client}
	resp, result, err := svc.AddRepresentative(ctx,
		retailstore.AddRepresentativeRequest{
			ActivityId: core.String("3118550000000004"),
			RepresentativeInfoList: []retailstore.RepresentativeInfo{retailstore.RepresentativeInfo{
				Openid: core.String("oK7fFt8zzEZ909XH-LE2#"),
			}},
			OutRequestNo: core.String("1002600620019090123143254436"),
			AddTime:      core.String("2015-05-20T13:29:35+08:00"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call AddRepresentative err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleRetailStoreActApiService_AddStores() {
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

	svc := retailstore.RetailStoreActApiService{Client: client}
	resp, result, err := svc.AddStores(ctx,
		retailstore.AddStoresRequest{
			BrandId:      core.String("1001"),
			OutRequestNo: core.String("1002600620019090123143254436"),
			AddTime:      core.String("2015-05-20T13:29:35+08:00"),
			Stores: []retailstore.RetailStoreInfo{retailstore.RetailStoreInfo{
				StoreCode: core.String("abc_001"),
				StoreName: core.String("幸福小店"),
			}},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call AddStores err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleRetailStoreActApiService_CreateMaterials() {
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

	svc := retailstore.RetailStoreActApiService{Client: client}
	resp, result, err := svc.CreateMaterials(ctx,
		retailstore.CreateMaterialsRequest{
			BrandId:      core.String("1001"),
			OutRequestNo: core.String("1002600620019090123143254436"),
			MaterialNum:  core.Int64(100),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateMaterials err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleRetailStoreActApiService_DeleteRepresentative() {
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

	svc := retailstore.RetailStoreActApiService{Client: client}
	resp, result, err := svc.DeleteRepresentative(ctx,
		retailstore.DeleteRepresentativeRequest{
			ActivityId: core.String("3118550000000004"),
			RepresentativeInfoList: []retailstore.RepresentativeInfo{retailstore.RepresentativeInfo{
				Openid: core.String("oK7fFt8zzEZ909XH-LE2#"),
			}},
			OutRequestNo: core.String("1002600620019090123143254436"),
			DeleteTime:   core.String("2015-05-20T13:29:35.120+08:00"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call DeleteRepresentative err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleRetailStoreActApiService_DeleteStores() {
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

	svc := retailstore.RetailStoreActApiService{Client: client}
	resp, result, err := svc.DeleteStores(ctx,
		retailstore.DeleteStoresRequest{
			BrandId:      core.String("1001"),
			OutRequestNo: core.String("1002600620019090123143254436"),
			DeleteTime:   core.String("2015-05-20T13:29:35+08:00"),
			Stores: []retailstore.RetailStoreInfo{retailstore.RetailStoreInfo{
				StoreCode: core.String("abc_001"),
				StoreName: core.String("幸福小店"),
			}},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call DeleteStores err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleRetailStoreActApiService_GetStore() {
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

	svc := retailstore.RetailStoreActApiService{Client: client}
	resp, result, err := svc.GetStore(ctx,
		retailstore.GetStoreRequest{
			BrandId:   core.String("1001"),
			StoreCode: core.String("abc_001"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetStore err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleRetailStoreActApiService_ListRepresentative() {
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

	svc := retailstore.RetailStoreActApiService{Client: client}
	resp, result, err := svc.ListRepresentative(ctx,
		retailstore.ListRepresentativeRequest{
			ActivityId: core.String("3118550000000004"),
			Offset:     core.Int64(0),
			Limit:      core.Int64(10),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ListRepresentative err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleRetailStoreActApiService_ListStore() {
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

	svc := retailstore.RetailStoreActApiService{Client: client}
	resp, result, err := svc.ListStore(ctx,
		retailstore.ListStoreRequest{
			BrandId: core.String("1001"),
			Offset:  core.Int64(0),
			Limit:   core.Int64(10),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ListStore err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
