//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UFS DescribeUFSVolumeMountpoint

package ufs

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUFSVolumeMountpointRequest is request schema for DescribeUFSVolumeMountpoint action
type DescribeUFSVolumeMountpointRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 文件系统ID
	VolumeId *string `required:"true"`
}

// DescribeUFSVolumeMountpointResponse is response schema for DescribeUFSVolumeMountpoint action
type DescribeUFSVolumeMountpointResponse struct {
	response.CommonBase

	//
	DataSet []MountPointInfo

	// 目前的挂载点总数
	TotalMountPointNum int

	// 文件系统能创建的最大挂载点数目
	MaxMountPointNum int
}

// NewDescribeUFSVolumeMountpointRequest will create request of DescribeUFSVolumeMountpoint action.
func (c *UFSClient) NewDescribeUFSVolumeMountpointRequest() *DescribeUFSVolumeMountpointRequest {
	req := &DescribeUFSVolumeMountpointRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUFSVolumeMountpoint - 获取文件系统挂载点信息
func (c *UFSClient) DescribeUFSVolumeMountpoint(req *DescribeUFSVolumeMountpointRequest) (*DescribeUFSVolumeMountpointResponse, error) {
	var err error
	var res DescribeUFSVolumeMountpointResponse

	err = c.Client.InvokeAction("DescribeUFSVolumeMountpoint", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
