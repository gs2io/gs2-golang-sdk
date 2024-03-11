/*
Copyright 2016 Game Server Services, Inc. or its affiliates. All Rights
Reserved.

Licensed under the Apache License, Version 2.0 (the "License").
You may not use this file except in compliance with the License.
A copy of the License is located at

 http://www.apache.org/licenses/LICENSE-2.0

or in the "license" file accompanying this file. This file is distributed
on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing
permissions and limitations under the License.
*/

package loginReward

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2LoginRewardRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2LoginRewardRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeNamespacesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeNamespacesResult
	if asyncResult.Err != nil {
		callback <- DescribeNamespacesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeNamespacesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeNamespacesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	path := "/"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeNamespacesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) DescribeNamespaces(
	request *DescribeNamespacesRequest,
) (*DescribeNamespacesResult, error) {
	callback := make(chan DescribeNamespacesAsyncResult, 1)
	go p.DescribeNamespacesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createNamespaceAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- CreateNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateNamespaceResult
	if asyncResult.Err != nil {
		callback <- CreateNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	path := "/"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.TransactionSetting != nil {
		bodies["transactionSetting"] = request.TransactionSetting.ToDict()
	}
	if request.ReceiveScript != nil {
		bodies["receiveScript"] = request.ReceiveScript.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) CreateNamespace(
	request *CreateNamespaceRequest,
) (*CreateNamespaceResult, error) {
	callback := make(chan CreateNamespaceAsyncResult, 1)
	go p.CreateNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getNamespaceStatusAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetNamespaceStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetNamespaceStatusResult
	if asyncResult.Err != nil {
		callback <- GetNamespaceStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetNamespaceStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetNamespaceStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	path := "/{namespaceName}/status"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getNamespaceStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) GetNamespaceStatus(
	request *GetNamespaceStatusRequest,
) (*GetNamespaceStatusResult, error) {
	callback := make(chan GetNamespaceStatusAsyncResult, 1)
	go p.GetNamespaceStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getNamespaceAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- GetNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetNamespaceResult
	if asyncResult.Err != nil {
		callback <- GetNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) GetNamespace(
	request *GetNamespaceRequest,
) (*GetNamespaceResult, error) {
	callback := make(chan GetNamespaceAsyncResult, 1)
	go p.GetNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateNamespaceAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateNamespaceResult
	if asyncResult.Err != nil {
		callback <- UpdateNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.TransactionSetting != nil {
		bodies["transactionSetting"] = request.TransactionSetting.ToDict()
	}
	if request.ReceiveScript != nil {
		bodies["receiveScript"] = request.ReceiveScript.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) UpdateNamespace(
	request *UpdateNamespaceRequest,
) (*UpdateNamespaceResult, error) {
	callback := make(chan UpdateNamespaceAsyncResult, 1)
	go p.UpdateNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteNamespaceAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteNamespaceResult
	if asyncResult.Err != nil {
		callback <- DeleteNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) DeleteNamespace(
	request *DeleteNamespaceRequest,
) (*DeleteNamespaceResult, error) {
	callback := make(chan DeleteNamespaceAsyncResult, 1)
	go p.DeleteNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func dumpUserDataByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DumpUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DumpUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- DumpUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DumpUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DumpUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) DumpUserDataByUserIdAsync(
	request *DumpUserDataByUserIdRequest,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	path := "/system/dump/user/{userId}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go dumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) DumpUserDataByUserId(
	request *DumpUserDataByUserIdRequest,
) (*DumpUserDataByUserIdResult, error) {
	callback := make(chan DumpUserDataByUserIdAsyncResult, 1)
	go p.DumpUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func checkDumpUserDataByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CheckDumpUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CheckDumpUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- CheckDumpUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckDumpUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CheckDumpUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) CheckDumpUserDataByUserIdAsync(
	request *CheckDumpUserDataByUserIdRequest,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	path := "/system/dump/user/{userId}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go checkDumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) CheckDumpUserDataByUserId(
	request *CheckDumpUserDataByUserIdRequest,
) (*CheckDumpUserDataByUserIdResult, error) {
	callback := make(chan CheckDumpUserDataByUserIdAsyncResult, 1)
	go p.CheckDumpUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func cleanUserDataByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CleanUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CleanUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- CleanUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CleanUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CleanUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) CleanUserDataByUserIdAsync(
	request *CleanUserDataByUserIdRequest,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	path := "/system/clean/user/{userId}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go cleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) CleanUserDataByUserId(
	request *CleanUserDataByUserIdRequest,
) (*CleanUserDataByUserIdResult, error) {
	callback := make(chan CleanUserDataByUserIdAsyncResult, 1)
	go p.CleanUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func checkCleanUserDataByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CheckCleanUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CheckCleanUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- CheckCleanUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckCleanUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CheckCleanUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) CheckCleanUserDataByUserIdAsync(
	request *CheckCleanUserDataByUserIdRequest,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	path := "/system/clean/user/{userId}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go checkCleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) CheckCleanUserDataByUserId(
	request *CheckCleanUserDataByUserIdRequest,
) (*CheckCleanUserDataByUserIdResult, error) {
	callback := make(chan CheckCleanUserDataByUserIdAsyncResult, 1)
	go p.CheckCleanUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareImportUserDataByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareImportUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareImportUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- PrepareImportUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareImportUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareImportUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) PrepareImportUserDataByUserIdAsync(
	request *PrepareImportUserDataByUserIdRequest,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	path := "/system/import/user/{userId}/prepare"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go prepareImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) PrepareImportUserDataByUserId(
	request *PrepareImportUserDataByUserIdRequest,
) (*PrepareImportUserDataByUserIdResult, error) {
	callback := make(chan PrepareImportUserDataByUserIdAsyncResult, 1)
	go p.PrepareImportUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func importUserDataByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ImportUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ImportUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- ImportUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ImportUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ImportUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) ImportUserDataByUserIdAsync(
	request *ImportUserDataByUserIdRequest,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	path := "/system/import/user/{userId}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go importUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) ImportUserDataByUserId(
	request *ImportUserDataByUserIdRequest,
) (*ImportUserDataByUserIdResult, error) {
	callback := make(chan ImportUserDataByUserIdAsyncResult, 1)
	go p.ImportUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func checkImportUserDataByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CheckImportUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CheckImportUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- CheckImportUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckImportUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CheckImportUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) CheckImportUserDataByUserIdAsync(
	request *CheckImportUserDataByUserIdRequest,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	path := "/system/import/user/{userId}/{uploadToken}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		path = strings.ReplaceAll(path, "{uploadToken}", core.ToString(*request.UploadToken))
	} else {
		path = strings.ReplaceAll(path, "{uploadToken}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go checkImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) CheckImportUserDataByUserId(
	request *CheckImportUserDataByUserIdRequest,
) (*CheckImportUserDataByUserIdResult, error) {
	callback := make(chan CheckImportUserDataByUserIdAsyncResult, 1)
	go p.CheckImportUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBonusModelMastersAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBonusModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBonusModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBonusModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeBonusModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBonusModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBonusModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) DescribeBonusModelMastersAsync(
	request *DescribeBonusModelMastersRequest,
	callback chan<- DescribeBonusModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/bonusModel"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeBonusModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) DescribeBonusModelMasters(
	request *DescribeBonusModelMastersRequest,
) (*DescribeBonusModelMastersResult, error) {
	callback := make(chan DescribeBonusModelMastersAsyncResult, 1)
	go p.DescribeBonusModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createBonusModelMasterAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- CreateBonusModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateBonusModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateBonusModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateBonusModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateBonusModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateBonusModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) CreateBonusModelMasterAsync(
	request *CreateBonusModelMasterRequest,
	callback chan<- CreateBonusModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/bonusModel"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.PeriodEventId != nil && *request.PeriodEventId != "" {
		bodies["periodEventId"] = *request.PeriodEventId
	}
	if request.ResetHour != nil {
		bodies["resetHour"] = *request.ResetHour
	}
	if request.Repeat != nil && *request.Repeat != "" {
		bodies["repeat"] = *request.Repeat
	}
	if request.Rewards != nil {
		var _rewards []interface{}
		for _, item := range request.Rewards {
			_rewards = append(_rewards, item)
		}
		bodies["rewards"] = _rewards
	}
	if request.MissedReceiveRelief != nil && *request.MissedReceiveRelief != "" {
		bodies["missedReceiveRelief"] = *request.MissedReceiveRelief
	}
	if request.MissedReceiveReliefConsumeActions != nil {
		var _missedReceiveReliefConsumeActions []interface{}
		for _, item := range request.MissedReceiveReliefConsumeActions {
			_missedReceiveReliefConsumeActions = append(_missedReceiveReliefConsumeActions, item)
		}
		bodies["missedReceiveReliefConsumeActions"] = _missedReceiveReliefConsumeActions
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createBonusModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) CreateBonusModelMaster(
	request *CreateBonusModelMasterRequest,
) (*CreateBonusModelMasterResult, error) {
	callback := make(chan CreateBonusModelMasterAsyncResult, 1)
	go p.CreateBonusModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBonusModelMasterAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- GetBonusModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBonusModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBonusModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetBonusModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBonusModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBonusModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) GetBonusModelMasterAsync(
	request *GetBonusModelMasterRequest,
	callback chan<- GetBonusModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/bonusModel/{bonusModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getBonusModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) GetBonusModelMaster(
	request *GetBonusModelMasterRequest,
) (*GetBonusModelMasterResult, error) {
	callback := make(chan GetBonusModelMasterAsyncResult, 1)
	go p.GetBonusModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateBonusModelMasterAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateBonusModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateBonusModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateBonusModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateBonusModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateBonusModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateBonusModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) UpdateBonusModelMasterAsync(
	request *UpdateBonusModelMasterRequest,
	callback chan<- UpdateBonusModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/bonusModel/{bonusModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.PeriodEventId != nil && *request.PeriodEventId != "" {
		bodies["periodEventId"] = *request.PeriodEventId
	}
	if request.ResetHour != nil {
		bodies["resetHour"] = *request.ResetHour
	}
	if request.Repeat != nil && *request.Repeat != "" {
		bodies["repeat"] = *request.Repeat
	}
	if request.Rewards != nil {
		var _rewards []interface{}
		for _, item := range request.Rewards {
			_rewards = append(_rewards, item)
		}
		bodies["rewards"] = _rewards
	}
	if request.MissedReceiveRelief != nil && *request.MissedReceiveRelief != "" {
		bodies["missedReceiveRelief"] = *request.MissedReceiveRelief
	}
	if request.MissedReceiveReliefConsumeActions != nil {
		var _missedReceiveReliefConsumeActions []interface{}
		for _, item := range request.MissedReceiveReliefConsumeActions {
			_missedReceiveReliefConsumeActions = append(_missedReceiveReliefConsumeActions, item)
		}
		bodies["missedReceiveReliefConsumeActions"] = _missedReceiveReliefConsumeActions
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateBonusModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) UpdateBonusModelMaster(
	request *UpdateBonusModelMasterRequest,
) (*UpdateBonusModelMasterResult, error) {
	callback := make(chan UpdateBonusModelMasterAsyncResult, 1)
	go p.UpdateBonusModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteBonusModelMasterAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteBonusModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteBonusModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteBonusModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteBonusModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteBonusModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteBonusModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) DeleteBonusModelMasterAsync(
	request *DeleteBonusModelMasterRequest,
	callback chan<- DeleteBonusModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/bonusModel/{bonusModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteBonusModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) DeleteBonusModelMaster(
	request *DeleteBonusModelMasterRequest,
) (*DeleteBonusModelMasterResult, error) {
	callback := make(chan DeleteBonusModelMasterAsyncResult, 1)
	go p.DeleteBonusModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- ExportMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ExportMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ExportMasterResult
	if asyncResult.Err != nil {
		callback <- ExportMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ExportMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ExportMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	path := "/{namespaceName}/master/export"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go exportMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) ExportMaster(
	request *ExportMasterRequest,
) (*ExportMasterResult, error) {
	callback := make(chan ExportMasterAsyncResult, 1)
	go p.ExportMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getCurrentBonusMasterAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentBonusMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentBonusMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentBonusMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentBonusMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentBonusMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentBonusMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) GetCurrentBonusMasterAsync(
	request *GetCurrentBonusMasterRequest,
	callback chan<- GetCurrentBonusMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getCurrentBonusMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) GetCurrentBonusMaster(
	request *GetCurrentBonusMasterRequest,
) (*GetCurrentBonusMasterResult, error) {
	callback := make(chan GetCurrentBonusMasterAsyncResult, 1)
	go p.GetCurrentBonusMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentBonusMasterAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentBonusMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentBonusMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentBonusMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentBonusMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentBonusMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentBonusMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) UpdateCurrentBonusMasterAsync(
	request *UpdateCurrentBonusMasterRequest,
	callback chan<- UpdateCurrentBonusMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Settings != nil && *request.Settings != "" {
		bodies["settings"] = *request.Settings
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateCurrentBonusMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) UpdateCurrentBonusMaster(
	request *UpdateCurrentBonusMasterRequest,
) (*UpdateCurrentBonusMasterResult, error) {
	callback := make(chan UpdateCurrentBonusMasterAsyncResult, 1)
	go p.UpdateCurrentBonusMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentBonusMasterFromGitHubAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentBonusMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentBonusMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentBonusMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentBonusMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentBonusMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentBonusMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) UpdateCurrentBonusMasterFromGitHubAsync(
	request *UpdateCurrentBonusMasterFromGitHubRequest,
	callback chan<- UpdateCurrentBonusMasterFromGitHubAsyncResult,
) {
	path := "/{namespaceName}/master/from_git_hub"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.CheckoutSetting != nil {
		bodies["checkoutSetting"] = request.CheckoutSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateCurrentBonusMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) UpdateCurrentBonusMasterFromGitHub(
	request *UpdateCurrentBonusMasterFromGitHubRequest,
) (*UpdateCurrentBonusMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentBonusMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentBonusMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBonusModelsAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBonusModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBonusModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBonusModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeBonusModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBonusModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBonusModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) DescribeBonusModelsAsync(
	request *DescribeBonusModelsRequest,
	callback chan<- DescribeBonusModelsAsyncResult,
) {
	path := "/{namespaceName}/model/bonusModel"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeBonusModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) DescribeBonusModels(
	request *DescribeBonusModelsRequest,
) (*DescribeBonusModelsResult, error) {
	callback := make(chan DescribeBonusModelsAsyncResult, 1)
	go p.DescribeBonusModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBonusModelAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- GetBonusModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBonusModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBonusModelResult
	if asyncResult.Err != nil {
		callback <- GetBonusModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBonusModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBonusModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) GetBonusModelAsync(
	request *GetBonusModelRequest,
	callback chan<- GetBonusModelAsyncResult,
) {
	path := "/{namespaceName}/model/bonusModel/{bonusModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getBonusModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) GetBonusModel(
	request *GetBonusModelRequest,
) (*GetBonusModelResult, error) {
	callback := make(chan GetBonusModelAsyncResult, 1)
	go p.GetBonusModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func receiveAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- ReceiveAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "loginReward.bonus.alreadyReceived" {
				asyncResult.Err = gs2err.SetClientError(AlreadyReceived{})
			}
		}
		callback <- ReceiveAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReceiveAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) ReceiveAsync(
	request *ReceiveRequest,
	callback chan<- ReceiveAsyncResult,
) {
	path := "/{namespaceName}/user/me/bonus/{bonusModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go receiveAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) Receive(
	request *ReceiveRequest,
) (*ReceiveResult, error) {
	callback := make(chan ReceiveAsyncResult, 1)
	go p.ReceiveAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func receiveByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- ReceiveByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "loginReward.bonus.alreadyReceived" {
				asyncResult.Err = gs2err.SetClientError(AlreadyReceived{})
			}
		}
		callback <- ReceiveByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReceiveByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) ReceiveByUserIdAsync(
	request *ReceiveByUserIdRequest,
	callback chan<- ReceiveByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/bonus/{bonusModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go receiveByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) ReceiveByUserId(
	request *ReceiveByUserIdRequest,
) (*ReceiveByUserIdResult, error) {
	callback := make(chan ReceiveByUserIdAsyncResult, 1)
	go p.ReceiveByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func missedReceiveAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- MissedReceiveAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MissedReceiveAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MissedReceiveResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "loginReward.bonus.alreadyReceived" {
				asyncResult.Err = gs2err.SetClientError(AlreadyReceived{})
			}
		}
		callback <- MissedReceiveAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MissedReceiveAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- MissedReceiveAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) MissedReceiveAsync(
	request *MissedReceiveRequest,
	callback chan<- MissedReceiveAsyncResult,
) {
	path := "/{namespaceName}/user/me/bonus/{bonusModelName}/missed"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StepNumber != nil {
		bodies["stepNumber"] = *request.StepNumber
	}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go missedReceiveAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) MissedReceive(
	request *MissedReceiveRequest,
) (*MissedReceiveResult, error) {
	callback := make(chan MissedReceiveAsyncResult, 1)
	go p.MissedReceiveAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func missedReceiveByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- MissedReceiveByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MissedReceiveByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MissedReceiveByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "loginReward.bonus.alreadyReceived" {
				asyncResult.Err = gs2err.SetClientError(AlreadyReceived{})
			}
		}
		callback <- MissedReceiveByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MissedReceiveByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- MissedReceiveByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) MissedReceiveByUserIdAsync(
	request *MissedReceiveByUserIdRequest,
	callback chan<- MissedReceiveByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/bonus/{bonusModelName}/missed"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StepNumber != nil {
		bodies["stepNumber"] = *request.StepNumber
	}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go missedReceiveByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) MissedReceiveByUserId(
	request *MissedReceiveByUserIdRequest,
) (*MissedReceiveByUserIdResult, error) {
	callback := make(chan MissedReceiveByUserIdAsyncResult, 1)
	go p.MissedReceiveByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeReceiveStatusesAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeReceiveStatusesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReceiveStatusesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReceiveStatusesResult
	if asyncResult.Err != nil {
		callback <- DescribeReceiveStatusesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeReceiveStatusesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeReceiveStatusesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) DescribeReceiveStatusesAsync(
	request *DescribeReceiveStatusesRequest,
	callback chan<- DescribeReceiveStatusesAsyncResult,
) {
	path := "/{namespaceName}/user/me/login_reward"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeReceiveStatusesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) DescribeReceiveStatuses(
	request *DescribeReceiveStatusesRequest,
) (*DescribeReceiveStatusesResult, error) {
	callback := make(chan DescribeReceiveStatusesAsyncResult, 1)
	go p.DescribeReceiveStatusesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeReceiveStatusesByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeReceiveStatusesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReceiveStatusesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReceiveStatusesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeReceiveStatusesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeReceiveStatusesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeReceiveStatusesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) DescribeReceiveStatusesByUserIdAsync(
	request *DescribeReceiveStatusesByUserIdRequest,
	callback chan<- DescribeReceiveStatusesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/login_reward"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeReceiveStatusesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) DescribeReceiveStatusesByUserId(
	request *DescribeReceiveStatusesByUserIdRequest,
) (*DescribeReceiveStatusesByUserIdResult, error) {
	callback := make(chan DescribeReceiveStatusesByUserIdAsyncResult, 1)
	go p.DescribeReceiveStatusesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getReceiveStatusAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- GetReceiveStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReceiveStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReceiveStatusResult
	if asyncResult.Err != nil {
		callback <- GetReceiveStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetReceiveStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetReceiveStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) GetReceiveStatusAsync(
	request *GetReceiveStatusRequest,
	callback chan<- GetReceiveStatusAsyncResult,
) {
	path := "/{namespaceName}/user/me/receiveStatus/{bonusModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getReceiveStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) GetReceiveStatus(
	request *GetReceiveStatusRequest,
) (*GetReceiveStatusResult, error) {
	callback := make(chan GetReceiveStatusAsyncResult, 1)
	go p.GetReceiveStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getReceiveStatusByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- GetReceiveStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReceiveStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReceiveStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetReceiveStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetReceiveStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetReceiveStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) GetReceiveStatusByUserIdAsync(
	request *GetReceiveStatusByUserIdRequest,
	callback chan<- GetReceiveStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/receiveStatus/{bonusModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go getReceiveStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) GetReceiveStatusByUserId(
	request *GetReceiveStatusByUserIdRequest,
) (*GetReceiveStatusByUserIdResult, error) {
	callback := make(chan GetReceiveStatusByUserIdAsyncResult, 1)
	go p.GetReceiveStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteReceiveStatusByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteReceiveStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteReceiveStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteReceiveStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteReceiveStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteReceiveStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteReceiveStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) DeleteReceiveStatusByUserIdAsync(
	request *DeleteReceiveStatusByUserIdRequest,
	callback chan<- DeleteReceiveStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/receiveStatus/{bonusModelName}/delete"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go deleteReceiveStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) DeleteReceiveStatusByUserId(
	request *DeleteReceiveStatusByUserIdRequest,
) (*DeleteReceiveStatusByUserIdResult, error) {
	callback := make(chan DeleteReceiveStatusByUserIdAsyncResult, 1)
	go p.DeleteReceiveStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteReceiveStatusByStampSheetAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteReceiveStatusByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteReceiveStatusByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteReceiveStatusByStampSheetResult
	if asyncResult.Err != nil {
		callback <- DeleteReceiveStatusByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteReceiveStatusByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteReceiveStatusByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) DeleteReceiveStatusByStampSheetAsync(
	request *DeleteReceiveStatusByStampSheetRequest,
	callback chan<- DeleteReceiveStatusByStampSheetAsyncResult,
) {
	path := "/receiveStatus/delete"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampSheet != nil && *request.StampSheet != "" {
		bodies["stampSheet"] = *request.StampSheet
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteReceiveStatusByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) DeleteReceiveStatusByStampSheet(
	request *DeleteReceiveStatusByStampSheetRequest,
) (*DeleteReceiveStatusByStampSheetResult, error) {
	callback := make(chan DeleteReceiveStatusByStampSheetAsyncResult, 1)
	go p.DeleteReceiveStatusByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func markReceivedAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- MarkReceivedAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MarkReceivedAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MarkReceivedResult
	if asyncResult.Err != nil {
		callback <- MarkReceivedAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MarkReceivedAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- MarkReceivedAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) MarkReceivedAsync(
	request *MarkReceivedRequest,
	callback chan<- MarkReceivedAsyncResult,
) {
	path := "/{namespaceName}/user/me/receiveStatus/{bonusModelName}/mark"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StepNumber != nil {
		bodies["stepNumber"] = *request.StepNumber
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go markReceivedAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) MarkReceived(
	request *MarkReceivedRequest,
) (*MarkReceivedResult, error) {
	callback := make(chan MarkReceivedAsyncResult, 1)
	go p.MarkReceivedAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func markReceivedByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- MarkReceivedByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MarkReceivedByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MarkReceivedByUserIdResult
	if asyncResult.Err != nil {
		callback <- MarkReceivedByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MarkReceivedByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- MarkReceivedByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) MarkReceivedByUserIdAsync(
	request *MarkReceivedByUserIdRequest,
	callback chan<- MarkReceivedByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/receiveStatus/{bonusModelName}/mark"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StepNumber != nil {
		bodies["stepNumber"] = *request.StepNumber
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go markReceivedByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) MarkReceivedByUserId(
	request *MarkReceivedByUserIdRequest,
) (*MarkReceivedByUserIdResult, error) {
	callback := make(chan MarkReceivedByUserIdAsyncResult, 1)
	go p.MarkReceivedByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func unmarkReceivedByUserIdAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- UnmarkReceivedByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UnmarkReceivedByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UnmarkReceivedByUserIdResult
	if asyncResult.Err != nil {
		callback <- UnmarkReceivedByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UnmarkReceivedByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UnmarkReceivedByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) UnmarkReceivedByUserIdAsync(
	request *UnmarkReceivedByUserIdRequest,
	callback chan<- UnmarkReceivedByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/receiveStatus/{bonusModelName}/unmark"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.BonusModelName != nil && *request.BonusModelName != "" {
		path = strings.ReplaceAll(path, "{bonusModelName}", core.ToString(*request.BonusModelName))
	} else {
		path = strings.ReplaceAll(path, "{bonusModelName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StepNumber != nil {
		bodies["stepNumber"] = *request.StepNumber
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go unmarkReceivedByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) UnmarkReceivedByUserId(
	request *UnmarkReceivedByUserIdRequest,
) (*UnmarkReceivedByUserIdResult, error) {
	callback := make(chan UnmarkReceivedByUserIdAsyncResult, 1)
	go p.UnmarkReceivedByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func markReceivedByStampTaskAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- MarkReceivedByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MarkReceivedByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MarkReceivedByStampTaskResult
	if asyncResult.Err != nil {
		callback <- MarkReceivedByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MarkReceivedByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- MarkReceivedByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) MarkReceivedByStampTaskAsync(
	request *MarkReceivedByStampTaskRequest,
	callback chan<- MarkReceivedByStampTaskAsyncResult,
) {
	path := "/receiveStatus/mark"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampTask != nil && *request.StampTask != "" {
		bodies["stampTask"] = *request.StampTask
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go markReceivedByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) MarkReceivedByStampTask(
	request *MarkReceivedByStampTaskRequest,
) (*MarkReceivedByStampTaskResult, error) {
	callback := make(chan MarkReceivedByStampTaskAsyncResult, 1)
	go p.MarkReceivedByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func unmarkReceivedByStampSheetAsyncHandler(
	client Gs2LoginRewardRestClient,
	job *core.NetworkJob,
	callback chan<- UnmarkReceivedByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UnmarkReceivedByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UnmarkReceivedByStampSheetResult
	if asyncResult.Err != nil {
		callback <- UnmarkReceivedByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UnmarkReceivedByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UnmarkReceivedByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LoginRewardRestClient) UnmarkReceivedByStampSheetAsync(
	request *UnmarkReceivedByStampSheetRequest,
	callback chan<- UnmarkReceivedByStampSheetAsyncResult,
) {
	path := "/receiveStatus/unmark"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampSheet != nil && *request.StampSheet != "" {
		bodies["stampSheet"] = *request.StampSheet
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go unmarkReceivedByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("login-reward").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LoginRewardRestClient) UnmarkReceivedByStampSheet(
	request *UnmarkReceivedByStampSheetRequest,
) (*UnmarkReceivedByStampSheetResult, error) {
	callback := make(chan UnmarkReceivedByStampSheetAsyncResult, 1)
	go p.UnmarkReceivedByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
