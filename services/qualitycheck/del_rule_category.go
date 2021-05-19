package qualitycheck

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

// DelRuleCategory invokes the qualitycheck.DelRuleCategory API synchronously
func (client *Client) DelRuleCategory(request *DelRuleCategoryRequest) (response *DelRuleCategoryResponse, err error) {
	response = CreateDelRuleCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// DelRuleCategoryWithChan invokes the qualitycheck.DelRuleCategory API asynchronously
func (client *Client) DelRuleCategoryWithChan(request *DelRuleCategoryRequest) (<-chan *DelRuleCategoryResponse, <-chan error) {
	responseChan := make(chan *DelRuleCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DelRuleCategory(request)
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

// DelRuleCategoryWithCallback invokes the qualitycheck.DelRuleCategory API asynchronously
func (client *Client) DelRuleCategoryWithCallback(request *DelRuleCategoryRequest, callback func(response *DelRuleCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DelRuleCategoryResponse
		var err error
		defer close(result)
		response, err = client.DelRuleCategory(request)
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

// DelRuleCategoryRequest is the request struct for api DelRuleCategory
type DelRuleCategoryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// DelRuleCategoryResponse is the response struct for api DelRuleCategory
type DelRuleCategoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDelRuleCategoryRequest creates a request to invoke DelRuleCategory API
func CreateDelRuleCategoryRequest() (request *DelRuleCategoryRequest) {
	request = &DelRuleCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "DelRuleCategory", "", "")
	request.Method = requests.POST
	return
}

// CreateDelRuleCategoryResponse creates a response to parse from DelRuleCategory response
func CreateDelRuleCategoryResponse() (response *DelRuleCategoryResponse) {
	response = &DelRuleCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
