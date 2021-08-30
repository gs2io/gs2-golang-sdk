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

package datastore

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2DatastoreWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2DatastoreWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2DatastoreWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2DatastoreWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
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

func (p Gs2DatastoreWebSocketClient) DescribeNamespaces(
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

func (p Gs2DatastoreWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2DatastoreWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
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
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.DoneUploadScript != nil {
		bodies["doneUploadScript"] = request.DoneUploadScript.ToDict()
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

func (p Gs2DatastoreWebSocketClient) CreateNamespace(
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

func (p Gs2DatastoreWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2DatastoreWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
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

func (p Gs2DatastoreWebSocketClient) GetNamespaceStatus(
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

func (p Gs2DatastoreWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2DatastoreWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
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

func (p Gs2DatastoreWebSocketClient) GetNamespace(
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

func (p Gs2DatastoreWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2DatastoreWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
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
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.DoneUploadScript != nil {
		bodies["doneUploadScript"] = request.DoneUploadScript.ToDict()
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

func (p Gs2DatastoreWebSocketClient) UpdateNamespace(
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

func (p Gs2DatastoreWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2DatastoreWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
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

func (p Gs2DatastoreWebSocketClient) DeleteNamespace(
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

func (p Gs2DatastoreWebSocketClient) describeDataObjectsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeDataObjectsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDataObjectsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDataObjectsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDataObjectsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeDataObjectsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) DescribeDataObjectsAsync(
	request *DescribeDataObjectsRequest,
	callback chan<- DescribeDataObjectsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "describeDataObjects",
			"contentType": "application/json",
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
	if request.Status != nil && *request.Status != "" {
		bodies["status"] = *request.Status
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

	go p.describeDataObjectsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) DescribeDataObjects(
	request *DescribeDataObjectsRequest,
) (*DescribeDataObjectsResult, error) {
	callback := make(chan DescribeDataObjectsAsyncResult, 1)
	go p.DescribeDataObjectsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) describeDataObjectsByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeDataObjectsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDataObjectsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDataObjectsByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDataObjectsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeDataObjectsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) DescribeDataObjectsByUserIdAsync(
	request *DescribeDataObjectsByUserIdRequest,
	callback chan<- DescribeDataObjectsByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "describeDataObjectsByUserId",
			"contentType": "application/json",
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
	if request.Status != nil && *request.Status != "" {
		bodies["status"] = *request.Status
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

	go p.describeDataObjectsByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) DescribeDataObjectsByUserId(
	request *DescribeDataObjectsByUserIdRequest,
) (*DescribeDataObjectsByUserIdResult, error) {
	callback := make(chan DescribeDataObjectsByUserIdAsyncResult, 1)
	go p.DescribeDataObjectsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) prepareUploadAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareUploadAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareUploadAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareUploadResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareUploadAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareUploadAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) PrepareUploadAsync(
	request *PrepareUploadRequest,
	callback chan<- PrepareUploadAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "prepareUpload",
			"contentType": "application/json",
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
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.ContentType != nil && *request.ContentType != "" {
		bodies["contentType"] = *request.ContentType
	}
	if request.Scope != nil && *request.Scope != "" {
		bodies["scope"] = *request.Scope
	}
	if request.AllowUserIds != nil {
		var _allowUserIds []interface{}
		for _, item := range request.AllowUserIds {
			_allowUserIds = append(_allowUserIds, item)
		}
		bodies["allowUserIds"] = _allowUserIds
	}
	if request.UpdateIfExists != nil {
		bodies["updateIfExists"] = *request.UpdateIfExists
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.prepareUploadAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) PrepareUpload(
	request *PrepareUploadRequest,
) (*PrepareUploadResult, error) {
	callback := make(chan PrepareUploadAsyncResult, 1)
	go p.PrepareUploadAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) prepareUploadByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareUploadByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareUploadByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareUploadByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareUploadByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareUploadByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) PrepareUploadByUserIdAsync(
	request *PrepareUploadByUserIdRequest,
	callback chan<- PrepareUploadByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "prepareUploadByUserId",
			"contentType": "application/json",
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
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.ContentType != nil && *request.ContentType != "" {
		bodies["contentType"] = *request.ContentType
	}
	if request.Scope != nil && *request.Scope != "" {
		bodies["scope"] = *request.Scope
	}
	if request.AllowUserIds != nil {
		var _allowUserIds []interface{}
		for _, item := range request.AllowUserIds {
			_allowUserIds = append(_allowUserIds, item)
		}
		bodies["allowUserIds"] = _allowUserIds
	}
	if request.UpdateIfExists != nil {
		bodies["updateIfExists"] = *request.UpdateIfExists
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.prepareUploadByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) PrepareUploadByUserId(
	request *PrepareUploadByUserIdRequest,
) (*PrepareUploadByUserIdResult, error) {
	callback := make(chan PrepareUploadByUserIdAsyncResult, 1)
	go p.PrepareUploadByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) updateDataObjectAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateDataObjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateDataObjectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateDataObjectResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateDataObjectAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateDataObjectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) UpdateDataObjectAsync(
	request *UpdateDataObjectRequest,
	callback chan<- UpdateDataObjectAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "updateDataObject",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.Scope != nil && *request.Scope != "" {
		bodies["scope"] = *request.Scope
	}
	if request.AllowUserIds != nil {
		var _allowUserIds []interface{}
		for _, item := range request.AllowUserIds {
			_allowUserIds = append(_allowUserIds, item)
		}
		bodies["allowUserIds"] = _allowUserIds
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.updateDataObjectAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) UpdateDataObject(
	request *UpdateDataObjectRequest,
) (*UpdateDataObjectResult, error) {
	callback := make(chan UpdateDataObjectAsyncResult, 1)
	go p.UpdateDataObjectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) updateDataObjectByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateDataObjectByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateDataObjectByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateDataObjectByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateDataObjectByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateDataObjectByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) UpdateDataObjectByUserIdAsync(
	request *UpdateDataObjectByUserIdRequest,
	callback chan<- UpdateDataObjectByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "updateDataObjectByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Scope != nil && *request.Scope != "" {
		bodies["scope"] = *request.Scope
	}
	if request.AllowUserIds != nil {
		var _allowUserIds []interface{}
		for _, item := range request.AllowUserIds {
			_allowUserIds = append(_allowUserIds, item)
		}
		bodies["allowUserIds"] = _allowUserIds
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.updateDataObjectByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) UpdateDataObjectByUserId(
	request *UpdateDataObjectByUserIdRequest,
) (*UpdateDataObjectByUserIdResult, error) {
	callback := make(chan UpdateDataObjectByUserIdAsyncResult, 1)
	go p.UpdateDataObjectByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) prepareReUploadAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareReUploadAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareReUploadAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareReUploadResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareReUploadAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareReUploadAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) PrepareReUploadAsync(
	request *PrepareReUploadRequest,
	callback chan<- PrepareReUploadAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "prepareReUpload",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.AccessToken != nil && *request.AccessToken != "" {
		bodies["accessToken"] = *request.AccessToken
	}
	if request.ContentType != nil && *request.ContentType != "" {
		bodies["contentType"] = *request.ContentType
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.prepareReUploadAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) PrepareReUpload(
	request *PrepareReUploadRequest,
) (*PrepareReUploadResult, error) {
	callback := make(chan PrepareReUploadAsyncResult, 1)
	go p.PrepareReUploadAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) prepareReUploadByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareReUploadByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareReUploadByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareReUploadByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareReUploadByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareReUploadByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) PrepareReUploadByUserIdAsync(
	request *PrepareReUploadByUserIdRequest,
	callback chan<- PrepareReUploadByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "prepareReUploadByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.ContentType != nil && *request.ContentType != "" {
		bodies["contentType"] = *request.ContentType
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.prepareReUploadByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) PrepareReUploadByUserId(
	request *PrepareReUploadByUserIdRequest,
) (*PrepareReUploadByUserIdResult, error) {
	callback := make(chan PrepareReUploadByUserIdAsyncResult, 1)
	go p.PrepareReUploadByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) doneUploadAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DoneUploadAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoneUploadAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoneUploadResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DoneUploadAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DoneUploadAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) DoneUploadAsync(
	request *DoneUploadRequest,
	callback chan<- DoneUploadAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "doneUpload",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
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

	go p.doneUploadAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) DoneUpload(
	request *DoneUploadRequest,
) (*DoneUploadResult, error) {
	callback := make(chan DoneUploadAsyncResult, 1)
	go p.DoneUploadAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) doneUploadByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DoneUploadByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoneUploadByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoneUploadByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DoneUploadByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DoneUploadByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) DoneUploadByUserIdAsync(
	request *DoneUploadByUserIdRequest,
	callback chan<- DoneUploadByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "doneUploadByUserId",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.doneUploadByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) DoneUploadByUserId(
	request *DoneUploadByUserIdRequest,
) (*DoneUploadByUserIdResult, error) {
	callback := make(chan DoneUploadByUserIdAsyncResult, 1)
	go p.DoneUploadByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) deleteDataObjectAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteDataObjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteDataObjectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteDataObjectResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteDataObjectAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteDataObjectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) DeleteDataObjectAsync(
	request *DeleteDataObjectRequest,
	callback chan<- DeleteDataObjectAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "deleteDataObject",
			"contentType": "application/json",
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
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.deleteDataObjectAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) DeleteDataObject(
	request *DeleteDataObjectRequest,
) (*DeleteDataObjectResult, error) {
	callback := make(chan DeleteDataObjectAsyncResult, 1)
	go p.DeleteDataObjectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) deleteDataObjectByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteDataObjectByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteDataObjectByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteDataObjectByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteDataObjectByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteDataObjectByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) DeleteDataObjectByUserIdAsync(
	request *DeleteDataObjectByUserIdRequest,
	callback chan<- DeleteDataObjectByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "deleteDataObjectByUserId",
			"contentType": "application/json",
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
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.deleteDataObjectByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) DeleteDataObjectByUserId(
	request *DeleteDataObjectByUserIdRequest,
) (*DeleteDataObjectByUserIdResult, error) {
	callback := make(chan DeleteDataObjectByUserIdAsyncResult, 1)
	go p.DeleteDataObjectByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) prepareDownloadAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareDownloadAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareDownloadAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareDownloadAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadAsync(
	request *PrepareDownloadRequest,
	callback chan<- PrepareDownloadAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "prepareDownload",
			"contentType": "application/json",
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
	if request.DataObjectId != nil && *request.DataObjectId != "" {
		bodies["dataObjectId"] = *request.DataObjectId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.prepareDownloadAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) PrepareDownload(
	request *PrepareDownloadRequest,
) (*PrepareDownloadResult, error) {
	callback := make(chan PrepareDownloadAsyncResult, 1)
	go p.PrepareDownloadAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) prepareDownloadByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareDownloadByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareDownloadByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareDownloadByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadByUserIdAsync(
	request *PrepareDownloadByUserIdRequest,
	callback chan<- PrepareDownloadByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "prepareDownloadByUserId",
			"contentType": "application/json",
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
	if request.DataObjectId != nil && *request.DataObjectId != "" {
		bodies["dataObjectId"] = *request.DataObjectId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.prepareDownloadByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadByUserId(
	request *PrepareDownloadByUserIdRequest,
) (*PrepareDownloadByUserIdResult, error) {
	callback := make(chan PrepareDownloadByUserIdAsyncResult, 1)
	go p.PrepareDownloadByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) prepareDownloadByGenerationAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareDownloadByGenerationAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadByGenerationAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadByGenerationResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareDownloadByGenerationAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareDownloadByGenerationAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadByGenerationAsync(
	request *PrepareDownloadByGenerationRequest,
	callback chan<- PrepareDownloadByGenerationAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "prepareDownloadByGeneration",
			"contentType": "application/json",
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
	if request.DataObjectId != nil && *request.DataObjectId != "" {
		bodies["dataObjectId"] = *request.DataObjectId
	}
	if request.Generation != nil && *request.Generation != "" {
		bodies["generation"] = *request.Generation
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.prepareDownloadByGenerationAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadByGeneration(
	request *PrepareDownloadByGenerationRequest,
) (*PrepareDownloadByGenerationResult, error) {
	callback := make(chan PrepareDownloadByGenerationAsyncResult, 1)
	go p.PrepareDownloadByGenerationAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) prepareDownloadByGenerationAndUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareDownloadByGenerationAndUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadByGenerationAndUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadByGenerationAndUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareDownloadByGenerationAndUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareDownloadByGenerationAndUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadByGenerationAndUserIdAsync(
	request *PrepareDownloadByGenerationAndUserIdRequest,
	callback chan<- PrepareDownloadByGenerationAndUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "prepareDownloadByGenerationAndUserId",
			"contentType": "application/json",
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
	if request.DataObjectId != nil && *request.DataObjectId != "" {
		bodies["dataObjectId"] = *request.DataObjectId
	}
	if request.Generation != nil && *request.Generation != "" {
		bodies["generation"] = *request.Generation
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.prepareDownloadByGenerationAndUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadByGenerationAndUserId(
	request *PrepareDownloadByGenerationAndUserIdRequest,
) (*PrepareDownloadByGenerationAndUserIdResult, error) {
	callback := make(chan PrepareDownloadByGenerationAndUserIdAsyncResult, 1)
	go p.PrepareDownloadByGenerationAndUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) prepareDownloadOwnDataAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareDownloadOwnDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadOwnDataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadOwnDataResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareDownloadOwnDataAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareDownloadOwnDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadOwnDataAsync(
	request *PrepareDownloadOwnDataRequest,
	callback chan<- PrepareDownloadOwnDataAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "prepareDownloadOwnData",
			"contentType": "application/json",
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
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.prepareDownloadOwnDataAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadOwnData(
	request *PrepareDownloadOwnDataRequest,
) (*PrepareDownloadOwnDataResult, error) {
	callback := make(chan PrepareDownloadOwnDataAsyncResult, 1)
	go p.PrepareDownloadOwnDataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) prepareDownloadByUserIdAndDataObjectNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareDownloadByUserIdAndDataObjectNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadByUserIdAndDataObjectNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadByUserIdAndDataObjectNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareDownloadByUserIdAndDataObjectNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareDownloadByUserIdAndDataObjectNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadByUserIdAndDataObjectNameAsync(
	request *PrepareDownloadByUserIdAndDataObjectNameRequest,
	callback chan<- PrepareDownloadByUserIdAndDataObjectNameAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "prepareDownloadByUserIdAndDataObjectName",
			"contentType": "application/json",
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
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.prepareDownloadByUserIdAndDataObjectNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadByUserIdAndDataObjectName(
	request *PrepareDownloadByUserIdAndDataObjectNameRequest,
) (*PrepareDownloadByUserIdAndDataObjectNameResult, error) {
	callback := make(chan PrepareDownloadByUserIdAndDataObjectNameAsyncResult, 1)
	go p.PrepareDownloadByUserIdAndDataObjectNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) prepareDownloadOwnDataByGenerationAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareDownloadOwnDataByGenerationAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadOwnDataByGenerationAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadOwnDataByGenerationResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareDownloadOwnDataByGenerationAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareDownloadOwnDataByGenerationAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadOwnDataByGenerationAsync(
	request *PrepareDownloadOwnDataByGenerationRequest,
	callback chan<- PrepareDownloadOwnDataByGenerationAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "prepareDownloadOwnDataByGeneration",
			"contentType": "application/json",
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
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.Generation != nil && *request.Generation != "" {
		bodies["generation"] = *request.Generation
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.prepareDownloadOwnDataByGenerationAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadOwnDataByGeneration(
	request *PrepareDownloadOwnDataByGenerationRequest,
) (*PrepareDownloadOwnDataByGenerationResult, error) {
	callback := make(chan PrepareDownloadOwnDataByGenerationAsyncResult, 1)
	go p.PrepareDownloadOwnDataByGenerationAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) prepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsync(
	request *PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest,
	callback chan<- PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "prepareDownloadByUserIdAndDataObjectNameAndGeneration",
			"contentType": "application/json",
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
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.Generation != nil && *request.Generation != "" {
		bodies["generation"] = *request.Generation
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.prepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) PrepareDownloadByUserIdAndDataObjectNameAndGeneration(
	request *PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest,
) (*PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult, error) {
	callback := make(chan PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult, 1)
	go p.PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) restoreDataObjectAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- RestoreDataObjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RestoreDataObjectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RestoreDataObjectResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RestoreDataObjectAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RestoreDataObjectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) RestoreDataObjectAsync(
	request *RestoreDataObjectRequest,
	callback chan<- RestoreDataObjectAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObject",
			"function":    "restoreDataObject",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.DataObjectId != nil && *request.DataObjectId != "" {
		bodies["dataObjectId"] = *request.DataObjectId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.restoreDataObjectAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) RestoreDataObject(
	request *RestoreDataObjectRequest,
) (*RestoreDataObjectResult, error) {
	callback := make(chan RestoreDataObjectAsyncResult, 1)
	go p.RestoreDataObjectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) describeDataObjectHistoriesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeDataObjectHistoriesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDataObjectHistoriesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDataObjectHistoriesResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDataObjectHistoriesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeDataObjectHistoriesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) DescribeDataObjectHistoriesAsync(
	request *DescribeDataObjectHistoriesRequest,
	callback chan<- DescribeDataObjectHistoriesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObjectHistory",
			"function":    "describeDataObjectHistories",
			"contentType": "application/json",
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
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
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

	go p.describeDataObjectHistoriesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) DescribeDataObjectHistories(
	request *DescribeDataObjectHistoriesRequest,
) (*DescribeDataObjectHistoriesResult, error) {
	callback := make(chan DescribeDataObjectHistoriesAsyncResult, 1)
	go p.DescribeDataObjectHistoriesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) describeDataObjectHistoriesByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeDataObjectHistoriesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDataObjectHistoriesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDataObjectHistoriesByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDataObjectHistoriesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeDataObjectHistoriesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) DescribeDataObjectHistoriesByUserIdAsync(
	request *DescribeDataObjectHistoriesByUserIdRequest,
	callback chan<- DescribeDataObjectHistoriesByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObjectHistory",
			"function":    "describeDataObjectHistoriesByUserId",
			"contentType": "application/json",
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
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
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

	go p.describeDataObjectHistoriesByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) DescribeDataObjectHistoriesByUserId(
	request *DescribeDataObjectHistoriesByUserIdRequest,
) (*DescribeDataObjectHistoriesByUserIdResult, error) {
	callback := make(chan DescribeDataObjectHistoriesByUserIdAsyncResult, 1)
	go p.DescribeDataObjectHistoriesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) getDataObjectHistoryAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetDataObjectHistoryAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDataObjectHistoryAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDataObjectHistoryResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDataObjectHistoryAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetDataObjectHistoryAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) GetDataObjectHistoryAsync(
	request *GetDataObjectHistoryRequest,
	callback chan<- GetDataObjectHistoryAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObjectHistory",
			"function":    "getDataObjectHistory",
			"contentType": "application/json",
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
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.Generation != nil && *request.Generation != "" {
		bodies["generation"] = *request.Generation
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.AccessToken != nil {
		bodies["xGs2AccessToken"] = string(*request.AccessToken)
	}

	go p.getDataObjectHistoryAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) GetDataObjectHistory(
	request *GetDataObjectHistoryRequest,
) (*GetDataObjectHistoryResult, error) {
	callback := make(chan GetDataObjectHistoryAsyncResult, 1)
	go p.GetDataObjectHistoryAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2DatastoreWebSocketClient) getDataObjectHistoryByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetDataObjectHistoryByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDataObjectHistoryByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDataObjectHistoryByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDataObjectHistoryByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetDataObjectHistoryByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreWebSocketClient) GetDataObjectHistoryByUserIdAsync(
	request *GetDataObjectHistoryByUserIdRequest,
	callback chan<- GetDataObjectHistoryByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "datastore",
			"component":   "dataObjectHistory",
			"function":    "getDataObjectHistoryByUserId",
			"contentType": "application/json",
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
	if request.DataObjectName != nil && *request.DataObjectName != "" {
		bodies["dataObjectName"] = *request.DataObjectName
	}
	if request.Generation != nil && *request.Generation != "" {
		bodies["generation"] = *request.Generation
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	go p.getDataObjectHistoryByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreWebSocketClient) GetDataObjectHistoryByUserId(
	request *GetDataObjectHistoryByUserIdRequest,
) (*GetDataObjectHistoryByUserIdResult, error) {
	callback := make(chan GetDataObjectHistoryByUserIdAsyncResult, 1)
	go p.GetDataObjectHistoryByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
