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

package guild

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2GuildWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2GuildWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2GuildWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2GuildWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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

func (p Gs2GuildWebSocketClient) DescribeNamespaces(
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

func (p Gs2GuildWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2GuildWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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
	if request.ChangeNotification != nil {
		bodies["changeNotification"] = request.ChangeNotification.ToDict()
	}
	if request.JoinNotification != nil {
		bodies["joinNotification"] = request.JoinNotification.ToDict()
	}
	if request.LeaveNotification != nil {
		bodies["leaveNotification"] = request.LeaveNotification.ToDict()
	}
	if request.ChangeMemberNotification != nil {
		bodies["changeMemberNotification"] = request.ChangeMemberNotification.ToDict()
	}
	if request.ReceiveRequestNotification != nil {
		bodies["receiveRequestNotification"] = request.ReceiveRequestNotification.ToDict()
	}
	if request.RemoveRequestNotification != nil {
		bodies["removeRequestNotification"] = request.RemoveRequestNotification.ToDict()
	}
	if request.CreateGuildScript != nil {
		bodies["createGuildScript"] = request.CreateGuildScript.ToDict()
	}
	if request.UpdateGuildScript != nil {
		bodies["updateGuildScript"] = request.UpdateGuildScript.ToDict()
	}
	if request.JoinGuildScript != nil {
		bodies["joinGuildScript"] = request.JoinGuildScript.ToDict()
	}
	if request.ReceiveJoinRequestScript != nil {
		bodies["receiveJoinRequestScript"] = request.ReceiveJoinRequestScript.ToDict()
	}
	if request.LeaveGuildScript != nil {
		bodies["leaveGuildScript"] = request.LeaveGuildScript.ToDict()
	}
	if request.ChangeRoleScript != nil {
		bodies["changeRoleScript"] = request.ChangeRoleScript.ToDict()
	}
	if request.DeleteGuildScript != nil {
		bodies["deleteGuildScript"] = request.DeleteGuildScript.ToDict()
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

func (p Gs2GuildWebSocketClient) CreateNamespace(
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

func (p Gs2GuildWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2GuildWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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

func (p Gs2GuildWebSocketClient) GetNamespaceStatus(
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

func (p Gs2GuildWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2GuildWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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

func (p Gs2GuildWebSocketClient) GetNamespace(
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

func (p Gs2GuildWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2GuildWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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
	if request.ChangeNotification != nil {
		bodies["changeNotification"] = request.ChangeNotification.ToDict()
	}
	if request.JoinNotification != nil {
		bodies["joinNotification"] = request.JoinNotification.ToDict()
	}
	if request.LeaveNotification != nil {
		bodies["leaveNotification"] = request.LeaveNotification.ToDict()
	}
	if request.ChangeMemberNotification != nil {
		bodies["changeMemberNotification"] = request.ChangeMemberNotification.ToDict()
	}
	if request.ReceiveRequestNotification != nil {
		bodies["receiveRequestNotification"] = request.ReceiveRequestNotification.ToDict()
	}
	if request.RemoveRequestNotification != nil {
		bodies["removeRequestNotification"] = request.RemoveRequestNotification.ToDict()
	}
	if request.CreateGuildScript != nil {
		bodies["createGuildScript"] = request.CreateGuildScript.ToDict()
	}
	if request.UpdateGuildScript != nil {
		bodies["updateGuildScript"] = request.UpdateGuildScript.ToDict()
	}
	if request.JoinGuildScript != nil {
		bodies["joinGuildScript"] = request.JoinGuildScript.ToDict()
	}
	if request.ReceiveJoinRequestScript != nil {
		bodies["receiveJoinRequestScript"] = request.ReceiveJoinRequestScript.ToDict()
	}
	if request.LeaveGuildScript != nil {
		bodies["leaveGuildScript"] = request.LeaveGuildScript.ToDict()
	}
	if request.ChangeRoleScript != nil {
		bodies["changeRoleScript"] = request.ChangeRoleScript.ToDict()
	}
	if request.DeleteGuildScript != nil {
		bodies["deleteGuildScript"] = request.DeleteGuildScript.ToDict()
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

func (p Gs2GuildWebSocketClient) UpdateNamespace(
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

func (p Gs2GuildWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2GuildWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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

func (p Gs2GuildWebSocketClient) DeleteNamespace(
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

func (p Gs2GuildWebSocketClient) getServiceVersionAsyncHandler(
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

func (p Gs2GuildWebSocketClient) GetServiceVersionAsync(
	request *GetServiceVersionRequest,
	callback chan<- GetServiceVersionAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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

func (p Gs2GuildWebSocketClient) GetServiceVersion(
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

func (p Gs2GuildWebSocketClient) dumpUserDataByUserIdAsyncHandler(
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

func (p Gs2GuildWebSocketClient) DumpUserDataByUserIdAsync(
	request *DumpUserDataByUserIdRequest,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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

func (p Gs2GuildWebSocketClient) DumpUserDataByUserId(
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

func (p Gs2GuildWebSocketClient) checkDumpUserDataByUserIdAsyncHandler(
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

func (p Gs2GuildWebSocketClient) CheckDumpUserDataByUserIdAsync(
	request *CheckDumpUserDataByUserIdRequest,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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

func (p Gs2GuildWebSocketClient) CheckDumpUserDataByUserId(
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

func (p Gs2GuildWebSocketClient) cleanUserDataByUserIdAsyncHandler(
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

func (p Gs2GuildWebSocketClient) CleanUserDataByUserIdAsync(
	request *CleanUserDataByUserIdRequest,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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

func (p Gs2GuildWebSocketClient) CleanUserDataByUserId(
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

func (p Gs2GuildWebSocketClient) checkCleanUserDataByUserIdAsyncHandler(
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

func (p Gs2GuildWebSocketClient) CheckCleanUserDataByUserIdAsync(
	request *CheckCleanUserDataByUserIdRequest,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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

func (p Gs2GuildWebSocketClient) CheckCleanUserDataByUserId(
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

func (p Gs2GuildWebSocketClient) prepareImportUserDataByUserIdAsyncHandler(
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

func (p Gs2GuildWebSocketClient) PrepareImportUserDataByUserIdAsync(
	request *PrepareImportUserDataByUserIdRequest,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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

func (p Gs2GuildWebSocketClient) PrepareImportUserDataByUserId(
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

func (p Gs2GuildWebSocketClient) importUserDataByUserIdAsyncHandler(
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

func (p Gs2GuildWebSocketClient) ImportUserDataByUserIdAsync(
	request *ImportUserDataByUserIdRequest,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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

func (p Gs2GuildWebSocketClient) ImportUserDataByUserId(
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

func (p Gs2GuildWebSocketClient) checkImportUserDataByUserIdAsyncHandler(
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

func (p Gs2GuildWebSocketClient) CheckImportUserDataByUserIdAsync(
	request *CheckImportUserDataByUserIdRequest,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
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

func (p Gs2GuildWebSocketClient) CheckImportUserDataByUserId(
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

func (p Gs2GuildWebSocketClient) describeGuildModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGuildModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGuildModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGuildModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGuildModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeGuildModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DescribeGuildModelMastersAsync(
	request *DescribeGuildModelMastersRequest,
	callback chan<- DescribeGuildModelMastersAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guildModelMaster",
			"function":    "describeGuildModelMasters",
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

	go p.describeGuildModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DescribeGuildModelMasters(
	request *DescribeGuildModelMastersRequest,
) (*DescribeGuildModelMastersResult, error) {
	callback := make(chan DescribeGuildModelMastersAsyncResult, 1)
	go p.DescribeGuildModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) createGuildModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateGuildModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGuildModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGuildModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGuildModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateGuildModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) CreateGuildModelMasterAsync(
	request *CreateGuildModelMasterRequest,
	callback chan<- CreateGuildModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guildModelMaster",
			"function":    "createGuildModelMaster",
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
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.DefaultMaximumMemberCount != nil {
		bodies["defaultMaximumMemberCount"] = *request.DefaultMaximumMemberCount
	}
	if request.MaximumMemberCount != nil {
		bodies["maximumMemberCount"] = *request.MaximumMemberCount
	}
	if request.InactivityPeriodDays != nil {
		bodies["inactivityPeriodDays"] = *request.InactivityPeriodDays
	}
	if request.Roles != nil {
		var _roles []interface{}
		for _, item := range request.Roles {
			_roles = append(_roles, item)
		}
		bodies["roles"] = _roles
	}
	if request.GuildMasterRole != nil && *request.GuildMasterRole != "" {
		bodies["guildMasterRole"] = *request.GuildMasterRole
	}
	if request.GuildMemberDefaultRole != nil && *request.GuildMemberDefaultRole != "" {
		bodies["guildMemberDefaultRole"] = *request.GuildMemberDefaultRole
	}
	if request.RejoinCoolTimeMinutes != nil {
		bodies["rejoinCoolTimeMinutes"] = *request.RejoinCoolTimeMinutes
	}
	if request.MaxConcurrentJoinGuilds != nil {
		bodies["maxConcurrentJoinGuilds"] = *request.MaxConcurrentJoinGuilds
	}
	if request.MaxConcurrentGuildMasterCount != nil {
		bodies["maxConcurrentGuildMasterCount"] = *request.MaxConcurrentGuildMasterCount
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

	go p.createGuildModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) CreateGuildModelMaster(
	request *CreateGuildModelMasterRequest,
) (*CreateGuildModelMasterResult, error) {
	callback := make(chan CreateGuildModelMasterAsyncResult, 1)
	go p.CreateGuildModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) getGuildModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGuildModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGuildModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGuildModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGuildModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGuildModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) GetGuildModelMasterAsync(
	request *GetGuildModelMasterRequest,
	callback chan<- GetGuildModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guildModelMaster",
			"function":    "getGuildModelMaster",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

	go p.getGuildModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) GetGuildModelMaster(
	request *GetGuildModelMasterRequest,
) (*GetGuildModelMasterResult, error) {
	callback := make(chan GetGuildModelMasterAsyncResult, 1)
	go p.GetGuildModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) updateGuildModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateGuildModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGuildModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGuildModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateGuildModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateGuildModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) UpdateGuildModelMasterAsync(
	request *UpdateGuildModelMasterRequest,
	callback chan<- UpdateGuildModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guildModelMaster",
			"function":    "updateGuildModelMaster",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.DefaultMaximumMemberCount != nil {
		bodies["defaultMaximumMemberCount"] = *request.DefaultMaximumMemberCount
	}
	if request.MaximumMemberCount != nil {
		bodies["maximumMemberCount"] = *request.MaximumMemberCount
	}
	if request.InactivityPeriodDays != nil {
		bodies["inactivityPeriodDays"] = *request.InactivityPeriodDays
	}
	if request.Roles != nil {
		var _roles []interface{}
		for _, item := range request.Roles {
			_roles = append(_roles, item)
		}
		bodies["roles"] = _roles
	}
	if request.GuildMasterRole != nil && *request.GuildMasterRole != "" {
		bodies["guildMasterRole"] = *request.GuildMasterRole
	}
	if request.GuildMemberDefaultRole != nil && *request.GuildMemberDefaultRole != "" {
		bodies["guildMemberDefaultRole"] = *request.GuildMemberDefaultRole
	}
	if request.RejoinCoolTimeMinutes != nil {
		bodies["rejoinCoolTimeMinutes"] = *request.RejoinCoolTimeMinutes
	}
	if request.MaxConcurrentJoinGuilds != nil {
		bodies["maxConcurrentJoinGuilds"] = *request.MaxConcurrentJoinGuilds
	}
	if request.MaxConcurrentGuildMasterCount != nil {
		bodies["maxConcurrentGuildMasterCount"] = *request.MaxConcurrentGuildMasterCount
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

	go p.updateGuildModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) UpdateGuildModelMaster(
	request *UpdateGuildModelMasterRequest,
) (*UpdateGuildModelMasterResult, error) {
	callback := make(chan UpdateGuildModelMasterAsyncResult, 1)
	go p.UpdateGuildModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) deleteGuildModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteGuildModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGuildModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGuildModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGuildModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteGuildModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DeleteGuildModelMasterAsync(
	request *DeleteGuildModelMasterRequest,
	callback chan<- DeleteGuildModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guildModelMaster",
			"function":    "deleteGuildModelMaster",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

	go p.deleteGuildModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DeleteGuildModelMaster(
	request *DeleteGuildModelMasterRequest,
) (*DeleteGuildModelMasterResult, error) {
	callback := make(chan DeleteGuildModelMasterAsyncResult, 1)
	go p.DeleteGuildModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) describeGuildModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGuildModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGuildModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGuildModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGuildModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeGuildModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DescribeGuildModelsAsync(
	request *DescribeGuildModelsRequest,
	callback chan<- DescribeGuildModelsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guildModel",
			"function":    "describeGuildModels",
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

	go p.describeGuildModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DescribeGuildModels(
	request *DescribeGuildModelsRequest,
) (*DescribeGuildModelsResult, error) {
	callback := make(chan DescribeGuildModelsAsyncResult, 1)
	go p.DescribeGuildModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) getGuildModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGuildModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGuildModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGuildModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGuildModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGuildModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) GetGuildModelAsync(
	request *GetGuildModelRequest,
	callback chan<- GetGuildModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guildModel",
			"function":    "getGuildModel",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

	go p.getGuildModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) GetGuildModel(
	request *GetGuildModelRequest,
) (*GetGuildModelResult, error) {
	callback := make(chan GetGuildModelAsyncResult, 1)
	go p.GetGuildModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) searchGuildsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SearchGuildsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SearchGuildsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SearchGuildsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SearchGuildsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- SearchGuildsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) SearchGuildsAsync(
	request *SearchGuildsRequest,
	callback chan<- SearchGuildsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "searchGuilds",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Attributes1 != nil {
		var _attributes1 []interface{}
		for _, item := range request.Attributes1 {
			_attributes1 = append(_attributes1, item)
		}
		bodies["attributes1"] = _attributes1
	}
	if request.Attributes2 != nil {
		var _attributes2 []interface{}
		for _, item := range request.Attributes2 {
			_attributes2 = append(_attributes2, item)
		}
		bodies["attributes2"] = _attributes2
	}
	if request.Attributes3 != nil {
		var _attributes3 []interface{}
		for _, item := range request.Attributes3 {
			_attributes3 = append(_attributes3, item)
		}
		bodies["attributes3"] = _attributes3
	}
	if request.Attributes4 != nil {
		var _attributes4 []interface{}
		for _, item := range request.Attributes4 {
			_attributes4 = append(_attributes4, item)
		}
		bodies["attributes4"] = _attributes4
	}
	if request.Attributes5 != nil {
		var _attributes5 []interface{}
		for _, item := range request.Attributes5 {
			_attributes5 = append(_attributes5, item)
		}
		bodies["attributes5"] = _attributes5
	}
	if request.JoinPolicies != nil {
		var _joinPolicies []interface{}
		for _, item := range request.JoinPolicies {
			_joinPolicies = append(_joinPolicies, item)
		}
		bodies["joinPolicies"] = _joinPolicies
	}
	if request.IncludeFullMembersGuild != nil {
		bodies["includeFullMembersGuild"] = *request.IncludeFullMembersGuild
	}
	if request.OrderBy != nil && *request.OrderBy != "" {
		bodies["orderBy"] = *request.OrderBy
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

	go p.searchGuildsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) SearchGuilds(
	request *SearchGuildsRequest,
) (*SearchGuildsResult, error) {
	callback := make(chan SearchGuildsAsyncResult, 1)
	go p.SearchGuildsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) searchGuildsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SearchGuildsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SearchGuildsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SearchGuildsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SearchGuildsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- SearchGuildsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) SearchGuildsByUserIdAsync(
	request *SearchGuildsByUserIdRequest,
	callback chan<- SearchGuildsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "searchGuildsByUserId",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Attributes1 != nil {
		var _attributes1 []interface{}
		for _, item := range request.Attributes1 {
			_attributes1 = append(_attributes1, item)
		}
		bodies["attributes1"] = _attributes1
	}
	if request.Attributes2 != nil {
		var _attributes2 []interface{}
		for _, item := range request.Attributes2 {
			_attributes2 = append(_attributes2, item)
		}
		bodies["attributes2"] = _attributes2
	}
	if request.Attributes3 != nil {
		var _attributes3 []interface{}
		for _, item := range request.Attributes3 {
			_attributes3 = append(_attributes3, item)
		}
		bodies["attributes3"] = _attributes3
	}
	if request.Attributes4 != nil {
		var _attributes4 []interface{}
		for _, item := range request.Attributes4 {
			_attributes4 = append(_attributes4, item)
		}
		bodies["attributes4"] = _attributes4
	}
	if request.Attributes5 != nil {
		var _attributes5 []interface{}
		for _, item := range request.Attributes5 {
			_attributes5 = append(_attributes5, item)
		}
		bodies["attributes5"] = _attributes5
	}
	if request.JoinPolicies != nil {
		var _joinPolicies []interface{}
		for _, item := range request.JoinPolicies {
			_joinPolicies = append(_joinPolicies, item)
		}
		bodies["joinPolicies"] = _joinPolicies
	}
	if request.IncludeFullMembersGuild != nil {
		bodies["includeFullMembersGuild"] = *request.IncludeFullMembersGuild
	}
	if request.OrderBy != nil && *request.OrderBy != "" {
		bodies["orderBy"] = *request.OrderBy
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

	go p.searchGuildsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) SearchGuildsByUserId(
	request *SearchGuildsByUserIdRequest,
) (*SearchGuildsByUserIdResult, error) {
	callback := make(chan SearchGuildsByUserIdAsyncResult, 1)
	go p.SearchGuildsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) createGuildAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateGuildAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGuildAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGuildResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGuildAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateGuildAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) CreateGuildAsync(
	request *CreateGuildRequest,
	callback chan<- CreateGuildAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "createGuild",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Attribute1 != nil {
		bodies["attribute1"] = *request.Attribute1
	}
	if request.Attribute2 != nil {
		bodies["attribute2"] = *request.Attribute2
	}
	if request.Attribute3 != nil {
		bodies["attribute3"] = *request.Attribute3
	}
	if request.Attribute4 != nil {
		bodies["attribute4"] = *request.Attribute4
	}
	if request.Attribute5 != nil {
		bodies["attribute5"] = *request.Attribute5
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.MemberMetadata != nil && *request.MemberMetadata != "" {
		bodies["memberMetadata"] = *request.MemberMetadata
	}
	if request.JoinPolicy != nil && *request.JoinPolicy != "" {
		bodies["joinPolicy"] = *request.JoinPolicy
	}
	if request.CustomRoles != nil {
		var _customRoles []interface{}
		for _, item := range request.CustomRoles {
			_customRoles = append(_customRoles, item)
		}
		bodies["customRoles"] = _customRoles
	}
	if request.GuildMemberDefaultRole != nil && *request.GuildMemberDefaultRole != "" {
		bodies["guildMemberDefaultRole"] = *request.GuildMemberDefaultRole
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

	go p.createGuildAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) CreateGuild(
	request *CreateGuildRequest,
) (*CreateGuildResult, error) {
	callback := make(chan CreateGuildAsyncResult, 1)
	go p.CreateGuildAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) createGuildByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateGuildByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGuildByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGuildByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGuildByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateGuildByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) CreateGuildByUserIdAsync(
	request *CreateGuildByUserIdRequest,
	callback chan<- CreateGuildByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "createGuildByUserId",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Attribute1 != nil {
		bodies["attribute1"] = *request.Attribute1
	}
	if request.Attribute2 != nil {
		bodies["attribute2"] = *request.Attribute2
	}
	if request.Attribute3 != nil {
		bodies["attribute3"] = *request.Attribute3
	}
	if request.Attribute4 != nil {
		bodies["attribute4"] = *request.Attribute4
	}
	if request.Attribute5 != nil {
		bodies["attribute5"] = *request.Attribute5
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.MemberMetadata != nil && *request.MemberMetadata != "" {
		bodies["memberMetadata"] = *request.MemberMetadata
	}
	if request.JoinPolicy != nil && *request.JoinPolicy != "" {
		bodies["joinPolicy"] = *request.JoinPolicy
	}
	if request.CustomRoles != nil {
		var _customRoles []interface{}
		for _, item := range request.CustomRoles {
			_customRoles = append(_customRoles, item)
		}
		bodies["customRoles"] = _customRoles
	}
	if request.GuildMemberDefaultRole != nil && *request.GuildMemberDefaultRole != "" {
		bodies["guildMemberDefaultRole"] = *request.GuildMemberDefaultRole
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

	go p.createGuildByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) CreateGuildByUserId(
	request *CreateGuildByUserIdRequest,
) (*CreateGuildByUserIdResult, error) {
	callback := make(chan CreateGuildByUserIdAsyncResult, 1)
	go p.CreateGuildByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) getGuildAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGuildAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGuildAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGuildResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGuildAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGuildAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) GetGuildAsync(
	request *GetGuildRequest,
	callback chan<- GetGuildAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "getGuild",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.getGuildAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) GetGuild(
	request *GetGuildRequest,
) (*GetGuildResult, error) {
	callback := make(chan GetGuildAsyncResult, 1)
	go p.GetGuildAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) getGuildByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGuildByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGuildByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGuildByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGuildByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGuildByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) GetGuildByUserIdAsync(
	request *GetGuildByUserIdRequest,
	callback chan<- GetGuildByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "getGuildByUserId",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.getGuildByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) GetGuildByUserId(
	request *GetGuildByUserIdRequest,
) (*GetGuildByUserIdResult, error) {
	callback := make(chan GetGuildByUserIdAsyncResult, 1)
	go p.GetGuildByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) updateGuildAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateGuildAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGuildAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGuildResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateGuildAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateGuildAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) UpdateGuildAsync(
	request *UpdateGuildRequest,
	callback chan<- UpdateGuildAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "updateGuild",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Attribute1 != nil {
		bodies["attribute1"] = *request.Attribute1
	}
	if request.Attribute2 != nil {
		bodies["attribute2"] = *request.Attribute2
	}
	if request.Attribute3 != nil {
		bodies["attribute3"] = *request.Attribute3
	}
	if request.Attribute4 != nil {
		bodies["attribute4"] = *request.Attribute4
	}
	if request.Attribute5 != nil {
		bodies["attribute5"] = *request.Attribute5
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.JoinPolicy != nil && *request.JoinPolicy != "" {
		bodies["joinPolicy"] = *request.JoinPolicy
	}
	if request.CustomRoles != nil {
		var _customRoles []interface{}
		for _, item := range request.CustomRoles {
			_customRoles = append(_customRoles, item)
		}
		bodies["customRoles"] = _customRoles
	}
	if request.GuildMemberDefaultRole != nil && *request.GuildMemberDefaultRole != "" {
		bodies["guildMemberDefaultRole"] = *request.GuildMemberDefaultRole
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

	go p.updateGuildAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) UpdateGuild(
	request *UpdateGuildRequest,
) (*UpdateGuildResult, error) {
	callback := make(chan UpdateGuildAsyncResult, 1)
	go p.UpdateGuildAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) updateGuildByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateGuildByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGuildByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGuildByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateGuildByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateGuildByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) UpdateGuildByGuildNameAsync(
	request *UpdateGuildByGuildNameRequest,
	callback chan<- UpdateGuildByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "updateGuildByGuildName",
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
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Attribute1 != nil {
		bodies["attribute1"] = *request.Attribute1
	}
	if request.Attribute2 != nil {
		bodies["attribute2"] = *request.Attribute2
	}
	if request.Attribute3 != nil {
		bodies["attribute3"] = *request.Attribute3
	}
	if request.Attribute4 != nil {
		bodies["attribute4"] = *request.Attribute4
	}
	if request.Attribute5 != nil {
		bodies["attribute5"] = *request.Attribute5
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.JoinPolicy != nil && *request.JoinPolicy != "" {
		bodies["joinPolicy"] = *request.JoinPolicy
	}
	if request.CustomRoles != nil {
		var _customRoles []interface{}
		for _, item := range request.CustomRoles {
			_customRoles = append(_customRoles, item)
		}
		bodies["customRoles"] = _customRoles
	}
	if request.GuildMemberDefaultRole != nil && *request.GuildMemberDefaultRole != "" {
		bodies["guildMemberDefaultRole"] = *request.GuildMemberDefaultRole
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

	go p.updateGuildByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) UpdateGuildByGuildName(
	request *UpdateGuildByGuildNameRequest,
) (*UpdateGuildByGuildNameResult, error) {
	callback := make(chan UpdateGuildByGuildNameAsyncResult, 1)
	go p.UpdateGuildByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) deleteMemberAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteMemberAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMemberAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMemberResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMemberAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.member.master.require" {
				asyncResult.Err = gs2err.SetClientError(GuildMasterRequired{})
			}
		}
	}
	callback <- DeleteMemberAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DeleteMemberAsync(
	request *DeleteMemberRequest,
	callback chan<- DeleteMemberAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "deleteMember",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

	go p.deleteMemberAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DeleteMember(
	request *DeleteMemberRequest,
) (*DeleteMemberResult, error) {
	callback := make(chan DeleteMemberAsyncResult, 1)
	go p.DeleteMemberAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) deleteMemberByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteMemberByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMemberByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMemberByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMemberByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.member.master.require" {
				asyncResult.Err = gs2err.SetClientError(GuildMasterRequired{})
			}
		}
	}
	callback <- DeleteMemberByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DeleteMemberByGuildNameAsync(
	request *DeleteMemberByGuildNameRequest,
	callback chan<- DeleteMemberByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "deleteMemberByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
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

	go p.deleteMemberByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DeleteMemberByGuildName(
	request *DeleteMemberByGuildNameRequest,
) (*DeleteMemberByGuildNameResult, error) {
	callback := make(chan DeleteMemberByGuildNameAsyncResult, 1)
	go p.DeleteMemberByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) updateMemberRoleAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateMemberRoleAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMemberRoleAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMemberRoleResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMemberRoleAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateMemberRoleAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) UpdateMemberRoleAsync(
	request *UpdateMemberRoleRequest,
	callback chan<- UpdateMemberRoleAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "updateMemberRole",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.RoleName != nil && *request.RoleName != "" {
		bodies["roleName"] = *request.RoleName
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

	go p.updateMemberRoleAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) UpdateMemberRole(
	request *UpdateMemberRoleRequest,
) (*UpdateMemberRoleResult, error) {
	callback := make(chan UpdateMemberRoleAsyncResult, 1)
	go p.UpdateMemberRoleAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) updateMemberRoleByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateMemberRoleByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMemberRoleByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMemberRoleByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMemberRoleByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateMemberRoleByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) UpdateMemberRoleByGuildNameAsync(
	request *UpdateMemberRoleByGuildNameRequest,
	callback chan<- UpdateMemberRoleByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "updateMemberRoleByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		bodies["targetUserId"] = *request.TargetUserId
	}
	if request.RoleName != nil && *request.RoleName != "" {
		bodies["roleName"] = *request.RoleName
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

	go p.updateMemberRoleByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) UpdateMemberRoleByGuildName(
	request *UpdateMemberRoleByGuildNameRequest,
) (*UpdateMemberRoleByGuildNameResult, error) {
	callback := make(chan UpdateMemberRoleByGuildNameAsyncResult, 1)
	go p.UpdateMemberRoleByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) batchUpdateMemberRoleAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- BatchUpdateMemberRoleAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BatchUpdateMemberRoleAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BatchUpdateMemberRoleResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- BatchUpdateMemberRoleAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- BatchUpdateMemberRoleAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) BatchUpdateMemberRoleAsync(
	request *BatchUpdateMemberRoleRequest,
	callback chan<- BatchUpdateMemberRoleAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "batchUpdateMemberRole",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Members != nil {
		var _members []interface{}
		for _, item := range request.Members {
			_members = append(_members, item)
		}
		bodies["members"] = _members
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

	go p.batchUpdateMemberRoleAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) BatchUpdateMemberRole(
	request *BatchUpdateMemberRoleRequest,
) (*BatchUpdateMemberRoleResult, error) {
	callback := make(chan BatchUpdateMemberRoleAsyncResult, 1)
	go p.BatchUpdateMemberRoleAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) batchUpdateMemberRoleByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- BatchUpdateMemberRoleByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BatchUpdateMemberRoleByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BatchUpdateMemberRoleByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- BatchUpdateMemberRoleByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- BatchUpdateMemberRoleByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) BatchUpdateMemberRoleByGuildNameAsync(
	request *BatchUpdateMemberRoleByGuildNameRequest,
	callback chan<- BatchUpdateMemberRoleByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "batchUpdateMemberRoleByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.Members != nil {
		var _members []interface{}
		for _, item := range request.Members {
			_members = append(_members, item)
		}
		bodies["members"] = _members
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

	go p.batchUpdateMemberRoleByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) BatchUpdateMemberRoleByGuildName(
	request *BatchUpdateMemberRoleByGuildNameRequest,
) (*BatchUpdateMemberRoleByGuildNameResult, error) {
	callback := make(chan BatchUpdateMemberRoleByGuildNameAsyncResult, 1)
	go p.BatchUpdateMemberRoleByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) deleteGuildAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteGuildAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGuildAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGuildResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGuildAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteGuildAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DeleteGuildAsync(
	request *DeleteGuildRequest,
	callback chan<- DeleteGuildAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "deleteGuild",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

	go p.deleteGuildAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DeleteGuild(
	request *DeleteGuildRequest,
) (*DeleteGuildResult, error) {
	callback := make(chan DeleteGuildAsyncResult, 1)
	go p.DeleteGuildAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) deleteGuildByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteGuildByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGuildByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGuildByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGuildByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteGuildByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DeleteGuildByGuildNameAsync(
	request *DeleteGuildByGuildNameRequest,
	callback chan<- DeleteGuildByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "deleteGuildByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.deleteGuildByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DeleteGuildByGuildName(
	request *DeleteGuildByGuildNameRequest,
) (*DeleteGuildByGuildNameResult, error) {
	callback := make(chan DeleteGuildByGuildNameAsyncResult, 1)
	go p.DeleteGuildByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) increaseMaximumCurrentMaximumMemberCountByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result IncreaseMaximumCurrentMaximumMemberCountByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsync(
	request *IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest,
	callback chan<- IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "increaseMaximumCurrentMaximumMemberCountByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.Value != nil {
		bodies["value"] = *request.Value
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

	go p.increaseMaximumCurrentMaximumMemberCountByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) IncreaseMaximumCurrentMaximumMemberCountByGuildName(
	request *IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest,
) (*IncreaseMaximumCurrentMaximumMemberCountByGuildNameResult, error) {
	callback := make(chan IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult, 1)
	go p.IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) decreaseMaximumCurrentMaximumMemberCountAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DecreaseMaximumCurrentMaximumMemberCountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DecreaseMaximumCurrentMaximumMemberCountAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DecreaseMaximumCurrentMaximumMemberCountResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseMaximumCurrentMaximumMemberCountAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DecreaseMaximumCurrentMaximumMemberCountAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DecreaseMaximumCurrentMaximumMemberCountAsync(
	request *DecreaseMaximumCurrentMaximumMemberCountRequest,
	callback chan<- DecreaseMaximumCurrentMaximumMemberCountAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "decreaseMaximumCurrentMaximumMemberCount",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Value != nil {
		bodies["value"] = *request.Value
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

	go p.decreaseMaximumCurrentMaximumMemberCountAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DecreaseMaximumCurrentMaximumMemberCount(
	request *DecreaseMaximumCurrentMaximumMemberCountRequest,
) (*DecreaseMaximumCurrentMaximumMemberCountResult, error) {
	callback := make(chan DecreaseMaximumCurrentMaximumMemberCountAsyncResult, 1)
	go p.DecreaseMaximumCurrentMaximumMemberCountAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) decreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DecreaseMaximumCurrentMaximumMemberCountByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsync(
	request *DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest,
	callback chan<- DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "decreaseMaximumCurrentMaximumMemberCountByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.Value != nil {
		bodies["value"] = *request.Value
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

	go p.decreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DecreaseMaximumCurrentMaximumMemberCountByGuildName(
	request *DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest,
) (*DecreaseMaximumCurrentMaximumMemberCountByGuildNameResult, error) {
	callback := make(chan DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult, 1)
	go p.DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) verifyCurrentMaximumMemberCountAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyCurrentMaximumMemberCountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCurrentMaximumMemberCountAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCurrentMaximumMemberCountResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCurrentMaximumMemberCountAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- VerifyCurrentMaximumMemberCountAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) VerifyCurrentMaximumMemberCountAsync(
	request *VerifyCurrentMaximumMemberCountRequest,
	callback chan<- VerifyCurrentMaximumMemberCountAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "verifyCurrentMaximumMemberCount",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		bodies["verifyType"] = *request.VerifyType
	}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go p.verifyCurrentMaximumMemberCountAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) VerifyCurrentMaximumMemberCount(
	request *VerifyCurrentMaximumMemberCountRequest,
) (*VerifyCurrentMaximumMemberCountResult, error) {
	callback := make(chan VerifyCurrentMaximumMemberCountAsyncResult, 1)
	go p.VerifyCurrentMaximumMemberCountAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) verifyCurrentMaximumMemberCountByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCurrentMaximumMemberCountByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCurrentMaximumMemberCountByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCurrentMaximumMemberCountByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- VerifyCurrentMaximumMemberCountByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) VerifyCurrentMaximumMemberCountByGuildNameAsync(
	request *VerifyCurrentMaximumMemberCountByGuildNameRequest,
	callback chan<- VerifyCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "verifyCurrentMaximumMemberCountByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		bodies["verifyType"] = *request.VerifyType
	}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go p.verifyCurrentMaximumMemberCountByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) VerifyCurrentMaximumMemberCountByGuildName(
	request *VerifyCurrentMaximumMemberCountByGuildNameRequest,
) (*VerifyCurrentMaximumMemberCountByGuildNameResult, error) {
	callback := make(chan VerifyCurrentMaximumMemberCountByGuildNameAsyncResult, 1)
	go p.VerifyCurrentMaximumMemberCountByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) verifyIncludeMemberAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyIncludeMemberAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyIncludeMemberAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyIncludeMemberResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyIncludeMemberAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- VerifyIncludeMemberAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) VerifyIncludeMemberAsync(
	request *VerifyIncludeMemberRequest,
	callback chan<- VerifyIncludeMemberAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "verifyIncludeMember",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		bodies["verifyType"] = *request.VerifyType
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

	go p.verifyIncludeMemberAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) VerifyIncludeMember(
	request *VerifyIncludeMemberRequest,
) (*VerifyIncludeMemberResult, error) {
	callback := make(chan VerifyIncludeMemberAsyncResult, 1)
	go p.VerifyIncludeMemberAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) verifyIncludeMemberByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyIncludeMemberByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyIncludeMemberByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyIncludeMemberByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyIncludeMemberByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- VerifyIncludeMemberByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) VerifyIncludeMemberByUserIdAsync(
	request *VerifyIncludeMemberByUserIdRequest,
	callback chan<- VerifyIncludeMemberByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "verifyIncludeMemberByUserId",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		bodies["verifyType"] = *request.VerifyType
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

	go p.verifyIncludeMemberByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) VerifyIncludeMemberByUserId(
	request *VerifyIncludeMemberByUserIdRequest,
) (*VerifyIncludeMemberByUserIdResult, error) {
	callback := make(chan VerifyIncludeMemberByUserIdAsyncResult, 1)
	go p.VerifyIncludeMemberByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) setMaximumCurrentMaximumMemberCountByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetMaximumCurrentMaximumMemberCountByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) SetMaximumCurrentMaximumMemberCountByGuildNameAsync(
	request *SetMaximumCurrentMaximumMemberCountByGuildNameRequest,
	callback chan<- SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "setMaximumCurrentMaximumMemberCountByGuildName",
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
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.Value != nil {
		bodies["value"] = *request.Value
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

	go p.setMaximumCurrentMaximumMemberCountByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) SetMaximumCurrentMaximumMemberCountByGuildName(
	request *SetMaximumCurrentMaximumMemberCountByGuildNameRequest,
) (*SetMaximumCurrentMaximumMemberCountByGuildNameResult, error) {
	callback := make(chan SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult, 1)
	go p.SetMaximumCurrentMaximumMemberCountByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) assumeAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AssumeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AssumeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AssumeResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AssumeAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.member.notFound" {
				asyncResult.Err = gs2err.SetClientError(NotIncludedGuildMember{})
			}
		}
	}
	callback <- AssumeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) AssumeAsync(
	request *AssumeRequest,
	callback chan<- AssumeAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "assume",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.assumeAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) Assume(
	request *AssumeRequest,
) (*AssumeResult, error) {
	callback := make(chan AssumeAsyncResult, 1)
	go p.AssumeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) assumeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AssumeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AssumeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AssumeByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AssumeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.member.notFound" {
				asyncResult.Err = gs2err.SetClientError(NotIncludedGuildMember{})
			}
		}
	}
	callback <- AssumeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) AssumeByUserIdAsync(
	request *AssumeByUserIdRequest,
	callback chan<- AssumeByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "assumeByUserId",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.assumeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) AssumeByUserId(
	request *AssumeByUserIdRequest,
) (*AssumeByUserIdResult, error) {
	callback := make(chan AssumeByUserIdAsyncResult, 1)
	go p.AssumeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) increaseMaximumCurrentMaximumMemberCountByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result IncreaseMaximumCurrentMaximumMemberCountByStampSheetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsync(
	request *IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest,
	callback chan<- IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "increaseMaximumCurrentMaximumMemberCountByStampSheet",
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

	go p.increaseMaximumCurrentMaximumMemberCountByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) IncreaseMaximumCurrentMaximumMemberCountByStampSheet(
	request *IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest,
) (*IncreaseMaximumCurrentMaximumMemberCountByStampSheetResult, error) {
	callback := make(chan IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult, 1)
	go p.IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) decreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DecreaseMaximumCurrentMaximumMemberCountByStampTaskResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsync(
	request *DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest,
	callback chan<- DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "decreaseMaximumCurrentMaximumMemberCountByStampTask",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.decreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DecreaseMaximumCurrentMaximumMemberCountByStampTask(
	request *DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest,
) (*DecreaseMaximumCurrentMaximumMemberCountByStampTaskResult, error) {
	callback := make(chan DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult, 1)
	go p.DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) setMaximumCurrentMaximumMemberCountByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetMaximumCurrentMaximumMemberCountByStampSheetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) SetMaximumCurrentMaximumMemberCountByStampSheetAsync(
	request *SetMaximumCurrentMaximumMemberCountByStampSheetRequest,
	callback chan<- SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "setMaximumCurrentMaximumMemberCountByStampSheet",
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

	go p.setMaximumCurrentMaximumMemberCountByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) SetMaximumCurrentMaximumMemberCountByStampSheet(
	request *SetMaximumCurrentMaximumMemberCountByStampSheetRequest,
) (*SetMaximumCurrentMaximumMemberCountByStampSheetResult, error) {
	callback := make(chan SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult, 1)
	go p.SetMaximumCurrentMaximumMemberCountByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) verifyCurrentMaximumMemberCountByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyCurrentMaximumMemberCountByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCurrentMaximumMemberCountByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCurrentMaximumMemberCountByStampTaskResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCurrentMaximumMemberCountByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- VerifyCurrentMaximumMemberCountByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) VerifyCurrentMaximumMemberCountByStampTaskAsync(
	request *VerifyCurrentMaximumMemberCountByStampTaskRequest,
	callback chan<- VerifyCurrentMaximumMemberCountByStampTaskAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "verifyCurrentMaximumMemberCountByStampTask",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.verifyCurrentMaximumMemberCountByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) VerifyCurrentMaximumMemberCountByStampTask(
	request *VerifyCurrentMaximumMemberCountByStampTaskRequest,
) (*VerifyCurrentMaximumMemberCountByStampTaskResult, error) {
	callback := make(chan VerifyCurrentMaximumMemberCountByStampTaskAsyncResult, 1)
	go p.VerifyCurrentMaximumMemberCountByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) verifyIncludeMemberByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyIncludeMemberByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyIncludeMemberByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyIncludeMemberByStampTaskResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyIncludeMemberByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- VerifyIncludeMemberByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) VerifyIncludeMemberByStampTaskAsync(
	request *VerifyIncludeMemberByStampTaskRequest,
	callback chan<- VerifyIncludeMemberByStampTaskAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "guild",
			"function":    "verifyIncludeMemberByStampTask",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.verifyIncludeMemberByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) VerifyIncludeMemberByStampTask(
	request *VerifyIncludeMemberByStampTaskRequest,
) (*VerifyIncludeMemberByStampTaskResult, error) {
	callback := make(chan VerifyIncludeMemberByStampTaskAsyncResult, 1)
	go p.VerifyIncludeMemberByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) describeJoinedGuildsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeJoinedGuildsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeJoinedGuildsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeJoinedGuildsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeJoinedGuildsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeJoinedGuildsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DescribeJoinedGuildsAsync(
	request *DescribeJoinedGuildsRequest,
	callback chan<- DescribeJoinedGuildsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "joinedGuild",
			"function":    "describeJoinedGuilds",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

	go p.describeJoinedGuildsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DescribeJoinedGuilds(
	request *DescribeJoinedGuildsRequest,
) (*DescribeJoinedGuildsResult, error) {
	callback := make(chan DescribeJoinedGuildsAsyncResult, 1)
	go p.DescribeJoinedGuildsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) describeJoinedGuildsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeJoinedGuildsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeJoinedGuildsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeJoinedGuildsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeJoinedGuildsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeJoinedGuildsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DescribeJoinedGuildsByUserIdAsync(
	request *DescribeJoinedGuildsByUserIdRequest,
	callback chan<- DescribeJoinedGuildsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "joinedGuild",
			"function":    "describeJoinedGuildsByUserId",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

	go p.describeJoinedGuildsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DescribeJoinedGuildsByUserId(
	request *DescribeJoinedGuildsByUserIdRequest,
) (*DescribeJoinedGuildsByUserIdResult, error) {
	callback := make(chan DescribeJoinedGuildsByUserIdAsyncResult, 1)
	go p.DescribeJoinedGuildsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) getJoinedGuildAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetJoinedGuildAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetJoinedGuildAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetJoinedGuildResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetJoinedGuildAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetJoinedGuildAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) GetJoinedGuildAsync(
	request *GetJoinedGuildRequest,
	callback chan<- GetJoinedGuildAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "joinedGuild",
			"function":    "getJoinedGuild",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.getJoinedGuildAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) GetJoinedGuild(
	request *GetJoinedGuildRequest,
) (*GetJoinedGuildResult, error) {
	callback := make(chan GetJoinedGuildAsyncResult, 1)
	go p.GetJoinedGuildAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) getJoinedGuildByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetJoinedGuildByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetJoinedGuildByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetJoinedGuildByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetJoinedGuildByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetJoinedGuildByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) GetJoinedGuildByUserIdAsync(
	request *GetJoinedGuildByUserIdRequest,
	callback chan<- GetJoinedGuildByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "joinedGuild",
			"function":    "getJoinedGuildByUserId",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.getJoinedGuildByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) GetJoinedGuildByUserId(
	request *GetJoinedGuildByUserIdRequest,
) (*GetJoinedGuildByUserIdResult, error) {
	callback := make(chan GetJoinedGuildByUserIdAsyncResult, 1)
	go p.GetJoinedGuildByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) updateMemberMetadataAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateMemberMetadataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMemberMetadataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMemberMetadataResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMemberMetadataAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateMemberMetadataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) UpdateMemberMetadataAsync(
	request *UpdateMemberMetadataRequest,
	callback chan<- UpdateMemberMetadataAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "joinedGuild",
			"function":    "updateMemberMetadata",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
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

	go p.updateMemberMetadataAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) UpdateMemberMetadata(
	request *UpdateMemberMetadataRequest,
) (*UpdateMemberMetadataResult, error) {
	callback := make(chan UpdateMemberMetadataAsyncResult, 1)
	go p.UpdateMemberMetadataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) updateMemberMetadataByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateMemberMetadataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMemberMetadataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMemberMetadataByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMemberMetadataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateMemberMetadataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) UpdateMemberMetadataByUserIdAsync(
	request *UpdateMemberMetadataByUserIdRequest,
	callback chan<- UpdateMemberMetadataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "joinedGuild",
			"function":    "updateMemberMetadataByUserId",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
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

	go p.updateMemberMetadataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) UpdateMemberMetadataByUserId(
	request *UpdateMemberMetadataByUserIdRequest,
) (*UpdateMemberMetadataByUserIdResult, error) {
	callback := make(chan UpdateMemberMetadataByUserIdAsyncResult, 1)
	go p.UpdateMemberMetadataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) withdrawalAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- WithdrawalAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WithdrawalAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WithdrawalResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WithdrawalAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.member.master.require" {
				asyncResult.Err = gs2err.SetClientError(GuildMasterRequired{})
			}
		}
	}
	callback <- WithdrawalAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) WithdrawalAsync(
	request *WithdrawalRequest,
	callback chan<- WithdrawalAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "joinedGuild",
			"function":    "withdrawal",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.withdrawalAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) Withdrawal(
	request *WithdrawalRequest,
) (*WithdrawalResult, error) {
	callback := make(chan WithdrawalAsyncResult, 1)
	go p.WithdrawalAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) withdrawalByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- WithdrawalByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WithdrawalByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WithdrawalByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WithdrawalByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.member.master.require" {
				asyncResult.Err = gs2err.SetClientError(GuildMasterRequired{})
			}
		}
	}
	callback <- WithdrawalByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) WithdrawalByUserIdAsync(
	request *WithdrawalByUserIdRequest,
	callback chan<- WithdrawalByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "joinedGuild",
			"function":    "withdrawalByUserId",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.withdrawalByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) WithdrawalByUserId(
	request *WithdrawalByUserIdRequest,
) (*WithdrawalByUserIdResult, error) {
	callback := make(chan WithdrawalByUserIdAsyncResult, 1)
	go p.WithdrawalByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) getLastGuildMasterActivityAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetLastGuildMasterActivityAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLastGuildMasterActivityAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLastGuildMasterActivityResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLastGuildMasterActivityAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetLastGuildMasterActivityAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) GetLastGuildMasterActivityAsync(
	request *GetLastGuildMasterActivityRequest,
	callback chan<- GetLastGuildMasterActivityAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "lastGuildMasterActivity",
			"function":    "getLastGuildMasterActivity",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

	go p.getLastGuildMasterActivityAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) GetLastGuildMasterActivity(
	request *GetLastGuildMasterActivityRequest,
) (*GetLastGuildMasterActivityResult, error) {
	callback := make(chan GetLastGuildMasterActivityAsyncResult, 1)
	go p.GetLastGuildMasterActivityAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) getLastGuildMasterActivityByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetLastGuildMasterActivityByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLastGuildMasterActivityByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLastGuildMasterActivityByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLastGuildMasterActivityByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetLastGuildMasterActivityByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) GetLastGuildMasterActivityByGuildNameAsync(
	request *GetLastGuildMasterActivityByGuildNameRequest,
	callback chan<- GetLastGuildMasterActivityByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "lastGuildMasterActivity",
			"function":    "getLastGuildMasterActivityByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.getLastGuildMasterActivityByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) GetLastGuildMasterActivityByGuildName(
	request *GetLastGuildMasterActivityByGuildNameRequest,
) (*GetLastGuildMasterActivityByGuildNameResult, error) {
	callback := make(chan GetLastGuildMasterActivityByGuildNameAsyncResult, 1)
	go p.GetLastGuildMasterActivityByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) promoteSeniorMemberAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PromoteSeniorMemberAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PromoteSeniorMemberAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PromoteSeniorMemberResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PromoteSeniorMemberAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PromoteSeniorMemberAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) PromoteSeniorMemberAsync(
	request *PromoteSeniorMemberRequest,
	callback chan<- PromoteSeniorMemberAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "lastGuildMasterActivity",
			"function":    "promoteSeniorMember",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

	go p.promoteSeniorMemberAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) PromoteSeniorMember(
	request *PromoteSeniorMemberRequest,
) (*PromoteSeniorMemberResult, error) {
	callback := make(chan PromoteSeniorMemberAsyncResult, 1)
	go p.PromoteSeniorMemberAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) promoteSeniorMemberByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PromoteSeniorMemberByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PromoteSeniorMemberByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PromoteSeniorMemberByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PromoteSeniorMemberByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PromoteSeniorMemberByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) PromoteSeniorMemberByGuildNameAsync(
	request *PromoteSeniorMemberByGuildNameRequest,
	callback chan<- PromoteSeniorMemberByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "lastGuildMasterActivity",
			"function":    "promoteSeniorMemberByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.promoteSeniorMemberByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) PromoteSeniorMemberByGuildName(
	request *PromoteSeniorMemberByGuildNameRequest,
) (*PromoteSeniorMemberByGuildNameResult, error) {
	callback := make(chan PromoteSeniorMemberByGuildNameAsyncResult, 1)
	go p.PromoteSeniorMemberByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) exportMasterAsyncHandler(
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

func (p Gs2GuildWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "currentGuildMaster",
			"function":    "exportMaster",
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

	go p.exportMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) ExportMaster(
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

func (p Gs2GuildWebSocketClient) getCurrentGuildMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentGuildMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentGuildMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentGuildMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentGuildMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetCurrentGuildMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) GetCurrentGuildMasterAsync(
	request *GetCurrentGuildMasterRequest,
	callback chan<- GetCurrentGuildMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "currentGuildMaster",
			"function":    "getCurrentGuildMaster",
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

	go p.getCurrentGuildMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) GetCurrentGuildMaster(
	request *GetCurrentGuildMasterRequest,
) (*GetCurrentGuildMasterResult, error) {
	callback := make(chan GetCurrentGuildMasterAsyncResult, 1)
	go p.GetCurrentGuildMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) preUpdateCurrentGuildMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PreUpdateCurrentGuildMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreUpdateCurrentGuildMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreUpdateCurrentGuildMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreUpdateCurrentGuildMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PreUpdateCurrentGuildMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) PreUpdateCurrentGuildMasterAsync(
	request *PreUpdateCurrentGuildMasterRequest,
	callback chan<- PreUpdateCurrentGuildMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "currentGuildMaster",
			"function":    "preUpdateCurrentGuildMaster",
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

	go p.preUpdateCurrentGuildMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) PreUpdateCurrentGuildMaster(
	request *PreUpdateCurrentGuildMasterRequest,
) (*PreUpdateCurrentGuildMasterResult, error) {
	callback := make(chan PreUpdateCurrentGuildMasterAsyncResult, 1)
	go p.PreUpdateCurrentGuildMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) updateCurrentGuildMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentGuildMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentGuildMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentGuildMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentGuildMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateCurrentGuildMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) UpdateCurrentGuildMasterAsync(
	request *UpdateCurrentGuildMasterRequest,
	callback chan<- UpdateCurrentGuildMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "currentGuildMaster",
			"function":    "updateCurrentGuildMaster",
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
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Settings != nil && *request.Settings != "" {
		bodies["settings"] = *request.Settings
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
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

	go p.updateCurrentGuildMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) UpdateCurrentGuildMaster(
	request *UpdateCurrentGuildMasterRequest,
) (*UpdateCurrentGuildMasterResult, error) {
	callback := make(chan UpdateCurrentGuildMasterAsyncResult, 1)
	go p.UpdateCurrentGuildMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) updateCurrentGuildMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentGuildMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentGuildMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentGuildMasterFromGitHubResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentGuildMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateCurrentGuildMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) UpdateCurrentGuildMasterFromGitHubAsync(
	request *UpdateCurrentGuildMasterFromGitHubRequest,
	callback chan<- UpdateCurrentGuildMasterFromGitHubAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "currentGuildMaster",
			"function":    "updateCurrentGuildMasterFromGitHub",
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
	if request.CheckoutSetting != nil {
		bodies["checkoutSetting"] = request.CheckoutSetting.ToDict()
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

	go p.updateCurrentGuildMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) UpdateCurrentGuildMasterFromGitHub(
	request *UpdateCurrentGuildMasterFromGitHubRequest,
) (*UpdateCurrentGuildMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentGuildMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentGuildMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) describeReceiveRequestsAsyncHandler(
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

func (p Gs2GuildWebSocketClient) DescribeReceiveRequestsAsync(
	request *DescribeReceiveRequestsRequest,
	callback chan<- DescribeReceiveRequestsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "receiveMemberRequest",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

	go p.describeReceiveRequestsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DescribeReceiveRequests(
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

func (p Gs2GuildWebSocketClient) describeReceiveRequestsByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeReceiveRequestsByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReceiveRequestsByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReceiveRequestsByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeReceiveRequestsByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeReceiveRequestsByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DescribeReceiveRequestsByGuildNameAsync(
	request *DescribeReceiveRequestsByGuildNameRequest,
	callback chan<- DescribeReceiveRequestsByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "receiveMemberRequest",
			"function":    "describeReceiveRequestsByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.describeReceiveRequestsByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DescribeReceiveRequestsByGuildName(
	request *DescribeReceiveRequestsByGuildNameRequest,
) (*DescribeReceiveRequestsByGuildNameResult, error) {
	callback := make(chan DescribeReceiveRequestsByGuildNameAsyncResult, 1)
	go p.DescribeReceiveRequestsByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) getReceiveRequestAsyncHandler(
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

func (p Gs2GuildWebSocketClient) GetReceiveRequestAsync(
	request *GetReceiveRequestRequest,
	callback chan<- GetReceiveRequestAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "receiveMemberRequest",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

func (p Gs2GuildWebSocketClient) GetReceiveRequest(
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

func (p Gs2GuildWebSocketClient) getReceiveRequestByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetReceiveRequestByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReceiveRequestByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReceiveRequestByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetReceiveRequestByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetReceiveRequestByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) GetReceiveRequestByGuildNameAsync(
	request *GetReceiveRequestByGuildNameRequest,
	callback chan<- GetReceiveRequestByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "receiveMemberRequest",
			"function":    "getReceiveRequestByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.FromUserId != nil && *request.FromUserId != "" {
		bodies["fromUserId"] = *request.FromUserId
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

	go p.getReceiveRequestByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) GetReceiveRequestByGuildName(
	request *GetReceiveRequestByGuildNameRequest,
) (*GetReceiveRequestByGuildNameResult, error) {
	callback := make(chan GetReceiveRequestByGuildNameAsyncResult, 1)
	go p.GetReceiveRequestByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) acceptRequestAsyncHandler(
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
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "user.joinedGuild.tooMany" {
				asyncResult.Err = gs2err.SetClientError(MaximumJoinedGuildsReached{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.members.tooMany" {
				asyncResult.Err = gs2err.SetClientError(MaximumMembersReached{})
			}
		}
	}
	callback <- AcceptRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) AcceptRequestAsync(
	request *AcceptRequestRequest,
	callback chan<- AcceptRequestAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "receiveMemberRequest",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

func (p Gs2GuildWebSocketClient) AcceptRequest(
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

func (p Gs2GuildWebSocketClient) acceptRequestByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AcceptRequestByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcceptRequestByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcceptRequestByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcceptRequestByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "user.joinedGuild.tooMany" {
				asyncResult.Err = gs2err.SetClientError(MaximumJoinedGuildsReached{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.members.tooMany" {
				asyncResult.Err = gs2err.SetClientError(MaximumMembersReached{})
			}
		}
	}
	callback <- AcceptRequestByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) AcceptRequestByGuildNameAsync(
	request *AcceptRequestByGuildNameRequest,
	callback chan<- AcceptRequestByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "receiveMemberRequest",
			"function":    "acceptRequestByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.FromUserId != nil && *request.FromUserId != "" {
		bodies["fromUserId"] = *request.FromUserId
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

	go p.acceptRequestByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) AcceptRequestByGuildName(
	request *AcceptRequestByGuildNameRequest,
) (*AcceptRequestByGuildNameResult, error) {
	callback := make(chan AcceptRequestByGuildNameAsyncResult, 1)
	go p.AcceptRequestByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) rejectRequestAsyncHandler(
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

func (p Gs2GuildWebSocketClient) RejectRequestAsync(
	request *RejectRequestRequest,
	callback chan<- RejectRequestAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "receiveMemberRequest",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

func (p Gs2GuildWebSocketClient) RejectRequest(
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

func (p Gs2GuildWebSocketClient) rejectRequestByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RejectRequestByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RejectRequestByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RejectRequestByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RejectRequestByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- RejectRequestByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) RejectRequestByGuildNameAsync(
	request *RejectRequestByGuildNameRequest,
	callback chan<- RejectRequestByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "receiveMemberRequest",
			"function":    "rejectRequestByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
	}
	if request.FromUserId != nil && *request.FromUserId != "" {
		bodies["fromUserId"] = *request.FromUserId
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

	go p.rejectRequestByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) RejectRequestByGuildName(
	request *RejectRequestByGuildNameRequest,
) (*RejectRequestByGuildNameResult, error) {
	callback := make(chan RejectRequestByGuildNameAsyncResult, 1)
	go p.RejectRequestByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) describeSendRequestsAsyncHandler(
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

func (p Gs2GuildWebSocketClient) DescribeSendRequestsAsync(
	request *DescribeSendRequestsRequest,
	callback chan<- DescribeSendRequestsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "sendMemberRequest",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

func (p Gs2GuildWebSocketClient) DescribeSendRequests(
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

func (p Gs2GuildWebSocketClient) describeSendRequestsByUserIdAsyncHandler(
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

func (p Gs2GuildWebSocketClient) DescribeSendRequestsByUserIdAsync(
	request *DescribeSendRequestsByUserIdRequest,
	callback chan<- DescribeSendRequestsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "sendMemberRequest",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

func (p Gs2GuildWebSocketClient) DescribeSendRequestsByUserId(
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

func (p Gs2GuildWebSocketClient) getSendRequestAsyncHandler(
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

func (p Gs2GuildWebSocketClient) GetSendRequestAsync(
	request *GetSendRequestRequest,
	callback chan<- GetSendRequestAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "sendMemberRequest",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.TargetGuildName != nil && *request.TargetGuildName != "" {
		bodies["targetGuildName"] = *request.TargetGuildName
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

func (p Gs2GuildWebSocketClient) GetSendRequest(
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

func (p Gs2GuildWebSocketClient) getSendRequestByUserIdAsyncHandler(
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

func (p Gs2GuildWebSocketClient) GetSendRequestByUserIdAsync(
	request *GetSendRequestByUserIdRequest,
	callback chan<- GetSendRequestByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "sendMemberRequest",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.TargetGuildName != nil && *request.TargetGuildName != "" {
		bodies["targetGuildName"] = *request.TargetGuildName
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

func (p Gs2GuildWebSocketClient) GetSendRequestByUserId(
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

func (p Gs2GuildWebSocketClient) sendRequestAsyncHandler(
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
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.members.tooMany" {
				asyncResult.Err = gs2err.SetClientError(MaximumMembersReached{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "user.joinedGuild.tooMany" {
				asyncResult.Err = gs2err.SetClientError(MaximumJoinedGuildsReached{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.receiveRequests.tooMany" {
				asyncResult.Err = gs2err.SetClientError(MaximumReceiveRequestsReached{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.sendRequests.tooMany" {
				asyncResult.Err = gs2err.SetClientError(MaximumSendRequestsReached{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.sendRequests.notMeetJoinRequirements" {
				asyncResult.Err = gs2err.SetClientError(DotMeetJoinRequirements{})
			}
		}
	}
	callback <- SendRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) SendRequestAsync(
	request *SendRequestRequest,
	callback chan<- SendRequestAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "sendMemberRequest",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.TargetGuildName != nil && *request.TargetGuildName != "" {
		bodies["targetGuildName"] = *request.TargetGuildName
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
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

func (p Gs2GuildWebSocketClient) SendRequest(
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

func (p Gs2GuildWebSocketClient) sendRequestByUserIdAsyncHandler(
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
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.members.tooMany" {
				asyncResult.Err = gs2err.SetClientError(MaximumMembersReached{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "user.joinedGuild.tooMany" {
				asyncResult.Err = gs2err.SetClientError(MaximumJoinedGuildsReached{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.receiveRequests.tooMany" {
				asyncResult.Err = gs2err.SetClientError(MaximumReceiveRequestsReached{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.sendRequests.tooMany" {
				asyncResult.Err = gs2err.SetClientError(MaximumSendRequestsReached{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "guild.sendRequests.notMeetJoinRequirements" {
				asyncResult.Err = gs2err.SetClientError(DotMeetJoinRequirements{})
			}
		}
	}
	callback <- SendRequestByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) SendRequestByUserIdAsync(
	request *SendRequestByUserIdRequest,
	callback chan<- SendRequestByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "sendMemberRequest",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.TargetGuildName != nil && *request.TargetGuildName != "" {
		bodies["targetGuildName"] = *request.TargetGuildName
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
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

func (p Gs2GuildWebSocketClient) SendRequestByUserId(
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

func (p Gs2GuildWebSocketClient) deleteRequestAsyncHandler(
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

func (p Gs2GuildWebSocketClient) DeleteRequestAsync(
	request *DeleteRequestRequest,
	callback chan<- DeleteRequestAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "sendMemberRequest",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.TargetGuildName != nil && *request.TargetGuildName != "" {
		bodies["targetGuildName"] = *request.TargetGuildName
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

func (p Gs2GuildWebSocketClient) DeleteRequest(
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

func (p Gs2GuildWebSocketClient) deleteRequestByUserIdAsyncHandler(
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

func (p Gs2GuildWebSocketClient) DeleteRequestByUserIdAsync(
	request *DeleteRequestByUserIdRequest,
	callback chan<- DeleteRequestByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "sendMemberRequest",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.TargetGuildName != nil && *request.TargetGuildName != "" {
		bodies["targetGuildName"] = *request.TargetGuildName
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

func (p Gs2GuildWebSocketClient) DeleteRequestByUserId(
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

func (p Gs2GuildWebSocketClient) describeIgnoreUsersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeIgnoreUsersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeIgnoreUsersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeIgnoreUsersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeIgnoreUsersAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeIgnoreUsersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DescribeIgnoreUsersAsync(
	request *DescribeIgnoreUsersRequest,
	callback chan<- DescribeIgnoreUsersAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "ignoreUser",
			"function":    "describeIgnoreUsers",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
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

	go p.describeIgnoreUsersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DescribeIgnoreUsers(
	request *DescribeIgnoreUsersRequest,
) (*DescribeIgnoreUsersResult, error) {
	callback := make(chan DescribeIgnoreUsersAsyncResult, 1)
	go p.DescribeIgnoreUsersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) describeIgnoreUsersByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeIgnoreUsersByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeIgnoreUsersByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeIgnoreUsersByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeIgnoreUsersByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeIgnoreUsersByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DescribeIgnoreUsersByGuildNameAsync(
	request *DescribeIgnoreUsersByGuildNameRequest,
	callback chan<- DescribeIgnoreUsersByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "ignoreUser",
			"function":    "describeIgnoreUsersByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.describeIgnoreUsersByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DescribeIgnoreUsersByGuildName(
	request *DescribeIgnoreUsersByGuildNameRequest,
) (*DescribeIgnoreUsersByGuildNameResult, error) {
	callback := make(chan DescribeIgnoreUsersByGuildNameAsyncResult, 1)
	go p.DescribeIgnoreUsersByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) getIgnoreUserAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetIgnoreUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetIgnoreUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetIgnoreUserResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetIgnoreUserAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetIgnoreUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) GetIgnoreUserAsync(
	request *GetIgnoreUserRequest,
	callback chan<- GetIgnoreUserAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "ignoreUser",
			"function":    "getIgnoreUser",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
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

	go p.getIgnoreUserAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) GetIgnoreUser(
	request *GetIgnoreUserRequest,
) (*GetIgnoreUserResult, error) {
	callback := make(chan GetIgnoreUserAsyncResult, 1)
	go p.GetIgnoreUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) getIgnoreUserByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetIgnoreUserByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetIgnoreUserByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetIgnoreUserByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetIgnoreUserByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetIgnoreUserByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) GetIgnoreUserByGuildNameAsync(
	request *GetIgnoreUserByGuildNameRequest,
	callback chan<- GetIgnoreUserByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "ignoreUser",
			"function":    "getIgnoreUserByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.getIgnoreUserByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) GetIgnoreUserByGuildName(
	request *GetIgnoreUserByGuildNameRequest,
) (*GetIgnoreUserByGuildNameResult, error) {
	callback := make(chan GetIgnoreUserByGuildNameAsyncResult, 1)
	go p.GetIgnoreUserByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) addIgnoreUserAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddIgnoreUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddIgnoreUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddIgnoreUserResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddIgnoreUserAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- AddIgnoreUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) AddIgnoreUserAsync(
	request *AddIgnoreUserRequest,
	callback chan<- AddIgnoreUserAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "ignoreUser",
			"function":    "addIgnoreUser",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
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

	go p.addIgnoreUserAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) AddIgnoreUser(
	request *AddIgnoreUserRequest,
) (*AddIgnoreUserResult, error) {
	callback := make(chan AddIgnoreUserAsyncResult, 1)
	go p.AddIgnoreUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) addIgnoreUserByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddIgnoreUserByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddIgnoreUserByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddIgnoreUserByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddIgnoreUserByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- AddIgnoreUserByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) AddIgnoreUserByGuildNameAsync(
	request *AddIgnoreUserByGuildNameRequest,
	callback chan<- AddIgnoreUserByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "ignoreUser",
			"function":    "addIgnoreUserByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.addIgnoreUserByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) AddIgnoreUserByGuildName(
	request *AddIgnoreUserByGuildNameRequest,
) (*AddIgnoreUserByGuildNameResult, error) {
	callback := make(chan AddIgnoreUserByGuildNameAsyncResult, 1)
	go p.AddIgnoreUserByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) deleteIgnoreUserAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteIgnoreUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteIgnoreUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteIgnoreUserResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteIgnoreUserAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteIgnoreUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DeleteIgnoreUserAsync(
	request *DeleteIgnoreUserRequest,
	callback chan<- DeleteIgnoreUserAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "ignoreUser",
			"function":    "deleteIgnoreUser",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
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

	go p.deleteIgnoreUserAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DeleteIgnoreUser(
	request *DeleteIgnoreUserRequest,
) (*DeleteIgnoreUserResult, error) {
	callback := make(chan DeleteIgnoreUserAsyncResult, 1)
	go p.DeleteIgnoreUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GuildWebSocketClient) deleteIgnoreUserByGuildNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteIgnoreUserByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteIgnoreUserByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteIgnoreUserByGuildNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteIgnoreUserByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteIgnoreUserByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildWebSocketClient) DeleteIgnoreUserByGuildNameAsync(
	request *DeleteIgnoreUserByGuildNameRequest,
	callback chan<- DeleteIgnoreUserByGuildNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "guild",
			"component":   "ignoreUser",
			"function":    "deleteIgnoreUserByGuildName",
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		bodies["guildModelName"] = *request.GuildModelName
	}
	if request.GuildName != nil && *request.GuildName != "" {
		bodies["guildName"] = *request.GuildName
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

	go p.deleteIgnoreUserByGuildNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GuildWebSocketClient) DeleteIgnoreUserByGuildName(
	request *DeleteIgnoreUserByGuildNameRequest,
) (*DeleteIgnoreUserByGuildNameResult, error) {
	callback := make(chan DeleteIgnoreUserByGuildNameAsyncResult, 1)
	go p.DeleteIgnoreUserByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
