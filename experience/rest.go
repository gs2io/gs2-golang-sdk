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

package experience

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2ExperienceRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2ExperienceRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2ExperienceRestClient,
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

func (p Gs2ExperienceRestClient) DescribeNamespacesAsync(
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeNamespacesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) DescribeNamespaces(
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
	client Gs2ExperienceRestClient,
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

func (p Gs2ExperienceRestClient) CreateNamespaceAsync(
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
	if request.ExperienceCapScriptId != nil && *request.ExperienceCapScriptId != "" {
		bodies["experienceCapScriptId"] = *request.ExperienceCapScriptId
	}
	if request.ChangeExperienceScript != nil {
		bodies["changeExperienceScript"] = request.ChangeExperienceScript.ToDict()
	}
	if request.ChangeRankScript != nil {
		bodies["changeRankScript"] = request.ChangeRankScript.ToDict()
	}
	if request.ChangeRankCapScript != nil {
		bodies["changeRankCapScript"] = request.ChangeRankCapScript.ToDict()
	}
	if request.OverflowExperienceScript != nil {
		bodies["overflowExperienceScript"] = request.OverflowExperienceScript.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) CreateNamespace(
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
	client Gs2ExperienceRestClient,
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

func (p Gs2ExperienceRestClient) GetNamespaceStatusAsync(
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getNamespaceStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) GetNamespaceStatus(
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
	client Gs2ExperienceRestClient,
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

func (p Gs2ExperienceRestClient) GetNamespaceAsync(
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) GetNamespace(
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
	client Gs2ExperienceRestClient,
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

func (p Gs2ExperienceRestClient) UpdateNamespaceAsync(
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
	if request.ExperienceCapScriptId != nil && *request.ExperienceCapScriptId != "" {
		bodies["experienceCapScriptId"] = *request.ExperienceCapScriptId
	}
	if request.ChangeExperienceScript != nil {
		bodies["changeExperienceScript"] = request.ChangeExperienceScript.ToDict()
	}
	if request.ChangeRankScript != nil {
		bodies["changeRankScript"] = request.ChangeRankScript.ToDict()
	}
	if request.ChangeRankCapScript != nil {
		bodies["changeRankCapScript"] = request.ChangeRankCapScript.ToDict()
	}
	if request.OverflowExperienceScript != nil {
		bodies["overflowExperienceScript"] = request.OverflowExperienceScript.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) UpdateNamespace(
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
	client Gs2ExperienceRestClient,
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

func (p Gs2ExperienceRestClient) DeleteNamespaceAsync(
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) DeleteNamespace(
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
	client Gs2ExperienceRestClient,
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

func (p Gs2ExperienceRestClient) DumpUserDataByUserIdAsync(
	request *DumpUserDataByUserIdRequest,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	path := "/system/user/{userId}/dump"
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go dumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) DumpUserDataByUserId(
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
	client Gs2ExperienceRestClient,
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

func (p Gs2ExperienceRestClient) CheckDumpUserDataByUserIdAsync(
	request *CheckDumpUserDataByUserIdRequest,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	path := "/system/user/{userId}/dump"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go checkDumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) CheckDumpUserDataByUserId(
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
	client Gs2ExperienceRestClient,
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

func (p Gs2ExperienceRestClient) CleanUserDataByUserIdAsync(
	request *CleanUserDataByUserIdRequest,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	path := "/system/user/{userId}/clean"
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go cleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) CleanUserDataByUserId(
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
	client Gs2ExperienceRestClient,
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

func (p Gs2ExperienceRestClient) CheckCleanUserDataByUserIdAsync(
	request *CheckCleanUserDataByUserIdRequest,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	path := "/system/user/{userId}/clean"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go checkCleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) CheckCleanUserDataByUserId(
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

func describeExperienceModelMastersAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeExperienceModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeExperienceModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeExperienceModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeExperienceModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeExperienceModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeExperienceModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) DescribeExperienceModelMastersAsync(
	request *DescribeExperienceModelMastersRequest,
	callback chan<- DescribeExperienceModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/model"
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeExperienceModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) DescribeExperienceModelMasters(
	request *DescribeExperienceModelMastersRequest,
) (*DescribeExperienceModelMastersResult, error) {
	callback := make(chan DescribeExperienceModelMastersAsyncResult, 1)
	go p.DescribeExperienceModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createExperienceModelMasterAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- CreateExperienceModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateExperienceModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateExperienceModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateExperienceModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateExperienceModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateExperienceModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) CreateExperienceModelMasterAsync(
	request *CreateExperienceModelMasterRequest,
	callback chan<- CreateExperienceModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model"
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
	if request.DefaultExperience != nil {
		bodies["defaultExperience"] = *request.DefaultExperience
	}
	if request.DefaultRankCap != nil {
		bodies["defaultRankCap"] = *request.DefaultRankCap
	}
	if request.MaxRankCap != nil {
		bodies["maxRankCap"] = *request.MaxRankCap
	}
	if request.RankThresholdName != nil && *request.RankThresholdName != "" {
		bodies["rankThresholdName"] = *request.RankThresholdName
	}
	if request.AcquireActionRates != nil {
		var _acquireActionRates []interface{}
		for _, item := range request.AcquireActionRates {
			_acquireActionRates = append(_acquireActionRates, item)
		}
		bodies["acquireActionRates"] = _acquireActionRates
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createExperienceModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) CreateExperienceModelMaster(
	request *CreateExperienceModelMasterRequest,
) (*CreateExperienceModelMasterResult, error) {
	callback := make(chan CreateExperienceModelMasterAsyncResult, 1)
	go p.CreateExperienceModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getExperienceModelMasterAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- GetExperienceModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetExperienceModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetExperienceModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetExperienceModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetExperienceModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetExperienceModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) GetExperienceModelMasterAsync(
	request *GetExperienceModelMasterRequest,
	callback chan<- GetExperienceModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{experienceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getExperienceModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) GetExperienceModelMaster(
	request *GetExperienceModelMasterRequest,
) (*GetExperienceModelMasterResult, error) {
	callback := make(chan GetExperienceModelMasterAsyncResult, 1)
	go p.GetExperienceModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateExperienceModelMasterAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateExperienceModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateExperienceModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateExperienceModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateExperienceModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateExperienceModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateExperienceModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) UpdateExperienceModelMasterAsync(
	request *UpdateExperienceModelMasterRequest,
	callback chan<- UpdateExperienceModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{experienceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.DefaultExperience != nil {
		bodies["defaultExperience"] = *request.DefaultExperience
	}
	if request.DefaultRankCap != nil {
		bodies["defaultRankCap"] = *request.DefaultRankCap
	}
	if request.MaxRankCap != nil {
		bodies["maxRankCap"] = *request.MaxRankCap
	}
	if request.RankThresholdName != nil && *request.RankThresholdName != "" {
		bodies["rankThresholdName"] = *request.RankThresholdName
	}
	if request.AcquireActionRates != nil {
		var _acquireActionRates []interface{}
		for _, item := range request.AcquireActionRates {
			_acquireActionRates = append(_acquireActionRates, item)
		}
		bodies["acquireActionRates"] = _acquireActionRates
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateExperienceModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) UpdateExperienceModelMaster(
	request *UpdateExperienceModelMasterRequest,
) (*UpdateExperienceModelMasterResult, error) {
	callback := make(chan UpdateExperienceModelMasterAsyncResult, 1)
	go p.UpdateExperienceModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteExperienceModelMasterAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteExperienceModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteExperienceModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteExperienceModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteExperienceModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteExperienceModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteExperienceModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) DeleteExperienceModelMasterAsync(
	request *DeleteExperienceModelMasterRequest,
	callback chan<- DeleteExperienceModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{experienceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteExperienceModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) DeleteExperienceModelMaster(
	request *DeleteExperienceModelMasterRequest,
) (*DeleteExperienceModelMasterResult, error) {
	callback := make(chan DeleteExperienceModelMasterAsyncResult, 1)
	go p.DeleteExperienceModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeExperienceModelsAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeExperienceModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeExperienceModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeExperienceModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeExperienceModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeExperienceModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeExperienceModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) DescribeExperienceModelsAsync(
	request *DescribeExperienceModelsRequest,
	callback chan<- DescribeExperienceModelsAsyncResult,
) {
	path := "/{namespaceName}/model"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeExperienceModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) DescribeExperienceModels(
	request *DescribeExperienceModelsRequest,
) (*DescribeExperienceModelsResult, error) {
	callback := make(chan DescribeExperienceModelsAsyncResult, 1)
	go p.DescribeExperienceModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getExperienceModelAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- GetExperienceModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetExperienceModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetExperienceModelResult
	if asyncResult.Err != nil {
		callback <- GetExperienceModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetExperienceModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetExperienceModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) GetExperienceModelAsync(
	request *GetExperienceModelRequest,
	callback chan<- GetExperienceModelAsyncResult,
) {
	path := "/{namespaceName}/model/{experienceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getExperienceModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) GetExperienceModel(
	request *GetExperienceModelRequest,
) (*GetExperienceModelResult, error) {
	callback := make(chan GetExperienceModelAsyncResult, 1)
	go p.GetExperienceModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeThresholdMastersAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeThresholdMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeThresholdMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeThresholdMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeThresholdMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeThresholdMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeThresholdMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) DescribeThresholdMastersAsync(
	request *DescribeThresholdMastersRequest,
	callback chan<- DescribeThresholdMastersAsyncResult,
) {
	path := "/{namespaceName}/master/threshold"
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeThresholdMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) DescribeThresholdMasters(
	request *DescribeThresholdMastersRequest,
) (*DescribeThresholdMastersResult, error) {
	callback := make(chan DescribeThresholdMastersAsyncResult, 1)
	go p.DescribeThresholdMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createThresholdMasterAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- CreateThresholdMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateThresholdMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateThresholdMasterResult
	if asyncResult.Err != nil {
		callback <- CreateThresholdMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateThresholdMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateThresholdMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) CreateThresholdMasterAsync(
	request *CreateThresholdMasterRequest,
	callback chan<- CreateThresholdMasterAsyncResult,
) {
	path := "/{namespaceName}/master/threshold"
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
	if request.Values != nil {
		var _values []interface{}
		for _, item := range request.Values {
			_values = append(_values, item)
		}
		bodies["values"] = _values
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createThresholdMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) CreateThresholdMaster(
	request *CreateThresholdMasterRequest,
) (*CreateThresholdMasterResult, error) {
	callback := make(chan CreateThresholdMasterAsyncResult, 1)
	go p.CreateThresholdMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getThresholdMasterAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- GetThresholdMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetThresholdMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetThresholdMasterResult
	if asyncResult.Err != nil {
		callback <- GetThresholdMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetThresholdMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetThresholdMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) GetThresholdMasterAsync(
	request *GetThresholdMasterRequest,
	callback chan<- GetThresholdMasterAsyncResult,
) {
	path := "/{namespaceName}/master/threshold/{thresholdName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ThresholdName != nil && *request.ThresholdName != "" {
		path = strings.ReplaceAll(path, "{thresholdName}", core.ToString(*request.ThresholdName))
	} else {
		path = strings.ReplaceAll(path, "{thresholdName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getThresholdMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) GetThresholdMaster(
	request *GetThresholdMasterRequest,
) (*GetThresholdMasterResult, error) {
	callback := make(chan GetThresholdMasterAsyncResult, 1)
	go p.GetThresholdMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateThresholdMasterAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateThresholdMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateThresholdMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateThresholdMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateThresholdMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateThresholdMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateThresholdMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) UpdateThresholdMasterAsync(
	request *UpdateThresholdMasterRequest,
	callback chan<- UpdateThresholdMasterAsyncResult,
) {
	path := "/{namespaceName}/master/threshold/{thresholdName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ThresholdName != nil && *request.ThresholdName != "" {
		path = strings.ReplaceAll(path, "{thresholdName}", core.ToString(*request.ThresholdName))
	} else {
		path = strings.ReplaceAll(path, "{thresholdName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Values != nil {
		var _values []interface{}
		for _, item := range request.Values {
			_values = append(_values, item)
		}
		bodies["values"] = _values
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateThresholdMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) UpdateThresholdMaster(
	request *UpdateThresholdMasterRequest,
) (*UpdateThresholdMasterResult, error) {
	callback := make(chan UpdateThresholdMasterAsyncResult, 1)
	go p.UpdateThresholdMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteThresholdMasterAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteThresholdMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteThresholdMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteThresholdMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteThresholdMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteThresholdMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteThresholdMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) DeleteThresholdMasterAsync(
	request *DeleteThresholdMasterRequest,
	callback chan<- DeleteThresholdMasterAsyncResult,
) {
	path := "/{namespaceName}/master/threshold/{thresholdName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ThresholdName != nil && *request.ThresholdName != "" {
		path = strings.ReplaceAll(path, "{thresholdName}", core.ToString(*request.ThresholdName))
	} else {
		path = strings.ReplaceAll(path, "{thresholdName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteThresholdMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) DeleteThresholdMaster(
	request *DeleteThresholdMasterRequest,
) (*DeleteThresholdMasterResult, error) {
	callback := make(chan DeleteThresholdMasterAsyncResult, 1)
	go p.DeleteThresholdMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2ExperienceRestClient,
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

func (p Gs2ExperienceRestClient) ExportMasterAsync(
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go exportMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) ExportMaster(
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

func getCurrentExperienceMasterAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentExperienceMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentExperienceMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentExperienceMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentExperienceMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentExperienceMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentExperienceMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) GetCurrentExperienceMasterAsync(
	request *GetCurrentExperienceMasterRequest,
	callback chan<- GetCurrentExperienceMasterAsyncResult,
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getCurrentExperienceMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) GetCurrentExperienceMaster(
	request *GetCurrentExperienceMasterRequest,
) (*GetCurrentExperienceMasterResult, error) {
	callback := make(chan GetCurrentExperienceMasterAsyncResult, 1)
	go p.GetCurrentExperienceMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentExperienceMasterAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentExperienceMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentExperienceMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentExperienceMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentExperienceMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentExperienceMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentExperienceMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) UpdateCurrentExperienceMasterAsync(
	request *UpdateCurrentExperienceMasterRequest,
	callback chan<- UpdateCurrentExperienceMasterAsyncResult,
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateCurrentExperienceMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) UpdateCurrentExperienceMaster(
	request *UpdateCurrentExperienceMasterRequest,
) (*UpdateCurrentExperienceMasterResult, error) {
	callback := make(chan UpdateCurrentExperienceMasterAsyncResult, 1)
	go p.UpdateCurrentExperienceMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentExperienceMasterFromGitHubAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentExperienceMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentExperienceMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentExperienceMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentExperienceMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentExperienceMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentExperienceMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) UpdateCurrentExperienceMasterFromGitHubAsync(
	request *UpdateCurrentExperienceMasterFromGitHubRequest,
	callback chan<- UpdateCurrentExperienceMasterFromGitHubAsyncResult,
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateCurrentExperienceMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) UpdateCurrentExperienceMasterFromGitHub(
	request *UpdateCurrentExperienceMasterFromGitHubRequest,
) (*UpdateCurrentExperienceMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentExperienceMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentExperienceMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeStatusesAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStatusesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStatusesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStatusesResult
	if asyncResult.Err != nil {
		callback <- DescribeStatusesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStatusesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStatusesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) DescribeStatusesAsync(
	request *DescribeStatusesRequest,
	callback chan<- DescribeStatusesAsyncResult,
) {
	path := "/{namespaceName}/user/me/status"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ExperienceName != nil {
		queryStrings["experienceName"] = core.ToString(*request.ExperienceName)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeStatusesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) DescribeStatuses(
	request *DescribeStatusesRequest,
) (*DescribeStatusesResult, error) {
	callback := make(chan DescribeStatusesAsyncResult, 1)
	go p.DescribeStatusesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeStatusesByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStatusesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStatusesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStatusesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeStatusesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStatusesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStatusesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) DescribeStatusesByUserIdAsync(
	request *DescribeStatusesByUserIdRequest,
	callback chan<- DescribeStatusesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status"
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
	if request.ExperienceName != nil {
		queryStrings["experienceName"] = core.ToString(*request.ExperienceName)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeStatusesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) DescribeStatusesByUserId(
	request *DescribeStatusesByUserIdRequest,
) (*DescribeStatusesByUserIdResult, error) {
	callback := make(chan DescribeStatusesByUserIdAsyncResult, 1)
	go p.DescribeStatusesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStatusAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- GetStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusResult
	if asyncResult.Err != nil {
		callback <- GetStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) GetStatusAsync(
	request *GetStatusRequest,
	callback chan<- GetStatusAsyncResult,
) {
	path := "/{namespaceName}/user/me/status/model/{experienceName}/property/{propertyId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) GetStatus(
	request *GetStatusRequest,
) (*GetStatusResult, error) {
	callback := make(chan GetStatusAsyncResult, 1)
	go p.GetStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStatusByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- GetStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) GetStatusByUserIdAsync(
	request *GetStatusByUserIdRequest,
	callback chan<- GetStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/model/{experienceName}/property/{propertyId}"
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
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) GetStatusByUserId(
	request *GetStatusByUserIdRequest,
) (*GetStatusByUserIdResult, error) {
	callback := make(chan GetStatusByUserIdAsyncResult, 1)
	go p.GetStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStatusWithSignatureAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- GetStatusWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusWithSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusWithSignatureResult
	if asyncResult.Err != nil {
		callback <- GetStatusWithSignatureAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStatusWithSignatureAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStatusWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) GetStatusWithSignatureAsync(
	request *GetStatusWithSignatureRequest,
	callback chan<- GetStatusWithSignatureAsyncResult,
) {
	path := "/{namespaceName}/user/me/status/model/{experienceName}/property/{propertyId}/signature"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.KeyId != nil {
		queryStrings["keyId"] = core.ToString(*request.KeyId)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getStatusWithSignatureAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) GetStatusWithSignature(
	request *GetStatusWithSignatureRequest,
) (*GetStatusWithSignatureResult, error) {
	callback := make(chan GetStatusWithSignatureAsyncResult, 1)
	go p.GetStatusWithSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStatusWithSignatureByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- GetStatusWithSignatureByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusWithSignatureByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusWithSignatureByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetStatusWithSignatureByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStatusWithSignatureByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStatusWithSignatureByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) GetStatusWithSignatureByUserIdAsync(
	request *GetStatusWithSignatureByUserIdRequest,
	callback chan<- GetStatusWithSignatureByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/model/{experienceName}/property/{propertyId}/signature"
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
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.KeyId != nil {
		queryStrings["keyId"] = core.ToString(*request.KeyId)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getStatusWithSignatureByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) GetStatusWithSignatureByUserId(
	request *GetStatusWithSignatureByUserIdRequest,
) (*GetStatusWithSignatureByUserIdResult, error) {
	callback := make(chan GetStatusWithSignatureByUserIdAsyncResult, 1)
	go p.GetStatusWithSignatureByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addExperienceByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- AddExperienceByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddExperienceByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddExperienceByUserIdResult
	if asyncResult.Err != nil {
		callback <- AddExperienceByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddExperienceByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddExperienceByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) AddExperienceByUserIdAsync(
	request *AddExperienceByUserIdRequest,
	callback chan<- AddExperienceByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/model/{experienceName}/property/{propertyId}"
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
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ExperienceValue != nil {
		bodies["experienceValue"] = *request.ExperienceValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go addExperienceByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) AddExperienceByUserId(
	request *AddExperienceByUserIdRequest,
) (*AddExperienceByUserIdResult, error) {
	callback := make(chan AddExperienceByUserIdAsyncResult, 1)
	go p.AddExperienceByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func subExperienceByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- SubExperienceByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubExperienceByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubExperienceByUserIdResult
	if asyncResult.Err != nil {
		callback <- SubExperienceByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SubExperienceByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SubExperienceByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) SubExperienceByUserIdAsync(
	request *SubExperienceByUserIdRequest,
	callback chan<- SubExperienceByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/model/{experienceName}/property/{propertyId}/sub"
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
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ExperienceValue != nil {
		bodies["experienceValue"] = *request.ExperienceValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go subExperienceByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) SubExperienceByUserId(
	request *SubExperienceByUserIdRequest,
) (*SubExperienceByUserIdResult, error) {
	callback := make(chan SubExperienceByUserIdAsyncResult, 1)
	go p.SubExperienceByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setExperienceByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- SetExperienceByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetExperienceByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetExperienceByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetExperienceByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetExperienceByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetExperienceByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) SetExperienceByUserIdAsync(
	request *SetExperienceByUserIdRequest,
	callback chan<- SetExperienceByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/model/{experienceName}/property/{propertyId}"
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
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ExperienceValue != nil {
		bodies["experienceValue"] = *request.ExperienceValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go setExperienceByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) SetExperienceByUserId(
	request *SetExperienceByUserIdRequest,
) (*SetExperienceByUserIdResult, error) {
	callback := make(chan SetExperienceByUserIdAsyncResult, 1)
	go p.SetExperienceByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addRankCapByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- AddRankCapByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddRankCapByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddRankCapByUserIdResult
	if asyncResult.Err != nil {
		callback <- AddRankCapByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddRankCapByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddRankCapByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) AddRankCapByUserIdAsync(
	request *AddRankCapByUserIdRequest,
	callback chan<- AddRankCapByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/model/{experienceName}/property/{propertyId}/cap"
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
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RankCapValue != nil {
		bodies["rankCapValue"] = *request.RankCapValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go addRankCapByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) AddRankCapByUserId(
	request *AddRankCapByUserIdRequest,
) (*AddRankCapByUserIdResult, error) {
	callback := make(chan AddRankCapByUserIdAsyncResult, 1)
	go p.AddRankCapByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func subRankCapByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- SubRankCapByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubRankCapByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubRankCapByUserIdResult
	if asyncResult.Err != nil {
		callback <- SubRankCapByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SubRankCapByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SubRankCapByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) SubRankCapByUserIdAsync(
	request *SubRankCapByUserIdRequest,
	callback chan<- SubRankCapByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/model/{experienceName}/property/{propertyId}/cap/sub"
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
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RankCapValue != nil {
		bodies["rankCapValue"] = *request.RankCapValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go subRankCapByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) SubRankCapByUserId(
	request *SubRankCapByUserIdRequest,
) (*SubRankCapByUserIdResult, error) {
	callback := make(chan SubRankCapByUserIdAsyncResult, 1)
	go p.SubRankCapByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setRankCapByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- SetRankCapByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRankCapByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRankCapByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetRankCapByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetRankCapByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetRankCapByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) SetRankCapByUserIdAsync(
	request *SetRankCapByUserIdRequest,
	callback chan<- SetRankCapByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/model/{experienceName}/property/{propertyId}/cap"
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
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RankCapValue != nil {
		bodies["rankCapValue"] = *request.RankCapValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go setRankCapByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) SetRankCapByUserId(
	request *SetRankCapByUserIdRequest,
) (*SetRankCapByUserIdResult, error) {
	callback := make(chan SetRankCapByUserIdAsyncResult, 1)
	go p.SetRankCapByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteStatusByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) DeleteStatusByUserIdAsync(
	request *DeleteStatusByUserIdRequest,
	callback chan<- DeleteStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/model/{experienceName}/property/{propertyId}"
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
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go deleteStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) DeleteStatusByUserId(
	request *DeleteStatusByUserIdRequest,
) (*DeleteStatusByUserIdResult, error) {
	callback := make(chan DeleteStatusByUserIdAsyncResult, 1)
	go p.DeleteStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyRankAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyRankAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyRankAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyRankResult
	if asyncResult.Err != nil {
		callback <- VerifyRankAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyRankAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyRankAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) VerifyRankAsync(
	request *VerifyRankRequest,
	callback chan<- VerifyRankAsyncResult,
) {
	path := "/{namespaceName}/user/me/status/{experienceName}/verify/rank/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.RankValue != nil {
		bodies["rankValue"] = *request.RankValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go verifyRankAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) VerifyRank(
	request *VerifyRankRequest,
) (*VerifyRankResult, error) {
	callback := make(chan VerifyRankAsyncResult, 1)
	go p.VerifyRankAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyRankByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyRankByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyRankByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyRankByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyRankByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyRankByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyRankByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) VerifyRankByUserIdAsync(
	request *VerifyRankByUserIdRequest,
	callback chan<- VerifyRankByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/{experienceName}/verify/rank/{verifyType}"
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
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.RankValue != nil {
		bodies["rankValue"] = *request.RankValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go verifyRankByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) VerifyRankByUserId(
	request *VerifyRankByUserIdRequest,
) (*VerifyRankByUserIdResult, error) {
	callback := make(chan VerifyRankByUserIdAsyncResult, 1)
	go p.VerifyRankByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyRankCapAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyRankCapAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyRankCapAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyRankCapResult
	if asyncResult.Err != nil {
		callback <- VerifyRankCapAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyRankCapAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyRankCapAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) VerifyRankCapAsync(
	request *VerifyRankCapRequest,
	callback chan<- VerifyRankCapAsyncResult,
) {
	path := "/{namespaceName}/user/me/status/{experienceName}/verify/rankCap/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.RankCapValue != nil {
		bodies["rankCapValue"] = *request.RankCapValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go verifyRankCapAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) VerifyRankCap(
	request *VerifyRankCapRequest,
) (*VerifyRankCapResult, error) {
	callback := make(chan VerifyRankCapAsyncResult, 1)
	go p.VerifyRankCapAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyRankCapByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyRankCapByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyRankCapByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyRankCapByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyRankCapByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyRankCapByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyRankCapByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) VerifyRankCapByUserIdAsync(
	request *VerifyRankCapByUserIdRequest,
	callback chan<- VerifyRankCapByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/{experienceName}/verify/rankCap/{verifyType}"
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
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.RankCapValue != nil {
		bodies["rankCapValue"] = *request.RankCapValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go verifyRankCapByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) VerifyRankCapByUserId(
	request *VerifyRankCapByUserIdRequest,
) (*VerifyRankCapByUserIdResult, error) {
	callback := make(chan VerifyRankCapByUserIdAsyncResult, 1)
	go p.VerifyRankCapByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addExperienceByStampSheetAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- AddExperienceByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddExperienceByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddExperienceByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AddExperienceByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddExperienceByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddExperienceByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) AddExperienceByStampSheetAsync(
	request *AddExperienceByStampSheetRequest,
	callback chan<- AddExperienceByStampSheetAsyncResult,
) {
	path := "/stamp/experience/add"

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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go addExperienceByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) AddExperienceByStampSheet(
	request *AddExperienceByStampSheetRequest,
) (*AddExperienceByStampSheetResult, error) {
	callback := make(chan AddExperienceByStampSheetAsyncResult, 1)
	go p.AddExperienceByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func subExperienceByStampTaskAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- SubExperienceByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubExperienceByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubExperienceByStampTaskResult
	if asyncResult.Err != nil {
		callback <- SubExperienceByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SubExperienceByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SubExperienceByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) SubExperienceByStampTaskAsync(
	request *SubExperienceByStampTaskRequest,
	callback chan<- SubExperienceByStampTaskAsyncResult,
) {
	path := "/stamp/experience/sub"

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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go subExperienceByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) SubExperienceByStampTask(
	request *SubExperienceByStampTaskRequest,
) (*SubExperienceByStampTaskResult, error) {
	callback := make(chan SubExperienceByStampTaskAsyncResult, 1)
	go p.SubExperienceByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addRankCapByStampSheetAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- AddRankCapByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddRankCapByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddRankCapByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AddRankCapByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddRankCapByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddRankCapByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) AddRankCapByStampSheetAsync(
	request *AddRankCapByStampSheetRequest,
	callback chan<- AddRankCapByStampSheetAsyncResult,
) {
	path := "/stamp/rankCap/add"

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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go addRankCapByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) AddRankCapByStampSheet(
	request *AddRankCapByStampSheetRequest,
) (*AddRankCapByStampSheetResult, error) {
	callback := make(chan AddRankCapByStampSheetAsyncResult, 1)
	go p.AddRankCapByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func subRankCapByStampTaskAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- SubRankCapByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubRankCapByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubRankCapByStampTaskResult
	if asyncResult.Err != nil {
		callback <- SubRankCapByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SubRankCapByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SubRankCapByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) SubRankCapByStampTaskAsync(
	request *SubRankCapByStampTaskRequest,
	callback chan<- SubRankCapByStampTaskAsyncResult,
) {
	path := "/stamp/rankCap/sub"

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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go subRankCapByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) SubRankCapByStampTask(
	request *SubRankCapByStampTaskRequest,
) (*SubRankCapByStampTaskResult, error) {
	callback := make(chan SubRankCapByStampTaskAsyncResult, 1)
	go p.SubRankCapByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setRankCapByStampSheetAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- SetRankCapByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRankCapByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRankCapByStampSheetResult
	if asyncResult.Err != nil {
		callback <- SetRankCapByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetRankCapByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetRankCapByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) SetRankCapByStampSheetAsync(
	request *SetRankCapByStampSheetRequest,
	callback chan<- SetRankCapByStampSheetAsyncResult,
) {
	path := "/stamp/rankCap/set"

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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go setRankCapByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) SetRankCapByStampSheet(
	request *SetRankCapByStampSheetRequest,
) (*SetRankCapByStampSheetResult, error) {
	callback := make(chan SetRankCapByStampSheetAsyncResult, 1)
	go p.SetRankCapByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func multiplyAcquireActionsByUserIdAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- MultiplyAcquireActionsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MultiplyAcquireActionsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MultiplyAcquireActionsByUserIdResult
	if asyncResult.Err != nil {
		callback <- MultiplyAcquireActionsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MultiplyAcquireActionsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- MultiplyAcquireActionsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) MultiplyAcquireActionsByUserIdAsync(
	request *MultiplyAcquireActionsByUserIdRequest,
	callback chan<- MultiplyAcquireActionsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/model/{experienceName}/property/{propertyId}/acquire/rate/{rateName}/multiply"
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
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}
	if request.RateName != nil && *request.RateName != "" {
		path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
	} else {
		path = strings.ReplaceAll(path, "{rateName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AcquireActions != nil {
		var _acquireActions []interface{}
		for _, item := range request.AcquireActions {
			_acquireActions = append(_acquireActions, item)
		}
		bodies["acquireActions"] = _acquireActions
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go multiplyAcquireActionsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) MultiplyAcquireActionsByUserId(
	request *MultiplyAcquireActionsByUserIdRequest,
) (*MultiplyAcquireActionsByUserIdResult, error) {
	callback := make(chan MultiplyAcquireActionsByUserIdAsyncResult, 1)
	go p.MultiplyAcquireActionsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func multiplyAcquireActionsByStampSheetAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- MultiplyAcquireActionsByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MultiplyAcquireActionsByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MultiplyAcquireActionsByStampSheetResult
	if asyncResult.Err != nil {
		callback <- MultiplyAcquireActionsByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MultiplyAcquireActionsByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- MultiplyAcquireActionsByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) MultiplyAcquireActionsByStampSheetAsync(
	request *MultiplyAcquireActionsByStampSheetRequest,
	callback chan<- MultiplyAcquireActionsByStampSheetAsyncResult,
) {
	path := "/stamp/form/acquire"

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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go multiplyAcquireActionsByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) MultiplyAcquireActionsByStampSheet(
	request *MultiplyAcquireActionsByStampSheetRequest,
) (*MultiplyAcquireActionsByStampSheetResult, error) {
	callback := make(chan MultiplyAcquireActionsByStampSheetAsyncResult, 1)
	go p.MultiplyAcquireActionsByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyRankByStampTaskAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyRankByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyRankByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyRankByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyRankByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyRankByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyRankByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) VerifyRankByStampTaskAsync(
	request *VerifyRankByStampTaskRequest,
	callback chan<- VerifyRankByStampTaskAsyncResult,
) {
	path := "/stamp/rank/verify"

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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go verifyRankByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) VerifyRankByStampTask(
	request *VerifyRankByStampTaskRequest,
) (*VerifyRankByStampTaskResult, error) {
	callback := make(chan VerifyRankByStampTaskAsyncResult, 1)
	go p.VerifyRankByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyRankCapByStampTaskAsyncHandler(
	client Gs2ExperienceRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyRankCapByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyRankCapByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyRankCapByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyRankCapByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyRankCapByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyRankCapByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceRestClient) VerifyRankCapByStampTaskAsync(
	request *VerifyRankCapByStampTaskRequest,
	callback chan<- VerifyRankCapByStampTaskAsyncResult,
) {
	path := "/stamp/rankCap/verify"

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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go verifyRankCapByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("experience").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceRestClient) VerifyRankCapByStampTask(
	request *VerifyRankCapByStampTaskRequest,
) (*VerifyRankCapByStampTaskResult, error) {
	callback := make(chan VerifyRankCapByStampTaskAsyncResult, 1)
	go p.VerifyRankCapByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
