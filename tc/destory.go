package tc

import (
	"fmt"
	"strings"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func (tc *TC) Destory(key, value string) {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	// 首先根据tag 查询对应的实例id
	instances := tc.Describe(key, value)
	if len(instances) == 0 {
		return
	}
	client := tc.Client
	request := cvm.NewTerminateInstancesRequest()

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request.InstanceIds = common.StringPtrs(instances)

	// 返回的resp是一个TerminateInstancesResponse的实例，与请求对象对应
	response, err := client.TerminateInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}

func (tc *TC) Describe(key, value string) (instanceids []string) {
	client := tc.Client
	request := cvm.NewDescribeInstancesRequest()

	if key != "" && value != "" {
		request.Filters = []*cvm.Filter{
			{
				Name:   common.StringPtr("tag-key"),
				Values: common.StringPtrs([]string{key}),
			},
			{
				Name:   common.StringPtr("tag-value"),
				Values: common.StringPtrs([]string{value}),
			},
		}
	}

	// 返回的resp是一个DescribeInstancesResponse的实例，与请求对象对应
	response, err := client.DescribeInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	for _, v := range response.Response.InstanceSet {
		fmt.Printf("包含实例ID: %s,公网地址是: %s\n", *v.InstanceId, strings.Join(common.StringValues(v.PublicIpAddresses), ","))
		instanceids = append(instanceids, *v.InstanceId)
	}
	if len(instanceids) == 0 {
		fmt.Println("没有查询到任何实例")
	}
	return instanceids
}
