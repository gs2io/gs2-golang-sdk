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
	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2ShowcaseWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2ShowcaseWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2ShowcaseWebSocketClient) describeNamespacesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "namespace",
    		"function": "describeNamespaces",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeNamespacesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) DescribeNamespaces(
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

func (p Gs2ShowcaseWebSocketClient) createNamespaceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "namespace",
    		"function": "createNamespace",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
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

	go p.createNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) CreateNamespace(
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

func (p Gs2ShowcaseWebSocketClient) getNamespaceStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "namespace",
    		"function": "getNamespaceStatus",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getNamespaceStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) GetNamespaceStatus(
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

func (p Gs2ShowcaseWebSocketClient) getNamespaceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "namespace",
    		"function": "getNamespace",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) GetNamespace(
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

func (p Gs2ShowcaseWebSocketClient) updateNamespaceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "namespace",
    		"function": "updateNamespace",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
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

	go p.updateNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) UpdateNamespace(
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

func (p Gs2ShowcaseWebSocketClient) deleteNamespaceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "namespace",
    		"function": "deleteNamespace",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) DeleteNamespace(
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

func (p Gs2ShowcaseWebSocketClient) describeSalesItemMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSalesItemMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) DescribeSalesItemMastersAsync(
	request *DescribeSalesItemMastersRequest,
	callback chan<- DescribeSalesItemMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "salesItemMaster",
    		"function": "describeSalesItemMasters",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeSalesItemMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) DescribeSalesItemMasters(
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

func (p Gs2ShowcaseWebSocketClient) createSalesItemMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateSalesItemMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) CreateSalesItemMasterAsync(
	request *CreateSalesItemMasterRequest,
	callback chan<- CreateSalesItemMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "salesItemMaster",
    		"function": "createSalesItemMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
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

	go p.createSalesItemMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) CreateSalesItemMaster(
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

func (p Gs2ShowcaseWebSocketClient) getSalesItemMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSalesItemMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) GetSalesItemMasterAsync(
	request *GetSalesItemMasterRequest,
	callback chan<- GetSalesItemMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "salesItemMaster",
    		"function": "getSalesItemMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.SalesItemName != nil && *request.SalesItemName != "" {
        bodies["salesItemName"] = *request.SalesItemName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getSalesItemMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) GetSalesItemMaster(
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

func (p Gs2ShowcaseWebSocketClient) updateSalesItemMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateSalesItemMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) UpdateSalesItemMasterAsync(
	request *UpdateSalesItemMasterRequest,
	callback chan<- UpdateSalesItemMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "salesItemMaster",
    		"function": "updateSalesItemMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.SalesItemName != nil && *request.SalesItemName != "" {
        bodies["salesItemName"] = *request.SalesItemName
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

	go p.updateSalesItemMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) UpdateSalesItemMaster(
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

func (p Gs2ShowcaseWebSocketClient) deleteSalesItemMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteSalesItemMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) DeleteSalesItemMasterAsync(
	request *DeleteSalesItemMasterRequest,
	callback chan<- DeleteSalesItemMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "salesItemMaster",
    		"function": "deleteSalesItemMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.SalesItemName != nil && *request.SalesItemName != "" {
        bodies["salesItemName"] = *request.SalesItemName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteSalesItemMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) DeleteSalesItemMaster(
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

func (p Gs2ShowcaseWebSocketClient) describeSalesItemGroupMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSalesItemGroupMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) DescribeSalesItemGroupMastersAsync(
	request *DescribeSalesItemGroupMastersRequest,
	callback chan<- DescribeSalesItemGroupMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "salesItemGroupMaster",
    		"function": "describeSalesItemGroupMasters",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeSalesItemGroupMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) DescribeSalesItemGroupMasters(
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

func (p Gs2ShowcaseWebSocketClient) createSalesItemGroupMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateSalesItemGroupMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) CreateSalesItemGroupMasterAsync(
	request *CreateSalesItemGroupMasterRequest,
	callback chan<- CreateSalesItemGroupMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "salesItemGroupMaster",
    		"function": "createSalesItemGroupMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
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

	go p.createSalesItemGroupMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) CreateSalesItemGroupMaster(
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

func (p Gs2ShowcaseWebSocketClient) getSalesItemGroupMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSalesItemGroupMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) GetSalesItemGroupMasterAsync(
	request *GetSalesItemGroupMasterRequest,
	callback chan<- GetSalesItemGroupMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "salesItemGroupMaster",
    		"function": "getSalesItemGroupMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.SalesItemGroupName != nil && *request.SalesItemGroupName != "" {
        bodies["salesItemGroupName"] = *request.SalesItemGroupName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getSalesItemGroupMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) GetSalesItemGroupMaster(
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

func (p Gs2ShowcaseWebSocketClient) updateSalesItemGroupMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateSalesItemGroupMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) UpdateSalesItemGroupMasterAsync(
	request *UpdateSalesItemGroupMasterRequest,
	callback chan<- UpdateSalesItemGroupMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "salesItemGroupMaster",
    		"function": "updateSalesItemGroupMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.SalesItemGroupName != nil && *request.SalesItemGroupName != "" {
        bodies["salesItemGroupName"] = *request.SalesItemGroupName
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

	go p.updateSalesItemGroupMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) UpdateSalesItemGroupMaster(
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

func (p Gs2ShowcaseWebSocketClient) deleteSalesItemGroupMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteSalesItemGroupMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) DeleteSalesItemGroupMasterAsync(
	request *DeleteSalesItemGroupMasterRequest,
	callback chan<- DeleteSalesItemGroupMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "salesItemGroupMaster",
    		"function": "deleteSalesItemGroupMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.SalesItemGroupName != nil && *request.SalesItemGroupName != "" {
        bodies["salesItemGroupName"] = *request.SalesItemGroupName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteSalesItemGroupMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) DeleteSalesItemGroupMaster(
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

func (p Gs2ShowcaseWebSocketClient) describeShowcaseMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeShowcaseMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) DescribeShowcaseMastersAsync(
	request *DescribeShowcaseMastersRequest,
	callback chan<- DescribeShowcaseMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "showcaseMaster",
    		"function": "describeShowcaseMasters",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeShowcaseMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) DescribeShowcaseMasters(
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

func (p Gs2ShowcaseWebSocketClient) createShowcaseMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateShowcaseMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) CreateShowcaseMasterAsync(
	request *CreateShowcaseMasterRequest,
	callback chan<- CreateShowcaseMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "showcaseMaster",
    		"function": "createShowcaseMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
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

	go p.createShowcaseMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) CreateShowcaseMaster(
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

func (p Gs2ShowcaseWebSocketClient) getShowcaseMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetShowcaseMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) GetShowcaseMasterAsync(
	request *GetShowcaseMasterRequest,
	callback chan<- GetShowcaseMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "showcaseMaster",
    		"function": "getShowcaseMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != "" {
        bodies["showcaseName"] = *request.ShowcaseName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getShowcaseMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) GetShowcaseMaster(
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

func (p Gs2ShowcaseWebSocketClient) updateShowcaseMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateShowcaseMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) UpdateShowcaseMasterAsync(
	request *UpdateShowcaseMasterRequest,
	callback chan<- UpdateShowcaseMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "showcaseMaster",
    		"function": "updateShowcaseMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != "" {
        bodies["showcaseName"] = *request.ShowcaseName
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

	go p.updateShowcaseMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) UpdateShowcaseMaster(
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

func (p Gs2ShowcaseWebSocketClient) deleteShowcaseMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteShowcaseMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) DeleteShowcaseMasterAsync(
	request *DeleteShowcaseMasterRequest,
	callback chan<- DeleteShowcaseMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "showcaseMaster",
    		"function": "deleteShowcaseMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != "" {
        bodies["showcaseName"] = *request.ShowcaseName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteShowcaseMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) DeleteShowcaseMaster(
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

func (p Gs2ShowcaseWebSocketClient) exportMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ExportMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "currentShowcaseMaster",
    		"function": "exportMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.exportMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) ExportMaster(
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

func (p Gs2ShowcaseWebSocketClient) getCurrentShowcaseMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentShowcaseMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) GetCurrentShowcaseMasterAsync(
	request *GetCurrentShowcaseMasterRequest,
	callback chan<- GetCurrentShowcaseMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "currentShowcaseMaster",
    		"function": "getCurrentShowcaseMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getCurrentShowcaseMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) GetCurrentShowcaseMaster(
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

func (p Gs2ShowcaseWebSocketClient) updateCurrentShowcaseMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentShowcaseMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) UpdateCurrentShowcaseMasterAsync(
	request *UpdateCurrentShowcaseMasterRequest,
	callback chan<- UpdateCurrentShowcaseMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "currentShowcaseMaster",
    		"function": "updateCurrentShowcaseMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.Settings != nil && *request.Settings != "" {
        bodies["settings"] = *request.Settings
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateCurrentShowcaseMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) UpdateCurrentShowcaseMaster(
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

func (p Gs2ShowcaseWebSocketClient) updateCurrentShowcaseMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentShowcaseMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) UpdateCurrentShowcaseMasterFromGitHubAsync(
	request *UpdateCurrentShowcaseMasterFromGitHubRequest,
	callback chan<- UpdateCurrentShowcaseMasterFromGitHubAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "currentShowcaseMaster",
    		"function": "updateCurrentShowcaseMasterFromGitHub",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CheckoutSetting != nil {
        bodies["checkoutSetting"] = request.CheckoutSetting.ToDict()
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateCurrentShowcaseMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) UpdateCurrentShowcaseMasterFromGitHub(
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

func (p Gs2ShowcaseWebSocketClient) describeShowcasesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeShowcasesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) DescribeShowcasesAsync(
	request *DescribeShowcasesRequest,
	callback chan<- DescribeShowcasesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "showcase",
    		"function": "describeShowcases",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.describeShowcasesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) DescribeShowcases(
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

func (p Gs2ShowcaseWebSocketClient) describeShowcasesByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeShowcasesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) DescribeShowcasesByUserIdAsync(
	request *DescribeShowcasesByUserIdRequest,
	callback chan<- DescribeShowcasesByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "showcase",
    		"function": "describeShowcasesByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeShowcasesByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) DescribeShowcasesByUserId(
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

func (p Gs2ShowcaseWebSocketClient) getShowcaseAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetShowcaseAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) GetShowcaseAsync(
	request *GetShowcaseRequest,
	callback chan<- GetShowcaseAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "showcase",
    		"function": "getShowcase",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != "" {
        bodies["showcaseName"] = *request.ShowcaseName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getShowcaseAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) GetShowcase(
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

func (p Gs2ShowcaseWebSocketClient) getShowcaseByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetShowcaseByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) GetShowcaseByUserIdAsync(
	request *GetShowcaseByUserIdRequest,
	callback chan<- GetShowcaseByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "showcase",
    		"function": "getShowcaseByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != "" {
        bodies["showcaseName"] = *request.ShowcaseName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getShowcaseByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) GetShowcaseByUserId(
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

func (p Gs2ShowcaseWebSocketClient) buyAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- BuyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) BuyAsync(
	request *BuyRequest,
	callback chan<- BuyAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "showcase",
    		"function": "buy",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != "" {
        bodies["showcaseName"] = *request.ShowcaseName
    }
    if request.DisplayItemId != nil && *request.DisplayItemId != "" {
        bodies["displayItemId"] = *request.DisplayItemId
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
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
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.buyAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) Buy(
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

func (p Gs2ShowcaseWebSocketClient) buyByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- BuyByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ShowcaseWebSocketClient) BuyByUserIdAsync(
	request *BuyByUserIdRequest,
	callback chan<- BuyByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "showcase",
    		"component": "showcase",
    		"function": "buyByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.ShowcaseName != nil && *request.ShowcaseName != "" {
        bodies["showcaseName"] = *request.ShowcaseName
    }
    if request.DisplayItemId != nil && *request.DisplayItemId != "" {
        bodies["displayItemId"] = *request.DisplayItemId
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
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

	go p.buyByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ShowcaseWebSocketClient) BuyByUserId(
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
