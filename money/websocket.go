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

package money

import (
	"encoding/json"
	"github.com/google/uuid"
	"core"
)

type Gs2MoneyWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2MoneyWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2MoneyWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2MoneyWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
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

func (p Gs2MoneyWebSocketClient) DescribeNamespaces(
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

func (p Gs2MoneyWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2MoneyWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
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
    if request.Priority != nil && *request.Priority != "" {
        bodies["priority"] = *request.Priority
    }
    if request.ShareFree != nil {
        bodies["shareFree"] = *request.ShareFree
    }
    if request.Currency != nil && *request.Currency != "" {
        bodies["currency"] = *request.Currency
    }
    if request.AppleKey != nil && *request.AppleKey != "" {
        bodies["appleKey"] = *request.AppleKey
    }
    if request.GoogleKey != nil && *request.GoogleKey != "" {
        bodies["googleKey"] = *request.GoogleKey
    }
    if request.EnableFakeReceipt != nil {
        bodies["enableFakeReceipt"] = *request.EnableFakeReceipt
    }
    if request.CreateWalletScript != nil {
        bodies["createWalletScript"] = request.CreateWalletScript.ToDict()
    }
    if request.DepositScript != nil {
        bodies["depositScript"] = request.DepositScript.ToDict()
    }
    if request.WithdrawScript != nil {
        bodies["withdrawScript"] = request.WithdrawScript.ToDict()
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

func (p Gs2MoneyWebSocketClient) CreateNamespace(
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

func (p Gs2MoneyWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2MoneyWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
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

func (p Gs2MoneyWebSocketClient) GetNamespaceStatus(
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

func (p Gs2MoneyWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2MoneyWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
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

func (p Gs2MoneyWebSocketClient) GetNamespace(
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

func (p Gs2MoneyWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2MoneyWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
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
    if request.Priority != nil && *request.Priority != "" {
        bodies["priority"] = *request.Priority
    }
    if request.AppleKey != nil && *request.AppleKey != "" {
        bodies["appleKey"] = *request.AppleKey
    }
    if request.GoogleKey != nil && *request.GoogleKey != "" {
        bodies["googleKey"] = *request.GoogleKey
    }
    if request.EnableFakeReceipt != nil {
        bodies["enableFakeReceipt"] = *request.EnableFakeReceipt
    }
    if request.CreateWalletScript != nil {
        bodies["createWalletScript"] = request.CreateWalletScript.ToDict()
    }
    if request.DepositScript != nil {
        bodies["depositScript"] = request.DepositScript.ToDict()
    }
    if request.WithdrawScript != nil {
        bodies["withdrawScript"] = request.WithdrawScript.ToDict()
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

func (p Gs2MoneyWebSocketClient) UpdateNamespace(
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

func (p Gs2MoneyWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2MoneyWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
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

func (p Gs2MoneyWebSocketClient) DeleteNamespace(
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

func (p Gs2MoneyWebSocketClient) describeWalletsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeWalletsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeWalletsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeWalletsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeWalletsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeWalletsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) DescribeWalletsAsync(
	request *DescribeWalletsRequest,
	callback chan<- DescribeWalletsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "wallet",
    		"function": "describeWallets",
            "contentType": "application/json",
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

	go p.describeWalletsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) DescribeWallets(
	request *DescribeWalletsRequest,
) (*DescribeWalletsResult, error) {
	callback := make(chan DescribeWalletsAsyncResult, 1)
	go p.DescribeWalletsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MoneyWebSocketClient) describeWalletsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeWalletsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeWalletsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeWalletsByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeWalletsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeWalletsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) DescribeWalletsByUserIdAsync(
	request *DescribeWalletsByUserIdRequest,
	callback chan<- DescribeWalletsByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "wallet",
    		"function": "describeWalletsByUserId",
            "contentType": "application/json",
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

	go p.describeWalletsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) DescribeWalletsByUserId(
	request *DescribeWalletsByUserIdRequest,
) (*DescribeWalletsByUserIdResult, error) {
	callback := make(chan DescribeWalletsByUserIdAsyncResult, 1)
	go p.DescribeWalletsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MoneyWebSocketClient) getWalletAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetWalletAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetWalletAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetWalletResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetWalletAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetWalletAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) GetWalletAsync(
	request *GetWalletRequest,
	callback chan<- GetWalletAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "wallet",
    		"function": "getWallet",
            "contentType": "application/json",
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
    if request.Slot != nil {
        bodies["slot"] = *request.Slot
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getWalletAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) GetWallet(
	request *GetWalletRequest,
) (*GetWalletResult, error) {
	callback := make(chan GetWalletAsyncResult, 1)
	go p.GetWalletAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MoneyWebSocketClient) getWalletByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetWalletByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetWalletByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetWalletByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetWalletByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetWalletByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) GetWalletByUserIdAsync(
	request *GetWalletByUserIdRequest,
	callback chan<- GetWalletByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "wallet",
    		"function": "getWalletByUserId",
            "contentType": "application/json",
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
    if request.Slot != nil {
        bodies["slot"] = *request.Slot
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getWalletByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) GetWalletByUserId(
	request *GetWalletByUserIdRequest,
) (*GetWalletByUserIdResult, error) {
	callback := make(chan GetWalletByUserIdAsyncResult, 1)
	go p.GetWalletByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MoneyWebSocketClient) depositByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DepositByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DepositByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DepositByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DepositByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DepositByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) DepositByUserIdAsync(
	request *DepositByUserIdRequest,
	callback chan<- DepositByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "wallet",
    		"function": "depositByUserId",
            "contentType": "application/json",
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
    if request.Slot != nil {
        bodies["slot"] = *request.Slot
    }
    if request.Price != nil {
        bodies["price"] = *request.Price
    }
    if request.Count != nil {
        bodies["count"] = *request.Count
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.depositByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) DepositByUserId(
	request *DepositByUserIdRequest,
) (*DepositByUserIdResult, error) {
	callback := make(chan DepositByUserIdAsyncResult, 1)
	go p.DepositByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MoneyWebSocketClient) withdrawAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- WithdrawAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WithdrawAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WithdrawResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- WithdrawAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- WithdrawAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) WithdrawAsync(
	request *WithdrawRequest,
	callback chan<- WithdrawAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "wallet",
    		"function": "withdraw",
            "contentType": "application/json",
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
    if request.Slot != nil {
        bodies["slot"] = *request.Slot
    }
    if request.Count != nil {
        bodies["count"] = *request.Count
    }
    if request.PaidOnly != nil {
        bodies["paidOnly"] = *request.PaidOnly
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.withdrawAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) Withdraw(
	request *WithdrawRequest,
) (*WithdrawResult, error) {
	callback := make(chan WithdrawAsyncResult, 1)
	go p.WithdrawAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MoneyWebSocketClient) withdrawByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- WithdrawByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WithdrawByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WithdrawByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- WithdrawByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- WithdrawByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) WithdrawByUserIdAsync(
	request *WithdrawByUserIdRequest,
	callback chan<- WithdrawByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "wallet",
    		"function": "withdrawByUserId",
            "contentType": "application/json",
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
    if request.Slot != nil {
        bodies["slot"] = *request.Slot
    }
    if request.Count != nil {
        bodies["count"] = *request.Count
    }
    if request.PaidOnly != nil {
        bodies["paidOnly"] = *request.PaidOnly
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.withdrawByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) WithdrawByUserId(
	request *WithdrawByUserIdRequest,
) (*WithdrawByUserIdResult, error) {
	callback := make(chan WithdrawByUserIdAsyncResult, 1)
	go p.WithdrawByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MoneyWebSocketClient) depositByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DepositByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DepositByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DepositByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DepositByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DepositByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) DepositByStampSheetAsync(
	request *DepositByStampSheetRequest,
	callback chan<- DepositByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "wallet",
    		"function": "depositByStampSheet",
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

	go p.depositByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) DepositByStampSheet(
	request *DepositByStampSheetRequest,
) (*DepositByStampSheetResult, error) {
	callback := make(chan DepositByStampSheetAsyncResult, 1)
	go p.DepositByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MoneyWebSocketClient) withdrawByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- WithdrawByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WithdrawByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WithdrawByStampTaskResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- WithdrawByStampTaskAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- WithdrawByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) WithdrawByStampTaskAsync(
	request *WithdrawByStampTaskRequest,
	callback chan<- WithdrawByStampTaskAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "wallet",
    		"function": "withdrawByStampTask",
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

	go p.withdrawByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) WithdrawByStampTask(
	request *WithdrawByStampTaskRequest,
) (*WithdrawByStampTaskResult, error) {
	callback := make(chan WithdrawByStampTaskAsyncResult, 1)
	go p.WithdrawByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MoneyWebSocketClient) describeReceiptsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeReceiptsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReceiptsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReceiptsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeReceiptsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeReceiptsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) DescribeReceiptsAsync(
	request *DescribeReceiptsRequest,
	callback chan<- DescribeReceiptsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "receipt",
    		"function": "describeReceipts",
            "contentType": "application/json",
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
    if request.Slot != nil {
        bodies["slot"] = *request.Slot
    }
    if request.Begin != nil {
        bodies["begin"] = *request.Begin
    }
    if request.End != nil {
        bodies["end"] = *request.End
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

	go p.describeReceiptsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) DescribeReceipts(
	request *DescribeReceiptsRequest,
) (*DescribeReceiptsResult, error) {
	callback := make(chan DescribeReceiptsAsyncResult, 1)
	go p.DescribeReceiptsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MoneyWebSocketClient) getByUserIdAndTransactionIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetByUserIdAndTransactionIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetByUserIdAndTransactionIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetByUserIdAndTransactionIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetByUserIdAndTransactionIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetByUserIdAndTransactionIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) GetByUserIdAndTransactionIdAsync(
	request *GetByUserIdAndTransactionIdRequest,
	callback chan<- GetByUserIdAndTransactionIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "receipt",
    		"function": "getByUserIdAndTransactionId",
            "contentType": "application/json",
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
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getByUserIdAndTransactionIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) GetByUserIdAndTransactionId(
	request *GetByUserIdAndTransactionIdRequest,
) (*GetByUserIdAndTransactionIdResult, error) {
	callback := make(chan GetByUserIdAndTransactionIdAsyncResult, 1)
	go p.GetByUserIdAndTransactionIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MoneyWebSocketClient) recordReceiptAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RecordReceiptAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RecordReceiptAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RecordReceiptResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- RecordReceiptAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- RecordReceiptAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) RecordReceiptAsync(
	request *RecordReceiptRequest,
	callback chan<- RecordReceiptAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "receipt",
    		"function": "recordReceipt",
            "contentType": "application/json",
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
    if request.ContentsId != nil && *request.ContentsId != "" {
        bodies["contentsId"] = *request.ContentsId
    }
    if request.Receipt != nil && *request.Receipt != "" {
        bodies["receipt"] = *request.Receipt
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.recordReceiptAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) RecordReceipt(
	request *RecordReceiptRequest,
) (*RecordReceiptResult, error) {
	callback := make(chan RecordReceiptAsyncResult, 1)
	go p.RecordReceiptAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MoneyWebSocketClient) recordReceiptByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RecordReceiptByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RecordReceiptByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RecordReceiptByStampTaskResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- RecordReceiptByStampTaskAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- RecordReceiptByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MoneyWebSocketClient) RecordReceiptByStampTaskAsync(
	request *RecordReceiptByStampTaskRequest,
	callback chan<- RecordReceiptByStampTaskAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "money",
    		"component": "receipt",
    		"function": "recordReceiptByStampTask",
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

	go p.recordReceiptByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MoneyWebSocketClient) RecordReceiptByStampTask(
	request *RecordReceiptByStampTaskRequest,
) (*RecordReceiptByStampTaskResult, error) {
	callback := make(chan RecordReceiptByStampTaskAsyncResult, 1)
	go p.RecordReceiptByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
