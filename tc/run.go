package tc

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func (tc *TC) Run() {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	client := tc.Client
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := cvm.NewRunInstancesRequest()

	request.InstanceChargeType = common.StringPtr("SPOTPAID")
	request.Placement = &cvm.Placement{
		Zone:      common.StringPtr("ap-hongkong-2"),
		ProjectId: common.Int64Ptr(0),
	}
	request.InstanceType = common.StringPtr("SA5.MEDIUM4")
	request.ImageId = common.StringPtr("img-487zeit5")
	request.SystemDisk = &cvm.SystemDisk{
		DiskType: common.StringPtr("CLOUD_BSSD"),
		DiskSize: common.Int64Ptr(20),
	}
	request.VirtualPrivateCloud = &cvm.VirtualPrivateCloud{
		VpcId:            common.StringPtr("vpc-dgt9qvcs"),
		SubnetId:         common.StringPtr("subnet-4npo7wfj"),
		AsVpcGateway:     common.BoolPtr(false),
		Ipv6AddressCount: common.Uint64Ptr(0),
	}
	request.InternetAccessible = &cvm.InternetAccessible{
		InternetChargeType:      common.StringPtr("TRAFFIC_POSTPAID_BY_HOUR"),
		InternetMaxBandwidthOut: common.Int64Ptr(5),
		PublicIpAssigned:        common.BoolPtr(true),
	}
	request.InstanceCount = common.Int64Ptr(1)
	request.InstanceName = common.StringPtr("ss-01")
	request.LoginSettings = &cvm.LoginSettings{
		KeyIds: common.StringPtrs([]string{"skey-kstravwz"}),
	}
	request.SecurityGroupIds = common.StringPtrs([]string{"sg-nhgo6dej"})
	request.EnhancedService = &cvm.EnhancedService{
		SecurityService: &cvm.RunSecurityServiceEnabled{
			Enabled: common.BoolPtr(true),
		},
		MonitorService: &cvm.RunMonitorServiceEnabled{
			Enabled: common.BoolPtr(true),
		},
	}
	request.TagSpecification = []*cvm.TagSpecification{
		{
			ResourceType: common.StringPtr("instance"),
			Tags: []*cvm.Tag{
				{
					Key:   common.StringPtr("role"),
					Value: common.StringPtr("server"),
				},
			},
		},
	}
	request.InstanceMarketOptions = &cvm.InstanceMarketOptionsRequest{
		SpotOptions: &cvm.SpotMarketOptions{
			MaxPrice: common.StringPtr("1"),
		},
	}
	request.UserData = common.StringPtr(execdata)

	// 返回的resp是一个RunInstancesResponse的实例，与请求对象对应
	response, err := client.RunInstances(request)
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
