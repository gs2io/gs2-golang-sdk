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

package enchant

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2EnchantWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2EnchantWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2EnchantWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2EnchantWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
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

func (p Gs2EnchantWebSocketClient) DescribeNamespaces(
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

func (p Gs2EnchantWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2EnchantWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
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

func (p Gs2EnchantWebSocketClient) CreateNamespace(
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

func (p Gs2EnchantWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2EnchantWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
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

func (p Gs2EnchantWebSocketClient) GetNamespaceStatus(
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

func (p Gs2EnchantWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2EnchantWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
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

func (p Gs2EnchantWebSocketClient) GetNamespace(
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

func (p Gs2EnchantWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2EnchantWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
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

func (p Gs2EnchantWebSocketClient) UpdateNamespace(
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

func (p Gs2EnchantWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2EnchantWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
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

func (p Gs2EnchantWebSocketClient) DeleteNamespace(
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

func (p Gs2EnchantWebSocketClient) describeBalanceParameterModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeBalanceParameterModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBalanceParameterModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBalanceParameterModelsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeBalanceParameterModelsAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeBalanceParameterModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) DescribeBalanceParameterModelsAsync(
	request *DescribeBalanceParameterModelsRequest,
	callback chan<- DescribeBalanceParameterModelsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterModel",
    		"function": "describeBalanceParameterModels",
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

	go p.describeBalanceParameterModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) DescribeBalanceParameterModels(
	request *DescribeBalanceParameterModelsRequest,
) (*DescribeBalanceParameterModelsResult, error) {
	callback := make(chan DescribeBalanceParameterModelsAsyncResult, 1)
	go p.DescribeBalanceParameterModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) getBalanceParameterModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetBalanceParameterModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBalanceParameterModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBalanceParameterModelResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetBalanceParameterModelAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetBalanceParameterModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) GetBalanceParameterModelAsync(
	request *GetBalanceParameterModelRequest,
	callback chan<- GetBalanceParameterModelAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterModel",
    		"function": "getBalanceParameterModel",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getBalanceParameterModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) GetBalanceParameterModel(
	request *GetBalanceParameterModelRequest,
) (*GetBalanceParameterModelResult, error) {
	callback := make(chan GetBalanceParameterModelAsyncResult, 1)
	go p.GetBalanceParameterModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) describeBalanceParameterModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeBalanceParameterModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBalanceParameterModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBalanceParameterModelMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeBalanceParameterModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeBalanceParameterModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) DescribeBalanceParameterModelMastersAsync(
	request *DescribeBalanceParameterModelMastersRequest,
	callback chan<- DescribeBalanceParameterModelMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterModelMaster",
    		"function": "describeBalanceParameterModelMasters",
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

	go p.describeBalanceParameterModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) DescribeBalanceParameterModelMasters(
	request *DescribeBalanceParameterModelMastersRequest,
) (*DescribeBalanceParameterModelMastersResult, error) {
	callback := make(chan DescribeBalanceParameterModelMastersAsyncResult, 1)
	go p.DescribeBalanceParameterModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) createBalanceParameterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateBalanceParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateBalanceParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateBalanceParameterModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateBalanceParameterModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- CreateBalanceParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) CreateBalanceParameterModelMasterAsync(
	request *CreateBalanceParameterModelMasterRequest,
	callback chan<- CreateBalanceParameterModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterModelMaster",
    		"function": "createBalanceParameterModelMaster",
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
    if request.TotalValue != nil {
        bodies["totalValue"] = *request.TotalValue
    }
    if request.InitialValueStrategy != nil && *request.InitialValueStrategy != "" {
        bodies["initialValueStrategy"] = *request.InitialValueStrategy
    }
    if request.Parameters != nil {
        var _parameters []interface {}
        for _, item := range request.Parameters {
            _parameters = append(_parameters, item)
        }
        bodies["parameters"] = _parameters
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createBalanceParameterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) CreateBalanceParameterModelMaster(
	request *CreateBalanceParameterModelMasterRequest,
) (*CreateBalanceParameterModelMasterResult, error) {
	callback := make(chan CreateBalanceParameterModelMasterAsyncResult, 1)
	go p.CreateBalanceParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) getBalanceParameterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetBalanceParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBalanceParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBalanceParameterModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetBalanceParameterModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetBalanceParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) GetBalanceParameterModelMasterAsync(
	request *GetBalanceParameterModelMasterRequest,
	callback chan<- GetBalanceParameterModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterModelMaster",
    		"function": "getBalanceParameterModelMaster",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getBalanceParameterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) GetBalanceParameterModelMaster(
	request *GetBalanceParameterModelMasterRequest,
) (*GetBalanceParameterModelMasterResult, error) {
	callback := make(chan GetBalanceParameterModelMasterAsyncResult, 1)
	go p.GetBalanceParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) updateBalanceParameterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateBalanceParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateBalanceParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateBalanceParameterModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateBalanceParameterModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- UpdateBalanceParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) UpdateBalanceParameterModelMasterAsync(
	request *UpdateBalanceParameterModelMasterRequest,
	callback chan<- UpdateBalanceParameterModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterModelMaster",
    		"function": "updateBalanceParameterModelMaster",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.TotalValue != nil {
        bodies["totalValue"] = *request.TotalValue
    }
    if request.InitialValueStrategy != nil && *request.InitialValueStrategy != "" {
        bodies["initialValueStrategy"] = *request.InitialValueStrategy
    }
    if request.Parameters != nil {
        var _parameters []interface {}
        for _, item := range request.Parameters {
            _parameters = append(_parameters, item)
        }
        bodies["parameters"] = _parameters
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateBalanceParameterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) UpdateBalanceParameterModelMaster(
	request *UpdateBalanceParameterModelMasterRequest,
) (*UpdateBalanceParameterModelMasterResult, error) {
	callback := make(chan UpdateBalanceParameterModelMasterAsyncResult, 1)
	go p.UpdateBalanceParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) deleteBalanceParameterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteBalanceParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteBalanceParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteBalanceParameterModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteBalanceParameterModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteBalanceParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) DeleteBalanceParameterModelMasterAsync(
	request *DeleteBalanceParameterModelMasterRequest,
	callback chan<- DeleteBalanceParameterModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterModelMaster",
    		"function": "deleteBalanceParameterModelMaster",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteBalanceParameterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) DeleteBalanceParameterModelMaster(
	request *DeleteBalanceParameterModelMasterRequest,
) (*DeleteBalanceParameterModelMasterResult, error) {
	callback := make(chan DeleteBalanceParameterModelMasterAsyncResult, 1)
	go p.DeleteBalanceParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) describeRarityParameterModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRarityParameterModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRarityParameterModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRarityParameterModelsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRarityParameterModelsAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeRarityParameterModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) DescribeRarityParameterModelsAsync(
	request *DescribeRarityParameterModelsRequest,
	callback chan<- DescribeRarityParameterModelsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterModel",
    		"function": "describeRarityParameterModels",
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

	go p.describeRarityParameterModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) DescribeRarityParameterModels(
	request *DescribeRarityParameterModelsRequest,
) (*DescribeRarityParameterModelsResult, error) {
	callback := make(chan DescribeRarityParameterModelsAsyncResult, 1)
	go p.DescribeRarityParameterModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) getRarityParameterModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRarityParameterModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRarityParameterModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRarityParameterModelResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRarityParameterModelAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetRarityParameterModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) GetRarityParameterModelAsync(
	request *GetRarityParameterModelRequest,
	callback chan<- GetRarityParameterModelAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterModel",
    		"function": "getRarityParameterModel",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getRarityParameterModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) GetRarityParameterModel(
	request *GetRarityParameterModelRequest,
) (*GetRarityParameterModelResult, error) {
	callback := make(chan GetRarityParameterModelAsyncResult, 1)
	go p.GetRarityParameterModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) describeRarityParameterModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRarityParameterModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRarityParameterModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRarityParameterModelMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRarityParameterModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeRarityParameterModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) DescribeRarityParameterModelMastersAsync(
	request *DescribeRarityParameterModelMastersRequest,
	callback chan<- DescribeRarityParameterModelMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterModelMaster",
    		"function": "describeRarityParameterModelMasters",
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

	go p.describeRarityParameterModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) DescribeRarityParameterModelMasters(
	request *DescribeRarityParameterModelMastersRequest,
) (*DescribeRarityParameterModelMastersResult, error) {
	callback := make(chan DescribeRarityParameterModelMastersAsyncResult, 1)
	go p.DescribeRarityParameterModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) createRarityParameterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateRarityParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateRarityParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateRarityParameterModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateRarityParameterModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- CreateRarityParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) CreateRarityParameterModelMasterAsync(
	request *CreateRarityParameterModelMasterRequest,
	callback chan<- CreateRarityParameterModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterModelMaster",
    		"function": "createRarityParameterModelMaster",
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
    if request.MaximumParameterCount != nil {
        bodies["maximumParameterCount"] = *request.MaximumParameterCount
    }
    if request.ParameterCounts != nil {
        var _parameterCounts []interface {}
        for _, item := range request.ParameterCounts {
            _parameterCounts = append(_parameterCounts, item)
        }
        bodies["parameterCounts"] = _parameterCounts
    }
    if request.Parameters != nil {
        var _parameters []interface {}
        for _, item := range request.Parameters {
            _parameters = append(_parameters, item)
        }
        bodies["parameters"] = _parameters
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createRarityParameterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) CreateRarityParameterModelMaster(
	request *CreateRarityParameterModelMasterRequest,
) (*CreateRarityParameterModelMasterResult, error) {
	callback := make(chan CreateRarityParameterModelMasterAsyncResult, 1)
	go p.CreateRarityParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) getRarityParameterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRarityParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRarityParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRarityParameterModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRarityParameterModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetRarityParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) GetRarityParameterModelMasterAsync(
	request *GetRarityParameterModelMasterRequest,
	callback chan<- GetRarityParameterModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterModelMaster",
    		"function": "getRarityParameterModelMaster",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getRarityParameterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) GetRarityParameterModelMaster(
	request *GetRarityParameterModelMasterRequest,
) (*GetRarityParameterModelMasterResult, error) {
	callback := make(chan GetRarityParameterModelMasterAsyncResult, 1)
	go p.GetRarityParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) updateRarityParameterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateRarityParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateRarityParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateRarityParameterModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateRarityParameterModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- UpdateRarityParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) UpdateRarityParameterModelMasterAsync(
	request *UpdateRarityParameterModelMasterRequest,
	callback chan<- UpdateRarityParameterModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterModelMaster",
    		"function": "updateRarityParameterModelMaster",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.MaximumParameterCount != nil {
        bodies["maximumParameterCount"] = *request.MaximumParameterCount
    }
    if request.ParameterCounts != nil {
        var _parameterCounts []interface {}
        for _, item := range request.ParameterCounts {
            _parameterCounts = append(_parameterCounts, item)
        }
        bodies["parameterCounts"] = _parameterCounts
    }
    if request.Parameters != nil {
        var _parameters []interface {}
        for _, item := range request.Parameters {
            _parameters = append(_parameters, item)
        }
        bodies["parameters"] = _parameters
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateRarityParameterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) UpdateRarityParameterModelMaster(
	request *UpdateRarityParameterModelMasterRequest,
) (*UpdateRarityParameterModelMasterResult, error) {
	callback := make(chan UpdateRarityParameterModelMasterAsyncResult, 1)
	go p.UpdateRarityParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) deleteRarityParameterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteRarityParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRarityParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRarityParameterModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteRarityParameterModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteRarityParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) DeleteRarityParameterModelMasterAsync(
	request *DeleteRarityParameterModelMasterRequest,
	callback chan<- DeleteRarityParameterModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterModelMaster",
    		"function": "deleteRarityParameterModelMaster",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteRarityParameterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) DeleteRarityParameterModelMaster(
	request *DeleteRarityParameterModelMasterRequest,
) (*DeleteRarityParameterModelMasterResult, error) {
	callback := make(chan DeleteRarityParameterModelMasterAsyncResult, 1)
	go p.DeleteRarityParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) exportMasterAsyncHandler(
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

func (p Gs2EnchantWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "currentParameterMaster",
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

func (p Gs2EnchantWebSocketClient) ExportMaster(
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

func (p Gs2EnchantWebSocketClient) getCurrentParameterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentParameterModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentParameterModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetCurrentParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) GetCurrentParameterModelMasterAsync(
	request *GetCurrentParameterModelMasterRequest,
	callback chan<- GetCurrentParameterModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "currentParameterMaster",
    		"function": "getCurrentParameterModelMaster",
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

	go p.getCurrentParameterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) GetCurrentParameterModelMaster(
	request *GetCurrentParameterModelMasterRequest,
) (*GetCurrentParameterModelMasterResult, error) {
	callback := make(chan GetCurrentParameterModelMasterAsyncResult, 1)
	go p.GetCurrentParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) updateCurrentParameterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentParameterModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentParameterModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- UpdateCurrentParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) UpdateCurrentParameterModelMasterAsync(
	request *UpdateCurrentParameterModelMasterRequest,
	callback chan<- UpdateCurrentParameterModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "currentParameterMaster",
    		"function": "updateCurrentParameterModelMaster",
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

	go p.updateCurrentParameterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) UpdateCurrentParameterModelMaster(
	request *UpdateCurrentParameterModelMasterRequest,
) (*UpdateCurrentParameterModelMasterResult, error) {
	callback := make(chan UpdateCurrentParameterModelMasterAsyncResult, 1)
	go p.UpdateCurrentParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) updateCurrentParameterModelMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentParameterModelMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentParameterModelMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentParameterModelMasterFromGitHubResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentParameterModelMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- UpdateCurrentParameterModelMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) UpdateCurrentParameterModelMasterFromGitHubAsync(
	request *UpdateCurrentParameterModelMasterFromGitHubRequest,
	callback chan<- UpdateCurrentParameterModelMasterFromGitHubAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "currentParameterMaster",
    		"function": "updateCurrentParameterModelMasterFromGitHub",
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

	go p.updateCurrentParameterModelMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) UpdateCurrentParameterModelMasterFromGitHub(
	request *UpdateCurrentParameterModelMasterFromGitHubRequest,
) (*UpdateCurrentParameterModelMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentParameterModelMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentParameterModelMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) describeBalanceParameterStatusesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeBalanceParameterStatusesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBalanceParameterStatusesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBalanceParameterStatusesResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeBalanceParameterStatusesAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeBalanceParameterStatusesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) DescribeBalanceParameterStatusesAsync(
	request *DescribeBalanceParameterStatusesRequest,
	callback chan<- DescribeBalanceParameterStatusesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterStatus",
    		"function": "describeBalanceParameterStatuses",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
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

	go p.describeBalanceParameterStatusesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) DescribeBalanceParameterStatuses(
	request *DescribeBalanceParameterStatusesRequest,
) (*DescribeBalanceParameterStatusesResult, error) {
	callback := make(chan DescribeBalanceParameterStatusesAsyncResult, 1)
	go p.DescribeBalanceParameterStatusesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) describeBalanceParameterStatusesByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeBalanceParameterStatusesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBalanceParameterStatusesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBalanceParameterStatusesByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeBalanceParameterStatusesByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeBalanceParameterStatusesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) DescribeBalanceParameterStatusesByUserIdAsync(
	request *DescribeBalanceParameterStatusesByUserIdRequest,
	callback chan<- DescribeBalanceParameterStatusesByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterStatus",
    		"function": "describeBalanceParameterStatusesByUserId",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
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

	go p.describeBalanceParameterStatusesByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) DescribeBalanceParameterStatusesByUserId(
	request *DescribeBalanceParameterStatusesByUserIdRequest,
) (*DescribeBalanceParameterStatusesByUserIdResult, error) {
	callback := make(chan DescribeBalanceParameterStatusesByUserIdAsyncResult, 1)
	go p.DescribeBalanceParameterStatusesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) getBalanceParameterStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetBalanceParameterStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBalanceParameterStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBalanceParameterStatusResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetBalanceParameterStatusAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetBalanceParameterStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) GetBalanceParameterStatusAsync(
	request *GetBalanceParameterStatusRequest,
	callback chan<- GetBalanceParameterStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterStatus",
    		"function": "getBalanceParameterStatus",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getBalanceParameterStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) GetBalanceParameterStatus(
	request *GetBalanceParameterStatusRequest,
) (*GetBalanceParameterStatusResult, error) {
	callback := make(chan GetBalanceParameterStatusAsyncResult, 1)
	go p.GetBalanceParameterStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) getBalanceParameterStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetBalanceParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBalanceParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBalanceParameterStatusByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetBalanceParameterStatusByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetBalanceParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) GetBalanceParameterStatusByUserIdAsync(
	request *GetBalanceParameterStatusByUserIdRequest,
	callback chan<- GetBalanceParameterStatusByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterStatus",
    		"function": "getBalanceParameterStatusByUserId",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getBalanceParameterStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) GetBalanceParameterStatusByUserId(
	request *GetBalanceParameterStatusByUserIdRequest,
) (*GetBalanceParameterStatusByUserIdResult, error) {
	callback := make(chan GetBalanceParameterStatusByUserIdAsyncResult, 1)
	go p.GetBalanceParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) deleteBalanceParameterStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteBalanceParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteBalanceParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteBalanceParameterStatusByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteBalanceParameterStatusByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteBalanceParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) DeleteBalanceParameterStatusByUserIdAsync(
	request *DeleteBalanceParameterStatusByUserIdRequest,
	callback chan<- DeleteBalanceParameterStatusByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterStatus",
    		"function": "deleteBalanceParameterStatusByUserId",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.deleteBalanceParameterStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) DeleteBalanceParameterStatusByUserId(
	request *DeleteBalanceParameterStatusByUserIdRequest,
) (*DeleteBalanceParameterStatusByUserIdResult, error) {
	callback := make(chan DeleteBalanceParameterStatusByUserIdAsyncResult, 1)
	go p.DeleteBalanceParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) reDrawBalanceParameterStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReDrawBalanceParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReDrawBalanceParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReDrawBalanceParameterStatusByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ReDrawBalanceParameterStatusByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- ReDrawBalanceParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) ReDrawBalanceParameterStatusByUserIdAsync(
	request *ReDrawBalanceParameterStatusByUserIdRequest,
	callback chan<- ReDrawBalanceParameterStatusByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterStatus",
    		"function": "reDrawBalanceParameterStatusByUserId",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
    if request.FixedParameterNames != nil {
        var _fixedParameterNames []interface {}
        for _, item := range request.FixedParameterNames {
            _fixedParameterNames = append(_fixedParameterNames, item)
        }
        bodies["fixedParameterNames"] = _fixedParameterNames
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.reDrawBalanceParameterStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) ReDrawBalanceParameterStatusByUserId(
	request *ReDrawBalanceParameterStatusByUserIdRequest,
) (*ReDrawBalanceParameterStatusByUserIdResult, error) {
	callback := make(chan ReDrawBalanceParameterStatusByUserIdAsyncResult, 1)
	go p.ReDrawBalanceParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) reDrawBalanceParameterStatusByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReDrawBalanceParameterStatusByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReDrawBalanceParameterStatusByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReDrawBalanceParameterStatusByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ReDrawBalanceParameterStatusByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- ReDrawBalanceParameterStatusByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) ReDrawBalanceParameterStatusByStampSheetAsync(
	request *ReDrawBalanceParameterStatusByStampSheetRequest,
	callback chan<- ReDrawBalanceParameterStatusByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "balanceParameterStatus",
    		"function": "reDrawBalanceParameterStatusByStampSheet",
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

	go p.reDrawBalanceParameterStatusByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) ReDrawBalanceParameterStatusByStampSheet(
	request *ReDrawBalanceParameterStatusByStampSheetRequest,
) (*ReDrawBalanceParameterStatusByStampSheetResult, error) {
	callback := make(chan ReDrawBalanceParameterStatusByStampSheetAsyncResult, 1)
	go p.ReDrawBalanceParameterStatusByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) describeRarityParameterStatusesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRarityParameterStatusesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRarityParameterStatusesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRarityParameterStatusesResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRarityParameterStatusesAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeRarityParameterStatusesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) DescribeRarityParameterStatusesAsync(
	request *DescribeRarityParameterStatusesRequest,
	callback chan<- DescribeRarityParameterStatusesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterStatus",
    		"function": "describeRarityParameterStatuses",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
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

	go p.describeRarityParameterStatusesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) DescribeRarityParameterStatuses(
	request *DescribeRarityParameterStatusesRequest,
) (*DescribeRarityParameterStatusesResult, error) {
	callback := make(chan DescribeRarityParameterStatusesAsyncResult, 1)
	go p.DescribeRarityParameterStatusesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) describeRarityParameterStatusesByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRarityParameterStatusesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRarityParameterStatusesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRarityParameterStatusesByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRarityParameterStatusesByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeRarityParameterStatusesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) DescribeRarityParameterStatusesByUserIdAsync(
	request *DescribeRarityParameterStatusesByUserIdRequest,
	callback chan<- DescribeRarityParameterStatusesByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterStatus",
    		"function": "describeRarityParameterStatusesByUserId",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
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

	go p.describeRarityParameterStatusesByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) DescribeRarityParameterStatusesByUserId(
	request *DescribeRarityParameterStatusesByUserIdRequest,
) (*DescribeRarityParameterStatusesByUserIdResult, error) {
	callback := make(chan DescribeRarityParameterStatusesByUserIdAsyncResult, 1)
	go p.DescribeRarityParameterStatusesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) getRarityParameterStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRarityParameterStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRarityParameterStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRarityParameterStatusResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRarityParameterStatusAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetRarityParameterStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) GetRarityParameterStatusAsync(
	request *GetRarityParameterStatusRequest,
	callback chan<- GetRarityParameterStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterStatus",
    		"function": "getRarityParameterStatus",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getRarityParameterStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) GetRarityParameterStatus(
	request *GetRarityParameterStatusRequest,
) (*GetRarityParameterStatusResult, error) {
	callback := make(chan GetRarityParameterStatusAsyncResult, 1)
	go p.GetRarityParameterStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) getRarityParameterStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRarityParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRarityParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRarityParameterStatusByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRarityParameterStatusByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetRarityParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) GetRarityParameterStatusByUserIdAsync(
	request *GetRarityParameterStatusByUserIdRequest,
	callback chan<- GetRarityParameterStatusByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterStatus",
    		"function": "getRarityParameterStatusByUserId",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getRarityParameterStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) GetRarityParameterStatusByUserId(
	request *GetRarityParameterStatusByUserIdRequest,
) (*GetRarityParameterStatusByUserIdResult, error) {
	callback := make(chan GetRarityParameterStatusByUserIdAsyncResult, 1)
	go p.GetRarityParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) deleteRarityParameterStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteRarityParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRarityParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRarityParameterStatusByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteRarityParameterStatusByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteRarityParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) DeleteRarityParameterStatusByUserIdAsync(
	request *DeleteRarityParameterStatusByUserIdRequest,
	callback chan<- DeleteRarityParameterStatusByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterStatus",
    		"function": "deleteRarityParameterStatusByUserId",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.deleteRarityParameterStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) DeleteRarityParameterStatusByUserId(
	request *DeleteRarityParameterStatusByUserIdRequest,
) (*DeleteRarityParameterStatusByUserIdResult, error) {
	callback := make(chan DeleteRarityParameterStatusByUserIdAsyncResult, 1)
	go p.DeleteRarityParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) reDrawRarityParameterStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReDrawRarityParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReDrawRarityParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReDrawRarityParameterStatusByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ReDrawRarityParameterStatusByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- ReDrawRarityParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) ReDrawRarityParameterStatusByUserIdAsync(
	request *ReDrawRarityParameterStatusByUserIdRequest,
	callback chan<- ReDrawRarityParameterStatusByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterStatus",
    		"function": "reDrawRarityParameterStatusByUserId",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
    if request.FixedParameterNames != nil {
        var _fixedParameterNames []interface {}
        for _, item := range request.FixedParameterNames {
            _fixedParameterNames = append(_fixedParameterNames, item)
        }
        bodies["fixedParameterNames"] = _fixedParameterNames
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.reDrawRarityParameterStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) ReDrawRarityParameterStatusByUserId(
	request *ReDrawRarityParameterStatusByUserIdRequest,
) (*ReDrawRarityParameterStatusByUserIdResult, error) {
	callback := make(chan ReDrawRarityParameterStatusByUserIdAsyncResult, 1)
	go p.ReDrawRarityParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) reDrawRarityParameterStatusByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReDrawRarityParameterStatusByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReDrawRarityParameterStatusByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReDrawRarityParameterStatusByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ReDrawRarityParameterStatusByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- ReDrawRarityParameterStatusByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) ReDrawRarityParameterStatusByStampSheetAsync(
	request *ReDrawRarityParameterStatusByStampSheetRequest,
	callback chan<- ReDrawRarityParameterStatusByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterStatus",
    		"function": "reDrawRarityParameterStatusByStampSheet",
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

	go p.reDrawRarityParameterStatusByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) ReDrawRarityParameterStatusByStampSheet(
	request *ReDrawRarityParameterStatusByStampSheetRequest,
) (*ReDrawRarityParameterStatusByStampSheetResult, error) {
	callback := make(chan ReDrawRarityParameterStatusByStampSheetAsyncResult, 1)
	go p.ReDrawRarityParameterStatusByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) addRarityParameterStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddRarityParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddRarityParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddRarityParameterStatusByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddRarityParameterStatusByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- AddRarityParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) AddRarityParameterStatusByUserIdAsync(
	request *AddRarityParameterStatusByUserIdRequest,
	callback chan<- AddRarityParameterStatusByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterStatus",
    		"function": "addRarityParameterStatusByUserId",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
    if request.Count != nil {
        bodies["count"] = *request.Count
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.addRarityParameterStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) AddRarityParameterStatusByUserId(
	request *AddRarityParameterStatusByUserIdRequest,
) (*AddRarityParameterStatusByUserIdResult, error) {
	callback := make(chan AddRarityParameterStatusByUserIdAsyncResult, 1)
	go p.AddRarityParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) addRarityParameterStatusByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddRarityParameterStatusByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddRarityParameterStatusByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddRarityParameterStatusByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddRarityParameterStatusByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- AddRarityParameterStatusByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) AddRarityParameterStatusByStampSheetAsync(
	request *AddRarityParameterStatusByStampSheetRequest,
	callback chan<- AddRarityParameterStatusByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterStatus",
    		"function": "addRarityParameterStatusByStampSheet",
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

	go p.addRarityParameterStatusByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) AddRarityParameterStatusByStampSheet(
	request *AddRarityParameterStatusByStampSheetRequest,
) (*AddRarityParameterStatusByStampSheetResult, error) {
	callback := make(chan AddRarityParameterStatusByStampSheetAsyncResult, 1)
	go p.AddRarityParameterStatusByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) verifyRarityParameterStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyRarityParameterStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyRarityParameterStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyRarityParameterStatusResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- VerifyRarityParameterStatusAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- VerifyRarityParameterStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) VerifyRarityParameterStatusAsync(
	request *VerifyRarityParameterStatusRequest,
	callback chan<- VerifyRarityParameterStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterStatus",
    		"function": "verifyRarityParameterStatus",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
    if request.VerifyType != nil && *request.VerifyType != "" {
        bodies["verifyType"] = *request.VerifyType
    }
    if request.ParameterValueName != nil && *request.ParameterValueName != "" {
        bodies["parameterValueName"] = *request.ParameterValueName
    }
    if request.ParameterCount != nil {
        bodies["parameterCount"] = *request.ParameterCount
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

	go p.verifyRarityParameterStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) VerifyRarityParameterStatus(
	request *VerifyRarityParameterStatusRequest,
) (*VerifyRarityParameterStatusResult, error) {
	callback := make(chan VerifyRarityParameterStatusAsyncResult, 1)
	go p.VerifyRarityParameterStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) verifyRarityParameterStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyRarityParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyRarityParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyRarityParameterStatusByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- VerifyRarityParameterStatusByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- VerifyRarityParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) VerifyRarityParameterStatusByUserIdAsync(
	request *VerifyRarityParameterStatusByUserIdRequest,
	callback chan<- VerifyRarityParameterStatusByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterStatus",
    		"function": "verifyRarityParameterStatusByUserId",
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
    if request.ParameterName != nil && *request.ParameterName != "" {
        bodies["parameterName"] = *request.ParameterName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
    if request.VerifyType != nil && *request.VerifyType != "" {
        bodies["verifyType"] = *request.VerifyType
    }
    if request.ParameterValueName != nil && *request.ParameterValueName != "" {
        bodies["parameterValueName"] = *request.ParameterValueName
    }
    if request.ParameterCount != nil {
        bodies["parameterCount"] = *request.ParameterCount
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.verifyRarityParameterStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) VerifyRarityParameterStatusByUserId(
	request *VerifyRarityParameterStatusByUserIdRequest,
) (*VerifyRarityParameterStatusByUserIdResult, error) {
	callback := make(chan VerifyRarityParameterStatusByUserIdAsyncResult, 1)
	go p.VerifyRarityParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2EnchantWebSocketClient) verifyRarityParameterStatusByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyRarityParameterStatusByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyRarityParameterStatusByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyRarityParameterStatusByStampTaskResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- VerifyRarityParameterStatusByStampTaskAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- VerifyRarityParameterStatusByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantWebSocketClient) VerifyRarityParameterStatusByStampTaskAsync(
	request *VerifyRarityParameterStatusByStampTaskRequest,
	callback chan<- VerifyRarityParameterStatusByStampTaskAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "enchant",
    		"component": "rarityParameterStatus",
    		"function": "verifyRarityParameterStatusByStampTask",
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

	go p.verifyRarityParameterStatusByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2EnchantWebSocketClient) VerifyRarityParameterStatusByStampTask(
	request *VerifyRarityParameterStatusByStampTaskRequest,
) (*VerifyRarityParameterStatusByStampTaskResult, error) {
	callback := make(chan VerifyRarityParameterStatusByStampTaskAsyncResult, 1)
	go p.VerifyRarityParameterStatusByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
