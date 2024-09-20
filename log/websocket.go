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
