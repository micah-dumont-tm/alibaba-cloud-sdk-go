package cdrs

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

// ListCorpTrackDetail invokes the cdrs.ListCorpTrackDetail API synchronously
func (client *Client) ListCorpTrackDetail(request *ListCorpTrackDetailRequest) (response *ListCorpTrackDetailResponse, err error) {
	response = CreateListCorpTrackDetailResponse()
	err = client.DoAction(request, response)
	return
}

// ListCorpTrackDetailWithChan invokes the cdrs.ListCorpTrackDetail API asynchronously
func (client *Client) ListCorpTrackDetailWithChan(request *ListCorpTrackDetailRequest) (<-chan *ListCorpTrackDetailResponse, <-chan error) {
	responseChan := make(chan *ListCorpTrackDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCorpTrackDetail(request)
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

// ListCorpTrackDetailWithCallback invokes the cdrs.ListCorpTrackDetail API asynchronously
func (client *Client) ListCorpTrackDetailWithCallback(request *ListCorpTrackDetailRequest, callback func(response *ListCorpTrackDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCorpTrackDetailResponse
		var err error
		defer close(result)
		response, err = client.ListCorpTrackDetail(request)
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

// ListCorpTrackDetailRequest is the request struct for api ListCorpTrackDetail
type ListCorpTrackDetailRequest struct {
	*requests.RpcRequest
	Schema       string `position:"Body" name:"Schema"`
	CorpId       string `position:"Body" name:"CorpId"`
	EndTime      string `position:"Body" name:"EndTime"`
	StartTime    string `position:"Body" name:"StartTime"`
	PageNumber   string `position:"Body" name:"PageNumber"`
	PageSize     string `position:"Body" name:"PageSize"`
	DataSourceId string `position:"Body" name:"DataSourceId"`
	PersonId     string `position:"Body" name:"PersonId"`
}

// ListCorpTrackDetailResponse is the response struct for api ListCorpTrackDetail
type ListCorpTrackDetailResponse struct {
	*responses.BaseResponse
	Code       string     `json:"Code" xml:"Code"`
	Message    string     `json:"Message" xml:"Message"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Success    string     `json:"Success" xml:"Success"`
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	Data       []DataItem `json:"Data" xml:"Data"`
}

// CreateListCorpTrackDetailRequest creates a request to invoke ListCorpTrackDetail API
func CreateListCorpTrackDetailRequest() (request *ListCorpTrackDetailRequest) {
	request = &ListCorpTrackDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListCorpTrackDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateListCorpTrackDetailResponse creates a response to parse from ListCorpTrackDetail response
func CreateListCorpTrackDetailResponse() (response *ListCorpTrackDetailResponse) {
	response = &ListCorpTrackDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
