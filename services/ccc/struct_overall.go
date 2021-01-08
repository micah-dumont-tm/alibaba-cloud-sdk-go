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

// Overall is a nested struct in ccc response
type Overall struct {
	AverageBreakTime             float64 `json:"AverageBreakTime" xml:"AverageBreakTime"`
	AverageReadyTime             float64 `json:"AverageReadyTime" xml:"AverageReadyTime"`
	AverageTalkTime              float64 `json:"AverageTalkTime" xml:"AverageTalkTime"`
	AverageWorkTime              float64 `json:"AverageWorkTime" xml:"AverageWorkTime"`
	MaxBreakTime                 int64   `json:"MaxBreakTime" xml:"MaxBreakTime"`
	MaxReadyTime                 int64   `json:"MaxReadyTime" xml:"MaxReadyTime"`
	MaxTalkTime                  int64   `json:"MaxTalkTime" xml:"MaxTalkTime"`
	MaxWorkTime                  int64   `json:"MaxWorkTime" xml:"MaxWorkTime"`
	OccupancyRate                float64 `json:"OccupancyRate" xml:"OccupancyRate"`
	SatisfactionIndex            float64 `json:"SatisfactionIndex" xml:"SatisfactionIndex"`
	SatisfactionSurveysOffered   int64   `json:"SatisfactionSurveysOffered" xml:"SatisfactionSurveysOffered"`
	SatisfactionSurveysResponded int64   `json:"SatisfactionSurveysResponded" xml:"SatisfactionSurveysResponded"`
	TotalBreakTime               int64   `json:"TotalBreakTime" xml:"TotalBreakTime"`
	TotalCalls                   int64   `json:"TotalCalls" xml:"TotalCalls"`
	TotalHoldTime                int64   `json:"TotalHoldTime" xml:"TotalHoldTime"`
	TotalLoggedInTime            int64   `json:"TotalLoggedInTime" xml:"TotalLoggedInTime"`
	TotalReadyTime               int64   `json:"TotalReadyTime" xml:"TotalReadyTime"`
	TotalTalkTime                int64   `json:"TotalTalkTime" xml:"TotalTalkTime"`
	TotalWorkTime                int64   `json:"TotalWorkTime" xml:"TotalWorkTime"`
}
