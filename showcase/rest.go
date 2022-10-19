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

package showcase

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

type Gs2ShowcaseRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2ShowcaseRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2ShowcaseRestClient,
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

func (p Gs2ShowcaseRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) DescribeNamespaces(
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
	client Gs2ShowcaseRestClient,
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

func (p Gs2ShowcaseRestClient) CreateNamespaceAsync(
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
    if request.BuyScript != nil {
        bodies["buyScript"] = request.BuyScript.ToDict()
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) CreateNamespace(
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
	client Gs2ShowcaseRestClient,
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

func (p Gs2ShowcaseRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) GetNamespaceStatus(
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
	client Gs2ShowcaseRestClient,
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

func (p Gs2ShowcaseRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) GetNamespace(
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
	client Gs2ShowcaseRestClient,
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

func (p Gs2ShowcaseRestClient) UpdateNamespaceAsync(
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
    if request.TransactionSetting != nil {
        bodies["transactionSetting"] = request.TransactionSetting.ToDict()
    }
    if request.BuyScript != nil {
        bodies["buyScript"] = request.BuyScript.ToDict()
    }
    if request.LogSetting != nil {
        bodies["logSetting"] = request.LogSetting.ToDict()
    }
    if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
        bodies["queueNamespaceId"] = *request.QueueNamespaceId
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

	go updateNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) UpdateNamespace(
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
	client Gs2ShowcaseRestClient,
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

func (p Gs2ShowcaseRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) DeleteNamespace(
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

func describeSalesItemMastersAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSalesItemMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSalesItemMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSalesItemMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeSalesItemMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeSalesItemMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeSalesItemMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) DescribeSalesItemMastersAsync(
	request *DescribeSalesItemMastersRequest,
	callback chan<- DescribeSalesItemMastersAsyncResult,
) {
	path := "/{namespaceName}/master/salesItem"
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

	go describeSalesItemMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) DescribeSalesItemMasters(
	request *DescribeSalesItemMastersRequest,
) (*DescribeSalesItemMastersResult, error) {
	callback := make(chan DescribeSalesItemMastersAsyncResult, 1)
	go p.DescribeSalesItemMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createSalesItemMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- CreateSalesItemMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateSalesItemMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateSalesItemMasterResult
	if asyncResult.Err != nil {
		callback <- CreateSalesItemMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateSalesItemMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateSalesItemMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) CreateSalesItemMasterAsync(
	request *CreateSalesItemMasterRequest,
	callback chan<- CreateSalesItemMasterAsyncResult,
) {
	path := "/{namespaceName}/master/salesItem"
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
    if request.ConsumeActions != nil {
        var _consumeActions []interface {}
        for _, item := range request.ConsumeActions {
            _consumeActions = append(_consumeActions, item)
        }
        bodies["consumeActions"] = _consumeActions
    }
    if request.AcquireActions != nil {
        var _acquireActions []interface {}
        for _, item := range request.AcquireActions {
            _acquireActions = append(_acquireActions, item)
        }
        bodies["acquireActions"] = _acquireActions
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createSalesItemMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) CreateSalesItemMaster(
	request *CreateSalesItemMasterRequest,
) (*CreateSalesItemMasterResult, error) {
	callback := make(chan CreateSalesItemMasterAsyncResult, 1)
	go p.CreateSalesItemMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSalesItemMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- GetSalesItemMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSalesItemMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSalesItemMasterResult
	if asyncResult.Err != nil {
		callback <- GetSalesItemMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetSalesItemMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetSalesItemMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) GetSalesItemMasterAsync(
	request *GetSalesItemMasterRequest,
	callback chan<- GetSalesItemMasterAsyncResult,
) {
	path := "/{namespaceName}/master/salesItem/{salesItemName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.SalesItemName != nil && *request.SalesItemName != ""  {
        path = strings.ReplaceAll(path, "{salesItemName}", core.ToString(*request.SalesItemName))
    } else {
        path = strings.ReplaceAll(path, "{salesItemName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getSalesItemMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) GetSalesItemMaster(
	request *GetSalesItemMasterRequest,
) (*GetSalesItemMasterResult, error) {
	callback := make(chan GetSalesItemMasterAsyncResult, 1)
	go p.GetSalesItemMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateSalesItemMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateSalesItemMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateSalesItemMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateSalesItemMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateSalesItemMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateSalesItemMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateSalesItemMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) UpdateSalesItemMasterAsync(
	request *UpdateSalesItemMasterRequest,
	callback chan<- UpdateSalesItemMasterAsyncResult,
) {
	path := "/{namespaceName}/master/salesItem/{salesItemName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.SalesItemName != nil && *request.SalesItemName != ""  {
        path = strings.ReplaceAll(path, "{salesItemName}", core.ToString(*request.SalesItemName))
    } else {
        path = strings.ReplaceAll(path, "{salesItemName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.ConsumeActions != nil {
        var _consumeActions []interface {}
        for _, item := range request.ConsumeActions {
            _consumeActions = append(_consumeActions, item)
        }
        bodies["consumeActions"] = _consumeActions
    }
    if request.AcquireActions != nil {
        var _acquireActions []interface {}
        for _, item := range request.AcquireActions {
            _acquireActions = append(_acquireActions, item)
        }
        bodies["acquireActions"] = _acquireActions
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateSalesItemMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) UpdateSalesItemMaster(
	request *UpdateSalesItemMasterRequest,
) (*UpdateSalesItemMasterResult, error) {
	callback := make(chan UpdateSalesItemMasterAsyncResult, 1)
	go p.UpdateSalesItemMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteSalesItemMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteSalesItemMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSalesItemMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSalesItemMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteSalesItemMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteSalesItemMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteSalesItemMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) DeleteSalesItemMasterAsync(
	request *DeleteSalesItemMasterRequest,
	callback chan<- DeleteSalesItemMasterAsyncResult,
) {
	path := "/{namespaceName}/master/salesItem/{salesItemName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.SalesItemName != nil && *request.SalesItemName != ""  {
        path = strings.ReplaceAll(path, "{salesItemName}", core.ToString(*request.SalesItemName))
    } else {
        path = strings.ReplaceAll(path, "{salesItemName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteSalesItemMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) DeleteSalesItemMaster(
	request *DeleteSalesItemMasterRequest,
) (*DeleteSalesItemMasterResult, error) {
	callback := make(chan DeleteSalesItemMasterAsyncResult, 1)
	go p.DeleteSalesItemMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSalesItemGroupMastersAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSalesItemGroupMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSalesItemGroupMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSalesItemGroupMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeSalesItemGroupMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeSalesItemGroupMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeSalesItemGroupMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) DescribeSalesItemGroupMastersAsync(
	request *DescribeSalesItemGroupMastersRequest,
	callback chan<- DescribeSalesItemGroupMastersAsyncResult,
) {
	path := "/{namespaceName}/master/group"
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

	go describeSalesItemGroupMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) DescribeSalesItemGroupMasters(
	request *DescribeSalesItemGroupMastersRequest,
) (*DescribeSalesItemGroupMastersResult, error) {
	callback := make(chan DescribeSalesItemGroupMastersAsyncResult, 1)
	go p.DescribeSalesItemGroupMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createSalesItemGroupMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- CreateSalesItemGroupMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateSalesItemGroupMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateSalesItemGroupMasterResult
	if asyncResult.Err != nil {
		callback <- CreateSalesItemGroupMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateSalesItemGroupMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateSalesItemGroupMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) CreateSalesItemGroupMasterAsync(
	request *CreateSalesItemGroupMasterRequest,
	callback chan<- CreateSalesItemGroupMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group"
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
    if request.SalesItemNames != nil {
        var _salesItemNames []interface {}
        for _, item := range request.SalesItemNames {
            _salesItemNames = append(_salesItemNames, item)
        }
        bodies["salesItemNames"] = _salesItemNames
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createSalesItemGroupMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) CreateSalesItemGroupMaster(
	request *CreateSalesItemGroupMasterRequest,
) (*CreateSalesItemGroupMasterResult, error) {
	callback := make(chan CreateSalesItemGroupMasterAsyncResult, 1)
	go p.CreateSalesItemGroupMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSalesItemGroupMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- GetSalesItemGroupMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSalesItemGroupMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSalesItemGroupMasterResult
	if asyncResult.Err != nil {
		callback <- GetSalesItemGroupMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetSalesItemGroupMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetSalesItemGroupMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) GetSalesItemGroupMasterAsync(
	request *GetSalesItemGroupMasterRequest,
	callback chan<- GetSalesItemGroupMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{salesItemGroupName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.SalesItemGroupName != nil && *request.SalesItemGroupName != ""  {
        path = strings.ReplaceAll(path, "{salesItemGroupName}", core.ToString(*request.SalesItemGroupName))
    } else {
        path = strings.ReplaceAll(path, "{salesItemGroupName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getSalesItemGroupMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) GetSalesItemGroupMaster(
	request *GetSalesItemGroupMasterRequest,
) (*GetSalesItemGroupMasterResult, error) {
	callback := make(chan GetSalesItemGroupMasterAsyncResult, 1)
	go p.GetSalesItemGroupMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateSalesItemGroupMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateSalesItemGroupMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateSalesItemGroupMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateSalesItemGroupMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateSalesItemGroupMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateSalesItemGroupMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateSalesItemGroupMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) UpdateSalesItemGroupMasterAsync(
	request *UpdateSalesItemGroupMasterRequest,
	callback chan<- UpdateSalesItemGroupMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{salesItemGroupName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.SalesItemGroupName != nil && *request.SalesItemGroupName != ""  {
        path = strings.ReplaceAll(path, "{salesItemGroupName}", core.ToString(*request.SalesItemGroupName))
    } else {
        path = strings.ReplaceAll(path, "{salesItemGroupName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.SalesItemNames != nil {
        var _salesItemNames []interface {}
        for _, item := range request.SalesItemNames {
            _salesItemNames = append(_salesItemNames, item)
        }
        bodies["salesItemNames"] = _salesItemNames
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateSalesItemGroupMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) UpdateSalesItemGroupMaster(
	request *UpdateSalesItemGroupMasterRequest,
) (*UpdateSalesItemGroupMasterResult, error) {
	callback := make(chan UpdateSalesItemGroupMasterAsyncResult, 1)
	go p.UpdateSalesItemGroupMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteSalesItemGroupMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteSalesItemGroupMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSalesItemGroupMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSalesItemGroupMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteSalesItemGroupMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteSalesItemGroupMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteSalesItemGroupMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) DeleteSalesItemGroupMasterAsync(
	request *DeleteSalesItemGroupMasterRequest,
	callback chan<- DeleteSalesItemGroupMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{salesItemGroupName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.SalesItemGroupName != nil && *request.SalesItemGroupName != ""  {
        path = strings.ReplaceAll(path, "{salesItemGroupName}", core.ToString(*request.SalesItemGroupName))
    } else {
        path = strings.ReplaceAll(path, "{salesItemGroupName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteSalesItemGroupMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) DeleteSalesItemGroupMaster(
	request *DeleteSalesItemGroupMasterRequest,
) (*DeleteSalesItemGroupMasterResult, error) {
	callback := make(chan DeleteSalesItemGroupMasterAsyncResult, 1)
	go p.DeleteSalesItemGroupMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeShowcaseMastersAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeShowcaseMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeShowcaseMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeShowcaseMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeShowcaseMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeShowcaseMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeShowcaseMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) DescribeShowcaseMastersAsync(
	request *DescribeShowcaseMastersRequest,
	callback chan<- DescribeShowcaseMastersAsyncResult,
) {
	path := "/{namespaceName}/master/showcase"
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

	go describeShowcaseMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) DescribeShowcaseMasters(
	request *DescribeShowcaseMastersRequest,
) (*DescribeShowcaseMastersResult, error) {
	callback := make(chan DescribeShowcaseMastersAsyncResult, 1)
	go p.DescribeShowcaseMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createShowcaseMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- CreateShowcaseMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateShowcaseMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateShowcaseMasterResult
	if asyncResult.Err != nil {
		callback <- CreateShowcaseMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateShowcaseMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateShowcaseMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) CreateShowcaseMasterAsync(
	request *CreateShowcaseMasterRequest,
	callback chan<- CreateShowcaseMasterAsyncResult,
) {
	path := "/{namespaceName}/master/showcase"
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
    if request.DisplayItems != nil {
        var _displayItems []interface {}
        for _, item := range request.DisplayItems {
            _displayItems = append(_displayItems, item)
        }
        bodies["displayItems"] = _displayItems
    }
    if request.SalesPeriodEventId != nil && *request.SalesPeriodEventId != "" {
        bodies["salesPeriodEventId"] = *request.SalesPeriodEventId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createShowcaseMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) CreateShowcaseMaster(
	request *CreateShowcaseMasterRequest,
) (*CreateShowcaseMasterResult, error) {
	callback := make(chan CreateShowcaseMasterAsyncResult, 1)
	go p.CreateShowcaseMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getShowcaseMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- GetShowcaseMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetShowcaseMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetShowcaseMasterResult
	if asyncResult.Err != nil {
		callback <- GetShowcaseMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetShowcaseMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetShowcaseMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) GetShowcaseMasterAsync(
	request *GetShowcaseMasterRequest,
	callback chan<- GetShowcaseMasterAsyncResult,
) {
	path := "/{namespaceName}/master/showcase/{showcaseName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != ""  {
        path = strings.ReplaceAll(path, "{showcaseName}", core.ToString(*request.ShowcaseName))
    } else {
        path = strings.ReplaceAll(path, "{showcaseName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getShowcaseMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) GetShowcaseMaster(
	request *GetShowcaseMasterRequest,
) (*GetShowcaseMasterResult, error) {
	callback := make(chan GetShowcaseMasterAsyncResult, 1)
	go p.GetShowcaseMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateShowcaseMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateShowcaseMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateShowcaseMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateShowcaseMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateShowcaseMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateShowcaseMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateShowcaseMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) UpdateShowcaseMasterAsync(
	request *UpdateShowcaseMasterRequest,
	callback chan<- UpdateShowcaseMasterAsyncResult,
) {
	path := "/{namespaceName}/master/showcase/{showcaseName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != ""  {
        path = strings.ReplaceAll(path, "{showcaseName}", core.ToString(*request.ShowcaseName))
    } else {
        path = strings.ReplaceAll(path, "{showcaseName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.DisplayItems != nil {
        var _displayItems []interface {}
        for _, item := range request.DisplayItems {
            _displayItems = append(_displayItems, item)
        }
        bodies["displayItems"] = _displayItems
    }
    if request.SalesPeriodEventId != nil && *request.SalesPeriodEventId != "" {
        bodies["salesPeriodEventId"] = *request.SalesPeriodEventId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateShowcaseMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) UpdateShowcaseMaster(
	request *UpdateShowcaseMasterRequest,
) (*UpdateShowcaseMasterResult, error) {
	callback := make(chan UpdateShowcaseMasterAsyncResult, 1)
	go p.UpdateShowcaseMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteShowcaseMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteShowcaseMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteShowcaseMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteShowcaseMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteShowcaseMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteShowcaseMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteShowcaseMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) DeleteShowcaseMasterAsync(
	request *DeleteShowcaseMasterRequest,
	callback chan<- DeleteShowcaseMasterAsyncResult,
) {
	path := "/{namespaceName}/master/showcase/{showcaseName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != ""  {
        path = strings.ReplaceAll(path, "{showcaseName}", core.ToString(*request.ShowcaseName))
    } else {
        path = strings.ReplaceAll(path, "{showcaseName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteShowcaseMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) DeleteShowcaseMaster(
	request *DeleteShowcaseMasterRequest,
) (*DeleteShowcaseMasterResult, error) {
	callback := make(chan DeleteShowcaseMasterAsyncResult, 1)
	go p.DeleteShowcaseMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
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

func (p Gs2ShowcaseRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) ExportMaster(
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

func getCurrentShowcaseMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentShowcaseMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentShowcaseMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentShowcaseMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentShowcaseMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentShowcaseMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentShowcaseMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) GetCurrentShowcaseMasterAsync(
	request *GetCurrentShowcaseMasterRequest,
	callback chan<- GetCurrentShowcaseMasterAsyncResult,
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

	go getCurrentShowcaseMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) GetCurrentShowcaseMaster(
	request *GetCurrentShowcaseMasterRequest,
) (*GetCurrentShowcaseMasterResult, error) {
	callback := make(chan GetCurrentShowcaseMasterAsyncResult, 1)
	go p.GetCurrentShowcaseMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentShowcaseMasterAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentShowcaseMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentShowcaseMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentShowcaseMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentShowcaseMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentShowcaseMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentShowcaseMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) UpdateCurrentShowcaseMasterAsync(
	request *UpdateCurrentShowcaseMasterRequest,
	callback chan<- UpdateCurrentShowcaseMasterAsyncResult,
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

	go updateCurrentShowcaseMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) UpdateCurrentShowcaseMaster(
	request *UpdateCurrentShowcaseMasterRequest,
) (*UpdateCurrentShowcaseMasterResult, error) {
	callback := make(chan UpdateCurrentShowcaseMasterAsyncResult, 1)
	go p.UpdateCurrentShowcaseMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentShowcaseMasterFromGitHubAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentShowcaseMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentShowcaseMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentShowcaseMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentShowcaseMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentShowcaseMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentShowcaseMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) UpdateCurrentShowcaseMasterFromGitHubAsync(
	request *UpdateCurrentShowcaseMasterFromGitHubRequest,
	callback chan<- UpdateCurrentShowcaseMasterFromGitHubAsyncResult,
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

	go updateCurrentShowcaseMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) UpdateCurrentShowcaseMasterFromGitHub(
	request *UpdateCurrentShowcaseMasterFromGitHubRequest,
) (*UpdateCurrentShowcaseMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentShowcaseMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentShowcaseMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeShowcasesAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeShowcasesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeShowcasesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeShowcasesResult
	if asyncResult.Err != nil {
		callback <- DescribeShowcasesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeShowcasesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeShowcasesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) DescribeShowcasesAsync(
	request *DescribeShowcasesRequest,
	callback chan<- DescribeShowcasesAsyncResult,
) {
	path := "/{namespaceName}/user/me/showcase"
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
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go describeShowcasesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) DescribeShowcases(
	request *DescribeShowcasesRequest,
) (*DescribeShowcasesResult, error) {
	callback := make(chan DescribeShowcasesAsyncResult, 1)
	go p.DescribeShowcasesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeShowcasesByUserIdAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeShowcasesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeShowcasesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeShowcasesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeShowcasesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeShowcasesByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeShowcasesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) DescribeShowcasesByUserIdAsync(
	request *DescribeShowcasesByUserIdRequest,
	callback chan<- DescribeShowcasesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/showcase"
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go describeShowcasesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) DescribeShowcasesByUserId(
	request *DescribeShowcasesByUserIdRequest,
) (*DescribeShowcasesByUserIdResult, error) {
	callback := make(chan DescribeShowcasesByUserIdAsyncResult, 1)
	go p.DescribeShowcasesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getShowcaseAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- GetShowcaseAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetShowcaseAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetShowcaseResult
	if asyncResult.Err != nil {
		callback <- GetShowcaseAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetShowcaseAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetShowcaseAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) GetShowcaseAsync(
	request *GetShowcaseRequest,
	callback chan<- GetShowcaseAsyncResult,
) {
	path := "/{namespaceName}/user/me/showcase/{showcaseName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != ""  {
        path = strings.ReplaceAll(path, "{showcaseName}", core.ToString(*request.ShowcaseName))
    } else {
        path = strings.ReplaceAll(path, "{showcaseName}", "null")
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

	go getShowcaseAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) GetShowcase(
	request *GetShowcaseRequest,
) (*GetShowcaseResult, error) {
	callback := make(chan GetShowcaseAsyncResult, 1)
	go p.GetShowcaseAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getShowcaseByUserIdAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- GetShowcaseByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetShowcaseByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetShowcaseByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetShowcaseByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetShowcaseByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetShowcaseByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) GetShowcaseByUserIdAsync(
	request *GetShowcaseByUserIdRequest,
	callback chan<- GetShowcaseByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/showcase/{showcaseName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != ""  {
        path = strings.ReplaceAll(path, "{showcaseName}", core.ToString(*request.ShowcaseName))
    } else {
        path = strings.ReplaceAll(path, "{showcaseName}", "null")
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

	go getShowcaseByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) GetShowcaseByUserId(
	request *GetShowcaseByUserIdRequest,
) (*GetShowcaseByUserIdResult, error) {
	callback := make(chan GetShowcaseByUserIdAsyncResult, 1)
	go p.GetShowcaseByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func buyAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- BuyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BuyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BuyResult
	if asyncResult.Err != nil {
		callback <- BuyAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- BuyAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- BuyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) BuyAsync(
	request *BuyRequest,
	callback chan<- BuyAsyncResult,
) {
	path := "/{namespaceName}/user/me/showcase/{showcaseName}/{displayItemId}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != ""  {
        path = strings.ReplaceAll(path, "{showcaseName}", core.ToString(*request.ShowcaseName))
    } else {
        path = strings.ReplaceAll(path, "{showcaseName}", "null")
    }
    if request.DisplayItemId != nil && *request.DisplayItemId != ""  {
        path = strings.ReplaceAll(path, "{displayItemId}", core.ToString(*request.DisplayItemId))
    } else {
        path = strings.ReplaceAll(path, "{displayItemId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Quantity != nil {
        bodies["quantity"] = *request.Quantity
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

	go buyAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) Buy(
	request *BuyRequest,
) (*BuyResult, error) {
	callback := make(chan BuyAsyncResult, 1)
	go p.BuyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func buyByUserIdAsyncHandler(
	client Gs2ShowcaseRestClient,
	job *core.NetworkJob,
	callback chan<- BuyByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BuyByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BuyByUserIdResult
	if asyncResult.Err != nil {
		callback <- BuyByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- BuyByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- BuyByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ShowcaseRestClient) BuyByUserIdAsync(
	request *BuyByUserIdRequest,
	callback chan<- BuyByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/showcase/{showcaseName}/{displayItemId}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != ""  {
        path = strings.ReplaceAll(path, "{showcaseName}", core.ToString(*request.ShowcaseName))
    } else {
        path = strings.ReplaceAll(path, "{showcaseName}", "null")
    }
    if request.DisplayItemId != nil && *request.DisplayItemId != ""  {
        path = strings.ReplaceAll(path, "{displayItemId}", core.ToString(*request.DisplayItemId))
    } else {
        path = strings.ReplaceAll(path, "{displayItemId}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Quantity != nil {
        bodies["quantity"] = *request.Quantity
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

	go buyByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("showcase").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseRestClient) BuyByUserId(
	request *BuyByUserIdRequest,
) (*BuyByUserIdResult, error) {
	callback := make(chan BuyByUserIdAsyncResult, 1)
	go p.BuyByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
