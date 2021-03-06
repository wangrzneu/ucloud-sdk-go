// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/unet"
	"github.com/ucloud/ucloud-sdk-go/services/vpc"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/functions"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario612(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "612",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Region": "cn-bj2",
				"Zone":   "cn-bj2-02",
			}
		},
		Owners: []string{"arno.gao@ucloud.cn"},
		Title:  "新版NAT网关-natgw自动化回归-基本操作-01-BGP线路",
		Steps: []*driver.Step{
			testStep612CreateVPC01,
			testStep612CreateSubnet02,
			testStep612AllocateEIP03,
			testStep612DescribeFirewall04,
			testStep612CreateNATGW05,
			testStep612DescribeEIP06,
			testStep612DescribeNATGW07,
			testStep612SetGwDefaultExport08,
			testStep612UpdateNATGW09,
			testStep612ListSubnetForNATGW10,
			testStep612UpdateNATGWSubnet11,
			testStep612DeleteNATGW12,
			testStep612ReleaseEIP13,
			testStep612DeleteSubnet14,
			testStep612DeleteVPC15,
		},
	})
}

var testStep612CreateVPC01 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateVPCRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"Network": []interface{}{
				"172.16.0.0/12",
			},
			"Name": "vpc-natgw-bgp",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateVPC(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("VPCId", step.Must(utils.GetValue(resp, "VPCId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "创建VPC",
	FastFail:      false,
}

var testStep612CreateSubnet02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateSubnetRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VPCId":      step.Scenario.GetVar("VPCId"),
			"SubnetName": "natgw-s1-bgp",
			"Subnet":     "172.16.0.0",
			"Region":     step.Scenario.GetVar("Region"),
			"Netmask":    21,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateSubnet(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("SubnetId", step.Must(utils.GetValue(resp, "SubnetId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "创建子网",
	FastFail:      false,
}

var testStep612AllocateEIP03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewAllocateEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Tag":          "Default",
			"Region":       step.Scenario.GetVar("Region"),
			"Quantity":     1,
			"PayMode":      "Bandwidth",
			"OperatorName": "Bgp",
			"Name":         "natgw-eip-bgp",
			"ChargeType":   "Month",
			"Bandwidth":    2,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.AllocateEIP(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("EIPId", step.Must(utils.GetValue(resp, "EIPSet.0.EIPId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "申请弹性IP",
	FastFail:      false,
}

var testStep612DescribeFirewall04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeFirewallRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeFirewall(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("FWId", step.Must(utils.GetValue(resp, "DataSet")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "获取防火墙信息",
	FastFail:      false,
}

var testStep612CreateNATGW05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateNATGWRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VPCId": step.Scenario.GetVar("VPCId"),
			"Tag":   "Default",
			"SubnetworkIds": []interface{}{
				step.Scenario.GetVar("SubnetId"),
			},
			"Remark":     "bgp",
			"Region":     step.Scenario.GetVar("Region"),
			"NATGWName":  "natgw-bgp",
			"IfOpen":     0,
			"FirewallId": step.Must(functions.SearchValue(step.Scenario.GetVar("FWId"), "Type", "recommend web", "FWId")),
			"EIPIds": []interface{}{
				step.Scenario.GetVar("EIPId"),
			},
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateNATGW(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 60 * time.Second,
	Title:         "创建NatGateway",
	FastFail:      false,
}

var testStep612DescribeEIP06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"EIPIds": []interface{}{
				step.Scenario.GetVar("EIPId"),
			},
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeEIP(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("NATGWId", step.Must(utils.GetValue(resp, "EIPSet.0.Resource.ResourceID")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("EIPSet.0.Resource.ResourceType", "natgw", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取弹性IP信息",
	FastFail:      false,
}

var testStep612DescribeNATGW07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDescribeNATGWRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"Offset": 0,
			"NATGWIds": []interface{}{
				step.Scenario.GetVar("NATGWId"),
			},
			"Limit": 60,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeNATGW(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "获取NatGateway信息",
	FastFail:      false,
}

var testStep612SetGwDefaultExport08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewSetGwDefaultExportRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":      step.Scenario.GetVar("Region"),
			"NATGWId":     step.Scenario.GetVar("NATGWId"),
			"ExportEipId": step.Scenario.GetVar("EIPId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.SetGwDefaultExport(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "SetGwDefaultExportResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "设置NAT网关的默认出口",
	FastFail:      false,
}

var testStep612UpdateNATGW09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("UpdateNATGW")
		err = req.SetPayload(map[string]interface{}{
			"Tag":       "Default",
			"Remark":    "bgp-gai",
			"Region":    step.Scenario.GetVar("Region"),
			"NATGWName": "natgw-bgp-gai",
			"NATGWId":   step.Scenario.GetVar("NATGWId"),
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "更新NatGateway",
	FastFail:      false,
}

var testStep612ListSubnetForNATGW10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewListSubnetForNATGWRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VPCId":  step.Scenario.GetVar("VPCId"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.ListSubnetForNATGW(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "给出NatGateway可以绑定的子网列表",
	FastFail:      false,
}

var testStep612UpdateNATGWSubnet11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewUpdateNATGWSubnetRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"SubnetworkIds": []interface{}{
				step.Scenario.GetVar("SubnetId"),
			},
			"Region":  step.Scenario.GetVar("Region"),
			"NATGWId": step.Scenario.GetVar("NATGWId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.UpdateNATGWSubnet(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "更新NatGateway绑定子网",
	FastFail:      false,
}

var testStep612DeleteNATGW12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteNATGWRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":  step.Scenario.GetVar("Region"),
			"NATGWId": step.Scenario.GetVar("NATGWId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteNATGW(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "删除NatGateway",
	FastFail:      false,
}

var testStep612ReleaseEIP13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewReleaseEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"EIPId":  step.Scenario.GetVar("EIPId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.ReleaseEIP(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "释放弹性IP",
	FastFail:      false,
}

var testStep612DeleteSubnet14 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteSubnetRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"SubnetId": step.Scenario.GetVar("SubnetId"),
			"Region":   step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteSubnet(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "删除子网",
	FastFail:      false,
}

var testStep612DeleteVPC15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteVPCRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VPCId":  step.Scenario.GetVar("VPCId"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteVPC(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "删除VPC",
	FastFail:      false,
}
