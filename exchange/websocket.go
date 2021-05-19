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
	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2ExchangeWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2ExchangeWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2ExchangeWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2ExchangeWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
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

func (p Gs2ExchangeWebSocketClient) DescribeNamespaces(
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

func (p Gs2ExchangeWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2ExchangeWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
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

func (p Gs2ExchangeWebSocketClient) CreateNamespace(
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

func (p Gs2ExchangeWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2ExchangeWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
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

func (p Gs2ExchangeWebSocketClient) GetNamespaceStatus(
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

func (p Gs2ExchangeWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2ExchangeWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
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

func (p Gs2ExchangeWebSocketClient) GetNamespace(
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

func (p Gs2ExchangeWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2ExchangeWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
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

func (p Gs2ExchangeWebSocketClient) UpdateNamespace(
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

func (p Gs2ExchangeWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2ExchangeWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
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

func (p Gs2ExchangeWebSocketClient) DeleteNamespace(
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

func (p Gs2ExchangeWebSocketClient) describeRateModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRateModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) DescribeRateModelsAsync(
	request *DescribeRateModelsRequest,
	callback chan<- DescribeRateModelsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "rateModel",
    		"function": "describeRateModels",
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

	go p.describeRateModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) DescribeRateModels(
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

func (p Gs2ExchangeWebSocketClient) getRateModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRateModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) GetRateModelAsync(
	request *GetRateModelRequest,
	callback chan<- GetRateModelAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "rateModel",
    		"function": "getRateModel",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getRateModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) GetRateModel(
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

func (p Gs2ExchangeWebSocketClient) describeRateModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRateModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) DescribeRateModelMastersAsync(
	request *DescribeRateModelMastersRequest,
	callback chan<- DescribeRateModelMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "rateModelMaster",
    		"function": "describeRateModelMasters",
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

	go p.describeRateModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) DescribeRateModelMasters(
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

func (p Gs2ExchangeWebSocketClient) createRateModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateRateModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) CreateRateModelMasterAsync(
	request *CreateRateModelMasterRequest,
	callback chan<- CreateRateModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "rateModelMaster",
    		"function": "createRateModelMaster",
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
        var _skipConsumeActions []*map[string]interface {}
        for _, item := range request.SkipConsumeActions {
            _skipConsumeActions = append(_skipConsumeActions, item.ToDict())
        }
        bodies["skipConsumeActions"] = _skipConsumeActions
    }
    if request.AcquireActions != nil {
        var _acquireActions []*map[string]interface {}
        for _, item := range request.AcquireActions {
            _acquireActions = append(_acquireActions, item.ToDict())
        }
        bodies["acquireActions"] = _acquireActions
    }
    if request.ConsumeActions != nil {
        var _consumeActions []*map[string]interface {}
        for _, item := range request.ConsumeActions {
            _consumeActions = append(_consumeActions, item.ToDict())
        }
        bodies["consumeActions"] = _consumeActions
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createRateModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) CreateRateModelMaster(
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

func (p Gs2ExchangeWebSocketClient) getRateModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRateModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) GetRateModelMasterAsync(
	request *GetRateModelMasterRequest,
	callback chan<- GetRateModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "rateModelMaster",
    		"function": "getRateModelMaster",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getRateModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) GetRateModelMaster(
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

func (p Gs2ExchangeWebSocketClient) updateRateModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateRateModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) UpdateRateModelMasterAsync(
	request *UpdateRateModelMasterRequest,
	callback chan<- UpdateRateModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "rateModelMaster",
    		"function": "updateRateModelMaster",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
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
        var _skipConsumeActions []*map[string]interface {}
        for _, item := range request.SkipConsumeActions {
            _skipConsumeActions = append(_skipConsumeActions, item.ToDict())
        }
        bodies["skipConsumeActions"] = _skipConsumeActions
    }
    if request.AcquireActions != nil {
        var _acquireActions []*map[string]interface {}
        for _, item := range request.AcquireActions {
            _acquireActions = append(_acquireActions, item.ToDict())
        }
        bodies["acquireActions"] = _acquireActions
    }
    if request.ConsumeActions != nil {
        var _consumeActions []*map[string]interface {}
        for _, item := range request.ConsumeActions {
            _consumeActions = append(_consumeActions, item.ToDict())
        }
        bodies["consumeActions"] = _consumeActions
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateRateModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) UpdateRateModelMaster(
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

func (p Gs2ExchangeWebSocketClient) deleteRateModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteRateModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) DeleteRateModelMasterAsync(
	request *DeleteRateModelMasterRequest,
	callback chan<- DeleteRateModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "rateModelMaster",
    		"function": "deleteRateModelMaster",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteRateModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) DeleteRateModelMaster(
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

func (p Gs2ExchangeWebSocketClient) exchangeAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ExchangeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) ExchangeAsync(
	request *ExchangeRequest,
	callback chan<- ExchangeAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "exchange",
    		"function": "exchange",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
    if request.Count != nil {
        bodies["count"] = *request.Count
    }
    if request.Config != nil {
        var _config []*map[string]interface {}
        for _, item := range request.Config {
            _config = append(_config, item.ToDict())
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.exchangeAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) Exchange(
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

func (p Gs2ExchangeWebSocketClient) exchangeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ExchangeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) ExchangeByUserIdAsync(
	request *ExchangeByUserIdRequest,
	callback chan<- ExchangeByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "exchange",
    		"function": "exchangeByUserId",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.Count != nil {
        bodies["count"] = *request.Count
    }
    if request.Config != nil {
        var _config []*map[string]interface {}
        for _, item := range request.Config {
            _config = append(_config, item.ToDict())
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.exchangeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) ExchangeByUserId(
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

func (p Gs2ExchangeWebSocketClient) exchangeByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ExchangeByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) ExchangeByStampSheetAsync(
	request *ExchangeByStampSheetRequest,
	callback chan<- ExchangeByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "exchange",
    		"function": "exchangeByStampSheet",
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

	go p.exchangeByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) ExchangeByStampSheet(
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

func (p Gs2ExchangeWebSocketClient) exportMasterAsyncHandler(
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

func (p Gs2ExchangeWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "currentRateMaster",
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

func (p Gs2ExchangeWebSocketClient) ExportMaster(
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

func (p Gs2ExchangeWebSocketClient) getCurrentRateMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentRateMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) GetCurrentRateMasterAsync(
	request *GetCurrentRateMasterRequest,
	callback chan<- GetCurrentRateMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "currentRateMaster",
    		"function": "getCurrentRateMaster",
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

	go p.getCurrentRateMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) GetCurrentRateMaster(
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

func (p Gs2ExchangeWebSocketClient) updateCurrentRateMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentRateMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) UpdateCurrentRateMasterAsync(
	request *UpdateCurrentRateMasterRequest,
	callback chan<- UpdateCurrentRateMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "currentRateMaster",
    		"function": "updateCurrentRateMaster",
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

	go p.updateCurrentRateMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) UpdateCurrentRateMaster(
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

func (p Gs2ExchangeWebSocketClient) updateCurrentRateMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentRateMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) UpdateCurrentRateMasterFromGitHubAsync(
	request *UpdateCurrentRateMasterFromGitHubRequest,
	callback chan<- UpdateCurrentRateMasterFromGitHubAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "currentRateMaster",
    		"function": "updateCurrentRateMasterFromGitHub",
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

	go p.updateCurrentRateMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) UpdateCurrentRateMasterFromGitHub(
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

func (p Gs2ExchangeWebSocketClient) createAwaitByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateAwaitByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) CreateAwaitByUserIdAsync(
	request *CreateAwaitByUserIdRequest,
	callback chan<- CreateAwaitByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "createAwaitByUserId",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
    if request.Count != nil {
        bodies["count"] = *request.Count
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createAwaitByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) CreateAwaitByUserId(
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

func (p Gs2ExchangeWebSocketClient) describeAwaitsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeAwaitsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) DescribeAwaitsAsync(
	request *DescribeAwaitsRequest,
	callback chan<- DescribeAwaitsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "describeAwaits",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
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

	go p.describeAwaitsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) DescribeAwaits(
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

func (p Gs2ExchangeWebSocketClient) describeAwaitsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeAwaitsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) DescribeAwaitsByUserIdAsync(
	request *DescribeAwaitsByUserIdRequest,
	callback chan<- DescribeAwaitsByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "describeAwaitsByUserId",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
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

	go p.describeAwaitsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) DescribeAwaitsByUserId(
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

func (p Gs2ExchangeWebSocketClient) getAwaitAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetAwaitAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) GetAwaitAsync(
	request *GetAwaitRequest,
	callback chan<- GetAwaitAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "getAwait",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
    if request.AwaitName != nil && *request.AwaitName != "" {
        bodies["awaitName"] = *request.AwaitName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getAwaitAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) GetAwait(
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

func (p Gs2ExchangeWebSocketClient) getAwaitByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetAwaitByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) GetAwaitByUserIdAsync(
	request *GetAwaitByUserIdRequest,
	callback chan<- GetAwaitByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "getAwaitByUserId",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
    if request.AwaitName != nil && *request.AwaitName != "" {
        bodies["awaitName"] = *request.AwaitName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getAwaitByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) GetAwaitByUserId(
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

func (p Gs2ExchangeWebSocketClient) acquireAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AcquireAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) AcquireAsync(
	request *AcquireRequest,
	callback chan<- AcquireAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "acquire",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
    if request.AwaitName != nil && *request.AwaitName != "" {
        bodies["awaitName"] = *request.AwaitName
    }
    if request.Config != nil {
        var _config []*map[string]interface {}
        for _, item := range request.Config {
            _config = append(_config, item.ToDict())
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.acquireAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) Acquire(
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

func (p Gs2ExchangeWebSocketClient) acquireByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AcquireByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) AcquireByUserIdAsync(
	request *AcquireByUserIdRequest,
	callback chan<- AcquireByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "acquireByUserId",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
    if request.AwaitName != nil && *request.AwaitName != "" {
        bodies["awaitName"] = *request.AwaitName
    }
    if request.Config != nil {
        var _config []*map[string]interface {}
        for _, item := range request.Config {
            _config = append(_config, item.ToDict())
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.acquireByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) AcquireByUserId(
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

func (p Gs2ExchangeWebSocketClient) acquireForceByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AcquireForceByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) AcquireForceByUserIdAsync(
	request *AcquireForceByUserIdRequest,
	callback chan<- AcquireForceByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "acquireForceByUserId",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
    if request.AwaitName != nil && *request.AwaitName != "" {
        bodies["awaitName"] = *request.AwaitName
    }
    if request.Config != nil {
        var _config []*map[string]interface {}
        for _, item := range request.Config {
            _config = append(_config, item.ToDict())
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.acquireForceByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) AcquireForceByUserId(
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

func (p Gs2ExchangeWebSocketClient) skipAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SkipAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) SkipAsync(
	request *SkipRequest,
	callback chan<- SkipAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "skip",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
    if request.AwaitName != nil && *request.AwaitName != "" {
        bodies["awaitName"] = *request.AwaitName
    }
    if request.Config != nil {
        var _config []*map[string]interface {}
        for _, item := range request.Config {
            _config = append(_config, item.ToDict())
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.skipAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) Skip(
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

func (p Gs2ExchangeWebSocketClient) skipByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SkipByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) SkipByUserIdAsync(
	request *SkipByUserIdRequest,
	callback chan<- SkipByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "skipByUserId",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
    if request.AwaitName != nil && *request.AwaitName != "" {
        bodies["awaitName"] = *request.AwaitName
    }
    if request.Config != nil {
        var _config []*map[string]interface {}
        for _, item := range request.Config {
            _config = append(_config, item.ToDict())
        }
        bodies["config"] = _config
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.skipByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) SkipByUserId(
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

func (p Gs2ExchangeWebSocketClient) deleteAwaitAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteAwaitAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) DeleteAwaitAsync(
	request *DeleteAwaitRequest,
	callback chan<- DeleteAwaitAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "deleteAwait",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
    if request.AwaitName != nil && *request.AwaitName != "" {
        bodies["awaitName"] = *request.AwaitName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.deleteAwaitAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) DeleteAwait(
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

func (p Gs2ExchangeWebSocketClient) deleteAwaitByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteAwaitByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) DeleteAwaitByUserIdAsync(
	request *DeleteAwaitByUserIdRequest,
	callback chan<- DeleteAwaitByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "deleteAwaitByUserId",
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
    if request.RateName != nil && *request.RateName != "" {
        bodies["rateName"] = *request.RateName
    }
    if request.AwaitName != nil && *request.AwaitName != "" {
        bodies["awaitName"] = *request.AwaitName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteAwaitByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) DeleteAwaitByUserId(
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

func (p Gs2ExchangeWebSocketClient) createAwaitByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateAwaitByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) CreateAwaitByStampSheetAsync(
	request *CreateAwaitByStampSheetRequest,
	callback chan<- CreateAwaitByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "createAwaitByStampSheet",
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

	go p.createAwaitByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) CreateAwaitByStampSheet(
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

func (p Gs2ExchangeWebSocketClient) deleteAwaitByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteAwaitByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ExchangeWebSocketClient) DeleteAwaitByStampTaskAsync(
	request *DeleteAwaitByStampTaskRequest,
	callback chan<- DeleteAwaitByStampTaskAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "exchange",
    		"component": "await",
    		"function": "deleteAwaitByStampTask",
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

	go p.deleteAwaitByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExchangeWebSocketClient) DeleteAwaitByStampTask(
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
