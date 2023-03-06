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
	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2InventoryWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2InventoryWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2InventoryWebSocketClient) describeNamespacesAsyncHandler(
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
    if asyncResult.Err != nil {
    }
	callback <- DescribeNamespacesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
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

func (p Gs2InventoryWebSocketClient) DescribeNamespaces(
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

func (p Gs2InventoryWebSocketClient) createNamespaceAsyncHandler(
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
    if asyncResult.Err != nil {
    }
	callback <- CreateNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
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

	go p.createNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) CreateNamespace(
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

func (p Gs2InventoryWebSocketClient) getNamespaceStatusAsyncHandler(
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
    if asyncResult.Err != nil {
    }
	callback <- GetNamespaceStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
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

func (p Gs2InventoryWebSocketClient) GetNamespaceStatus(
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

func (p Gs2InventoryWebSocketClient) getNamespaceAsyncHandler(
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
    if asyncResult.Err != nil {
    }
	callback <- GetNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
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

func (p Gs2InventoryWebSocketClient) GetNamespace(
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

func (p Gs2InventoryWebSocketClient) updateNamespaceAsyncHandler(
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
    if asyncResult.Err != nil {
    }
	callback <- UpdateNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
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

	go p.updateNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) UpdateNamespace(
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

func (p Gs2InventoryWebSocketClient) deleteNamespaceAsyncHandler(
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
    if asyncResult.Err != nil {
    }
	callback <- DeleteNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
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

func (p Gs2InventoryWebSocketClient) DeleteNamespace(
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

func (p Gs2InventoryWebSocketClient) describeInventoryModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeInventoryModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeInventoryModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeInventoryModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DescribeInventoryModelMastersAsync(
	request *DescribeInventoryModelMastersRequest,
	callback chan<- DescribeInventoryModelMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventoryModelMaster",
    		"function": "describeInventoryModelMasters",
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

	go p.describeInventoryModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DescribeInventoryModelMasters(
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

func (p Gs2InventoryWebSocketClient) createInventoryModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateInventoryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- CreateInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) CreateInventoryModelMasterAsync(
	request *CreateInventoryModelMasterRequest,
	callback chan<- CreateInventoryModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventoryModelMaster",
    		"function": "createInventoryModelMaster",
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

	go p.createInventoryModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) CreateInventoryModelMaster(
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

func (p Gs2InventoryWebSocketClient) getInventoryModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetInventoryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetInventoryModelMasterAsync(
	request *GetInventoryModelMasterRequest,
	callback chan<- GetInventoryModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventoryModelMaster",
    		"function": "getInventoryModelMaster",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getInventoryModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetInventoryModelMaster(
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

func (p Gs2InventoryWebSocketClient) updateInventoryModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateInventoryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- UpdateInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) UpdateInventoryModelMasterAsync(
	request *UpdateInventoryModelMasterRequest,
	callback chan<- UpdateInventoryModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventoryModelMaster",
    		"function": "updateInventoryModelMaster",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
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

	go p.updateInventoryModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) UpdateInventoryModelMaster(
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

func (p Gs2InventoryWebSocketClient) deleteInventoryModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteInventoryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DeleteInventoryModelMasterAsync(
	request *DeleteInventoryModelMasterRequest,
	callback chan<- DeleteInventoryModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventoryModelMaster",
    		"function": "deleteInventoryModelMaster",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteInventoryModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DeleteInventoryModelMaster(
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

func (p Gs2InventoryWebSocketClient) describeInventoryModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeInventoryModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeInventoryModelsAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeInventoryModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DescribeInventoryModelsAsync(
	request *DescribeInventoryModelsRequest,
	callback chan<- DescribeInventoryModelsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventoryModel",
    		"function": "describeInventoryModels",
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

	go p.describeInventoryModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DescribeInventoryModels(
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

func (p Gs2InventoryWebSocketClient) getInventoryModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetInventoryModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetInventoryModelAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetInventoryModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetInventoryModelAsync(
	request *GetInventoryModelRequest,
	callback chan<- GetInventoryModelAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventoryModel",
    		"function": "getInventoryModel",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getInventoryModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetInventoryModel(
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

func (p Gs2InventoryWebSocketClient) describeItemModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeItemModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeItemModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeItemModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DescribeItemModelMastersAsync(
	request *DescribeItemModelMastersRequest,
	callback chan<- DescribeItemModelMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemModelMaster",
    		"function": "describeItemModelMasters",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
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

	go p.describeItemModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DescribeItemModelMasters(
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

func (p Gs2InventoryWebSocketClient) createItemModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateItemModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- CreateItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) CreateItemModelMasterAsync(
	request *CreateItemModelMasterRequest,
	callback chan<- CreateItemModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemModelMaster",
    		"function": "createItemModelMaster",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
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

	go p.createItemModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) CreateItemModelMaster(
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

func (p Gs2InventoryWebSocketClient) getItemModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetItemModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetItemModelMasterAsync(
	request *GetItemModelMasterRequest,
	callback chan<- GetItemModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemModelMaster",
    		"function": "getItemModelMaster",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getItemModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetItemModelMaster(
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

func (p Gs2InventoryWebSocketClient) updateItemModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateItemModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- UpdateItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) UpdateItemModelMasterAsync(
	request *UpdateItemModelMasterRequest,
	callback chan<- UpdateItemModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemModelMaster",
    		"function": "updateItemModelMaster",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
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

	go p.updateItemModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) UpdateItemModelMaster(
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

func (p Gs2InventoryWebSocketClient) deleteItemModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteItemModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DeleteItemModelMasterAsync(
	request *DeleteItemModelMasterRequest,
	callback chan<- DeleteItemModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemModelMaster",
    		"function": "deleteItemModelMaster",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteItemModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DeleteItemModelMaster(
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

func (p Gs2InventoryWebSocketClient) describeItemModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeItemModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeItemModelsAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeItemModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DescribeItemModelsAsync(
	request *DescribeItemModelsRequest,
	callback chan<- DescribeItemModelsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemModel",
    		"function": "describeItemModels",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeItemModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DescribeItemModels(
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

func (p Gs2InventoryWebSocketClient) getItemModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetItemModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetItemModelAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetItemModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetItemModelAsync(
	request *GetItemModelRequest,
	callback chan<- GetItemModelAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemModel",
    		"function": "getItemModel",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getItemModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetItemModel(
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

func (p Gs2InventoryWebSocketClient) exportMasterAsyncHandler(
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
    if asyncResult.Err != nil {
    }
	callback <- ExportMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "currentItemModelMaster",
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

func (p Gs2InventoryWebSocketClient) ExportMaster(
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

func (p Gs2InventoryWebSocketClient) getCurrentItemModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentItemModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetCurrentItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetCurrentItemModelMasterAsync(
	request *GetCurrentItemModelMasterRequest,
	callback chan<- GetCurrentItemModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "currentItemModelMaster",
    		"function": "getCurrentItemModelMaster",
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

	go p.getCurrentItemModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetCurrentItemModelMaster(
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

func (p Gs2InventoryWebSocketClient) updateCurrentItemModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentItemModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- UpdateCurrentItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) UpdateCurrentItemModelMasterAsync(
	request *UpdateCurrentItemModelMasterRequest,
	callback chan<- UpdateCurrentItemModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "currentItemModelMaster",
    		"function": "updateCurrentItemModelMaster",
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

	go p.updateCurrentItemModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) UpdateCurrentItemModelMaster(
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

func (p Gs2InventoryWebSocketClient) updateCurrentItemModelMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentItemModelMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentItemModelMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- UpdateCurrentItemModelMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) UpdateCurrentItemModelMasterFromGitHubAsync(
	request *UpdateCurrentItemModelMasterFromGitHubRequest,
	callback chan<- UpdateCurrentItemModelMasterFromGitHubAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "currentItemModelMaster",
    		"function": "updateCurrentItemModelMasterFromGitHub",
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

	go p.updateCurrentItemModelMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) UpdateCurrentItemModelMasterFromGitHub(
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

func (p Gs2InventoryWebSocketClient) describeInventoriesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeInventoriesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeInventoriesAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeInventoriesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DescribeInventoriesAsync(
	request *DescribeInventoriesRequest,
	callback chan<- DescribeInventoriesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventory",
    		"function": "describeInventories",
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
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.describeInventoriesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DescribeInventories(
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

func (p Gs2InventoryWebSocketClient) describeInventoriesByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeInventoriesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeInventoriesByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeInventoriesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DescribeInventoriesByUserIdAsync(
	request *DescribeInventoriesByUserIdRequest,
	callback chan<- DescribeInventoriesByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventory",
    		"function": "describeInventoriesByUserId",
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
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeInventoriesByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DescribeInventoriesByUserId(
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

func (p Gs2InventoryWebSocketClient) getInventoryAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetInventoryAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetInventoryAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetInventoryAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetInventoryAsync(
	request *GetInventoryRequest,
	callback chan<- GetInventoryAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventory",
    		"function": "getInventory",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
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

	go p.getInventoryAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetInventory(
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

func (p Gs2InventoryWebSocketClient) getInventoryByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetInventoryByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetInventoryByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetInventoryByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetInventoryByUserIdAsync(
	request *GetInventoryByUserIdRequest,
	callback chan<- GetInventoryByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventory",
    		"function": "getInventoryByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getInventoryByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetInventoryByUserId(
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

func (p Gs2InventoryWebSocketClient) addCapacityByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddCapacityByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddCapacityByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- AddCapacityByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) AddCapacityByUserIdAsync(
	request *AddCapacityByUserIdRequest,
	callback chan<- AddCapacityByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventory",
    		"function": "addCapacityByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.AddCapacityValue != nil {
        bodies["addCapacityValue"] = *request.AddCapacityValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.addCapacityByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) AddCapacityByUserId(
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

func (p Gs2InventoryWebSocketClient) setCapacityByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetCapacityByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetCapacityByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- SetCapacityByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) SetCapacityByUserIdAsync(
	request *SetCapacityByUserIdRequest,
	callback chan<- SetCapacityByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventory",
    		"function": "setCapacityByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.NewCapacityValue != nil {
        bodies["newCapacityValue"] = *request.NewCapacityValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.setCapacityByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) SetCapacityByUserId(
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

func (p Gs2InventoryWebSocketClient) deleteInventoryByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteInventoryByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteInventoryByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteInventoryByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DeleteInventoryByUserIdAsync(
	request *DeleteInventoryByUserIdRequest,
	callback chan<- DeleteInventoryByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventory",
    		"function": "deleteInventoryByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.deleteInventoryByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DeleteInventoryByUserId(
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

func (p Gs2InventoryWebSocketClient) addCapacityByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddCapacityByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddCapacityByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- AddCapacityByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) AddCapacityByStampSheetAsync(
	request *AddCapacityByStampSheetRequest,
	callback chan<- AddCapacityByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventory",
    		"function": "addCapacityByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.StampSheet != nil && *request.StampSheet != "" {
        bodies["stampSheet"] = *request.StampSheet
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.addCapacityByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) AddCapacityByStampSheet(
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

func (p Gs2InventoryWebSocketClient) setCapacityByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetCapacityByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetCapacityByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- SetCapacityByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) SetCapacityByStampSheetAsync(
	request *SetCapacityByStampSheetRequest,
	callback chan<- SetCapacityByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "inventory",
    		"function": "setCapacityByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.StampSheet != nil && *request.StampSheet != "" {
        bodies["stampSheet"] = *request.StampSheet
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.setCapacityByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) SetCapacityByStampSheet(
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

func (p Gs2InventoryWebSocketClient) describeItemSetsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeItemSetsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeItemSetsAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeItemSetsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DescribeItemSetsAsync(
	request *DescribeItemSetsRequest,
	callback chan<- DescribeItemSetsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemSet",
    		"function": "describeItemSets",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
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
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.describeItemSetsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DescribeItemSets(
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

func (p Gs2InventoryWebSocketClient) describeItemSetsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeItemSetsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeItemSetsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeItemSetsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DescribeItemSetsByUserIdAsync(
	request *DescribeItemSetsByUserIdRequest,
	callback chan<- DescribeItemSetsByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemSet",
    		"function": "describeItemSetsByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
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

	go p.describeItemSetsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DescribeItemSetsByUserId(
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

func (p Gs2InventoryWebSocketClient) getItemSetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetItemSetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetItemSetAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetItemSetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetItemSetAsync(
	request *GetItemSetRequest,
	callback chan<- GetItemSetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemSet",
    		"function": "getItemSet",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getItemSetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetItemSet(
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

func (p Gs2InventoryWebSocketClient) getItemSetByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetItemSetByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetItemSetByUserIdAsync(
	request *GetItemSetByUserIdRequest,
	callback chan<- GetItemSetByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemSet",
    		"function": "getItemSetByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getItemSetByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetItemSetByUserId(
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

func (p Gs2InventoryWebSocketClient) getItemWithSignatureAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetItemWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetItemWithSignatureAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetItemWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetItemWithSignatureAsync(
	request *GetItemWithSignatureRequest,
	callback chan<- GetItemWithSignatureAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemSet",
    		"function": "getItemWithSignature",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getItemWithSignatureAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetItemWithSignature(
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

func (p Gs2InventoryWebSocketClient) getItemWithSignatureByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetItemWithSignatureByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetItemWithSignatureByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetItemWithSignatureByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetItemWithSignatureByUserIdAsync(
	request *GetItemWithSignatureByUserIdRequest,
	callback chan<- GetItemWithSignatureByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemSet",
    		"function": "getItemWithSignatureByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getItemWithSignatureByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetItemWithSignatureByUserId(
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

func (p Gs2InventoryWebSocketClient) acquireItemSetByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AcquireItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AcquireItemSetByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
        gs2err, ok := asyncResult.Err.(core.Gs2Exception)
        if ok {
            if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
            }
        }
    }
	callback <- AcquireItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) AcquireItemSetByUserIdAsync(
	request *AcquireItemSetByUserIdRequest,
	callback chan<- AcquireItemSetByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemSet",
    		"function": "acquireItemSetByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
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
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.acquireItemSetByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) AcquireItemSetByUserId(
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

func (p Gs2InventoryWebSocketClient) consumeItemSetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ConsumeItemSetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ConsumeItemSetAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
        gs2err, ok := asyncResult.Err.(core.Gs2Exception)
        if ok {
            if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
            }
            if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.count.insufficient" {
				asyncResult.Err = gs2err.SetClientError(Insufficient{})
            }
        }
    }
	callback <- ConsumeItemSetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) ConsumeItemSetAsync(
	request *ConsumeItemSetRequest,
	callback chan<- ConsumeItemSetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemSet",
    		"function": "consumeItemSet",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ConsumeCount != nil {
        bodies["consumeCount"] = *request.ConsumeCount
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.consumeItemSetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) ConsumeItemSet(
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

func (p Gs2InventoryWebSocketClient) consumeItemSetByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ConsumeItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ConsumeItemSetByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
        gs2err, ok := asyncResult.Err.(core.Gs2Exception)
        if ok {
            if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
            }
            if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.count.insufficient" {
				asyncResult.Err = gs2err.SetClientError(Insufficient{})
            }
        }
    }
	callback <- ConsumeItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) ConsumeItemSetByUserIdAsync(
	request *ConsumeItemSetByUserIdRequest,
	callback chan<- ConsumeItemSetByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemSet",
    		"function": "consumeItemSetByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ConsumeCount != nil {
        bodies["consumeCount"] = *request.ConsumeCount
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.consumeItemSetByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) ConsumeItemSetByUserId(
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

func (p Gs2InventoryWebSocketClient) deleteItemSetByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteItemSetByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DeleteItemSetByUserIdAsync(
	request *DeleteItemSetByUserIdRequest,
	callback chan<- DeleteItemSetByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemSet",
    		"function": "deleteItemSetByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.deleteItemSetByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DeleteItemSetByUserId(
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

func (p Gs2InventoryWebSocketClient) acquireItemSetByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AcquireItemSetByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AcquireItemSetByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- AcquireItemSetByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) AcquireItemSetByStampSheetAsync(
	request *AcquireItemSetByStampSheetRequest,
	callback chan<- AcquireItemSetByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemSet",
    		"function": "acquireItemSetByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.StampSheet != nil && *request.StampSheet != "" {
        bodies["stampSheet"] = *request.StampSheet
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.acquireItemSetByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) AcquireItemSetByStampSheet(
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

func (p Gs2InventoryWebSocketClient) consumeItemSetByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ConsumeItemSetByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ConsumeItemSetByStampTaskAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- ConsumeItemSetByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) ConsumeItemSetByStampTaskAsync(
	request *ConsumeItemSetByStampTaskRequest,
	callback chan<- ConsumeItemSetByStampTaskAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "itemSet",
    		"function": "consumeItemSetByStampTask",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.StampTask != nil && *request.StampTask != "" {
        bodies["stampTask"] = *request.StampTask
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.consumeItemSetByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) ConsumeItemSetByStampTask(
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

func (p Gs2InventoryWebSocketClient) describeReferenceOfAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeReferenceOfAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DescribeReferenceOfAsync(
	request *DescribeReferenceOfRequest,
	callback chan<- DescribeReferenceOfAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "describeReferenceOf",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.describeReferenceOfAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DescribeReferenceOf(
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

func (p Gs2InventoryWebSocketClient) describeReferenceOfByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeReferenceOfByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DescribeReferenceOfByUserIdAsync(
	request *DescribeReferenceOfByUserIdRequest,
	callback chan<- DescribeReferenceOfByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "describeReferenceOfByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeReferenceOfByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DescribeReferenceOfByUserId(
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

func (p Gs2InventoryWebSocketClient) getReferenceOfAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetReferenceOfAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetReferenceOfAsync(
	request *GetReferenceOfRequest,
	callback chan<- GetReferenceOfAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "getReferenceOf",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != "" {
        bodies["referenceOf"] = *request.ReferenceOf
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getReferenceOfAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetReferenceOf(
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

func (p Gs2InventoryWebSocketClient) getReferenceOfByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetReferenceOfByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) GetReferenceOfByUserIdAsync(
	request *GetReferenceOfByUserIdRequest,
	callback chan<- GetReferenceOfByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "getReferenceOfByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != "" {
        bodies["referenceOf"] = *request.ReferenceOf
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getReferenceOfByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) GetReferenceOfByUserId(
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

func (p Gs2InventoryWebSocketClient) verifyReferenceOfAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- VerifyReferenceOfAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- VerifyReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) VerifyReferenceOfAsync(
	request *VerifyReferenceOfRequest,
	callback chan<- VerifyReferenceOfAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "verifyReferenceOf",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != "" {
        bodies["referenceOf"] = *request.ReferenceOf
    }
    if request.VerifyType != nil && *request.VerifyType != "" {
        bodies["verifyType"] = *request.VerifyType
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.verifyReferenceOfAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) VerifyReferenceOf(
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

func (p Gs2InventoryWebSocketClient) verifyReferenceOfByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- VerifyReferenceOfByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- VerifyReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) VerifyReferenceOfByUserIdAsync(
	request *VerifyReferenceOfByUserIdRequest,
	callback chan<- VerifyReferenceOfByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "verifyReferenceOfByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != "" {
        bodies["referenceOf"] = *request.ReferenceOf
    }
    if request.VerifyType != nil && *request.VerifyType != "" {
        bodies["verifyType"] = *request.VerifyType
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.verifyReferenceOfByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) VerifyReferenceOfByUserId(
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

func (p Gs2InventoryWebSocketClient) addReferenceOfAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddReferenceOfAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- AddReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) AddReferenceOfAsync(
	request *AddReferenceOfRequest,
	callback chan<- AddReferenceOfAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "addReferenceOf",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != "" {
        bodies["referenceOf"] = *request.ReferenceOf
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.addReferenceOfAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) AddReferenceOf(
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

func (p Gs2InventoryWebSocketClient) addReferenceOfByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddReferenceOfByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- AddReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) AddReferenceOfByUserIdAsync(
	request *AddReferenceOfByUserIdRequest,
	callback chan<- AddReferenceOfByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "addReferenceOfByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != "" {
        bodies["referenceOf"] = *request.ReferenceOf
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.addReferenceOfByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) AddReferenceOfByUserId(
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

func (p Gs2InventoryWebSocketClient) deleteReferenceOfAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteReferenceOfAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DeleteReferenceOfAsync(
	request *DeleteReferenceOfRequest,
	callback chan<- DeleteReferenceOfAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "deleteReferenceOf",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != "" {
        bodies["referenceOf"] = *request.ReferenceOf
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.deleteReferenceOfAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DeleteReferenceOf(
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

func (p Gs2InventoryWebSocketClient) deleteReferenceOfByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteReferenceOfByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DeleteReferenceOfByUserIdAsync(
	request *DeleteReferenceOfByUserIdRequest,
	callback chan<- DeleteReferenceOfByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "deleteReferenceOfByUserId",
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
    if request.InventoryName != nil && *request.InventoryName != "" {
        bodies["inventoryName"] = *request.InventoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ItemName != nil && *request.ItemName != "" {
        bodies["itemName"] = *request.ItemName
    }
    if request.ItemSetName != nil && *request.ItemSetName != "" {
        bodies["itemSetName"] = *request.ItemSetName
    }
    if request.ReferenceOf != nil && *request.ReferenceOf != "" {
        bodies["referenceOf"] = *request.ReferenceOf
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.deleteReferenceOfByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DeleteReferenceOfByUserId(
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

func (p Gs2InventoryWebSocketClient) addReferenceOfItemSetByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddReferenceOfItemSetByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddReferenceOfItemSetByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- AddReferenceOfItemSetByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) AddReferenceOfItemSetByStampSheetAsync(
	request *AddReferenceOfItemSetByStampSheetRequest,
	callback chan<- AddReferenceOfItemSetByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "addReferenceOfItemSetByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.StampSheet != nil && *request.StampSheet != "" {
        bodies["stampSheet"] = *request.StampSheet
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.addReferenceOfItemSetByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) AddReferenceOfItemSetByStampSheet(
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

func (p Gs2InventoryWebSocketClient) deleteReferenceOfItemSetByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteReferenceOfItemSetByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteReferenceOfItemSetByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteReferenceOfItemSetByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) DeleteReferenceOfItemSetByStampSheetAsync(
	request *DeleteReferenceOfItemSetByStampSheetRequest,
	callback chan<- DeleteReferenceOfItemSetByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "deleteReferenceOfItemSetByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.StampSheet != nil && *request.StampSheet != "" {
        bodies["stampSheet"] = *request.StampSheet
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteReferenceOfItemSetByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) DeleteReferenceOfItemSetByStampSheet(
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

func (p Gs2InventoryWebSocketClient) verifyReferenceOfByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyReferenceOfByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- VerifyReferenceOfByStampTaskAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- VerifyReferenceOfByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryWebSocketClient) VerifyReferenceOfByStampTaskAsync(
	request *VerifyReferenceOfByStampTaskRequest,
	callback chan<- VerifyReferenceOfByStampTaskAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inventory",
    		"component": "referenceOf",
    		"function": "verifyReferenceOfByStampTask",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.StampTask != nil && *request.StampTask != "" {
        bodies["stampTask"] = *request.StampTask
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.verifyReferenceOfByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InventoryWebSocketClient) VerifyReferenceOfByStampTask(
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
