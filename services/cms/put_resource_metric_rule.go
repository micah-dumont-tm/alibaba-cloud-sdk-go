package cms

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

// PutResourceMetricRule invokes the cms.PutResourceMetricRule API synchronously
func (client *Client) PutResourceMetricRule(request *PutResourceMetricRuleRequest) (response *PutResourceMetricRuleResponse, err error) {
	response = CreatePutResourceMetricRuleResponse()
	err = client.DoAction(request, response)
	return
}

// PutResourceMetricRuleWithChan invokes the cms.PutResourceMetricRule API asynchronously
func (client *Client) PutResourceMetricRuleWithChan(request *PutResourceMetricRuleRequest) (<-chan *PutResourceMetricRuleResponse, <-chan error) {
	responseChan := make(chan *PutResourceMetricRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutResourceMetricRule(request)
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

// PutResourceMetricRuleWithCallback invokes the cms.PutResourceMetricRule API asynchronously
func (client *Client) PutResourceMetricRuleWithCallback(request *PutResourceMetricRuleRequest, callback func(response *PutResourceMetricRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutResourceMetricRuleResponse
		var err error
		defer close(result)
		response, err = client.PutResourceMetricRule(request)
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

// PutResourceMetricRuleRequest is the request struct for api PutResourceMetricRule
type PutResourceMetricRuleRequest struct {
	*requests.RpcRequest
	Webhook                               string           `position:"Query" name:"Webhook"`
	EscalationsWarnComparisonOperator     string           `position:"Query" name:"Escalations.Warn.ComparisonOperator"`
	RuleName                              string           `position:"Query" name:"RuleName"`
	EffectiveInterval                     string           `position:"Query" name:"EffectiveInterval"`
	NoDataPolicy                          string           `position:"Query" name:"NoDataPolicy"`
	NoEffectiveInterval                   string           `position:"Query" name:"NoEffectiveInterval"`
	EmailSubject                          string           `position:"Query" name:"EmailSubject"`
	Options                               string           `position:"Query" name:"Options"`
	MetricName                            string           `position:"Query" name:"MetricName"`
	EscalationsWarnTimes                  requests.Integer `position:"Query" name:"Escalations.Warn.Times"`
	Period                                string           `position:"Query" name:"Period"`
	EscalationsWarnThreshold              string           `position:"Query" name:"Escalations.Warn.Threshold"`
	ContactGroups                         string           `position:"Query" name:"ContactGroups"`
	EscalationsCriticalStatistics         string           `position:"Query" name:"Escalations.Critical.Statistics"`
	GroupId                               string           `position:"Query" name:"GroupId"`
	GroupName                             string           `position:"Query" name:"GroupName"`
	Labels                                string           `position:"Query" name:"Labels"`
	Interval                              string           `position:"Query" name:"Interval"`
	RuleId                                string           `position:"Query" name:"RuleId"`
	EscalationsCriticalThreshold          string           `position:"Query" name:"Escalations.Critical.Threshold"`
	EscalationsInfoStatistics             string           `position:"Query" name:"Escalations.Info.Statistics"`
	EscalationsInfoComparisonOperator     string           `position:"Query" name:"Escalations.Info.ComparisonOperator"`
	SilenceTime                           requests.Integer `position:"Query" name:"SilenceTime"`
	Prometheus                            string           `position:"Query" name:"Prometheus"`
	CompositeExpression                   string           `position:"Query" name:"CompositeExpression"`
	Resources                             string           `position:"Query" name:"Resources"`
	EscalationsInfoTimes                  requests.Integer `position:"Query" name:"Escalations.Info.Times"`
	GroupBy                               string           `position:"Query" name:"GroupBy"`
	EscalationsCriticalTimes              requests.Integer `position:"Query" name:"Escalations.Critical.Times"`
	EscalationsWarnStatistics             string           `position:"Query" name:"Escalations.Warn.Statistics"`
	EscalationsInfoThreshold              string           `position:"Query" name:"Escalations.Info.Threshold"`
	Namespace                             string           `position:"Query" name:"Namespace"`
	EscalationsCriticalComparisonOperator string           `position:"Query" name:"Escalations.Critical.ComparisonOperator"`
}

// PutResourceMetricRuleResponse is the response struct for api PutResourceMetricRule
type PutResourceMetricRuleResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreatePutResourceMetricRuleRequest creates a request to invoke PutResourceMetricRule API
func CreatePutResourceMetricRuleRequest() (request *PutResourceMetricRuleRequest) {
	request = &PutResourceMetricRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "PutResourceMetricRule", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePutResourceMetricRuleResponse creates a response to parse from PutResourceMetricRule response
func CreatePutResourceMetricRuleResponse() (response *PutResourceMetricRuleResponse) {
	response = &PutResourceMetricRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
