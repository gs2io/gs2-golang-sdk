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

package chat

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2ChatRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2ChatRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2ChatRestClient,
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

func (p Gs2ChatRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) DescribeNamespaces(
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
	client Gs2ChatRestClient,
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

func (p Gs2ChatRestClient) CreateNamespaceAsync(
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
	if request.AllowCreateRoom != nil {
		bodies["allowCreateRoom"] = *request.AllowCreateRoom
	}
	if request.PostMessageScript != nil {
		bodies["postMessageScript"] = request.PostMessageScript.ToDict()
	}
	if request.CreateRoomScript != nil {
		bodies["createRoomScript"] = request.CreateRoomScript.ToDict()
	}
	if request.DeleteRoomScript != nil {
		bodies["deleteRoomScript"] = request.DeleteRoomScript.ToDict()
	}
	if request.SubscribeRoomScript != nil {
		bodies["subscribeRoomScript"] = request.SubscribeRoomScript.ToDict()
	}
	if request.UnsubscribeRoomScript != nil {
		bodies["unsubscribeRoomScript"] = request.UnsubscribeRoomScript.ToDict()
	}
	if request.PostNotification != nil {
		bodies["postNotification"] = request.PostNotification.ToDict()
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
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) CreateNamespace(
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
	client Gs2ChatRestClient,
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

func (p Gs2ChatRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) GetNamespaceStatus(
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
	client Gs2ChatRestClient,
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

func (p Gs2ChatRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) GetNamespace(
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
	client Gs2ChatRestClient,
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

func (p Gs2ChatRestClient) UpdateNamespaceAsync(
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
	if request.AllowCreateRoom != nil {
		bodies["allowCreateRoom"] = *request.AllowCreateRoom
	}
	if request.PostMessageScript != nil {
		bodies["postMessageScript"] = request.PostMessageScript.ToDict()
	}
	if request.CreateRoomScript != nil {
		bodies["createRoomScript"] = request.CreateRoomScript.ToDict()
	}
	if request.DeleteRoomScript != nil {
		bodies["deleteRoomScript"] = request.DeleteRoomScript.ToDict()
	}
	if request.SubscribeRoomScript != nil {
		bodies["subscribeRoomScript"] = request.SubscribeRoomScript.ToDict()
	}
	if request.UnsubscribeRoomScript != nil {
		bodies["unsubscribeRoomScript"] = request.UnsubscribeRoomScript.ToDict()
	}
	if request.PostNotification != nil {
		bodies["postNotification"] = request.PostNotification.ToDict()
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
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) UpdateNamespace(
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
	client Gs2ChatRestClient,
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

func (p Gs2ChatRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) DeleteNamespace(
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
	client Gs2ChatRestClient,
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

func (p Gs2ChatRestClient) DumpUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) DumpUserDataByUserId(
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
	client Gs2ChatRestClient,
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

func (p Gs2ChatRestClient) CheckDumpUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) CheckDumpUserDataByUserId(
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
	client Gs2ChatRestClient,
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

func (p Gs2ChatRestClient) CleanUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) CleanUserDataByUserId(
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
	client Gs2ChatRestClient,
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

func (p Gs2ChatRestClient) CheckCleanUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) CheckCleanUserDataByUserId(
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

func describeRoomsAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRoomsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRoomsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRoomsResult
	if asyncResult.Err != nil {
		callback <- DescribeRoomsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRoomsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRoomsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) DescribeRoomsAsync(
	request *DescribeRoomsRequest,
	callback chan<- DescribeRoomsAsyncResult,
) {
	path := "/{namespaceName}/room"
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

	go describeRoomsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) DescribeRooms(
	request *DescribeRoomsRequest,
) (*DescribeRoomsResult, error) {
	callback := make(chan DescribeRoomsAsyncResult, 1)
	go p.DescribeRoomsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createRoomAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- CreateRoomAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateRoomAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateRoomResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.allowUserIds.notInclude" {
				asyncResult.Err = gs2err.SetClientError(NoAccessPrivileges{})
			}
		}
		callback <- CreateRoomAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateRoomAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateRoomAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) CreateRoomAsync(
	request *CreateRoomRequest,
	callback chan<- CreateRoomAsyncResult,
) {
	path := "/{namespaceName}/room/user"
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
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
	}
	if request.WhiteListUserIds != nil {
		var _whiteListUserIds []interface{}
		for _, item := range request.WhiteListUserIds {
			_whiteListUserIds = append(_whiteListUserIds, item)
		}
		bodies["whiteListUserIds"] = _whiteListUserIds
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

	go createRoomAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) CreateRoom(
	request *CreateRoomRequest,
) (*CreateRoomResult, error) {
	callback := make(chan CreateRoomAsyncResult, 1)
	go p.CreateRoomAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createRoomFromBackendAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- CreateRoomFromBackendAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateRoomFromBackendAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateRoomFromBackendResult
	if asyncResult.Err != nil {
		callback <- CreateRoomFromBackendAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateRoomFromBackendAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateRoomFromBackendAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) CreateRoomFromBackendAsync(
	request *CreateRoomFromBackendRequest,
	callback chan<- CreateRoomFromBackendAsyncResult,
) {
	path := "/{namespaceName}/room"
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
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
	}
	if request.WhiteListUserIds != nil {
		var _whiteListUserIds []interface{}
		for _, item := range request.WhiteListUserIds {
			_whiteListUserIds = append(_whiteListUserIds, item)
		}
		bodies["whiteListUserIds"] = _whiteListUserIds
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

	go createRoomFromBackendAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) CreateRoomFromBackend(
	request *CreateRoomFromBackendRequest,
) (*CreateRoomFromBackendResult, error) {
	callback := make(chan CreateRoomFromBackendAsyncResult, 1)
	go p.CreateRoomFromBackendAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRoomAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- GetRoomAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRoomAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRoomResult
	if asyncResult.Err != nil {
		callback <- GetRoomAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRoomAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRoomAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) GetRoomAsync(
	request *GetRoomRequest,
	callback chan<- GetRoomAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getRoomAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) GetRoom(
	request *GetRoomRequest,
) (*GetRoomResult, error) {
	callback := make(chan GetRoomAsyncResult, 1)
	go p.GetRoomAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateRoomAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateRoomAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateRoomAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateRoomResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.allowUserIds.notInclude" {
				asyncResult.Err = gs2err.SetClientError(NoAccessPrivileges{})
			}
		}
		callback <- UpdateRoomAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateRoomAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateRoomAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) UpdateRoomAsync(
	request *UpdateRoomRequest,
	callback chan<- UpdateRoomAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}/user"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
	}
	if request.WhiteListUserIds != nil {
		var _whiteListUserIds []interface{}
		for _, item := range request.WhiteListUserIds {
			_whiteListUserIds = append(_whiteListUserIds, item)
		}
		bodies["whiteListUserIds"] = _whiteListUserIds
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

	go updateRoomAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) UpdateRoom(
	request *UpdateRoomRequest,
) (*UpdateRoomResult, error) {
	callback := make(chan UpdateRoomAsyncResult, 1)
	go p.UpdateRoomAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateRoomFromBackendAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateRoomFromBackendAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateRoomFromBackendAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateRoomFromBackendResult
	if asyncResult.Err != nil {
		callback <- UpdateRoomFromBackendAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateRoomFromBackendAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateRoomFromBackendAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) UpdateRoomFromBackendAsync(
	request *UpdateRoomFromBackendRequest,
	callback chan<- UpdateRoomFromBackendAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
	}
	if request.WhiteListUserIds != nil {
		var _whiteListUserIds []interface{}
		for _, item := range request.WhiteListUserIds {
			_whiteListUserIds = append(_whiteListUserIds, item)
		}
		bodies["whiteListUserIds"] = _whiteListUserIds
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
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

	go updateRoomFromBackendAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) UpdateRoomFromBackend(
	request *UpdateRoomFromBackendRequest,
) (*UpdateRoomFromBackendResult, error) {
	callback := make(chan UpdateRoomFromBackendAsyncResult, 1)
	go p.UpdateRoomFromBackendAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteRoomAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteRoomAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRoomAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRoomResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.allowUserIds.notInclude" {
				asyncResult.Err = gs2err.SetClientError(NoAccessPrivileges{})
			}
		}
		callback <- DeleteRoomAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteRoomAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteRoomAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) DeleteRoomAsync(
	request *DeleteRoomRequest,
	callback chan<- DeleteRoomAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}/user"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go deleteRoomAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) DeleteRoom(
	request *DeleteRoomRequest,
) (*DeleteRoomResult, error) {
	callback := make(chan DeleteRoomAsyncResult, 1)
	go p.DeleteRoomAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteRoomFromBackendAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteRoomFromBackendAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRoomFromBackendAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRoomFromBackendResult
	if asyncResult.Err != nil {
		callback <- DeleteRoomFromBackendAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteRoomFromBackendAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteRoomFromBackendAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) DeleteRoomFromBackendAsync(
	request *DeleteRoomFromBackendRequest,
	callback chan<- DeleteRoomFromBackendAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go deleteRoomFromBackendAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) DeleteRoomFromBackend(
	request *DeleteRoomFromBackendRequest,
) (*DeleteRoomFromBackendResult, error) {
	callback := make(chan DeleteRoomFromBackendAsyncResult, 1)
	go p.DeleteRoomFromBackendAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeMessagesAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMessagesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMessagesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMessagesResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.allowUserIds.notInclude" {
				asyncResult.Err = gs2err.SetClientError(NoAccessPrivileges{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.password.require" {
				asyncResult.Err = gs2err.SetClientError(PasswordRequired{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
			}
		}
		callback <- DescribeMessagesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMessagesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMessagesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) DescribeMessagesAsync(
	request *DescribeMessagesRequest,
	callback chan<- DescribeMessagesAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}/message"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Password != nil {
		queryStrings["password"] = core.ToString(*request.Password)
	}
	if request.StartAt != nil {
		queryStrings["startAt"] = core.ToString(*request.StartAt)
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

	go describeMessagesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) DescribeMessages(
	request *DescribeMessagesRequest,
) (*DescribeMessagesResult, error) {
	callback := make(chan DescribeMessagesAsyncResult, 1)
	go p.DescribeMessagesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeMessagesByUserIdAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMessagesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMessagesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMessagesByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.allowUserIds.notInclude" {
				asyncResult.Err = gs2err.SetClientError(NoAccessPrivileges{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.password.require" {
				asyncResult.Err = gs2err.SetClientError(PasswordRequired{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
			}
		}
		callback <- DescribeMessagesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMessagesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMessagesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) DescribeMessagesByUserIdAsync(
	request *DescribeMessagesByUserIdRequest,
	callback chan<- DescribeMessagesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}/message/get"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Password != nil {
		queryStrings["password"] = core.ToString(*request.Password)
	}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}
	if request.StartAt != nil {
		queryStrings["startAt"] = core.ToString(*request.StartAt)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeMessagesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) DescribeMessagesByUserId(
	request *DescribeMessagesByUserIdRequest,
) (*DescribeMessagesByUserIdResult, error) {
	callback := make(chan DescribeMessagesByUserIdAsyncResult, 1)
	go p.DescribeMessagesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func postAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- PostAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PostAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PostResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.allowUserIds.notInclude" {
				asyncResult.Err = gs2err.SetClientError(NoAccessPrivileges{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.password.require" {
				asyncResult.Err = gs2err.SetClientError(PasswordRequired{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
			}
		}
		callback <- PostAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PostAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PostAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) PostAsync(
	request *PostRequest,
	callback chan<- PostAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}/message"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Category != nil {
		bodies["category"] = *request.Category
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
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

	go postAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) Post(
	request *PostRequest,
) (*PostResult, error) {
	callback := make(chan PostAsyncResult, 1)
	go p.PostAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func postByUserIdAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- PostByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PostByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PostByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.allowUserIds.notInclude" {
				asyncResult.Err = gs2err.SetClientError(NoAccessPrivileges{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.password.require" {
				asyncResult.Err = gs2err.SetClientError(PasswordRequired{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
			}
		}
		callback <- PostByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PostByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PostByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) PostByUserIdAsync(
	request *PostByUserIdRequest,
	callback chan<- PostByUserIdAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}/message/user/{userId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Category != nil {
		bodies["category"] = *request.Category
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
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

	go postByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) PostByUserId(
	request *PostByUserIdRequest,
) (*PostByUserIdResult, error) {
	callback := make(chan PostByUserIdAsyncResult, 1)
	go p.PostByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getMessageAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- GetMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMessageResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.allowUserIds.notInclude" {
				asyncResult.Err = gs2err.SetClientError(NoAccessPrivileges{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.password.require" {
				asyncResult.Err = gs2err.SetClientError(PasswordRequired{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
			}
		}
		callback <- GetMessageAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMessageAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) GetMessageAsync(
	request *GetMessageRequest,
	callback chan<- GetMessageAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}/message/{messageName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}
	if request.MessageName != nil && *request.MessageName != "" {
		path = strings.ReplaceAll(path, "{messageName}", core.ToString(*request.MessageName))
	} else {
		path = strings.ReplaceAll(path, "{messageName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Password != nil {
		queryStrings["password"] = core.ToString(*request.Password)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getMessageAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) GetMessage(
	request *GetMessageRequest,
) (*GetMessageResult, error) {
	callback := make(chan GetMessageAsyncResult, 1)
	go p.GetMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getMessageByUserIdAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- GetMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMessageByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.allowUserIds.notInclude" {
				asyncResult.Err = gs2err.SetClientError(NoAccessPrivileges{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.password.require" {
				asyncResult.Err = gs2err.SetClientError(PasswordRequired{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "room.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
			}
		}
		callback <- GetMessageByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMessageByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) GetMessageByUserIdAsync(
	request *GetMessageByUserIdRequest,
	callback chan<- GetMessageByUserIdAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}/message/{messageName}/get"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}
	if request.MessageName != nil && *request.MessageName != "" {
		path = strings.ReplaceAll(path, "{messageName}", core.ToString(*request.MessageName))
	} else {
		path = strings.ReplaceAll(path, "{messageName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Password != nil {
		queryStrings["password"] = core.ToString(*request.Password)
	}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getMessageByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) GetMessageByUserId(
	request *GetMessageByUserIdRequest,
) (*GetMessageByUserIdResult, error) {
	callback := make(chan GetMessageByUserIdAsyncResult, 1)
	go p.GetMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteMessageAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMessageResult
	if asyncResult.Err != nil {
		callback <- DeleteMessageAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMessageAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) DeleteMessageAsync(
	request *DeleteMessageRequest,
	callback chan<- DeleteMessageAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}/message/{messageName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}
	if request.MessageName != nil && *request.MessageName != "" {
		path = strings.ReplaceAll(path, "{messageName}", core.ToString(*request.MessageName))
	} else {
		path = strings.ReplaceAll(path, "{messageName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go deleteMessageAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) DeleteMessage(
	request *DeleteMessageRequest,
) (*DeleteMessageResult, error) {
	callback := make(chan DeleteMessageAsyncResult, 1)
	go p.DeleteMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribesAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribesResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscribesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) DescribeSubscribesAsync(
	request *DescribeSubscribesRequest,
	callback chan<- DescribeSubscribesAsyncResult,
) {
	path := "/{namespaceName}/user/me/room/subscribe"
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeSubscribesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) DescribeSubscribes(
	request *DescribeSubscribesRequest,
) (*DescribeSubscribesResult, error) {
	callback := make(chan DescribeSubscribesAsyncResult, 1)
	go p.DescribeSubscribesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribesByUserIdAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscribesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) DescribeSubscribesByUserIdAsync(
	request *DescribeSubscribesByUserIdRequest,
	callback chan<- DescribeSubscribesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/room/subscribe"
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeSubscribesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) DescribeSubscribesByUserId(
	request *DescribeSubscribesByUserIdRequest,
) (*DescribeSubscribesByUserIdResult, error) {
	callback := make(chan DescribeSubscribesByUserIdAsyncResult, 1)
	go p.DescribeSubscribesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribesByRoomNameAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribesByRoomNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribesByRoomNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribesByRoomNameResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscribesByRoomNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribesByRoomNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribesByRoomNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) DescribeSubscribesByRoomNameAsync(
	request *DescribeSubscribesByRoomNameRequest,
	callback chan<- DescribeSubscribesByRoomNameAsyncResult,
) {
	path := "/{namespaceName}/room/{roomName}/subscribe"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
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

	go describeSubscribesByRoomNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) DescribeSubscribesByRoomName(
	request *DescribeSubscribesByRoomNameRequest,
) (*DescribeSubscribesByRoomNameResult, error) {
	callback := make(chan DescribeSubscribesByRoomNameAsyncResult, 1)
	go p.DescribeSubscribesByRoomNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func subscribeAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- SubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubscribeResult
	if asyncResult.Err != nil {
		callback <- SubscribeAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SubscribeAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) SubscribeAsync(
	request *SubscribeRequest,
	callback chan<- SubscribeAsyncResult,
) {
	path := "/{namespaceName}/user/me/room/{roomName}/subscribe"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.NotificationTypes != nil {
		var _notificationTypes []interface{}
		for _, item := range request.NotificationTypes {
			_notificationTypes = append(_notificationTypes, item)
		}
		bodies["notificationTypes"] = _notificationTypes
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

	go subscribeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) Subscribe(
	request *SubscribeRequest,
) (*SubscribeResult, error) {
	callback := make(chan SubscribeAsyncResult, 1)
	go p.SubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func subscribeByUserIdAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- SubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubscribeByUserIdResult
	if asyncResult.Err != nil {
		callback <- SubscribeByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SubscribeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) SubscribeByUserIdAsync(
	request *SubscribeByUserIdRequest,
	callback chan<- SubscribeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/room/{roomName}/subscribe"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.NotificationTypes != nil {
		var _notificationTypes []interface{}
		for _, item := range request.NotificationTypes {
			_notificationTypes = append(_notificationTypes, item)
		}
		bodies["notificationTypes"] = _notificationTypes
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

	go subscribeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) SubscribeByUserId(
	request *SubscribeByUserIdRequest,
) (*SubscribeByUserIdResult, error) {
	callback := make(chan SubscribeByUserIdAsyncResult, 1)
	go p.SubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSubscribeAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- GetSubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeResult
	if asyncResult.Err != nil {
		callback <- GetSubscribeAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) GetSubscribeAsync(
	request *GetSubscribeRequest,
	callback chan<- GetSubscribeAsyncResult,
) {
	path := "/{namespaceName}/user/me/room/{roomName}/subscribe"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
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

	go getSubscribeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) GetSubscribe(
	request *GetSubscribeRequest,
) (*GetSubscribeResult, error) {
	callback := make(chan GetSubscribeAsyncResult, 1)
	go p.GetSubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSubscribeByUserIdAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- GetSubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetSubscribeByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) GetSubscribeByUserIdAsync(
	request *GetSubscribeByUserIdRequest,
	callback chan<- GetSubscribeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/room/{roomName}/subscribe"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}
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

	go getSubscribeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) GetSubscribeByUserId(
	request *GetSubscribeByUserIdRequest,
) (*GetSubscribeByUserIdResult, error) {
	callback := make(chan GetSubscribeByUserIdAsyncResult, 1)
	go p.GetSubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateNotificationTypeAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateNotificationTypeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateNotificationTypeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateNotificationTypeResult
	if asyncResult.Err != nil {
		callback <- UpdateNotificationTypeAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateNotificationTypeAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateNotificationTypeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) UpdateNotificationTypeAsync(
	request *UpdateNotificationTypeRequest,
	callback chan<- UpdateNotificationTypeAsyncResult,
) {
	path := "/{namespaceName}/user/me/room/{roomName}/subscribe/notification"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.NotificationTypes != nil {
		var _notificationTypes []interface{}
		for _, item := range request.NotificationTypes {
			_notificationTypes = append(_notificationTypes, item)
		}
		bodies["notificationTypes"] = _notificationTypes
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

	go updateNotificationTypeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) UpdateNotificationType(
	request *UpdateNotificationTypeRequest,
) (*UpdateNotificationTypeResult, error) {
	callback := make(chan UpdateNotificationTypeAsyncResult, 1)
	go p.UpdateNotificationTypeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateNotificationTypeByUserIdAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateNotificationTypeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateNotificationTypeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateNotificationTypeByUserIdResult
	if asyncResult.Err != nil {
		callback <- UpdateNotificationTypeByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateNotificationTypeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateNotificationTypeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) UpdateNotificationTypeByUserIdAsync(
	request *UpdateNotificationTypeByUserIdRequest,
	callback chan<- UpdateNotificationTypeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/room/{roomName}/subscribe/notification"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.NotificationTypes != nil {
		var _notificationTypes []interface{}
		for _, item := range request.NotificationTypes {
			_notificationTypes = append(_notificationTypes, item)
		}
		bodies["notificationTypes"] = _notificationTypes
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

	go updateNotificationTypeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) UpdateNotificationTypeByUserId(
	request *UpdateNotificationTypeByUserIdRequest,
) (*UpdateNotificationTypeByUserIdResult, error) {
	callback := make(chan UpdateNotificationTypeByUserIdAsyncResult, 1)
	go p.UpdateNotificationTypeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func unsubscribeAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- UnsubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UnsubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UnsubscribeResult
	if asyncResult.Err != nil {
		callback <- UnsubscribeAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UnsubscribeAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UnsubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) UnsubscribeAsync(
	request *UnsubscribeRequest,
	callback chan<- UnsubscribeAsyncResult,
) {
	path := "/{namespaceName}/user/me/room/{roomName}/subscribe"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go unsubscribeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) Unsubscribe(
	request *UnsubscribeRequest,
) (*UnsubscribeResult, error) {
	callback := make(chan UnsubscribeAsyncResult, 1)
	go p.UnsubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func unsubscribeByUserIdAsyncHandler(
	client Gs2ChatRestClient,
	job *core.NetworkJob,
	callback chan<- UnsubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UnsubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UnsubscribeByUserIdResult
	if asyncResult.Err != nil {
		callback <- UnsubscribeByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UnsubscribeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UnsubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ChatRestClient) UnsubscribeByUserIdAsync(
	request *UnsubscribeByUserIdRequest,
	callback chan<- UnsubscribeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/room/{roomName}/subscribe"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RoomName != nil && *request.RoomName != "" {
		path = strings.ReplaceAll(path, "{roomName}", core.ToString(*request.RoomName))
	} else {
		path = strings.ReplaceAll(path, "{roomName}", "null")
	}
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

	go unsubscribeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("chat").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ChatRestClient) UnsubscribeByUserId(
	request *UnsubscribeByUserIdRequest,
) (*UnsubscribeByUserIdResult, error) {
	callback := make(chan UnsubscribeByUserIdAsyncResult, 1)
	go p.UnsubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
