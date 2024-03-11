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

package grade

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2GradeWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2GradeWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2GradeWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2GradeWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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

func (p Gs2GradeWebSocketClient) DescribeNamespaces(
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

func (p Gs2GradeWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2GradeWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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
	if request.ChangeGradeScript != nil {
		bodies["changeGradeScript"] = request.ChangeGradeScript.ToDict()
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

func (p Gs2GradeWebSocketClient) CreateNamespace(
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

func (p Gs2GradeWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2GradeWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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

func (p Gs2GradeWebSocketClient) GetNamespaceStatus(
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

func (p Gs2GradeWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2GradeWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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

func (p Gs2GradeWebSocketClient) GetNamespace(
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

func (p Gs2GradeWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2GradeWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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
	if request.ChangeGradeScript != nil {
		bodies["changeGradeScript"] = request.ChangeGradeScript.ToDict()
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

func (p Gs2GradeWebSocketClient) UpdateNamespace(
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

func (p Gs2GradeWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2GradeWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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

func (p Gs2GradeWebSocketClient) DeleteNamespace(
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

func (p Gs2GradeWebSocketClient) dumpUserDataByUserIdAsyncHandler(
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

func (p Gs2GradeWebSocketClient) DumpUserDataByUserIdAsync(
	request *DumpUserDataByUserIdRequest,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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

func (p Gs2GradeWebSocketClient) DumpUserDataByUserId(
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

func (p Gs2GradeWebSocketClient) checkDumpUserDataByUserIdAsyncHandler(
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

func (p Gs2GradeWebSocketClient) CheckDumpUserDataByUserIdAsync(
	request *CheckDumpUserDataByUserIdRequest,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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

func (p Gs2GradeWebSocketClient) CheckDumpUserDataByUserId(
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

func (p Gs2GradeWebSocketClient) cleanUserDataByUserIdAsyncHandler(
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

func (p Gs2GradeWebSocketClient) CleanUserDataByUserIdAsync(
	request *CleanUserDataByUserIdRequest,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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

func (p Gs2GradeWebSocketClient) CleanUserDataByUserId(
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

func (p Gs2GradeWebSocketClient) checkCleanUserDataByUserIdAsyncHandler(
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

func (p Gs2GradeWebSocketClient) CheckCleanUserDataByUserIdAsync(
	request *CheckCleanUserDataByUserIdRequest,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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

func (p Gs2GradeWebSocketClient) CheckCleanUserDataByUserId(
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

func (p Gs2GradeWebSocketClient) prepareImportUserDataByUserIdAsyncHandler(
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

func (p Gs2GradeWebSocketClient) PrepareImportUserDataByUserIdAsync(
	request *PrepareImportUserDataByUserIdRequest,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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

func (p Gs2GradeWebSocketClient) PrepareImportUserDataByUserId(
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

func (p Gs2GradeWebSocketClient) importUserDataByUserIdAsyncHandler(
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

func (p Gs2GradeWebSocketClient) ImportUserDataByUserIdAsync(
	request *ImportUserDataByUserIdRequest,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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

func (p Gs2GradeWebSocketClient) ImportUserDataByUserId(
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

func (p Gs2GradeWebSocketClient) checkImportUserDataByUserIdAsyncHandler(
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

func (p Gs2GradeWebSocketClient) CheckImportUserDataByUserIdAsync(
	request *CheckImportUserDataByUserIdRequest,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
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

func (p Gs2GradeWebSocketClient) CheckImportUserDataByUserId(
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

func (p Gs2GradeWebSocketClient) describeGradeModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGradeModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGradeModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGradeModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGradeModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeGradeModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) DescribeGradeModelMastersAsync(
	request *DescribeGradeModelMastersRequest,
	callback chan<- DescribeGradeModelMastersAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "gradeModelMaster",
			"function":    "describeGradeModelMasters",
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

	go p.describeGradeModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) DescribeGradeModelMasters(
	request *DescribeGradeModelMastersRequest,
) (*DescribeGradeModelMastersResult, error) {
	callback := make(chan DescribeGradeModelMastersAsyncResult, 1)
	go p.DescribeGradeModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) createGradeModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateGradeModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGradeModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGradeModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGradeModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateGradeModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) CreateGradeModelMasterAsync(
	request *CreateGradeModelMasterRequest,
	callback chan<- CreateGradeModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "gradeModelMaster",
			"function":    "createGradeModelMaster",
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
	if request.DefaultGrades != nil {
		var _defaultGrades []interface{}
		for _, item := range request.DefaultGrades {
			_defaultGrades = append(_defaultGrades, item)
		}
		bodies["defaultGrades"] = _defaultGrades
	}
	if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
		bodies["experienceModelId"] = *request.ExperienceModelId
	}
	if request.GradeEntries != nil {
		var _gradeEntries []interface{}
		for _, item := range request.GradeEntries {
			_gradeEntries = append(_gradeEntries, item)
		}
		bodies["gradeEntries"] = _gradeEntries
	}
	if request.AcquireActionRates != nil {
		var _acquireActionRates []interface{}
		for _, item := range request.AcquireActionRates {
			_acquireActionRates = append(_acquireActionRates, item)
		}
		bodies["acquireActionRates"] = _acquireActionRates
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.createGradeModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) CreateGradeModelMaster(
	request *CreateGradeModelMasterRequest,
) (*CreateGradeModelMasterResult, error) {
	callback := make(chan CreateGradeModelMasterAsyncResult, 1)
	go p.CreateGradeModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) getGradeModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGradeModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGradeModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGradeModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGradeModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGradeModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) GetGradeModelMasterAsync(
	request *GetGradeModelMasterRequest,
	callback chan<- GetGradeModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "gradeModelMaster",
			"function":    "getGradeModelMaster",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getGradeModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) GetGradeModelMaster(
	request *GetGradeModelMasterRequest,
) (*GetGradeModelMasterResult, error) {
	callback := make(chan GetGradeModelMasterAsyncResult, 1)
	go p.GetGradeModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) updateGradeModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateGradeModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGradeModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGradeModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateGradeModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateGradeModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) UpdateGradeModelMasterAsync(
	request *UpdateGradeModelMasterRequest,
	callback chan<- UpdateGradeModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "gradeModelMaster",
			"function":    "updateGradeModelMaster",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.DefaultGrades != nil {
		var _defaultGrades []interface{}
		for _, item := range request.DefaultGrades {
			_defaultGrades = append(_defaultGrades, item)
		}
		bodies["defaultGrades"] = _defaultGrades
	}
	if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
		bodies["experienceModelId"] = *request.ExperienceModelId
	}
	if request.GradeEntries != nil {
		var _gradeEntries []interface{}
		for _, item := range request.GradeEntries {
			_gradeEntries = append(_gradeEntries, item)
		}
		bodies["gradeEntries"] = _gradeEntries
	}
	if request.AcquireActionRates != nil {
		var _acquireActionRates []interface{}
		for _, item := range request.AcquireActionRates {
			_acquireActionRates = append(_acquireActionRates, item)
		}
		bodies["acquireActionRates"] = _acquireActionRates
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateGradeModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) UpdateGradeModelMaster(
	request *UpdateGradeModelMasterRequest,
) (*UpdateGradeModelMasterResult, error) {
	callback := make(chan UpdateGradeModelMasterAsyncResult, 1)
	go p.UpdateGradeModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) deleteGradeModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteGradeModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGradeModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGradeModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGradeModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteGradeModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) DeleteGradeModelMasterAsync(
	request *DeleteGradeModelMasterRequest,
	callback chan<- DeleteGradeModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "gradeModelMaster",
			"function":    "deleteGradeModelMaster",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteGradeModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) DeleteGradeModelMaster(
	request *DeleteGradeModelMasterRequest,
) (*DeleteGradeModelMasterResult, error) {
	callback := make(chan DeleteGradeModelMasterAsyncResult, 1)
	go p.DeleteGradeModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) describeGradeModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGradeModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGradeModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGradeModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGradeModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeGradeModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) DescribeGradeModelsAsync(
	request *DescribeGradeModelsRequest,
	callback chan<- DescribeGradeModelsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "gradeModel",
			"function":    "describeGradeModels",
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

	go p.describeGradeModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) DescribeGradeModels(
	request *DescribeGradeModelsRequest,
) (*DescribeGradeModelsResult, error) {
	callback := make(chan DescribeGradeModelsAsyncResult, 1)
	go p.DescribeGradeModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) getGradeModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGradeModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGradeModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGradeModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGradeModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGradeModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) GetGradeModelAsync(
	request *GetGradeModelRequest,
	callback chan<- GetGradeModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "gradeModel",
			"function":    "getGradeModel",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getGradeModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) GetGradeModel(
	request *GetGradeModelRequest,
) (*GetGradeModelResult, error) {
	callback := make(chan GetGradeModelAsyncResult, 1)
	go p.GetGradeModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) describeStatusesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStatusesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStatusesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStatusesResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStatusesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeStatusesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) DescribeStatusesAsync(
	request *DescribeStatusesRequest,
	callback chan<- DescribeStatusesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "describeStatuses",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
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

	go p.describeStatusesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) DescribeStatuses(
	request *DescribeStatusesRequest,
) (*DescribeStatusesResult, error) {
	callback := make(chan DescribeStatusesAsyncResult, 1)
	go p.DescribeStatusesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) describeStatusesByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStatusesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStatusesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStatusesByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStatusesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeStatusesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) DescribeStatusesByUserIdAsync(
	request *DescribeStatusesByUserIdRequest,
	callback chan<- DescribeStatusesByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "describeStatusesByUserId",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
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

	go p.describeStatusesByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) DescribeStatusesByUserId(
	request *DescribeStatusesByUserIdRequest,
) (*DescribeStatusesByUserIdResult, error) {
	callback := make(chan DescribeStatusesByUserIdAsyncResult, 1)
	go p.DescribeStatusesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) getStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) GetStatusAsync(
	request *GetStatusRequest,
	callback chan<- GetStatusAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "getStatus",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) GetStatus(
	request *GetStatusRequest,
) (*GetStatusResult, error) {
	callback := make(chan GetStatusAsyncResult, 1)
	go p.GetStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) getStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) GetStatusByUserIdAsync(
	request *GetStatusByUserIdRequest,
	callback chan<- GetStatusByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "getStatusByUserId",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) GetStatusByUserId(
	request *GetStatusByUserIdRequest,
) (*GetStatusByUserIdResult, error) {
	callback := make(chan GetStatusByUserIdAsyncResult, 1)
	go p.GetStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) addGradeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddGradeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddGradeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddGradeByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddGradeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- AddGradeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) AddGradeByUserIdAsync(
	request *AddGradeByUserIdRequest,
	callback chan<- AddGradeByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "addGradeByUserId",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.GradeValue != nil {
		bodies["gradeValue"] = *request.GradeValue
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

	go p.addGradeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) AddGradeByUserId(
	request *AddGradeByUserIdRequest,
) (*AddGradeByUserIdResult, error) {
	callback := make(chan AddGradeByUserIdAsyncResult, 1)
	go p.AddGradeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) subGradeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SubGradeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubGradeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubGradeByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SubGradeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- SubGradeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) SubGradeByUserIdAsync(
	request *SubGradeByUserIdRequest,
	callback chan<- SubGradeByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "subGradeByUserId",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.GradeValue != nil {
		bodies["gradeValue"] = *request.GradeValue
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

	go p.subGradeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) SubGradeByUserId(
	request *SubGradeByUserIdRequest,
) (*SubGradeByUserIdResult, error) {
	callback := make(chan SubGradeByUserIdAsyncResult, 1)
	go p.SubGradeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) setGradeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetGradeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetGradeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetGradeByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetGradeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- SetGradeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) SetGradeByUserIdAsync(
	request *SetGradeByUserIdRequest,
	callback chan<- SetGradeByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "setGradeByUserId",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.GradeValue != nil {
		bodies["gradeValue"] = *request.GradeValue
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

	go p.setGradeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) SetGradeByUserId(
	request *SetGradeByUserIdRequest,
) (*SetGradeByUserIdResult, error) {
	callback := make(chan SetGradeByUserIdAsyncResult, 1)
	go p.SetGradeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) applyRankCapAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ApplyRankCapAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ApplyRankCapAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ApplyRankCapResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ApplyRankCapAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ApplyRankCapAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) ApplyRankCapAsync(
	request *ApplyRankCapRequest,
	callback chan<- ApplyRankCapAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "applyRankCap",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
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

	go p.applyRankCapAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) ApplyRankCap(
	request *ApplyRankCapRequest,
) (*ApplyRankCapResult, error) {
	callback := make(chan ApplyRankCapAsyncResult, 1)
	go p.ApplyRankCapAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) applyRankCapByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ApplyRankCapByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ApplyRankCapByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ApplyRankCapByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ApplyRankCapByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ApplyRankCapByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) ApplyRankCapByUserIdAsync(
	request *ApplyRankCapByUserIdRequest,
	callback chan<- ApplyRankCapByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "applyRankCapByUserId",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
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

	go p.applyRankCapByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) ApplyRankCapByUserId(
	request *ApplyRankCapByUserIdRequest,
) (*ApplyRankCapByUserIdResult, error) {
	callback := make(chan ApplyRankCapByUserIdAsyncResult, 1)
	go p.ApplyRankCapByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) deleteStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStatusByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) DeleteStatusByUserIdAsync(
	request *DeleteStatusByUserIdRequest,
	callback chan<- DeleteStatusByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "deleteStatusByUserId",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
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

	go p.deleteStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) DeleteStatusByUserId(
	request *DeleteStatusByUserIdRequest,
) (*DeleteStatusByUserIdResult, error) {
	callback := make(chan DeleteStatusByUserIdAsyncResult, 1)
	go p.DeleteStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) verifyGradeAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyGradeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyGradeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyGradeResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyGradeAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- VerifyGradeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) VerifyGradeAsync(
	request *VerifyGradeRequest,
	callback chan<- VerifyGradeAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "verifyGrade",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		bodies["verifyType"] = *request.VerifyType
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.GradeValue != nil {
		bodies["gradeValue"] = *request.GradeValue
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

	go p.verifyGradeAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) VerifyGrade(
	request *VerifyGradeRequest,
) (*VerifyGradeResult, error) {
	callback := make(chan VerifyGradeAsyncResult, 1)
	go p.VerifyGradeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) verifyGradeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyGradeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyGradeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyGradeByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyGradeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- VerifyGradeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) VerifyGradeByUserIdAsync(
	request *VerifyGradeByUserIdRequest,
	callback chan<- VerifyGradeByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "verifyGradeByUserId",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		bodies["verifyType"] = *request.VerifyType
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.GradeValue != nil {
		bodies["gradeValue"] = *request.GradeValue
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go p.verifyGradeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) VerifyGradeByUserId(
	request *VerifyGradeByUserIdRequest,
) (*VerifyGradeByUserIdResult, error) {
	callback := make(chan VerifyGradeByUserIdAsyncResult, 1)
	go p.VerifyGradeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) verifyGradeUpMaterialAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyGradeUpMaterialAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyGradeUpMaterialAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyGradeUpMaterialResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyGradeUpMaterialAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- VerifyGradeUpMaterialAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) VerifyGradeUpMaterialAsync(
	request *VerifyGradeUpMaterialRequest,
	callback chan<- VerifyGradeUpMaterialAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "verifyGradeUpMaterial",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		bodies["verifyType"] = *request.VerifyType
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.MaterialPropertyId != nil && *request.MaterialPropertyId != "" {
		bodies["materialPropertyId"] = *request.MaterialPropertyId
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

	go p.verifyGradeUpMaterialAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) VerifyGradeUpMaterial(
	request *VerifyGradeUpMaterialRequest,
) (*VerifyGradeUpMaterialResult, error) {
	callback := make(chan VerifyGradeUpMaterialAsyncResult, 1)
	go p.VerifyGradeUpMaterialAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) verifyGradeUpMaterialByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyGradeUpMaterialByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyGradeUpMaterialByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyGradeUpMaterialByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyGradeUpMaterialByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- VerifyGradeUpMaterialByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) VerifyGradeUpMaterialByUserIdAsync(
	request *VerifyGradeUpMaterialByUserIdRequest,
	callback chan<- VerifyGradeUpMaterialByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "verifyGradeUpMaterialByUserId",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		bodies["verifyType"] = *request.VerifyType
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.MaterialPropertyId != nil && *request.MaterialPropertyId != "" {
		bodies["materialPropertyId"] = *request.MaterialPropertyId
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

	go p.verifyGradeUpMaterialByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) VerifyGradeUpMaterialByUserId(
	request *VerifyGradeUpMaterialByUserIdRequest,
) (*VerifyGradeUpMaterialByUserIdResult, error) {
	callback := make(chan VerifyGradeUpMaterialByUserIdAsyncResult, 1)
	go p.VerifyGradeUpMaterialByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) addGradeByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddGradeByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddGradeByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddGradeByStampSheetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddGradeByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- AddGradeByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) AddGradeByStampSheetAsync(
	request *AddGradeByStampSheetRequest,
	callback chan<- AddGradeByStampSheetAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "addGradeByStampSheet",
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

	go p.addGradeByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) AddGradeByStampSheet(
	request *AddGradeByStampSheetRequest,
) (*AddGradeByStampSheetResult, error) {
	callback := make(chan AddGradeByStampSheetAsyncResult, 1)
	go p.AddGradeByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) applyRankCapByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ApplyRankCapByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ApplyRankCapByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ApplyRankCapByStampSheetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ApplyRankCapByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ApplyRankCapByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) ApplyRankCapByStampSheetAsync(
	request *ApplyRankCapByStampSheetRequest,
	callback chan<- ApplyRankCapByStampSheetAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "applyRankCapByStampSheet",
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

	go p.applyRankCapByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) ApplyRankCapByStampSheet(
	request *ApplyRankCapByStampSheetRequest,
) (*ApplyRankCapByStampSheetResult, error) {
	callback := make(chan ApplyRankCapByStampSheetAsyncResult, 1)
	go p.ApplyRankCapByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) subGradeByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SubGradeByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubGradeByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubGradeByStampTaskResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SubGradeByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- SubGradeByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) SubGradeByStampTaskAsync(
	request *SubGradeByStampTaskRequest,
	callback chan<- SubGradeByStampTaskAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "subGradeByStampTask",
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

	go p.subGradeByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) SubGradeByStampTask(
	request *SubGradeByStampTaskRequest,
) (*SubGradeByStampTaskResult, error) {
	callback := make(chan SubGradeByStampTaskAsyncResult, 1)
	go p.SubGradeByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) multiplyAcquireActionsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- MultiplyAcquireActionsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MultiplyAcquireActionsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MultiplyAcquireActionsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MultiplyAcquireActionsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- MultiplyAcquireActionsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) MultiplyAcquireActionsByUserIdAsync(
	request *MultiplyAcquireActionsByUserIdRequest,
	callback chan<- MultiplyAcquireActionsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "multiplyAcquireActionsByUserId",
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
	if request.GradeName != nil && *request.GradeName != "" {
		bodies["gradeName"] = *request.GradeName
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		bodies["propertyId"] = *request.PropertyId
	}
	if request.RateName != nil && *request.RateName != "" {
		bodies["rateName"] = *request.RateName
	}
	if request.AcquireActions != nil {
		var _acquireActions []interface{}
		for _, item := range request.AcquireActions {
			_acquireActions = append(_acquireActions, item)
		}
		bodies["acquireActions"] = _acquireActions
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

	go p.multiplyAcquireActionsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) MultiplyAcquireActionsByUserId(
	request *MultiplyAcquireActionsByUserIdRequest,
) (*MultiplyAcquireActionsByUserIdResult, error) {
	callback := make(chan MultiplyAcquireActionsByUserIdAsyncResult, 1)
	go p.MultiplyAcquireActionsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) multiplyAcquireActionsByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- MultiplyAcquireActionsByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- MultiplyAcquireActionsByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result MultiplyAcquireActionsByStampSheetResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- MultiplyAcquireActionsByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- MultiplyAcquireActionsByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) MultiplyAcquireActionsByStampSheetAsync(
	request *MultiplyAcquireActionsByStampSheetRequest,
	callback chan<- MultiplyAcquireActionsByStampSheetAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "multiplyAcquireActionsByStampSheet",
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

	go p.multiplyAcquireActionsByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) MultiplyAcquireActionsByStampSheet(
	request *MultiplyAcquireActionsByStampSheetRequest,
) (*MultiplyAcquireActionsByStampSheetResult, error) {
	callback := make(chan MultiplyAcquireActionsByStampSheetAsyncResult, 1)
	go p.MultiplyAcquireActionsByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) verifyGradeByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyGradeByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyGradeByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyGradeByStampTaskResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyGradeByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- VerifyGradeByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) VerifyGradeByStampTaskAsync(
	request *VerifyGradeByStampTaskRequest,
	callback chan<- VerifyGradeByStampTaskAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "verifyGradeByStampTask",
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

	go p.verifyGradeByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) VerifyGradeByStampTask(
	request *VerifyGradeByStampTaskRequest,
) (*VerifyGradeByStampTaskResult, error) {
	callback := make(chan VerifyGradeByStampTaskAsyncResult, 1)
	go p.VerifyGradeByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) verifyGradeUpMaterialByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyGradeUpMaterialByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyGradeUpMaterialByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyGradeUpMaterialByStampTaskResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyGradeUpMaterialByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- VerifyGradeUpMaterialByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) VerifyGradeUpMaterialByStampTaskAsync(
	request *VerifyGradeUpMaterialByStampTaskRequest,
	callback chan<- VerifyGradeUpMaterialByStampTaskAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "status",
			"function":    "verifyGradeUpMaterialByStampTask",
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

	go p.verifyGradeUpMaterialByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) VerifyGradeUpMaterialByStampTask(
	request *VerifyGradeUpMaterialByStampTaskRequest,
) (*VerifyGradeUpMaterialByStampTaskResult, error) {
	callback := make(chan VerifyGradeUpMaterialByStampTaskAsyncResult, 1)
	go p.VerifyGradeUpMaterialByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) exportMasterAsyncHandler(
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

func (p Gs2GradeWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "currentGradeMaster",
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

func (p Gs2GradeWebSocketClient) ExportMaster(
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

func (p Gs2GradeWebSocketClient) getCurrentGradeMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentGradeMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentGradeMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentGradeMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentGradeMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetCurrentGradeMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) GetCurrentGradeMasterAsync(
	request *GetCurrentGradeMasterRequest,
	callback chan<- GetCurrentGradeMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "currentGradeMaster",
			"function":    "getCurrentGradeMaster",
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

	go p.getCurrentGradeMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) GetCurrentGradeMaster(
	request *GetCurrentGradeMasterRequest,
) (*GetCurrentGradeMasterResult, error) {
	callback := make(chan GetCurrentGradeMasterAsyncResult, 1)
	go p.GetCurrentGradeMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) updateCurrentGradeMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentGradeMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentGradeMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentGradeMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentGradeMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateCurrentGradeMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) UpdateCurrentGradeMasterAsync(
	request *UpdateCurrentGradeMasterRequest,
	callback chan<- UpdateCurrentGradeMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "currentGradeMaster",
			"function":    "updateCurrentGradeMaster",
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

	go p.updateCurrentGradeMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) UpdateCurrentGradeMaster(
	request *UpdateCurrentGradeMasterRequest,
) (*UpdateCurrentGradeMasterResult, error) {
	callback := make(chan UpdateCurrentGradeMasterAsyncResult, 1)
	go p.UpdateCurrentGradeMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2GradeWebSocketClient) updateCurrentGradeMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentGradeMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentGradeMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentGradeMasterFromGitHubResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentGradeMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateCurrentGradeMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GradeWebSocketClient) UpdateCurrentGradeMasterFromGitHubAsync(
	request *UpdateCurrentGradeMasterFromGitHubRequest,
	callback chan<- UpdateCurrentGradeMasterFromGitHubAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "grade",
			"component":   "currentGradeMaster",
			"function":    "updateCurrentGradeMasterFromGitHub",
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

	go p.updateCurrentGradeMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2GradeWebSocketClient) UpdateCurrentGradeMasterFromGitHub(
	request *UpdateCurrentGradeMasterFromGitHubRequest,
) (*UpdateCurrentGradeMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentGradeMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentGradeMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
