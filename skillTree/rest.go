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

package skillTree

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2SkillTreeRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2SkillTreeRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2SkillTreeRestClient,
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

func (p Gs2SkillTreeRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) DescribeNamespaces(
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
	client Gs2SkillTreeRestClient,
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

func (p Gs2SkillTreeRestClient) CreateNamespaceAsync(
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
	if request.ReleaseScript != nil {
		bodies["releaseScript"] = request.ReleaseScript.ToDict()
	}
	if request.RestrainScript != nil {
		bodies["restrainScript"] = request.RestrainScript.ToDict()
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
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) CreateNamespace(
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
	client Gs2SkillTreeRestClient,
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

func (p Gs2SkillTreeRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) GetNamespaceStatus(
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
	client Gs2SkillTreeRestClient,
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

func (p Gs2SkillTreeRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) GetNamespace(
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
	client Gs2SkillTreeRestClient,
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

func (p Gs2SkillTreeRestClient) UpdateNamespaceAsync(
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
	if request.ReleaseScript != nil {
		bodies["releaseScript"] = request.ReleaseScript.ToDict()
	}
	if request.RestrainScript != nil {
		bodies["restrainScript"] = request.RestrainScript.ToDict()
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
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) UpdateNamespace(
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
	client Gs2SkillTreeRestClient,
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

func (p Gs2SkillTreeRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) DeleteNamespace(
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

func describeStatusModelsAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStatusModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStatusModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStatusModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeStatusModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStatusModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStatusModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) DescribeStatusModelsAsync(
	request *DescribeStatusModelsRequest,
	callback chan<- DescribeStatusModelsAsyncResult,
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

	go describeStatusModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) DescribeStatusModels(
	request *DescribeStatusModelsRequest,
) (*DescribeStatusModelsResult, error) {
	callback := make(chan DescribeStatusModelsAsyncResult, 1)
	go p.DescribeStatusModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStatusModelAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- GetStatusModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusModelResult
	if asyncResult.Err != nil {
		callback <- GetStatusModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStatusModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStatusModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) GetStatusModelAsync(
	request *GetStatusModelRequest,
	callback chan<- GetStatusModelAsyncResult,
) {
	path := "/{namespaceName}/model/{nodeModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.NodeModelName != nil && *request.NodeModelName != "" {
		path = strings.ReplaceAll(path, "{nodeModelName}", core.ToString(*request.NodeModelName))
	} else {
		path = strings.ReplaceAll(path, "{nodeModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getStatusModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) GetStatusModel(
	request *GetStatusModelRequest,
) (*GetStatusModelResult, error) {
	callback := make(chan GetStatusModelAsyncResult, 1)
	go p.GetStatusModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeStatusModelMastersAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStatusModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStatusModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStatusModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeStatusModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStatusModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStatusModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) DescribeStatusModelMastersAsync(
	request *DescribeStatusModelMastersRequest,
	callback chan<- DescribeStatusModelMastersAsyncResult,
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

	go describeStatusModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) DescribeStatusModelMasters(
	request *DescribeStatusModelMastersRequest,
) (*DescribeStatusModelMastersResult, error) {
	callback := make(chan DescribeStatusModelMastersAsyncResult, 1)
	go p.DescribeStatusModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createStatusModelMasterAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- CreateStatusModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateStatusModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateStatusModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateStatusModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateStatusModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateStatusModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) CreateStatusModelMasterAsync(
	request *CreateStatusModelMasterRequest,
	callback chan<- CreateStatusModelMasterAsyncResult,
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
	if request.ReleaseConsumeActions != nil {
		var _releaseConsumeActions []interface{}
		for _, item := range request.ReleaseConsumeActions {
			_releaseConsumeActions = append(_releaseConsumeActions, item)
		}
		bodies["releaseConsumeActions"] = _releaseConsumeActions
	}
	if request.RestrainReturnRate != nil {
		bodies["restrainReturnRate"] = *request.RestrainReturnRate
	}
	if request.PremiseNodeNames != nil {
		var _premiseNodeNames []interface{}
		for _, item := range request.PremiseNodeNames {
			_premiseNodeNames = append(_premiseNodeNames, item)
		}
		bodies["premiseNodeNames"] = _premiseNodeNames
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createStatusModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) CreateStatusModelMaster(
	request *CreateStatusModelMasterRequest,
) (*CreateStatusModelMasterResult, error) {
	callback := make(chan CreateStatusModelMasterAsyncResult, 1)
	go p.CreateStatusModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStatusModelMasterAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- GetStatusModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetStatusModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStatusModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStatusModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) GetStatusModelMasterAsync(
	request *GetStatusModelMasterRequest,
	callback chan<- GetStatusModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{nodeModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.NodeModelName != nil && *request.NodeModelName != "" {
		path = strings.ReplaceAll(path, "{nodeModelName}", core.ToString(*request.NodeModelName))
	} else {
		path = strings.ReplaceAll(path, "{nodeModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getStatusModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) GetStatusModelMaster(
	request *GetStatusModelMasterRequest,
) (*GetStatusModelMasterResult, error) {
	callback := make(chan GetStatusModelMasterAsyncResult, 1)
	go p.GetStatusModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateStatusModelMasterAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateStatusModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateStatusModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateStatusModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateStatusModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateStatusModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateStatusModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) UpdateStatusModelMasterAsync(
	request *UpdateStatusModelMasterRequest,
	callback chan<- UpdateStatusModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{nodeModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.NodeModelName != nil && *request.NodeModelName != "" {
		path = strings.ReplaceAll(path, "{nodeModelName}", core.ToString(*request.NodeModelName))
	} else {
		path = strings.ReplaceAll(path, "{nodeModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ReleaseConsumeActions != nil {
		var _releaseConsumeActions []interface{}
		for _, item := range request.ReleaseConsumeActions {
			_releaseConsumeActions = append(_releaseConsumeActions, item)
		}
		bodies["releaseConsumeActions"] = _releaseConsumeActions
	}
	if request.RestrainReturnRate != nil {
		bodies["restrainReturnRate"] = *request.RestrainReturnRate
	}
	if request.PremiseNodeNames != nil {
		var _premiseNodeNames []interface{}
		for _, item := range request.PremiseNodeNames {
			_premiseNodeNames = append(_premiseNodeNames, item)
		}
		bodies["premiseNodeNames"] = _premiseNodeNames
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateStatusModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) UpdateStatusModelMaster(
	request *UpdateStatusModelMasterRequest,
) (*UpdateStatusModelMasterResult, error) {
	callback := make(chan UpdateStatusModelMasterAsyncResult, 1)
	go p.UpdateStatusModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteStatusModelMasterAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteStatusModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStatusModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStatusModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteStatusModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStatusModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteStatusModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) DeleteStatusModelMasterAsync(
	request *DeleteStatusModelMasterRequest,
	callback chan<- DeleteStatusModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{nodeModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.NodeModelName != nil && *request.NodeModelName != "" {
		path = strings.ReplaceAll(path, "{nodeModelName}", core.ToString(*request.NodeModelName))
	} else {
		path = strings.ReplaceAll(path, "{nodeModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteStatusModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) DeleteStatusModelMaster(
	request *DeleteStatusModelMasterRequest,
) (*DeleteStatusModelMasterResult, error) {
	callback := make(chan DeleteStatusModelMasterAsyncResult, 1)
	go p.DeleteStatusModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func markReleaseByUserIdAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- MarkReleaseByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MarkReleaseByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MarkReleaseByUserIdResult
	if asyncResult.Err != nil {
		callback <- MarkReleaseByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MarkReleaseByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- MarkReleaseByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) MarkReleaseByUserIdAsync(
	request *MarkReleaseByUserIdRequest,
	callback chan<- MarkReleaseByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/node/release/mark"
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
	var bodies = core.Bodies{}
	if request.NodeModelNames != nil {
		var _nodeModelNames []interface{}
		for _, item := range request.NodeModelNames {
			_nodeModelNames = append(_nodeModelNames, item)
		}
		bodies["nodeModelNames"] = _nodeModelNames
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

	go markReleaseByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) MarkReleaseByUserId(
	request *MarkReleaseByUserIdRequest,
) (*MarkReleaseByUserIdResult, error) {
	callback := make(chan MarkReleaseByUserIdAsyncResult, 1)
	go p.MarkReleaseByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func releaseAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- ReleaseAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReleaseAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReleaseResult
	if asyncResult.Err != nil {
		callback <- ReleaseAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReleaseAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReleaseAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) ReleaseAsync(
	request *ReleaseRequest,
	callback chan<- ReleaseAsyncResult,
) {
	path := "/{namespaceName}/user/me/status/node/release"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.NodeModelNames != nil {
		var _nodeModelNames []interface{}
		for _, item := range request.NodeModelNames {
			_nodeModelNames = append(_nodeModelNames, item)
		}
		bodies["nodeModelNames"] = _nodeModelNames
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go releaseAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) Release(
	request *ReleaseRequest,
) (*ReleaseResult, error) {
	callback := make(chan ReleaseAsyncResult, 1)
	go p.ReleaseAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func releaseByUserIdAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- ReleaseByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReleaseByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReleaseByUserIdResult
	if asyncResult.Err != nil {
		callback <- ReleaseByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReleaseByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReleaseByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) ReleaseByUserIdAsync(
	request *ReleaseByUserIdRequest,
	callback chan<- ReleaseByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/node/release"
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
	var bodies = core.Bodies{}
	if request.NodeModelNames != nil {
		var _nodeModelNames []interface{}
		for _, item := range request.NodeModelNames {
			_nodeModelNames = append(_nodeModelNames, item)
		}
		bodies["nodeModelNames"] = _nodeModelNames
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go releaseByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) ReleaseByUserId(
	request *ReleaseByUserIdRequest,
) (*ReleaseByUserIdResult, error) {
	callback := make(chan ReleaseByUserIdAsyncResult, 1)
	go p.ReleaseByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func markRestrainByUserIdAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- MarkRestrainByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MarkRestrainByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MarkRestrainByUserIdResult
	if asyncResult.Err != nil {
		callback <- MarkRestrainByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MarkRestrainByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- MarkRestrainByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) MarkRestrainByUserIdAsync(
	request *MarkRestrainByUserIdRequest,
	callback chan<- MarkRestrainByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/node/restrain/mark"
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
	var bodies = core.Bodies{}
	if request.NodeModelNames != nil {
		var _nodeModelNames []interface{}
		for _, item := range request.NodeModelNames {
			_nodeModelNames = append(_nodeModelNames, item)
		}
		bodies["nodeModelNames"] = _nodeModelNames
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

	go markRestrainByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) MarkRestrainByUserId(
	request *MarkRestrainByUserIdRequest,
) (*MarkRestrainByUserIdResult, error) {
	callback := make(chan MarkRestrainByUserIdAsyncResult, 1)
	go p.MarkRestrainByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func restrainAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- RestrainAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RestrainAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RestrainResult
	if asyncResult.Err != nil {
		callback <- RestrainAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RestrainAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RestrainAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) RestrainAsync(
	request *RestrainRequest,
	callback chan<- RestrainAsyncResult,
) {
	path := "/{namespaceName}/user/me/status/node/restrain"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.NodeModelNames != nil {
		var _nodeModelNames []interface{}
		for _, item := range request.NodeModelNames {
			_nodeModelNames = append(_nodeModelNames, item)
		}
		bodies["nodeModelNames"] = _nodeModelNames
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go restrainAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) Restrain(
	request *RestrainRequest,
) (*RestrainResult, error) {
	callback := make(chan RestrainAsyncResult, 1)
	go p.RestrainAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func restrainByUserIdAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- RestrainByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RestrainByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RestrainByUserIdResult
	if asyncResult.Err != nil {
		callback <- RestrainByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RestrainByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RestrainByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) RestrainByUserIdAsync(
	request *RestrainByUserIdRequest,
	callback chan<- RestrainByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/node/restrain"
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
	var bodies = core.Bodies{}
	if request.NodeModelNames != nil {
		var _nodeModelNames []interface{}
		for _, item := range request.NodeModelNames {
			_nodeModelNames = append(_nodeModelNames, item)
		}
		bodies["nodeModelNames"] = _nodeModelNames
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go restrainByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) RestrainByUserId(
	request *RestrainByUserIdRequest,
) (*RestrainByUserIdResult, error) {
	callback := make(chan RestrainByUserIdAsyncResult, 1)
	go p.RestrainByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStatusAsyncHandler(
	client Gs2SkillTreeRestClient,
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

func (p Gs2SkillTreeRestClient) GetStatusAsync(
	request *GetStatusRequest,
	callback chan<- GetStatusAsyncResult,
) {
	path := "/{namespaceName}/user/me/status"
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) GetStatus(
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
	client Gs2SkillTreeRestClient,
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

func (p Gs2SkillTreeRestClient) GetStatusByUserIdAsync(
	request *GetStatusByUserIdRequest,
	callback chan<- GetStatusByUserIdAsyncResult,
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) GetStatusByUserId(
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

func resetAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- ResetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetResult
	if asyncResult.Err != nil {
		callback <- ResetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ResetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) ResetAsync(
	request *ResetRequest,
	callback chan<- ResetAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/reset"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go resetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) Reset(
	request *ResetRequest,
) (*ResetResult, error) {
	callback := make(chan ResetAsyncResult, 1)
	go p.ResetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func resetByUserIdAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- ResetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetByUserIdResult
	if asyncResult.Err != nil {
		callback <- ResetByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ResetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) ResetByUserIdAsync(
	request *ResetByUserIdRequest,
	callback chan<- ResetByUserIdAsyncResult,
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go resetByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) ResetByUserId(
	request *ResetByUserIdRequest,
) (*ResetByUserIdResult, error) {
	callback := make(chan ResetByUserIdAsyncResult, 1)
	go p.ResetByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func markReleaseByStampSheetAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- MarkReleaseByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MarkReleaseByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MarkReleaseByStampSheetResult
	if asyncResult.Err != nil {
		callback <- MarkReleaseByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MarkReleaseByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- MarkReleaseByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) MarkReleaseByStampSheetAsync(
	request *MarkReleaseByStampSheetRequest,
	callback chan<- MarkReleaseByStampSheetAsyncResult,
) {
	path := "/stamp/status/release"

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

	go markReleaseByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) MarkReleaseByStampSheet(
	request *MarkReleaseByStampSheetRequest,
) (*MarkReleaseByStampSheetResult, error) {
	callback := make(chan MarkReleaseByStampSheetAsyncResult, 1)
	go p.MarkReleaseByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func markRestrainByStampTaskAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- MarkRestrainByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MarkRestrainByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MarkRestrainByStampTaskResult
	if asyncResult.Err != nil {
		callback <- MarkRestrainByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MarkRestrainByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- MarkRestrainByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) MarkRestrainByStampTaskAsync(
	request *MarkRestrainByStampTaskRequest,
	callback chan<- MarkRestrainByStampTaskAsyncResult,
) {
	path := "/stamp/status/restrain"

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

	go markRestrainByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) MarkRestrainByStampTask(
	request *MarkRestrainByStampTaskRequest,
) (*MarkRestrainByStampTaskResult, error) {
	callback := make(chan MarkRestrainByStampTaskAsyncResult, 1)
	go p.MarkRestrainByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2SkillTreeRestClient,
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

func (p Gs2SkillTreeRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) ExportMaster(
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

func getCurrentStatusMasterAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentStatusMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentStatusMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentStatusMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentStatusMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentStatusMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentStatusMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) GetCurrentStatusMasterAsync(
	request *GetCurrentStatusMasterRequest,
	callback chan<- GetCurrentStatusMasterAsyncResult,
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

	go getCurrentStatusMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) GetCurrentStatusMaster(
	request *GetCurrentStatusMasterRequest,
) (*GetCurrentStatusMasterResult, error) {
	callback := make(chan GetCurrentStatusMasterAsyncResult, 1)
	go p.GetCurrentStatusMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentStatusMasterAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentStatusMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentStatusMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentStatusMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentStatusMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentStatusMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentStatusMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) UpdateCurrentStatusMasterAsync(
	request *UpdateCurrentStatusMasterRequest,
	callback chan<- UpdateCurrentStatusMasterAsyncResult,
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

	go updateCurrentStatusMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) UpdateCurrentStatusMaster(
	request *UpdateCurrentStatusMasterRequest,
) (*UpdateCurrentStatusMasterResult, error) {
	callback := make(chan UpdateCurrentStatusMasterAsyncResult, 1)
	go p.UpdateCurrentStatusMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentStatusMasterFromGitHubAsyncHandler(
	client Gs2SkillTreeRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentStatusMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentStatusMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentStatusMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentStatusMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentStatusMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentStatusMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SkillTreeRestClient) UpdateCurrentStatusMasterFromGitHubAsync(
	request *UpdateCurrentStatusMasterFromGitHubRequest,
	callback chan<- UpdateCurrentStatusMasterFromGitHubAsyncResult,
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

	go updateCurrentStatusMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("skill-tree").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2SkillTreeRestClient) UpdateCurrentStatusMasterFromGitHub(
	request *UpdateCurrentStatusMasterFromGitHubRequest,
) (*UpdateCurrentStatusMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentStatusMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentStatusMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
