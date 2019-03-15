package finmall

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

// GetDocumentDownloadUrl invokes the finmall.GetDocumentDownloadUrl API synchronously
// api document: https://help.aliyun.com/api/finmall/getdocumentdownloadurl.html
func (client *Client) GetDocumentDownloadUrl(request *GetDocumentDownloadUrlRequest) (response *GetDocumentDownloadUrlResponse, err error) {
	response = CreateGetDocumentDownloadUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetDocumentDownloadUrlWithChan invokes the finmall.GetDocumentDownloadUrl API asynchronously
// api document: https://help.aliyun.com/api/finmall/getdocumentdownloadurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDocumentDownloadUrlWithChan(request *GetDocumentDownloadUrlRequest) (<-chan *GetDocumentDownloadUrlResponse, <-chan error) {
	responseChan := make(chan *GetDocumentDownloadUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDocumentDownloadUrl(request)
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

// GetDocumentDownloadUrlWithCallback invokes the finmall.GetDocumentDownloadUrl API asynchronously
// api document: https://help.aliyun.com/api/finmall/getdocumentdownloadurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDocumentDownloadUrlWithCallback(request *GetDocumentDownloadUrlRequest, callback func(response *GetDocumentDownloadUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDocumentDownloadUrlResponse
		var err error
		defer close(result)
		response, err = client.GetDocumentDownloadUrl(request)
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

// GetDocumentDownloadUrlRequest is the request struct for api GetDocumentDownloadUrl
type GetDocumentDownloadUrlRequest struct {
	*requests.RpcRequest
	BizType    string `position:"Query" name:"BizType"`
	BizId      string `position:"Query" name:"BizId"`
	DocumentId string `position:"Query" name:"DocumentId"`
	UserId     string `position:"Query" name:"UserId"`
}

// GetDocumentDownloadUrlResponse is the response struct for api GetDocumentDownloadUrl
type GetDocumentDownloadUrlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetDocumentDownloadUrlRequest creates a request to invoke GetDocumentDownloadUrl API
func CreateGetDocumentDownloadUrlRequest() (request *GetDocumentDownloadUrlRequest) {
	request = &GetDocumentDownloadUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("finmall", "2018-07-23", "GetDocumentDownloadUrl", "finmall", "openAPI")
	return
}

// CreateGetDocumentDownloadUrlResponse creates a response to parse from GetDocumentDownloadUrl response
func CreateGetDocumentDownloadUrlResponse() (response *GetDocumentDownloadUrlResponse) {
	response = &GetDocumentDownloadUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
