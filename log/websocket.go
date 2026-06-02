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

package log

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2LogWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2LogWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2LogWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2LogWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
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

func (p Gs2LogWebSocketClient) DescribeNamespaces(
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

func (p Gs2LogWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2LogWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
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
	if request.Type != nil && *request.Type != "" {
		bodies["type"] = *request.Type
	}
	if request.GcpCredentialJson != nil && *request.GcpCredentialJson != "" {
		bodies["gcpCredentialJson"] = *request.GcpCredentialJson
	}
	if request.BigQueryDatasetName != nil && *request.BigQueryDatasetName != "" {
		bodies["bigQueryDatasetName"] = *request.BigQueryDatasetName
	}
	if request.LogExpireDays != nil {
		bodies["logExpireDays"] = *request.LogExpireDays
	}
	if request.AwsRegion != nil && *request.AwsRegion != "" {
		bodies["awsRegion"] = *request.AwsRegion
	}
	if request.AwsAccessKeyId != nil && *request.AwsAccessKeyId != "" {
		bodies["awsAccessKeyId"] = *request.AwsAccessKeyId
	}
	if request.AwsSecretAccessKey != nil && *request.AwsSecretAccessKey != "" {
		bodies["awsSecretAccessKey"] = *request.AwsSecretAccessKey
	}
	if request.FirehoseStreamName != nil && *request.FirehoseStreamName != "" {
		bodies["firehoseStreamName"] = *request.FirehoseStreamName
	}
	if request.FirehoseCompressData != nil && *request.FirehoseCompressData != "" {
		bodies["firehoseCompressData"] = *request.FirehoseCompressData
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

func (p Gs2LogWebSocketClient) CreateNamespace(
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

func (p Gs2LogWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2LogWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
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

func (p Gs2LogWebSocketClient) GetNamespaceStatus(
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

func (p Gs2LogWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2LogWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
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

func (p Gs2LogWebSocketClient) GetNamespace(
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

func (p Gs2LogWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2LogWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
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
	if request.Type != nil && *request.Type != "" {
		bodies["type"] = *request.Type
	}
	if request.GcpCredentialJson != nil && *request.GcpCredentialJson != "" {
		bodies["gcpCredentialJson"] = *request.GcpCredentialJson
	}
	if request.BigQueryDatasetName != nil && *request.BigQueryDatasetName != "" {
		bodies["bigQueryDatasetName"] = *request.BigQueryDatasetName
	}
	if request.LogExpireDays != nil {
		bodies["logExpireDays"] = *request.LogExpireDays
	}
	if request.AwsRegion != nil && *request.AwsRegion != "" {
		bodies["awsRegion"] = *request.AwsRegion
	}
	if request.AwsAccessKeyId != nil && *request.AwsAccessKeyId != "" {
		bodies["awsAccessKeyId"] = *request.AwsAccessKeyId
	}
	if request.AwsSecretAccessKey != nil && *request.AwsSecretAccessKey != "" {
		bodies["awsSecretAccessKey"] = *request.AwsSecretAccessKey
	}
	if request.FirehoseStreamName != nil && *request.FirehoseStreamName != "" {
		bodies["firehoseStreamName"] = *request.FirehoseStreamName
	}
	if request.FirehoseCompressData != nil && *request.FirehoseCompressData != "" {
		bodies["firehoseCompressData"] = *request.FirehoseCompressData
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

func (p Gs2LogWebSocketClient) UpdateNamespace(
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

func (p Gs2LogWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2LogWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
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

func (p Gs2LogWebSocketClient) DeleteNamespace(
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

func (p Gs2LogWebSocketClient) getServiceVersionAsyncHandler(
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

func (p Gs2LogWebSocketClient) GetServiceVersionAsync(
	request *GetServiceVersionRequest,
	callback chan<- GetServiceVersionAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
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

func (p Gs2LogWebSocketClient) GetServiceVersion(
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

func (p Gs2LogWebSocketClient) queryAccessLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- QueryAccessLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- QueryAccessLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result QueryAccessLogResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryAccessLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- QueryAccessLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) QueryAccessLogAsync(
	request *QueryAccessLogRequest,
	callback chan<- QueryAccessLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "accessLog",
			"function":    "queryAccessLog",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Service != nil && *request.Service != "" {
		bodies["service"] = *request.Service
	}
	if request.Method != nil && *request.Method != "" {
		bodies["method"] = *request.Method
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.LongTerm != nil {
		bodies["longTerm"] = *request.LongTerm
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

	go p.queryAccessLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) QueryAccessLog(
	request *QueryAccessLogRequest,
) (*QueryAccessLogResult, error) {
	callback := make(chan QueryAccessLogAsyncResult, 1)
	go p.QueryAccessLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) countAccessLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CountAccessLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CountAccessLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CountAccessLogResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountAccessLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CountAccessLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) CountAccessLogAsync(
	request *CountAccessLogRequest,
	callback chan<- CountAccessLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "accessLog",
			"function":    "countAccessLog",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Service != nil {
		bodies["service"] = *request.Service
	}
	if request.Method != nil {
		bodies["method"] = *request.Method
	}
	if request.UserId != nil {
		bodies["userId"] = *request.UserId
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.LongTerm != nil {
		bodies["longTerm"] = *request.LongTerm
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

	go p.countAccessLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) CountAccessLog(
	request *CountAccessLogRequest,
) (*CountAccessLogResult, error) {
	callback := make(chan CountAccessLogAsyncResult, 1)
	go p.CountAccessLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) queryIssueStampSheetLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- QueryIssueStampSheetLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- QueryIssueStampSheetLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result QueryIssueStampSheetLogResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryIssueStampSheetLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- QueryIssueStampSheetLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) QueryIssueStampSheetLogAsync(
	request *QueryIssueStampSheetLogRequest,
	callback chan<- QueryIssueStampSheetLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "issueStampSheetLog",
			"function":    "queryIssueStampSheetLog",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Service != nil && *request.Service != "" {
		bodies["service"] = *request.Service
	}
	if request.Method != nil && *request.Method != "" {
		bodies["method"] = *request.Method
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Action != nil && *request.Action != "" {
		bodies["action"] = *request.Action
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.LongTerm != nil {
		bodies["longTerm"] = *request.LongTerm
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

	go p.queryIssueStampSheetLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) QueryIssueStampSheetLog(
	request *QueryIssueStampSheetLogRequest,
) (*QueryIssueStampSheetLogResult, error) {
	callback := make(chan QueryIssueStampSheetLogAsyncResult, 1)
	go p.QueryIssueStampSheetLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) countIssueStampSheetLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CountIssueStampSheetLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CountIssueStampSheetLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CountIssueStampSheetLogResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountIssueStampSheetLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CountIssueStampSheetLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) CountIssueStampSheetLogAsync(
	request *CountIssueStampSheetLogRequest,
	callback chan<- CountIssueStampSheetLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "issueStampSheetLog",
			"function":    "countIssueStampSheetLog",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Service != nil {
		bodies["service"] = *request.Service
	}
	if request.Method != nil {
		bodies["method"] = *request.Method
	}
	if request.UserId != nil {
		bodies["userId"] = *request.UserId
	}
	if request.Action != nil {
		bodies["action"] = *request.Action
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.LongTerm != nil {
		bodies["longTerm"] = *request.LongTerm
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

	go p.countIssueStampSheetLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) CountIssueStampSheetLog(
	request *CountIssueStampSheetLogRequest,
) (*CountIssueStampSheetLogResult, error) {
	callback := make(chan CountIssueStampSheetLogAsyncResult, 1)
	go p.CountIssueStampSheetLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) queryExecuteStampSheetLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- QueryExecuteStampSheetLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- QueryExecuteStampSheetLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result QueryExecuteStampSheetLogResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryExecuteStampSheetLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- QueryExecuteStampSheetLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) QueryExecuteStampSheetLogAsync(
	request *QueryExecuteStampSheetLogRequest,
	callback chan<- QueryExecuteStampSheetLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "executeStampSheetLog",
			"function":    "queryExecuteStampSheetLog",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Service != nil && *request.Service != "" {
		bodies["service"] = *request.Service
	}
	if request.Method != nil && *request.Method != "" {
		bodies["method"] = *request.Method
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Action != nil && *request.Action != "" {
		bodies["action"] = *request.Action
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.LongTerm != nil {
		bodies["longTerm"] = *request.LongTerm
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

	go p.queryExecuteStampSheetLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) QueryExecuteStampSheetLog(
	request *QueryExecuteStampSheetLogRequest,
) (*QueryExecuteStampSheetLogResult, error) {
	callback := make(chan QueryExecuteStampSheetLogAsyncResult, 1)
	go p.QueryExecuteStampSheetLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) countExecuteStampSheetLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CountExecuteStampSheetLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CountExecuteStampSheetLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CountExecuteStampSheetLogResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountExecuteStampSheetLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CountExecuteStampSheetLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) CountExecuteStampSheetLogAsync(
	request *CountExecuteStampSheetLogRequest,
	callback chan<- CountExecuteStampSheetLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "executeStampSheetLog",
			"function":    "countExecuteStampSheetLog",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Service != nil {
		bodies["service"] = *request.Service
	}
	if request.Method != nil {
		bodies["method"] = *request.Method
	}
	if request.UserId != nil {
		bodies["userId"] = *request.UserId
	}
	if request.Action != nil {
		bodies["action"] = *request.Action
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.LongTerm != nil {
		bodies["longTerm"] = *request.LongTerm
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

	go p.countExecuteStampSheetLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) CountExecuteStampSheetLog(
	request *CountExecuteStampSheetLogRequest,
) (*CountExecuteStampSheetLogResult, error) {
	callback := make(chan CountExecuteStampSheetLogAsyncResult, 1)
	go p.CountExecuteStampSheetLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) queryExecuteStampTaskLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- QueryExecuteStampTaskLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- QueryExecuteStampTaskLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result QueryExecuteStampTaskLogResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryExecuteStampTaskLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- QueryExecuteStampTaskLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) QueryExecuteStampTaskLogAsync(
	request *QueryExecuteStampTaskLogRequest,
	callback chan<- QueryExecuteStampTaskLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "executeStampTaskLog",
			"function":    "queryExecuteStampTaskLog",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Service != nil && *request.Service != "" {
		bodies["service"] = *request.Service
	}
	if request.Method != nil && *request.Method != "" {
		bodies["method"] = *request.Method
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.Action != nil && *request.Action != "" {
		bodies["action"] = *request.Action
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.LongTerm != nil {
		bodies["longTerm"] = *request.LongTerm
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

	go p.queryExecuteStampTaskLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) QueryExecuteStampTaskLog(
	request *QueryExecuteStampTaskLogRequest,
) (*QueryExecuteStampTaskLogResult, error) {
	callback := make(chan QueryExecuteStampTaskLogAsyncResult, 1)
	go p.QueryExecuteStampTaskLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) countExecuteStampTaskLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CountExecuteStampTaskLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CountExecuteStampTaskLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CountExecuteStampTaskLogResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountExecuteStampTaskLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CountExecuteStampTaskLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) CountExecuteStampTaskLogAsync(
	request *CountExecuteStampTaskLogRequest,
	callback chan<- CountExecuteStampTaskLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "executeStampTaskLog",
			"function":    "countExecuteStampTaskLog",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Service != nil {
		bodies["service"] = *request.Service
	}
	if request.Method != nil {
		bodies["method"] = *request.Method
	}
	if request.UserId != nil {
		bodies["userId"] = *request.UserId
	}
	if request.Action != nil {
		bodies["action"] = *request.Action
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.LongTerm != nil {
		bodies["longTerm"] = *request.LongTerm
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

	go p.countExecuteStampTaskLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) CountExecuteStampTaskLog(
	request *CountExecuteStampTaskLogRequest,
) (*CountExecuteStampTaskLogResult, error) {
	callback := make(chan CountExecuteStampTaskLogAsyncResult, 1)
	go p.CountExecuteStampTaskLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) queryInGameLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- QueryInGameLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- QueryInGameLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result QueryInGameLogResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryInGameLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- QueryInGameLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) QueryInGameLogAsync(
	request *QueryInGameLogRequest,
	callback chan<- QueryInGameLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "inGameLog",
			"function":    "queryInGameLog",
			"contentType": "application/json",
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
	if request.Tags != nil {
		var _tags []interface{}
		for _, item := range request.Tags {
			_tags = append(_tags, item)
		}
		bodies["tags"] = _tags
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.LongTerm != nil {
		bodies["longTerm"] = *request.LongTerm
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

	go p.queryInGameLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) QueryInGameLog(
	request *QueryInGameLogRequest,
) (*QueryInGameLogResult, error) {
	callback := make(chan QueryInGameLogAsyncResult, 1)
	go p.QueryInGameLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) sendInGameLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SendInGameLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SendInGameLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SendInGameLogResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SendInGameLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- SendInGameLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) SendInGameLogAsync(
	request *SendInGameLogRequest,
	callback chan<- SendInGameLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "inGameLog",
			"function":    "sendInGameLog",
			"contentType": "application/json",
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
	if request.Tags != nil {
		var _tags []interface{}
		for _, item := range request.Tags {
			_tags = append(_tags, item)
		}
		bodies["tags"] = _tags
	}
	if request.Payload != nil && *request.Payload != "" {
		bodies["payload"] = *request.Payload
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

	go p.sendInGameLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) SendInGameLog(
	request *SendInGameLogRequest,
) (*SendInGameLogResult, error) {
	callback := make(chan SendInGameLogAsyncResult, 1)
	go p.SendInGameLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) sendInGameLogByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SendInGameLogByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SendInGameLogByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SendInGameLogByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SendInGameLogByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- SendInGameLogByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) SendInGameLogByUserIdAsync(
	request *SendInGameLogByUserIdRequest,
	callback chan<- SendInGameLogByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "inGameLog",
			"function":    "sendInGameLogByUserId",
			"contentType": "application/json",
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
	if request.Tags != nil {
		var _tags []interface{}
		for _, item := range request.Tags {
			_tags = append(_tags, item)
		}
		bodies["tags"] = _tags
	}
	if request.Payload != nil && *request.Payload != "" {
		bodies["payload"] = *request.Payload
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

	go p.sendInGameLogByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) SendInGameLogByUserId(
	request *SendInGameLogByUserIdRequest,
) (*SendInGameLogByUserIdResult, error) {
	callback := make(chan SendInGameLogByUserIdAsyncResult, 1)
	go p.SendInGameLogByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) queryAccessLogWithTelemetryAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- QueryAccessLogWithTelemetryAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- QueryAccessLogWithTelemetryAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result QueryAccessLogWithTelemetryResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryAccessLogWithTelemetryAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- QueryAccessLogWithTelemetryAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) QueryAccessLogWithTelemetryAsync(
	request *QueryAccessLogWithTelemetryRequest,
	callback chan<- QueryAccessLogWithTelemetryAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "accessLogWithTelemetry",
			"function":    "queryAccessLogWithTelemetry",
			"contentType": "application/json",
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
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.LongTerm != nil {
		bodies["longTerm"] = *request.LongTerm
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

	go p.queryAccessLogWithTelemetryAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) QueryAccessLogWithTelemetry(
	request *QueryAccessLogWithTelemetryRequest,
) (*QueryAccessLogWithTelemetryResult, error) {
	callback := make(chan QueryAccessLogWithTelemetryAsyncResult, 1)
	go p.QueryAccessLogWithTelemetryAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) describeInsightsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeInsightsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInsightsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInsightsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInsightsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeInsightsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) DescribeInsightsAsync(
	request *DescribeInsightsRequest,
	callback chan<- DescribeInsightsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "insight",
			"function":    "describeInsights",
			"contentType": "application/json",
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

	go p.describeInsightsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) DescribeInsights(
	request *DescribeInsightsRequest,
) (*DescribeInsightsResult, error) {
	callback := make(chan DescribeInsightsAsyncResult, 1)
	go p.DescribeInsightsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) createInsightAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateInsightAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateInsightAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateInsightResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateInsightAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateInsightAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) CreateInsightAsync(
	request *CreateInsightRequest,
	callback chan<- CreateInsightAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "insight",
			"function":    "createInsight",
			"contentType": "application/json",
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

	go p.createInsightAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) CreateInsight(
	request *CreateInsightRequest,
) (*CreateInsightResult, error) {
	callback := make(chan CreateInsightAsyncResult, 1)
	go p.CreateInsightAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) getInsightAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetInsightAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetInsightAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetInsightResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetInsightAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetInsightAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) GetInsightAsync(
	request *GetInsightRequest,
	callback chan<- GetInsightAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "insight",
			"function":    "getInsight",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.InsightName != nil && *request.InsightName != "" {
		bodies["insightName"] = *request.InsightName
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

	go p.getInsightAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) GetInsight(
	request *GetInsightRequest,
) (*GetInsightResult, error) {
	callback := make(chan GetInsightAsyncResult, 1)
	go p.GetInsightAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) deleteInsightAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteInsightAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteInsightAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteInsightResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteInsightAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteInsightAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) DeleteInsightAsync(
	request *DeleteInsightRequest,
	callback chan<- DeleteInsightAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "insight",
			"function":    "deleteInsight",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.InsightName != nil && *request.InsightName != "" {
		bodies["insightName"] = *request.InsightName
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

	go p.deleteInsightAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) DeleteInsight(
	request *DeleteInsightRequest,
) (*DeleteInsightResult, error) {
	callback := make(chan DeleteInsightAsyncResult, 1)
	go p.DeleteInsightAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) describeFacetModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeFacetModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFacetModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFacetModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFacetModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeFacetModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) DescribeFacetModelsAsync(
	request *DescribeFacetModelsRequest,
	callback chan<- DescribeFacetModelsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "facetModel",
			"function":    "describeFacetModels",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.NamePrefix != nil && *request.NamePrefix != "" {
		bodies["namePrefix"] = *request.NamePrefix
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

	go p.describeFacetModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) DescribeFacetModels(
	request *DescribeFacetModelsRequest,
) (*DescribeFacetModelsResult, error) {
	callback := make(chan DescribeFacetModelsAsyncResult, 1)
	go p.DescribeFacetModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) createFacetModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateFacetModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateFacetModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateFacetModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateFacetModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateFacetModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) CreateFacetModelAsync(
	request *CreateFacetModelRequest,
	callback chan<- CreateFacetModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "facetModel",
			"function":    "createFacetModel",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Field != nil && *request.Field != "" {
		bodies["field"] = *request.Field
	}
	if request.Type != nil && *request.Type != "" {
		bodies["type"] = *request.Type
	}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Order != nil {
		bodies["order"] = *request.Order
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

	go p.createFacetModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) CreateFacetModel(
	request *CreateFacetModelRequest,
) (*CreateFacetModelResult, error) {
	callback := make(chan CreateFacetModelAsyncResult, 1)
	go p.CreateFacetModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) getFacetModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetFacetModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFacetModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFacetModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFacetModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetFacetModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) GetFacetModelAsync(
	request *GetFacetModelRequest,
	callback chan<- GetFacetModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "facetModel",
			"function":    "getFacetModel",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Field != nil && *request.Field != "" {
		bodies["field"] = *request.Field
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

	go p.getFacetModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) GetFacetModel(
	request *GetFacetModelRequest,
) (*GetFacetModelResult, error) {
	callback := make(chan GetFacetModelAsyncResult, 1)
	go p.GetFacetModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) updateFacetModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateFacetModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateFacetModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateFacetModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateFacetModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateFacetModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) UpdateFacetModelAsync(
	request *UpdateFacetModelRequest,
	callback chan<- UpdateFacetModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "facetModel",
			"function":    "updateFacetModel",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Field != nil && *request.Field != "" {
		bodies["field"] = *request.Field
	}
	if request.Type != nil && *request.Type != "" {
		bodies["type"] = *request.Type
	}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Order != nil {
		bodies["order"] = *request.Order
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

	go p.updateFacetModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) UpdateFacetModel(
	request *UpdateFacetModelRequest,
) (*UpdateFacetModelResult, error) {
	callback := make(chan UpdateFacetModelAsyncResult, 1)
	go p.UpdateFacetModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) deleteFacetModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteFacetModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteFacetModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteFacetModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteFacetModelAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteFacetModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) DeleteFacetModelAsync(
	request *DeleteFacetModelRequest,
	callback chan<- DeleteFacetModelAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "facetModel",
			"function":    "deleteFacetModel",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Field != nil && *request.Field != "" {
		bodies["field"] = *request.Field
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

	go p.deleteFacetModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) DeleteFacetModel(
	request *DeleteFacetModelRequest,
) (*DeleteFacetModelResult, error) {
	callback := make(chan DeleteFacetModelAsyncResult, 1)
	go p.DeleteFacetModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) describeDashboardsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeDashboardsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDashboardsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDashboardsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDashboardsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeDashboardsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) DescribeDashboardsAsync(
	request *DescribeDashboardsRequest,
	callback chan<- DescribeDashboardsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "dashboard",
			"function":    "describeDashboards",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.NamePrefix != nil && *request.NamePrefix != "" {
		bodies["namePrefix"] = *request.NamePrefix
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

	go p.describeDashboardsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) DescribeDashboards(
	request *DescribeDashboardsRequest,
) (*DescribeDashboardsResult, error) {
	callback := make(chan DescribeDashboardsAsyncResult, 1)
	go p.DescribeDashboardsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) createDashboardAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateDashboardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateDashboardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateDashboardResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateDashboardAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CreateDashboardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) CreateDashboardAsync(
	request *CreateDashboardRequest,
	callback chan<- CreateDashboardAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "dashboard",
			"function":    "createDashboard",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
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

	go p.createDashboardAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) CreateDashboard(
	request *CreateDashboardRequest,
) (*CreateDashboardResult, error) {
	callback := make(chan CreateDashboardAsyncResult, 1)
	go p.CreateDashboardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) getDashboardAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetDashboardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDashboardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDashboardResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDashboardAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetDashboardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) GetDashboardAsync(
	request *GetDashboardRequest,
	callback chan<- GetDashboardAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "dashboard",
			"function":    "getDashboard",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.DashboardName != nil && *request.DashboardName != "" {
		bodies["dashboardName"] = *request.DashboardName
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

	go p.getDashboardAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) GetDashboard(
	request *GetDashboardRequest,
) (*GetDashboardResult, error) {
	callback := make(chan GetDashboardAsyncResult, 1)
	go p.GetDashboardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) updateDashboardAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateDashboardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateDashboardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateDashboardResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateDashboardAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- UpdateDashboardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) UpdateDashboardAsync(
	request *UpdateDashboardRequest,
	callback chan<- UpdateDashboardAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "dashboard",
			"function":    "updateDashboard",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.DashboardName != nil && *request.DashboardName != "" {
		bodies["dashboardName"] = *request.DashboardName
	}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Payload != nil && *request.Payload != "" {
		bodies["payload"] = *request.Payload
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

	go p.updateDashboardAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) UpdateDashboard(
	request *UpdateDashboardRequest,
) (*UpdateDashboardResult, error) {
	callback := make(chan UpdateDashboardAsyncResult, 1)
	go p.UpdateDashboardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) duplicateDashboardAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DuplicateDashboardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DuplicateDashboardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DuplicateDashboardResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DuplicateDashboardAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DuplicateDashboardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) DuplicateDashboardAsync(
	request *DuplicateDashboardRequest,
	callback chan<- DuplicateDashboardAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "dashboard",
			"function":    "duplicateDashboard",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.DashboardName != nil && *request.DashboardName != "" {
		bodies["dashboardName"] = *request.DashboardName
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

	go p.duplicateDashboardAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) DuplicateDashboard(
	request *DuplicateDashboardRequest,
) (*DuplicateDashboardResult, error) {
	callback := make(chan DuplicateDashboardAsyncResult, 1)
	go p.DuplicateDashboardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) deleteDashboardAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteDashboardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteDashboardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteDashboardResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteDashboardAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DeleteDashboardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) DeleteDashboardAsync(
	request *DeleteDashboardRequest,
	callback chan<- DeleteDashboardAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "dashboard",
			"function":    "deleteDashboard",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.DashboardName != nil && *request.DashboardName != "" {
		bodies["dashboardName"] = *request.DashboardName
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

	go p.deleteDashboardAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) DeleteDashboard(
	request *DeleteDashboardRequest,
) (*DeleteDashboardResult, error) {
	callback := make(chan DeleteDashboardAsyncResult, 1)
	go p.DeleteDashboardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) queryLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- QueryLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- QueryLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result QueryLogResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- QueryLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) QueryLogAsync(
	request *QueryLogRequest,
	callback chan<- QueryLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "logEntry",
			"function":    "queryLog",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.Query != nil && *request.Query != "" {
		bodies["query"] = *request.Query
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

	go p.queryLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) QueryLog(
	request *QueryLogRequest,
) (*QueryLogResult, error) {
	callback := make(chan QueryLogAsyncResult, 1)
	go p.QueryLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) getLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLogResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) GetLogAsync(
	request *GetLogRequest,
	callback chan<- GetLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "logEntry",
			"function":    "getLog",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.LogRequestId != nil && *request.LogRequestId != "" {
		bodies["logRequestId"] = *request.LogRequestId
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
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

	go p.getLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) GetLog(
	request *GetLogRequest,
) (*GetLogResult, error) {
	callback := make(chan GetLogAsyncResult, 1)
	go p.GetLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) queryFacetsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- QueryFacetsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- QueryFacetsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result QueryFacetsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryFacetsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- QueryFacetsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) QueryFacetsAsync(
	request *QueryFacetsRequest,
	callback chan<- QueryFacetsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "logEntry",
			"function":    "queryFacets",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.Query != nil && *request.Query != "" {
		bodies["query"] = *request.Query
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

	go p.queryFacetsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) QueryFacets(
	request *QueryFacetsRequest,
) (*QueryFacetsResult, error) {
	callback := make(chan QueryFacetsAsyncResult, 1)
	go p.QueryFacetsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) queryTimeseriesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- QueryTimeseriesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- QueryTimeseriesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result QueryTimeseriesResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryTimeseriesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- QueryTimeseriesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) QueryTimeseriesAsync(
	request *QueryTimeseriesRequest,
	callback chan<- QueryTimeseriesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "logEntry",
			"function":    "queryTimeseries",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.Query != nil && *request.Query != "" {
		bodies["query"] = *request.Query
	}
	if request.GroupBy != nil {
		var _groupBy []interface{}
		for _, item := range request.GroupBy {
			_groupBy = append(_groupBy, item)
		}
		bodies["groupBy"] = _groupBy
	}
	if request.Aggregation != nil {
		bodies["aggregation"] = request.Aggregation.ToDict()
	}
	if request.Interval != nil {
		bodies["interval"] = *request.Interval
	}
	if request.SeriesLimit != nil {
		bodies["seriesLimit"] = *request.SeriesLimit
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

	go p.queryTimeseriesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) QueryTimeseries(
	request *QueryTimeseriesRequest,
) (*QueryTimeseriesResult, error) {
	callback := make(chan QueryTimeseriesAsyncResult, 1)
	go p.QueryTimeseriesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) getTraceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetTraceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetTraceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetTraceResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetTraceAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetTraceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) GetTraceAsync(
	request *GetTraceRequest,
	callback chan<- GetTraceAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "logEntry",
			"function":    "getTrace",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.TraceId != nil && *request.TraceId != "" {
		bodies["traceId"] = *request.TraceId
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
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

	go p.getTraceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) GetTrace(
	request *GetTraceRequest,
) (*GetTraceResult, error) {
	callback := make(chan GetTraceAsyncResult, 1)
	go p.GetTraceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) queryMetricsTimeseriesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- QueryMetricsTimeseriesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- QueryMetricsTimeseriesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result QueryMetricsTimeseriesResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryMetricsTimeseriesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- QueryMetricsTimeseriesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) QueryMetricsTimeseriesAsync(
	request *QueryMetricsTimeseriesRequest,
	callback chan<- QueryMetricsTimeseriesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "metricModel",
			"function":    "queryMetricsTimeseries",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.Query != nil && *request.Query != "" {
		bodies["query"] = *request.Query
	}
	if request.GroupBy != nil {
		var _groupBy []interface{}
		for _, item := range request.GroupBy {
			_groupBy = append(_groupBy, item)
		}
		bodies["groupBy"] = _groupBy
	}
	if request.Aggregations != nil {
		var _aggregations []interface{}
		for _, item := range request.Aggregations {
			_aggregations = append(_aggregations, item)
		}
		bodies["aggregations"] = _aggregations
	}
	if request.Interval != nil {
		bodies["interval"] = *request.Interval
	}
	if request.SeriesLimit != nil {
		bodies["seriesLimit"] = *request.SeriesLimit
	}
	if request.OrderKey != nil && *request.OrderKey != "" {
		bodies["orderKey"] = *request.OrderKey
	}
	if request.OrderBy != nil && *request.OrderBy != "" {
		bodies["orderBy"] = *request.OrderBy
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

	go p.queryMetricsTimeseriesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) QueryMetricsTimeseries(
	request *QueryMetricsTimeseriesRequest,
) (*QueryMetricsTimeseriesResult, error) {
	callback := make(chan QueryMetricsTimeseriesAsyncResult, 1)
	go p.QueryMetricsTimeseriesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) describeMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) DescribeMetricsAsync(
	request *DescribeMetricsRequest,
	callback chan<- DescribeMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "metricModel",
			"function":    "describeMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.NamePrefix != nil && *request.NamePrefix != "" {
		bodies["namePrefix"] = *request.NamePrefix
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

	go p.describeMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) DescribeMetrics(
	request *DescribeMetricsRequest,
) (*DescribeMetricsResult, error) {
	callback := make(chan DescribeMetricsAsyncResult, 1)
	go p.DescribeMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2LogWebSocketClient) describeLabelValuesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeLabelValuesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLabelValuesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLabelValuesResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLabelValuesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeLabelValuesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogWebSocketClient) DescribeLabelValuesAsync(
	request *DescribeLabelValuesRequest,
	callback chan<- DescribeLabelValuesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "log",
			"component":   "metricModel",
			"function":    "describeLabelValues",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.MetricName != nil && *request.MetricName != "" {
		bodies["metricName"] = *request.MetricName
	}
	if request.LabelNamePrefix != nil && *request.LabelNamePrefix != "" {
		bodies["labelNamePrefix"] = *request.LabelNamePrefix
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

	go p.describeLabelValuesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2LogWebSocketClient) DescribeLabelValues(
	request *DescribeLabelValuesRequest,
) (*DescribeLabelValuesResult, error) {
	callback := make(chan DescribeLabelValuesAsyncResult, 1)
	go p.DescribeLabelValuesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
