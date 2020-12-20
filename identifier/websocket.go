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

package identifier

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2IdentifierWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2IdentifierWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2IdentifierWebSocketClient) describeUsersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeUsersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeUsersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeUsersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeUsersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeUsersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) DescribeUsersAsync(
	request *DescribeUsersRequest,
	callback chan<- DescribeUsersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "user",
    		"function": "describeUsers",
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

	go p.describeUsersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) DescribeUsers(
	request *DescribeUsersRequest,
) (*DescribeUsersResult, error) {
	callback := make(chan DescribeUsersAsyncResult, 1)
	go p.DescribeUsersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) createUserAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateUserResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateUserAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) CreateUserAsync(
	request *CreateUserRequest,
	callback chan<- CreateUserAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "user",
    		"function": "createUser",
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
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createUserAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) CreateUser(
	request *CreateUserRequest,
) (*CreateUserResult, error) {
	callback := make(chan CreateUserAsyncResult, 1)
	go p.CreateUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) updateUserAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateUserResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateUserAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) UpdateUserAsync(
	request *UpdateUserRequest,
	callback chan<- UpdateUserAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "user",
    		"function": "updateUser",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateUserAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) UpdateUser(
	request *UpdateUserRequest,
) (*UpdateUserResult, error) {
	callback := make(chan UpdateUserAsyncResult, 1)
	go p.UpdateUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) getUserAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetUserResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetUserAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) GetUserAsync(
	request *GetUserRequest,
	callback chan<- GetUserAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "user",
    		"function": "getUser",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getUserAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) GetUser(
	request *GetUserRequest,
) (*GetUserResult, error) {
	callback := make(chan GetUserAsyncResult, 1)
	go p.GetUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) deleteUserAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteUserResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteUserAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) DeleteUserAsync(
	request *DeleteUserRequest,
	callback chan<- DeleteUserAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "user",
    		"function": "deleteUser",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteUserAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) DeleteUser(
	request *DeleteUserRequest,
) (*DeleteUserResult, error) {
	callback := make(chan DeleteUserAsyncResult, 1)
	go p.DeleteUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) describeSecurityPoliciesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSecurityPoliciesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSecurityPoliciesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSecurityPoliciesResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeSecurityPoliciesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeSecurityPoliciesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) DescribeSecurityPoliciesAsync(
	request *DescribeSecurityPoliciesRequest,
	callback chan<- DescribeSecurityPoliciesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "securityPolicy",
    		"function": "describeSecurityPolicies",
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

	go p.describeSecurityPoliciesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) DescribeSecurityPolicies(
	request *DescribeSecurityPoliciesRequest,
) (*DescribeSecurityPoliciesResult, error) {
	callback := make(chan DescribeSecurityPoliciesAsyncResult, 1)
	go p.DescribeSecurityPoliciesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) describeCommonSecurityPoliciesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeCommonSecurityPoliciesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCommonSecurityPoliciesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCommonSecurityPoliciesResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeCommonSecurityPoliciesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeCommonSecurityPoliciesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) DescribeCommonSecurityPoliciesAsync(
	request *DescribeCommonSecurityPoliciesRequest,
	callback chan<- DescribeCommonSecurityPoliciesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "securityPolicy",
    		"function": "describeCommonSecurityPolicies",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.OwnerId != nil && *request.OwnerId != "" {
        bodies["ownerId"] = *request.OwnerId
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

	go p.describeCommonSecurityPoliciesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) DescribeCommonSecurityPolicies(
	request *DescribeCommonSecurityPoliciesRequest,
) (*DescribeCommonSecurityPoliciesResult, error) {
	callback := make(chan DescribeCommonSecurityPoliciesAsyncResult, 1)
	go p.DescribeCommonSecurityPoliciesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) createSecurityPolicyAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateSecurityPolicyResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateSecurityPolicyAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) CreateSecurityPolicyAsync(
	request *CreateSecurityPolicyRequest,
	callback chan<- CreateSecurityPolicyAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "securityPolicy",
    		"function": "createSecurityPolicy",
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
    if request.Policy != nil && *request.Policy != "" {
        bodies["policy"] = *request.Policy
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createSecurityPolicyAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) CreateSecurityPolicy(
	request *CreateSecurityPolicyRequest,
) (*CreateSecurityPolicyResult, error) {
	callback := make(chan CreateSecurityPolicyAsyncResult, 1)
	go p.CreateSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) updateSecurityPolicyAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateSecurityPolicyResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateSecurityPolicyAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) UpdateSecurityPolicyAsync(
	request *UpdateSecurityPolicyRequest,
	callback chan<- UpdateSecurityPolicyAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "securityPolicy",
    		"function": "updateSecurityPolicy",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.SecurityPolicyName != nil && *request.SecurityPolicyName != "" {
        bodies["securityPolicyName"] = *request.SecurityPolicyName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Policy != nil && *request.Policy != "" {
        bodies["policy"] = *request.Policy
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateSecurityPolicyAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) UpdateSecurityPolicy(
	request *UpdateSecurityPolicyRequest,
) (*UpdateSecurityPolicyResult, error) {
	callback := make(chan UpdateSecurityPolicyAsyncResult, 1)
	go p.UpdateSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) getSecurityPolicyAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSecurityPolicyResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetSecurityPolicyAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) GetSecurityPolicyAsync(
	request *GetSecurityPolicyRequest,
	callback chan<- GetSecurityPolicyAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "securityPolicy",
    		"function": "getSecurityPolicy",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.SecurityPolicyName != nil && *request.SecurityPolicyName != "" {
        bodies["securityPolicyName"] = *request.SecurityPolicyName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getSecurityPolicyAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) GetSecurityPolicy(
	request *GetSecurityPolicyRequest,
) (*GetSecurityPolicyResult, error) {
	callback := make(chan GetSecurityPolicyAsyncResult, 1)
	go p.GetSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) deleteSecurityPolicyAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSecurityPolicyResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteSecurityPolicyAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) DeleteSecurityPolicyAsync(
	request *DeleteSecurityPolicyRequest,
	callback chan<- DeleteSecurityPolicyAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "securityPolicy",
    		"function": "deleteSecurityPolicy",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.SecurityPolicyName != nil && *request.SecurityPolicyName != "" {
        bodies["securityPolicyName"] = *request.SecurityPolicyName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteSecurityPolicyAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) DeleteSecurityPolicy(
	request *DeleteSecurityPolicyRequest,
) (*DeleteSecurityPolicyResult, error) {
	callback := make(chan DeleteSecurityPolicyAsyncResult, 1)
	go p.DeleteSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) describeIdentifiersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeIdentifiersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeIdentifiersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeIdentifiersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeIdentifiersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeIdentifiersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) DescribeIdentifiersAsync(
	request *DescribeIdentifiersRequest,
	callback chan<- DescribeIdentifiersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "identifier",
    		"function": "describeIdentifiers",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
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

	go p.describeIdentifiersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) DescribeIdentifiers(
	request *DescribeIdentifiersRequest,
) (*DescribeIdentifiersResult, error) {
	callback := make(chan DescribeIdentifiersAsyncResult, 1)
	go p.DescribeIdentifiersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) createIdentifierAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateIdentifierAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateIdentifierAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateIdentifierResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateIdentifierAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateIdentifierAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) CreateIdentifierAsync(
	request *CreateIdentifierRequest,
	callback chan<- CreateIdentifierAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "identifier",
    		"function": "createIdentifier",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createIdentifierAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) CreateIdentifier(
	request *CreateIdentifierRequest,
) (*CreateIdentifierResult, error) {
	callback := make(chan CreateIdentifierAsyncResult, 1)
	go p.CreateIdentifierAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) getIdentifierAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetIdentifierAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetIdentifierAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetIdentifierResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetIdentifierAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetIdentifierAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) GetIdentifierAsync(
	request *GetIdentifierRequest,
	callback chan<- GetIdentifierAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "identifier",
    		"function": "getIdentifier",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
    if request.ClientId != nil && *request.ClientId != "" {
        bodies["clientId"] = *request.ClientId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getIdentifierAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) GetIdentifier(
	request *GetIdentifierRequest,
) (*GetIdentifierResult, error) {
	callback := make(chan GetIdentifierAsyncResult, 1)
	go p.GetIdentifierAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) deleteIdentifierAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteIdentifierAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteIdentifierAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteIdentifierResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteIdentifierAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteIdentifierAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) DeleteIdentifierAsync(
	request *DeleteIdentifierRequest,
	callback chan<- DeleteIdentifierAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "identifier",
    		"function": "deleteIdentifier",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
    if request.ClientId != nil && *request.ClientId != "" {
        bodies["clientId"] = *request.ClientId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteIdentifierAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) DeleteIdentifier(
	request *DeleteIdentifierRequest,
) (*DeleteIdentifierResult, error) {
	callback := make(chan DeleteIdentifierAsyncResult, 1)
	go p.DeleteIdentifierAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) describePasswordsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribePasswordsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePasswordsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePasswordsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribePasswordsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribePasswordsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) DescribePasswordsAsync(
	request *DescribePasswordsRequest,
	callback chan<- DescribePasswordsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "password",
    		"function": "describePasswords",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
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

	go p.describePasswordsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) DescribePasswords(
	request *DescribePasswordsRequest,
) (*DescribePasswordsResult, error) {
	callback := make(chan DescribePasswordsAsyncResult, 1)
	go p.DescribePasswordsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) createPasswordAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreatePasswordAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreatePasswordAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreatePasswordResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreatePasswordAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreatePasswordAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) CreatePasswordAsync(
	request *CreatePasswordRequest,
	callback chan<- CreatePasswordAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "password",
    		"function": "createPassword",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
    if request.Password != nil && *request.Password != "" {
        bodies["password"] = *request.Password
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createPasswordAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) CreatePassword(
	request *CreatePasswordRequest,
) (*CreatePasswordResult, error) {
	callback := make(chan CreatePasswordAsyncResult, 1)
	go p.CreatePasswordAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) getPasswordAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetPasswordAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPasswordAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPasswordResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetPasswordAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetPasswordAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) GetPasswordAsync(
	request *GetPasswordRequest,
	callback chan<- GetPasswordAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "password",
    		"function": "getPassword",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getPasswordAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) GetPassword(
	request *GetPasswordRequest,
) (*GetPasswordResult, error) {
	callback := make(chan GetPasswordAsyncResult, 1)
	go p.GetPasswordAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) deletePasswordAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeletePasswordAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeletePasswordAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeletePasswordResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeletePasswordAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeletePasswordAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) DeletePasswordAsync(
	request *DeletePasswordRequest,
	callback chan<- DeletePasswordAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "password",
    		"function": "deletePassword",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deletePasswordAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) DeletePassword(
	request *DeletePasswordRequest,
) (*DeletePasswordResult, error) {
	callback := make(chan DeletePasswordAsyncResult, 1)
	go p.DeletePasswordAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) getHasSecurityPolicyAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetHasSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetHasSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetHasSecurityPolicyResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetHasSecurityPolicyAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetHasSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) GetHasSecurityPolicyAsync(
	request *GetHasSecurityPolicyRequest,
	callback chan<- GetHasSecurityPolicyAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "attachSecurityPolicy",
    		"function": "getHasSecurityPolicy",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getHasSecurityPolicyAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) GetHasSecurityPolicy(
	request *GetHasSecurityPolicyRequest,
) (*GetHasSecurityPolicyResult, error) {
	callback := make(chan GetHasSecurityPolicyAsyncResult, 1)
	go p.GetHasSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) attachSecurityPolicyAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AttachSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AttachSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AttachSecurityPolicyResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AttachSecurityPolicyAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AttachSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) AttachSecurityPolicyAsync(
	request *AttachSecurityPolicyRequest,
	callback chan<- AttachSecurityPolicyAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "attachSecurityPolicy",
    		"function": "attachSecurityPolicy",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
    if request.SecurityPolicyId != nil && *request.SecurityPolicyId != "" {
        bodies["securityPolicyId"] = *request.SecurityPolicyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.attachSecurityPolicyAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) AttachSecurityPolicy(
	request *AttachSecurityPolicyRequest,
) (*AttachSecurityPolicyResult, error) {
	callback := make(chan AttachSecurityPolicyAsyncResult, 1)
	go p.AttachSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) detachSecurityPolicyAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DetachSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DetachSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DetachSecurityPolicyResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DetachSecurityPolicyAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DetachSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) DetachSecurityPolicyAsync(
	request *DetachSecurityPolicyRequest,
	callback chan<- DetachSecurityPolicyAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "attachSecurityPolicy",
    		"function": "detachSecurityPolicy",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
    if request.SecurityPolicyId != nil && *request.SecurityPolicyId != "" {
        bodies["securityPolicyId"] = *request.SecurityPolicyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.detachSecurityPolicyAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) DetachSecurityPolicy(
	request *DetachSecurityPolicyRequest,
) (*DetachSecurityPolicyResult, error) {
	callback := make(chan DetachSecurityPolicyAsyncResult, 1)
	go p.DetachSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) loginAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- LoginAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- LoginAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result LoginResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- LoginAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- LoginAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) LoginAsync(
	request *LoginRequest,
	callback chan<- LoginAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "projectToken",
    		"function": "login",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.ClientId != nil && *request.ClientId != "" {
        bodies["client_id"] = *request.ClientId
    }
    if request.ClientSecret != nil && *request.ClientSecret != "" {
        bodies["client_secret"] = *request.ClientSecret
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.loginAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) Login(
	request *LoginRequest,
) (*LoginResult, error) {
	callback := make(chan LoginAsyncResult, 1)
	go p.LoginAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2IdentifierWebSocketClient) loginByUserAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- LoginByUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- LoginByUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result LoginByUserResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- LoginByUserAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- LoginByUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierWebSocketClient) LoginByUserAsync(
	request *LoginByUserRequest,
	callback chan<- LoginByUserAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "identifier",
    		"component": "projectToken",
    		"function": "loginByUser",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
    if request.Password != nil && *request.Password != "" {
        bodies["password"] = *request.Password
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.loginByUserAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierWebSocketClient) LoginByUser(
	request *LoginByUserRequest,
) (*LoginByUserResult, error) {
	callback := make(chan LoginByUserAsyncResult, 1)
	go p.LoginByUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
