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

// AssignReviewer invokes the qualitycheck.AssignReviewer API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/assignreviewer.html
func (client *Client) AssignReviewer(request *AssignReviewerRequest) (response *AssignReviewerResponse, err error) {
	response = CreateAssignReviewerResponse()
	err = client.DoAction(request, response)
	return
}

// AssignReviewerWithChan invokes the qualitycheck.AssignReviewer API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/assignreviewer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssignReviewerWithChan(request *AssignReviewerRequest) (<-chan *AssignReviewerResponse, <-chan error) {
	responseChan := make(chan *AssignReviewerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssignReviewer(request)
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

// AssignReviewerWithCallback invokes the qualitycheck.AssignReviewer API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/assignreviewer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssignReviewerWithCallback(request *AssignReviewerRequest, callback func(response *AssignReviewerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssignReviewerResponse
		var err error
		defer close(result)
		response, err = client.AssignReviewer(request)
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

// AssignReviewerRequest is the request struct for api AssignReviewer
type AssignReviewerRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// AssignReviewerResponse is the response struct for api AssignReviewer
type AssignReviewerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateAssignReviewerRequest creates a request to invoke AssignReviewer API
func CreateAssignReviewerRequest() (request *AssignReviewerRequest) {
	request = &AssignReviewerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "AssignReviewer", "", "")
	return
}

// CreateAssignReviewerResponse creates a response to parse from AssignReviewer response
func CreateAssignReviewerResponse() (response *AssignReviewerResponse) {
	response = &AssignReviewerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}