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

package mission

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2MissionWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2MissionWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2MissionWebSocketClient) describeCompletesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeCompletesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCompletesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCompletesResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCompletesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeCompletesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DescribeCompletesAsync(
	request *DescribeCompletesRequest,
	callback chan<- DescribeCompletesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "complete",
			"function":    "describeCompletes",
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

	go p.describeCompletesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DescribeCompletes(
	request *DescribeCompletesRequest,
) (*DescribeCompletesResult, error) {
	callback := make(chan DescribeCompletesAsyncResult, 1)
	go p.DescribeCompletesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) describeCompletesByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeCompletesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCompletesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCompletesByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCompletesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeCompletesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DescribeCompletesByUserIdAsync(
	request *DescribeCompletesByUserIdRequest,
	callback chan<- DescribeCompletesByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "complete",
			"function":    "describeCompletesByUserId",
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

	go p.describeCompletesByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DescribeCompletesByUserId(
	request *DescribeCompletesByUserIdRequest,
) (*DescribeCompletesByUserIdResult, error) {
	callback := make(chan DescribeCompletesByUserIdAsyncResult, 1)
	go p.DescribeCompletesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) completeAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CompleteAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CompleteAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CompleteResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CompleteAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CompleteAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) CompleteAsync(
	request *CompleteRequest,
	callback chan<- CompleteAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "complete",
			"function":    "complete",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		bodies["missionTaskName"] = *request.MissionTaskName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
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

	go p.completeAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) Complete(
	request *CompleteRequest,
) (*CompleteResult, error) {
	callback := make(chan CompleteAsyncResult, 1)
	go p.CompleteAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) completeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CompleteByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CompleteByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CompleteByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CompleteByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CompleteByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) CompleteByUserIdAsync(
	request *CompleteByUserIdRequest,
	callback chan<- CompleteByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "complete",
			"function":    "completeByUserId",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		bodies["missionTaskName"] = *request.MissionTaskName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
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

	go p.completeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) CompleteByUserId(
	request *CompleteByUserIdRequest,
) (*CompleteByUserIdResult, error) {
	callback := make(chan CompleteByUserIdAsyncResult, 1)
	go p.CompleteByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) receiveByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReceiveByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ReceiveByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) ReceiveByUserIdAsync(
	request *ReceiveByUserIdRequest,
	callback chan<- ReceiveByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "complete",
			"function":    "receiveByUserId",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		bodies["missionTaskName"] = *request.MissionTaskName
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

	go p.receiveByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) ReceiveByUserId(
	request *ReceiveByUserIdRequest,
) (*ReceiveByUserIdResult, error) {
	callback := make(chan ReceiveByUserIdAsyncResult, 1)
	go p.ReceiveByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) revertReceiveByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RevertReceiveByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RevertReceiveByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RevertReceiveByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RevertReceiveByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- RevertReceiveByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) RevertReceiveByUserIdAsync(
	request *RevertReceiveByUserIdRequest,
	callback chan<- RevertReceiveByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "complete",
			"function":    "revertReceiveByUserId",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		bodies["missionTaskName"] = *request.MissionTaskName
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

	go p.revertReceiveByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) RevertReceiveByUserId(
	request *RevertReceiveByUserIdRequest,
) (*RevertReceiveByUserIdResult, error) {
	callback := make(chan RevertReceiveByUserIdAsyncResult, 1)
	go p.RevertReceiveByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) getCompleteAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCompleteAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCompleteAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCompleteResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCompleteAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetCompleteAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) GetCompleteAsync(
	request *GetCompleteRequest,
	callback chan<- GetCompleteAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "complete",
			"function":    "getComplete",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
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

	go p.getCompleteAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) GetComplete(
	request *GetCompleteRequest,
) (*GetCompleteResult, error) {
	callback := make(chan GetCompleteAsyncResult, 1)
	go p.GetCompleteAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) getCompleteByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCompleteByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCompleteByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCompleteByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCompleteByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetCompleteByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) GetCompleteByUserIdAsync(
	request *GetCompleteByUserIdRequest,
	callback chan<- GetCompleteByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "complete",
			"function":    "getCompleteByUserId",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
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

	go p.getCompleteByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) GetCompleteByUserId(
	request *GetCompleteByUserIdRequest,
) (*GetCompleteByUserIdResult, error) {
	callback := make(chan GetCompleteByUserIdAsyncResult, 1)
	go p.GetCompleteByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) deleteCompleteByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteCompleteByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteCompleteByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteCompleteByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteCompleteByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteCompleteByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DeleteCompleteByUserIdAsync(
	request *DeleteCompleteByUserIdRequest,
	callback chan<- DeleteCompleteByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "complete",
			"function":    "deleteCompleteByUserId",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
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

	go p.deleteCompleteByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DeleteCompleteByUserId(
	request *DeleteCompleteByUserIdRequest,
) (*DeleteCompleteByUserIdResult, error) {
	callback := make(chan DeleteCompleteByUserIdAsyncResult, 1)
	go p.DeleteCompleteByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) receiveByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReceiveByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveByStampTaskResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ReceiveByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) ReceiveByStampTaskAsync(
	request *ReceiveByStampTaskRequest,
	callback chan<- ReceiveByStampTaskAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "complete",
			"function":    "receiveByStampTask",
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

	go p.receiveByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) ReceiveByStampTask(
	request *ReceiveByStampTaskRequest,
) (*ReceiveByStampTaskResult, error) {
	callback := make(chan ReceiveByStampTaskAsyncResult, 1)
	go p.ReceiveByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) revertReceiveByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RevertReceiveByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RevertReceiveByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RevertReceiveByStampSheetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RevertReceiveByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- RevertReceiveByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) RevertReceiveByStampSheetAsync(
	request *RevertReceiveByStampSheetRequest,
	callback chan<- RevertReceiveByStampSheetAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "complete",
			"function":    "revertReceiveByStampSheet",
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

	go p.revertReceiveByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) RevertReceiveByStampSheet(
	request *RevertReceiveByStampSheetRequest,
) (*RevertReceiveByStampSheetResult, error) {
	callback := make(chan RevertReceiveByStampSheetAsyncResult, 1)
	go p.RevertReceiveByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) describeCounterModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeCounterModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCounterModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCounterModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCounterModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeCounterModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DescribeCounterModelMastersAsync(
	request *DescribeCounterModelMastersRequest,
	callback chan<- DescribeCounterModelMastersAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counterModelMaster",
			"function":    "describeCounterModelMasters",
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

	go p.describeCounterModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DescribeCounterModelMasters(
	request *DescribeCounterModelMastersRequest,
) (*DescribeCounterModelMastersResult, error) {
	callback := make(chan DescribeCounterModelMastersAsyncResult, 1)
	go p.DescribeCounterModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) createCounterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateCounterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateCounterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateCounterModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateCounterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateCounterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) CreateCounterModelMasterAsync(
	request *CreateCounterModelMasterRequest,
	callback chan<- CreateCounterModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counterModelMaster",
			"function":    "createCounterModelMaster",
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
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Scopes != nil {
		var _scopes []interface{}
		for _, item := range request.Scopes {
			_scopes = append(_scopes, item)
		}
		bodies["scopes"] = _scopes
	}
	if request.ChallengePeriodEventId != nil && *request.ChallengePeriodEventId != "" {
		bodies["challengePeriodEventId"] = *request.ChallengePeriodEventId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.createCounterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) CreateCounterModelMaster(
	request *CreateCounterModelMasterRequest,
) (*CreateCounterModelMasterResult, error) {
	callback := make(chan CreateCounterModelMasterAsyncResult, 1)
	go p.CreateCounterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) getCounterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCounterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCounterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCounterModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCounterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetCounterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) GetCounterModelMasterAsync(
	request *GetCounterModelMasterRequest,
	callback chan<- GetCounterModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counterModelMaster",
			"function":    "getCounterModelMaster",
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
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getCounterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) GetCounterModelMaster(
	request *GetCounterModelMasterRequest,
) (*GetCounterModelMasterResult, error) {
	callback := make(chan GetCounterModelMasterAsyncResult, 1)
	go p.GetCounterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) updateCounterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCounterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCounterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCounterModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCounterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateCounterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) UpdateCounterModelMasterAsync(
	request *UpdateCounterModelMasterRequest,
	callback chan<- UpdateCounterModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counterModelMaster",
			"function":    "updateCounterModelMaster",
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
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Scopes != nil {
		var _scopes []interface{}
		for _, item := range request.Scopes {
			_scopes = append(_scopes, item)
		}
		bodies["scopes"] = _scopes
	}
	if request.ChallengePeriodEventId != nil && *request.ChallengePeriodEventId != "" {
		bodies["challengePeriodEventId"] = *request.ChallengePeriodEventId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateCounterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) UpdateCounterModelMaster(
	request *UpdateCounterModelMasterRequest,
) (*UpdateCounterModelMasterResult, error) {
	callback := make(chan UpdateCounterModelMasterAsyncResult, 1)
	go p.UpdateCounterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) deleteCounterModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteCounterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteCounterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteCounterModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteCounterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteCounterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DeleteCounterModelMasterAsync(
	request *DeleteCounterModelMasterRequest,
	callback chan<- DeleteCounterModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counterModelMaster",
			"function":    "deleteCounterModelMaster",
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
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteCounterModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DeleteCounterModelMaster(
	request *DeleteCounterModelMasterRequest,
) (*DeleteCounterModelMasterResult, error) {
	callback := make(chan DeleteCounterModelMasterAsyncResult, 1)
	go p.DeleteCounterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) describeMissionGroupModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMissionGroupModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMissionGroupModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMissionGroupModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionGroupModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeMissionGroupModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DescribeMissionGroupModelMastersAsync(
	request *DescribeMissionGroupModelMastersRequest,
	callback chan<- DescribeMissionGroupModelMastersAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionGroupModelMaster",
			"function":    "describeMissionGroupModelMasters",
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

	go p.describeMissionGroupModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DescribeMissionGroupModelMasters(
	request *DescribeMissionGroupModelMastersRequest,
) (*DescribeMissionGroupModelMastersResult, error) {
	callback := make(chan DescribeMissionGroupModelMastersAsyncResult, 1)
	go p.DescribeMissionGroupModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) createMissionGroupModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateMissionGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateMissionGroupModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateMissionGroupModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateMissionGroupModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateMissionGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) CreateMissionGroupModelMasterAsync(
	request *CreateMissionGroupModelMasterRequest,
	callback chan<- CreateMissionGroupModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionGroupModelMaster",
			"function":    "createMissionGroupModelMaster",
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
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.ResetType != nil && *request.ResetType != "" {
		bodies["resetType"] = *request.ResetType
	}
	if request.ResetDayOfMonth != nil {
		bodies["resetDayOfMonth"] = *request.ResetDayOfMonth
	}
	if request.ResetDayOfWeek != nil && *request.ResetDayOfWeek != "" {
		bodies["resetDayOfWeek"] = *request.ResetDayOfWeek
	}
	if request.ResetHour != nil {
		bodies["resetHour"] = *request.ResetHour
	}
	if request.CompleteNotificationNamespaceId != nil && *request.CompleteNotificationNamespaceId != "" {
		bodies["completeNotificationNamespaceId"] = *request.CompleteNotificationNamespaceId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.createMissionGroupModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) CreateMissionGroupModelMaster(
	request *CreateMissionGroupModelMasterRequest,
) (*CreateMissionGroupModelMasterResult, error) {
	callback := make(chan CreateMissionGroupModelMasterAsyncResult, 1)
	go p.CreateMissionGroupModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) getMissionGroupModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMissionGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMissionGroupModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMissionGroupModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMissionGroupModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetMissionGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) GetMissionGroupModelMasterAsync(
	request *GetMissionGroupModelMasterRequest,
	callback chan<- GetMissionGroupModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionGroupModelMaster",
			"function":    "getMissionGroupModelMaster",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getMissionGroupModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) GetMissionGroupModelMaster(
	request *GetMissionGroupModelMasterRequest,
) (*GetMissionGroupModelMasterResult, error) {
	callback := make(chan GetMissionGroupModelMasterAsyncResult, 1)
	go p.GetMissionGroupModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) updateMissionGroupModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateMissionGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMissionGroupModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMissionGroupModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMissionGroupModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateMissionGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) UpdateMissionGroupModelMasterAsync(
	request *UpdateMissionGroupModelMasterRequest,
	callback chan<- UpdateMissionGroupModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionGroupModelMaster",
			"function":    "updateMissionGroupModelMaster",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.ResetType != nil && *request.ResetType != "" {
		bodies["resetType"] = *request.ResetType
	}
	if request.ResetDayOfMonth != nil {
		bodies["resetDayOfMonth"] = *request.ResetDayOfMonth
	}
	if request.ResetDayOfWeek != nil && *request.ResetDayOfWeek != "" {
		bodies["resetDayOfWeek"] = *request.ResetDayOfWeek
	}
	if request.ResetHour != nil {
		bodies["resetHour"] = *request.ResetHour
	}
	if request.CompleteNotificationNamespaceId != nil && *request.CompleteNotificationNamespaceId != "" {
		bodies["completeNotificationNamespaceId"] = *request.CompleteNotificationNamespaceId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateMissionGroupModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) UpdateMissionGroupModelMaster(
	request *UpdateMissionGroupModelMasterRequest,
) (*UpdateMissionGroupModelMasterResult, error) {
	callback := make(chan UpdateMissionGroupModelMasterAsyncResult, 1)
	go p.UpdateMissionGroupModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) deleteMissionGroupModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteMissionGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMissionGroupModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMissionGroupModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMissionGroupModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteMissionGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DeleteMissionGroupModelMasterAsync(
	request *DeleteMissionGroupModelMasterRequest,
	callback chan<- DeleteMissionGroupModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionGroupModelMaster",
			"function":    "deleteMissionGroupModelMaster",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteMissionGroupModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DeleteMissionGroupModelMaster(
	request *DeleteMissionGroupModelMasterRequest,
) (*DeleteMissionGroupModelMasterResult, error) {
	callback := make(chan DeleteMissionGroupModelMasterAsyncResult, 1)
	go p.DeleteMissionGroupModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2MissionWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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

func (p Gs2MissionWebSocketClient) DescribeNamespaces(
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

func (p Gs2MissionWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2MissionWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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
	if request.MissionCompleteScript != nil {
		bodies["missionCompleteScript"] = request.MissionCompleteScript.ToDict()
	}
	if request.CounterIncrementScript != nil {
		bodies["counterIncrementScript"] = request.CounterIncrementScript.ToDict()
	}
	if request.ReceiveRewardsScript != nil {
		bodies["receiveRewardsScript"] = request.ReceiveRewardsScript.ToDict()
	}
	if request.CompleteNotification != nil {
		bodies["completeNotification"] = request.CompleteNotification.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
		bodies["queueNamespaceId"] = *request.QueueNamespaceId
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

func (p Gs2MissionWebSocketClient) CreateNamespace(
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

func (p Gs2MissionWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2MissionWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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

func (p Gs2MissionWebSocketClient) GetNamespaceStatus(
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

func (p Gs2MissionWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2MissionWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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

func (p Gs2MissionWebSocketClient) GetNamespace(
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

func (p Gs2MissionWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2MissionWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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
	if request.MissionCompleteScript != nil {
		bodies["missionCompleteScript"] = request.MissionCompleteScript.ToDict()
	}
	if request.CounterIncrementScript != nil {
		bodies["counterIncrementScript"] = request.CounterIncrementScript.ToDict()
	}
	if request.ReceiveRewardsScript != nil {
		bodies["receiveRewardsScript"] = request.ReceiveRewardsScript.ToDict()
	}
	if request.CompleteNotification != nil {
		bodies["completeNotification"] = request.CompleteNotification.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
		bodies["queueNamespaceId"] = *request.QueueNamespaceId
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

func (p Gs2MissionWebSocketClient) UpdateNamespace(
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

func (p Gs2MissionWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2MissionWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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

func (p Gs2MissionWebSocketClient) DeleteNamespace(
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

func (p Gs2MissionWebSocketClient) dumpUserDataByUserIdAsyncHandler(
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

func (p Gs2MissionWebSocketClient) DumpUserDataByUserIdAsync(
	request *DumpUserDataByUserIdRequest,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.dumpUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DumpUserDataByUserId(
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

func (p Gs2MissionWebSocketClient) checkDumpUserDataByUserIdAsyncHandler(
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

func (p Gs2MissionWebSocketClient) CheckDumpUserDataByUserIdAsync(
	request *CheckDumpUserDataByUserIdRequest,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.checkDumpUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) CheckDumpUserDataByUserId(
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

func (p Gs2MissionWebSocketClient) cleanUserDataByUserIdAsyncHandler(
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

func (p Gs2MissionWebSocketClient) CleanUserDataByUserIdAsync(
	request *CleanUserDataByUserIdRequest,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.cleanUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) CleanUserDataByUserId(
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

func (p Gs2MissionWebSocketClient) checkCleanUserDataByUserIdAsyncHandler(
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

func (p Gs2MissionWebSocketClient) CheckCleanUserDataByUserIdAsync(
	request *CheckCleanUserDataByUserIdRequest,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.checkCleanUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) CheckCleanUserDataByUserId(
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

func (p Gs2MissionWebSocketClient) prepareImportUserDataByUserIdAsyncHandler(
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

func (p Gs2MissionWebSocketClient) PrepareImportUserDataByUserIdAsync(
	request *PrepareImportUserDataByUserIdRequest,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.prepareImportUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) PrepareImportUserDataByUserId(
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

func (p Gs2MissionWebSocketClient) importUserDataByUserIdAsyncHandler(
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

func (p Gs2MissionWebSocketClient) ImportUserDataByUserIdAsync(
	request *ImportUserDataByUserIdRequest,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.importUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) ImportUserDataByUserId(
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

func (p Gs2MissionWebSocketClient) checkImportUserDataByUserIdAsyncHandler(
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

func (p Gs2MissionWebSocketClient) CheckImportUserDataByUserIdAsync(
	request *CheckImportUserDataByUserIdRequest,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
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
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}

	go p.checkImportUserDataByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) CheckImportUserDataByUserId(
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

func (p Gs2MissionWebSocketClient) describeCountersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeCountersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCountersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCountersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCountersAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeCountersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DescribeCountersAsync(
	request *DescribeCountersRequest,
	callback chan<- DescribeCountersAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counter",
			"function":    "describeCounters",
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

	go p.describeCountersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DescribeCounters(
	request *DescribeCountersRequest,
) (*DescribeCountersResult, error) {
	callback := make(chan DescribeCountersAsyncResult, 1)
	go p.DescribeCountersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) describeCountersByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeCountersByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCountersByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCountersByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCountersByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeCountersByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DescribeCountersByUserIdAsync(
	request *DescribeCountersByUserIdRequest,
	callback chan<- DescribeCountersByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counter",
			"function":    "describeCountersByUserId",
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

	go p.describeCountersByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DescribeCountersByUserId(
	request *DescribeCountersByUserIdRequest,
) (*DescribeCountersByUserIdResult, error) {
	callback := make(chan DescribeCountersByUserIdAsyncResult, 1)
	go p.DescribeCountersByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) increaseCounterByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- IncreaseCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- IncreaseCounterByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result IncreaseCounterByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IncreaseCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "counter.increase.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
		}
	}
	callback <- IncreaseCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) IncreaseCounterByUserIdAsync(
	request *IncreaseCounterByUserIdRequest,
	callback chan<- IncreaseCounterByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counter",
			"function":    "increaseCounterByUserId",
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
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Value != nil {
		bodies["value"] = *request.Value
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

	go p.increaseCounterByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) IncreaseCounterByUserId(
	request *IncreaseCounterByUserIdRequest,
) (*IncreaseCounterByUserIdResult, error) {
	callback := make(chan IncreaseCounterByUserIdAsyncResult, 1)
	go p.IncreaseCounterByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) decreaseCounterByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DecreaseCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DecreaseCounterByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DecreaseCounterByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "counter.increase.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
		}
	}
	callback <- DecreaseCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DecreaseCounterByUserIdAsync(
	request *DecreaseCounterByUserIdRequest,
	callback chan<- DecreaseCounterByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counter",
			"function":    "decreaseCounterByUserId",
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
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Value != nil {
		bodies["value"] = *request.Value
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

	go p.decreaseCounterByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DecreaseCounterByUserId(
	request *DecreaseCounterByUserIdRequest,
) (*DecreaseCounterByUserIdResult, error) {
	callback := make(chan DecreaseCounterByUserIdAsyncResult, 1)
	go p.DecreaseCounterByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) getCounterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCounterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCounterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCounterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCounterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetCounterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) GetCounterAsync(
	request *GetCounterRequest,
	callback chan<- GetCounterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counter",
			"function":    "getCounter",
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
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
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

	go p.getCounterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) GetCounter(
	request *GetCounterRequest,
) (*GetCounterResult, error) {
	callback := make(chan GetCounterAsyncResult, 1)
	go p.GetCounterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) getCounterByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCounterByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCounterByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) GetCounterByUserIdAsync(
	request *GetCounterByUserIdRequest,
	callback chan<- GetCounterByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counter",
			"function":    "getCounterByUserId",
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
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
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

	go p.getCounterByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) GetCounterByUserId(
	request *GetCounterByUserIdRequest,
) (*GetCounterByUserIdResult, error) {
	callback := make(chan GetCounterByUserIdAsyncResult, 1)
	go p.GetCounterByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) deleteCounterByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteCounterByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteCounterByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DeleteCounterByUserIdAsync(
	request *DeleteCounterByUserIdRequest,
	callback chan<- DeleteCounterByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counter",
			"function":    "deleteCounterByUserId",
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
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
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

	go p.deleteCounterByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DeleteCounterByUserId(
	request *DeleteCounterByUserIdRequest,
) (*DeleteCounterByUserIdResult, error) {
	callback := make(chan DeleteCounterByUserIdAsyncResult, 1)
	go p.DeleteCounterByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) increaseByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- IncreaseByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- IncreaseByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result IncreaseByStampSheetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IncreaseByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- IncreaseByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) IncreaseByStampSheetAsync(
	request *IncreaseByStampSheetRequest,
	callback chan<- IncreaseByStampSheetAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counter",
			"function":    "increaseByStampSheet",
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

	go p.increaseByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) IncreaseByStampSheet(
	request *IncreaseByStampSheetRequest,
) (*IncreaseByStampSheetResult, error) {
	callback := make(chan IncreaseByStampSheetAsyncResult, 1)
	go p.IncreaseByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) decreaseByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DecreaseByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DecreaseByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DecreaseByStampTaskResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DecreaseByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DecreaseByStampTaskAsync(
	request *DecreaseByStampTaskRequest,
	callback chan<- DecreaseByStampTaskAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counter",
			"function":    "decreaseByStampTask",
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

	go p.decreaseByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DecreaseByStampTask(
	request *DecreaseByStampTaskRequest,
) (*DecreaseByStampTaskResult, error) {
	callback := make(chan DecreaseByStampTaskAsyncResult, 1)
	go p.DecreaseByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) exportMasterAsyncHandler(
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

func (p Gs2MissionWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "currentMissionMaster",
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

func (p Gs2MissionWebSocketClient) ExportMaster(
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

func (p Gs2MissionWebSocketClient) getCurrentMissionMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentMissionMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentMissionMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentMissionMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentMissionMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetCurrentMissionMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) GetCurrentMissionMasterAsync(
	request *GetCurrentMissionMasterRequest,
	callback chan<- GetCurrentMissionMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "currentMissionMaster",
			"function":    "getCurrentMissionMaster",
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

	go p.getCurrentMissionMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) GetCurrentMissionMaster(
	request *GetCurrentMissionMasterRequest,
) (*GetCurrentMissionMasterResult, error) {
	callback := make(chan GetCurrentMissionMasterAsyncResult, 1)
	go p.GetCurrentMissionMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) updateCurrentMissionMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentMissionMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentMissionMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentMissionMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentMissionMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateCurrentMissionMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) UpdateCurrentMissionMasterAsync(
	request *UpdateCurrentMissionMasterRequest,
	callback chan<- UpdateCurrentMissionMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "currentMissionMaster",
			"function":    "updateCurrentMissionMaster",
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

	go p.updateCurrentMissionMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) UpdateCurrentMissionMaster(
	request *UpdateCurrentMissionMasterRequest,
) (*UpdateCurrentMissionMasterResult, error) {
	callback := make(chan UpdateCurrentMissionMasterAsyncResult, 1)
	go p.UpdateCurrentMissionMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) updateCurrentMissionMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentMissionMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentMissionMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentMissionMasterFromGitHubResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentMissionMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateCurrentMissionMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) UpdateCurrentMissionMasterFromGitHubAsync(
	request *UpdateCurrentMissionMasterFromGitHubRequest,
	callback chan<- UpdateCurrentMissionMasterFromGitHubAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "currentMissionMaster",
			"function":    "updateCurrentMissionMasterFromGitHub",
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

	go p.updateCurrentMissionMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) UpdateCurrentMissionMasterFromGitHub(
	request *UpdateCurrentMissionMasterFromGitHubRequest,
) (*UpdateCurrentMissionMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentMissionMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentMissionMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) describeCounterModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeCounterModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCounterModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCounterModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCounterModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeCounterModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DescribeCounterModelsAsync(
	request *DescribeCounterModelsRequest,
	callback chan<- DescribeCounterModelsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counterModel",
			"function":    "describeCounterModels",
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

	go p.describeCounterModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DescribeCounterModels(
	request *DescribeCounterModelsRequest,
) (*DescribeCounterModelsResult, error) {
	callback := make(chan DescribeCounterModelsAsyncResult, 1)
	go p.DescribeCounterModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) getCounterModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCounterModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCounterModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCounterModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCounterModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetCounterModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) GetCounterModelAsync(
	request *GetCounterModelRequest,
	callback chan<- GetCounterModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "counterModel",
			"function":    "getCounterModel",
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
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getCounterModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) GetCounterModel(
	request *GetCounterModelRequest,
) (*GetCounterModelResult, error) {
	callback := make(chan GetCounterModelAsyncResult, 1)
	go p.GetCounterModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) describeMissionGroupModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMissionGroupModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMissionGroupModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMissionGroupModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionGroupModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeMissionGroupModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DescribeMissionGroupModelsAsync(
	request *DescribeMissionGroupModelsRequest,
	callback chan<- DescribeMissionGroupModelsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionGroupModel",
			"function":    "describeMissionGroupModels",
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

	go p.describeMissionGroupModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DescribeMissionGroupModels(
	request *DescribeMissionGroupModelsRequest,
) (*DescribeMissionGroupModelsResult, error) {
	callback := make(chan DescribeMissionGroupModelsAsyncResult, 1)
	go p.DescribeMissionGroupModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) getMissionGroupModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMissionGroupModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMissionGroupModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMissionGroupModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMissionGroupModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetMissionGroupModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) GetMissionGroupModelAsync(
	request *GetMissionGroupModelRequest,
	callback chan<- GetMissionGroupModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionGroupModel",
			"function":    "getMissionGroupModel",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getMissionGroupModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) GetMissionGroupModel(
	request *GetMissionGroupModelRequest,
) (*GetMissionGroupModelResult, error) {
	callback := make(chan GetMissionGroupModelAsyncResult, 1)
	go p.GetMissionGroupModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) describeMissionTaskModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMissionTaskModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMissionTaskModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMissionTaskModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionTaskModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeMissionTaskModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DescribeMissionTaskModelsAsync(
	request *DescribeMissionTaskModelsRequest,
	callback chan<- DescribeMissionTaskModelsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionTaskModel",
			"function":    "describeMissionTaskModels",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.describeMissionTaskModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DescribeMissionTaskModels(
	request *DescribeMissionTaskModelsRequest,
) (*DescribeMissionTaskModelsResult, error) {
	callback := make(chan DescribeMissionTaskModelsAsyncResult, 1)
	go p.DescribeMissionTaskModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) getMissionTaskModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMissionTaskModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMissionTaskModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMissionTaskModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMissionTaskModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetMissionTaskModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) GetMissionTaskModelAsync(
	request *GetMissionTaskModelRequest,
	callback chan<- GetMissionTaskModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionTaskModel",
			"function":    "getMissionTaskModel",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		bodies["missionTaskName"] = *request.MissionTaskName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getMissionTaskModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) GetMissionTaskModel(
	request *GetMissionTaskModelRequest,
) (*GetMissionTaskModelResult, error) {
	callback := make(chan GetMissionTaskModelAsyncResult, 1)
	go p.GetMissionTaskModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) describeMissionTaskModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMissionTaskModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMissionTaskModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMissionTaskModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionTaskModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeMissionTaskModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DescribeMissionTaskModelMastersAsync(
	request *DescribeMissionTaskModelMastersRequest,
	callback chan<- DescribeMissionTaskModelMastersAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionTaskModelMaster",
			"function":    "describeMissionTaskModelMasters",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
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

	go p.describeMissionTaskModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DescribeMissionTaskModelMasters(
	request *DescribeMissionTaskModelMastersRequest,
) (*DescribeMissionTaskModelMastersResult, error) {
	callback := make(chan DescribeMissionTaskModelMastersAsyncResult, 1)
	go p.DescribeMissionTaskModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) createMissionTaskModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateMissionTaskModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateMissionTaskModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateMissionTaskModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateMissionTaskModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateMissionTaskModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) CreateMissionTaskModelMasterAsync(
	request *CreateMissionTaskModelMasterRequest,
	callback chan<- CreateMissionTaskModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionTaskModelMaster",
			"function":    "createMissionTaskModelMaster",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
	}
	if request.TargetResetType != nil && *request.TargetResetType != "" {
		bodies["targetResetType"] = *request.TargetResetType
	}
	if request.TargetValue != nil {
		bodies["targetValue"] = *request.TargetValue
	}
	if request.CompleteAcquireActions != nil {
		var _completeAcquireActions []interface{}
		for _, item := range request.CompleteAcquireActions {
			_completeAcquireActions = append(_completeAcquireActions, item)
		}
		bodies["completeAcquireActions"] = _completeAcquireActions
	}
	if request.ChallengePeriodEventId != nil && *request.ChallengePeriodEventId != "" {
		bodies["challengePeriodEventId"] = *request.ChallengePeriodEventId
	}
	if request.PremiseMissionTaskName != nil && *request.PremiseMissionTaskName != "" {
		bodies["premiseMissionTaskName"] = *request.PremiseMissionTaskName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.createMissionTaskModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) CreateMissionTaskModelMaster(
	request *CreateMissionTaskModelMasterRequest,
) (*CreateMissionTaskModelMasterResult, error) {
	callback := make(chan CreateMissionTaskModelMasterAsyncResult, 1)
	go p.CreateMissionTaskModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) getMissionTaskModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMissionTaskModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMissionTaskModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMissionTaskModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMissionTaskModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetMissionTaskModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) GetMissionTaskModelMasterAsync(
	request *GetMissionTaskModelMasterRequest,
	callback chan<- GetMissionTaskModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionTaskModelMaster",
			"function":    "getMissionTaskModelMaster",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		bodies["missionTaskName"] = *request.MissionTaskName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getMissionTaskModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) GetMissionTaskModelMaster(
	request *GetMissionTaskModelMasterRequest,
) (*GetMissionTaskModelMasterResult, error) {
	callback := make(chan GetMissionTaskModelMasterAsyncResult, 1)
	go p.GetMissionTaskModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) updateMissionTaskModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateMissionTaskModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMissionTaskModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMissionTaskModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMissionTaskModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateMissionTaskModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) UpdateMissionTaskModelMasterAsync(
	request *UpdateMissionTaskModelMasterRequest,
	callback chan<- UpdateMissionTaskModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionTaskModelMaster",
			"function":    "updateMissionTaskModelMaster",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		bodies["missionTaskName"] = *request.MissionTaskName
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
	}
	if request.TargetResetType != nil && *request.TargetResetType != "" {
		bodies["targetResetType"] = *request.TargetResetType
	}
	if request.TargetValue != nil {
		bodies["targetValue"] = *request.TargetValue
	}
	if request.CompleteAcquireActions != nil {
		var _completeAcquireActions []interface{}
		for _, item := range request.CompleteAcquireActions {
			_completeAcquireActions = append(_completeAcquireActions, item)
		}
		bodies["completeAcquireActions"] = _completeAcquireActions
	}
	if request.ChallengePeriodEventId != nil && *request.ChallengePeriodEventId != "" {
		bodies["challengePeriodEventId"] = *request.ChallengePeriodEventId
	}
	if request.PremiseMissionTaskName != nil && *request.PremiseMissionTaskName != "" {
		bodies["premiseMissionTaskName"] = *request.PremiseMissionTaskName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateMissionTaskModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) UpdateMissionTaskModelMaster(
	request *UpdateMissionTaskModelMasterRequest,
) (*UpdateMissionTaskModelMasterResult, error) {
	callback := make(chan UpdateMissionTaskModelMasterAsyncResult, 1)
	go p.UpdateMissionTaskModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2MissionWebSocketClient) deleteMissionTaskModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteMissionTaskModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMissionTaskModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMissionTaskModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMissionTaskModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteMissionTaskModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionWebSocketClient) DeleteMissionTaskModelMasterAsync(
	request *DeleteMissionTaskModelMasterRequest,
	callback chan<- DeleteMissionTaskModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "mission",
			"component":   "missionTaskModelMaster",
			"function":    "deleteMissionTaskModelMaster",
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		bodies["missionTaskName"] = *request.MissionTaskName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteMissionTaskModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2MissionWebSocketClient) DeleteMissionTaskModelMaster(
	request *DeleteMissionTaskModelMasterRequest,
) (*DeleteMissionTaskModelMasterResult, error) {
	callback := make(chan DeleteMissionTaskModelMasterAsyncResult, 1)
	go p.DeleteMissionTaskModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
