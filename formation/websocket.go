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

package formation

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2FormationWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2FormationWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2FormationWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2FormationWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
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

func (p Gs2FormationWebSocketClient) DescribeNamespaces(
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

func (p Gs2FormationWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2FormationWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
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
    if request.UpdateMoldScript != nil {
        bodies["updateMoldScript"] = request.UpdateMoldScript.ToDict()
    }
    if request.UpdateFormScript != nil {
        bodies["updateFormScript"] = request.UpdateFormScript.ToDict()
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

func (p Gs2FormationWebSocketClient) CreateNamespace(
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

func (p Gs2FormationWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2FormationWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
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

func (p Gs2FormationWebSocketClient) GetNamespaceStatus(
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

func (p Gs2FormationWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2FormationWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
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

func (p Gs2FormationWebSocketClient) GetNamespace(
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

func (p Gs2FormationWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2FormationWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
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
    if request.UpdateMoldScript != nil {
        bodies["updateMoldScript"] = request.UpdateMoldScript.ToDict()
    }
    if request.UpdateFormScript != nil {
        bodies["updateFormScript"] = request.UpdateFormScript.ToDict()
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

func (p Gs2FormationWebSocketClient) UpdateNamespace(
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

func (p Gs2FormationWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2FormationWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
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

func (p Gs2FormationWebSocketClient) DeleteNamespace(
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

func (p Gs2FormationWebSocketClient) describeFormModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeFormModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFormModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFormModelMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeFormModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeFormModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DescribeFormModelMastersAsync(
	request *DescribeFormModelMastersRequest,
	callback chan<- DescribeFormModelMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "formModelMaster",
    		"function": "describeFormModelMasters",
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

	go p.describeFormModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DescribeFormModelMasters(
	request *DescribeFormModelMastersRequest,
) (*DescribeFormModelMastersResult, error) {
	callback := make(chan DescribeFormModelMastersAsyncResult, 1)
	go p.DescribeFormModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) createFormModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateFormModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateFormModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateFormModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateFormModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateFormModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) CreateFormModelMasterAsync(
	request *CreateFormModelMasterRequest,
	callback chan<- CreateFormModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "formModelMaster",
    		"function": "createFormModelMaster",
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
    if request.Slots != nil {
        var _slots []interface {}
        for _, item := range request.Slots {
            _slots = append(_slots, item)
        }
        bodies["slots"] = _slots
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createFormModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) CreateFormModelMaster(
	request *CreateFormModelMasterRequest,
) (*CreateFormModelMasterResult, error) {
	callback := make(chan CreateFormModelMasterAsyncResult, 1)
	go p.CreateFormModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) getFormModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetFormModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFormModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFormModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetFormModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetFormModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) GetFormModelMasterAsync(
	request *GetFormModelMasterRequest,
	callback chan<- GetFormModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "formModelMaster",
    		"function": "getFormModelMaster",
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
    if request.FormModelName != nil && *request.FormModelName != "" {
        bodies["formModelName"] = *request.FormModelName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getFormModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) GetFormModelMaster(
	request *GetFormModelMasterRequest,
) (*GetFormModelMasterResult, error) {
	callback := make(chan GetFormModelMasterAsyncResult, 1)
	go p.GetFormModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) updateFormModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateFormModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateFormModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateFormModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateFormModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateFormModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) UpdateFormModelMasterAsync(
	request *UpdateFormModelMasterRequest,
	callback chan<- UpdateFormModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "formModelMaster",
    		"function": "updateFormModelMaster",
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
    if request.FormModelName != nil && *request.FormModelName != "" {
        bodies["formModelName"] = *request.FormModelName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.Slots != nil {
        var _slots []interface {}
        for _, item := range request.Slots {
            _slots = append(_slots, item)
        }
        bodies["slots"] = _slots
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateFormModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) UpdateFormModelMaster(
	request *UpdateFormModelMasterRequest,
) (*UpdateFormModelMasterResult, error) {
	callback := make(chan UpdateFormModelMasterAsyncResult, 1)
	go p.UpdateFormModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) deleteFormModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteFormModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteFormModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteFormModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteFormModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteFormModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DeleteFormModelMasterAsync(
	request *DeleteFormModelMasterRequest,
	callback chan<- DeleteFormModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "formModelMaster",
    		"function": "deleteFormModelMaster",
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
    if request.FormModelName != nil && *request.FormModelName != "" {
        bodies["formModelName"] = *request.FormModelName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteFormModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DeleteFormModelMaster(
	request *DeleteFormModelMasterRequest,
) (*DeleteFormModelMasterResult, error) {
	callback := make(chan DeleteFormModelMasterAsyncResult, 1)
	go p.DeleteFormModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) describeMoldModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMoldModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMoldModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMoldModelsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeMoldModelsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeMoldModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DescribeMoldModelsAsync(
	request *DescribeMoldModelsRequest,
	callback chan<- DescribeMoldModelsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "moldModel",
    		"function": "describeMoldModels",
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

	go p.describeMoldModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DescribeMoldModels(
	request *DescribeMoldModelsRequest,
) (*DescribeMoldModelsResult, error) {
	callback := make(chan DescribeMoldModelsAsyncResult, 1)
	go p.DescribeMoldModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) getMoldModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMoldModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMoldModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMoldModelResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetMoldModelAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetMoldModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) GetMoldModelAsync(
	request *GetMoldModelRequest,
	callback chan<- GetMoldModelAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "moldModel",
    		"function": "getMoldModel",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getMoldModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) GetMoldModel(
	request *GetMoldModelRequest,
) (*GetMoldModelResult, error) {
	callback := make(chan GetMoldModelAsyncResult, 1)
	go p.GetMoldModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) describeMoldModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMoldModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMoldModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMoldModelMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeMoldModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeMoldModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DescribeMoldModelMastersAsync(
	request *DescribeMoldModelMastersRequest,
	callback chan<- DescribeMoldModelMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "moldModelMaster",
    		"function": "describeMoldModelMasters",
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

	go p.describeMoldModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DescribeMoldModelMasters(
	request *DescribeMoldModelMastersRequest,
) (*DescribeMoldModelMastersResult, error) {
	callback := make(chan DescribeMoldModelMastersAsyncResult, 1)
	go p.DescribeMoldModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) createMoldModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateMoldModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateMoldModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateMoldModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateMoldModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateMoldModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) CreateMoldModelMasterAsync(
	request *CreateMoldModelMasterRequest,
	callback chan<- CreateMoldModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "moldModelMaster",
    		"function": "createMoldModelMaster",
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
    if request.FormModelName != nil && *request.FormModelName != "" {
        bodies["formModelName"] = *request.FormModelName
    }
    if request.InitialMaxCapacity != nil {
        bodies["initialMaxCapacity"] = *request.InitialMaxCapacity
    }
    if request.MaxCapacity != nil {
        bodies["maxCapacity"] = *request.MaxCapacity
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createMoldModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) CreateMoldModelMaster(
	request *CreateMoldModelMasterRequest,
) (*CreateMoldModelMasterResult, error) {
	callback := make(chan CreateMoldModelMasterAsyncResult, 1)
	go p.CreateMoldModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) getMoldModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMoldModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMoldModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMoldModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetMoldModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetMoldModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) GetMoldModelMasterAsync(
	request *GetMoldModelMasterRequest,
	callback chan<- GetMoldModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "moldModelMaster",
    		"function": "getMoldModelMaster",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getMoldModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) GetMoldModelMaster(
	request *GetMoldModelMasterRequest,
) (*GetMoldModelMasterResult, error) {
	callback := make(chan GetMoldModelMasterAsyncResult, 1)
	go p.GetMoldModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) updateMoldModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateMoldModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMoldModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMoldModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateMoldModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateMoldModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) UpdateMoldModelMasterAsync(
	request *UpdateMoldModelMasterRequest,
	callback chan<- UpdateMoldModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "moldModelMaster",
    		"function": "updateMoldModelMaster",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.FormModelName != nil && *request.FormModelName != "" {
        bodies["formModelName"] = *request.FormModelName
    }
    if request.InitialMaxCapacity != nil {
        bodies["initialMaxCapacity"] = *request.InitialMaxCapacity
    }
    if request.MaxCapacity != nil {
        bodies["maxCapacity"] = *request.MaxCapacity
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateMoldModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) UpdateMoldModelMaster(
	request *UpdateMoldModelMasterRequest,
) (*UpdateMoldModelMasterResult, error) {
	callback := make(chan UpdateMoldModelMasterAsyncResult, 1)
	go p.UpdateMoldModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) deleteMoldModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteMoldModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMoldModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMoldModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteMoldModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteMoldModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DeleteMoldModelMasterAsync(
	request *DeleteMoldModelMasterRequest,
	callback chan<- DeleteMoldModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "moldModelMaster",
    		"function": "deleteMoldModelMaster",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteMoldModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DeleteMoldModelMaster(
	request *DeleteMoldModelMasterRequest,
) (*DeleteMoldModelMasterResult, error) {
	callback := make(chan DeleteMoldModelMasterAsyncResult, 1)
	go p.DeleteMoldModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) exportMasterAsyncHandler(
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

func (p Gs2FormationWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "currentFormMaster",
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

func (p Gs2FormationWebSocketClient) ExportMaster(
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

func (p Gs2FormationWebSocketClient) getCurrentFormMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentFormMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentFormMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentFormMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentFormMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentFormMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) GetCurrentFormMasterAsync(
	request *GetCurrentFormMasterRequest,
	callback chan<- GetCurrentFormMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "currentFormMaster",
    		"function": "getCurrentFormMaster",
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

	go p.getCurrentFormMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) GetCurrentFormMaster(
	request *GetCurrentFormMasterRequest,
) (*GetCurrentFormMasterResult, error) {
	callback := make(chan GetCurrentFormMasterAsyncResult, 1)
	go p.GetCurrentFormMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) updateCurrentFormMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentFormMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentFormMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentFormMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentFormMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentFormMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) UpdateCurrentFormMasterAsync(
	request *UpdateCurrentFormMasterRequest,
	callback chan<- UpdateCurrentFormMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "currentFormMaster",
    		"function": "updateCurrentFormMaster",
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

	go p.updateCurrentFormMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) UpdateCurrentFormMaster(
	request *UpdateCurrentFormMasterRequest,
) (*UpdateCurrentFormMasterResult, error) {
	callback := make(chan UpdateCurrentFormMasterAsyncResult, 1)
	go p.UpdateCurrentFormMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) updateCurrentFormMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentFormMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentFormMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentFormMasterFromGitHubResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentFormMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentFormMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) UpdateCurrentFormMasterFromGitHubAsync(
	request *UpdateCurrentFormMasterFromGitHubRequest,
	callback chan<- UpdateCurrentFormMasterFromGitHubAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "currentFormMaster",
    		"function": "updateCurrentFormMasterFromGitHub",
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

	go p.updateCurrentFormMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) UpdateCurrentFormMasterFromGitHub(
	request *UpdateCurrentFormMasterFromGitHubRequest,
) (*UpdateCurrentFormMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentFormMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentFormMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) describeMoldsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMoldsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMoldsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMoldsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeMoldsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeMoldsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DescribeMoldsAsync(
	request *DescribeMoldsRequest,
	callback chan<- DescribeMoldsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "mold",
    		"function": "describeMolds",
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

	go p.describeMoldsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DescribeMolds(
	request *DescribeMoldsRequest,
) (*DescribeMoldsResult, error) {
	callback := make(chan DescribeMoldsAsyncResult, 1)
	go p.DescribeMoldsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) describeMoldsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMoldsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMoldsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMoldsByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeMoldsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeMoldsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DescribeMoldsByUserIdAsync(
	request *DescribeMoldsByUserIdRequest,
	callback chan<- DescribeMoldsByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "mold",
    		"function": "describeMoldsByUserId",
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

	go p.describeMoldsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DescribeMoldsByUserId(
	request *DescribeMoldsByUserIdRequest,
) (*DescribeMoldsByUserIdResult, error) {
	callback := make(chan DescribeMoldsByUserIdAsyncResult, 1)
	go p.DescribeMoldsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) getMoldAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMoldAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMoldAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMoldResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetMoldAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetMoldAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) GetMoldAsync(
	request *GetMoldRequest,
	callback chan<- GetMoldAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "mold",
    		"function": "getMold",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getMoldAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) GetMold(
	request *GetMoldRequest,
) (*GetMoldResult, error) {
	callback := make(chan GetMoldAsyncResult, 1)
	go p.GetMoldAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) getMoldByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMoldByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMoldByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMoldByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetMoldByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetMoldByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) GetMoldByUserIdAsync(
	request *GetMoldByUserIdRequest,
	callback chan<- GetMoldByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "mold",
    		"function": "getMoldByUserId",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getMoldByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) GetMoldByUserId(
	request *GetMoldByUserIdRequest,
) (*GetMoldByUserIdResult, error) {
	callback := make(chan GetMoldByUserIdAsyncResult, 1)
	go p.GetMoldByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) setMoldCapacityByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetMoldCapacityByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetMoldCapacityByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetMoldCapacityByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetMoldCapacityByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetMoldCapacityByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) SetMoldCapacityByUserIdAsync(
	request *SetMoldCapacityByUserIdRequest,
	callback chan<- SetMoldCapacityByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "mold",
    		"function": "setMoldCapacityByUserId",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
    if request.Capacity != nil {
        bodies["capacity"] = *request.Capacity
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.setMoldCapacityByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) SetMoldCapacityByUserId(
	request *SetMoldCapacityByUserIdRequest,
) (*SetMoldCapacityByUserIdResult, error) {
	callback := make(chan SetMoldCapacityByUserIdAsyncResult, 1)
	go p.SetMoldCapacityByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) addMoldCapacityByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddMoldCapacityByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddMoldCapacityByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddMoldCapacityByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddMoldCapacityByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AddMoldCapacityByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) AddMoldCapacityByUserIdAsync(
	request *AddMoldCapacityByUserIdRequest,
	callback chan<- AddMoldCapacityByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "mold",
    		"function": "addMoldCapacityByUserId",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
    if request.Capacity != nil {
        bodies["capacity"] = *request.Capacity
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.addMoldCapacityByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) AddMoldCapacityByUserId(
	request *AddMoldCapacityByUserIdRequest,
) (*AddMoldCapacityByUserIdResult, error) {
	callback := make(chan AddMoldCapacityByUserIdAsyncResult, 1)
	go p.AddMoldCapacityByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) deleteMoldAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteMoldAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMoldAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMoldResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteMoldAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteMoldAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DeleteMoldAsync(
	request *DeleteMoldRequest,
	callback chan<- DeleteMoldAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "mold",
    		"function": "deleteMold",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.deleteMoldAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DeleteMold(
	request *DeleteMoldRequest,
) (*DeleteMoldResult, error) {
	callback := make(chan DeleteMoldAsyncResult, 1)
	go p.DeleteMoldAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) deleteMoldByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteMoldByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMoldByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMoldByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteMoldByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteMoldByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DeleteMoldByUserIdAsync(
	request *DeleteMoldByUserIdRequest,
	callback chan<- DeleteMoldByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "mold",
    		"function": "deleteMoldByUserId",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.deleteMoldByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DeleteMoldByUserId(
	request *DeleteMoldByUserIdRequest,
) (*DeleteMoldByUserIdResult, error) {
	callback := make(chan DeleteMoldByUserIdAsyncResult, 1)
	go p.DeleteMoldByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) addCapacityByStampSheetAsyncHandler(
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
	callback <- AddCapacityByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) AddCapacityByStampSheetAsync(
	request *AddCapacityByStampSheetRequest,
	callback chan<- AddCapacityByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "mold",
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

func (p Gs2FormationWebSocketClient) AddCapacityByStampSheet(
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

func (p Gs2FormationWebSocketClient) setCapacityByStampSheetAsyncHandler(
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
	callback <- SetCapacityByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) SetCapacityByStampSheetAsync(
	request *SetCapacityByStampSheetRequest,
	callback chan<- SetCapacityByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "mold",
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

func (p Gs2FormationWebSocketClient) SetCapacityByStampSheet(
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

func (p Gs2FormationWebSocketClient) describeFormsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeFormsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFormsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFormsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeFormsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeFormsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DescribeFormsAsync(
	request *DescribeFormsRequest,
	callback chan<- DescribeFormsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "form",
    		"function": "describeForms",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
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

	go p.describeFormsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DescribeForms(
	request *DescribeFormsRequest,
) (*DescribeFormsResult, error) {
	callback := make(chan DescribeFormsAsyncResult, 1)
	go p.DescribeFormsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) describeFormsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeFormsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFormsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFormsByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeFormsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeFormsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DescribeFormsByUserIdAsync(
	request *DescribeFormsByUserIdRequest,
	callback chan<- DescribeFormsByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "form",
    		"function": "describeFormsByUserId",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
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

	go p.describeFormsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DescribeFormsByUserId(
	request *DescribeFormsByUserIdRequest,
) (*DescribeFormsByUserIdResult, error) {
	callback := make(chan DescribeFormsByUserIdAsyncResult, 1)
	go p.DescribeFormsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) getFormAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetFormAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFormAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFormResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetFormAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetFormAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) GetFormAsync(
	request *GetFormRequest,
	callback chan<- GetFormAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "form",
    		"function": "getForm",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
    if request.Index != nil {
        bodies["index"] = *request.Index
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getFormAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) GetForm(
	request *GetFormRequest,
) (*GetFormResult, error) {
	callback := make(chan GetFormAsyncResult, 1)
	go p.GetFormAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) getFormByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetFormByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFormByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFormByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetFormByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetFormByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) GetFormByUserIdAsync(
	request *GetFormByUserIdRequest,
	callback chan<- GetFormByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "form",
    		"function": "getFormByUserId",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
    if request.Index != nil {
        bodies["index"] = *request.Index
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getFormByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) GetFormByUserId(
	request *GetFormByUserIdRequest,
) (*GetFormByUserIdResult, error) {
	callback := make(chan GetFormByUserIdAsyncResult, 1)
	go p.GetFormByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) getFormWithSignatureAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetFormWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFormWithSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFormWithSignatureResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetFormWithSignatureAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetFormWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) GetFormWithSignatureAsync(
	request *GetFormWithSignatureRequest,
	callback chan<- GetFormWithSignatureAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "form",
    		"function": "getFormWithSignature",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
    if request.Index != nil {
        bodies["index"] = *request.Index
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

	go p.getFormWithSignatureAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) GetFormWithSignature(
	request *GetFormWithSignatureRequest,
) (*GetFormWithSignatureResult, error) {
	callback := make(chan GetFormWithSignatureAsyncResult, 1)
	go p.GetFormWithSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) getFormWithSignatureByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetFormWithSignatureByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFormWithSignatureByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFormWithSignatureByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetFormWithSignatureByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetFormWithSignatureByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) GetFormWithSignatureByUserIdAsync(
	request *GetFormWithSignatureByUserIdRequest,
	callback chan<- GetFormWithSignatureByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "form",
    		"function": "getFormWithSignatureByUserId",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
    if request.Index != nil {
        bodies["index"] = *request.Index
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getFormWithSignatureByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) GetFormWithSignatureByUserId(
	request *GetFormWithSignatureByUserIdRequest,
) (*GetFormWithSignatureByUserIdResult, error) {
	callback := make(chan GetFormWithSignatureByUserIdAsyncResult, 1)
	go p.GetFormWithSignatureByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) setFormByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetFormByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetFormByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetFormByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetFormByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetFormByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) SetFormByUserIdAsync(
	request *SetFormByUserIdRequest,
	callback chan<- SetFormByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "form",
    		"function": "setFormByUserId",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
    if request.Index != nil {
        bodies["index"] = *request.Index
    }
    if request.Slots != nil {
        var _slots []interface {}
        for _, item := range request.Slots {
            _slots = append(_slots, item)
        }
        bodies["slots"] = _slots
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.setFormByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) SetFormByUserId(
	request *SetFormByUserIdRequest,
) (*SetFormByUserIdResult, error) {
	callback := make(chan SetFormByUserIdAsyncResult, 1)
	go p.SetFormByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) setFormWithSignatureAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetFormWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetFormWithSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetFormWithSignatureResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetFormWithSignatureAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetFormWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) SetFormWithSignatureAsync(
	request *SetFormWithSignatureRequest,
	callback chan<- SetFormWithSignatureAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "form",
    		"function": "setFormWithSignature",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
    if request.Index != nil {
        bodies["index"] = *request.Index
    }
    if request.Slots != nil {
        var _slots []interface {}
        for _, item := range request.Slots {
            _slots = append(_slots, item)
        }
        bodies["slots"] = _slots
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

	go p.setFormWithSignatureAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) SetFormWithSignature(
	request *SetFormWithSignatureRequest,
) (*SetFormWithSignatureResult, error) {
	callback := make(chan SetFormWithSignatureAsyncResult, 1)
	go p.SetFormWithSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) acquireActionsToFormPropertiesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AcquireActionsToFormPropertiesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireActionsToFormPropertiesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireActionsToFormPropertiesResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AcquireActionsToFormPropertiesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AcquireActionsToFormPropertiesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) AcquireActionsToFormPropertiesAsync(
	request *AcquireActionsToFormPropertiesRequest,
	callback chan<- AcquireActionsToFormPropertiesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "form",
    		"function": "acquireActionsToFormProperties",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
    if request.Index != nil {
        bodies["index"] = *request.Index
    }
    if request.AcquireAction != nil {
        bodies["acquireAction"] = request.AcquireAction.ToDict()
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
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.acquireActionsToFormPropertiesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) AcquireActionsToFormProperties(
	request *AcquireActionsToFormPropertiesRequest,
) (*AcquireActionsToFormPropertiesResult, error) {
	callback := make(chan AcquireActionsToFormPropertiesAsyncResult, 1)
	go p.AcquireActionsToFormPropertiesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) deleteFormAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteFormAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteFormAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteFormResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteFormAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteFormAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DeleteFormAsync(
	request *DeleteFormRequest,
	callback chan<- DeleteFormAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "form",
    		"function": "deleteForm",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
    if request.Index != nil {
        bodies["index"] = *request.Index
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.deleteFormAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DeleteForm(
	request *DeleteFormRequest,
) (*DeleteFormResult, error) {
	callback := make(chan DeleteFormAsyncResult, 1)
	go p.DeleteFormAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) deleteFormByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteFormByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteFormByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteFormByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteFormByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteFormByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) DeleteFormByUserIdAsync(
	request *DeleteFormByUserIdRequest,
	callback chan<- DeleteFormByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "form",
    		"function": "deleteFormByUserId",
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
    if request.MoldName != nil && *request.MoldName != "" {
        bodies["moldName"] = *request.MoldName
    }
    if request.Index != nil {
        bodies["index"] = *request.Index
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.deleteFormByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) DeleteFormByUserId(
	request *DeleteFormByUserIdRequest,
) (*DeleteFormByUserIdResult, error) {
	callback := make(chan DeleteFormByUserIdAsyncResult, 1)
	go p.DeleteFormByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FormationWebSocketClient) acquireActionToFormPropertiesByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AcquireActionToFormPropertiesByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireActionToFormPropertiesByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireActionToFormPropertiesByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AcquireActionToFormPropertiesByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AcquireActionToFormPropertiesByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationWebSocketClient) AcquireActionToFormPropertiesByStampSheetAsync(
	request *AcquireActionToFormPropertiesByStampSheetRequest,
	callback chan<- AcquireActionToFormPropertiesByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "formation",
    		"component": "form",
    		"function": "acquireActionToFormPropertiesByStampSheet",
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

	go p.acquireActionToFormPropertiesByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2FormationWebSocketClient) AcquireActionToFormPropertiesByStampSheet(
	request *AcquireActionToFormPropertiesByStampSheetRequest,
) (*AcquireActionToFormPropertiesByStampSheetResult, error) {
	callback := make(chan AcquireActionToFormPropertiesByStampSheetAsyncResult, 1)
	go p.AcquireActionToFormPropertiesByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
