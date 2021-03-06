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

// GetFileDimension invokes the qualitycheck.GetFileDimension API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/getfiledimension.html
func (client *Client) GetFileDimension(request *GetFileDimensionRequest) (response *GetFileDimensionResponse, err error) {
	response = CreateGetFileDimensionResponse()
	err = client.DoAction(request, response)
	return
}

// GetFileDimensionWithChan invokes the qualitycheck.GetFileDimension API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getfiledimension.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFileDimensionWithChan(request *GetFileDimensionRequest) (<-chan *GetFileDimensionResponse, <-chan error) {
	responseChan := make(chan *GetFileDimensionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFileDimension(request)
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

// GetFileDimensionWithCallback invokes the qualitycheck.GetFileDimension API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getfiledimension.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFileDimensionWithCallback(request *GetFileDimensionRequest, callback func(response *GetFileDimensionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFileDimensionResponse
		var err error
		defer close(result)
		response, err = client.GetFileDimension(request)
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

// GetFileDimensionRequest is the request struct for api GetFileDimension
type GetFileDimensionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetFileDimensionResponse is the response struct for api GetFileDimension
type GetFileDimensionResponse struct {
	*responses.BaseResponse
	RequestId  string                 `json:"RequestId" xml:"RequestId"`
	Success    bool                   `json:"Success" xml:"Success"`
	Code       string                 `json:"Code" xml:"Code"`
	Message    string                 `json:"Message" xml:"Message"`
	PageSize   int                    `json:"PageSize" xml:"PageSize"`
	DataSize   int                    `json:"DataSize" xml:"DataSize"`
	TotalCount int                    `json:"TotalCount" xml:"TotalCount"`
	Data       DataInGetFileDimension `json:"Data" xml:"Data"`
}

// CreateGetFileDimensionRequest creates a request to invoke GetFileDimension API
func CreateGetFileDimensionRequest() (request *GetFileDimensionRequest) {
	request = &GetFileDimensionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetFileDimension", "", "")
	return
}

// CreateGetFileDimensionResponse creates a response to parse from GetFileDimension response
func CreateGetFileDimensionResponse() (response *GetFileDimensionResponse) {
	response = &GetFileDimensionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
