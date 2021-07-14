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

package lottery

import (
	"encoding/json"
	"core"
	"strings"
)

type Gs2LotteryRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2LotteryRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeNamespaces(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) CreateNamespaceAsync(
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
    if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
        bodies["queueNamespaceId"] = *request.QueueNamespaceId
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
    if request.LotteryTriggerScriptId != nil && *request.LotteryTriggerScriptId != "" {
        bodies["lotteryTriggerScriptId"] = *request.LotteryTriggerScriptId
    }
    if request.ChoicePrizeTableScriptId != nil && *request.ChoicePrizeTableScriptId != "" {
        bodies["choicePrizeTableScriptId"] = *request.ChoicePrizeTableScriptId
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) CreateNamespace(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetNamespaceStatus(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetNamespace(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) UpdateNamespaceAsync(
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
    if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
        bodies["queueNamespaceId"] = *request.QueueNamespaceId
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
    if request.LotteryTriggerScriptId != nil && *request.LotteryTriggerScriptId != "" {
        bodies["lotteryTriggerScriptId"] = *request.LotteryTriggerScriptId
    }
    if request.ChoicePrizeTableScriptId != nil && *request.ChoicePrizeTableScriptId != "" {
        bodies["choicePrizeTableScriptId"] = *request.ChoicePrizeTableScriptId
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) UpdateNamespace(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DeleteNamespace(
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

func describeLotteryModelMastersAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLotteryModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLotteryModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLotteryModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeLotteryModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeLotteryModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeLotteryModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribeLotteryModelMastersAsync(
	request *DescribeLotteryModelMastersRequest,
	callback chan<- DescribeLotteryModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/lottery"
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

	go describeLotteryModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeLotteryModelMasters(
	request *DescribeLotteryModelMastersRequest,
) (*DescribeLotteryModelMastersResult, error) {
	callback := make(chan DescribeLotteryModelMastersAsyncResult, 1)
	go p.DescribeLotteryModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createLotteryModelMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- CreateLotteryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateLotteryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateLotteryModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateLotteryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateLotteryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateLotteryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) CreateLotteryModelMasterAsync(
	request *CreateLotteryModelMasterRequest,
	callback chan<- CreateLotteryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/lottery"
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
    if request.Mode != nil && *request.Mode != "" {
        bodies["mode"] = *request.Mode
    }
    if request.Method != nil && *request.Method != "" {
        bodies["method"] = *request.Method
    }
    if request.PrizeTableName != nil && *request.PrizeTableName != "" {
        bodies["prizeTableName"] = *request.PrizeTableName
    }
    if request.ChoicePrizeTableScriptId != nil && *request.ChoicePrizeTableScriptId != "" {
        bodies["choicePrizeTableScriptId"] = *request.ChoicePrizeTableScriptId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createLotteryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) CreateLotteryModelMaster(
	request *CreateLotteryModelMasterRequest,
) (*CreateLotteryModelMasterResult, error) {
	callback := make(chan CreateLotteryModelMasterAsyncResult, 1)
	go p.CreateLotteryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getLotteryModelMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetLotteryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLotteryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLotteryModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetLotteryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetLotteryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetLotteryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetLotteryModelMasterAsync(
	request *GetLotteryModelMasterRequest,
	callback chan<- GetLotteryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/lottery/{lotteryName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.LotteryName != nil && *request.LotteryName != ""  {
        path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
    } else {
        path = strings.ReplaceAll(path, "{lotteryName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getLotteryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetLotteryModelMaster(
	request *GetLotteryModelMasterRequest,
) (*GetLotteryModelMasterResult, error) {
	callback := make(chan GetLotteryModelMasterAsyncResult, 1)
	go p.GetLotteryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateLotteryModelMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateLotteryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateLotteryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateLotteryModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateLotteryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateLotteryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateLotteryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) UpdateLotteryModelMasterAsync(
	request *UpdateLotteryModelMasterRequest,
	callback chan<- UpdateLotteryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/lottery/{lotteryName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.LotteryName != nil && *request.LotteryName != ""  {
        path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
    } else {
        path = strings.ReplaceAll(path, "{lotteryName}", "null")
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
    if request.Method != nil && *request.Method != "" {
        bodies["method"] = *request.Method
    }
    if request.PrizeTableName != nil && *request.PrizeTableName != "" {
        bodies["prizeTableName"] = *request.PrizeTableName
    }
    if request.ChoicePrizeTableScriptId != nil && *request.ChoicePrizeTableScriptId != "" {
        bodies["choicePrizeTableScriptId"] = *request.ChoicePrizeTableScriptId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateLotteryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) UpdateLotteryModelMaster(
	request *UpdateLotteryModelMasterRequest,
) (*UpdateLotteryModelMasterResult, error) {
	callback := make(chan UpdateLotteryModelMasterAsyncResult, 1)
	go p.UpdateLotteryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteLotteryModelMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteLotteryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteLotteryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteLotteryModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteLotteryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteLotteryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteLotteryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DeleteLotteryModelMasterAsync(
	request *DeleteLotteryModelMasterRequest,
	callback chan<- DeleteLotteryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/lottery/{lotteryName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.LotteryName != nil && *request.LotteryName != ""  {
        path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
    } else {
        path = strings.ReplaceAll(path, "{lotteryName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteLotteryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DeleteLotteryModelMaster(
	request *DeleteLotteryModelMasterRequest,
) (*DeleteLotteryModelMasterResult, error) {
	callback := make(chan DeleteLotteryModelMasterAsyncResult, 1)
	go p.DeleteLotteryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describePrizeTableMastersAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribePrizeTableMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePrizeTableMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePrizeTableMastersResult
	if asyncResult.Err != nil {
		callback <- DescribePrizeTableMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribePrizeTableMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribePrizeTableMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribePrizeTableMastersAsync(
	request *DescribePrizeTableMastersRequest,
	callback chan<- DescribePrizeTableMastersAsyncResult,
) {
	path := "/{namespaceName}/master/table"
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

	go describePrizeTableMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribePrizeTableMasters(
	request *DescribePrizeTableMastersRequest,
) (*DescribePrizeTableMastersResult, error) {
	callback := make(chan DescribePrizeTableMastersAsyncResult, 1)
	go p.DescribePrizeTableMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createPrizeTableMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- CreatePrizeTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreatePrizeTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreatePrizeTableMasterResult
	if asyncResult.Err != nil {
		callback <- CreatePrizeTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreatePrizeTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreatePrizeTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) CreatePrizeTableMasterAsync(
	request *CreatePrizeTableMasterRequest,
	callback chan<- CreatePrizeTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/table"
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
    if request.Prizes != nil {
        var _prizes []interface {}
        for _, item := range request.Prizes {
            _prizes = append(_prizes, item)
        }
        bodies["prizes"] = _prizes
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createPrizeTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) CreatePrizeTableMaster(
	request *CreatePrizeTableMasterRequest,
) (*CreatePrizeTableMasterResult, error) {
	callback := make(chan CreatePrizeTableMasterAsyncResult, 1)
	go p.CreatePrizeTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPrizeTableMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetPrizeTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPrizeTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPrizeTableMasterResult
	if asyncResult.Err != nil {
		callback <- GetPrizeTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetPrizeTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetPrizeTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetPrizeTableMasterAsync(
	request *GetPrizeTableMasterRequest,
	callback chan<- GetPrizeTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/table/{prizeTableName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.PrizeTableName != nil && *request.PrizeTableName != ""  {
        path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
    } else {
        path = strings.ReplaceAll(path, "{prizeTableName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getPrizeTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetPrizeTableMaster(
	request *GetPrizeTableMasterRequest,
) (*GetPrizeTableMasterResult, error) {
	callback := make(chan GetPrizeTableMasterAsyncResult, 1)
	go p.GetPrizeTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updatePrizeTableMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdatePrizeTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdatePrizeTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdatePrizeTableMasterResult
	if asyncResult.Err != nil {
		callback <- UpdatePrizeTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdatePrizeTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdatePrizeTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) UpdatePrizeTableMasterAsync(
	request *UpdatePrizeTableMasterRequest,
	callback chan<- UpdatePrizeTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/table/{prizeTableName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.PrizeTableName != nil && *request.PrizeTableName != ""  {
        path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
    } else {
        path = strings.ReplaceAll(path, "{prizeTableName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.Prizes != nil {
        var _prizes []interface {}
        for _, item := range request.Prizes {
            _prizes = append(_prizes, item)
        }
        bodies["prizes"] = _prizes
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updatePrizeTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) UpdatePrizeTableMaster(
	request *UpdatePrizeTableMasterRequest,
) (*UpdatePrizeTableMasterResult, error) {
	callback := make(chan UpdatePrizeTableMasterAsyncResult, 1)
	go p.UpdatePrizeTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deletePrizeTableMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DeletePrizeTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeletePrizeTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeletePrizeTableMasterResult
	if asyncResult.Err != nil {
		callback <- DeletePrizeTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeletePrizeTableMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeletePrizeTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DeletePrizeTableMasterAsync(
	request *DeletePrizeTableMasterRequest,
	callback chan<- DeletePrizeTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/table/{prizeTableName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.PrizeTableName != nil && *request.PrizeTableName != ""  {
        path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
    } else {
        path = strings.ReplaceAll(path, "{prizeTableName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deletePrizeTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DeletePrizeTableMaster(
	request *DeletePrizeTableMasterRequest,
) (*DeletePrizeTableMasterResult, error) {
	callback := make(chan DeletePrizeTableMasterAsyncResult, 1)
	go p.DeletePrizeTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBoxesAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBoxesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBoxesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBoxesResult
	if asyncResult.Err != nil {
		callback <- DescribeBoxesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeBoxesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeBoxesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribeBoxesAsync(
	request *DescribeBoxesRequest,
	callback chan<- DescribeBoxesAsyncResult,
) {
	path := "/{namespaceName}/user/me/box"
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
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go describeBoxesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeBoxes(
	request *DescribeBoxesRequest,
) (*DescribeBoxesResult, error) {
	callback := make(chan DescribeBoxesAsyncResult, 1)
	go p.DescribeBoxesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBoxesByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBoxesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBoxesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBoxesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeBoxesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeBoxesByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeBoxesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribeBoxesByUserIdAsync(
	request *DescribeBoxesByUserIdRequest,
	callback chan<- DescribeBoxesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/box"
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

	go describeBoxesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeBoxesByUserId(
	request *DescribeBoxesByUserIdRequest,
) (*DescribeBoxesByUserIdResult, error) {
	callback := make(chan DescribeBoxesByUserIdAsyncResult, 1)
	go p.DescribeBoxesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBoxAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetBoxAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBoxAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBoxResult
	if asyncResult.Err != nil {
		callback <- GetBoxAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetBoxAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetBoxAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetBoxAsync(
	request *GetBoxRequest,
	callback chan<- GetBoxAsyncResult,
) {
	path := "/{namespaceName}/user/me/box/{prizeTableName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.PrizeTableName != nil && *request.PrizeTableName != ""  {
        path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
    } else {
        path = strings.ReplaceAll(path, "{prizeTableName}", "null")
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

	go getBoxAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetBox(
	request *GetBoxRequest,
) (*GetBoxResult, error) {
	callback := make(chan GetBoxAsyncResult, 1)
	go p.GetBoxAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBoxByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetBoxByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBoxByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBoxByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetBoxByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetBoxByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetBoxByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetBoxByUserIdAsync(
	request *GetBoxByUserIdRequest,
	callback chan<- GetBoxByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/box/{prizeTableName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.PrizeTableName != nil && *request.PrizeTableName != ""  {
        path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
    } else {
        path = strings.ReplaceAll(path, "{prizeTableName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
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

	go getBoxByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetBoxByUserId(
	request *GetBoxByUserIdRequest,
) (*GetBoxByUserIdResult, error) {
	callback := make(chan GetBoxByUserIdAsyncResult, 1)
	go p.GetBoxByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRawBoxByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetRawBoxByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRawBoxByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRawBoxByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetRawBoxByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRawBoxByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetRawBoxByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetRawBoxByUserIdAsync(
	request *GetRawBoxByUserIdRequest,
	callback chan<- GetRawBoxByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/box/{prizeTableName}/raw"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.PrizeTableName != nil && *request.PrizeTableName != ""  {
        path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
    } else {
        path = strings.ReplaceAll(path, "{prizeTableName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
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

	go getRawBoxByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetRawBoxByUserId(
	request *GetRawBoxByUserIdRequest,
) (*GetRawBoxByUserIdResult, error) {
	callback := make(chan GetRawBoxByUserIdAsyncResult, 1)
	go p.GetRawBoxByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func resetBoxAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- ResetBoxAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetBoxAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetBoxResult
	if asyncResult.Err != nil {
		callback <- ResetBoxAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ResetBoxAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ResetBoxAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) ResetBoxAsync(
	request *ResetBoxRequest,
	callback chan<- ResetBoxAsyncResult,
) {
	path := "/{namespaceName}/user/me/box/{prizeTableName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.PrizeTableName != nil && *request.PrizeTableName != ""  {
        path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
    } else {
        path = strings.ReplaceAll(path, "{prizeTableName}", "null")
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

	go resetBoxAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) ResetBox(
	request *ResetBoxRequest,
) (*ResetBoxResult, error) {
	callback := make(chan ResetBoxAsyncResult, 1)
	go p.ResetBoxAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func resetBoxByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- ResetBoxByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetBoxByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetBoxByUserIdResult
	if asyncResult.Err != nil {
		callback <- ResetBoxByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ResetBoxByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ResetBoxByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) ResetBoxByUserIdAsync(
	request *ResetBoxByUserIdRequest,
	callback chan<- ResetBoxByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/box/{prizeTableName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.PrizeTableName != nil && *request.PrizeTableName != ""  {
        path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
    } else {
        path = strings.ReplaceAll(path, "{prizeTableName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
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

	go resetBoxByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) ResetBoxByUserId(
	request *ResetBoxByUserIdRequest,
) (*ResetBoxByUserIdResult, error) {
	callback := make(chan ResetBoxByUserIdAsyncResult, 1)
	go p.ResetBoxByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeLotteryModelsAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLotteryModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLotteryModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLotteryModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeLotteryModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeLotteryModelsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeLotteryModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribeLotteryModelsAsync(
	request *DescribeLotteryModelsRequest,
	callback chan<- DescribeLotteryModelsAsyncResult,
) {
	path := "/{namespaceName}/lottery"
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

	go describeLotteryModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeLotteryModels(
	request *DescribeLotteryModelsRequest,
) (*DescribeLotteryModelsResult, error) {
	callback := make(chan DescribeLotteryModelsAsyncResult, 1)
	go p.DescribeLotteryModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getLotteryModelAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetLotteryModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLotteryModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLotteryModelResult
	if asyncResult.Err != nil {
		callback <- GetLotteryModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetLotteryModelAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetLotteryModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetLotteryModelAsync(
	request *GetLotteryModelRequest,
	callback chan<- GetLotteryModelAsyncResult,
) {
	path := "/{namespaceName}/lottery/{lotteryName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.LotteryName != nil && *request.LotteryName != ""  {
        path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
    } else {
        path = strings.ReplaceAll(path, "{lotteryName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getLotteryModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetLotteryModel(
	request *GetLotteryModelRequest,
) (*GetLotteryModelResult, error) {
	callback := make(chan GetLotteryModelAsyncResult, 1)
	go p.GetLotteryModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describePrizeTablesAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribePrizeTablesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePrizeTablesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePrizeTablesResult
	if asyncResult.Err != nil {
		callback <- DescribePrizeTablesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribePrizeTablesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribePrizeTablesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribePrizeTablesAsync(
	request *DescribePrizeTablesRequest,
	callback chan<- DescribePrizeTablesAsyncResult,
) {
	path := "/{namespaceName}/table"
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

	go describePrizeTablesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribePrizeTables(
	request *DescribePrizeTablesRequest,
) (*DescribePrizeTablesResult, error) {
	callback := make(chan DescribePrizeTablesAsyncResult, 1)
	go p.DescribePrizeTablesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPrizeTableAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetPrizeTableAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPrizeTableAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPrizeTableResult
	if asyncResult.Err != nil {
		callback <- GetPrizeTableAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetPrizeTableAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetPrizeTableAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetPrizeTableAsync(
	request *GetPrizeTableRequest,
	callback chan<- GetPrizeTableAsyncResult,
) {
	path := "/{namespaceName}/table/{prizeTableName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.PrizeTableName != nil && *request.PrizeTableName != ""  {
        path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
    } else {
        path = strings.ReplaceAll(path, "{prizeTableName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getPrizeTableAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetPrizeTable(
	request *GetPrizeTableRequest,
) (*GetPrizeTableResult, error) {
	callback := make(chan GetPrizeTableAsyncResult, 1)
	go p.GetPrizeTableAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func drawByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DrawByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DrawByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DrawByUserIdResult
	if asyncResult.Err != nil {
		callback <- DrawByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DrawByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DrawByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DrawByUserIdAsync(
	request *DrawByUserIdRequest,
	callback chan<- DrawByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/lottery/{lotteryName}/draw"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.LotteryName != nil && *request.LotteryName != ""  {
        path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
    } else {
        path = strings.ReplaceAll(path, "{lotteryName}", "null")
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

	go drawByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DrawByUserId(
	request *DrawByUserIdRequest,
) (*DrawByUserIdResult, error) {
	callback := make(chan DrawByUserIdAsyncResult, 1)
	go p.DrawByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeProbabilitiesAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeProbabilitiesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeProbabilitiesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeProbabilitiesResult
	if asyncResult.Err != nil {
		callback <- DescribeProbabilitiesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeProbabilitiesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeProbabilitiesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribeProbabilitiesAsync(
	request *DescribeProbabilitiesRequest,
	callback chan<- DescribeProbabilitiesAsyncResult,
) {
	path := "/{namespaceName}/user/me/lottery/{lotteryName}/probability"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.LotteryName != nil && *request.LotteryName != ""  {
        path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
    } else {
        path = strings.ReplaceAll(path, "{lotteryName}", "null")
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

	go describeProbabilitiesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeProbabilities(
	request *DescribeProbabilitiesRequest,
) (*DescribeProbabilitiesResult, error) {
	callback := make(chan DescribeProbabilitiesAsyncResult, 1)
	go p.DescribeProbabilitiesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeProbabilitiesByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeProbabilitiesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeProbabilitiesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeProbabilitiesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeProbabilitiesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeProbabilitiesByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeProbabilitiesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribeProbabilitiesByUserIdAsync(
	request *DescribeProbabilitiesByUserIdRequest,
	callback chan<- DescribeProbabilitiesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/lottery/{lotteryName}/probability"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.LotteryName != nil && *request.LotteryName != ""  {
        path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
    } else {
        path = strings.ReplaceAll(path, "{lotteryName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
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

	go describeProbabilitiesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeProbabilitiesByUserId(
	request *DescribeProbabilitiesByUserIdRequest,
) (*DescribeProbabilitiesByUserIdResult, error) {
	callback := make(chan DescribeProbabilitiesByUserIdAsyncResult, 1)
	go p.DescribeProbabilitiesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func drawByStampSheetAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DrawByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DrawByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DrawByStampSheetResult
	if asyncResult.Err != nil {
		callback <- DrawByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DrawByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DrawByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DrawByStampSheetAsync(
	request *DrawByStampSheetRequest,
	callback chan<- DrawByStampSheetAsyncResult,
) {
	path := "/stamp/draw"

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

	go drawByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DrawByStampSheet(
	request *DrawByStampSheetRequest,
) (*DrawByStampSheetResult, error) {
	callback := make(chan DrawByStampSheetAsyncResult, 1)
	go p.DrawByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) ExportMaster(
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

func getCurrentLotteryMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentLotteryMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentLotteryMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentLotteryMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentLotteryMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentLotteryMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentLotteryMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetCurrentLotteryMasterAsync(
	request *GetCurrentLotteryMasterRequest,
	callback chan<- GetCurrentLotteryMasterAsyncResult,
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

	go getCurrentLotteryMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetCurrentLotteryMaster(
	request *GetCurrentLotteryMasterRequest,
) (*GetCurrentLotteryMasterResult, error) {
	callback := make(chan GetCurrentLotteryMasterAsyncResult, 1)
	go p.GetCurrentLotteryMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentLotteryMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentLotteryMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentLotteryMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentLotteryMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentLotteryMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentLotteryMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentLotteryMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) UpdateCurrentLotteryMasterAsync(
	request *UpdateCurrentLotteryMasterRequest,
	callback chan<- UpdateCurrentLotteryMasterAsyncResult,
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

	go updateCurrentLotteryMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) UpdateCurrentLotteryMaster(
	request *UpdateCurrentLotteryMasterRequest,
) (*UpdateCurrentLotteryMasterResult, error) {
	callback := make(chan UpdateCurrentLotteryMasterAsyncResult, 1)
	go p.UpdateCurrentLotteryMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentLotteryMasterFromGitHubAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentLotteryMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentLotteryMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentLotteryMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentLotteryMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentLotteryMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentLotteryMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) UpdateCurrentLotteryMasterFromGitHubAsync(
	request *UpdateCurrentLotteryMasterFromGitHubRequest,
	callback chan<- UpdateCurrentLotteryMasterFromGitHubAsyncResult,
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

	go updateCurrentLotteryMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) UpdateCurrentLotteryMasterFromGitHub(
	request *UpdateCurrentLotteryMasterFromGitHubRequest,
) (*UpdateCurrentLotteryMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentLotteryMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentLotteryMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
