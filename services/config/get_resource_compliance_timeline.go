package config

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

// GetResourceComplianceTimeline invokes the config.GetResourceComplianceTimeline API synchronously
func (client *Client) GetResourceComplianceTimeline(request *GetResourceComplianceTimelineRequest) (response *GetResourceComplianceTimelineResponse, err error) {
	response = CreateGetResourceComplianceTimelineResponse()
	err = client.DoAction(request, response)
	return
}

// GetResourceComplianceTimelineWithChan invokes the config.GetResourceComplianceTimeline API asynchronously
func (client *Client) GetResourceComplianceTimelineWithChan(request *GetResourceComplianceTimelineRequest) (<-chan *GetResourceComplianceTimelineResponse, <-chan error) {
	responseChan := make(chan *GetResourceComplianceTimelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResourceComplianceTimeline(request)
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

// GetResourceComplianceTimelineWithCallback invokes the config.GetResourceComplianceTimeline API asynchronously
func (client *Client) GetResourceComplianceTimelineWithCallback(request *GetResourceComplianceTimelineRequest, callback func(response *GetResourceComplianceTimelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResourceComplianceTimelineResponse
		var err error
		defer close(result)
		response, err = client.GetResourceComplianceTimeline(request)
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

// GetResourceComplianceTimelineRequest is the request struct for api GetResourceComplianceTimeline
type GetResourceComplianceTimelineRequest struct {
	*requests.RpcRequest
	MultiAccount requests.Boolean `position:"Query" name:"MultiAccount"`
	StartTime    requests.Integer `position:"Query" name:"StartTime"`
	NextToken    string           `position:"Query" name:"NextToken"`
	Limit        requests.Integer `position:"Query" name:"Limit"`
	ResourceId   string           `position:"Query" name:"ResourceId"`
	EndTime      requests.Integer `position:"Query" name:"EndTime"`
	ResourceType string           `position:"Query" name:"ResourceType"`
	Region       string           `position:"Query" name:"Region"`
	MemberId     string           `position:"Query" name:"MemberId"`
}

// GetResourceComplianceTimelineResponse is the response struct for api GetResourceComplianceTimeline
type GetResourceComplianceTimelineResponse struct {
	*responses.BaseResponse
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	ResourceComplianceTimeline ResourceComplianceTimeline `json:"ResourceComplianceTimeline" xml:"ResourceComplianceTimeline"`
}

// CreateGetResourceComplianceTimelineRequest creates a request to invoke GetResourceComplianceTimeline API
func CreateGetResourceComplianceTimelineRequest() (request *GetResourceComplianceTimelineRequest) {
	request = &GetResourceComplianceTimelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2019-01-08", "GetResourceComplianceTimeline", "", "")
	request.Method = requests.GET
	return
}

// CreateGetResourceComplianceTimelineResponse creates a response to parse from GetResourceComplianceTimeline response
func CreateGetResourceComplianceTimelineResponse() (response *GetResourceComplianceTimelineResponse) {
	response = &GetResourceComplianceTimelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
