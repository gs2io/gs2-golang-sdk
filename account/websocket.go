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

package account

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2AccountWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2AccountWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2AccountWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2AccountWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
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

func (p Gs2AccountWebSocketClient) DescribeNamespaces(
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

func (p Gs2AccountWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2AccountWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
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
    if request.ChangePasswordIfTakeOver != nil {
        bodies["changePasswordIfTakeOver"] = *request.ChangePasswordIfTakeOver
    }
    if request.DifferentUserIdForLoginAndDataRetention != nil {
        bodies["differentUserIdForLoginAndDataRetention"] = *request.DifferentUserIdForLoginAndDataRetention
    }
    if request.CreateAccountScript != nil {
        bodies["createAccountScript"] = request.CreateAccountScript.ToDict()
    }
    if request.AuthenticationScript != nil {
        bodies["authenticationScript"] = request.AuthenticationScript.ToDict()
    }
    if request.CreateTakeOverScript != nil {
        bodies["createTakeOverScript"] = request.CreateTakeOverScript.ToDict()
    }
    if request.DoTakeOverScript != nil {
        bodies["doTakeOverScript"] = request.DoTakeOverScript.ToDict()
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

func (p Gs2AccountWebSocketClient) CreateNamespace(
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

func (p Gs2AccountWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2AccountWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
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

func (p Gs2AccountWebSocketClient) GetNamespaceStatus(
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

func (p Gs2AccountWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2AccountWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
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

func (p Gs2AccountWebSocketClient) GetNamespace(
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

func (p Gs2AccountWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2AccountWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
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
    if request.ChangePasswordIfTakeOver != nil {
        bodies["changePasswordIfTakeOver"] = *request.ChangePasswordIfTakeOver
    }
    if request.CreateAccountScript != nil {
        bodies["createAccountScript"] = request.CreateAccountScript.ToDict()
    }
    if request.AuthenticationScript != nil {
        bodies["authenticationScript"] = request.AuthenticationScript.ToDict()
    }
    if request.CreateTakeOverScript != nil {
        bodies["createTakeOverScript"] = request.CreateTakeOverScript.ToDict()
    }
    if request.DoTakeOverScript != nil {
        bodies["doTakeOverScript"] = request.DoTakeOverScript.ToDict()
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

func (p Gs2AccountWebSocketClient) UpdateNamespace(
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

func (p Gs2AccountWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2AccountWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
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

func (p Gs2AccountWebSocketClient) DeleteNamespace(
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

func (p Gs2AccountWebSocketClient) describeAccountsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeAccountsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeAccountsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeAccountsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeAccountsAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeAccountsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) DescribeAccountsAsync(
	request *DescribeAccountsRequest,
	callback chan<- DescribeAccountsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "account",
    		"function": "describeAccounts",
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

	go p.describeAccountsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) DescribeAccounts(
	request *DescribeAccountsRequest,
) (*DescribeAccountsResult, error) {
	callback := make(chan DescribeAccountsAsyncResult, 1)
	go p.DescribeAccountsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) createAccountAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateAccountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateAccountAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateAccountResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateAccountAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- CreateAccountAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) CreateAccountAsync(
	request *CreateAccountRequest,
	callback chan<- CreateAccountAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "account",
    		"function": "createAccount",
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

	go p.createAccountAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) CreateAccount(
	request *CreateAccountRequest,
) (*CreateAccountResult, error) {
	callback := make(chan CreateAccountAsyncResult, 1)
	go p.CreateAccountAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) updateTimeOffsetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateTimeOffsetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateTimeOffsetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateTimeOffsetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateTimeOffsetAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- UpdateTimeOffsetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) UpdateTimeOffsetAsync(
	request *UpdateTimeOffsetRequest,
	callback chan<- UpdateTimeOffsetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "account",
    		"function": "updateTimeOffset",
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
    if request.TimeOffset != nil {
        bodies["timeOffset"] = *request.TimeOffset
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.updateTimeOffsetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) UpdateTimeOffset(
	request *UpdateTimeOffsetRequest,
) (*UpdateTimeOffsetResult, error) {
	callback := make(chan UpdateTimeOffsetAsyncResult, 1)
	go p.UpdateTimeOffsetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) updateBannedAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateBannedAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateBannedAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateBannedResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateBannedAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- UpdateBannedAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) UpdateBannedAsync(
	request *UpdateBannedRequest,
	callback chan<- UpdateBannedAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "account",
    		"function": "updateBanned",
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
    if request.Banned != nil {
        bodies["banned"] = *request.Banned
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.updateBannedAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) UpdateBanned(
	request *UpdateBannedRequest,
) (*UpdateBannedResult, error) {
	callback := make(chan UpdateBannedAsyncResult, 1)
	go p.UpdateBannedAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) getAccountAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetAccountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetAccountAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetAccountResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetAccountAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetAccountAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) GetAccountAsync(
	request *GetAccountRequest,
	callback chan<- GetAccountAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "account",
    		"function": "getAccount",
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

	go p.getAccountAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) GetAccount(
	request *GetAccountRequest,
) (*GetAccountResult, error) {
	callback := make(chan GetAccountAsyncResult, 1)
	go p.GetAccountAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) deleteAccountAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteAccountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteAccountAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteAccountResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteAccountAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteAccountAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) DeleteAccountAsync(
	request *DeleteAccountRequest,
	callback chan<- DeleteAccountAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "account",
    		"function": "deleteAccount",
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
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.deleteAccountAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) DeleteAccount(
	request *DeleteAccountRequest,
) (*DeleteAccountResult, error) {
	callback := make(chan DeleteAccountAsyncResult, 1)
	go p.DeleteAccountAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) authenticationAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AuthenticationAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AuthenticationAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AuthenticationResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AuthenticationAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
        gs2err, ok := asyncResult.Err.(core.Gs2Exception)
        if ok {
            if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "account.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
            }
            if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "account.banned.infinity" {
				asyncResult.Err = gs2err.SetClientError(BannedInfinity{})
            }
        }
    }
	callback <- AuthenticationAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) AuthenticationAsync(
	request *AuthenticationRequest,
	callback chan<- AuthenticationAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "account",
    		"function": "authentication",
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
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
    if request.Password != nil && *request.Password != "" {
        bodies["password"] = *request.Password
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.authenticationAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) Authentication(
	request *AuthenticationRequest,
) (*AuthenticationResult, error) {
	callback := make(chan AuthenticationAsyncResult, 1)
	go p.AuthenticationAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) describeTakeOversAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeTakeOversAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeTakeOversAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeTakeOversResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeTakeOversAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeTakeOversAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) DescribeTakeOversAsync(
	request *DescribeTakeOversRequest,
	callback chan<- DescribeTakeOversAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "takeOver",
    		"function": "describeTakeOvers",
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

	go p.describeTakeOversAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) DescribeTakeOvers(
	request *DescribeTakeOversRequest,
) (*DescribeTakeOversResult, error) {
	callback := make(chan DescribeTakeOversAsyncResult, 1)
	go p.DescribeTakeOversAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) describeTakeOversByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeTakeOversByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeTakeOversByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeTakeOversByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeTakeOversByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DescribeTakeOversByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) DescribeTakeOversByUserIdAsync(
	request *DescribeTakeOversByUserIdRequest,
	callback chan<- DescribeTakeOversByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "takeOver",
    		"function": "describeTakeOversByUserId",
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

	go p.describeTakeOversByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) DescribeTakeOversByUserId(
	request *DescribeTakeOversByUserIdRequest,
) (*DescribeTakeOversByUserIdResult, error) {
	callback := make(chan DescribeTakeOversByUserIdAsyncResult, 1)
	go p.DescribeTakeOversByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) createTakeOverAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateTakeOverAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateTakeOverAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateTakeOverResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateTakeOverAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- CreateTakeOverAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) CreateTakeOverAsync(
	request *CreateTakeOverRequest,
	callback chan<- CreateTakeOverAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "takeOver",
    		"function": "createTakeOver",
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
    if request.Type != nil {
        bodies["type"] = *request.Type
    }
    if request.UserIdentifier != nil && *request.UserIdentifier != "" {
        bodies["userIdentifier"] = *request.UserIdentifier
    }
    if request.Password != nil && *request.Password != "" {
        bodies["password"] = *request.Password
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

	go p.createTakeOverAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) CreateTakeOver(
	request *CreateTakeOverRequest,
) (*CreateTakeOverResult, error) {
	callback := make(chan CreateTakeOverAsyncResult, 1)
	go p.CreateTakeOverAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) createTakeOverByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateTakeOverByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateTakeOverByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateTakeOverByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateTakeOverByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- CreateTakeOverByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) CreateTakeOverByUserIdAsync(
	request *CreateTakeOverByUserIdRequest,
	callback chan<- CreateTakeOverByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "takeOver",
    		"function": "createTakeOverByUserId",
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
    if request.Type != nil {
        bodies["type"] = *request.Type
    }
    if request.UserIdentifier != nil && *request.UserIdentifier != "" {
        bodies["userIdentifier"] = *request.UserIdentifier
    }
    if request.Password != nil && *request.Password != "" {
        bodies["password"] = *request.Password
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.createTakeOverByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) CreateTakeOverByUserId(
	request *CreateTakeOverByUserIdRequest,
) (*CreateTakeOverByUserIdResult, error) {
	callback := make(chan CreateTakeOverByUserIdAsyncResult, 1)
	go p.CreateTakeOverByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) getTakeOverAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetTakeOverAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetTakeOverAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetTakeOverResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetTakeOverAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetTakeOverAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) GetTakeOverAsync(
	request *GetTakeOverRequest,
	callback chan<- GetTakeOverAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "takeOver",
    		"function": "getTakeOver",
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
    if request.Type != nil {
        bodies["type"] = *request.Type
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getTakeOverAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) GetTakeOver(
	request *GetTakeOverRequest,
) (*GetTakeOverResult, error) {
	callback := make(chan GetTakeOverAsyncResult, 1)
	go p.GetTakeOverAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) getTakeOverByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetTakeOverByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetTakeOverByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetTakeOverByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetTakeOverByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetTakeOverByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) GetTakeOverByUserIdAsync(
	request *GetTakeOverByUserIdRequest,
	callback chan<- GetTakeOverByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "takeOver",
    		"function": "getTakeOverByUserId",
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
    if request.Type != nil {
        bodies["type"] = *request.Type
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getTakeOverByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) GetTakeOverByUserId(
	request *GetTakeOverByUserIdRequest,
) (*GetTakeOverByUserIdResult, error) {
	callback := make(chan GetTakeOverByUserIdAsyncResult, 1)
	go p.GetTakeOverByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) updateTakeOverAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateTakeOverAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateTakeOverAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateTakeOverResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateTakeOverAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
        gs2err, ok := asyncResult.Err.(core.Gs2Exception)
        if ok {
            if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "account.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
            }
        }
    }
	callback <- UpdateTakeOverAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) UpdateTakeOverAsync(
	request *UpdateTakeOverRequest,
	callback chan<- UpdateTakeOverAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "takeOver",
    		"function": "updateTakeOver",
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
    if request.Type != nil {
        bodies["type"] = *request.Type
    }
    if request.OldPassword != nil && *request.OldPassword != "" {
        bodies["oldPassword"] = *request.OldPassword
    }
    if request.Password != nil && *request.Password != "" {
        bodies["password"] = *request.Password
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

	go p.updateTakeOverAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) UpdateTakeOver(
	request *UpdateTakeOverRequest,
) (*UpdateTakeOverResult, error) {
	callback := make(chan UpdateTakeOverAsyncResult, 1)
	go p.UpdateTakeOverAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) updateTakeOverByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateTakeOverByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateTakeOverByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateTakeOverByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateTakeOverByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
        gs2err, ok := asyncResult.Err.(core.Gs2Exception)
        if ok {
            if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "account.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
            }
        }
    }
	callback <- UpdateTakeOverByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) UpdateTakeOverByUserIdAsync(
	request *UpdateTakeOverByUserIdRequest,
	callback chan<- UpdateTakeOverByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "takeOver",
    		"function": "updateTakeOverByUserId",
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
    if request.Type != nil {
        bodies["type"] = *request.Type
    }
    if request.OldPassword != nil && *request.OldPassword != "" {
        bodies["oldPassword"] = *request.OldPassword
    }
    if request.Password != nil && *request.Password != "" {
        bodies["password"] = *request.Password
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.updateTakeOverByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) UpdateTakeOverByUserId(
	request *UpdateTakeOverByUserIdRequest,
) (*UpdateTakeOverByUserIdResult, error) {
	callback := make(chan UpdateTakeOverByUserIdAsyncResult, 1)
	go p.UpdateTakeOverByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) deleteTakeOverAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteTakeOverAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteTakeOverAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteTakeOverResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteTakeOverAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteTakeOverAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) DeleteTakeOverAsync(
	request *DeleteTakeOverRequest,
	callback chan<- DeleteTakeOverAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "takeOver",
    		"function": "deleteTakeOver",
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
    if request.Type != nil {
        bodies["type"] = *request.Type
    }
    if request.UserIdentifier != nil && *request.UserIdentifier != "" {
        bodies["userIdentifier"] = *request.UserIdentifier
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

	go p.deleteTakeOverAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) DeleteTakeOver(
	request *DeleteTakeOverRequest,
) (*DeleteTakeOverResult, error) {
	callback := make(chan DeleteTakeOverAsyncResult, 1)
	go p.DeleteTakeOverAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) deleteTakeOverByUserIdentifierAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteTakeOverByUserIdentifierAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteTakeOverByUserIdentifierAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteTakeOverByUserIdentifierResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteTakeOverByUserIdentifierAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteTakeOverByUserIdentifierAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) DeleteTakeOverByUserIdentifierAsync(
	request *DeleteTakeOverByUserIdentifierRequest,
	callback chan<- DeleteTakeOverByUserIdentifierAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "takeOver",
    		"function": "deleteTakeOverByUserIdentifier",
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
    if request.Type != nil {
        bodies["type"] = *request.Type
    }
    if request.UserIdentifier != nil && *request.UserIdentifier != "" {
        bodies["userIdentifier"] = *request.UserIdentifier
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.deleteTakeOverByUserIdentifierAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) DeleteTakeOverByUserIdentifier(
	request *DeleteTakeOverByUserIdentifierRequest,
) (*DeleteTakeOverByUserIdentifierResult, error) {
	callback := make(chan DeleteTakeOverByUserIdentifierAsyncResult, 1)
	go p.DeleteTakeOverByUserIdentifierAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) doTakeOverAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DoTakeOverAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoTakeOverAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoTakeOverResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DoTakeOverAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
        gs2err, ok := asyncResult.Err.(core.Gs2Exception)
        if ok {
            if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "account.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
            }
        }
    }
	callback <- DoTakeOverAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) DoTakeOverAsync(
	request *DoTakeOverRequest,
	callback chan<- DoTakeOverAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "takeOver",
    		"function": "doTakeOver",
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
    if request.Type != nil {
        bodies["type"] = *request.Type
    }
    if request.UserIdentifier != nil && *request.UserIdentifier != "" {
        bodies["userIdentifier"] = *request.UserIdentifier
    }
    if request.Password != nil && *request.Password != "" {
        bodies["password"] = *request.Password
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.doTakeOverAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) DoTakeOver(
	request *DoTakeOverRequest,
) (*DoTakeOverResult, error) {
	callback := make(chan DoTakeOverAsyncResult, 1)
	go p.DoTakeOverAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) getDataOwnerByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetDataOwnerByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDataOwnerByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDataOwnerByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetDataOwnerByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- GetDataOwnerByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) GetDataOwnerByUserIdAsync(
	request *GetDataOwnerByUserIdRequest,
	callback chan<- GetDataOwnerByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "dataOwner",
    		"function": "getDataOwnerByUserId",
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

	go p.getDataOwnerByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) GetDataOwnerByUserId(
	request *GetDataOwnerByUserIdRequest,
) (*GetDataOwnerByUserIdResult, error) {
	callback := make(chan GetDataOwnerByUserIdAsyncResult, 1)
	go p.GetDataOwnerByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AccountWebSocketClient) deleteDataOwnerByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteDataOwnerByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteDataOwnerByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteDataOwnerByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteDataOwnerByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
    if asyncResult.Err != nil {
    }
	callback <- DeleteDataOwnerByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountWebSocketClient) DeleteDataOwnerByUserIdAsync(
	request *DeleteDataOwnerByUserIdRequest,
	callback chan<- DeleteDataOwnerByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "account",
    		"component": "dataOwner",
    		"function": "deleteDataOwnerByUserId",
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
    if request.DuplicationAvoider != nil {
      bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
    }

	go p.deleteDataOwnerByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2AccountWebSocketClient) DeleteDataOwnerByUserId(
	request *DeleteDataOwnerByUserIdRequest,
) (*DeleteDataOwnerByUserIdResult, error) {
	callback := make(chan DeleteDataOwnerByUserIdAsyncResult, 1)
	go p.DeleteDataOwnerByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
