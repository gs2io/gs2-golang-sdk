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

package matchmaking

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2MatchmakingWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2MatchmakingWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2MatchmakingWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2MatchmakingWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
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

	go p.describeNamespacesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DescribeNamespaces(
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

func (p Gs2MatchmakingWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2MatchmakingWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
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
	if request.EnableRating != nil {
		bodies["enableRating"] = *request.EnableRating
	}
	if request.CreateGatheringTriggerType != nil && *request.CreateGatheringTriggerType != "" {
		bodies["createGatheringTriggerType"] = *request.CreateGatheringTriggerType
	}
	if request.CreateGatheringTriggerRealtimeNamespaceId != nil && *request.CreateGatheringTriggerRealtimeNamespaceId != "" {
		bodies["createGatheringTriggerRealtimeNamespaceId"] = *request.CreateGatheringTriggerRealtimeNamespaceId
	}
	if request.CreateGatheringTriggerScriptId != nil && *request.CreateGatheringTriggerScriptId != "" {
		bodies["createGatheringTriggerScriptId"] = *request.CreateGatheringTriggerScriptId
	}
	if request.CompleteMatchmakingTriggerType != nil && *request.CompleteMatchmakingTriggerType != "" {
		bodies["completeMatchmakingTriggerType"] = *request.CompleteMatchmakingTriggerType
	}
	if request.CompleteMatchmakingTriggerRealtimeNamespaceId != nil && *request.CompleteMatchmakingTriggerRealtimeNamespaceId != "" {
		bodies["completeMatchmakingTriggerRealtimeNamespaceId"] = *request.CompleteMatchmakingTriggerRealtimeNamespaceId
	}
	if request.CompleteMatchmakingTriggerScriptId != nil && *request.CompleteMatchmakingTriggerScriptId != "" {
		bodies["completeMatchmakingTriggerScriptId"] = *request.CompleteMatchmakingTriggerScriptId
	}
	if request.JoinNotification != nil {
		bodies["joinNotification"] = request.JoinNotification.ToDict()
	}
	if request.LeaveNotification != nil {
		bodies["leaveNotification"] = request.LeaveNotification.ToDict()
	}
	if request.CompleteNotification != nil {
		bodies["completeNotification"] = request.CompleteNotification.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.createNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) CreateNamespace(
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

func (p Gs2MatchmakingWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2MatchmakingWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
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

	go p.getNamespaceStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) GetNamespaceStatus(
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

func (p Gs2MatchmakingWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2MatchmakingWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
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

	go p.getNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) GetNamespace(
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

func (p Gs2MatchmakingWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2MatchmakingWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
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
	if request.EnableRating != nil {
		bodies["enableRating"] = *request.EnableRating
	}
	if request.CreateGatheringTriggerType != nil && *request.CreateGatheringTriggerType != "" {
		bodies["createGatheringTriggerType"] = *request.CreateGatheringTriggerType
	}
	if request.CreateGatheringTriggerRealtimeNamespaceId != nil && *request.CreateGatheringTriggerRealtimeNamespaceId != "" {
		bodies["createGatheringTriggerRealtimeNamespaceId"] = *request.CreateGatheringTriggerRealtimeNamespaceId
	}
	if request.CreateGatheringTriggerScriptId != nil && *request.CreateGatheringTriggerScriptId != "" {
		bodies["createGatheringTriggerScriptId"] = *request.CreateGatheringTriggerScriptId
	}
	if request.CompleteMatchmakingTriggerType != nil && *request.CompleteMatchmakingTriggerType != "" {
		bodies["completeMatchmakingTriggerType"] = *request.CompleteMatchmakingTriggerType
	}
	if request.CompleteMatchmakingTriggerRealtimeNamespaceId != nil && *request.CompleteMatchmakingTriggerRealtimeNamespaceId != "" {
		bodies["completeMatchmakingTriggerRealtimeNamespaceId"] = *request.CompleteMatchmakingTriggerRealtimeNamespaceId
	}
	if request.CompleteMatchmakingTriggerScriptId != nil && *request.CompleteMatchmakingTriggerScriptId != "" {
		bodies["completeMatchmakingTriggerScriptId"] = *request.CompleteMatchmakingTriggerScriptId
	}
	if request.JoinNotification != nil {
		bodies["joinNotification"] = request.JoinNotification.ToDict()
	}
	if request.LeaveNotification != nil {
		bodies["leaveNotification"] = request.LeaveNotification.ToDict()
	}
	if request.CompleteNotification != nil {
		bodies["completeNotification"] = request.CompleteNotification.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) UpdateNamespace(
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

func (p Gs2MatchmakingWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2MatchmakingWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
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

	go p.deleteNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DeleteNamespace(
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

func (p Gs2MatchmakingWebSocketClient) describeGatheringsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGatheringsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGatheringsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGatheringsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGatheringsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGatheringsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) DescribeGatheringsAsync(
	request *DescribeGatheringsRequest,
	callback chan<- DescribeGatheringsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "gathering",
			"function":    "describeGatherings",
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

	go p.describeGatheringsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DescribeGatherings(
	request *DescribeGatheringsRequest,
) (*DescribeGatheringsResult, error) {
	callback := make(chan DescribeGatheringsAsyncResult, 1)
	go p.DescribeGatheringsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) createGatheringAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateGatheringAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGatheringAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGatheringResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGatheringAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateGatheringAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) CreateGatheringAsync(
	request *CreateGatheringRequest,
	callback chan<- CreateGatheringAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "gathering",
			"function":    "createGathering",
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
	if request.Player != nil {
		bodies["player"] = request.Player.ToDict()
	}
	if request.AttributeRanges != nil {
		var _attributeRanges []interface{}
		for _, item := range request.AttributeRanges {
			_attributeRanges = append(_attributeRanges, item)
		}
		bodies["attributeRanges"] = _attributeRanges
	}
	if request.CapacityOfRoles != nil {
		var _capacityOfRoles []interface{}
		for _, item := range request.CapacityOfRoles {
			_capacityOfRoles = append(_capacityOfRoles, item)
		}
		bodies["capacityOfRoles"] = _capacityOfRoles
	}
	if request.AllowUserIds != nil {
		var _allowUserIds []interface{}
		for _, item := range request.AllowUserIds {
			_allowUserIds = append(_allowUserIds, item)
		}
		bodies["allowUserIds"] = _allowUserIds
	}
	if request.ExpiresAt != nil {
		bodies["expiresAt"] = *request.ExpiresAt
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.createGatheringAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) CreateGathering(
	request *CreateGatheringRequest,
) (*CreateGatheringResult, error) {
	callback := make(chan CreateGatheringAsyncResult, 1)
	go p.CreateGatheringAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) createGatheringByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateGatheringByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGatheringByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGatheringByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGatheringByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateGatheringByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) CreateGatheringByUserIdAsync(
	request *CreateGatheringByUserIdRequest,
	callback chan<- CreateGatheringByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "gathering",
			"function":    "createGatheringByUserId",
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
	if request.Player != nil {
		bodies["player"] = request.Player.ToDict()
	}
	if request.AttributeRanges != nil {
		var _attributeRanges []interface{}
		for _, item := range request.AttributeRanges {
			_attributeRanges = append(_attributeRanges, item)
		}
		bodies["attributeRanges"] = _attributeRanges
	}
	if request.CapacityOfRoles != nil {
		var _capacityOfRoles []interface{}
		for _, item := range request.CapacityOfRoles {
			_capacityOfRoles = append(_capacityOfRoles, item)
		}
		bodies["capacityOfRoles"] = _capacityOfRoles
	}
	if request.AllowUserIds != nil {
		var _allowUserIds []interface{}
		for _, item := range request.AllowUserIds {
			_allowUserIds = append(_allowUserIds, item)
		}
		bodies["allowUserIds"] = _allowUserIds
	}
	if request.ExpiresAt != nil {
		bodies["expiresAt"] = *request.ExpiresAt
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.createGatheringByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) CreateGatheringByUserId(
	request *CreateGatheringByUserIdRequest,
) (*CreateGatheringByUserIdResult, error) {
	callback := make(chan CreateGatheringByUserIdAsyncResult, 1)
	go p.CreateGatheringByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) updateGatheringAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateGatheringAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGatheringAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGatheringResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateGatheringAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateGatheringAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) UpdateGatheringAsync(
	request *UpdateGatheringRequest,
	callback chan<- UpdateGatheringAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "gathering",
			"function":    "updateGathering",
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
	if request.GatheringName != nil && *request.GatheringName != "" {
		bodies["gatheringName"] = *request.GatheringName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.AttributeRanges != nil {
		var _attributeRanges []interface{}
		for _, item := range request.AttributeRanges {
			_attributeRanges = append(_attributeRanges, item)
		}
		bodies["attributeRanges"] = _attributeRanges
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.updateGatheringAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) UpdateGathering(
	request *UpdateGatheringRequest,
) (*UpdateGatheringResult, error) {
	callback := make(chan UpdateGatheringAsyncResult, 1)
	go p.UpdateGatheringAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) updateGatheringByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateGatheringByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGatheringByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGatheringByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateGatheringByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateGatheringByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) UpdateGatheringByUserIdAsync(
	request *UpdateGatheringByUserIdRequest,
	callback chan<- UpdateGatheringByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "gathering",
			"function":    "updateGatheringByUserId",
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
	if request.GatheringName != nil && *request.GatheringName != "" {
		bodies["gatheringName"] = *request.GatheringName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.AttributeRanges != nil {
		var _attributeRanges []interface{}
		for _, item := range request.AttributeRanges {
			_attributeRanges = append(_attributeRanges, item)
		}
		bodies["attributeRanges"] = _attributeRanges
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateGatheringByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) UpdateGatheringByUserId(
	request *UpdateGatheringByUserIdRequest,
) (*UpdateGatheringByUserIdResult, error) {
	callback := make(chan UpdateGatheringByUserIdAsyncResult, 1)
	go p.UpdateGatheringByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) doMatchmakingByPlayerAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DoMatchmakingByPlayerAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoMatchmakingByPlayerAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoMatchmakingByPlayerResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DoMatchmakingByPlayerAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DoMatchmakingByPlayerAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) DoMatchmakingByPlayerAsync(
	request *DoMatchmakingByPlayerRequest,
	callback chan<- DoMatchmakingByPlayerAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "gathering",
			"function":    "doMatchmakingByPlayer",
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
	if request.Player != nil {
		bodies["player"] = request.Player.ToDict()
	}
	if request.MatchmakingContextToken != nil && *request.MatchmakingContextToken != "" {
		bodies["matchmakingContextToken"] = *request.MatchmakingContextToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.doMatchmakingByPlayerAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DoMatchmakingByPlayer(
	request *DoMatchmakingByPlayerRequest,
) (*DoMatchmakingByPlayerResult, error) {
	callback := make(chan DoMatchmakingByPlayerAsyncResult, 1)
	go p.DoMatchmakingByPlayerAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) doMatchmakingAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DoMatchmakingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoMatchmakingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoMatchmakingResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DoMatchmakingAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DoMatchmakingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) DoMatchmakingAsync(
	request *DoMatchmakingRequest,
	callback chan<- DoMatchmakingAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "gathering",
			"function":    "doMatchmaking",
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
	if request.Player != nil {
		bodies["player"] = request.Player.ToDict()
	}
	if request.MatchmakingContextToken != nil && *request.MatchmakingContextToken != "" {
		bodies["matchmakingContextToken"] = *request.MatchmakingContextToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.doMatchmakingAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DoMatchmaking(
	request *DoMatchmakingRequest,
) (*DoMatchmakingResult, error) {
	callback := make(chan DoMatchmakingAsyncResult, 1)
	go p.DoMatchmakingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) doMatchmakingByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DoMatchmakingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoMatchmakingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoMatchmakingByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DoMatchmakingByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DoMatchmakingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) DoMatchmakingByUserIdAsync(
	request *DoMatchmakingByUserIdRequest,
	callback chan<- DoMatchmakingByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "gathering",
			"function":    "doMatchmakingByUserId",
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
	if request.Player != nil {
		bodies["player"] = request.Player.ToDict()
	}
	if request.MatchmakingContextToken != nil && *request.MatchmakingContextToken != "" {
		bodies["matchmakingContextToken"] = *request.MatchmakingContextToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.doMatchmakingByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DoMatchmakingByUserId(
	request *DoMatchmakingByUserIdRequest,
) (*DoMatchmakingByUserIdResult, error) {
	callback := make(chan DoMatchmakingByUserIdAsyncResult, 1)
	go p.DoMatchmakingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) getGatheringAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGatheringAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGatheringAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGatheringResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGatheringAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGatheringAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) GetGatheringAsync(
	request *GetGatheringRequest,
	callback chan<- GetGatheringAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "gathering",
			"function":    "getGathering",
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
	if request.GatheringName != nil && *request.GatheringName != "" {
		bodies["gatheringName"] = *request.GatheringName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getGatheringAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) GetGathering(
	request *GetGatheringRequest,
) (*GetGatheringResult, error) {
	callback := make(chan GetGatheringAsyncResult, 1)
	go p.GetGatheringAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) cancelMatchmakingAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CancelMatchmakingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CancelMatchmakingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CancelMatchmakingResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CancelMatchmakingAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CancelMatchmakingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) CancelMatchmakingAsync(
	request *CancelMatchmakingRequest,
	callback chan<- CancelMatchmakingAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "gathering",
			"function":    "cancelMatchmaking",
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
	if request.GatheringName != nil && *request.GatheringName != "" {
		bodies["gatheringName"] = *request.GatheringName
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

	go p.cancelMatchmakingAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) CancelMatchmaking(
	request *CancelMatchmakingRequest,
) (*CancelMatchmakingResult, error) {
	callback := make(chan CancelMatchmakingAsyncResult, 1)
	go p.CancelMatchmakingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) cancelMatchmakingByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CancelMatchmakingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CancelMatchmakingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CancelMatchmakingByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CancelMatchmakingByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CancelMatchmakingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) CancelMatchmakingByUserIdAsync(
	request *CancelMatchmakingByUserIdRequest,
	callback chan<- CancelMatchmakingByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "gathering",
			"function":    "cancelMatchmakingByUserId",
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
	if request.GatheringName != nil && *request.GatheringName != "" {
		bodies["gatheringName"] = *request.GatheringName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.cancelMatchmakingByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) CancelMatchmakingByUserId(
	request *CancelMatchmakingByUserIdRequest,
) (*CancelMatchmakingByUserIdResult, error) {
	callback := make(chan CancelMatchmakingByUserIdAsyncResult, 1)
	go p.CancelMatchmakingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) deleteGatheringAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteGatheringAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGatheringAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGatheringResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGatheringAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteGatheringAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) DeleteGatheringAsync(
	request *DeleteGatheringRequest,
	callback chan<- DeleteGatheringAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "gathering",
			"function":    "deleteGathering",
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
	if request.GatheringName != nil && *request.GatheringName != "" {
		bodies["gatheringName"] = *request.GatheringName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteGatheringAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DeleteGathering(
	request *DeleteGatheringRequest,
) (*DeleteGatheringResult, error) {
	callback := make(chan DeleteGatheringAsyncResult, 1)
	go p.DeleteGatheringAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) describeRatingModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRatingModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRatingModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRatingModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRatingModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRatingModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) DescribeRatingModelMastersAsync(
	request *DescribeRatingModelMastersRequest,
	callback chan<- DescribeRatingModelMastersAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "ratingModelMaster",
			"function":    "describeRatingModelMasters",
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

	go p.describeRatingModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DescribeRatingModelMasters(
	request *DescribeRatingModelMastersRequest,
) (*DescribeRatingModelMastersResult, error) {
	callback := make(chan DescribeRatingModelMastersAsyncResult, 1)
	go p.DescribeRatingModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) createRatingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateRatingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateRatingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateRatingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateRatingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateRatingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) CreateRatingModelMasterAsync(
	request *CreateRatingModelMasterRequest,
	callback chan<- CreateRatingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "ratingModelMaster",
			"function":    "createRatingModelMaster",
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
	if request.Volatility != nil {
		bodies["volatility"] = *request.Volatility
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.createRatingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) CreateRatingModelMaster(
	request *CreateRatingModelMasterRequest,
) (*CreateRatingModelMasterResult, error) {
	callback := make(chan CreateRatingModelMasterAsyncResult, 1)
	go p.CreateRatingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) getRatingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRatingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRatingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRatingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRatingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRatingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) GetRatingModelMasterAsync(
	request *GetRatingModelMasterRequest,
	callback chan<- GetRatingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "ratingModelMaster",
			"function":    "getRatingModelMaster",
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
	if request.RatingName != nil && *request.RatingName != "" {
		bodies["ratingName"] = *request.RatingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getRatingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) GetRatingModelMaster(
	request *GetRatingModelMasterRequest,
) (*GetRatingModelMasterResult, error) {
	callback := make(chan GetRatingModelMasterAsyncResult, 1)
	go p.GetRatingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) updateRatingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateRatingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateRatingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateRatingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateRatingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateRatingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) UpdateRatingModelMasterAsync(
	request *UpdateRatingModelMasterRequest,
	callback chan<- UpdateRatingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "ratingModelMaster",
			"function":    "updateRatingModelMaster",
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
	if request.RatingName != nil && *request.RatingName != "" {
		bodies["ratingName"] = *request.RatingName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Volatility != nil {
		bodies["volatility"] = *request.Volatility
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateRatingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) UpdateRatingModelMaster(
	request *UpdateRatingModelMasterRequest,
) (*UpdateRatingModelMasterResult, error) {
	callback := make(chan UpdateRatingModelMasterAsyncResult, 1)
	go p.UpdateRatingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) deleteRatingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteRatingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRatingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRatingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteRatingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteRatingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) DeleteRatingModelMasterAsync(
	request *DeleteRatingModelMasterRequest,
	callback chan<- DeleteRatingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "ratingModelMaster",
			"function":    "deleteRatingModelMaster",
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
	if request.RatingName != nil && *request.RatingName != "" {
		bodies["ratingName"] = *request.RatingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteRatingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DeleteRatingModelMaster(
	request *DeleteRatingModelMasterRequest,
) (*DeleteRatingModelMasterResult, error) {
	callback := make(chan DeleteRatingModelMasterAsyncResult, 1)
	go p.DeleteRatingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) describeRatingModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRatingModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRatingModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRatingModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRatingModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRatingModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) DescribeRatingModelsAsync(
	request *DescribeRatingModelsRequest,
	callback chan<- DescribeRatingModelsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "ratingModel",
			"function":    "describeRatingModels",
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

	go p.describeRatingModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DescribeRatingModels(
	request *DescribeRatingModelsRequest,
) (*DescribeRatingModelsResult, error) {
	callback := make(chan DescribeRatingModelsAsyncResult, 1)
	go p.DescribeRatingModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) getRatingModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRatingModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRatingModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRatingModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRatingModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRatingModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) GetRatingModelAsync(
	request *GetRatingModelRequest,
	callback chan<- GetRatingModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "ratingModel",
			"function":    "getRatingModel",
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
	if request.RatingName != nil && *request.RatingName != "" {
		bodies["ratingName"] = *request.RatingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getRatingModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) GetRatingModel(
	request *GetRatingModelRequest,
) (*GetRatingModelResult, error) {
	callback := make(chan GetRatingModelAsyncResult, 1)
	go p.GetRatingModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) exportMasterAsyncHandler(
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

func (p Gs2MatchmakingWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "currentRatingModelMaster",
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

	go p.exportMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) ExportMaster(
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

func (p Gs2MatchmakingWebSocketClient) getCurrentRatingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentRatingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentRatingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentRatingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentRatingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentRatingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) GetCurrentRatingModelMasterAsync(
	request *GetCurrentRatingModelMasterRequest,
	callback chan<- GetCurrentRatingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "currentRatingModelMaster",
			"function":    "getCurrentRatingModelMaster",
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

	go p.getCurrentRatingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) GetCurrentRatingModelMaster(
	request *GetCurrentRatingModelMasterRequest,
) (*GetCurrentRatingModelMasterResult, error) {
	callback := make(chan GetCurrentRatingModelMasterAsyncResult, 1)
	go p.GetCurrentRatingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) updateCurrentRatingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentRatingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentRatingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentRatingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentRatingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentRatingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) UpdateCurrentRatingModelMasterAsync(
	request *UpdateCurrentRatingModelMasterRequest,
	callback chan<- UpdateCurrentRatingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "currentRatingModelMaster",
			"function":    "updateCurrentRatingModelMaster",
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
	if request.Settings != nil && *request.Settings != "" {
		bodies["settings"] = *request.Settings
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateCurrentRatingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) UpdateCurrentRatingModelMaster(
	request *UpdateCurrentRatingModelMasterRequest,
) (*UpdateCurrentRatingModelMasterResult, error) {
	callback := make(chan UpdateCurrentRatingModelMasterAsyncResult, 1)
	go p.UpdateCurrentRatingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) updateCurrentRatingModelMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentRatingModelMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentRatingModelMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentRatingModelMasterFromGitHubResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentRatingModelMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentRatingModelMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) UpdateCurrentRatingModelMasterFromGitHubAsync(
	request *UpdateCurrentRatingModelMasterFromGitHubRequest,
	callback chan<- UpdateCurrentRatingModelMasterFromGitHubAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "currentRatingModelMaster",
			"function":    "updateCurrentRatingModelMasterFromGitHub",
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

	go p.updateCurrentRatingModelMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) UpdateCurrentRatingModelMasterFromGitHub(
	request *UpdateCurrentRatingModelMasterFromGitHubRequest,
) (*UpdateCurrentRatingModelMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentRatingModelMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentRatingModelMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) describeRatingsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRatingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRatingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRatingsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRatingsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRatingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) DescribeRatingsAsync(
	request *DescribeRatingsRequest,
	callback chan<- DescribeRatingsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "rating",
			"function":    "describeRatings",
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

	go p.describeRatingsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DescribeRatings(
	request *DescribeRatingsRequest,
) (*DescribeRatingsResult, error) {
	callback := make(chan DescribeRatingsAsyncResult, 1)
	go p.DescribeRatingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) describeRatingsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRatingsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRatingsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRatingsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRatingsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRatingsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) DescribeRatingsByUserIdAsync(
	request *DescribeRatingsByUserIdRequest,
	callback chan<- DescribeRatingsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "rating",
			"function":    "describeRatingsByUserId",
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
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.describeRatingsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DescribeRatingsByUserId(
	request *DescribeRatingsByUserIdRequest,
) (*DescribeRatingsByUserIdResult, error) {
	callback := make(chan DescribeRatingsByUserIdAsyncResult, 1)
	go p.DescribeRatingsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) getRatingAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRatingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRatingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRatingResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRatingAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRatingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) GetRatingAsync(
	request *GetRatingRequest,
	callback chan<- GetRatingAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "rating",
			"function":    "getRating",
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
	if request.RatingName != nil && *request.RatingName != "" {
		bodies["ratingName"] = *request.RatingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getRatingAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) GetRating(
	request *GetRatingRequest,
) (*GetRatingResult, error) {
	callback := make(chan GetRatingAsyncResult, 1)
	go p.GetRatingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) getRatingByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRatingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRatingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRatingByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRatingByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRatingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) GetRatingByUserIdAsync(
	request *GetRatingByUserIdRequest,
	callback chan<- GetRatingByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "rating",
			"function":    "getRatingByUserId",
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
	if request.RatingName != nil && *request.RatingName != "" {
		bodies["ratingName"] = *request.RatingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getRatingByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) GetRatingByUserId(
	request *GetRatingByUserIdRequest,
) (*GetRatingByUserIdResult, error) {
	callback := make(chan GetRatingByUserIdAsyncResult, 1)
	go p.GetRatingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) putResultAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PutResultAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutResultAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutResultResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutResultAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PutResultAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) PutResultAsync(
	request *PutResultRequest,
	callback chan<- PutResultAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "rating",
			"function":    "putResult",
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
	if request.RatingName != nil && *request.RatingName != "" {
		bodies["ratingName"] = *request.RatingName
	}
	if request.GameResults != nil {
		var _gameResults []interface{}
		for _, item := range request.GameResults {
			_gameResults = append(_gameResults, item)
		}
		bodies["gameResults"] = _gameResults
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.putResultAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) PutResult(
	request *PutResultRequest,
) (*PutResultResult, error) {
	callback := make(chan PutResultAsyncResult, 1)
	go p.PutResultAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) deleteRatingAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteRatingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRatingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRatingResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteRatingAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteRatingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) DeleteRatingAsync(
	request *DeleteRatingRequest,
	callback chan<- DeleteRatingAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "rating",
			"function":    "deleteRating",
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
	if request.RatingName != nil && *request.RatingName != "" {
		bodies["ratingName"] = *request.RatingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteRatingAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) DeleteRating(
	request *DeleteRatingRequest,
) (*DeleteRatingResult, error) {
	callback := make(chan DeleteRatingAsyncResult, 1)
	go p.DeleteRatingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) getBallotAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetBallotAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBallotAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBallotResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBallotAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBallotAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) GetBallotAsync(
	request *GetBallotRequest,
	callback chan<- GetBallotAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "vote",
			"function":    "getBallot",
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
	if request.RatingName != nil && *request.RatingName != "" {
		bodies["ratingName"] = *request.RatingName
	}
	if request.GatheringName != nil && *request.GatheringName != "" {
		bodies["gatheringName"] = *request.GatheringName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.NumberOfPlayer != nil {
		bodies["numberOfPlayer"] = *request.NumberOfPlayer
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getBallotAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) GetBallot(
	request *GetBallotRequest,
) (*GetBallotResult, error) {
	callback := make(chan GetBallotAsyncResult, 1)
	go p.GetBallotAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) getBallotByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetBallotByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBallotByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBallotByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBallotByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBallotByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) GetBallotByUserIdAsync(
	request *GetBallotByUserIdRequest,
	callback chan<- GetBallotByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "vote",
			"function":    "getBallotByUserId",
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
	if request.RatingName != nil && *request.RatingName != "" {
		bodies["ratingName"] = *request.RatingName
	}
	if request.GatheringName != nil && *request.GatheringName != "" {
		bodies["gatheringName"] = *request.GatheringName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.NumberOfPlayer != nil {
		bodies["numberOfPlayer"] = *request.NumberOfPlayer
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getBallotByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) GetBallotByUserId(
	request *GetBallotByUserIdRequest,
) (*GetBallotByUserIdResult, error) {
	callback := make(chan GetBallotByUserIdAsyncResult, 1)
	go p.GetBallotByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) voteAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VoteAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VoteAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VoteResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VoteAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VoteAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) VoteAsync(
	request *VoteRequest,
	callback chan<- VoteAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "vote",
			"function":    "vote",
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
	if request.BallotBody != nil && *request.BallotBody != "" {
		bodies["ballotBody"] = *request.BallotBody
	}
	if request.BallotSignature != nil && *request.BallotSignature != "" {
		bodies["ballotSignature"] = *request.BallotSignature
	}
	if request.GameResults != nil {
		var _gameResults []interface{}
		for _, item := range request.GameResults {
			_gameResults = append(_gameResults, item)
		}
		bodies["gameResults"] = _gameResults
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.voteAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) Vote(
	request *VoteRequest,
) (*VoteResult, error) {
	callback := make(chan VoteAsyncResult, 1)
	go p.VoteAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) voteMultipleAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VoteMultipleAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VoteMultipleAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VoteMultipleResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VoteMultipleAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VoteMultipleAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) VoteMultipleAsync(
	request *VoteMultipleRequest,
	callback chan<- VoteMultipleAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "vote",
			"function":    "voteMultiple",
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
	if request.SignedBallots != nil {
		var _signedBallots []interface{}
		for _, item := range request.SignedBallots {
			_signedBallots = append(_signedBallots, item)
		}
		bodies["signedBallots"] = _signedBallots
	}
	if request.GameResults != nil {
		var _gameResults []interface{}
		for _, item := range request.GameResults {
			_gameResults = append(_gameResults, item)
		}
		bodies["gameResults"] = _gameResults
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.voteMultipleAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) VoteMultiple(
	request *VoteMultipleRequest,
) (*VoteMultipleResult, error) {
	callback := make(chan VoteMultipleAsyncResult, 1)
	go p.VoteMultipleAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MatchmakingWebSocketClient) commitVoteAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CommitVoteAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CommitVoteAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CommitVoteResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CommitVoteAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CommitVoteAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingWebSocketClient) CommitVoteAsync(
	request *CommitVoteRequest,
	callback chan<- CommitVoteAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "matchmaking",
			"component":   "vote",
			"function":    "commitVote",
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
	if request.RatingName != nil && *request.RatingName != "" {
		bodies["ratingName"] = *request.RatingName
	}
	if request.GatheringName != nil && *request.GatheringName != "" {
		bodies["gatheringName"] = *request.GatheringName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.commitVoteAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingWebSocketClient) CommitVote(
	request *CommitVoteRequest,
) (*CommitVoteResult, error) {
	callback := make(chan CommitVoteAsyncResult, 1)
	go p.CommitVoteAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
