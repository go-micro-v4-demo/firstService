package main

import (
	"go-micro.dev/v4/auth"
	"testing"
)

func TestAuth(t *testing.T) {
	//1-资源对象
	payServiceGetPayChiocesResource := &auth.Resource{
		Type:     "payService",
		Name:     "payService",
		Endpoint: "payService.getPayChioces",
	}
	payServiceCreatePayResource := &auth.Resource{
		Type:     "payService",
		Name:     "payService",
		Endpoint: "payService.createPay",
	}

	payServiceQueryOrderResource := &auth.Resource{
		Type:     "payService",
		Name:     "payService",
		Endpoint: "payService.queryOrder",
	}
	//2-资源规则
	getPayChiocesRule := &auth.Rule{
		Scope:    "getPayChioces",
		Access:   0, //规则有效
		Resource: payServiceGetPayChiocesResource,
	}
	createPayChiocesRule := &auth.Rule{
		Scope:    "createPay",
		Access:   0, //规则有效
		Resource: payServiceCreatePayResource,
	}
	queryOrderRule := &auth.Rule{
		Scope:    "queryOrder", //规则有效
		Access:   0,
		Resource: payServiceQueryOrderResource,
	}
	//3-用户权限
	orderServiceAccount := &auth.Account{
		ID:     "1",
		Scopes: []string{"getPayChioces", "createPay", "queryOrder"},
	}
	skuServiceAccount := &auth.Account{
		ID:     "1",
		Scopes: []string{"getPayChioces"},
	}
	//上面那些struct都是我们在业务数据库中查询完成后序列化出来的

	//表格驱动测试
	tableTestCases := []struct {
		Name     string
		Rules    []*auth.Rule
		Account  *auth.Account
		Resource *auth.Resource
		Error    error
	}{ //orderServiceAccount的测试用例
		{
			Name:     "orderServiceAccount能都否调用支付系统的getPayChioces",
			Resource: payServiceGetPayChiocesResource,
			Account:  orderServiceAccount,
			Rules:    []*auth.Rule{getPayChiocesRule},
			Error:    nil, //不返回错误 权限通过
		},
		{
			Name:     "orderServiceAccount能都否调用支付系统的createPay",
			Resource: payServiceCreatePayResource,
			Account:  orderServiceAccount,
			Rules:    []*auth.Rule{createPayChiocesRule},
			Error:    nil, //不返回错误 权限通过
		},
		{
			Name:     "orderServiceAccount能都否调用支付系统的queryOrder",
			Resource: payServiceQueryOrderResource,
			Account:  orderServiceAccount,
			Rules:    []*auth.Rule{queryOrderRule},
			Error:    nil, //不返回错误 权限通过
		},
		//skuServiceAccount的测试用例
		{
			Name:     "skuServiceAccount能都否调用支付系统的getPayChioces",
			Resource: payServiceGetPayChiocesResource,
			Account:  skuServiceAccount,
			Rules:    []*auth.Rule{getPayChiocesRule},
			Error:    nil, //不返回错误 权限通过
		},
		{
			Name:     "skuServiceAccount能都否调用支付系统的createPay",
			Resource: payServiceCreatePayResource,
			Account:  skuServiceAccount,
			Rules:    []*auth.Rule{createPayChiocesRule},
			Error:    auth.ErrForbidden, //返回ErrForbidden 错误
		},
	}

	for _, testCase := range tableTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if err := auth.Verify(testCase.Rules, testCase.Account, testCase.Resource); err != testCase.Error {
				t.Errorf("期望是： %v 但是得到的结果是: %v", testCase.Error, err)
			}
		})
	}
}

//https://github.com/go-micro/go-micro/blob/master/auth/rules_test.go 官方的单元测试例子
// cd 到当前目录
// go test
