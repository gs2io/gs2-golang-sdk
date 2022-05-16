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

package exchange

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

type Gs2ExchangeRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2ExchangeRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2ExchangeRestClient,
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

func (p Gs2ExchangeRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) DescribeNamespaces(
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
	client Gs2ExchangeRestClient,
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

func (p Gs2ExchangeRestClient) CreateNamespaceAsync(
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
    if request.EnableAwaitExchange != nil {
        bodies["enableAwaitExchange"] = *request.EnableAwaitExchange
    }
    if request.EnableDirectExchange != nil {
        bodies["enableDirectExchange"] = *request.EnableDirectExchange
    }
    if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
        bodies["queueNamespaceId"] = *request.QueueNamespaceId
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
    if request.ExchangeScript != nil {
        bodies["exchangeScript"] = request.ExchangeScript.ToDict()
    }
    if request.LogSetting != nil {
        bodies["logSetting"] = request.LogSetting.ToDict()
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) CreateNamespace(
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
	client Gs2ExchangeRestClient,
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

func (p Gs2ExchangeRestClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	path := "/{namespaceName}/status"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) GetNamespaceStatus(
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
	client Gs2ExchangeRestClient,
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

func (p Gs2ExchangeRestClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) GetNamespace(
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
	client Gs2ExchangeRestClient,
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

func (p Gs2ExchangeRestClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.EnableAwaitExchange != nil {
        bodies["enableAwaitExchange"] = *request.EnableAwaitExchange
    }
    if request.EnableDirectExchange != nil {
        bodies["enableDirectExchange"] = *request.EnableDirectExchange
    }
    if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
        bodies["queueNamespaceId"] = *request.QueueNamespaceId
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
    if request.ExchangeScript != nil {
        bodies["exchangeScript"] = request.ExchangeScript.ToDict()
    }
    if request.LogSetting != nil {
        bodies["logSetting"] = request.LogSetting.ToDict()
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) UpdateNamespace(
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
	client Gs2ExchangeRestClient,
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

func (p Gs2ExchangeRestClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) DeleteNamespace(
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

func describeRateModelsAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRateModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRateModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRateModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeRateModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRateModelsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeRateModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) DescribeRateModelsAsync(
	request *DescribeRateModelsRequest,
	callback chan<- DescribeRateModelsAsyncResult,
) {
	path := "/{namespaceName}/model"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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

	go describeRateModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) DescribeRateModels(
	request *DescribeRateModelsRequest,
) (*DescribeRateModelsResult, error) {
	callback := make(chan DescribeRateModelsAsyncResult, 1)
	go p.DescribeRateModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRateModelAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- GetRateModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRateModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRateModelResult
	if asyncResult.Err != nil {
		callback <- GetRateModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRateModelAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetRateModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) GetRateModelAsync(
	request *GetRateModelRequest,
	callback chan<- GetRateModelAsyncResult,
) {
	path := "/{namespaceName}/model/{rateName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getRateModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) GetRateModel(
	request *GetRateModelRequest,
) (*GetRateModelResult, error) {
	callback := make(chan GetRateModelAsyncResult, 1)
	go p.GetRateModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRateModelMastersAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRateModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRateModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRateModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeRateModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRateModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeRateModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) DescribeRateModelMastersAsync(
	request *DescribeRateModelMastersRequest,
	callback chan<- DescribeRateModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/model"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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

	go describeRateModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) DescribeRateModelMasters(
	request *DescribeRateModelMastersRequest,
) (*DescribeRateModelMastersResult, error) {
	callback := make(chan DescribeRateModelMastersAsyncResult, 1)
	go p.DescribeRateModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createRateModelMasterAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- CreateRateModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateRateModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateRateModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateRateModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateRateModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateRateModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) CreateRateModelMasterAsync(
	request *CreateRateModelMasterRequest,
	callback chan<- CreateRateModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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
    if request.TimingType != nil && *request.TimingType != "" {
        bodies["timingType"] = *request.TimingType
    }
    if request.LockTime != nil {
        bodies["lockTime"] = *request.LockTime
    }
    if request.EnableSkip != nil {
        bodies["enableSkip"] = *request.EnableSkip
    }
    if request.SkipConsumeActions != nil {
        var _skipConsumeActions []interface {}
        for _, item := range request.SkipConsumeActions {
            _skipConsumeActions = append(_skipConsumeActions, item)
        }
        bodies["skipConsumeActions"] = _skipConsumeActions
    }
    if request.AcquireActions != nil {
        var _acquireActions []interface {}
        for _, item := range request.AcquireActions {
            _acquireActions = append(_acquireActions, item)
        }
        bodies["acquireActions"] = _acquireActions
    }
    if request.ConsumeActions != nil {
        var _consumeActions []interface {}
        for _, item := range request.ConsumeActions {
            _consumeActions = append(_consumeActions, item)
        }
        bodies["consumeActions"] = _consumeActions
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createRateModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) CreateRateModelMaster(
	request *CreateRateModelMasterRequest,
) (*CreateRateModelMasterResult, error) {
	callback := make(chan CreateRateModelMasterAsyncResult, 1)
	go p.CreateRateModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRateModelMasterAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- GetRateModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRateModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRateModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetRateModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRateModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetRateModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) GetRateModelMasterAsync(
	request *GetRateModelMasterRequest,
	callback chan<- GetRateModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{rateName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getRateModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) GetRateModelMaster(
	request *GetRateModelMasterRequest,
) (*GetRateModelMasterResult, error) {
	callback := make(chan GetRateModelMasterAsyncResult, 1)
	go p.GetRateModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateRateModelMasterAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateRateModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateRateModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateRateModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateRateModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateRateModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateRateModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) UpdateRateModelMasterAsync(
	request *UpdateRateModelMasterRequest,
	callback chan<- UpdateRateModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{rateName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.TimingType != nil && *request.TimingType != "" {
        bodies["timingType"] = *request.TimingType
    }
    if request.LockTime != nil {
        bodies["lockTime"] = *request.LockTime
    }
    if request.EnableSkip != nil {
        bodies["enableSkip"] = *request.EnableSkip
    }
    if request.SkipConsumeActions != nil {
        var _skipConsumeActions []interface {}
        for _, item := range request.SkipConsumeActions {
            _skipConsumeActions = append(_skipConsumeActions, item)
        }
        bodies["skipConsumeActions"] = _skipConsumeActions
    }
    if request.AcquireActions != nil {
        var _acquireActions []interface {}
        for _, item := range request.AcquireActions {
            _acquireActions = append(_acquireActions, item)
        }
        bodies["acquireActions"] = _acquireActions
    }
    if request.ConsumeActions != nil {
        var _consumeActions []interface {}
        for _, item := range request.ConsumeActions {
            _consumeActions = append(_consumeActions, item)
        }
        bodies["consumeActions"] = _consumeActions
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateRateModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) UpdateRateModelMaster(
	request *UpdateRateModelMasterRequest,
) (*UpdateRateModelMasterResult, error) {
	callback := make(chan UpdateRateModelMasterAsyncResult, 1)
	go p.UpdateRateModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteRateModelMasterAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteRateModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRateModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRateModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteRateModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteRateModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteRateModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) DeleteRateModelMasterAsync(
	request *DeleteRateModelMasterRequest,
	callback chan<- DeleteRateModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{rateName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteRateModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) DeleteRateModelMaster(
	request *DeleteRateModelMasterRequest,
) (*DeleteRateModelMasterResult, error) {
	callback := make(chan DeleteRateModelMasterAsyncResult, 1)
	go p.DeleteRateModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exchangeAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- ExchangeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ExchangeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ExchangeResult
	if asyncResult.Err != nil {
		callback <- ExchangeAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ExchangeAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ExchangeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) ExchangeAsync(
	request *ExchangeRequest,
	callback chan<- ExchangeAsyncResult,
) {
	path := "/{namespaceName}/user/me/exchange/{rateName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Count != nil {
        bodies["count"] = *request.Count
    }
    if request.Config != nil {
        var _config []interface {}
        for _, item := range request.Config {
            _config = append(_config, item)
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go exchangeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) Exchange(
	request *ExchangeRequest,
) (*ExchangeResult, error) {
	callback := make(chan ExchangeAsyncResult, 1)
	go p.ExchangeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exchangeByUserIdAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- ExchangeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ExchangeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ExchangeByUserIdResult
	if asyncResult.Err != nil {
		callback <- ExchangeByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ExchangeByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ExchangeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) ExchangeByUserIdAsync(
	request *ExchangeByUserIdRequest,
	callback chan<- ExchangeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/exchange/{rateName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Count != nil {
        bodies["count"] = *request.Count
    }
    if request.Config != nil {
        var _config []interface {}
        for _, item := range request.Config {
            _config = append(_config, item)
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go exchangeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) ExchangeByUserId(
	request *ExchangeByUserIdRequest,
) (*ExchangeByUserIdResult, error) {
	callback := make(chan ExchangeByUserIdAsyncResult, 1)
	go p.ExchangeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exchangeByStampSheetAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- ExchangeByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ExchangeByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ExchangeByStampSheetResult
	if asyncResult.Err != nil {
		callback <- ExchangeByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ExchangeByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ExchangeByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) ExchangeByStampSheetAsync(
	request *ExchangeByStampSheetRequest,
	callback chan<- ExchangeByStampSheetAsyncResult,
) {
	path := "/stamp/exchange"

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.StampSheet != nil && *request.StampSheet != "" {
        bodies["stampSheet"] = *request.StampSheet
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go exchangeByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) ExchangeByStampSheet(
	request *ExchangeByStampSheetRequest,
) (*ExchangeByStampSheetResult, error) {
	callback := make(chan ExchangeByStampSheetAsyncResult, 1)
	go p.ExchangeByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2ExchangeRestClient,
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

func (p Gs2ExchangeRestClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	path := "/{namespaceName}/master/export"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) ExportMaster(
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

func getCurrentRateMasterAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentRateMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentRateMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentRateMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentRateMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentRateMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentRateMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) GetCurrentRateMasterAsync(
	request *GetCurrentRateMasterRequest,
	callback chan<- GetCurrentRateMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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

	go getCurrentRateMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) GetCurrentRateMaster(
	request *GetCurrentRateMasterRequest,
) (*GetCurrentRateMasterResult, error) {
	callback := make(chan GetCurrentRateMasterAsyncResult, 1)
	go p.GetCurrentRateMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentRateMasterAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentRateMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentRateMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentRateMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentRateMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentRateMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentRateMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) UpdateCurrentRateMasterAsync(
	request *UpdateCurrentRateMasterRequest,
	callback chan<- UpdateCurrentRateMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateCurrentRateMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) UpdateCurrentRateMaster(
	request *UpdateCurrentRateMasterRequest,
) (*UpdateCurrentRateMasterResult, error) {
	callback := make(chan UpdateCurrentRateMasterAsyncResult, 1)
	go p.UpdateCurrentRateMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentRateMasterFromGitHubAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentRateMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentRateMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentRateMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentRateMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentRateMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentRateMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) UpdateCurrentRateMasterFromGitHubAsync(
	request *UpdateCurrentRateMasterFromGitHubRequest,
	callback chan<- UpdateCurrentRateMasterFromGitHubAsyncResult,
) {
	path := "/{namespaceName}/master/from_git_hub"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateCurrentRateMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) UpdateCurrentRateMasterFromGitHub(
	request *UpdateCurrentRateMasterFromGitHubRequest,
) (*UpdateCurrentRateMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentRateMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentRateMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createAwaitByUserIdAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- CreateAwaitByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateAwaitByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateAwaitByUserIdResult
	if asyncResult.Err != nil {
		callback <- CreateAwaitByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateAwaitByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateAwaitByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) CreateAwaitByUserIdAsync(
	request *CreateAwaitByUserIdRequest,
	callback chan<- CreateAwaitByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/exchange/{rateName}/await"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Count != nil {
        bodies["count"] = *request.Count
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go createAwaitByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) CreateAwaitByUserId(
	request *CreateAwaitByUserIdRequest,
) (*CreateAwaitByUserIdResult, error) {
	callback := make(chan CreateAwaitByUserIdAsyncResult, 1)
	go p.CreateAwaitByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeAwaitsAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeAwaitsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeAwaitsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeAwaitsResult
	if asyncResult.Err != nil {
		callback <- DescribeAwaitsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeAwaitsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeAwaitsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) DescribeAwaitsAsync(
	request *DescribeAwaitsRequest,
	callback chan<- DescribeAwaitsAsyncResult,
) {
	path := "/{namespaceName}/user/me/exchange/await"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.RateName != nil {
		queryStrings["rateName"] = core.ToString(*request.RateName)
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

	go describeAwaitsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) DescribeAwaits(
	request *DescribeAwaitsRequest,
) (*DescribeAwaitsResult, error) {
	callback := make(chan DescribeAwaitsAsyncResult, 1)
	go p.DescribeAwaitsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeAwaitsByUserIdAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeAwaitsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeAwaitsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeAwaitsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeAwaitsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeAwaitsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeAwaitsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) DescribeAwaitsByUserIdAsync(
	request *DescribeAwaitsByUserIdRequest,
	callback chan<- DescribeAwaitsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/exchange/await"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.RateName != nil {
		queryStrings["rateName"] = core.ToString(*request.RateName)
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

	go describeAwaitsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) DescribeAwaitsByUserId(
	request *DescribeAwaitsByUserIdRequest,
) (*DescribeAwaitsByUserIdResult, error) {
	callback := make(chan DescribeAwaitsByUserIdAsyncResult, 1)
	go p.DescribeAwaitsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getAwaitAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- GetAwaitAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetAwaitAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetAwaitResult
	if asyncResult.Err != nil {
		callback <- GetAwaitAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetAwaitAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetAwaitAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) GetAwaitAsync(
	request *GetAwaitRequest,
	callback chan<- GetAwaitAsyncResult,
) {
	path := "/{namespaceName}/user/me/exchange/{rateName}/await/{awaitName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }
    if request.AwaitName != nil && *request.AwaitName != ""  {
        path = strings.ReplaceAll(path, "{awaitName}", core.ToString(*request.AwaitName))
    } else {
        path = strings.ReplaceAll(path, "{awaitName}", "null")
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

	go getAwaitAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) GetAwait(
	request *GetAwaitRequest,
) (*GetAwaitResult, error) {
	callback := make(chan GetAwaitAsyncResult, 1)
	go p.GetAwaitAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getAwaitByUserIdAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- GetAwaitByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetAwaitByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetAwaitByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetAwaitByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetAwaitByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetAwaitByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) GetAwaitByUserIdAsync(
	request *GetAwaitByUserIdRequest,
	callback chan<- GetAwaitByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/exchange/{rateName}/await/{awaitName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }
    if request.AwaitName != nil && *request.AwaitName != ""  {
        path = strings.ReplaceAll(path, "{awaitName}", core.ToString(*request.AwaitName))
    } else {
        path = strings.ReplaceAll(path, "{awaitName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getAwaitByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) GetAwaitByUserId(
	request *GetAwaitByUserIdRequest,
) (*GetAwaitByUserIdResult, error) {
	callback := make(chan GetAwaitByUserIdAsyncResult, 1)
	go p.GetAwaitByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireResult
	if asyncResult.Err != nil {
		callback <- AcquireAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AcquireAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AcquireAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) AcquireAsync(
	request *AcquireRequest,
	callback chan<- AcquireAsyncResult,
) {
	path := "/{namespaceName}/user/me/exchange/{rateName}/await/{awaitName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }
    if request.AwaitName != nil && *request.AwaitName != ""  {
        path = strings.ReplaceAll(path, "{awaitName}", core.ToString(*request.AwaitName))
    } else {
        path = strings.ReplaceAll(path, "{awaitName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Config != nil {
        var _config []interface {}
        for _, item := range request.Config {
            _config = append(_config, item)
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go acquireAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) Acquire(
	request *AcquireRequest,
) (*AcquireResult, error) {
	callback := make(chan AcquireAsyncResult, 1)
	go p.AcquireAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireByUserIdAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireByUserIdResult
	if asyncResult.Err != nil {
		callback <- AcquireByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AcquireByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AcquireByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) AcquireByUserIdAsync(
	request *AcquireByUserIdRequest,
	callback chan<- AcquireByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/exchange/{rateName}/await/{awaitName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }
    if request.AwaitName != nil && *request.AwaitName != ""  {
        path = strings.ReplaceAll(path, "{awaitName}", core.ToString(*request.AwaitName))
    } else {
        path = strings.ReplaceAll(path, "{awaitName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Config != nil {
        var _config []interface {}
        for _, item := range request.Config {
            _config = append(_config, item)
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go acquireByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) AcquireByUserId(
	request *AcquireByUserIdRequest,
) (*AcquireByUserIdResult, error) {
	callback := make(chan AcquireByUserIdAsyncResult, 1)
	go p.AcquireByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireForceByUserIdAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireForceByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireForceByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireForceByUserIdResult
	if asyncResult.Err != nil {
		callback <- AcquireForceByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AcquireForceByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AcquireForceByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) AcquireForceByUserIdAsync(
	request *AcquireForceByUserIdRequest,
	callback chan<- AcquireForceByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/exchange/{rateName}/await/{awaitName}/force"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }
    if request.AwaitName != nil && *request.AwaitName != ""  {
        path = strings.ReplaceAll(path, "{awaitName}", core.ToString(*request.AwaitName))
    } else {
        path = strings.ReplaceAll(path, "{awaitName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Config != nil {
        var _config []interface {}
        for _, item := range request.Config {
            _config = append(_config, item)
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go acquireForceByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) AcquireForceByUserId(
	request *AcquireForceByUserIdRequest,
) (*AcquireForceByUserIdResult, error) {
	callback := make(chan AcquireForceByUserIdAsyncResult, 1)
	go p.AcquireForceByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func skipAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- SkipAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SkipAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SkipResult
	if asyncResult.Err != nil {
		callback <- SkipAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SkipAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SkipAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) SkipAsync(
	request *SkipRequest,
	callback chan<- SkipAsyncResult,
) {
	path := "/{namespaceName}/user/me/exchange/{rateName}/await/{awaitName}/skip"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }
    if request.AwaitName != nil && *request.AwaitName != ""  {
        path = strings.ReplaceAll(path, "{awaitName}", core.ToString(*request.AwaitName))
    } else {
        path = strings.ReplaceAll(path, "{awaitName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Config != nil {
        var _config []interface {}
        for _, item := range request.Config {
            _config = append(_config, item)
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go skipAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) Skip(
	request *SkipRequest,
) (*SkipResult, error) {
	callback := make(chan SkipAsyncResult, 1)
	go p.SkipAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func skipByUserIdAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- SkipByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SkipByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SkipByUserIdResult
	if asyncResult.Err != nil {
		callback <- SkipByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SkipByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SkipByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) SkipByUserIdAsync(
	request *SkipByUserIdRequest,
	callback chan<- SkipByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/exchange/{rateName}/await/{awaitName}/skip"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }
    if request.AwaitName != nil && *request.AwaitName != ""  {
        path = strings.ReplaceAll(path, "{awaitName}", core.ToString(*request.AwaitName))
    } else {
        path = strings.ReplaceAll(path, "{awaitName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Config != nil {
        var _config []interface {}
        for _, item := range request.Config {
            _config = append(_config, item)
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go skipByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) SkipByUserId(
	request *SkipByUserIdRequest,
) (*SkipByUserIdResult, error) {
	callback := make(chan SkipByUserIdAsyncResult, 1)
	go p.SkipByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteAwaitAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteAwaitAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteAwaitAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteAwaitResult
	if asyncResult.Err != nil {
		callback <- DeleteAwaitAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteAwaitAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteAwaitAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) DeleteAwaitAsync(
	request *DeleteAwaitRequest,
	callback chan<- DeleteAwaitAsyncResult,
) {
	path := "/{namespaceName}/user/me/exchange/{rateName}/await/{awaitName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }
    if request.AwaitName != nil && *request.AwaitName != ""  {
        path = strings.ReplaceAll(path, "{awaitName}", core.ToString(*request.AwaitName))
    } else {
        path = strings.ReplaceAll(path, "{awaitName}", "null")
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

	go deleteAwaitAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) DeleteAwait(
	request *DeleteAwaitRequest,
) (*DeleteAwaitResult, error) {
	callback := make(chan DeleteAwaitAsyncResult, 1)
	go p.DeleteAwaitAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteAwaitByUserIdAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteAwaitByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteAwaitByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteAwaitByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteAwaitByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteAwaitByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteAwaitByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) DeleteAwaitByUserIdAsync(
	request *DeleteAwaitByUserIdRequest,
	callback chan<- DeleteAwaitByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/exchange/{rateName}/await/{awaitName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.RateName != nil && *request.RateName != ""  {
        path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
    } else {
        path = strings.ReplaceAll(path, "{rateName}", "null")
    }
    if request.AwaitName != nil && *request.AwaitName != ""  {
        path = strings.ReplaceAll(path, "{awaitName}", core.ToString(*request.AwaitName))
    } else {
        path = strings.ReplaceAll(path, "{awaitName}", "null")
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

	go deleteAwaitByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) DeleteAwaitByUserId(
	request *DeleteAwaitByUserIdRequest,
) (*DeleteAwaitByUserIdResult, error) {
	callback := make(chan DeleteAwaitByUserIdAsyncResult, 1)
	go p.DeleteAwaitByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createAwaitByStampSheetAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- CreateAwaitByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateAwaitByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateAwaitByStampSheetResult
	if asyncResult.Err != nil {
		callback <- CreateAwaitByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateAwaitByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateAwaitByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) CreateAwaitByStampSheetAsync(
	request *CreateAwaitByStampSheetRequest,
	callback chan<- CreateAwaitByStampSheetAsyncResult,
) {
	path := "/stamp/await/create"

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.StampSheet != nil && *request.StampSheet != "" {
        bodies["stampSheet"] = *request.StampSheet
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createAwaitByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) CreateAwaitByStampSheet(
	request *CreateAwaitByStampSheetRequest,
) (*CreateAwaitByStampSheetResult, error) {
	callback := make(chan CreateAwaitByStampSheetAsyncResult, 1)
	go p.CreateAwaitByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteAwaitByStampTaskAsyncHandler(
	client Gs2ExchangeRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteAwaitByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteAwaitByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteAwaitByStampTaskResult
	if asyncResult.Err != nil {
		callback <- DeleteAwaitByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteAwaitByStampTaskAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteAwaitByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExchangeRestClient) DeleteAwaitByStampTaskAsync(
	request *DeleteAwaitByStampTaskRequest,
	callback chan<- DeleteAwaitByStampTaskAsyncResult,
) {
	path := "/stamp/await/delete"

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.StampTask != nil && *request.StampTask != "" {
        bodies["stampTask"] = *request.StampTask
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteAwaitByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("exchange").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeRestClient) DeleteAwaitByStampTask(
	request *DeleteAwaitByStampTaskRequest,
) (*DeleteAwaitByStampTaskResult, error) {
	callback := make(chan DeleteAwaitByStampTaskAsyncResult, 1)
	go p.DeleteAwaitByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
