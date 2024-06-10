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

package ranking2

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2Ranking2WebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2Ranking2WebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2Ranking2WebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) DescribeNamespaces(
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

func (p Gs2Ranking2WebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) CreateNamespace(
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

func (p Gs2Ranking2WebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) GetNamespaceStatus(
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

func (p Gs2Ranking2WebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) GetNamespace(
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

func (p Gs2Ranking2WebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) UpdateNamespace(
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

func (p Gs2Ranking2WebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) DeleteNamespace(
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

func (p Gs2Ranking2WebSocketClient) dumpUserDataByUserIdAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) DumpUserDataByUserIdAsync(
	request *DumpUserDataByUserIdRequest,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) DumpUserDataByUserId(
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

func (p Gs2Ranking2WebSocketClient) checkDumpUserDataByUserIdAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) CheckDumpUserDataByUserIdAsync(
	request *CheckDumpUserDataByUserIdRequest,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) CheckDumpUserDataByUserId(
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

func (p Gs2Ranking2WebSocketClient) cleanUserDataByUserIdAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) CleanUserDataByUserIdAsync(
	request *CleanUserDataByUserIdRequest,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) CleanUserDataByUserId(
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

func (p Gs2Ranking2WebSocketClient) checkCleanUserDataByUserIdAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) CheckCleanUserDataByUserIdAsync(
	request *CheckCleanUserDataByUserIdRequest,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) CheckCleanUserDataByUserId(
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

func (p Gs2Ranking2WebSocketClient) prepareImportUserDataByUserIdAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) PrepareImportUserDataByUserIdAsync(
	request *PrepareImportUserDataByUserIdRequest,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) PrepareImportUserDataByUserId(
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

func (p Gs2Ranking2WebSocketClient) importUserDataByUserIdAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) ImportUserDataByUserIdAsync(
	request *ImportUserDataByUserIdRequest,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) ImportUserDataByUserId(
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

func (p Gs2Ranking2WebSocketClient) checkImportUserDataByUserIdAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) CheckImportUserDataByUserIdAsync(
	request *CheckImportUserDataByUserIdRequest,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
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

func (p Gs2Ranking2WebSocketClient) CheckImportUserDataByUserId(
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

func (p Gs2Ranking2WebSocketClient) describeGlobalRankingModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGlobalRankingModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeGlobalRankingModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingModelsAsync(
	request *DescribeGlobalRankingModelsRequest,
	callback chan<- DescribeGlobalRankingModelsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingModel",
			"function":    "describeGlobalRankingModels",
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

	go p.describeGlobalRankingModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingModels(
	request *DescribeGlobalRankingModelsRequest,
) (*DescribeGlobalRankingModelsResult, error) {
	callback := make(chan DescribeGlobalRankingModelsAsyncResult, 1)
	go p.DescribeGlobalRankingModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getGlobalRankingModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGlobalRankingModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGlobalRankingModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingModelAsync(
	request *GetGlobalRankingModelRequest,
	callback chan<- GetGlobalRankingModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingModel",
			"function":    "getGlobalRankingModel",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getGlobalRankingModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingModel(
	request *GetGlobalRankingModelRequest,
) (*GetGlobalRankingModelResult, error) {
	callback := make(chan GetGlobalRankingModelAsyncResult, 1)
	go p.GetGlobalRankingModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeGlobalRankingModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGlobalRankingModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeGlobalRankingModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingModelMastersAsync(
	request *DescribeGlobalRankingModelMastersRequest,
	callback chan<- DescribeGlobalRankingModelMastersAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingModelMaster",
			"function":    "describeGlobalRankingModelMasters",
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

	go p.describeGlobalRankingModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingModelMasters(
	request *DescribeGlobalRankingModelMastersRequest,
) (*DescribeGlobalRankingModelMastersResult, error) {
	callback := make(chan DescribeGlobalRankingModelMastersAsyncResult, 1)
	go p.DescribeGlobalRankingModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) createGlobalRankingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateGlobalRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGlobalRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGlobalRankingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGlobalRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateGlobalRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) CreateGlobalRankingModelMasterAsync(
	request *CreateGlobalRankingModelMasterRequest,
	callback chan<- CreateGlobalRankingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingModelMaster",
			"function":    "createGlobalRankingModelMaster",
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
	if request.MinimumValue != nil {
		bodies["minimumValue"] = *request.MinimumValue
	}
	if request.MaximumValue != nil {
		bodies["maximumValue"] = *request.MaximumValue
	}
	if request.Sum != nil {
		bodies["sum"] = *request.Sum
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
	}
	if request.RankingRewards != nil {
		var _rankingRewards []interface{}
		for _, item := range request.RankingRewards {
			_rankingRewards = append(_rankingRewards, item)
		}
		bodies["rankingRewards"] = _rankingRewards
	}
	if request.EntryPeriodEventId != nil && *request.EntryPeriodEventId != "" {
		bodies["entryPeriodEventId"] = *request.EntryPeriodEventId
	}
	if request.AccessPeriodEventId != nil && *request.AccessPeriodEventId != "" {
		bodies["accessPeriodEventId"] = *request.AccessPeriodEventId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.createGlobalRankingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) CreateGlobalRankingModelMaster(
	request *CreateGlobalRankingModelMasterRequest,
) (*CreateGlobalRankingModelMasterResult, error) {
	callback := make(chan CreateGlobalRankingModelMasterAsyncResult, 1)
	go p.CreateGlobalRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getGlobalRankingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGlobalRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGlobalRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingModelMasterAsync(
	request *GetGlobalRankingModelMasterRequest,
	callback chan<- GetGlobalRankingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingModelMaster",
			"function":    "getGlobalRankingModelMaster",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getGlobalRankingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingModelMaster(
	request *GetGlobalRankingModelMasterRequest,
) (*GetGlobalRankingModelMasterResult, error) {
	callback := make(chan GetGlobalRankingModelMasterAsyncResult, 1)
	go p.GetGlobalRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) updateGlobalRankingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateGlobalRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGlobalRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGlobalRankingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateGlobalRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateGlobalRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) UpdateGlobalRankingModelMasterAsync(
	request *UpdateGlobalRankingModelMasterRequest,
	callback chan<- UpdateGlobalRankingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingModelMaster",
			"function":    "updateGlobalRankingModelMaster",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.MinimumValue != nil {
		bodies["minimumValue"] = *request.MinimumValue
	}
	if request.MaximumValue != nil {
		bodies["maximumValue"] = *request.MaximumValue
	}
	if request.Sum != nil {
		bodies["sum"] = *request.Sum
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
	}
	if request.RankingRewards != nil {
		var _rankingRewards []interface{}
		for _, item := range request.RankingRewards {
			_rankingRewards = append(_rankingRewards, item)
		}
		bodies["rankingRewards"] = _rankingRewards
	}
	if request.EntryPeriodEventId != nil && *request.EntryPeriodEventId != "" {
		bodies["entryPeriodEventId"] = *request.EntryPeriodEventId
	}
	if request.AccessPeriodEventId != nil && *request.AccessPeriodEventId != "" {
		bodies["accessPeriodEventId"] = *request.AccessPeriodEventId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateGlobalRankingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) UpdateGlobalRankingModelMaster(
	request *UpdateGlobalRankingModelMasterRequest,
) (*UpdateGlobalRankingModelMasterResult, error) {
	callback := make(chan UpdateGlobalRankingModelMasterAsyncResult, 1)
	go p.UpdateGlobalRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) deleteGlobalRankingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteGlobalRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGlobalRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGlobalRankingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGlobalRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteGlobalRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DeleteGlobalRankingModelMasterAsync(
	request *DeleteGlobalRankingModelMasterRequest,
	callback chan<- DeleteGlobalRankingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingModelMaster",
			"function":    "deleteGlobalRankingModelMaster",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteGlobalRankingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DeleteGlobalRankingModelMaster(
	request *DeleteGlobalRankingModelMasterRequest,
) (*DeleteGlobalRankingModelMasterResult, error) {
	callback := make(chan DeleteGlobalRankingModelMasterAsyncResult, 1)
	go p.DeleteGlobalRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeGlobalRankingScoresAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGlobalRankingScoresAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingScoresAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingScoresResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingScoresAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeGlobalRankingScoresAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingScoresAsync(
	request *DescribeGlobalRankingScoresRequest,
	callback chan<- DescribeGlobalRankingScoresAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingScore",
			"function":    "describeGlobalRankingScores",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
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

	go p.describeGlobalRankingScoresAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingScores(
	request *DescribeGlobalRankingScoresRequest,
) (*DescribeGlobalRankingScoresResult, error) {
	callback := make(chan DescribeGlobalRankingScoresAsyncResult, 1)
	go p.DescribeGlobalRankingScoresAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeGlobalRankingScoresByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGlobalRankingScoresByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingScoresByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingScoresByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingScoresByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeGlobalRankingScoresByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingScoresByUserIdAsync(
	request *DescribeGlobalRankingScoresByUserIdRequest,
	callback chan<- DescribeGlobalRankingScoresByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingScore",
			"function":    "describeGlobalRankingScoresByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
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

	go p.describeGlobalRankingScoresByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingScoresByUserId(
	request *DescribeGlobalRankingScoresByUserIdRequest,
) (*DescribeGlobalRankingScoresByUserIdResult, error) {
	callback := make(chan DescribeGlobalRankingScoresByUserIdAsyncResult, 1)
	go p.DescribeGlobalRankingScoresByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) putGlobalRankingScoreAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PutGlobalRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutGlobalRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutGlobalRankingScoreResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutGlobalRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PutGlobalRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) PutGlobalRankingScoreAsync(
	request *PutGlobalRankingScoreRequest,
	callback chan<- PutGlobalRankingScoreAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingScore",
			"function":    "putGlobalRankingScore",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Score != nil {
		bodies["score"] = *request.Score
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

	go p.putGlobalRankingScoreAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) PutGlobalRankingScore(
	request *PutGlobalRankingScoreRequest,
) (*PutGlobalRankingScoreResult, error) {
	callback := make(chan PutGlobalRankingScoreAsyncResult, 1)
	go p.PutGlobalRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) putGlobalRankingScoreByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PutGlobalRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutGlobalRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutGlobalRankingScoreByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutGlobalRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PutGlobalRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) PutGlobalRankingScoreByUserIdAsync(
	request *PutGlobalRankingScoreByUserIdRequest,
	callback chan<- PutGlobalRankingScoreByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingScore",
			"function":    "putGlobalRankingScoreByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Score != nil {
		bodies["score"] = *request.Score
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

	go p.putGlobalRankingScoreByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) PutGlobalRankingScoreByUserId(
	request *PutGlobalRankingScoreByUserIdRequest,
) (*PutGlobalRankingScoreByUserIdResult, error) {
	callback := make(chan PutGlobalRankingScoreByUserIdAsyncResult, 1)
	go p.PutGlobalRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getGlobalRankingScoreAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGlobalRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingScoreResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGlobalRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingScoreAsync(
	request *GetGlobalRankingScoreRequest,
	callback chan<- GetGlobalRankingScoreAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingScore",
			"function":    "getGlobalRankingScore",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getGlobalRankingScoreAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingScore(
	request *GetGlobalRankingScoreRequest,
) (*GetGlobalRankingScoreResult, error) {
	callback := make(chan GetGlobalRankingScoreAsyncResult, 1)
	go p.GetGlobalRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getGlobalRankingScoreByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGlobalRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingScoreByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGlobalRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingScoreByUserIdAsync(
	request *GetGlobalRankingScoreByUserIdRequest,
	callback chan<- GetGlobalRankingScoreByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingScore",
			"function":    "getGlobalRankingScoreByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getGlobalRankingScoreByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingScoreByUserId(
	request *GetGlobalRankingScoreByUserIdRequest,
) (*GetGlobalRankingScoreByUserIdResult, error) {
	callback := make(chan GetGlobalRankingScoreByUserIdAsyncResult, 1)
	go p.GetGlobalRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) deleteGlobalRankingScoreByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteGlobalRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGlobalRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGlobalRankingScoreByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGlobalRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteGlobalRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DeleteGlobalRankingScoreByUserIdAsync(
	request *DeleteGlobalRankingScoreByUserIdRequest,
	callback chan<- DeleteGlobalRankingScoreByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingScore",
			"function":    "deleteGlobalRankingScoreByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.deleteGlobalRankingScoreByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DeleteGlobalRankingScoreByUserId(
	request *DeleteGlobalRankingScoreByUserIdRequest,
) (*DeleteGlobalRankingScoreByUserIdResult, error) {
	callback := make(chan DeleteGlobalRankingScoreByUserIdAsyncResult, 1)
	go p.DeleteGlobalRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeGlobalRankingReceivedRewardsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGlobalRankingReceivedRewardsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingReceivedRewardsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingReceivedRewardsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingReceivedRewardsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeGlobalRankingReceivedRewardsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingReceivedRewardsAsync(
	request *DescribeGlobalRankingReceivedRewardsRequest,
	callback chan<- DescribeGlobalRankingReceivedRewardsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingReceivedReward",
			"function":    "describeGlobalRankingReceivedRewards",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.describeGlobalRankingReceivedRewardsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingReceivedRewards(
	request *DescribeGlobalRankingReceivedRewardsRequest,
) (*DescribeGlobalRankingReceivedRewardsResult, error) {
	callback := make(chan DescribeGlobalRankingReceivedRewardsAsyncResult, 1)
	go p.DescribeGlobalRankingReceivedRewardsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeGlobalRankingReceivedRewardsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingReceivedRewardsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingReceivedRewardsByUserIdAsync(
	request *DescribeGlobalRankingReceivedRewardsByUserIdRequest,
	callback chan<- DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingReceivedReward",
			"function":    "describeGlobalRankingReceivedRewardsByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.describeGlobalRankingReceivedRewardsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingReceivedRewardsByUserId(
	request *DescribeGlobalRankingReceivedRewardsByUserIdRequest,
) (*DescribeGlobalRankingReceivedRewardsByUserIdResult, error) {
	callback := make(chan DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult, 1)
	go p.DescribeGlobalRankingReceivedRewardsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) createGlobalRankingReceivedRewardAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateGlobalRankingReceivedRewardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGlobalRankingReceivedRewardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGlobalRankingReceivedRewardResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGlobalRankingReceivedRewardAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateGlobalRankingReceivedRewardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) CreateGlobalRankingReceivedRewardAsync(
	request *CreateGlobalRankingReceivedRewardRequest,
	callback chan<- CreateGlobalRankingReceivedRewardAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingReceivedReward",
			"function":    "createGlobalRankingReceivedReward",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.createGlobalRankingReceivedRewardAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) CreateGlobalRankingReceivedReward(
	request *CreateGlobalRankingReceivedRewardRequest,
) (*CreateGlobalRankingReceivedRewardResult, error) {
	callback := make(chan CreateGlobalRankingReceivedRewardAsyncResult, 1)
	go p.CreateGlobalRankingReceivedRewardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) createGlobalRankingReceivedRewardByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGlobalRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGlobalRankingReceivedRewardByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGlobalRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateGlobalRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) CreateGlobalRankingReceivedRewardByUserIdAsync(
	request *CreateGlobalRankingReceivedRewardByUserIdRequest,
	callback chan<- CreateGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingReceivedReward",
			"function":    "createGlobalRankingReceivedRewardByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.createGlobalRankingReceivedRewardByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) CreateGlobalRankingReceivedRewardByUserId(
	request *CreateGlobalRankingReceivedRewardByUserIdRequest,
) (*CreateGlobalRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan CreateGlobalRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.CreateGlobalRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) receiveGlobalRankingReceivedRewardAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReceiveGlobalRankingReceivedRewardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveGlobalRankingReceivedRewardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveGlobalRankingReceivedRewardResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveGlobalRankingReceivedRewardAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ReceiveGlobalRankingReceivedRewardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) ReceiveGlobalRankingReceivedRewardAsync(
	request *ReceiveGlobalRankingReceivedRewardRequest,
	callback chan<- ReceiveGlobalRankingReceivedRewardAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingReceivedReward",
			"function":    "receiveGlobalRankingReceivedReward",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.receiveGlobalRankingReceivedRewardAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) ReceiveGlobalRankingReceivedReward(
	request *ReceiveGlobalRankingReceivedRewardRequest,
) (*ReceiveGlobalRankingReceivedRewardResult, error) {
	callback := make(chan ReceiveGlobalRankingReceivedRewardAsyncResult, 1)
	go p.ReceiveGlobalRankingReceivedRewardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) receiveGlobalRankingReceivedRewardByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveGlobalRankingReceivedRewardByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) ReceiveGlobalRankingReceivedRewardByUserIdAsync(
	request *ReceiveGlobalRankingReceivedRewardByUserIdRequest,
	callback chan<- ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingReceivedReward",
			"function":    "receiveGlobalRankingReceivedRewardByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.receiveGlobalRankingReceivedRewardByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) ReceiveGlobalRankingReceivedRewardByUserId(
	request *ReceiveGlobalRankingReceivedRewardByUserIdRequest,
) (*ReceiveGlobalRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.ReceiveGlobalRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getGlobalRankingReceivedRewardAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGlobalRankingReceivedRewardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingReceivedRewardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingReceivedRewardResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingReceivedRewardAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGlobalRankingReceivedRewardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingReceivedRewardAsync(
	request *GetGlobalRankingReceivedRewardRequest,
	callback chan<- GetGlobalRankingReceivedRewardAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingReceivedReward",
			"function":    "getGlobalRankingReceivedReward",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getGlobalRankingReceivedRewardAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingReceivedReward(
	request *GetGlobalRankingReceivedRewardRequest,
) (*GetGlobalRankingReceivedRewardResult, error) {
	callback := make(chan GetGlobalRankingReceivedRewardAsyncResult, 1)
	go p.GetGlobalRankingReceivedRewardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getGlobalRankingReceivedRewardByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingReceivedRewardByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGlobalRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingReceivedRewardByUserIdAsync(
	request *GetGlobalRankingReceivedRewardByUserIdRequest,
	callback chan<- GetGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingReceivedReward",
			"function":    "getGlobalRankingReceivedRewardByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getGlobalRankingReceivedRewardByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingReceivedRewardByUserId(
	request *GetGlobalRankingReceivedRewardByUserIdRequest,
) (*GetGlobalRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan GetGlobalRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.GetGlobalRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) deleteGlobalRankingReceivedRewardByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGlobalRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGlobalRankingReceivedRewardByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGlobalRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteGlobalRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DeleteGlobalRankingReceivedRewardByUserIdAsync(
	request *DeleteGlobalRankingReceivedRewardByUserIdRequest,
	callback chan<- DeleteGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingReceivedReward",
			"function":    "deleteGlobalRankingReceivedRewardByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.deleteGlobalRankingReceivedRewardByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DeleteGlobalRankingReceivedRewardByUserId(
	request *DeleteGlobalRankingReceivedRewardByUserIdRequest,
) (*DeleteGlobalRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan DeleteGlobalRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.DeleteGlobalRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) createGlobalRankingReceivedRewardByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateGlobalRankingReceivedRewardByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGlobalRankingReceivedRewardByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGlobalRankingReceivedRewardByStampTaskResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGlobalRankingReceivedRewardByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateGlobalRankingReceivedRewardByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) CreateGlobalRankingReceivedRewardByStampTaskAsync(
	request *CreateGlobalRankingReceivedRewardByStampTaskRequest,
	callback chan<- CreateGlobalRankingReceivedRewardByStampTaskAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingReceivedReward",
			"function":    "createGlobalRankingReceivedRewardByStampTask",
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

	go p.createGlobalRankingReceivedRewardByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) CreateGlobalRankingReceivedRewardByStampTask(
	request *CreateGlobalRankingReceivedRewardByStampTaskRequest,
) (*CreateGlobalRankingReceivedRewardByStampTaskResult, error) {
	callback := make(chan CreateGlobalRankingReceivedRewardByStampTaskAsyncResult, 1)
	go p.CreateGlobalRankingReceivedRewardByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeGlobalRankingsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGlobalRankingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeGlobalRankingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingsAsync(
	request *DescribeGlobalRankingsRequest,
	callback chan<- DescribeGlobalRankingsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingData",
			"function":    "describeGlobalRankings",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.describeGlobalRankingsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankings(
	request *DescribeGlobalRankingsRequest,
) (*DescribeGlobalRankingsResult, error) {
	callback := make(chan DescribeGlobalRankingsAsyncResult, 1)
	go p.DescribeGlobalRankingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeGlobalRankingsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeGlobalRankingsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeGlobalRankingsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingsByUserIdAsync(
	request *DescribeGlobalRankingsByUserIdRequest,
	callback chan<- DescribeGlobalRankingsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingData",
			"function":    "describeGlobalRankingsByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.describeGlobalRankingsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeGlobalRankingsByUserId(
	request *DescribeGlobalRankingsByUserIdRequest,
) (*DescribeGlobalRankingsByUserIdResult, error) {
	callback := make(chan DescribeGlobalRankingsByUserIdAsyncResult, 1)
	go p.DescribeGlobalRankingsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getGlobalRankingAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGlobalRankingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGlobalRankingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingAsync(
	request *GetGlobalRankingRequest,
	callback chan<- GetGlobalRankingAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingData",
			"function":    "getGlobalRanking",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getGlobalRankingAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetGlobalRanking(
	request *GetGlobalRankingRequest,
) (*GetGlobalRankingResult, error) {
	callback := make(chan GetGlobalRankingAsyncResult, 1)
	go p.GetGlobalRankingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getGlobalRankingByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGlobalRankingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGlobalRankingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingByUserIdAsync(
	request *GetGlobalRankingByUserIdRequest,
	callback chan<- GetGlobalRankingByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "globalRankingData",
			"function":    "getGlobalRankingByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getGlobalRankingByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetGlobalRankingByUserId(
	request *GetGlobalRankingByUserIdRequest,
) (*GetGlobalRankingByUserIdResult, error) {
	callback := make(chan GetGlobalRankingByUserIdAsyncResult, 1)
	go p.GetGlobalRankingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeClusterRankingModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeClusterRankingModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeClusterRankingModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingModelsAsync(
	request *DescribeClusterRankingModelsRequest,
	callback chan<- DescribeClusterRankingModelsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingModel",
			"function":    "describeClusterRankingModels",
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

	go p.describeClusterRankingModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingModels(
	request *DescribeClusterRankingModelsRequest,
) (*DescribeClusterRankingModelsResult, error) {
	callback := make(chan DescribeClusterRankingModelsAsyncResult, 1)
	go p.DescribeClusterRankingModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getClusterRankingModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetClusterRankingModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetClusterRankingModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingModelAsync(
	request *GetClusterRankingModelRequest,
	callback chan<- GetClusterRankingModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingModel",
			"function":    "getClusterRankingModel",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getClusterRankingModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingModel(
	request *GetClusterRankingModelRequest,
) (*GetClusterRankingModelResult, error) {
	callback := make(chan GetClusterRankingModelAsyncResult, 1)
	go p.GetClusterRankingModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeClusterRankingModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeClusterRankingModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeClusterRankingModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingModelMastersAsync(
	request *DescribeClusterRankingModelMastersRequest,
	callback chan<- DescribeClusterRankingModelMastersAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingModelMaster",
			"function":    "describeClusterRankingModelMasters",
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

	go p.describeClusterRankingModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingModelMasters(
	request *DescribeClusterRankingModelMastersRequest,
) (*DescribeClusterRankingModelMastersResult, error) {
	callback := make(chan DescribeClusterRankingModelMastersAsyncResult, 1)
	go p.DescribeClusterRankingModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) createClusterRankingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateClusterRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateClusterRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateClusterRankingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateClusterRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateClusterRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) CreateClusterRankingModelMasterAsync(
	request *CreateClusterRankingModelMasterRequest,
	callback chan<- CreateClusterRankingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingModelMaster",
			"function":    "createClusterRankingModelMaster",
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
	if request.ClusterType != nil && *request.ClusterType != "" {
		bodies["clusterType"] = *request.ClusterType
	}
	if request.MinimumValue != nil {
		bodies["minimumValue"] = *request.MinimumValue
	}
	if request.MaximumValue != nil {
		bodies["maximumValue"] = *request.MaximumValue
	}
	if request.Sum != nil {
		bodies["sum"] = *request.Sum
	}
	if request.ScoreTtlDays != nil {
		bodies["scoreTtlDays"] = *request.ScoreTtlDays
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
	}
	if request.RankingRewards != nil {
		var _rankingRewards []interface{}
		for _, item := range request.RankingRewards {
			_rankingRewards = append(_rankingRewards, item)
		}
		bodies["rankingRewards"] = _rankingRewards
	}
	if request.EntryPeriodEventId != nil && *request.EntryPeriodEventId != "" {
		bodies["entryPeriodEventId"] = *request.EntryPeriodEventId
	}
	if request.AccessPeriodEventId != nil && *request.AccessPeriodEventId != "" {
		bodies["accessPeriodEventId"] = *request.AccessPeriodEventId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.createClusterRankingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) CreateClusterRankingModelMaster(
	request *CreateClusterRankingModelMasterRequest,
) (*CreateClusterRankingModelMasterResult, error) {
	callback := make(chan CreateClusterRankingModelMasterAsyncResult, 1)
	go p.CreateClusterRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getClusterRankingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetClusterRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetClusterRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingModelMasterAsync(
	request *GetClusterRankingModelMasterRequest,
	callback chan<- GetClusterRankingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingModelMaster",
			"function":    "getClusterRankingModelMaster",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getClusterRankingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingModelMaster(
	request *GetClusterRankingModelMasterRequest,
) (*GetClusterRankingModelMasterResult, error) {
	callback := make(chan GetClusterRankingModelMasterAsyncResult, 1)
	go p.GetClusterRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) updateClusterRankingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateClusterRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateClusterRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateClusterRankingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateClusterRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateClusterRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) UpdateClusterRankingModelMasterAsync(
	request *UpdateClusterRankingModelMasterRequest,
	callback chan<- UpdateClusterRankingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingModelMaster",
			"function":    "updateClusterRankingModelMaster",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ClusterType != nil && *request.ClusterType != "" {
		bodies["clusterType"] = *request.ClusterType
	}
	if request.MinimumValue != nil {
		bodies["minimumValue"] = *request.MinimumValue
	}
	if request.MaximumValue != nil {
		bodies["maximumValue"] = *request.MaximumValue
	}
	if request.Sum != nil {
		bodies["sum"] = *request.Sum
	}
	if request.ScoreTtlDays != nil {
		bodies["scoreTtlDays"] = *request.ScoreTtlDays
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
	}
	if request.RankingRewards != nil {
		var _rankingRewards []interface{}
		for _, item := range request.RankingRewards {
			_rankingRewards = append(_rankingRewards, item)
		}
		bodies["rankingRewards"] = _rankingRewards
	}
	if request.EntryPeriodEventId != nil && *request.EntryPeriodEventId != "" {
		bodies["entryPeriodEventId"] = *request.EntryPeriodEventId
	}
	if request.AccessPeriodEventId != nil && *request.AccessPeriodEventId != "" {
		bodies["accessPeriodEventId"] = *request.AccessPeriodEventId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateClusterRankingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) UpdateClusterRankingModelMaster(
	request *UpdateClusterRankingModelMasterRequest,
) (*UpdateClusterRankingModelMasterResult, error) {
	callback := make(chan UpdateClusterRankingModelMasterAsyncResult, 1)
	go p.UpdateClusterRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) deleteClusterRankingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteClusterRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteClusterRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteClusterRankingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteClusterRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteClusterRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DeleteClusterRankingModelMasterAsync(
	request *DeleteClusterRankingModelMasterRequest,
	callback chan<- DeleteClusterRankingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingModelMaster",
			"function":    "deleteClusterRankingModelMaster",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteClusterRankingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DeleteClusterRankingModelMaster(
	request *DeleteClusterRankingModelMasterRequest,
) (*DeleteClusterRankingModelMasterResult, error) {
	callback := make(chan DeleteClusterRankingModelMasterAsyncResult, 1)
	go p.DeleteClusterRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeClusterRankingScoresAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeClusterRankingScoresAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingScoresAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingScoresResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingScoresAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeClusterRankingScoresAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingScoresAsync(
	request *DescribeClusterRankingScoresRequest,
	callback chan<- DescribeClusterRankingScoresAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingScore",
			"function":    "describeClusterRankingScores",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.describeClusterRankingScoresAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingScores(
	request *DescribeClusterRankingScoresRequest,
) (*DescribeClusterRankingScoresResult, error) {
	callback := make(chan DescribeClusterRankingScoresAsyncResult, 1)
	go p.DescribeClusterRankingScoresAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeClusterRankingScoresByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeClusterRankingScoresByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingScoresByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingScoresByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingScoresByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeClusterRankingScoresByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingScoresByUserIdAsync(
	request *DescribeClusterRankingScoresByUserIdRequest,
	callback chan<- DescribeClusterRankingScoresByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingScore",
			"function":    "describeClusterRankingScoresByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.describeClusterRankingScoresByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingScoresByUserId(
	request *DescribeClusterRankingScoresByUserIdRequest,
) (*DescribeClusterRankingScoresByUserIdResult, error) {
	callback := make(chan DescribeClusterRankingScoresByUserIdAsyncResult, 1)
	go p.DescribeClusterRankingScoresByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) putClusterRankingScoreAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PutClusterRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutClusterRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutClusterRankingScoreResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutClusterRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PutClusterRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) PutClusterRankingScoreAsync(
	request *PutClusterRankingScoreRequest,
	callback chan<- PutClusterRankingScoreAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingScore",
			"function":    "putClusterRankingScore",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Score != nil {
		bodies["score"] = *request.Score
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

	go p.putClusterRankingScoreAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) PutClusterRankingScore(
	request *PutClusterRankingScoreRequest,
) (*PutClusterRankingScoreResult, error) {
	callback := make(chan PutClusterRankingScoreAsyncResult, 1)
	go p.PutClusterRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) putClusterRankingScoreByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PutClusterRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutClusterRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutClusterRankingScoreByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutClusterRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PutClusterRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) PutClusterRankingScoreByUserIdAsync(
	request *PutClusterRankingScoreByUserIdRequest,
	callback chan<- PutClusterRankingScoreByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingScore",
			"function":    "putClusterRankingScoreByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Score != nil {
		bodies["score"] = *request.Score
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

	go p.putClusterRankingScoreByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) PutClusterRankingScoreByUserId(
	request *PutClusterRankingScoreByUserIdRequest,
) (*PutClusterRankingScoreByUserIdResult, error) {
	callback := make(chan PutClusterRankingScoreByUserIdAsyncResult, 1)
	go p.PutClusterRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getClusterRankingScoreAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetClusterRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingScoreResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetClusterRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingScoreAsync(
	request *GetClusterRankingScoreRequest,
	callback chan<- GetClusterRankingScoreAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingScore",
			"function":    "getClusterRankingScore",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getClusterRankingScoreAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingScore(
	request *GetClusterRankingScoreRequest,
) (*GetClusterRankingScoreResult, error) {
	callback := make(chan GetClusterRankingScoreAsyncResult, 1)
	go p.GetClusterRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getClusterRankingScoreByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetClusterRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingScoreByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetClusterRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingScoreByUserIdAsync(
	request *GetClusterRankingScoreByUserIdRequest,
	callback chan<- GetClusterRankingScoreByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingScore",
			"function":    "getClusterRankingScoreByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getClusterRankingScoreByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingScoreByUserId(
	request *GetClusterRankingScoreByUserIdRequest,
) (*GetClusterRankingScoreByUserIdResult, error) {
	callback := make(chan GetClusterRankingScoreByUserIdAsyncResult, 1)
	go p.GetClusterRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) deleteClusterRankingScoreByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteClusterRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteClusterRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteClusterRankingScoreByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteClusterRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteClusterRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DeleteClusterRankingScoreByUserIdAsync(
	request *DeleteClusterRankingScoreByUserIdRequest,
	callback chan<- DeleteClusterRankingScoreByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingScore",
			"function":    "deleteClusterRankingScoreByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.deleteClusterRankingScoreByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DeleteClusterRankingScoreByUserId(
	request *DeleteClusterRankingScoreByUserIdRequest,
) (*DeleteClusterRankingScoreByUserIdResult, error) {
	callback := make(chan DeleteClusterRankingScoreByUserIdAsyncResult, 1)
	go p.DeleteClusterRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeClusterRankingReceivedRewardsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeClusterRankingReceivedRewardsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingReceivedRewardsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingReceivedRewardsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingReceivedRewardsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeClusterRankingReceivedRewardsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingReceivedRewardsAsync(
	request *DescribeClusterRankingReceivedRewardsRequest,
	callback chan<- DescribeClusterRankingReceivedRewardsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingReceivedReward",
			"function":    "describeClusterRankingReceivedRewards",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.describeClusterRankingReceivedRewardsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingReceivedRewards(
	request *DescribeClusterRankingReceivedRewardsRequest,
) (*DescribeClusterRankingReceivedRewardsResult, error) {
	callback := make(chan DescribeClusterRankingReceivedRewardsAsyncResult, 1)
	go p.DescribeClusterRankingReceivedRewardsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeClusterRankingReceivedRewardsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeClusterRankingReceivedRewardsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingReceivedRewardsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingReceivedRewardsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingReceivedRewardsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeClusterRankingReceivedRewardsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingReceivedRewardsByUserIdAsync(
	request *DescribeClusterRankingReceivedRewardsByUserIdRequest,
	callback chan<- DescribeClusterRankingReceivedRewardsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingReceivedReward",
			"function":    "describeClusterRankingReceivedRewardsByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.describeClusterRankingReceivedRewardsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingReceivedRewardsByUserId(
	request *DescribeClusterRankingReceivedRewardsByUserIdRequest,
) (*DescribeClusterRankingReceivedRewardsByUserIdResult, error) {
	callback := make(chan DescribeClusterRankingReceivedRewardsByUserIdAsyncResult, 1)
	go p.DescribeClusterRankingReceivedRewardsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) createClusterRankingReceivedRewardAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateClusterRankingReceivedRewardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateClusterRankingReceivedRewardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateClusterRankingReceivedRewardResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateClusterRankingReceivedRewardAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateClusterRankingReceivedRewardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) CreateClusterRankingReceivedRewardAsync(
	request *CreateClusterRankingReceivedRewardRequest,
	callback chan<- CreateClusterRankingReceivedRewardAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingReceivedReward",
			"function":    "createClusterRankingReceivedReward",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.createClusterRankingReceivedRewardAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) CreateClusterRankingReceivedReward(
	request *CreateClusterRankingReceivedRewardRequest,
) (*CreateClusterRankingReceivedRewardResult, error) {
	callback := make(chan CreateClusterRankingReceivedRewardAsyncResult, 1)
	go p.CreateClusterRankingReceivedRewardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) createClusterRankingReceivedRewardByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateClusterRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateClusterRankingReceivedRewardByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateClusterRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateClusterRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) CreateClusterRankingReceivedRewardByUserIdAsync(
	request *CreateClusterRankingReceivedRewardByUserIdRequest,
	callback chan<- CreateClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingReceivedReward",
			"function":    "createClusterRankingReceivedRewardByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.createClusterRankingReceivedRewardByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) CreateClusterRankingReceivedRewardByUserId(
	request *CreateClusterRankingReceivedRewardByUserIdRequest,
) (*CreateClusterRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan CreateClusterRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.CreateClusterRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) receiveClusterRankingReceivedRewardAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReceiveClusterRankingReceivedRewardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveClusterRankingReceivedRewardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveClusterRankingReceivedRewardResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveClusterRankingReceivedRewardAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ReceiveClusterRankingReceivedRewardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) ReceiveClusterRankingReceivedRewardAsync(
	request *ReceiveClusterRankingReceivedRewardRequest,
	callback chan<- ReceiveClusterRankingReceivedRewardAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingReceivedReward",
			"function":    "receiveClusterRankingReceivedReward",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.receiveClusterRankingReceivedRewardAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) ReceiveClusterRankingReceivedReward(
	request *ReceiveClusterRankingReceivedRewardRequest,
) (*ReceiveClusterRankingReceivedRewardResult, error) {
	callback := make(chan ReceiveClusterRankingReceivedRewardAsyncResult, 1)
	go p.ReceiveClusterRankingReceivedRewardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) receiveClusterRankingReceivedRewardByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ReceiveClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveClusterRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveClusterRankingReceivedRewardByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveClusterRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ReceiveClusterRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) ReceiveClusterRankingReceivedRewardByUserIdAsync(
	request *ReceiveClusterRankingReceivedRewardByUserIdRequest,
	callback chan<- ReceiveClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingReceivedReward",
			"function":    "receiveClusterRankingReceivedRewardByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.receiveClusterRankingReceivedRewardByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) ReceiveClusterRankingReceivedRewardByUserId(
	request *ReceiveClusterRankingReceivedRewardByUserIdRequest,
) (*ReceiveClusterRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan ReceiveClusterRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.ReceiveClusterRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getClusterRankingReceivedRewardAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetClusterRankingReceivedRewardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingReceivedRewardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingReceivedRewardResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingReceivedRewardAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetClusterRankingReceivedRewardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingReceivedRewardAsync(
	request *GetClusterRankingReceivedRewardRequest,
	callback chan<- GetClusterRankingReceivedRewardAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingReceivedReward",
			"function":    "getClusterRankingReceivedReward",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getClusterRankingReceivedRewardAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingReceivedReward(
	request *GetClusterRankingReceivedRewardRequest,
) (*GetClusterRankingReceivedRewardResult, error) {
	callback := make(chan GetClusterRankingReceivedRewardAsyncResult, 1)
	go p.GetClusterRankingReceivedRewardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getClusterRankingReceivedRewardByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingReceivedRewardByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetClusterRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingReceivedRewardByUserIdAsync(
	request *GetClusterRankingReceivedRewardByUserIdRequest,
	callback chan<- GetClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingReceivedReward",
			"function":    "getClusterRankingReceivedRewardByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getClusterRankingReceivedRewardByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingReceivedRewardByUserId(
	request *GetClusterRankingReceivedRewardByUserIdRequest,
) (*GetClusterRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan GetClusterRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.GetClusterRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) deleteClusterRankingReceivedRewardByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteClusterRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteClusterRankingReceivedRewardByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteClusterRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteClusterRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DeleteClusterRankingReceivedRewardByUserIdAsync(
	request *DeleteClusterRankingReceivedRewardByUserIdRequest,
	callback chan<- DeleteClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingReceivedReward",
			"function":    "deleteClusterRankingReceivedRewardByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.deleteClusterRankingReceivedRewardByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DeleteClusterRankingReceivedRewardByUserId(
	request *DeleteClusterRankingReceivedRewardByUserIdRequest,
) (*DeleteClusterRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan DeleteClusterRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.DeleteClusterRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) createClusterRankingReceivedRewardByStampTaskAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateClusterRankingReceivedRewardByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateClusterRankingReceivedRewardByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateClusterRankingReceivedRewardByStampTaskResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateClusterRankingReceivedRewardByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateClusterRankingReceivedRewardByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) CreateClusterRankingReceivedRewardByStampTaskAsync(
	request *CreateClusterRankingReceivedRewardByStampTaskRequest,
	callback chan<- CreateClusterRankingReceivedRewardByStampTaskAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingReceivedReward",
			"function":    "createClusterRankingReceivedRewardByStampTask",
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

	go p.createClusterRankingReceivedRewardByStampTaskAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) CreateClusterRankingReceivedRewardByStampTask(
	request *CreateClusterRankingReceivedRewardByStampTaskRequest,
) (*CreateClusterRankingReceivedRewardByStampTaskResult, error) {
	callback := make(chan CreateClusterRankingReceivedRewardByStampTaskAsyncResult, 1)
	go p.CreateClusterRankingReceivedRewardByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeClusterRankingsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeClusterRankingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeClusterRankingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingsAsync(
	request *DescribeClusterRankingsRequest,
	callback chan<- DescribeClusterRankingsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingData",
			"function":    "describeClusterRankings",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.describeClusterRankingsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankings(
	request *DescribeClusterRankingsRequest,
) (*DescribeClusterRankingsResult, error) {
	callback := make(chan DescribeClusterRankingsAsyncResult, 1)
	go p.DescribeClusterRankingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeClusterRankingsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeClusterRankingsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeClusterRankingsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingsByUserIdAsync(
	request *DescribeClusterRankingsByUserIdRequest,
	callback chan<- DescribeClusterRankingsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingData",
			"function":    "describeClusterRankingsByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.describeClusterRankingsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeClusterRankingsByUserId(
	request *DescribeClusterRankingsByUserIdRequest,
) (*DescribeClusterRankingsByUserIdResult, error) {
	callback := make(chan DescribeClusterRankingsByUserIdAsyncResult, 1)
	go p.DescribeClusterRankingsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getClusterRankingAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetClusterRankingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetClusterRankingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingAsync(
	request *GetClusterRankingRequest,
	callback chan<- GetClusterRankingAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingData",
			"function":    "getClusterRanking",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getClusterRankingAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetClusterRanking(
	request *GetClusterRankingRequest,
) (*GetClusterRankingResult, error) {
	callback := make(chan GetClusterRankingAsyncResult, 1)
	go p.GetClusterRankingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getClusterRankingByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetClusterRankingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetClusterRankingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingByUserIdAsync(
	request *GetClusterRankingByUserIdRequest,
	callback chan<- GetClusterRankingByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "clusterRankingData",
			"function":    "getClusterRankingByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		bodies["clusterName"] = *request.ClusterName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getClusterRankingByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetClusterRankingByUserId(
	request *GetClusterRankingByUserIdRequest,
) (*GetClusterRankingByUserIdResult, error) {
	callback := make(chan GetClusterRankingByUserIdAsyncResult, 1)
	go p.GetClusterRankingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeSubscribeRankingModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSubscribeRankingModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribeRankingModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribeRankingModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribeRankingModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeSubscribeRankingModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribeRankingModelsAsync(
	request *DescribeSubscribeRankingModelsRequest,
	callback chan<- DescribeSubscribeRankingModelsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingModel",
			"function":    "describeSubscribeRankingModels",
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

	go p.describeSubscribeRankingModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribeRankingModels(
	request *DescribeSubscribeRankingModelsRequest,
) (*DescribeSubscribeRankingModelsResult, error) {
	callback := make(chan DescribeSubscribeRankingModelsAsyncResult, 1)
	go p.DescribeSubscribeRankingModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getSubscribeRankingModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSubscribeRankingModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeRankingModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeRankingModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeRankingModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetSubscribeRankingModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetSubscribeRankingModelAsync(
	request *GetSubscribeRankingModelRequest,
	callback chan<- GetSubscribeRankingModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingModel",
			"function":    "getSubscribeRankingModel",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getSubscribeRankingModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetSubscribeRankingModel(
	request *GetSubscribeRankingModelRequest,
) (*GetSubscribeRankingModelResult, error) {
	callback := make(chan GetSubscribeRankingModelAsyncResult, 1)
	go p.GetSubscribeRankingModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeSubscribeRankingModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSubscribeRankingModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribeRankingModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribeRankingModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribeRankingModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeSubscribeRankingModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribeRankingModelMastersAsync(
	request *DescribeSubscribeRankingModelMastersRequest,
	callback chan<- DescribeSubscribeRankingModelMastersAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingModelMaster",
			"function":    "describeSubscribeRankingModelMasters",
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

	go p.describeSubscribeRankingModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribeRankingModelMasters(
	request *DescribeSubscribeRankingModelMastersRequest,
) (*DescribeSubscribeRankingModelMastersResult, error) {
	callback := make(chan DescribeSubscribeRankingModelMastersAsyncResult, 1)
	go p.DescribeSubscribeRankingModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) createSubscribeRankingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateSubscribeRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateSubscribeRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateSubscribeRankingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateSubscribeRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateSubscribeRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) CreateSubscribeRankingModelMasterAsync(
	request *CreateSubscribeRankingModelMasterRequest,
	callback chan<- CreateSubscribeRankingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingModelMaster",
			"function":    "createSubscribeRankingModelMaster",
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
	if request.MinimumValue != nil {
		bodies["minimumValue"] = *request.MinimumValue
	}
	if request.MaximumValue != nil {
		bodies["maximumValue"] = *request.MaximumValue
	}
	if request.Sum != nil {
		bodies["sum"] = *request.Sum
	}
	if request.ScoreTtlDays != nil {
		bodies["scoreTtlDays"] = *request.ScoreTtlDays
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
	}
	if request.EntryPeriodEventId != nil && *request.EntryPeriodEventId != "" {
		bodies["entryPeriodEventId"] = *request.EntryPeriodEventId
	}
	if request.AccessPeriodEventId != nil && *request.AccessPeriodEventId != "" {
		bodies["accessPeriodEventId"] = *request.AccessPeriodEventId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.createSubscribeRankingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) CreateSubscribeRankingModelMaster(
	request *CreateSubscribeRankingModelMasterRequest,
) (*CreateSubscribeRankingModelMasterResult, error) {
	callback := make(chan CreateSubscribeRankingModelMasterAsyncResult, 1)
	go p.CreateSubscribeRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getSubscribeRankingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSubscribeRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeRankingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetSubscribeRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetSubscribeRankingModelMasterAsync(
	request *GetSubscribeRankingModelMasterRequest,
	callback chan<- GetSubscribeRankingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingModelMaster",
			"function":    "getSubscribeRankingModelMaster",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getSubscribeRankingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetSubscribeRankingModelMaster(
	request *GetSubscribeRankingModelMasterRequest,
) (*GetSubscribeRankingModelMasterResult, error) {
	callback := make(chan GetSubscribeRankingModelMasterAsyncResult, 1)
	go p.GetSubscribeRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) updateSubscribeRankingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateSubscribeRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateSubscribeRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateSubscribeRankingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateSubscribeRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateSubscribeRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) UpdateSubscribeRankingModelMasterAsync(
	request *UpdateSubscribeRankingModelMasterRequest,
	callback chan<- UpdateSubscribeRankingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingModelMaster",
			"function":    "updateSubscribeRankingModelMaster",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.MinimumValue != nil {
		bodies["minimumValue"] = *request.MinimumValue
	}
	if request.MaximumValue != nil {
		bodies["maximumValue"] = *request.MaximumValue
	}
	if request.Sum != nil {
		bodies["sum"] = *request.Sum
	}
	if request.ScoreTtlDays != nil {
		bodies["scoreTtlDays"] = *request.ScoreTtlDays
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
	}
	if request.EntryPeriodEventId != nil && *request.EntryPeriodEventId != "" {
		bodies["entryPeriodEventId"] = *request.EntryPeriodEventId
	}
	if request.AccessPeriodEventId != nil && *request.AccessPeriodEventId != "" {
		bodies["accessPeriodEventId"] = *request.AccessPeriodEventId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateSubscribeRankingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) UpdateSubscribeRankingModelMaster(
	request *UpdateSubscribeRankingModelMasterRequest,
) (*UpdateSubscribeRankingModelMasterResult, error) {
	callback := make(chan UpdateSubscribeRankingModelMasterAsyncResult, 1)
	go p.UpdateSubscribeRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) deleteSubscribeRankingModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteSubscribeRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSubscribeRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSubscribeRankingModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteSubscribeRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteSubscribeRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DeleteSubscribeRankingModelMasterAsync(
	request *DeleteSubscribeRankingModelMasterRequest,
	callback chan<- DeleteSubscribeRankingModelMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingModelMaster",
			"function":    "deleteSubscribeRankingModelMaster",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteSubscribeRankingModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DeleteSubscribeRankingModelMaster(
	request *DeleteSubscribeRankingModelMasterRequest,
) (*DeleteSubscribeRankingModelMasterResult, error) {
	callback := make(chan DeleteSubscribeRankingModelMasterAsyncResult, 1)
	go p.DeleteSubscribeRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeSubscribesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSubscribesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribesResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeSubscribesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribesAsync(
	request *DescribeSubscribesRequest,
	callback chan<- DescribeSubscribesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribe",
			"function":    "describeSubscribes",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
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

	go p.describeSubscribesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribes(
	request *DescribeSubscribesRequest,
) (*DescribeSubscribesResult, error) {
	callback := make(chan DescribeSubscribesAsyncResult, 1)
	go p.DescribeSubscribesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeSubscribesByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSubscribesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribesByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeSubscribesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribesByUserIdAsync(
	request *DescribeSubscribesByUserIdRequest,
	callback chan<- DescribeSubscribesByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribe",
			"function":    "describeSubscribesByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
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

	go p.describeSubscribesByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribesByUserId(
	request *DescribeSubscribesByUserIdRequest,
) (*DescribeSubscribesByUserIdResult, error) {
	callback := make(chan DescribeSubscribesByUserIdAsyncResult, 1)
	go p.DescribeSubscribesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) addSubscribeAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddSubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddSubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddSubscribeResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddSubscribeAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- AddSubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) AddSubscribeAsync(
	request *AddSubscribeRequest,
	callback chan<- AddSubscribeAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribe",
			"function":    "addSubscribe",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
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

	go p.addSubscribeAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) AddSubscribe(
	request *AddSubscribeRequest,
) (*AddSubscribeResult, error) {
	callback := make(chan AddSubscribeAsyncResult, 1)
	go p.AddSubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) addSubscribeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddSubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddSubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddSubscribeByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddSubscribeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- AddSubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) AddSubscribeByUserIdAsync(
	request *AddSubscribeByUserIdRequest,
	callback chan<- AddSubscribeByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribe",
			"function":    "addSubscribeByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
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

	go p.addSubscribeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) AddSubscribeByUserId(
	request *AddSubscribeByUserIdRequest,
) (*AddSubscribeByUserIdResult, error) {
	callback := make(chan AddSubscribeByUserIdAsyncResult, 1)
	go p.AddSubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeSubscribeRankingScoresAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSubscribeRankingScoresAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribeRankingScoresAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribeRankingScoresResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribeRankingScoresAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeSubscribeRankingScoresAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribeRankingScoresAsync(
	request *DescribeSubscribeRankingScoresRequest,
	callback chan<- DescribeSubscribeRankingScoresAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingScore",
			"function":    "describeSubscribeRankingScores",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
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

	go p.describeSubscribeRankingScoresAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribeRankingScores(
	request *DescribeSubscribeRankingScoresRequest,
) (*DescribeSubscribeRankingScoresResult, error) {
	callback := make(chan DescribeSubscribeRankingScoresAsyncResult, 1)
	go p.DescribeSubscribeRankingScoresAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeSubscribeRankingScoresByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSubscribeRankingScoresByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribeRankingScoresByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribeRankingScoresByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribeRankingScoresByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeSubscribeRankingScoresByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribeRankingScoresByUserIdAsync(
	request *DescribeSubscribeRankingScoresByUserIdRequest,
	callback chan<- DescribeSubscribeRankingScoresByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingScore",
			"function":    "describeSubscribeRankingScoresByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
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

	go p.describeSubscribeRankingScoresByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribeRankingScoresByUserId(
	request *DescribeSubscribeRankingScoresByUserIdRequest,
) (*DescribeSubscribeRankingScoresByUserIdResult, error) {
	callback := make(chan DescribeSubscribeRankingScoresByUserIdAsyncResult, 1)
	go p.DescribeSubscribeRankingScoresByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) putSubscribeRankingScoreAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PutSubscribeRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutSubscribeRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutSubscribeRankingScoreResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutSubscribeRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PutSubscribeRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) PutSubscribeRankingScoreAsync(
	request *PutSubscribeRankingScoreRequest,
	callback chan<- PutSubscribeRankingScoreAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingScore",
			"function":    "putSubscribeRankingScore",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Score != nil {
		bodies["score"] = *request.Score
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

	go p.putSubscribeRankingScoreAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) PutSubscribeRankingScore(
	request *PutSubscribeRankingScoreRequest,
) (*PutSubscribeRankingScoreResult, error) {
	callback := make(chan PutSubscribeRankingScoreAsyncResult, 1)
	go p.PutSubscribeRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) putSubscribeRankingScoreByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PutSubscribeRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutSubscribeRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutSubscribeRankingScoreByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutSubscribeRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PutSubscribeRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) PutSubscribeRankingScoreByUserIdAsync(
	request *PutSubscribeRankingScoreByUserIdRequest,
	callback chan<- PutSubscribeRankingScoreByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingScore",
			"function":    "putSubscribeRankingScoreByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Score != nil {
		bodies["score"] = *request.Score
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

	go p.putSubscribeRankingScoreByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) PutSubscribeRankingScoreByUserId(
	request *PutSubscribeRankingScoreByUserIdRequest,
) (*PutSubscribeRankingScoreByUserIdResult, error) {
	callback := make(chan PutSubscribeRankingScoreByUserIdAsyncResult, 1)
	go p.PutSubscribeRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getSubscribeRankingScoreAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSubscribeRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeRankingScoreResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetSubscribeRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetSubscribeRankingScoreAsync(
	request *GetSubscribeRankingScoreRequest,
	callback chan<- GetSubscribeRankingScoreAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingScore",
			"function":    "getSubscribeRankingScore",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getSubscribeRankingScoreAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetSubscribeRankingScore(
	request *GetSubscribeRankingScoreRequest,
) (*GetSubscribeRankingScoreResult, error) {
	callback := make(chan GetSubscribeRankingScoreAsyncResult, 1)
	go p.GetSubscribeRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getSubscribeRankingScoreByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSubscribeRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeRankingScoreByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetSubscribeRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetSubscribeRankingScoreByUserIdAsync(
	request *GetSubscribeRankingScoreByUserIdRequest,
	callback chan<- GetSubscribeRankingScoreByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingScore",
			"function":    "getSubscribeRankingScoreByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getSubscribeRankingScoreByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetSubscribeRankingScoreByUserId(
	request *GetSubscribeRankingScoreByUserIdRequest,
) (*GetSubscribeRankingScoreByUserIdResult, error) {
	callback := make(chan GetSubscribeRankingScoreByUserIdAsyncResult, 1)
	go p.GetSubscribeRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) deleteSubscribeRankingScoreByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteSubscribeRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSubscribeRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSubscribeRankingScoreByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteSubscribeRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteSubscribeRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DeleteSubscribeRankingScoreByUserIdAsync(
	request *DeleteSubscribeRankingScoreByUserIdRequest,
	callback chan<- DeleteSubscribeRankingScoreByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingScore",
			"function":    "deleteSubscribeRankingScoreByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.deleteSubscribeRankingScoreByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DeleteSubscribeRankingScoreByUserId(
	request *DeleteSubscribeRankingScoreByUserIdRequest,
) (*DeleteSubscribeRankingScoreByUserIdResult, error) {
	callback := make(chan DeleteSubscribeRankingScoreByUserIdAsyncResult, 1)
	go p.DeleteSubscribeRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeSubscribeRankingsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSubscribeRankingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribeRankingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribeRankingsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribeRankingsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeSubscribeRankingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribeRankingsAsync(
	request *DescribeSubscribeRankingsRequest,
	callback chan<- DescribeSubscribeRankingsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingData",
			"function":    "describeSubscribeRankings",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.describeSubscribeRankingsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribeRankings(
	request *DescribeSubscribeRankingsRequest,
) (*DescribeSubscribeRankingsResult, error) {
	callback := make(chan DescribeSubscribeRankingsAsyncResult, 1)
	go p.DescribeSubscribeRankingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) describeSubscribeRankingsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSubscribeRankingsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribeRankingsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribeRankingsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribeRankingsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeSubscribeRankingsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribeRankingsByUserIdAsync(
	request *DescribeSubscribeRankingsByUserIdRequest,
	callback chan<- DescribeSubscribeRankingsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingData",
			"function":    "describeSubscribeRankingsByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go p.describeSubscribeRankingsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DescribeSubscribeRankingsByUserId(
	request *DescribeSubscribeRankingsByUserIdRequest,
) (*DescribeSubscribeRankingsByUserIdResult, error) {
	callback := make(chan DescribeSubscribeRankingsByUserIdAsyncResult, 1)
	go p.DescribeSubscribeRankingsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getSubscribeRankingAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSubscribeRankingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeRankingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeRankingResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeRankingAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetSubscribeRankingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetSubscribeRankingAsync(
	request *GetSubscribeRankingRequest,
	callback chan<- GetSubscribeRankingAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingData",
			"function":    "getSubscribeRanking",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.ScorerUserId != nil && *request.ScorerUserId != "" {
		bodies["scorerUserId"] = *request.ScorerUserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getSubscribeRankingAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetSubscribeRanking(
	request *GetSubscribeRankingRequest,
) (*GetSubscribeRankingResult, error) {
	callback := make(chan GetSubscribeRankingAsyncResult, 1)
	go p.GetSubscribeRankingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getSubscribeRankingByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSubscribeRankingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeRankingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeRankingByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeRankingByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetSubscribeRankingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetSubscribeRankingByUserIdAsync(
	request *GetSubscribeRankingByUserIdRequest,
	callback chan<- GetSubscribeRankingByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeRankingData",
			"function":    "getSubscribeRankingByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.ScorerUserId != nil && *request.ScorerUserId != "" {
		bodies["scorerUserId"] = *request.ScorerUserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getSubscribeRankingByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetSubscribeRankingByUserId(
	request *GetSubscribeRankingByUserIdRequest,
) (*GetSubscribeRankingByUserIdResult, error) {
	callback := make(chan GetSubscribeRankingByUserIdAsyncResult, 1)
	go p.GetSubscribeRankingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) exportMasterAsyncHandler(
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

func (p Gs2Ranking2WebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "currentRankingMaster",
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

func (p Gs2Ranking2WebSocketClient) ExportMaster(
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

func (p Gs2Ranking2WebSocketClient) getCurrentRankingMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentRankingMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentRankingMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentRankingMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentRankingMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetCurrentRankingMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetCurrentRankingMasterAsync(
	request *GetCurrentRankingMasterRequest,
	callback chan<- GetCurrentRankingMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "currentRankingMaster",
			"function":    "getCurrentRankingMaster",
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

	go p.getCurrentRankingMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetCurrentRankingMaster(
	request *GetCurrentRankingMasterRequest,
) (*GetCurrentRankingMasterResult, error) {
	callback := make(chan GetCurrentRankingMasterAsyncResult, 1)
	go p.GetCurrentRankingMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) updateCurrentRankingMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentRankingMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentRankingMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentRankingMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentRankingMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateCurrentRankingMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) UpdateCurrentRankingMasterAsync(
	request *UpdateCurrentRankingMasterRequest,
	callback chan<- UpdateCurrentRankingMasterAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "currentRankingMaster",
			"function":    "updateCurrentRankingMaster",
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

	go p.updateCurrentRankingMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) UpdateCurrentRankingMaster(
	request *UpdateCurrentRankingMasterRequest,
) (*UpdateCurrentRankingMasterResult, error) {
	callback := make(chan UpdateCurrentRankingMasterAsyncResult, 1)
	go p.UpdateCurrentRankingMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) updateCurrentRankingMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentRankingMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentRankingMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentRankingMasterFromGitHubResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentRankingMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateCurrentRankingMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) UpdateCurrentRankingMasterFromGitHubAsync(
	request *UpdateCurrentRankingMasterFromGitHubRequest,
	callback chan<- UpdateCurrentRankingMasterFromGitHubAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "currentRankingMaster",
			"function":    "updateCurrentRankingMasterFromGitHub",
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

	go p.updateCurrentRankingMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) UpdateCurrentRankingMasterFromGitHub(
	request *UpdateCurrentRankingMasterFromGitHubRequest,
) (*UpdateCurrentRankingMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentRankingMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentRankingMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getSubscribeAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetSubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetSubscribeAsync(
	request *GetSubscribeRequest,
	callback chan<- GetSubscribeAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeUser",
			"function":    "getSubscribe",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
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

	go p.getSubscribeAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetSubscribe(
	request *GetSubscribeRequest,
) (*GetSubscribeResult, error) {
	callback := make(chan GetSubscribeAsyncResult, 1)
	go p.GetSubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) getSubscribeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetSubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) GetSubscribeByUserIdAsync(
	request *GetSubscribeByUserIdRequest,
	callback chan<- GetSubscribeByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeUser",
			"function":    "getSubscribeByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
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

	go p.getSubscribeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) GetSubscribeByUserId(
	request *GetSubscribeByUserIdRequest,
) (*GetSubscribeByUserIdResult, error) {
	callback := make(chan GetSubscribeByUserIdAsyncResult, 1)
	go p.GetSubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) deleteSubscribeAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteSubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSubscribeResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteSubscribeAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteSubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DeleteSubscribeAsync(
	request *DeleteSubscribeRequest,
	callback chan<- DeleteSubscribeAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeUser",
			"function":    "deleteSubscribe",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
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

	go p.deleteSubscribeAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DeleteSubscribe(
	request *DeleteSubscribeRequest,
) (*DeleteSubscribeResult, error) {
	callback := make(chan DeleteSubscribeAsyncResult, 1)
	go p.DeleteSubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2Ranking2WebSocketClient) deleteSubscribeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteSubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSubscribeByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteSubscribeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteSubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2WebSocketClient) DeleteSubscribeByUserIdAsync(
	request *DeleteSubscribeByUserIdRequest,
	callback chan<- DeleteSubscribeByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "ranking2",
			"component":   "subscribeUser",
			"function":    "deleteSubscribeByUserId",
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
	if request.RankingName != nil && *request.RankingName != "" {
		bodies["rankingName"] = *request.RankingName
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

	go p.deleteSubscribeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2WebSocketClient) DeleteSubscribeByUserId(
	request *DeleteSubscribeByUserIdRequest,
) (*DeleteSubscribeByUserIdResult, error) {
	callback := make(chan DeleteSubscribeByUserIdAsyncResult, 1)
	go p.DeleteSubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
