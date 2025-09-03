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

package friend

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2FriendWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2FriendWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2FriendWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2FriendWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "describeNamespaces",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeNamespacesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DescribeNamespaces(
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

func (p Gs2FriendWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2FriendWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "createNamespace",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.FollowScript != nil {
		bodies["followScript"] = request.FollowScript.ToDict()
	}
	if request.UnfollowScript != nil {
		bodies["unfollowScript"] = request.UnfollowScript.ToDict()
	}
	if request.SendRequestScript != nil {
		bodies["sendRequestScript"] = request.SendRequestScript.ToDict()
	}
	if request.CancelRequestScript != nil {
		bodies["cancelRequestScript"] = request.CancelRequestScript.ToDict()
	}
	if request.AcceptRequestScript != nil {
		bodies["acceptRequestScript"] = request.AcceptRequestScript.ToDict()
	}
	if request.RejectRequestScript != nil {
		bodies["rejectRequestScript"] = request.RejectRequestScript.ToDict()
	}
	if request.DeleteFriendScript != nil {
		bodies["deleteFriendScript"] = request.DeleteFriendScript.ToDict()
	}
	if request.UpdateProfileScript != nil {
		bodies["updateProfileScript"] = request.UpdateProfileScript.ToDict()
	}
	if request.FollowNotification != nil {
		bodies["followNotification"] = request.FollowNotification.ToDict()
	}
	if request.ReceiveRequestNotification != nil {
		bodies["receiveRequestNotification"] = request.ReceiveRequestNotification.ToDict()
	}
	if request.CancelRequestNotification != nil {
		bodies["cancelRequestNotification"] = request.CancelRequestNotification.ToDict()
	}
	if request.AcceptRequestNotification != nil {
		bodies["acceptRequestNotification"] = request.AcceptRequestNotification.ToDict()
	}
	if request.RejectRequestNotification != nil {
		bodies["rejectRequestNotification"] = request.RejectRequestNotification.ToDict()
	}
	if request.DeleteFriendNotification != nil {
		bodies["deleteFriendNotification"] = request.DeleteFriendNotification.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.createNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) CreateNamespace(
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

func (p Gs2FriendWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2FriendWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "getNamespaceStatus",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getNamespaceStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetNamespaceStatus(
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

func (p Gs2FriendWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2FriendWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "getNamespace",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetNamespace(
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

func (p Gs2FriendWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2FriendWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "updateNamespace",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.FollowScript != nil {
		bodies["followScript"] = request.FollowScript.ToDict()
	}
	if request.UnfollowScript != nil {
		bodies["unfollowScript"] = request.UnfollowScript.ToDict()
	}
	if request.SendRequestScript != nil {
		bodies["sendRequestScript"] = request.SendRequestScript.ToDict()
	}
	if request.CancelRequestScript != nil {
		bodies["cancelRequestScript"] = request.CancelRequestScript.ToDict()
	}
	if request.AcceptRequestScript != nil {
		bodies["acceptRequestScript"] = request.AcceptRequestScript.ToDict()
	}
	if request.RejectRequestScript != nil {
		bodies["rejectRequestScript"] = request.RejectRequestScript.ToDict()
	}
	if request.DeleteFriendScript != nil {
		bodies["deleteFriendScript"] = request.DeleteFriendScript.ToDict()
	}
	if request.UpdateProfileScript != nil {
		bodies["updateProfileScript"] = request.UpdateProfileScript.ToDict()
	}
	if request.FollowNotification != nil {
		bodies["followNotification"] = request.FollowNotification.ToDict()
	}
	if request.ReceiveRequestNotification != nil {
		bodies["receiveRequestNotification"] = request.ReceiveRequestNotification.ToDict()
	}
	if request.CancelRequestNotification != nil {
		bodies["cancelRequestNotification"] = request.CancelRequestNotification.ToDict()
	}
	if request.AcceptRequestNotification != nil {
		bodies["acceptRequestNotification"] = request.AcceptRequestNotification.ToDict()
	}
	if request.RejectRequestNotification != nil {
		bodies["rejectRequestNotification"] = request.RejectRequestNotification.ToDict()
	}
	if request.DeleteFriendNotification != nil {
		bodies["deleteFriendNotification"] = request.DeleteFriendNotification.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.updateNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) UpdateNamespace(
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

func (p Gs2FriendWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2FriendWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "deleteNamespace",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.deleteNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DeleteNamespace(
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

func (p Gs2FriendWebSocketClient) getServiceVersionAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetServiceVersionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetServiceVersionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetServiceVersionResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetServiceVersionAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetServiceVersionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) GetServiceVersionAsync(
	request *GetServiceVersionRequest,
	callback chan<- GetServiceVersionAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "getServiceVersion",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getServiceVersionAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetServiceVersion(
	request *GetServiceVersionRequest,
) (*GetServiceVersionResult, error) {
	callback := make(chan GetServiceVersionAsyncResult, 1)
	go p.GetServiceVersionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) dumpUserDataByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DumpUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DumpUserDataByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DumpUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DumpUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DumpUserDataByUserIdAsync(
	request *DumpUserDataByUserIdRequest,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "dumpUserDataByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.dumpUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DumpUserDataByUserId(
	request *DumpUserDataByUserIdRequest,
) (*DumpUserDataByUserIdResult, error) {
	callback := make(chan DumpUserDataByUserIdAsyncResult, 1)
	go p.DumpUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) checkDumpUserDataByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CheckDumpUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CheckDumpUserDataByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckDumpUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CheckDumpUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) CheckDumpUserDataByUserIdAsync(
	request *CheckDumpUserDataByUserIdRequest,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "checkDumpUserDataByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.checkDumpUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) CheckDumpUserDataByUserId(
	request *CheckDumpUserDataByUserIdRequest,
) (*CheckDumpUserDataByUserIdResult, error) {
	callback := make(chan CheckDumpUserDataByUserIdAsyncResult, 1)
	go p.CheckDumpUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) cleanUserDataByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CleanUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CleanUserDataByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CleanUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CleanUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) CleanUserDataByUserIdAsync(
	request *CleanUserDataByUserIdRequest,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "cleanUserDataByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.cleanUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) CleanUserDataByUserId(
	request *CleanUserDataByUserIdRequest,
) (*CleanUserDataByUserIdResult, error) {
	callback := make(chan CleanUserDataByUserIdAsyncResult, 1)
	go p.CleanUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) checkCleanUserDataByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CheckCleanUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CheckCleanUserDataByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckCleanUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CheckCleanUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) CheckCleanUserDataByUserIdAsync(
	request *CheckCleanUserDataByUserIdRequest,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "checkCleanUserDataByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.checkCleanUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) CheckCleanUserDataByUserId(
	request *CheckCleanUserDataByUserIdRequest,
) (*CheckCleanUserDataByUserIdResult, error) {
	callback := make(chan CheckCleanUserDataByUserIdAsyncResult, 1)
	go p.CheckCleanUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) prepareImportUserDataByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareImportUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareImportUserDataByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareImportUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PrepareImportUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) PrepareImportUserDataByUserIdAsync(
	request *PrepareImportUserDataByUserIdRequest,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "prepareImportUserDataByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.prepareImportUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) PrepareImportUserDataByUserId(
	request *PrepareImportUserDataByUserIdRequest,
) (*PrepareImportUserDataByUserIdResult, error) {
	callback := make(chan PrepareImportUserDataByUserIdAsyncResult, 1)
	go p.PrepareImportUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) importUserDataByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ImportUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ImportUserDataByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ImportUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ImportUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) ImportUserDataByUserIdAsync(
	request *ImportUserDataByUserIdRequest,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "importUserDataByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.importUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) ImportUserDataByUserId(
	request *ImportUserDataByUserIdRequest,
) (*ImportUserDataByUserIdResult, error) {
	callback := make(chan ImportUserDataByUserIdAsyncResult, 1)
	go p.ImportUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) checkImportUserDataByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CheckImportUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CheckImportUserDataByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckImportUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CheckImportUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) CheckImportUserDataByUserIdAsync(
	request *CheckImportUserDataByUserIdRequest,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "namespace",
			"function":    "checkImportUserDataByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.checkImportUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) CheckImportUserDataByUserId(
	request *CheckImportUserDataByUserIdRequest,
) (*CheckImportUserDataByUserIdResult, error) {
	callback := make(chan CheckImportUserDataByUserIdAsyncResult, 1)
	go p.CheckImportUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) getProfileAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetProfileAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetProfileAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetProfileResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetProfileAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetProfileAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) GetProfileAsync(
	request *GetProfileRequest,
	callback chan<- GetProfileAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "profile",
			"function":    "getProfile",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getProfileAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetProfile(
	request *GetProfileRequest,
) (*GetProfileResult, error) {
	callback := make(chan GetProfileAsyncResult, 1)
	go p.GetProfileAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) getProfileByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetProfileByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetProfileByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetProfileByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetProfileByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetProfileByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) GetProfileByUserIdAsync(
	request *GetProfileByUserIdRequest,
	callback chan<- GetProfileByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "profile",
			"function":    "getProfileByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getProfileByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetProfileByUserId(
	request *GetProfileByUserIdRequest,
) (*GetProfileByUserIdResult, error) {
	callback := make(chan GetProfileByUserIdAsyncResult, 1)
	go p.GetProfileByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) updateProfileAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateProfileAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateProfileAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateProfileResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateProfileAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateProfileAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) UpdateProfileAsync(
	request *UpdateProfileRequest,
	callback chan<- UpdateProfileAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "profile",
			"function":    "updateProfile",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.PublicProfile != nil && *request.PublicProfile != "" {
		bodies["publicProfile"] = *request.PublicProfile
	}
	if request.FollowerProfile != nil && *request.FollowerProfile != "" {
		bodies["followerProfile"] = *request.FollowerProfile
	}
	if request.FriendProfile != nil && *request.FriendProfile != "" {
		bodies["friendProfile"] = *request.FriendProfile
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.updateProfileAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) UpdateProfile(
	request *UpdateProfileRequest,
) (*UpdateProfileResult, error) {
	callback := make(chan UpdateProfileAsyncResult, 1)
	go p.UpdateProfileAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) updateProfileByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateProfileByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateProfileByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateProfileByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateProfileByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateProfileByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) UpdateProfileByUserIdAsync(
	request *UpdateProfileByUserIdRequest,
	callback chan<- UpdateProfileByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "profile",
			"function":    "updateProfileByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.PublicProfile != nil && *request.PublicProfile != "" {
		bodies["publicProfile"] = *request.PublicProfile
	}
	if request.FollowerProfile != nil && *request.FollowerProfile != "" {
		bodies["followerProfile"] = *request.FollowerProfile
	}
	if request.FriendProfile != nil && *request.FriendProfile != "" {
		bodies["friendProfile"] = *request.FriendProfile
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.updateProfileByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) UpdateProfileByUserId(
	request *UpdateProfileByUserIdRequest,
) (*UpdateProfileByUserIdResult, error) {
	callback := make(chan UpdateProfileByUserIdAsyncResult, 1)
	go p.UpdateProfileByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) deleteProfileByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteProfileByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteProfileByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteProfileByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteProfileByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteProfileByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DeleteProfileByUserIdAsync(
	request *DeleteProfileByUserIdRequest,
	callback chan<- DeleteProfileByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "profile",
			"function":    "deleteProfileByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.deleteProfileByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DeleteProfileByUserId(
	request *DeleteProfileByUserIdRequest,
) (*DeleteProfileByUserIdResult, error) {
	callback := make(chan DeleteProfileByUserIdAsyncResult, 1)
	go p.DeleteProfileByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) updateProfileByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateProfileByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateProfileByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateProfileByStampSheetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateProfileByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateProfileByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) UpdateProfileByStampSheetAsync(
	request *UpdateProfileByStampSheetRequest,
	callback chan<- UpdateProfileByStampSheetAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "profile",
			"function":    "updateProfileByStampSheet",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.updateProfileByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) UpdateProfileByStampSheet(
	request *UpdateProfileByStampSheetRequest,
) (*UpdateProfileByStampSheetResult, error) {
	callback := make(chan UpdateProfileByStampSheetAsyncResult, 1)
	go p.UpdateProfileByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) describeFriendsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeFriendsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFriendsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFriendsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFriendsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeFriendsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DescribeFriendsAsync(
	request *DescribeFriendsRequest,
	callback chan<- DescribeFriendsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "friend",
			"function":    "describeFriends",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeFriendsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DescribeFriends(
	request *DescribeFriendsRequest,
) (*DescribeFriendsResult, error) {
	callback := make(chan DescribeFriendsAsyncResult, 1)
	go p.DescribeFriendsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) describeFriendsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeFriendsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFriendsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFriendsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFriendsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeFriendsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DescribeFriendsByUserIdAsync(
	request *DescribeFriendsByUserIdRequest,
	callback chan<- DescribeFriendsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "friend",
			"function":    "describeFriendsByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeFriendsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DescribeFriendsByUserId(
	request *DescribeFriendsByUserIdRequest,
) (*DescribeFriendsByUserIdResult, error) {
	callback := make(chan DescribeFriendsByUserIdAsyncResult, 1)
	go p.DescribeFriendsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) describeBlackListAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeBlackListAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBlackListAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBlackListResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBlackListAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeBlackListAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DescribeBlackListAsync(
	request *DescribeBlackListRequest,
	callback chan<- DescribeBlackListAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "blackList",
			"function":    "describeBlackList",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeBlackListAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DescribeBlackList(
	request *DescribeBlackListRequest,
) (*DescribeBlackListResult, error) {
	callback := make(chan DescribeBlackListAsyncResult, 1)
	go p.DescribeBlackListAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) describeBlackListByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeBlackListByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBlackListByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBlackListByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBlackListByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeBlackListByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DescribeBlackListByUserIdAsync(
	request *DescribeBlackListByUserIdRequest,
	callback chan<- DescribeBlackListByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "blackList",
			"function":    "describeBlackListByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeBlackListByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DescribeBlackListByUserId(
	request *DescribeBlackListByUserIdRequest,
) (*DescribeBlackListByUserIdResult, error) {
	callback := make(chan DescribeBlackListByUserIdAsyncResult, 1)
	go p.DescribeBlackListByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) registerBlackListAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RegisterBlackListAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RegisterBlackListAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RegisterBlackListResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RegisterBlackListAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- RegisterBlackListAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) RegisterBlackListAsync(
	request *RegisterBlackListRequest,
	callback chan<- RegisterBlackListAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "blackList",
			"function":    "registerBlackList",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.registerBlackListAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) RegisterBlackList(
	request *RegisterBlackListRequest,
) (*RegisterBlackListResult, error) {
	callback := make(chan RegisterBlackListAsyncResult, 1)
	go p.RegisterBlackListAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) registerBlackListByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RegisterBlackListByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RegisterBlackListByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RegisterBlackListByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RegisterBlackListByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- RegisterBlackListByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) RegisterBlackListByUserIdAsync(
	request *RegisterBlackListByUserIdRequest,
	callback chan<- RegisterBlackListByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "blackList",
			"function":    "registerBlackListByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.registerBlackListByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) RegisterBlackListByUserId(
	request *RegisterBlackListByUserIdRequest,
) (*RegisterBlackListByUserIdResult, error) {
	callback := make(chan RegisterBlackListByUserIdAsyncResult, 1)
	go p.RegisterBlackListByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) unregisterBlackListAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UnregisterBlackListAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UnregisterBlackListAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UnregisterBlackListResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UnregisterBlackListAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UnregisterBlackListAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) UnregisterBlackListAsync(
	request *UnregisterBlackListRequest,
	callback chan<- UnregisterBlackListAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "blackList",
			"function":    "unregisterBlackList",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.unregisterBlackListAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) UnregisterBlackList(
	request *UnregisterBlackListRequest,
) (*UnregisterBlackListResult, error) {
	callback := make(chan UnregisterBlackListAsyncResult, 1)
	go p.UnregisterBlackListAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) unregisterBlackListByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UnregisterBlackListByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UnregisterBlackListByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UnregisterBlackListByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UnregisterBlackListByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UnregisterBlackListByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) UnregisterBlackListByUserIdAsync(
	request *UnregisterBlackListByUserIdRequest,
	callback chan<- UnregisterBlackListByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "blackList",
			"function":    "unregisterBlackListByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.unregisterBlackListByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) UnregisterBlackListByUserId(
	request *UnregisterBlackListByUserIdRequest,
) (*UnregisterBlackListByUserIdResult, error) {
	callback := make(chan UnregisterBlackListByUserIdAsyncResult, 1)
	go p.UnregisterBlackListByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) describeFollowsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeFollowsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFollowsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFollowsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFollowsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeFollowsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DescribeFollowsAsync(
	request *DescribeFollowsRequest,
	callback chan<- DescribeFollowsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "followUser",
			"function":    "describeFollows",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeFollowsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DescribeFollows(
	request *DescribeFollowsRequest,
) (*DescribeFollowsResult, error) {
	callback := make(chan DescribeFollowsAsyncResult, 1)
	go p.DescribeFollowsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) describeFollowsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeFollowsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFollowsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFollowsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFollowsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeFollowsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DescribeFollowsByUserIdAsync(
	request *DescribeFollowsByUserIdRequest,
	callback chan<- DescribeFollowsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "followUser",
			"function":    "describeFollowsByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeFollowsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DescribeFollowsByUserId(
	request *DescribeFollowsByUserIdRequest,
) (*DescribeFollowsByUserIdResult, error) {
	callback := make(chan DescribeFollowsByUserIdAsyncResult, 1)
	go p.DescribeFollowsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) getFollowAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetFollowAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFollowAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFollowResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFollowAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetFollowAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) GetFollowAsync(
	request *GetFollowRequest,
	callback chan<- GetFollowAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "followUser",
			"function":    "getFollow",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getFollowAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetFollow(
	request *GetFollowRequest,
) (*GetFollowResult, error) {
	callback := make(chan GetFollowAsyncResult, 1)
	go p.GetFollowAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) getFollowByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetFollowByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFollowByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFollowByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFollowByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetFollowByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) GetFollowByUserIdAsync(
	request *GetFollowByUserIdRequest,
	callback chan<- GetFollowByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "followUser",
			"function":    "getFollowByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getFollowByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetFollowByUserId(
	request *GetFollowByUserIdRequest,
) (*GetFollowByUserIdResult, error) {
	callback := make(chan GetFollowByUserIdAsyncResult, 1)
	go p.GetFollowByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) followAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- FollowAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- FollowAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result FollowResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- FollowAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- FollowAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) FollowAsync(
	request *FollowRequest,
	callback chan<- FollowAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "followUser",
			"function":    "follow",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.followAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) Follow(
	request *FollowRequest,
) (*FollowResult, error) {
	callback := make(chan FollowAsyncResult, 1)
	go p.FollowAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) followByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- FollowByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- FollowByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result FollowByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- FollowByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- FollowByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) FollowByUserIdAsync(
	request *FollowByUserIdRequest,
	callback chan<- FollowByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "followUser",
			"function":    "followByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.followByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) FollowByUserId(
	request *FollowByUserIdRequest,
) (*FollowByUserIdResult, error) {
	callback := make(chan FollowByUserIdAsyncResult, 1)
	go p.FollowByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) unfollowAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UnfollowAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UnfollowAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UnfollowResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UnfollowAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UnfollowAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) UnfollowAsync(
	request *UnfollowRequest,
	callback chan<- UnfollowAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "followUser",
			"function":    "unfollow",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.unfollowAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) Unfollow(
	request *UnfollowRequest,
) (*UnfollowResult, error) {
	callback := make(chan UnfollowAsyncResult, 1)
	go p.UnfollowAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) unfollowByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UnfollowByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UnfollowByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UnfollowByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UnfollowByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UnfollowByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) UnfollowByUserIdAsync(
	request *UnfollowByUserIdRequest,
	callback chan<- UnfollowByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "followUser",
			"function":    "unfollowByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.unfollowByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) UnfollowByUserId(
	request *UnfollowByUserIdRequest,
) (*UnfollowByUserIdResult, error) {
	callback := make(chan UnfollowByUserIdAsyncResult, 1)
	go p.UnfollowByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) getFriendAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetFriendAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFriendAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFriendResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFriendAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetFriendAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) GetFriendAsync(
	request *GetFriendRequest,
	callback chan<- GetFriendAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "friendUser",
			"function":    "getFriend",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getFriendAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetFriend(
	request *GetFriendRequest,
) (*GetFriendResult, error) {
	callback := make(chan GetFriendAsyncResult, 1)
	go p.GetFriendAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) getFriendByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetFriendByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFriendByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFriendByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFriendByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetFriendByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) GetFriendByUserIdAsync(
	request *GetFriendByUserIdRequest,
	callback chan<- GetFriendByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "friendUser",
			"function":    "getFriendByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getFriendByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetFriendByUserId(
	request *GetFriendByUserIdRequest,
) (*GetFriendByUserIdResult, error) {
	callback := make(chan GetFriendByUserIdAsyncResult, 1)
	go p.GetFriendByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) addFriendAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddFriendAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddFriendAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddFriendResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddFriendAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- AddFriendAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) AddFriendAsync(
	request *AddFriendRequest,
	callback chan<- AddFriendAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "friendUser",
			"function":    "addFriend",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.addFriendAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) AddFriend(
	request *AddFriendRequest,
) (*AddFriendResult, error) {
	callback := make(chan AddFriendAsyncResult, 1)
	go p.AddFriendAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) addFriendByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddFriendByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddFriendByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddFriendByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddFriendByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- AddFriendByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) AddFriendByUserIdAsync(
	request *AddFriendByUserIdRequest,
	callback chan<- AddFriendByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "friendUser",
			"function":    "addFriendByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.addFriendByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) AddFriendByUserId(
	request *AddFriendByUserIdRequest,
) (*AddFriendByUserIdResult, error) {
	callback := make(chan AddFriendByUserIdAsyncResult, 1)
	go p.AddFriendByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) deleteFriendAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteFriendAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteFriendAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteFriendResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteFriendAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteFriendAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DeleteFriendAsync(
	request *DeleteFriendRequest,
	callback chan<- DeleteFriendAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "friendUser",
			"function":    "deleteFriend",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.deleteFriendAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DeleteFriend(
	request *DeleteFriendRequest,
) (*DeleteFriendResult, error) {
	callback := make(chan DeleteFriendAsyncResult, 1)
	go p.DeleteFriendAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) deleteFriendByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteFriendByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteFriendByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteFriendByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteFriendByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteFriendByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DeleteFriendByUserIdAsync(
	request *DeleteFriendByUserIdRequest,
	callback chan<- DeleteFriendByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "friendUser",
			"function":    "deleteFriendByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.deleteFriendByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DeleteFriendByUserId(
	request *DeleteFriendByUserIdRequest,
) (*DeleteFriendByUserIdResult, error) {
	callback := make(chan DeleteFriendByUserIdAsyncResult, 1)
	go p.DeleteFriendByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) describeSendRequestsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSendRequestsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSendRequestsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSendRequestsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSendRequestsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeSendRequestsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DescribeSendRequestsAsync(
	request *DescribeSendRequestsRequest,
	callback chan<- DescribeSendRequestsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "sendFriendRequest",
			"function":    "describeSendRequests",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeSendRequestsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DescribeSendRequests(
	request *DescribeSendRequestsRequest,
) (*DescribeSendRequestsResult, error) {
	callback := make(chan DescribeSendRequestsAsyncResult, 1)
	go p.DescribeSendRequestsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) describeSendRequestsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSendRequestsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSendRequestsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSendRequestsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSendRequestsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeSendRequestsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DescribeSendRequestsByUserIdAsync(
	request *DescribeSendRequestsByUserIdRequest,
	callback chan<- DescribeSendRequestsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "sendFriendRequest",
			"function":    "describeSendRequestsByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeSendRequestsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DescribeSendRequestsByUserId(
	request *DescribeSendRequestsByUserIdRequest,
) (*DescribeSendRequestsByUserIdResult, error) {
	callback := make(chan DescribeSendRequestsByUserIdAsyncResult, 1)
	go p.DescribeSendRequestsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) getSendRequestAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSendRequestAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSendRequestAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSendRequestResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSendRequestAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetSendRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) GetSendRequestAsync(
	request *GetSendRequestRequest,
	callback chan<- GetSendRequestAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "sendFriendRequest",
			"function":    "getSendRequest",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getSendRequestAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetSendRequest(
	request *GetSendRequestRequest,
) (*GetSendRequestResult, error) {
	callback := make(chan GetSendRequestAsyncResult, 1)
	go p.GetSendRequestAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) getSendRequestByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSendRequestByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSendRequestByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSendRequestByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSendRequestByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetSendRequestByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) GetSendRequestByUserIdAsync(
	request *GetSendRequestByUserIdRequest,
	callback chan<- GetSendRequestByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "sendFriendRequest",
			"function":    "getSendRequestByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getSendRequestByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetSendRequestByUserId(
	request *GetSendRequestByUserIdRequest,
) (*GetSendRequestByUserIdResult, error) {
	callback := make(chan GetSendRequestByUserIdAsyncResult, 1)
	go p.GetSendRequestByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) sendRequestAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SendRequestAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SendRequestAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SendRequestResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SendRequestAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- SendRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) SendRequestAsync(
	request *SendRequestRequest,
	callback chan<- SendRequestAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "sendFriendRequest",
			"function":    "sendRequest",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.sendRequestAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) SendRequest(
	request *SendRequestRequest,
) (*SendRequestResult, error) {
	callback := make(chan SendRequestAsyncResult, 1)
	go p.SendRequestAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) sendRequestByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SendRequestByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SendRequestByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SendRequestByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SendRequestByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- SendRequestByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) SendRequestByUserIdAsync(
	request *SendRequestByUserIdRequest,
	callback chan<- SendRequestByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "sendFriendRequest",
			"function":    "sendRequestByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.sendRequestByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) SendRequestByUserId(
	request *SendRequestByUserIdRequest,
) (*SendRequestByUserIdResult, error) {
	callback := make(chan SendRequestByUserIdAsyncResult, 1)
	go p.SendRequestByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) deleteRequestAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteRequestAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRequestAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRequestResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteRequestAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DeleteRequestAsync(
	request *DeleteRequestRequest,
	callback chan<- DeleteRequestAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "sendFriendRequest",
			"function":    "deleteRequest",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.deleteRequestAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DeleteRequest(
	request *DeleteRequestRequest,
) (*DeleteRequestResult, error) {
	callback := make(chan DeleteRequestAsyncResult, 1)
	go p.DeleteRequestAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) deleteRequestByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteRequestByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRequestByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRequestByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteRequestByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteRequestByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DeleteRequestByUserIdAsync(
	request *DeleteRequestByUserIdRequest,
	callback chan<- DeleteRequestByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "sendFriendRequest",
			"function":    "deleteRequestByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.deleteRequestByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DeleteRequestByUserId(
	request *DeleteRequestByUserIdRequest,
) (*DeleteRequestByUserIdResult, error) {
	callback := make(chan DeleteRequestByUserIdAsyncResult, 1)
	go p.DeleteRequestByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) describeReceiveRequestsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeReceiveRequestsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReceiveRequestsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReceiveRequestsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeReceiveRequestsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeReceiveRequestsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DescribeReceiveRequestsAsync(
	request *DescribeReceiveRequestsRequest,
	callback chan<- DescribeReceiveRequestsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "receiveFriendRequest",
			"function":    "describeReceiveRequests",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeReceiveRequestsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DescribeReceiveRequests(
	request *DescribeReceiveRequestsRequest,
) (*DescribeReceiveRequestsResult, error) {
	callback := make(chan DescribeReceiveRequestsAsyncResult, 1)
	go p.DescribeReceiveRequestsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) describeReceiveRequestsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeReceiveRequestsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReceiveRequestsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReceiveRequestsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeReceiveRequestsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeReceiveRequestsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) DescribeReceiveRequestsByUserIdAsync(
	request *DescribeReceiveRequestsByUserIdRequest,
	callback chan<- DescribeReceiveRequestsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "receiveFriendRequest",
			"function":    "describeReceiveRequestsByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeReceiveRequestsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) DescribeReceiveRequestsByUserId(
	request *DescribeReceiveRequestsByUserIdRequest,
) (*DescribeReceiveRequestsByUserIdResult, error) {
	callback := make(chan DescribeReceiveRequestsByUserIdAsyncResult, 1)
	go p.DescribeReceiveRequestsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) getReceiveRequestAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetReceiveRequestAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReceiveRequestAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReceiveRequestResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetReceiveRequestAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetReceiveRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) GetReceiveRequestAsync(
	request *GetReceiveRequestRequest,
	callback chan<- GetReceiveRequestAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "receiveFriendRequest",
			"function":    "getReceiveRequest",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.FromUserId != nil && *request.FromUserId != "" {
		bodies["fromUserId"] = *request.FromUserId
	}
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getReceiveRequestAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetReceiveRequest(
	request *GetReceiveRequestRequest,
) (*GetReceiveRequestResult, error) {
	callback := make(chan GetReceiveRequestAsyncResult, 1)
	go p.GetReceiveRequestAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) getReceiveRequestByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetReceiveRequestByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReceiveRequestByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReceiveRequestByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetReceiveRequestByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetReceiveRequestByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) GetReceiveRequestByUserIdAsync(
	request *GetReceiveRequestByUserIdRequest,
	callback chan<- GetReceiveRequestByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "receiveFriendRequest",
			"function":    "getReceiveRequestByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.FromUserId != nil && *request.FromUserId != "" {
		bodies["fromUserId"] = *request.FromUserId
	}
	if request.WithProfile != nil {
		bodies["withProfile"] = *request.WithProfile
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getReceiveRequestByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetReceiveRequestByUserId(
	request *GetReceiveRequestByUserIdRequest,
) (*GetReceiveRequestByUserIdResult, error) {
	callback := make(chan GetReceiveRequestByUserIdAsyncResult, 1)
	go p.GetReceiveRequestByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) acceptRequestAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AcceptRequestAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcceptRequestAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcceptRequestResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcceptRequestAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- AcceptRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) AcceptRequestAsync(
	request *AcceptRequestRequest,
	callback chan<- AcceptRequestAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "receiveFriendRequest",
			"function":    "acceptRequest",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.FromUserId != nil && *request.FromUserId != "" {
		bodies["fromUserId"] = *request.FromUserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.acceptRequestAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) AcceptRequest(
	request *AcceptRequestRequest,
) (*AcceptRequestResult, error) {
	callback := make(chan AcceptRequestAsyncResult, 1)
	go p.AcceptRequestAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) acceptRequestByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AcceptRequestByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcceptRequestByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcceptRequestByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcceptRequestByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- AcceptRequestByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) AcceptRequestByUserIdAsync(
	request *AcceptRequestByUserIdRequest,
	callback chan<- AcceptRequestByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "receiveFriendRequest",
			"function":    "acceptRequestByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.FromUserId != nil && *request.FromUserId != "" {
		bodies["fromUserId"] = *request.FromUserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.acceptRequestByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) AcceptRequestByUserId(
	request *AcceptRequestByUserIdRequest,
) (*AcceptRequestByUserIdResult, error) {
	callback := make(chan AcceptRequestByUserIdAsyncResult, 1)
	go p.AcceptRequestByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) rejectRequestAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RejectRequestAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RejectRequestAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RejectRequestResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RejectRequestAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- RejectRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) RejectRequestAsync(
	request *RejectRequestRequest,
	callback chan<- RejectRequestAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "receiveFriendRequest",
			"function":    "rejectRequest",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.FromUserId != nil && *request.FromUserId != "" {
		bodies["fromUserId"] = *request.FromUserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.rejectRequestAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) RejectRequest(
	request *RejectRequestRequest,
) (*RejectRequestResult, error) {
	callback := make(chan RejectRequestAsyncResult, 1)
	go p.RejectRequestAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) rejectRequestByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RejectRequestByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RejectRequestByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RejectRequestByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RejectRequestByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- RejectRequestByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) RejectRequestByUserIdAsync(
	request *RejectRequestByUserIdRequest,
	callback chan<- RejectRequestByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "receiveFriendRequest",
			"function":    "rejectRequestByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.FromUserId != nil && *request.FromUserId != "" {
		bodies["fromUserId"] = *request.FromUserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.rejectRequestByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) RejectRequestByUserId(
	request *RejectRequestByUserIdRequest,
) (*RejectRequestByUserIdResult, error) {
	callback := make(chan RejectRequestByUserIdAsyncResult, 1)
	go p.RejectRequestByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2FriendWebSocketClient) getPublicProfileAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetPublicProfileAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPublicProfileAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPublicProfileResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPublicProfileAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetPublicProfileAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FriendWebSocketClient) GetPublicProfileAsync(
	request *GetPublicProfileRequest,
	callback chan<- GetPublicProfileAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "friend",
			"component":   "publicProfile",
			"function":    "getPublicProfile",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getPublicProfileAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2FriendWebSocketClient) GetPublicProfile(
	request *GetPublicProfileRequest,
) (*GetPublicProfileResult, error) {
	callback := make(chan GetPublicProfileAsyncResult, 1)
	go p.GetPublicProfileAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
