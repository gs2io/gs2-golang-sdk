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

package inbox

import (
	"encoding/json"
	"github.com/google/uuid"
	"core"
)

type Gs2InboxWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2InboxWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2InboxWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2InboxWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
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

func (p Gs2InboxWebSocketClient) DescribeNamespaces(
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

func (p Gs2InboxWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2InboxWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
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
    if request.IsAutomaticDeletingEnabled != nil {
        bodies["isAutomaticDeletingEnabled"] = *request.IsAutomaticDeletingEnabled
    }
    if request.ReceiveMessageScript != nil {
        bodies["receiveMessageScript"] = request.ReceiveMessageScript.ToDict()
    }
    if request.ReadMessageScript != nil {
        bodies["readMessageScript"] = request.ReadMessageScript.ToDict()
    }
    if request.DeleteMessageScript != nil {
        bodies["deleteMessageScript"] = request.DeleteMessageScript.ToDict()
    }
    if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
        bodies["queueNamespaceId"] = *request.QueueNamespaceId
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
    if request.ReceiveNotification != nil {
        bodies["receiveNotification"] = request.ReceiveNotification.ToDict()
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

func (p Gs2InboxWebSocketClient) CreateNamespace(
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

func (p Gs2InboxWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2InboxWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
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

func (p Gs2InboxWebSocketClient) GetNamespaceStatus(
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

func (p Gs2InboxWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2InboxWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
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

func (p Gs2InboxWebSocketClient) GetNamespace(
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

func (p Gs2InboxWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2InboxWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
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
    if request.IsAutomaticDeletingEnabled != nil {
        bodies["isAutomaticDeletingEnabled"] = *request.IsAutomaticDeletingEnabled
    }
    if request.ReceiveMessageScript != nil {
        bodies["receiveMessageScript"] = request.ReceiveMessageScript.ToDict()
    }
    if request.ReadMessageScript != nil {
        bodies["readMessageScript"] = request.ReadMessageScript.ToDict()
    }
    if request.DeleteMessageScript != nil {
        bodies["deleteMessageScript"] = request.DeleteMessageScript.ToDict()
    }
    if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
        bodies["queueNamespaceId"] = *request.QueueNamespaceId
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
    if request.ReceiveNotification != nil {
        bodies["receiveNotification"] = request.ReceiveNotification.ToDict()
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

func (p Gs2InboxWebSocketClient) UpdateNamespace(
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

func (p Gs2InboxWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2InboxWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
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

func (p Gs2InboxWebSocketClient) DeleteNamespace(
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

func (p Gs2InboxWebSocketClient) describeMessagesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMessagesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMessagesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMessagesResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeMessagesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeMessagesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) DescribeMessagesAsync(
	request *DescribeMessagesRequest,
	callback chan<- DescribeMessagesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "describeMessages",
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

	go p.describeMessagesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) DescribeMessages(
	request *DescribeMessagesRequest,
) (*DescribeMessagesResult, error) {
	callback := make(chan DescribeMessagesAsyncResult, 1)
	go p.DescribeMessagesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) describeMessagesByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMessagesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMessagesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMessagesByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeMessagesByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeMessagesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) DescribeMessagesByUserIdAsync(
	request *DescribeMessagesByUserIdRequest,
	callback chan<- DescribeMessagesByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "describeMessagesByUserId",
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

	go p.describeMessagesByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) DescribeMessagesByUserId(
	request *DescribeMessagesByUserIdRequest,
) (*DescribeMessagesByUserIdResult, error) {
	callback := make(chan DescribeMessagesByUserIdAsyncResult, 1)
	go p.DescribeMessagesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) sendMessageByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SendMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SendMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SendMessageByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SendMessageByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SendMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) SendMessageByUserIdAsync(
	request *SendMessageByUserIdRequest,
	callback chan<- SendMessageByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "sendMessageByUserId",
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
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.ReadAcquireActions != nil {
        var _readAcquireActions []interface {}
        for _, item := range request.ReadAcquireActions {
            _readAcquireActions = append(_readAcquireActions, item)
        }
        bodies["readAcquireActions"] = _readAcquireActions
    }
    if request.ExpiresAt != nil {
        bodies["expiresAt"] = *request.ExpiresAt
    }
    if request.ExpiresTimeSpan != nil {
        bodies["expiresTimeSpan"] = request.ExpiresTimeSpan.ToDict()
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.sendMessageByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) SendMessageByUserId(
	request *SendMessageByUserIdRequest,
) (*SendMessageByUserIdResult, error) {
	callback := make(chan SendMessageByUserIdAsyncResult, 1)
	go p.SendMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) getMessageAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMessageResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetMessageAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) GetMessageAsync(
	request *GetMessageRequest,
	callback chan<- GetMessageAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "getMessage",
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
    if request.MessageName != nil && *request.MessageName != "" {
        bodies["messageName"] = *request.MessageName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getMessageAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) GetMessage(
	request *GetMessageRequest,
) (*GetMessageResult, error) {
	callback := make(chan GetMessageAsyncResult, 1)
	go p.GetMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) getMessageByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMessageByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetMessageByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) GetMessageByUserIdAsync(
	request *GetMessageByUserIdRequest,
	callback chan<- GetMessageByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "getMessageByUserId",
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
    if request.MessageName != nil && *request.MessageName != "" {
        bodies["messageName"] = *request.MessageName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getMessageByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) GetMessageByUserId(
	request *GetMessageByUserIdRequest,
) (*GetMessageByUserIdResult, error) {
	callback := make(chan GetMessageByUserIdAsyncResult, 1)
	go p.GetMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) receiveGlobalMessageAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReceiveGlobalMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveGlobalMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveGlobalMessageResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ReceiveGlobalMessageAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ReceiveGlobalMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) ReceiveGlobalMessageAsync(
	request *ReceiveGlobalMessageRequest,
	callback chan<- ReceiveGlobalMessageAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "receiveGlobalMessage",
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

	go p.receiveGlobalMessageAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) ReceiveGlobalMessage(
	request *ReceiveGlobalMessageRequest,
) (*ReceiveGlobalMessageResult, error) {
	callback := make(chan ReceiveGlobalMessageAsyncResult, 1)
	go p.ReceiveGlobalMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) receiveGlobalMessageByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReceiveGlobalMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveGlobalMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveGlobalMessageByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ReceiveGlobalMessageByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ReceiveGlobalMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) ReceiveGlobalMessageByUserIdAsync(
	request *ReceiveGlobalMessageByUserIdRequest,
	callback chan<- ReceiveGlobalMessageByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "receiveGlobalMessageByUserId",
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

	go p.receiveGlobalMessageByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) ReceiveGlobalMessageByUserId(
	request *ReceiveGlobalMessageByUserIdRequest,
) (*ReceiveGlobalMessageByUserIdResult, error) {
	callback := make(chan ReceiveGlobalMessageByUserIdAsyncResult, 1)
	go p.ReceiveGlobalMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) openMessageAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- OpenMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- OpenMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result OpenMessageResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- OpenMessageAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- OpenMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) OpenMessageAsync(
	request *OpenMessageRequest,
	callback chan<- OpenMessageAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "openMessage",
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
    if request.MessageName != nil && *request.MessageName != "" {
        bodies["messageName"] = *request.MessageName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.openMessageAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) OpenMessage(
	request *OpenMessageRequest,
) (*OpenMessageResult, error) {
	callback := make(chan OpenMessageAsyncResult, 1)
	go p.OpenMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) openMessageByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- OpenMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- OpenMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result OpenMessageByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- OpenMessageByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- OpenMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) OpenMessageByUserIdAsync(
	request *OpenMessageByUserIdRequest,
	callback chan<- OpenMessageByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "openMessageByUserId",
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
    if request.MessageName != nil && *request.MessageName != "" {
        bodies["messageName"] = *request.MessageName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.openMessageByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) OpenMessageByUserId(
	request *OpenMessageByUserIdRequest,
) (*OpenMessageByUserIdResult, error) {
	callback := make(chan OpenMessageByUserIdAsyncResult, 1)
	go p.OpenMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) readMessageAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReadMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReadMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReadMessageResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ReadMessageAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ReadMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) ReadMessageAsync(
	request *ReadMessageRequest,
	callback chan<- ReadMessageAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "readMessage",
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
    if request.MessageName != nil && *request.MessageName != "" {
        bodies["messageName"] = *request.MessageName
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

	go p.readMessageAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) ReadMessage(
	request *ReadMessageRequest,
) (*ReadMessageResult, error) {
	callback := make(chan ReadMessageAsyncResult, 1)
	go p.ReadMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) readMessageByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReadMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReadMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReadMessageByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ReadMessageByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ReadMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) ReadMessageByUserIdAsync(
	request *ReadMessageByUserIdRequest,
	callback chan<- ReadMessageByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "readMessageByUserId",
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
    if request.MessageName != nil && *request.MessageName != "" {
        bodies["messageName"] = *request.MessageName
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

	go p.readMessageByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) ReadMessageByUserId(
	request *ReadMessageByUserIdRequest,
) (*ReadMessageByUserIdResult, error) {
	callback := make(chan ReadMessageByUserIdAsyncResult, 1)
	go p.ReadMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) deleteMessageAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMessageResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteMessageAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) DeleteMessageAsync(
	request *DeleteMessageRequest,
	callback chan<- DeleteMessageAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "deleteMessage",
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
    if request.MessageName != nil && *request.MessageName != "" {
        bodies["messageName"] = *request.MessageName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.deleteMessageAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) DeleteMessage(
	request *DeleteMessageRequest,
) (*DeleteMessageResult, error) {
	callback := make(chan DeleteMessageAsyncResult, 1)
	go p.DeleteMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) deleteMessageByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMessageByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteMessageByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) DeleteMessageByUserIdAsync(
	request *DeleteMessageByUserIdRequest,
	callback chan<- DeleteMessageByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "deleteMessageByUserId",
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
    if request.MessageName != nil && *request.MessageName != "" {
        bodies["messageName"] = *request.MessageName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteMessageByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) DeleteMessageByUserId(
	request *DeleteMessageByUserIdRequest,
) (*DeleteMessageByUserIdResult, error) {
	callback := make(chan DeleteMessageByUserIdAsyncResult, 1)
	go p.DeleteMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) openByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- OpenByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- OpenByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result OpenByStampTaskResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- OpenByStampTaskAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- OpenByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) OpenByStampTaskAsync(
	request *OpenByStampTaskRequest,
	callback chan<- OpenByStampTaskAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "message",
    		"function": "openByStampTask",
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

	go p.openByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) OpenByStampTask(
	request *OpenByStampTaskRequest,
) (*OpenByStampTaskResult, error) {
	callback := make(chan OpenByStampTaskAsyncResult, 1)
	go p.OpenByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) exportMasterAsyncHandler(
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

func (p Gs2InboxWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "currentMessageMaster",
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

func (p Gs2InboxWebSocketClient) ExportMaster(
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

func (p Gs2InboxWebSocketClient) getCurrentMessageMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentMessageMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentMessageMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) GetCurrentMessageMasterAsync(
	request *GetCurrentMessageMasterRequest,
	callback chan<- GetCurrentMessageMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "currentMessageMaster",
    		"function": "getCurrentMessageMaster",
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

	go p.getCurrentMessageMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) GetCurrentMessageMaster(
	request *GetCurrentMessageMasterRequest,
) (*GetCurrentMessageMasterResult, error) {
	callback := make(chan GetCurrentMessageMasterAsyncResult, 1)
	go p.GetCurrentMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) updateCurrentMessageMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentMessageMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentMessageMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) UpdateCurrentMessageMasterAsync(
	request *UpdateCurrentMessageMasterRequest,
	callback chan<- UpdateCurrentMessageMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "currentMessageMaster",
    		"function": "updateCurrentMessageMaster",
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

	go p.updateCurrentMessageMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) UpdateCurrentMessageMaster(
	request *UpdateCurrentMessageMasterRequest,
) (*UpdateCurrentMessageMasterResult, error) {
	callback := make(chan UpdateCurrentMessageMasterAsyncResult, 1)
	go p.UpdateCurrentMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) updateCurrentMessageMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentMessageMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentMessageMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentMessageMasterFromGitHubResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentMessageMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentMessageMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) UpdateCurrentMessageMasterFromGitHubAsync(
	request *UpdateCurrentMessageMasterFromGitHubRequest,
	callback chan<- UpdateCurrentMessageMasterFromGitHubAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "currentMessageMaster",
    		"function": "updateCurrentMessageMasterFromGitHub",
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

	go p.updateCurrentMessageMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) UpdateCurrentMessageMasterFromGitHub(
	request *UpdateCurrentMessageMasterFromGitHubRequest,
) (*UpdateCurrentMessageMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentMessageMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentMessageMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) describeGlobalMessageMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGlobalMessageMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalMessageMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalMessageMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeGlobalMessageMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeGlobalMessageMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) DescribeGlobalMessageMastersAsync(
	request *DescribeGlobalMessageMastersRequest,
	callback chan<- DescribeGlobalMessageMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "globalMessageMaster",
    		"function": "describeGlobalMessageMasters",
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

	go p.describeGlobalMessageMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) DescribeGlobalMessageMasters(
	request *DescribeGlobalMessageMastersRequest,
) (*DescribeGlobalMessageMastersResult, error) {
	callback := make(chan DescribeGlobalMessageMastersAsyncResult, 1)
	go p.DescribeGlobalMessageMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) createGlobalMessageMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateGlobalMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGlobalMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGlobalMessageMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateGlobalMessageMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateGlobalMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) CreateGlobalMessageMasterAsync(
	request *CreateGlobalMessageMasterRequest,
	callback chan<- CreateGlobalMessageMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "globalMessageMaster",
    		"function": "createGlobalMessageMaster",
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
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.ReadAcquireActions != nil {
        var _readAcquireActions []interface {}
        for _, item := range request.ReadAcquireActions {
            _readAcquireActions = append(_readAcquireActions, item)
        }
        bodies["readAcquireActions"] = _readAcquireActions
    }
    if request.ExpiresTimeSpan != nil {
        bodies["expiresTimeSpan"] = request.ExpiresTimeSpan.ToDict()
    }
    if request.ExpiresAt != nil {
        bodies["expiresAt"] = *request.ExpiresAt
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createGlobalMessageMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) CreateGlobalMessageMaster(
	request *CreateGlobalMessageMasterRequest,
) (*CreateGlobalMessageMasterResult, error) {
	callback := make(chan CreateGlobalMessageMasterAsyncResult, 1)
	go p.CreateGlobalMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) getGlobalMessageMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGlobalMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalMessageMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetGlobalMessageMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetGlobalMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) GetGlobalMessageMasterAsync(
	request *GetGlobalMessageMasterRequest,
	callback chan<- GetGlobalMessageMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "globalMessageMaster",
    		"function": "getGlobalMessageMaster",
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
    if request.GlobalMessageName != nil && *request.GlobalMessageName != "" {
        bodies["globalMessageName"] = *request.GlobalMessageName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getGlobalMessageMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) GetGlobalMessageMaster(
	request *GetGlobalMessageMasterRequest,
) (*GetGlobalMessageMasterResult, error) {
	callback := make(chan GetGlobalMessageMasterAsyncResult, 1)
	go p.GetGlobalMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) updateGlobalMessageMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateGlobalMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGlobalMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGlobalMessageMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateGlobalMessageMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateGlobalMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) UpdateGlobalMessageMasterAsync(
	request *UpdateGlobalMessageMasterRequest,
	callback chan<- UpdateGlobalMessageMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "globalMessageMaster",
    		"function": "updateGlobalMessageMaster",
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
    if request.GlobalMessageName != nil && *request.GlobalMessageName != "" {
        bodies["globalMessageName"] = *request.GlobalMessageName
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.ReadAcquireActions != nil {
        var _readAcquireActions []interface {}
        for _, item := range request.ReadAcquireActions {
            _readAcquireActions = append(_readAcquireActions, item)
        }
        bodies["readAcquireActions"] = _readAcquireActions
    }
    if request.ExpiresTimeSpan != nil {
        bodies["expiresTimeSpan"] = request.ExpiresTimeSpan.ToDict()
    }
    if request.ExpiresAt != nil {
        bodies["expiresAt"] = *request.ExpiresAt
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateGlobalMessageMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) UpdateGlobalMessageMaster(
	request *UpdateGlobalMessageMasterRequest,
) (*UpdateGlobalMessageMasterResult, error) {
	callback := make(chan UpdateGlobalMessageMasterAsyncResult, 1)
	go p.UpdateGlobalMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) deleteGlobalMessageMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteGlobalMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGlobalMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGlobalMessageMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteGlobalMessageMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteGlobalMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) DeleteGlobalMessageMasterAsync(
	request *DeleteGlobalMessageMasterRequest,
	callback chan<- DeleteGlobalMessageMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "globalMessageMaster",
    		"function": "deleteGlobalMessageMaster",
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
    if request.GlobalMessageName != nil && *request.GlobalMessageName != "" {
        bodies["globalMessageName"] = *request.GlobalMessageName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteGlobalMessageMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) DeleteGlobalMessageMaster(
	request *DeleteGlobalMessageMasterRequest,
) (*DeleteGlobalMessageMasterResult, error) {
	callback := make(chan DeleteGlobalMessageMasterAsyncResult, 1)
	go p.DeleteGlobalMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) describeGlobalMessagesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGlobalMessagesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalMessagesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalMessagesResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeGlobalMessagesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeGlobalMessagesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) DescribeGlobalMessagesAsync(
	request *DescribeGlobalMessagesRequest,
	callback chan<- DescribeGlobalMessagesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "globalMessage",
    		"function": "describeGlobalMessages",
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

	go p.describeGlobalMessagesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) DescribeGlobalMessages(
	request *DescribeGlobalMessagesRequest,
) (*DescribeGlobalMessagesResult, error) {
	callback := make(chan DescribeGlobalMessagesAsyncResult, 1)
	go p.DescribeGlobalMessagesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) getGlobalMessageAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGlobalMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalMessageResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetGlobalMessageAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetGlobalMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) GetGlobalMessageAsync(
	request *GetGlobalMessageRequest,
	callback chan<- GetGlobalMessageAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "globalMessage",
    		"function": "getGlobalMessage",
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
    if request.GlobalMessageName != nil && *request.GlobalMessageName != "" {
        bodies["globalMessageName"] = *request.GlobalMessageName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getGlobalMessageAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) GetGlobalMessage(
	request *GetGlobalMessageRequest,
) (*GetGlobalMessageResult, error) {
	callback := make(chan GetGlobalMessageAsyncResult, 1)
	go p.GetGlobalMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) getReceivedByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetReceivedByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReceivedByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReceivedByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetReceivedByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetReceivedByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) GetReceivedByUserIdAsync(
	request *GetReceivedByUserIdRequest,
	callback chan<- GetReceivedByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "received",
    		"function": "getReceivedByUserId",
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

	go p.getReceivedByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) GetReceivedByUserId(
	request *GetReceivedByUserIdRequest,
) (*GetReceivedByUserIdResult, error) {
	callback := make(chan GetReceivedByUserIdAsyncResult, 1)
	go p.GetReceivedByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) updateReceivedByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateReceivedByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateReceivedByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateReceivedByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateReceivedByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateReceivedByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) UpdateReceivedByUserIdAsync(
	request *UpdateReceivedByUserIdRequest,
	callback chan<- UpdateReceivedByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "received",
    		"function": "updateReceivedByUserId",
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
    if request.ReceivedGlobalMessageNames != nil {
        var _receivedGlobalMessageNames []interface {}
        for _, item := range request.ReceivedGlobalMessageNames {
            _receivedGlobalMessageNames = append(_receivedGlobalMessageNames, item)
        }
        bodies["receivedGlobalMessageNames"] = _receivedGlobalMessageNames
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateReceivedByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) UpdateReceivedByUserId(
	request *UpdateReceivedByUserIdRequest,
) (*UpdateReceivedByUserIdResult, error) {
	callback := make(chan UpdateReceivedByUserIdAsyncResult, 1)
	go p.UpdateReceivedByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2InboxWebSocketClient) deleteReceivedByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteReceivedByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteReceivedByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteReceivedByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteReceivedByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteReceivedByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxWebSocketClient) DeleteReceivedByUserIdAsync(
	request *DeleteReceivedByUserIdRequest,
	callback chan<- DeleteReceivedByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "inbox",
    		"component": "received",
    		"function": "deleteReceivedByUserId",
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

	go p.deleteReceivedByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2InboxWebSocketClient) DeleteReceivedByUserId(
	request *DeleteReceivedByUserIdRequest,
) (*DeleteReceivedByUserIdResult, error) {
	callback := make(chan DeleteReceivedByUserIdAsyncResult, 1)
	go p.DeleteReceivedByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
