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

package inventory

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

type Gs2InventoryRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2InventoryRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeNamespaces(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) CreateNamespaceAsync(
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
    if request.AcquireScript != nil {
        bodies["acquireScript"] = request.AcquireScript.ToDict()
    }
    if request.OverflowScript != nil {
        bodies["overflowScript"] = request.OverflowScript.ToDict()
    }
    if request.ConsumeScript != nil {
        bodies["consumeScript"] = request.ConsumeScript.ToDict()
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CreateNamespace(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetNamespaceStatus(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetNamespace(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) UpdateNamespaceAsync(
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
    if request.AcquireScript != nil {
        bodies["acquireScript"] = request.AcquireScript.ToDict()
    }
    if request.OverflowScript != nil {
        bodies["overflowScript"] = request.OverflowScript.ToDict()
    }
    if request.ConsumeScript != nil {
        bodies["consumeScript"] = request.ConsumeScript.ToDict()
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateNamespace(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteNamespace(
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

func describeInventoryModelMastersAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInventoryModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInventoryModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInventoryModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeInventoryModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeInventoryModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeInventoryModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeInventoryModelMastersAsync(
	request *DescribeInventoryModelMastersRequest,
	callback chan<- DescribeInventoryModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/inventory"
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

	go describeInventoryModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeInventoryModelMasters(
	request *DescribeInventoryModelMastersRequest,
) (*DescribeInventoryModelMastersResult, error) {
	callback := make(chan DescribeInventoryModelMastersAsyncResult, 1)
	go p.DescribeInventoryModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- CreateInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateInventoryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) CreateInventoryModelMasterAsync(
	request *CreateInventoryModelMasterRequest,
	callback chan<- CreateInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory"
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
    if request.InitialCapacity != nil {
        bodies["initialCapacity"] = *request.InitialCapacity
    }
    if request.MaxCapacity != nil {
        bodies["maxCapacity"] = *request.MaxCapacity
    }
    if request.ProtectReferencedItem != nil {
        bodies["protectReferencedItem"] = *request.ProtectReferencedItem
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CreateInventoryModelMaster(
	request *CreateInventoryModelMasterRequest,
) (*CreateInventoryModelMasterResult, error) {
	callback := make(chan CreateInventoryModelMasterAsyncResult, 1)
	go p.CreateInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetInventoryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetInventoryModelMasterAsync(
	request *GetInventoryModelMasterRequest,
	callback chan<- GetInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetInventoryModelMaster(
	request *GetInventoryModelMasterRequest,
) (*GetInventoryModelMasterResult, error) {
	callback := make(chan GetInventoryModelMasterAsyncResult, 1)
	go p.GetInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateInventoryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) UpdateInventoryModelMasterAsync(
	request *UpdateInventoryModelMasterRequest,
	callback chan<- UpdateInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.InitialCapacity != nil {
        bodies["initialCapacity"] = *request.InitialCapacity
    }
    if request.MaxCapacity != nil {
        bodies["maxCapacity"] = *request.MaxCapacity
    }
    if request.ProtectReferencedItem != nil {
        bodies["protectReferencedItem"] = *request.ProtectReferencedItem
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateInventoryModelMaster(
	request *UpdateInventoryModelMasterRequest,
) (*UpdateInventoryModelMasterResult, error) {
	callback := make(chan UpdateInventoryModelMasterAsyncResult, 1)
	go p.UpdateInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteInventoryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteInventoryModelMasterAsync(
	request *DeleteInventoryModelMasterRequest,
	callback chan<- DeleteInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteInventoryModelMaster(
	request *DeleteInventoryModelMasterRequest,
) (*DeleteInventoryModelMasterResult, error) {
	callback := make(chan DeleteInventoryModelMasterAsyncResult, 1)
	go p.DeleteInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeInventoryModelsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInventoryModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInventoryModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInventoryModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeInventoryModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeInventoryModelsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeInventoryModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeInventoryModelsAsync(
	request *DescribeInventoryModelsRequest,
	callback chan<- DescribeInventoryModelsAsyncResult,
) {
	path := "/{namespaceName}/inventory"
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

	go describeInventoryModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeInventoryModels(
	request *DescribeInventoryModelsRequest,
) (*DescribeInventoryModelsResult, error) {
	callback := make(chan DescribeInventoryModelsAsyncResult, 1)
	go p.DescribeInventoryModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getInventoryModelAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetInventoryModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetInventoryModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetInventoryModelResult
	if asyncResult.Err != nil {
		callback <- GetInventoryModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetInventoryModelAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetInventoryModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetInventoryModelAsync(
	request *GetInventoryModelRequest,
	callback chan<- GetInventoryModelAsyncResult,
) {
	path := "/{namespaceName}/inventory/{inventoryName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getInventoryModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetInventoryModel(
	request *GetInventoryModelRequest,
) (*GetInventoryModelResult, error) {
	callback := make(chan GetInventoryModelAsyncResult, 1)
	go p.GetInventoryModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeItemModelMastersAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeItemModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeItemModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeItemModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeItemModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeItemModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeItemModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeItemModelMastersAsync(
	request *DescribeItemModelMastersRequest,
	callback chan<- DescribeItemModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}/item"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
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

	go describeItemModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeItemModelMasters(
	request *DescribeItemModelMastersRequest,
) (*DescribeItemModelMastersResult, error) {
	callback := make(chan DescribeItemModelMastersAsyncResult, 1)
	go p.DescribeItemModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- CreateItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateItemModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) CreateItemModelMasterAsync(
	request *CreateItemModelMasterRequest,
	callback chan<- CreateItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}/item"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
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
    if request.StackingLimit != nil {
        bodies["stackingLimit"] = *request.StackingLimit
    }
    if request.AllowMultipleStacks != nil {
        bodies["allowMultipleStacks"] = *request.AllowMultipleStacks
    }
    if request.SortValue != nil {
        bodies["sortValue"] = *request.SortValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CreateItemModelMaster(
	request *CreateItemModelMasterRequest,
) (*CreateItemModelMasterResult, error) {
	callback := make(chan CreateItemModelMasterAsyncResult, 1)
	go p.CreateItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetItemModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetItemModelMasterAsync(
	request *GetItemModelMasterRequest,
	callback chan<- GetItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}/item/{itemName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetItemModelMaster(
	request *GetItemModelMasterRequest,
) (*GetItemModelMasterResult, error) {
	callback := make(chan GetItemModelMasterAsyncResult, 1)
	go p.GetItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateItemModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) UpdateItemModelMasterAsync(
	request *UpdateItemModelMasterRequest,
	callback chan<- UpdateItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}/item/{itemName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.StackingLimit != nil {
        bodies["stackingLimit"] = *request.StackingLimit
    }
    if request.AllowMultipleStacks != nil {
        bodies["allowMultipleStacks"] = *request.AllowMultipleStacks
    }
    if request.SortValue != nil {
        bodies["sortValue"] = *request.SortValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateItemModelMaster(
	request *UpdateItemModelMasterRequest,
) (*UpdateItemModelMasterResult, error) {
	callback := make(chan UpdateItemModelMasterAsyncResult, 1)
	go p.UpdateItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteItemModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteItemModelMasterAsync(
	request *DeleteItemModelMasterRequest,
	callback chan<- DeleteItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}/item/{itemName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteItemModelMaster(
	request *DeleteItemModelMasterRequest,
) (*DeleteItemModelMasterResult, error) {
	callback := make(chan DeleteItemModelMasterAsyncResult, 1)
	go p.DeleteItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeItemModelsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeItemModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeItemModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeItemModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeItemModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeItemModelsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeItemModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeItemModelsAsync(
	request *DescribeItemModelsRequest,
	callback chan<- DescribeItemModelsAsyncResult,
) {
	path := "/{namespaceName}/inventory/{inventoryName}/item"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go describeItemModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeItemModels(
	request *DescribeItemModelsRequest,
) (*DescribeItemModelsResult, error) {
	callback := make(chan DescribeItemModelsAsyncResult, 1)
	go p.DescribeItemModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getItemModelAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetItemModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetItemModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetItemModelResult
	if asyncResult.Err != nil {
		callback <- GetItemModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetItemModelAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetItemModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetItemModelAsync(
	request *GetItemModelRequest,
	callback chan<- GetItemModelAsyncResult,
) {
	path := "/{namespaceName}/inventory/{inventoryName}/item/{itemName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getItemModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetItemModel(
	request *GetItemModelRequest,
) (*GetItemModelResult, error) {
	callback := make(chan GetItemModelAsyncResult, 1)
	go p.GetItemModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ExportMaster(
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

func getCurrentItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentItemModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetCurrentItemModelMasterAsync(
	request *GetCurrentItemModelMasterRequest,
	callback chan<- GetCurrentItemModelMasterAsyncResult,
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

	go getCurrentItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetCurrentItemModelMaster(
	request *GetCurrentItemModelMasterRequest,
) (*GetCurrentItemModelMasterResult, error) {
	callback := make(chan GetCurrentItemModelMasterAsyncResult, 1)
	go p.GetCurrentItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentItemModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) UpdateCurrentItemModelMasterAsync(
	request *UpdateCurrentItemModelMasterRequest,
	callback chan<- UpdateCurrentItemModelMasterAsyncResult,
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

	go updateCurrentItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateCurrentItemModelMaster(
	request *UpdateCurrentItemModelMasterRequest,
) (*UpdateCurrentItemModelMasterResult, error) {
	callback := make(chan UpdateCurrentItemModelMasterAsyncResult, 1)
	go p.UpdateCurrentItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentItemModelMasterFromGitHubAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentItemModelMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentItemModelMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentItemModelMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentItemModelMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentItemModelMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentItemModelMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) UpdateCurrentItemModelMasterFromGitHubAsync(
	request *UpdateCurrentItemModelMasterFromGitHubRequest,
	callback chan<- UpdateCurrentItemModelMasterFromGitHubAsyncResult,
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

	go updateCurrentItemModelMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateCurrentItemModelMasterFromGitHub(
	request *UpdateCurrentItemModelMasterFromGitHubRequest,
) (*UpdateCurrentItemModelMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentItemModelMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentItemModelMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeInventoriesAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInventoriesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInventoriesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInventoriesResult
	if asyncResult.Err != nil {
		callback <- DescribeInventoriesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeInventoriesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeInventoriesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeInventoriesAsync(
	request *DescribeInventoriesRequest,
	callback chan<- DescribeInventoriesAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory"
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

	go describeInventoriesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeInventories(
	request *DescribeInventoriesRequest,
) (*DescribeInventoriesResult, error) {
	callback := make(chan DescribeInventoriesAsyncResult, 1)
	go p.DescribeInventoriesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeInventoriesByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInventoriesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInventoriesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInventoriesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeInventoriesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeInventoriesByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeInventoriesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeInventoriesByUserIdAsync(
	request *DescribeInventoriesByUserIdRequest,
	callback chan<- DescribeInventoriesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory"
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

	go describeInventoriesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeInventoriesByUserId(
	request *DescribeInventoriesByUserIdRequest,
) (*DescribeInventoriesByUserIdResult, error) {
	callback := make(chan DescribeInventoriesByUserIdAsyncResult, 1)
	go p.DescribeInventoriesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getInventoryAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetInventoryAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetInventoryAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetInventoryResult
	if asyncResult.Err != nil {
		callback <- GetInventoryAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetInventoryAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetInventoryAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetInventoryAsync(
	request *GetInventoryRequest,
	callback chan<- GetInventoryAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
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

	go getInventoryAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetInventory(
	request *GetInventoryRequest,
) (*GetInventoryResult, error) {
	callback := make(chan GetInventoryAsyncResult, 1)
	go p.GetInventoryAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getInventoryByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetInventoryByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetInventoryByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetInventoryByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetInventoryByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetInventoryByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetInventoryByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetInventoryByUserIdAsync(
	request *GetInventoryByUserIdRequest,
	callback chan<- GetInventoryByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
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

	go getInventoryByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetInventoryByUserId(
	request *GetInventoryByUserIdRequest,
) (*GetInventoryByUserIdResult, error) {
	callback := make(chan GetInventoryByUserIdAsyncResult, 1)
	go p.GetInventoryByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addCapacityByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AddCapacityByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddCapacityByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddCapacityByUserIdResult
	if asyncResult.Err != nil {
		callback <- AddCapacityByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddCapacityByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AddCapacityByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AddCapacityByUserIdAsync(
	request *AddCapacityByUserIdRequest,
	callback chan<- AddCapacityByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/capacity"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.AddCapacityValue != nil {
        bodies["addCapacityValue"] = *request.AddCapacityValue
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

	go addCapacityByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AddCapacityByUserId(
	request *AddCapacityByUserIdRequest,
) (*AddCapacityByUserIdResult, error) {
	callback := make(chan AddCapacityByUserIdAsyncResult, 1)
	go p.AddCapacityByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setCapacityByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- SetCapacityByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetCapacityByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetCapacityByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetCapacityByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetCapacityByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetCapacityByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) SetCapacityByUserIdAsync(
	request *SetCapacityByUserIdRequest,
	callback chan<- SetCapacityByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/capacity"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.NewCapacityValue != nil {
        bodies["newCapacityValue"] = *request.NewCapacityValue
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

	go setCapacityByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) SetCapacityByUserId(
	request *SetCapacityByUserIdRequest,
) (*SetCapacityByUserIdResult, error) {
	callback := make(chan SetCapacityByUserIdAsyncResult, 1)
	go p.SetCapacityByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteInventoryByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteInventoryByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteInventoryByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteInventoryByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteInventoryByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteInventoryByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteInventoryByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteInventoryByUserIdAsync(
	request *DeleteInventoryByUserIdRequest,
	callback chan<- DeleteInventoryByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
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
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go deleteInventoryByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteInventoryByUserId(
	request *DeleteInventoryByUserIdRequest,
) (*DeleteInventoryByUserIdResult, error) {
	callback := make(chan DeleteInventoryByUserIdAsyncResult, 1)
	go p.DeleteInventoryByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addCapacityByStampSheetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AddCapacityByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddCapacityByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddCapacityByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AddCapacityByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddCapacityByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AddCapacityByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AddCapacityByStampSheetAsync(
	request *AddCapacityByStampSheetRequest,
	callback chan<- AddCapacityByStampSheetAsyncResult,
) {
	path := "/stamp/inventory/capacity/add"

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

	go addCapacityByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AddCapacityByStampSheet(
	request *AddCapacityByStampSheetRequest,
) (*AddCapacityByStampSheetResult, error) {
	callback := make(chan AddCapacityByStampSheetAsyncResult, 1)
	go p.AddCapacityByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setCapacityByStampSheetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- SetCapacityByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetCapacityByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetCapacityByStampSheetResult
	if asyncResult.Err != nil {
		callback <- SetCapacityByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetCapacityByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetCapacityByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) SetCapacityByStampSheetAsync(
	request *SetCapacityByStampSheetRequest,
	callback chan<- SetCapacityByStampSheetAsyncResult,
) {
	path := "/stamp/inventory/capacity/set"

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

	go setCapacityByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) SetCapacityByStampSheet(
	request *SetCapacityByStampSheetRequest,
) (*SetCapacityByStampSheetResult, error) {
	callback := make(chan SetCapacityByStampSheetAsyncResult, 1)
	go p.SetCapacityByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeItemSetsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeItemSetsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeItemSetsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeItemSetsResult
	if asyncResult.Err != nil {
		callback <- DescribeItemSetsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeItemSetsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeItemSetsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeItemSetsAsync(
	request *DescribeItemSetsRequest,
	callback chan<- DescribeItemSetsAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
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

	go describeItemSetsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeItemSets(
	request *DescribeItemSetsRequest,
) (*DescribeItemSetsResult, error) {
	callback := make(chan DescribeItemSetsAsyncResult, 1)
	go p.DescribeItemSetsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeItemSetsByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeItemSetsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeItemSetsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeItemSetsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeItemSetsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeItemSetsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeItemSetsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeItemSetsByUserIdAsync(
	request *DescribeItemSetsByUserIdRequest,
	callback chan<- DescribeItemSetsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
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

	go describeItemSetsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeItemSetsByUserId(
	request *DescribeItemSetsByUserIdRequest,
) (*DescribeItemSetsByUserIdResult, error) {
	callback := make(chan DescribeItemSetsByUserIdAsyncResult, 1)
	go p.DescribeItemSetsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getItemSetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetItemSetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetItemSetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetItemSetResult
	if asyncResult.Err != nil {
		callback <- GetItemSetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetItemSetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetItemSetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetItemSetAsync(
	request *GetItemSetRequest,
	callback chan<- GetItemSetAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ItemSetName != nil {
		queryStrings["itemSetName"] = core.ToString(*request.ItemSetName)
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go getItemSetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetItemSet(
	request *GetItemSetRequest,
) (*GetItemSetResult, error) {
	callback := make(chan GetItemSetAsyncResult, 1)
	go p.GetItemSetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getItemSetByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetItemSetByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetItemSetByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetItemSetByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetItemSetByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetItemSetByUserIdAsync(
	request *GetItemSetByUserIdRequest,
	callback chan<- GetItemSetByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ItemSetName != nil {
		queryStrings["itemSetName"] = core.ToString(*request.ItemSetName)
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getItemSetByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetItemSetByUserId(
	request *GetItemSetByUserIdRequest,
) (*GetItemSetByUserIdResult, error) {
	callback := make(chan GetItemSetByUserIdAsyncResult, 1)
	go p.GetItemSetByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getItemWithSignatureAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetItemWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetItemWithSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetItemWithSignatureResult
	if asyncResult.Err != nil {
		callback <- GetItemWithSignatureAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetItemWithSignatureAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetItemWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetItemWithSignatureAsync(
	request *GetItemWithSignatureRequest,
	callback chan<- GetItemWithSignatureAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/signature"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ItemSetName != nil {
		queryStrings["itemSetName"] = core.ToString(*request.ItemSetName)
	}
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

	go getItemWithSignatureAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetItemWithSignature(
	request *GetItemWithSignatureRequest,
) (*GetItemWithSignatureResult, error) {
	callback := make(chan GetItemWithSignatureAsyncResult, 1)
	go p.GetItemWithSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getItemWithSignatureByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetItemWithSignatureByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetItemWithSignatureByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetItemWithSignatureByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetItemWithSignatureByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetItemWithSignatureByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetItemWithSignatureByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetItemWithSignatureByUserIdAsync(
	request *GetItemWithSignatureByUserIdRequest,
	callback chan<- GetItemWithSignatureByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/signature"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ItemSetName != nil {
		queryStrings["itemSetName"] = core.ToString(*request.ItemSetName)
	}
	if request.KeyId != nil {
		queryStrings["keyId"] = core.ToString(*request.KeyId)
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getItemWithSignatureByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetItemWithSignatureByUserId(
	request *GetItemWithSignatureByUserIdRequest,
) (*GetItemWithSignatureByUserIdResult, error) {
	callback := make(chan GetItemWithSignatureByUserIdAsyncResult, 1)
	go p.GetItemWithSignatureByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireItemSetByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireItemSetByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireItemSetByUserIdResult
	if asyncResult.Err != nil {
		callback <- AcquireItemSetByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AcquireItemSetByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AcquireItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AcquireItemSetByUserIdAsync(
	request *AcquireItemSetByUserIdRequest,
	callback chan<- AcquireItemSetByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/acquire"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.AcquireCount != nil {
        bodies["acquireCount"] = *request.AcquireCount
    }
    if request.ExpiresAt != nil {
        bodies["expiresAt"] = *request.ExpiresAt
    }
    if request.CreateNewItemSet != nil {
        bodies["createNewItemSet"] = *request.CreateNewItemSet
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
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

	go acquireItemSetByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AcquireItemSetByUserId(
	request *AcquireItemSetByUserIdRequest,
) (*AcquireItemSetByUserIdResult, error) {
	callback := make(chan AcquireItemSetByUserIdAsyncResult, 1)
	go p.AcquireItemSetByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeItemSetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeItemSetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeItemSetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeItemSetResult
	if asyncResult.Err != nil {
		callback <- ConsumeItemSetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ConsumeItemSetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ConsumeItemSetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ConsumeItemSetAsync(
	request *ConsumeItemSetRequest,
	callback chan<- ConsumeItemSetAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/consume"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.ConsumeCount != nil {
        bodies["consumeCount"] = *request.ConsumeCount
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
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

	go consumeItemSetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ConsumeItemSet(
	request *ConsumeItemSetRequest,
) (*ConsumeItemSetResult, error) {
	callback := make(chan ConsumeItemSetAsyncResult, 1)
	go p.ConsumeItemSetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeItemSetByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeItemSetByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeItemSetByUserIdResult
	if asyncResult.Err != nil {
		callback <- ConsumeItemSetByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ConsumeItemSetByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ConsumeItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ConsumeItemSetByUserIdAsync(
	request *ConsumeItemSetByUserIdRequest,
	callback chan<- ConsumeItemSetByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/consume"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.ConsumeCount != nil {
        bodies["consumeCount"] = *request.ConsumeCount
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
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

	go consumeItemSetByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ConsumeItemSetByUserId(
	request *ConsumeItemSetByUserIdRequest,
) (*ConsumeItemSetByUserIdResult, error) {
	callback := make(chan ConsumeItemSetByUserIdAsyncResult, 1)
	go p.ConsumeItemSetByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteItemSetByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteItemSetByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteItemSetByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteItemSetByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteItemSetByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteItemSetByUserIdAsync(
	request *DeleteItemSetByUserIdRequest,
	callback chan<- DeleteItemSetByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ItemSetName != nil {
		queryStrings["itemSetName"] = core.ToString(*request.ItemSetName)
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go deleteItemSetByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteItemSetByUserId(
	request *DeleteItemSetByUserIdRequest,
) (*DeleteItemSetByUserIdResult, error) {
	callback := make(chan DeleteItemSetByUserIdAsyncResult, 1)
	go p.DeleteItemSetByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireItemSetByStampSheetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireItemSetByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireItemSetByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireItemSetByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AcquireItemSetByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AcquireItemSetByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AcquireItemSetByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AcquireItemSetByStampSheetAsync(
	request *AcquireItemSetByStampSheetRequest,
	callback chan<- AcquireItemSetByStampSheetAsyncResult,
) {
	path := "/stamp/item/acquire"

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

	go acquireItemSetByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AcquireItemSetByStampSheet(
	request *AcquireItemSetByStampSheetRequest,
) (*AcquireItemSetByStampSheetResult, error) {
	callback := make(chan AcquireItemSetByStampSheetAsyncResult, 1)
	go p.AcquireItemSetByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeItemSetByStampTaskAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeItemSetByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeItemSetByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeItemSetByStampTaskResult
	if asyncResult.Err != nil {
		callback <- ConsumeItemSetByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ConsumeItemSetByStampTaskAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ConsumeItemSetByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ConsumeItemSetByStampTaskAsync(
	request *ConsumeItemSetByStampTaskRequest,
	callback chan<- ConsumeItemSetByStampTaskAsyncResult,
) {
	path := "/stamp/item/consume"

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

	go consumeItemSetByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ConsumeItemSetByStampTask(
	request *ConsumeItemSetByStampTaskRequest,
) (*ConsumeItemSetByStampTaskResult, error) {
	callback := make(chan ConsumeItemSetByStampTaskAsyncResult, 1)
	go p.ConsumeItemSetByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeReferenceOfAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReferenceOfAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReferenceOfResult
	if asyncResult.Err != nil {
		callback <- DescribeReferenceOfAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeReferenceOfAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeReferenceOfAsync(
	request *DescribeReferenceOfRequest,
	callback chan<- DescribeReferenceOfAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }
    if request.ItemSetName != nil && *request.ItemSetName != ""  {
        path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
    } else {
        path = strings.ReplaceAll(path, "{itemSetName}", "null")
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

	go describeReferenceOfAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeReferenceOf(
	request *DescribeReferenceOfRequest,
) (*DescribeReferenceOfResult, error) {
	callback := make(chan DescribeReferenceOfAsyncResult, 1)
	go p.DescribeReferenceOfAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeReferenceOfByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReferenceOfByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReferenceOfByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeReferenceOfByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeReferenceOfByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeReferenceOfByUserIdAsync(
	request *DescribeReferenceOfByUserIdRequest,
	callback chan<- DescribeReferenceOfByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }
    if request.ItemSetName != nil && *request.ItemSetName != ""  {
        path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
    } else {
        path = strings.ReplaceAll(path, "{itemSetName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go describeReferenceOfByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeReferenceOfByUserId(
	request *DescribeReferenceOfByUserIdRequest,
) (*DescribeReferenceOfByUserIdResult, error) {
	callback := make(chan DescribeReferenceOfByUserIdAsyncResult, 1)
	go p.DescribeReferenceOfByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getReferenceOfAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReferenceOfAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReferenceOfResult
	if asyncResult.Err != nil {
		callback <- GetReferenceOfAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetReferenceOfAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetReferenceOfAsync(
	request *GetReferenceOfRequest,
	callback chan<- GetReferenceOfAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference/{referenceOf}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }
    if request.ItemSetName != nil && *request.ItemSetName != ""  {
        path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
    } else {
        path = strings.ReplaceAll(path, "{itemSetName}", "null")
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != ""  {
        path = strings.ReplaceAll(path, "{referenceOf}", core.ToString(*request.ReferenceOf))
    } else {
        path = strings.ReplaceAll(path, "{referenceOf}", "null")
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

	go getReferenceOfAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetReferenceOf(
	request *GetReferenceOfRequest,
) (*GetReferenceOfResult, error) {
	callback := make(chan GetReferenceOfAsyncResult, 1)
	go p.GetReferenceOfAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getReferenceOfByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReferenceOfByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReferenceOfByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetReferenceOfByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetReferenceOfByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetReferenceOfByUserIdAsync(
	request *GetReferenceOfByUserIdRequest,
	callback chan<- GetReferenceOfByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference/{referenceOf}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }
    if request.ItemSetName != nil && *request.ItemSetName != ""  {
        path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
    } else {
        path = strings.ReplaceAll(path, "{itemSetName}", "null")
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != ""  {
        path = strings.ReplaceAll(path, "{referenceOf}", core.ToString(*request.ReferenceOf))
    } else {
        path = strings.ReplaceAll(path, "{referenceOf}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getReferenceOfByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetReferenceOfByUserId(
	request *GetReferenceOfByUserIdRequest,
) (*GetReferenceOfByUserIdResult, error) {
	callback := make(chan GetReferenceOfByUserIdAsyncResult, 1)
	go p.GetReferenceOfByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyReferenceOfAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyReferenceOfAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyReferenceOfResult
	if asyncResult.Err != nil {
		callback <- VerifyReferenceOfAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- VerifyReferenceOfAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- VerifyReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyReferenceOfAsync(
	request *VerifyReferenceOfRequest,
	callback chan<- VerifyReferenceOfAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference/{referenceOf}/verify/{verifyType}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }
    if request.ItemSetName != nil && *request.ItemSetName != ""  {
        path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
    } else {
        path = strings.ReplaceAll(path, "{itemSetName}", "null")
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != ""  {
        path = strings.ReplaceAll(path, "{referenceOf}", core.ToString(*request.ReferenceOf))
    } else {
        path = strings.ReplaceAll(path, "{referenceOf}", "null")
    }
    if request.VerifyType != nil && *request.VerifyType != ""  {
        path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
    } else {
        path = strings.ReplaceAll(path, "{verifyType}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
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

	go verifyReferenceOfAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyReferenceOf(
	request *VerifyReferenceOfRequest,
) (*VerifyReferenceOfResult, error) {
	callback := make(chan VerifyReferenceOfAsyncResult, 1)
	go p.VerifyReferenceOfAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyReferenceOfByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyReferenceOfByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyReferenceOfByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyReferenceOfByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- VerifyReferenceOfByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- VerifyReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyReferenceOfByUserIdAsync(
	request *VerifyReferenceOfByUserIdRequest,
	callback chan<- VerifyReferenceOfByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference/{referenceOf}/verify/{verifyType}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }
    if request.ItemSetName != nil && *request.ItemSetName != ""  {
        path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
    } else {
        path = strings.ReplaceAll(path, "{itemSetName}", "null")
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != ""  {
        path = strings.ReplaceAll(path, "{referenceOf}", core.ToString(*request.ReferenceOf))
    } else {
        path = strings.ReplaceAll(path, "{referenceOf}", "null")
    }
    if request.VerifyType != nil && *request.VerifyType != ""  {
        path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
    } else {
        path = strings.ReplaceAll(path, "{verifyType}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
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

	go verifyReferenceOfByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyReferenceOfByUserId(
	request *VerifyReferenceOfByUserIdRequest,
) (*VerifyReferenceOfByUserIdResult, error) {
	callback := make(chan VerifyReferenceOfByUserIdAsyncResult, 1)
	go p.VerifyReferenceOfByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addReferenceOfAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AddReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddReferenceOfAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddReferenceOfResult
	if asyncResult.Err != nil {
		callback <- AddReferenceOfAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddReferenceOfAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AddReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AddReferenceOfAsync(
	request *AddReferenceOfRequest,
	callback chan<- AddReferenceOfAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }
    if request.ItemSetName != nil && *request.ItemSetName != ""  {
        path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
    } else {
        path = strings.ReplaceAll(path, "{itemSetName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.ReferenceOf != nil && *request.ReferenceOf != "" {
        bodies["referenceOf"] = *request.ReferenceOf
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

	go addReferenceOfAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AddReferenceOf(
	request *AddReferenceOfRequest,
) (*AddReferenceOfResult, error) {
	callback := make(chan AddReferenceOfAsyncResult, 1)
	go p.AddReferenceOfAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addReferenceOfByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AddReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddReferenceOfByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddReferenceOfByUserIdResult
	if asyncResult.Err != nil {
		callback <- AddReferenceOfByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddReferenceOfByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AddReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AddReferenceOfByUserIdAsync(
	request *AddReferenceOfByUserIdRequest,
	callback chan<- AddReferenceOfByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }
    if request.ItemSetName != nil && *request.ItemSetName != ""  {
        path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
    } else {
        path = strings.ReplaceAll(path, "{itemSetName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.ReferenceOf != nil && *request.ReferenceOf != "" {
        bodies["referenceOf"] = *request.ReferenceOf
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

	go addReferenceOfByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AddReferenceOfByUserId(
	request *AddReferenceOfByUserIdRequest,
) (*AddReferenceOfByUserIdResult, error) {
	callback := make(chan AddReferenceOfByUserIdAsyncResult, 1)
	go p.AddReferenceOfByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteReferenceOfAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteReferenceOfAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteReferenceOfResult
	if asyncResult.Err != nil {
		callback <- DeleteReferenceOfAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteReferenceOfAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteReferenceOfAsync(
	request *DeleteReferenceOfRequest,
	callback chan<- DeleteReferenceOfAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference/{referenceOf}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }
    if request.ItemSetName != nil && *request.ItemSetName != ""  {
        path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
    } else {
        path = strings.ReplaceAll(path, "{itemSetName}", "null")
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != ""  {
        path = strings.ReplaceAll(path, "{referenceOf}", core.ToString(*request.ReferenceOf))
    } else {
        path = strings.ReplaceAll(path, "{referenceOf}", "null")
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

	go deleteReferenceOfAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteReferenceOf(
	request *DeleteReferenceOfRequest,
) (*DeleteReferenceOfResult, error) {
	callback := make(chan DeleteReferenceOfAsyncResult, 1)
	go p.DeleteReferenceOfAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteReferenceOfByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteReferenceOfByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteReferenceOfByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteReferenceOfByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteReferenceOfByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteReferenceOfByUserIdAsync(
	request *DeleteReferenceOfByUserIdRequest,
	callback chan<- DeleteReferenceOfByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference/{referenceOf}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InventoryName != nil && *request.InventoryName != ""  {
        path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
    } else {
        path = strings.ReplaceAll(path, "{inventoryName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.ItemName != nil && *request.ItemName != ""  {
        path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
    } else {
        path = strings.ReplaceAll(path, "{itemName}", "null")
    }
    if request.ItemSetName != nil && *request.ItemSetName != ""  {
        path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
    } else {
        path = strings.ReplaceAll(path, "{itemSetName}", "null")
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != ""  {
        path = strings.ReplaceAll(path, "{referenceOf}", core.ToString(*request.ReferenceOf))
    } else {
        path = strings.ReplaceAll(path, "{referenceOf}", "null")
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

	go deleteReferenceOfByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteReferenceOfByUserId(
	request *DeleteReferenceOfByUserIdRequest,
) (*DeleteReferenceOfByUserIdResult, error) {
	callback := make(chan DeleteReferenceOfByUserIdAsyncResult, 1)
	go p.DeleteReferenceOfByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addReferenceOfItemSetByStampSheetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AddReferenceOfItemSetByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddReferenceOfItemSetByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddReferenceOfItemSetByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AddReferenceOfItemSetByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddReferenceOfItemSetByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AddReferenceOfItemSetByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AddReferenceOfItemSetByStampSheetAsync(
	request *AddReferenceOfItemSetByStampSheetRequest,
	callback chan<- AddReferenceOfItemSetByStampSheetAsyncResult,
) {
	path := "/stamp/item/reference/add"

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

	go addReferenceOfItemSetByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AddReferenceOfItemSetByStampSheet(
	request *AddReferenceOfItemSetByStampSheetRequest,
) (*AddReferenceOfItemSetByStampSheetResult, error) {
	callback := make(chan AddReferenceOfItemSetByStampSheetAsyncResult, 1)
	go p.AddReferenceOfItemSetByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteReferenceOfItemSetByStampSheetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteReferenceOfItemSetByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteReferenceOfItemSetByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteReferenceOfItemSetByStampSheetResult
	if asyncResult.Err != nil {
		callback <- DeleteReferenceOfItemSetByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteReferenceOfItemSetByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteReferenceOfItemSetByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteReferenceOfItemSetByStampSheetAsync(
	request *DeleteReferenceOfItemSetByStampSheetRequest,
	callback chan<- DeleteReferenceOfItemSetByStampSheetAsyncResult,
) {
	path := "/stamp/item/reference/delete"

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

	go deleteReferenceOfItemSetByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteReferenceOfItemSetByStampSheet(
	request *DeleteReferenceOfItemSetByStampSheetRequest,
) (*DeleteReferenceOfItemSetByStampSheetResult, error) {
	callback := make(chan DeleteReferenceOfItemSetByStampSheetAsyncResult, 1)
	go p.DeleteReferenceOfItemSetByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyReferenceOfByStampTaskAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyReferenceOfByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyReferenceOfByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyReferenceOfByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyReferenceOfByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- VerifyReferenceOfByStampTaskAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- VerifyReferenceOfByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyReferenceOfByStampTaskAsync(
	request *VerifyReferenceOfByStampTaskRequest,
	callback chan<- VerifyReferenceOfByStampTaskAsyncResult,
) {
	path := "/stamp/item/verify"

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

	go verifyReferenceOfByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyReferenceOfByStampTask(
	request *VerifyReferenceOfByStampTaskRequest,
) (*VerifyReferenceOfByStampTaskResult, error) {
	callback := make(chan VerifyReferenceOfByStampTaskAsyncResult, 1)
	go p.VerifyReferenceOfByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
