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

package quest

import (
	"encoding/json"
	"github.com/google/uuid"
	"core"
)

type Gs2QuestWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2QuestWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2QuestWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2QuestWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
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

func (p Gs2QuestWebSocketClient) DescribeNamespaces(
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

func (p Gs2QuestWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2QuestWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
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
    if request.StartQuestScript != nil {
        bodies["startQuestScript"] = request.StartQuestScript.ToDict()
    }
    if request.CompleteQuestScript != nil {
        bodies["completeQuestScript"] = request.CompleteQuestScript.ToDict()
    }
    if request.FailedQuestScript != nil {
        bodies["failedQuestScript"] = request.FailedQuestScript.ToDict()
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

func (p Gs2QuestWebSocketClient) CreateNamespace(
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

func (p Gs2QuestWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2QuestWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
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

func (p Gs2QuestWebSocketClient) GetNamespaceStatus(
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

func (p Gs2QuestWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2QuestWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
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

func (p Gs2QuestWebSocketClient) GetNamespace(
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

func (p Gs2QuestWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2QuestWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
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
    if request.StartQuestScript != nil {
        bodies["startQuestScript"] = request.StartQuestScript.ToDict()
    }
    if request.CompleteQuestScript != nil {
        bodies["completeQuestScript"] = request.CompleteQuestScript.ToDict()
    }
    if request.FailedQuestScript != nil {
        bodies["failedQuestScript"] = request.FailedQuestScript.ToDict()
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

func (p Gs2QuestWebSocketClient) UpdateNamespace(
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

func (p Gs2QuestWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2QuestWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
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

func (p Gs2QuestWebSocketClient) DeleteNamespace(
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

func (p Gs2QuestWebSocketClient) describeQuestGroupModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeQuestGroupModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeQuestGroupModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeQuestGroupModelMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeQuestGroupModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeQuestGroupModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DescribeQuestGroupModelMastersAsync(
	request *DescribeQuestGroupModelMastersRequest,
	callback chan<- DescribeQuestGroupModelMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questGroupModelMaster",
    		"function": "describeQuestGroupModelMasters",
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

	go p.describeQuestGroupModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DescribeQuestGroupModelMasters(
	request *DescribeQuestGroupModelMastersRequest,
) (*DescribeQuestGroupModelMastersResult, error) {
	callback := make(chan DescribeQuestGroupModelMastersAsyncResult, 1)
	go p.DescribeQuestGroupModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) createQuestGroupModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateQuestGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateQuestGroupModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateQuestGroupModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateQuestGroupModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateQuestGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) CreateQuestGroupModelMasterAsync(
	request *CreateQuestGroupModelMasterRequest,
	callback chan<- CreateQuestGroupModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questGroupModelMaster",
    		"function": "createQuestGroupModelMaster",
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
    if request.ChallengePeriodEventId != nil && *request.ChallengePeriodEventId != "" {
        bodies["challengePeriodEventId"] = *request.ChallengePeriodEventId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createQuestGroupModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) CreateQuestGroupModelMaster(
	request *CreateQuestGroupModelMasterRequest,
) (*CreateQuestGroupModelMasterResult, error) {
	callback := make(chan CreateQuestGroupModelMasterAsyncResult, 1)
	go p.CreateQuestGroupModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) getQuestGroupModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetQuestGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetQuestGroupModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetQuestGroupModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetQuestGroupModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetQuestGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) GetQuestGroupModelMasterAsync(
	request *GetQuestGroupModelMasterRequest,
	callback chan<- GetQuestGroupModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questGroupModelMaster",
    		"function": "getQuestGroupModelMaster",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getQuestGroupModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) GetQuestGroupModelMaster(
	request *GetQuestGroupModelMasterRequest,
) (*GetQuestGroupModelMasterResult, error) {
	callback := make(chan GetQuestGroupModelMasterAsyncResult, 1)
	go p.GetQuestGroupModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) updateQuestGroupModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateQuestGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateQuestGroupModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateQuestGroupModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateQuestGroupModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateQuestGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) UpdateQuestGroupModelMasterAsync(
	request *UpdateQuestGroupModelMasterRequest,
	callback chan<- UpdateQuestGroupModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questGroupModelMaster",
    		"function": "updateQuestGroupModelMaster",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.ChallengePeriodEventId != nil && *request.ChallengePeriodEventId != "" {
        bodies["challengePeriodEventId"] = *request.ChallengePeriodEventId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateQuestGroupModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) UpdateQuestGroupModelMaster(
	request *UpdateQuestGroupModelMasterRequest,
) (*UpdateQuestGroupModelMasterResult, error) {
	callback := make(chan UpdateQuestGroupModelMasterAsyncResult, 1)
	go p.UpdateQuestGroupModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) deleteQuestGroupModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteQuestGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteQuestGroupModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteQuestGroupModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteQuestGroupModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteQuestGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DeleteQuestGroupModelMasterAsync(
	request *DeleteQuestGroupModelMasterRequest,
	callback chan<- DeleteQuestGroupModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questGroupModelMaster",
    		"function": "deleteQuestGroupModelMaster",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteQuestGroupModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DeleteQuestGroupModelMaster(
	request *DeleteQuestGroupModelMasterRequest,
) (*DeleteQuestGroupModelMasterResult, error) {
	callback := make(chan DeleteQuestGroupModelMasterAsyncResult, 1)
	go p.DeleteQuestGroupModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) describeQuestModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeQuestModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeQuestModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeQuestModelMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeQuestModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeQuestModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DescribeQuestModelMastersAsync(
	request *DescribeQuestModelMastersRequest,
	callback chan<- DescribeQuestModelMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questModelMaster",
    		"function": "describeQuestModelMasters",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
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

	go p.describeQuestModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DescribeQuestModelMasters(
	request *DescribeQuestModelMastersRequest,
) (*DescribeQuestModelMastersResult, error) {
	callback := make(chan DescribeQuestModelMastersAsyncResult, 1)
	go p.DescribeQuestModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) createQuestModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateQuestModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateQuestModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateQuestModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateQuestModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateQuestModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) CreateQuestModelMasterAsync(
	request *CreateQuestModelMasterRequest,
	callback chan<- CreateQuestModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questModelMaster",
    		"function": "createQuestModelMaster",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
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
    if request.Contents != nil {
        var _contents []interface {}
        for _, item := range request.Contents {
            _contents = append(_contents, item)
        }
        bodies["contents"] = _contents
    }
    if request.ChallengePeriodEventId != nil && *request.ChallengePeriodEventId != "" {
        bodies["challengePeriodEventId"] = *request.ChallengePeriodEventId
    }
    if request.ConsumeActions != nil {
        var _consumeActions []interface {}
        for _, item := range request.ConsumeActions {
            _consumeActions = append(_consumeActions, item)
        }
        bodies["consumeActions"] = _consumeActions
    }
    if request.FailedAcquireActions != nil {
        var _failedAcquireActions []interface {}
        for _, item := range request.FailedAcquireActions {
            _failedAcquireActions = append(_failedAcquireActions, item)
        }
        bodies["failedAcquireActions"] = _failedAcquireActions
    }
    if request.PremiseQuestNames != nil {
        var _premiseQuestNames []interface {}
        for _, item := range request.PremiseQuestNames {
            _premiseQuestNames = append(_premiseQuestNames, item)
        }
        bodies["premiseQuestNames"] = _premiseQuestNames
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createQuestModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) CreateQuestModelMaster(
	request *CreateQuestModelMasterRequest,
) (*CreateQuestModelMasterResult, error) {
	callback := make(chan CreateQuestModelMasterAsyncResult, 1)
	go p.CreateQuestModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) getQuestModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetQuestModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetQuestModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetQuestModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetQuestModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetQuestModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) GetQuestModelMasterAsync(
	request *GetQuestModelMasterRequest,
	callback chan<- GetQuestModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questModelMaster",
    		"function": "getQuestModelMaster",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
    if request.QuestName != nil && *request.QuestName != "" {
        bodies["questName"] = *request.QuestName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getQuestModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) GetQuestModelMaster(
	request *GetQuestModelMasterRequest,
) (*GetQuestModelMasterResult, error) {
	callback := make(chan GetQuestModelMasterAsyncResult, 1)
	go p.GetQuestModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) updateQuestModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateQuestModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateQuestModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateQuestModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateQuestModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateQuestModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) UpdateQuestModelMasterAsync(
	request *UpdateQuestModelMasterRequest,
	callback chan<- UpdateQuestModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questModelMaster",
    		"function": "updateQuestModelMaster",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
    if request.QuestName != nil && *request.QuestName != "" {
        bodies["questName"] = *request.QuestName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.Contents != nil {
        var _contents []interface {}
        for _, item := range request.Contents {
            _contents = append(_contents, item)
        }
        bodies["contents"] = _contents
    }
    if request.ChallengePeriodEventId != nil && *request.ChallengePeriodEventId != "" {
        bodies["challengePeriodEventId"] = *request.ChallengePeriodEventId
    }
    if request.ConsumeActions != nil {
        var _consumeActions []interface {}
        for _, item := range request.ConsumeActions {
            _consumeActions = append(_consumeActions, item)
        }
        bodies["consumeActions"] = _consumeActions
    }
    if request.FailedAcquireActions != nil {
        var _failedAcquireActions []interface {}
        for _, item := range request.FailedAcquireActions {
            _failedAcquireActions = append(_failedAcquireActions, item)
        }
        bodies["failedAcquireActions"] = _failedAcquireActions
    }
    if request.PremiseQuestNames != nil {
        var _premiseQuestNames []interface {}
        for _, item := range request.PremiseQuestNames {
            _premiseQuestNames = append(_premiseQuestNames, item)
        }
        bodies["premiseQuestNames"] = _premiseQuestNames
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateQuestModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) UpdateQuestModelMaster(
	request *UpdateQuestModelMasterRequest,
) (*UpdateQuestModelMasterResult, error) {
	callback := make(chan UpdateQuestModelMasterAsyncResult, 1)
	go p.UpdateQuestModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) deleteQuestModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteQuestModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteQuestModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteQuestModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteQuestModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteQuestModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DeleteQuestModelMasterAsync(
	request *DeleteQuestModelMasterRequest,
	callback chan<- DeleteQuestModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questModelMaster",
    		"function": "deleteQuestModelMaster",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
    if request.QuestName != nil && *request.QuestName != "" {
        bodies["questName"] = *request.QuestName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteQuestModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DeleteQuestModelMaster(
	request *DeleteQuestModelMasterRequest,
) (*DeleteQuestModelMasterResult, error) {
	callback := make(chan DeleteQuestModelMasterAsyncResult, 1)
	go p.DeleteQuestModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) exportMasterAsyncHandler(
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

func (p Gs2QuestWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "currentQuestMaster",
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

func (p Gs2QuestWebSocketClient) ExportMaster(
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

func (p Gs2QuestWebSocketClient) getCurrentQuestMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentQuestMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentQuestMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentQuestMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentQuestMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentQuestMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) GetCurrentQuestMasterAsync(
	request *GetCurrentQuestMasterRequest,
	callback chan<- GetCurrentQuestMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "currentQuestMaster",
    		"function": "getCurrentQuestMaster",
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

	go p.getCurrentQuestMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) GetCurrentQuestMaster(
	request *GetCurrentQuestMasterRequest,
) (*GetCurrentQuestMasterResult, error) {
	callback := make(chan GetCurrentQuestMasterAsyncResult, 1)
	go p.GetCurrentQuestMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) updateCurrentQuestMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentQuestMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentQuestMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentQuestMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentQuestMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentQuestMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) UpdateCurrentQuestMasterAsync(
	request *UpdateCurrentQuestMasterRequest,
	callback chan<- UpdateCurrentQuestMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "currentQuestMaster",
    		"function": "updateCurrentQuestMaster",
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

	go p.updateCurrentQuestMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) UpdateCurrentQuestMaster(
	request *UpdateCurrentQuestMasterRequest,
) (*UpdateCurrentQuestMasterResult, error) {
	callback := make(chan UpdateCurrentQuestMasterAsyncResult, 1)
	go p.UpdateCurrentQuestMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) updateCurrentQuestMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentQuestMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentQuestMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentQuestMasterFromGitHubResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentQuestMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentQuestMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) UpdateCurrentQuestMasterFromGitHubAsync(
	request *UpdateCurrentQuestMasterFromGitHubRequest,
	callback chan<- UpdateCurrentQuestMasterFromGitHubAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "currentQuestMaster",
    		"function": "updateCurrentQuestMasterFromGitHub",
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

	go p.updateCurrentQuestMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) UpdateCurrentQuestMasterFromGitHub(
	request *UpdateCurrentQuestMasterFromGitHubRequest,
) (*UpdateCurrentQuestMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentQuestMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentQuestMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) describeProgressesByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeProgressesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeProgressesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeProgressesByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeProgressesByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeProgressesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DescribeProgressesByUserIdAsync(
	request *DescribeProgressesByUserIdRequest,
	callback chan<- DescribeProgressesByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "progress",
    		"function": "describeProgressesByUserId",
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

	go p.describeProgressesByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DescribeProgressesByUserId(
	request *DescribeProgressesByUserIdRequest,
) (*DescribeProgressesByUserIdResult, error) {
	callback := make(chan DescribeProgressesByUserIdAsyncResult, 1)
	go p.DescribeProgressesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) createProgressByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateProgressByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateProgressByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateProgressByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateProgressByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateProgressByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) CreateProgressByUserIdAsync(
	request *CreateProgressByUserIdRequest,
	callback chan<- CreateProgressByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "progress",
    		"function": "createProgressByUserId",
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
    if request.QuestModelId != nil && *request.QuestModelId != "" {
        bodies["questModelId"] = *request.QuestModelId
    }
    if request.Force != nil {
        bodies["force"] = *request.Force
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

	go p.createProgressByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) CreateProgressByUserId(
	request *CreateProgressByUserIdRequest,
) (*CreateProgressByUserIdResult, error) {
	callback := make(chan CreateProgressByUserIdAsyncResult, 1)
	go p.CreateProgressByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) getProgressAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetProgressAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetProgressAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetProgressResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetProgressAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetProgressAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) GetProgressAsync(
	request *GetProgressRequest,
	callback chan<- GetProgressAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "progress",
    		"function": "getProgress",
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

	go p.getProgressAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) GetProgress(
	request *GetProgressRequest,
) (*GetProgressResult, error) {
	callback := make(chan GetProgressAsyncResult, 1)
	go p.GetProgressAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) getProgressByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetProgressByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetProgressByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetProgressByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetProgressByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetProgressByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) GetProgressByUserIdAsync(
	request *GetProgressByUserIdRequest,
	callback chan<- GetProgressByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "progress",
    		"function": "getProgressByUserId",
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

	go p.getProgressByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) GetProgressByUserId(
	request *GetProgressByUserIdRequest,
) (*GetProgressByUserIdResult, error) {
	callback := make(chan GetProgressByUserIdAsyncResult, 1)
	go p.GetProgressByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) startAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- StartAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- StartAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result StartResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- StartAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- StartAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) StartAsync(
	request *StartRequest,
	callback chan<- StartAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "progress",
    		"function": "start",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
    if request.QuestName != nil && *request.QuestName != "" {
        bodies["questName"] = *request.QuestName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.Force != nil {
        bodies["force"] = *request.Force
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

	go p.startAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) Start(
	request *StartRequest,
) (*StartResult, error) {
	callback := make(chan StartAsyncResult, 1)
	go p.StartAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) startByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- StartByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- StartByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result StartByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- StartByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- StartByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) StartByUserIdAsync(
	request *StartByUserIdRequest,
	callback chan<- StartByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "progress",
    		"function": "startByUserId",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
    if request.QuestName != nil && *request.QuestName != "" {
        bodies["questName"] = *request.QuestName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.Force != nil {
        bodies["force"] = *request.Force
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

	go p.startByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) StartByUserId(
	request *StartByUserIdRequest,
) (*StartByUserIdResult, error) {
	callback := make(chan StartByUserIdAsyncResult, 1)
	go p.StartByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) endAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- EndAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- EndAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result EndResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- EndAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- EndAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) EndAsync(
	request *EndRequest,
	callback chan<- EndAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "progress",
    		"function": "end",
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
    if request.TransactionId != nil && *request.TransactionId != "" {
        bodies["transactionId"] = *request.TransactionId
    }
    if request.Rewards != nil {
        var _rewards []interface {}
        for _, item := range request.Rewards {
            _rewards = append(_rewards, item)
        }
        bodies["rewards"] = _rewards
    }
    if request.IsComplete != nil {
        bodies["isComplete"] = *request.IsComplete
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

	go p.endAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) End(
	request *EndRequest,
) (*EndResult, error) {
	callback := make(chan EndAsyncResult, 1)
	go p.EndAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) endByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- EndByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- EndByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result EndByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- EndByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- EndByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) EndByUserIdAsync(
	request *EndByUserIdRequest,
	callback chan<- EndByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "progress",
    		"function": "endByUserId",
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
    if request.TransactionId != nil && *request.TransactionId != "" {
        bodies["transactionId"] = *request.TransactionId
    }
    if request.Rewards != nil {
        var _rewards []interface {}
        for _, item := range request.Rewards {
            _rewards = append(_rewards, item)
        }
        bodies["rewards"] = _rewards
    }
    if request.IsComplete != nil {
        bodies["isComplete"] = *request.IsComplete
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

	go p.endByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) EndByUserId(
	request *EndByUserIdRequest,
) (*EndByUserIdResult, error) {
	callback := make(chan EndByUserIdAsyncResult, 1)
	go p.EndByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) deleteProgressAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteProgressAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteProgressAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteProgressResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteProgressAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteProgressAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DeleteProgressAsync(
	request *DeleteProgressRequest,
	callback chan<- DeleteProgressAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "progress",
    		"function": "deleteProgress",
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

	go p.deleteProgressAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DeleteProgress(
	request *DeleteProgressRequest,
) (*DeleteProgressResult, error) {
	callback := make(chan DeleteProgressAsyncResult, 1)
	go p.DeleteProgressAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) deleteProgressByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteProgressByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteProgressByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteProgressByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteProgressByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteProgressByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DeleteProgressByUserIdAsync(
	request *DeleteProgressByUserIdRequest,
	callback chan<- DeleteProgressByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "progress",
    		"function": "deleteProgressByUserId",
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

	go p.deleteProgressByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DeleteProgressByUserId(
	request *DeleteProgressByUserIdRequest,
) (*DeleteProgressByUserIdResult, error) {
	callback := make(chan DeleteProgressByUserIdAsyncResult, 1)
	go p.DeleteProgressByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) createProgressByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateProgressByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateProgressByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateProgressByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateProgressByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateProgressByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) CreateProgressByStampSheetAsync(
	request *CreateProgressByStampSheetRequest,
	callback chan<- CreateProgressByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "progress",
    		"function": "createProgressByStampSheet",
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

	go p.createProgressByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) CreateProgressByStampSheet(
	request *CreateProgressByStampSheetRequest,
) (*CreateProgressByStampSheetResult, error) {
	callback := make(chan CreateProgressByStampSheetAsyncResult, 1)
	go p.CreateProgressByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) deleteProgressByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteProgressByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteProgressByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteProgressByStampTaskResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteProgressByStampTaskAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteProgressByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DeleteProgressByStampTaskAsync(
	request *DeleteProgressByStampTaskRequest,
	callback chan<- DeleteProgressByStampTaskAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "progress",
    		"function": "deleteProgressByStampTask",
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

	go p.deleteProgressByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DeleteProgressByStampTask(
	request *DeleteProgressByStampTaskRequest,
) (*DeleteProgressByStampTaskResult, error) {
	callback := make(chan DeleteProgressByStampTaskAsyncResult, 1)
	go p.DeleteProgressByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) describeCompletedQuestListsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeCompletedQuestListsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCompletedQuestListsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCompletedQuestListsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeCompletedQuestListsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeCompletedQuestListsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DescribeCompletedQuestListsAsync(
	request *DescribeCompletedQuestListsRequest,
	callback chan<- DescribeCompletedQuestListsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "completedQuestList",
    		"function": "describeCompletedQuestLists",
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

	go p.describeCompletedQuestListsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DescribeCompletedQuestLists(
	request *DescribeCompletedQuestListsRequest,
) (*DescribeCompletedQuestListsResult, error) {
	callback := make(chan DescribeCompletedQuestListsAsyncResult, 1)
	go p.DescribeCompletedQuestListsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) describeCompletedQuestListsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeCompletedQuestListsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCompletedQuestListsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCompletedQuestListsByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeCompletedQuestListsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeCompletedQuestListsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DescribeCompletedQuestListsByUserIdAsync(
	request *DescribeCompletedQuestListsByUserIdRequest,
	callback chan<- DescribeCompletedQuestListsByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "completedQuestList",
    		"function": "describeCompletedQuestListsByUserId",
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

	go p.describeCompletedQuestListsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DescribeCompletedQuestListsByUserId(
	request *DescribeCompletedQuestListsByUserIdRequest,
) (*DescribeCompletedQuestListsByUserIdResult, error) {
	callback := make(chan DescribeCompletedQuestListsByUserIdAsyncResult, 1)
	go p.DescribeCompletedQuestListsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) getCompletedQuestListAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCompletedQuestListAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCompletedQuestListAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCompletedQuestListResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCompletedQuestListAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCompletedQuestListAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) GetCompletedQuestListAsync(
	request *GetCompletedQuestListRequest,
	callback chan<- GetCompletedQuestListAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "completedQuestList",
    		"function": "getCompletedQuestList",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
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

	go p.getCompletedQuestListAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) GetCompletedQuestList(
	request *GetCompletedQuestListRequest,
) (*GetCompletedQuestListResult, error) {
	callback := make(chan GetCompletedQuestListAsyncResult, 1)
	go p.GetCompletedQuestListAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) getCompletedQuestListByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCompletedQuestListByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCompletedQuestListByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCompletedQuestListByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCompletedQuestListByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCompletedQuestListByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) GetCompletedQuestListByUserIdAsync(
	request *GetCompletedQuestListByUserIdRequest,
	callback chan<- GetCompletedQuestListByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "completedQuestList",
    		"function": "getCompletedQuestListByUserId",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getCompletedQuestListByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) GetCompletedQuestListByUserId(
	request *GetCompletedQuestListByUserIdRequest,
) (*GetCompletedQuestListByUserIdResult, error) {
	callback := make(chan GetCompletedQuestListByUserIdAsyncResult, 1)
	go p.GetCompletedQuestListByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) deleteCompletedQuestListByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteCompletedQuestListByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteCompletedQuestListByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteCompletedQuestListByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteCompletedQuestListByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteCompletedQuestListByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DeleteCompletedQuestListByUserIdAsync(
	request *DeleteCompletedQuestListByUserIdRequest,
	callback chan<- DeleteCompletedQuestListByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "completedQuestList",
    		"function": "deleteCompletedQuestListByUserId",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteCompletedQuestListByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DeleteCompletedQuestListByUserId(
	request *DeleteCompletedQuestListByUserIdRequest,
) (*DeleteCompletedQuestListByUserIdResult, error) {
	callback := make(chan DeleteCompletedQuestListByUserIdAsyncResult, 1)
	go p.DeleteCompletedQuestListByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) describeQuestGroupModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeQuestGroupModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeQuestGroupModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeQuestGroupModelsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeQuestGroupModelsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeQuestGroupModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DescribeQuestGroupModelsAsync(
	request *DescribeQuestGroupModelsRequest,
	callback chan<- DescribeQuestGroupModelsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questGroupModel",
    		"function": "describeQuestGroupModels",
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

	go p.describeQuestGroupModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DescribeQuestGroupModels(
	request *DescribeQuestGroupModelsRequest,
) (*DescribeQuestGroupModelsResult, error) {
	callback := make(chan DescribeQuestGroupModelsAsyncResult, 1)
	go p.DescribeQuestGroupModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) getQuestGroupModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetQuestGroupModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetQuestGroupModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetQuestGroupModelResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetQuestGroupModelAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetQuestGroupModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) GetQuestGroupModelAsync(
	request *GetQuestGroupModelRequest,
	callback chan<- GetQuestGroupModelAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questGroupModel",
    		"function": "getQuestGroupModel",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getQuestGroupModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) GetQuestGroupModel(
	request *GetQuestGroupModelRequest,
) (*GetQuestGroupModelResult, error) {
	callback := make(chan GetQuestGroupModelAsyncResult, 1)
	go p.GetQuestGroupModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) describeQuestModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeQuestModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeQuestModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeQuestModelsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeQuestModelsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeQuestModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) DescribeQuestModelsAsync(
	request *DescribeQuestModelsRequest,
	callback chan<- DescribeQuestModelsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questModel",
    		"function": "describeQuestModels",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeQuestModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) DescribeQuestModels(
	request *DescribeQuestModelsRequest,
) (*DescribeQuestModelsResult, error) {
	callback := make(chan DescribeQuestModelsAsyncResult, 1)
	go p.DescribeQuestModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2QuestWebSocketClient) getQuestModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetQuestModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetQuestModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetQuestModelResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetQuestModelAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetQuestModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2QuestWebSocketClient) GetQuestModelAsync(
	request *GetQuestModelRequest,
	callback chan<- GetQuestModelAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "quest",
    		"component": "questModel",
    		"function": "getQuestModel",
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
    if request.QuestGroupName != nil && *request.QuestGroupName != "" {
        bodies["questGroupName"] = *request.QuestGroupName
    }
    if request.QuestName != nil && *request.QuestName != "" {
        bodies["questName"] = *request.QuestName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getQuestModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2QuestWebSocketClient) GetQuestModel(
	request *GetQuestModelRequest,
) (*GetQuestModelResult, error) {
	callback := make(chan GetQuestModelAsyncResult, 1)
	go p.GetQuestModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
