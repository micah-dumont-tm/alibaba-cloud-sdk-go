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

// SaveAuthenticationInfo invokes the finmall.SaveAuthenticationInfo API synchronously
// api document: https://help.aliyun.com/api/finmall/saveauthenticationinfo.html
func (client *Client) SaveAuthenticationInfo(request *SaveAuthenticationInfoRequest) (response *SaveAuthenticationInfoResponse, err error) {
	response = CreateSaveAuthenticationInfoResponse()
	err = client.DoAction(request, response)
	return
}

// SaveAuthenticationInfoWithChan invokes the finmall.SaveAuthenticationInfo API asynchronously
// api document: https://help.aliyun.com/api/finmall/saveauthenticationinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveAuthenticationInfoWithChan(request *SaveAuthenticationInfoRequest) (<-chan *SaveAuthenticationInfoResponse, <-chan error) {
	responseChan := make(chan *SaveAuthenticationInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveAuthenticationInfo(request)
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

// SaveAuthenticationInfoWithCallback invokes the finmall.SaveAuthenticationInfo API asynchronously
// api document: https://help.aliyun.com/api/finmall/saveauthenticationinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveAuthenticationInfoWithCallback(request *SaveAuthenticationInfoRequest, callback func(response *SaveAuthenticationInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveAuthenticationInfoResponse
		var err error
		defer close(result)
		response, err = client.SaveAuthenticationInfo(request)
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

// SaveAuthenticationInfoRequest is the request struct for api SaveAuthenticationInfo
type SaveAuthenticationInfoRequest struct {
	*requests.RpcRequest
	IdCardNumber         string `position:"Query" name:"IdCardNumber"`
	Address              string `position:"Query" name:"Address"`
	EmployeeEmail        string `position:"Query" name:"EmployeeEmail"`
	EmployeePhoneNumber  string `position:"Query" name:"EmployeePhoneNumber"`
	PhoneNumber          string `position:"Query" name:"PhoneNumber"`
	BusinessLicense      string `position:"Query" name:"BusinessLicense"`
	LegalPersonName      string `position:"Query" name:"LegalPersonName"`
	EnterpriseName       string `position:"Query" name:"EnterpriseName"`
	AuthenticateType     string `position:"Query" name:"AuthenticateType"`
	UserId               string `position:"Query" name:"UserId"`
	ZhimaReturnUrl       string `position:"Query" name:"ZhimaReturnUrl"`
	BankCard             string `position:"Query" name:"BankCard"`
	Email                string `position:"Query" name:"Email"`
	EmployeeName         string `position:"Query" name:"EmployeeName"`
	EmployeeIdCardNumber string `position:"Query" name:"EmployeeIdCardNumber"`
}

// SaveAuthenticationInfoResponse is the response struct for api SaveAuthenticationInfo
type SaveAuthenticationInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateSaveAuthenticationInfoRequest creates a request to invoke SaveAuthenticationInfo API
func CreateSaveAuthenticationInfoRequest() (request *SaveAuthenticationInfoRequest) {
	request = &SaveAuthenticationInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("finmall", "2018-07-23", "SaveAuthenticationInfo", "finmall", "openAPI")
	return
}

// CreateSaveAuthenticationInfoResponse creates a response to parse from SaveAuthenticationInfo response
func CreateSaveAuthenticationInfoResponse() (response *SaveAuthenticationInfoResponse) {
	response = &SaveAuthenticationInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
