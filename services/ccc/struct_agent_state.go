package ccc

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

// AgentState is a nested struct in ccc response
type AgentState struct {
	AgentId          string   `json:"AgentId" xml:"AgentId"`
	AgentName        string   `json:"AgentName" xml:"AgentName"`
	CounterParty     string   `json:"CounterParty" xml:"CounterParty"`
	Extension        string   `json:"Extension" xml:"Extension"`
	InstanceId       string   `json:"InstanceId" xml:"InstanceId"`
	State            string   `json:"State" xml:"State"`
	StateCode        string   `json:"StateCode" xml:"StateCode"`
	StateTime        int64    `json:"StateTime" xml:"StateTime"`
	SkillGroupIdList []string `json:"SkillGroupIdList" xml:"SkillGroupIdList"`
}
