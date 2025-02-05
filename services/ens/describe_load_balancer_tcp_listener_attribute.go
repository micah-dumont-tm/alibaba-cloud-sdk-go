package ens

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeLoadBalancerTCPListenerAttribute invokes the ens.DescribeLoadBalancerTCPListenerAttribute API synchronously
func (client *Client) DescribeLoadBalancerTCPListenerAttribute(request *DescribeLoadBalancerTCPListenerAttributeRequest) (response *DescribeLoadBalancerTCPListenerAttributeResponse, err error) {
	response = CreateDescribeLoadBalancerTCPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoadBalancerTCPListenerAttributeWithChan invokes the ens.DescribeLoadBalancerTCPListenerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerTCPListenerAttributeWithChan(request *DescribeLoadBalancerTCPListenerAttributeRequest) (<-chan *DescribeLoadBalancerTCPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerTCPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerTCPListenerAttribute(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeLoadBalancerTCPListenerAttributeWithCallback invokes the ens.DescribeLoadBalancerTCPListenerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerTCPListenerAttributeWithCallback(request *DescribeLoadBalancerTCPListenerAttributeRequest, callback func(response *DescribeLoadBalancerTCPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerTCPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerTCPListenerAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeLoadBalancerTCPListenerAttributeRequest is the request struct for api DescribeLoadBalancerTCPListenerAttribute
type DescribeLoadBalancerTCPListenerAttributeRequest struct {
	*requests.RpcRequest
	Protocol       string           `position:"Query" name:"Protocol"`
	ListenerPort   requests.Integer `position:"Query" name:"ListenerPort"`
	LoadBalancerId string           `position:"Query" name:"LoadBalancerId"`
}

// DescribeLoadBalancerTCPListenerAttributeResponse is the response struct for api DescribeLoadBalancerTCPListenerAttribute
type DescribeLoadBalancerTCPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId                 string `json:"RequestId" xml:"RequestId"`
	ListenerPort              int    `json:"ListenerPort" xml:"ListenerPort"`
	Status                    string `json:"Status" xml:"Status"`
	Bandwidth                 int    `json:"Bandwidth" xml:"Bandwidth"`
	Scheduler                 string `json:"Scheduler" xml:"Scheduler"`
	PersistenceTimeout        int    `json:"PersistenceTimeout" xml:"PersistenceTimeout"`
	EstablishedTimeout        int    `json:"EstablishedTimeout" xml:"EstablishedTimeout"`
	HealthCheck               string `json:"HealthCheck" xml:"HealthCheck"`
	HealthyThreshold          int    `json:"HealthyThreshold" xml:"HealthyThreshold"`
	UnhealthyThreshold        int    `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	HealthCheckConnectTimeout int    `json:"HealthCheckConnectTimeout" xml:"HealthCheckConnectTimeout"`
	HealthCheckInterval       int    `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	Description               string `json:"Description" xml:"Description"`
	HealthCheckHttpCode       string `json:"HealthCheckHttpCode" xml:"HealthCheckHttpCode"`
	HealthCheckDomain         string `json:"HealthCheckDomain" xml:"HealthCheckDomain"`
	HealthCheckURI            string `json:"HealthCheckURI" xml:"HealthCheckURI"`
	HealthCheckType           string `json:"HealthCheckType" xml:"HealthCheckType"`
	BackendServerPort         int    `json:"BackendServerPort" xml:"BackendServerPort"`
	HealthCheckConnectPort    int    `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
}

// CreateDescribeLoadBalancerTCPListenerAttributeRequest creates a request to invoke DescribeLoadBalancerTCPListenerAttribute API
func CreateDescribeLoadBalancerTCPListenerAttributeRequest() (request *DescribeLoadBalancerTCPListenerAttributeRequest) {
	request = &DescribeLoadBalancerTCPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeLoadBalancerTCPListenerAttribute", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLoadBalancerTCPListenerAttributeResponse creates a response to parse from DescribeLoadBalancerTCPListenerAttribute response
func CreateDescribeLoadBalancerTCPListenerAttributeResponse() (response *DescribeLoadBalancerTCPListenerAttributeResponse) {
	response = &DescribeLoadBalancerTCPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
