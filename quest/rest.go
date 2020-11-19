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

package quest

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

type Gs2QuestRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2QuestRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2QuestRestClient,
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

func (p Gs2QuestRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DescribeNamespaces(
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
	client Gs2QuestRestClient,
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

func (p Gs2QuestRestClient) CreateNamespaceAsync(
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
	if request.StartQuestScript != nil {
		bodies["startQuestScript"] = request.StartQuestScript.ToDict()
	}
	if request.CompleteQuestScript != nil {
		bodies["completeQuestScript"] = request.CompleteQuestScript.ToDict()
	}
	if request.FailedQuestScript != nil {
		bodies["failedQuestScript"] = request.FailedQuestScript.ToDict()
	}
	if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
		bodies["queueNamespaceId"] = *request.QueueNamespaceId
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) CreateNamespace(
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
	client Gs2QuestRestClient,
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

func (p Gs2QuestRestClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	path := "/{namespaceName}/status"
	if request.NamespaceName != nil {
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
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) GetNamespaceStatus(
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
	client Gs2QuestRestClient,
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

func (p Gs2QuestRestClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil {
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
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) GetNamespace(
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
	client Gs2QuestRestClient,
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

func (p Gs2QuestRestClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.StartQuestScript != nil {
		bodies["startQuestScript"] = request.StartQuestScript.ToDict()
	}
	if request.CompleteQuestScript != nil {
		bodies["completeQuestScript"] = request.CompleteQuestScript.ToDict()
	}
	if request.FailedQuestScript != nil {
		bodies["failedQuestScript"] = request.FailedQuestScript.ToDict()
	}
	if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
		bodies["queueNamespaceId"] = *request.QueueNamespaceId
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) UpdateNamespace(
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
	client Gs2QuestRestClient,
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

func (p Gs2QuestRestClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil {
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
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DeleteNamespace(
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

func describeQuestGroupModelMastersAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeQuestGroupModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeQuestGroupModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeQuestGroupModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeQuestGroupModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeQuestGroupModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DescribeQuestGroupModelMastersAsync(
	request *DescribeQuestGroupModelMastersRequest,
	callback chan<- DescribeQuestGroupModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/group"
	if request.NamespaceName != nil {
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

	go describeQuestGroupModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DescribeQuestGroupModelMasters(
	request *DescribeQuestGroupModelMastersRequest,
) (*DescribeQuestGroupModelMastersResult, error) {
	callback := make(chan DescribeQuestGroupModelMastersAsyncResult, 1)
	go p.DescribeQuestGroupModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createQuestGroupModelMasterAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- CreateQuestGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateQuestGroupModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateQuestGroupModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateQuestGroupModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateQuestGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) CreateQuestGroupModelMasterAsync(
	request *CreateQuestGroupModelMasterRequest,
	callback chan<- CreateQuestGroupModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group"
	if request.NamespaceName != nil {
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
	if request.ChallengePeriodEventId != nil && *request.ChallengePeriodEventId != "" {
		bodies["challengePeriodEventId"] = *request.ChallengePeriodEventId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createQuestGroupModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) CreateQuestGroupModelMaster(
	request *CreateQuestGroupModelMasterRequest,
) (*CreateQuestGroupModelMasterResult, error) {
	callback := make(chan CreateQuestGroupModelMasterAsyncResult, 1)
	go p.CreateQuestGroupModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getQuestGroupModelMasterAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- GetQuestGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetQuestGroupModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetQuestGroupModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetQuestGroupModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetQuestGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) GetQuestGroupModelMasterAsync(
	request *GetQuestGroupModelMasterRequest,
	callback chan<- GetQuestGroupModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{questGroupName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getQuestGroupModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) GetQuestGroupModelMaster(
	request *GetQuestGroupModelMasterRequest,
) (*GetQuestGroupModelMasterResult, error) {
	callback := make(chan GetQuestGroupModelMasterAsyncResult, 1)
	go p.GetQuestGroupModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateQuestGroupModelMasterAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateQuestGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateQuestGroupModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateQuestGroupModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateQuestGroupModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateQuestGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) UpdateQuestGroupModelMasterAsync(
	request *UpdateQuestGroupModelMasterRequest,
	callback chan<- UpdateQuestGroupModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{questGroupName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ChallengePeriodEventId != nil && *request.ChallengePeriodEventId != "" {
		bodies["challengePeriodEventId"] = *request.ChallengePeriodEventId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateQuestGroupModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) UpdateQuestGroupModelMaster(
	request *UpdateQuestGroupModelMasterRequest,
) (*UpdateQuestGroupModelMasterResult, error) {
	callback := make(chan UpdateQuestGroupModelMasterAsyncResult, 1)
	go p.UpdateQuestGroupModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteQuestGroupModelMasterAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteQuestGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteQuestGroupModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteQuestGroupModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteQuestGroupModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteQuestGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DeleteQuestGroupModelMasterAsync(
	request *DeleteQuestGroupModelMasterRequest,
	callback chan<- DeleteQuestGroupModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{questGroupName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteQuestGroupModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DeleteQuestGroupModelMaster(
	request *DeleteQuestGroupModelMasterRequest,
) (*DeleteQuestGroupModelMasterResult, error) {
	callback := make(chan DeleteQuestGroupModelMasterAsyncResult, 1)
	go p.DeleteQuestGroupModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeQuestModelMastersAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeQuestModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeQuestModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeQuestModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeQuestModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeQuestModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DescribeQuestModelMastersAsync(
	request *DescribeQuestModelMastersRequest,
	callback chan<- DescribeQuestModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/group/{questGroupName}/quest"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
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

	go describeQuestModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DescribeQuestModelMasters(
	request *DescribeQuestModelMastersRequest,
) (*DescribeQuestModelMastersResult, error) {
	callback := make(chan DescribeQuestModelMastersAsyncResult, 1)
	go p.DescribeQuestModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createQuestModelMasterAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- CreateQuestModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateQuestModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateQuestModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateQuestModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateQuestModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) CreateQuestModelMasterAsync(
	request *CreateQuestModelMasterRequest,
	callback chan<- CreateQuestModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{questGroupName}/quest"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
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
	if request.Contents != nil {
		var _contents []*map[string]interface{}
		for _, item := range *request.Contents {
			_contents = append(_contents, item.ToDict())
		}
		bodies["contents"] = _contents
	}
	if request.ChallengePeriodEventId != nil && *request.ChallengePeriodEventId != "" {
		bodies["challengePeriodEventId"] = *request.ChallengePeriodEventId
	}
	if request.ConsumeActions != nil {
		var _consumeActions []*map[string]interface{}
		for _, item := range *request.ConsumeActions {
			_consumeActions = append(_consumeActions, item.ToDict())
		}
		bodies["consumeActions"] = _consumeActions
	}
	if request.FailedAcquireActions != nil {
		var _failedAcquireActions []*map[string]interface{}
		for _, item := range *request.FailedAcquireActions {
			_failedAcquireActions = append(_failedAcquireActions, item.ToDict())
		}
		bodies["failedAcquireActions"] = _failedAcquireActions
	}
	if request.PremiseQuestNames != nil {
		var _premiseQuestNames []string
		for _, item := range *request.PremiseQuestNames {
			_premiseQuestNames = append(_premiseQuestNames, item)
		}
		bodies["premiseQuestNames"] = _premiseQuestNames
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createQuestModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) CreateQuestModelMaster(
	request *CreateQuestModelMasterRequest,
) (*CreateQuestModelMasterResult, error) {
	callback := make(chan CreateQuestModelMasterAsyncResult, 1)
	go p.CreateQuestModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getQuestModelMasterAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- GetQuestModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetQuestModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetQuestModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetQuestModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetQuestModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) GetQuestModelMasterAsync(
	request *GetQuestModelMasterRequest,
	callback chan<- GetQuestModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{questGroupName}/quest/{questName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}
	if request.QuestName != nil {
		path = strings.ReplaceAll(path, "{questName}", core.ToString(*request.QuestName))
	} else {
		path = strings.ReplaceAll(path, "{questName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getQuestModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) GetQuestModelMaster(
	request *GetQuestModelMasterRequest,
) (*GetQuestModelMasterResult, error) {
	callback := make(chan GetQuestModelMasterAsyncResult, 1)
	go p.GetQuestModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateQuestModelMasterAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateQuestModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateQuestModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateQuestModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateQuestModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateQuestModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) UpdateQuestModelMasterAsync(
	request *UpdateQuestModelMasterRequest,
	callback chan<- UpdateQuestModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{questGroupName}/quest/{questName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}
	if request.QuestName != nil {
		path = strings.ReplaceAll(path, "{questName}", core.ToString(*request.QuestName))
	} else {
		path = strings.ReplaceAll(path, "{questName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Contents != nil {
		var _contents []*map[string]interface{}
		for _, item := range *request.Contents {
			_contents = append(_contents, item.ToDict())
		}
		bodies["contents"] = _contents
	}
	if request.ChallengePeriodEventId != nil && *request.ChallengePeriodEventId != "" {
		bodies["challengePeriodEventId"] = *request.ChallengePeriodEventId
	}
	if request.ConsumeActions != nil {
		var _consumeActions []*map[string]interface{}
		for _, item := range *request.ConsumeActions {
			_consumeActions = append(_consumeActions, item.ToDict())
		}
		bodies["consumeActions"] = _consumeActions
	}
	if request.FailedAcquireActions != nil {
		var _failedAcquireActions []*map[string]interface{}
		for _, item := range *request.FailedAcquireActions {
			_failedAcquireActions = append(_failedAcquireActions, item.ToDict())
		}
		bodies["failedAcquireActions"] = _failedAcquireActions
	}
	if request.PremiseQuestNames != nil {
		var _premiseQuestNames []string
		for _, item := range *request.PremiseQuestNames {
			_premiseQuestNames = append(_premiseQuestNames, item)
		}
		bodies["premiseQuestNames"] = _premiseQuestNames
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateQuestModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) UpdateQuestModelMaster(
	request *UpdateQuestModelMasterRequest,
) (*UpdateQuestModelMasterResult, error) {
	callback := make(chan UpdateQuestModelMasterAsyncResult, 1)
	go p.UpdateQuestModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteQuestModelMasterAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteQuestModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteQuestModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteQuestModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteQuestModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteQuestModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DeleteQuestModelMasterAsync(
	request *DeleteQuestModelMasterRequest,
	callback chan<- DeleteQuestModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{questGroupName}/quest/{questName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}
	if request.QuestName != nil {
		path = strings.ReplaceAll(path, "{questName}", core.ToString(*request.QuestName))
	} else {
		path = strings.ReplaceAll(path, "{questName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteQuestModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DeleteQuestModelMaster(
	request *DeleteQuestModelMasterRequest,
) (*DeleteQuestModelMasterResult, error) {
	callback := make(chan DeleteQuestModelMasterAsyncResult, 1)
	go p.DeleteQuestModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2QuestRestClient,
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

func (p Gs2QuestRestClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	path := "/{namespaceName}/master/export"
	if request.NamespaceName != nil {
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
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) ExportMaster(
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

func getCurrentQuestMasterAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentQuestMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentQuestMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentQuestMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentQuestMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentQuestMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) GetCurrentQuestMasterAsync(
	request *GetCurrentQuestMasterRequest,
	callback chan<- GetCurrentQuestMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil {
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

	go getCurrentQuestMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) GetCurrentQuestMaster(
	request *GetCurrentQuestMasterRequest,
) (*GetCurrentQuestMasterResult, error) {
	callback := make(chan GetCurrentQuestMasterAsyncResult, 1)
	go p.GetCurrentQuestMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentQuestMasterAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentQuestMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentQuestMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentQuestMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentQuestMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentQuestMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) UpdateCurrentQuestMasterAsync(
	request *UpdateCurrentQuestMasterRequest,
	callback chan<- UpdateCurrentQuestMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil {
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

	go updateCurrentQuestMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) UpdateCurrentQuestMaster(
	request *UpdateCurrentQuestMasterRequest,
) (*UpdateCurrentQuestMasterResult, error) {
	callback := make(chan UpdateCurrentQuestMasterAsyncResult, 1)
	go p.UpdateCurrentQuestMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentQuestMasterFromGitHubAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentQuestMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentQuestMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentQuestMasterFromGitHubResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentQuestMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentQuestMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) UpdateCurrentQuestMasterFromGitHubAsync(
	request *UpdateCurrentQuestMasterFromGitHubRequest,
	callback chan<- UpdateCurrentQuestMasterFromGitHubAsyncResult,
) {
	path := "/{namespaceName}/master/from_git_hub"
	if request.NamespaceName != nil {
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

	go updateCurrentQuestMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) UpdateCurrentQuestMasterFromGitHub(
	request *UpdateCurrentQuestMasterFromGitHubRequest,
) (*UpdateCurrentQuestMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentQuestMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentQuestMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeProgressesByUserIdAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeProgressesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeProgressesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeProgressesByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeProgressesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeProgressesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DescribeProgressesByUserIdAsync(
	request *DescribeProgressesByUserIdRequest,
	callback chan<- DescribeProgressesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/progress"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
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

	go describeProgressesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DescribeProgressesByUserId(
	request *DescribeProgressesByUserIdRequest,
) (*DescribeProgressesByUserIdResult, error) {
	callback := make(chan DescribeProgressesByUserIdAsyncResult, 1)
	go p.DescribeProgressesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createProgressByUserIdAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- CreateProgressByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateProgressByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateProgressByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateProgressByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateProgressByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) CreateProgressByUserIdAsync(
	request *CreateProgressByUserIdRequest,
	callback chan<- CreateProgressByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/progress"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.UserId != nil {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.QuestModelId != nil && *request.QuestModelId != "" {
		bodies["questModelId"] = *request.QuestModelId
	}
	if request.Force != nil {
		bodies["force"] = *request.Force
	}
	if request.Config != nil {
		var _config []*map[string]interface{}
		for _, item := range *request.Config {
			_config = append(_config, item.ToDict())
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createProgressByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) CreateProgressByUserId(
	request *CreateProgressByUserIdRequest,
) (*CreateProgressByUserIdResult, error) {
	callback := make(chan CreateProgressByUserIdAsyncResult, 1)
	go p.CreateProgressByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getProgressAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- GetProgressAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetProgressAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetProgressResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetProgressAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetProgressAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) GetProgressAsync(
	request *GetProgressRequest,
	callback chan<- GetProgressAsyncResult,
) {
	path := "/{namespaceName}/user/me/progress"
	if request.NamespaceName != nil {
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getProgressAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) GetProgress(
	request *GetProgressRequest,
) (*GetProgressResult, error) {
	callback := make(chan GetProgressAsyncResult, 1)
	go p.GetProgressAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getProgressByUserIdAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- GetProgressByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetProgressByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetProgressByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetProgressByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetProgressByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) GetProgressByUserIdAsync(
	request *GetProgressByUserIdRequest,
	callback chan<- GetProgressByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/progress"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.UserId != nil {
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

	go getProgressByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) GetProgressByUserId(
	request *GetProgressByUserIdRequest,
) (*GetProgressByUserIdResult, error) {
	callback := make(chan GetProgressByUserIdAsyncResult, 1)
	go p.GetProgressByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func startAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- StartAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- StartAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result StartResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- StartAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- StartAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) StartAsync(
	request *StartRequest,
	callback chan<- StartAsyncResult,
) {
	path := "/{namespaceName}/user/me/progress/group/{questGroupName}/quest/{questName}/start"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}
	if request.QuestName != nil {
		path = strings.ReplaceAll(path, "{questName}", core.ToString(*request.QuestName))
	} else {
		path = strings.ReplaceAll(path, "{questName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Force != nil {
		bodies["force"] = *request.Force
	}
	if request.Config != nil {
		var _config []*map[string]interface{}
		for _, item := range *request.Config {
			_config = append(_config, item.ToDict())
		}
		bodies["config"] = _config
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

	go startAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) Start(
	request *StartRequest,
) (*StartResult, error) {
	callback := make(chan StartAsyncResult, 1)
	go p.StartAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func startByUserIdAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- StartByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- StartByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result StartByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- StartByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- StartByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) StartByUserIdAsync(
	request *StartByUserIdRequest,
	callback chan<- StartByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/progress/group/{questGroupName}/quest/{questName}/start"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}
	if request.QuestName != nil {
		path = strings.ReplaceAll(path, "{questName}", core.ToString(*request.QuestName))
	} else {
		path = strings.ReplaceAll(path, "{questName}", "null")
	}
	if request.UserId != nil {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Force != nil {
		bodies["force"] = *request.Force
	}
	if request.Config != nil {
		var _config []*map[string]interface{}
		for _, item := range *request.Config {
			_config = append(_config, item.ToDict())
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go startByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) StartByUserId(
	request *StartByUserIdRequest,
) (*StartByUserIdResult, error) {
	callback := make(chan StartByUserIdAsyncResult, 1)
	go p.StartByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func endAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- EndAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- EndAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result EndResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- EndAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- EndAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) EndAsync(
	request *EndRequest,
	callback chan<- EndAsyncResult,
) {
	path := "/{namespaceName}/user/me/progress/end"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.TransactionId != nil && *request.TransactionId != "" {
		bodies["transactionId"] = *request.TransactionId
	}
	if request.Rewards != nil {
		var _rewards []*map[string]interface{}
		for _, item := range *request.Rewards {
			_rewards = append(_rewards, item.ToDict())
		}
		bodies["rewards"] = _rewards
	}
	if request.IsComplete != nil {
		bodies["isComplete"] = *request.IsComplete
	}
	if request.Config != nil {
		var _config []*map[string]interface{}
		for _, item := range *request.Config {
			_config = append(_config, item.ToDict())
		}
		bodies["config"] = _config
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

	go endAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) End(
	request *EndRequest,
) (*EndResult, error) {
	callback := make(chan EndAsyncResult, 1)
	go p.EndAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func endByUserIdAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- EndByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- EndByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result EndByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- EndByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- EndByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) EndByUserIdAsync(
	request *EndByUserIdRequest,
	callback chan<- EndByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/progress/end"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.UserId != nil {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.TransactionId != nil && *request.TransactionId != "" {
		bodies["transactionId"] = *request.TransactionId
	}
	if request.Rewards != nil {
		var _rewards []*map[string]interface{}
		for _, item := range *request.Rewards {
			_rewards = append(_rewards, item.ToDict())
		}
		bodies["rewards"] = _rewards
	}
	if request.IsComplete != nil {
		bodies["isComplete"] = *request.IsComplete
	}
	if request.Config != nil {
		var _config []*map[string]interface{}
		for _, item := range *request.Config {
			_config = append(_config, item.ToDict())
		}
		bodies["config"] = _config
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go endByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) EndByUserId(
	request *EndByUserIdRequest,
) (*EndByUserIdResult, error) {
	callback := make(chan EndByUserIdAsyncResult, 1)
	go p.EndByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteProgressAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteProgressAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteProgressAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteProgressResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteProgressAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteProgressAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DeleteProgressAsync(
	request *DeleteProgressRequest,
	callback chan<- DeleteProgressAsyncResult,
) {
	path := "/{namespaceName}/user/me/progress"
	if request.NamespaceName != nil {
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go deleteProgressAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DeleteProgress(
	request *DeleteProgressRequest,
) (*DeleteProgressResult, error) {
	callback := make(chan DeleteProgressAsyncResult, 1)
	go p.DeleteProgressAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteProgressByUserIdAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteProgressByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteProgressByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteProgressByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteProgressByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteProgressByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DeleteProgressByUserIdAsync(
	request *DeleteProgressByUserIdRequest,
	callback chan<- DeleteProgressByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/progress"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.UserId != nil {
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

	go deleteProgressByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DeleteProgressByUserId(
	request *DeleteProgressByUserIdRequest,
) (*DeleteProgressByUserIdResult, error) {
	callback := make(chan DeleteProgressByUserIdAsyncResult, 1)
	go p.DeleteProgressByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createProgressByStampSheetAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- CreateProgressByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateProgressByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateProgressByStampSheetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateProgressByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateProgressByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) CreateProgressByStampSheetAsync(
	request *CreateProgressByStampSheetRequest,
	callback chan<- CreateProgressByStampSheetAsyncResult,
) {
	path := "/stamp/progress/create"

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

	go createProgressByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) CreateProgressByStampSheet(
	request *CreateProgressByStampSheetRequest,
) (*CreateProgressByStampSheetResult, error) {
	callback := make(chan CreateProgressByStampSheetAsyncResult, 1)
	go p.CreateProgressByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteProgressByStampTaskAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteProgressByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteProgressByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteProgressByStampTaskResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteProgressByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteProgressByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DeleteProgressByStampTaskAsync(
	request *DeleteProgressByStampTaskRequest,
	callback chan<- DeleteProgressByStampTaskAsyncResult,
) {
	path := "/stamp/progress/delete"

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

	go deleteProgressByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DeleteProgressByStampTask(
	request *DeleteProgressByStampTaskRequest,
) (*DeleteProgressByStampTaskResult, error) {
	callback := make(chan DeleteProgressByStampTaskAsyncResult, 1)
	go p.DeleteProgressByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeCompletedQuestListsAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCompletedQuestListsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCompletedQuestListsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCompletedQuestListsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCompletedQuestListsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCompletedQuestListsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DescribeCompletedQuestListsAsync(
	request *DescribeCompletedQuestListsRequest,
	callback chan<- DescribeCompletedQuestListsAsyncResult,
) {
	path := "/{namespaceName}/user/me/completed"
	if request.NamespaceName != nil {
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

	go describeCompletedQuestListsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DescribeCompletedQuestLists(
	request *DescribeCompletedQuestListsRequest,
) (*DescribeCompletedQuestListsResult, error) {
	callback := make(chan DescribeCompletedQuestListsAsyncResult, 1)
	go p.DescribeCompletedQuestListsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeCompletedQuestListsByUserIdAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCompletedQuestListsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCompletedQuestListsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCompletedQuestListsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCompletedQuestListsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCompletedQuestListsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DescribeCompletedQuestListsByUserIdAsync(
	request *DescribeCompletedQuestListsByUserIdRequest,
	callback chan<- DescribeCompletedQuestListsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/completed"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.UserId != nil {
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

	go describeCompletedQuestListsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DescribeCompletedQuestListsByUserId(
	request *DescribeCompletedQuestListsByUserIdRequest,
) (*DescribeCompletedQuestListsByUserIdResult, error) {
	callback := make(chan DescribeCompletedQuestListsByUserIdAsyncResult, 1)
	go p.DescribeCompletedQuestListsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getCompletedQuestListAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- GetCompletedQuestListAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCompletedQuestListAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCompletedQuestListResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCompletedQuestListAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCompletedQuestListAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) GetCompletedQuestListAsync(
	request *GetCompletedQuestListRequest,
	callback chan<- GetCompletedQuestListAsyncResult,
) {
	path := "/{namespaceName}/user/me/completed/group/{questGroupName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
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

	go getCompletedQuestListAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) GetCompletedQuestList(
	request *GetCompletedQuestListRequest,
) (*GetCompletedQuestListResult, error) {
	callback := make(chan GetCompletedQuestListAsyncResult, 1)
	go p.GetCompletedQuestListAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getCompletedQuestListByUserIdAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- GetCompletedQuestListByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCompletedQuestListByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCompletedQuestListByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCompletedQuestListByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCompletedQuestListByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) GetCompletedQuestListByUserIdAsync(
	request *GetCompletedQuestListByUserIdRequest,
	callback chan<- GetCompletedQuestListByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/completed/group/{questGroupName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}
	if request.UserId != nil {
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

	go getCompletedQuestListByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) GetCompletedQuestListByUserId(
	request *GetCompletedQuestListByUserIdRequest,
) (*GetCompletedQuestListByUserIdResult, error) {
	callback := make(chan GetCompletedQuestListByUserIdAsyncResult, 1)
	go p.GetCompletedQuestListByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteCompletedQuestListByUserIdAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteCompletedQuestListByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteCompletedQuestListByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteCompletedQuestListByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteCompletedQuestListByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteCompletedQuestListByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DeleteCompletedQuestListByUserIdAsync(
	request *DeleteCompletedQuestListByUserIdRequest,
	callback chan<- DeleteCompletedQuestListByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/completed/group/{questGroupName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}
	if request.UserId != nil {
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

	go deleteCompletedQuestListByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DeleteCompletedQuestListByUserId(
	request *DeleteCompletedQuestListByUserIdRequest,
) (*DeleteCompletedQuestListByUserIdResult, error) {
	callback := make(chan DeleteCompletedQuestListByUserIdAsyncResult, 1)
	go p.DeleteCompletedQuestListByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeQuestGroupModelsAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeQuestGroupModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeQuestGroupModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeQuestGroupModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeQuestGroupModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeQuestGroupModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DescribeQuestGroupModelsAsync(
	request *DescribeQuestGroupModelsRequest,
	callback chan<- DescribeQuestGroupModelsAsyncResult,
) {
	path := "/{namespaceName}/group"
	if request.NamespaceName != nil {
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

	go describeQuestGroupModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DescribeQuestGroupModels(
	request *DescribeQuestGroupModelsRequest,
) (*DescribeQuestGroupModelsResult, error) {
	callback := make(chan DescribeQuestGroupModelsAsyncResult, 1)
	go p.DescribeQuestGroupModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getQuestGroupModelAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- GetQuestGroupModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetQuestGroupModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetQuestGroupModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetQuestGroupModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetQuestGroupModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) GetQuestGroupModelAsync(
	request *GetQuestGroupModelRequest,
	callback chan<- GetQuestGroupModelAsyncResult,
) {
	path := "/{namespaceName}/group/{questGroupName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getQuestGroupModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) GetQuestGroupModel(
	request *GetQuestGroupModelRequest,
) (*GetQuestGroupModelResult, error) {
	callback := make(chan GetQuestGroupModelAsyncResult, 1)
	go p.GetQuestGroupModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeQuestModelsAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeQuestModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeQuestModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeQuestModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeQuestModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeQuestModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) DescribeQuestModelsAsync(
	request *DescribeQuestModelsRequest,
	callback chan<- DescribeQuestModelsAsyncResult,
) {
	path := "/{namespaceName}/group/{questGroupName}/quest"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeQuestModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) DescribeQuestModels(
	request *DescribeQuestModelsRequest,
) (*DescribeQuestModelsResult, error) {
	callback := make(chan DescribeQuestModelsAsyncResult, 1)
	go p.DescribeQuestModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getQuestModelAsyncHandler(
	client Gs2QuestRestClient,
	job *core.NetworkJob,
	callback chan<- GetQuestModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetQuestModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetQuestModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetQuestModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetQuestModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestRestClient) GetQuestModelAsync(
	request *GetQuestModelRequest,
	callback chan<- GetQuestModelAsyncResult,
) {
	path := "/{namespaceName}/group/{questGroupName}/quest/{questName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}
	if request.QuestName != nil {
		path = strings.ReplaceAll(path, "{questName}", core.ToString(*request.QuestName))
	} else {
		path = strings.ReplaceAll(path, "{questName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getQuestModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("quest").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2QuestRestClient) GetQuestModel(
	request *GetQuestModelRequest,
) (*GetQuestModelResult, error) {
	callback := make(chan GetQuestModelAsyncResult, 1)
	go p.GetQuestModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
