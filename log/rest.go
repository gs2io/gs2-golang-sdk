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
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

var EndpointHost *string = nil

type Gs2LogRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2LogRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeNamespacesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2LogRestClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	path := "/"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeNamespacesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) DescribeNamespaces(
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

func createNamespaceAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- CreateNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CreateNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2LogRestClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	path := "/"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go createNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) CreateNamespace(
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

func getNamespaceStatusAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetNamespaceStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2LogRestClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	path := "/{namespaceName}/status"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getNamespaceStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) GetNamespaceStatus(
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

func getNamespaceAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- GetNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2LogRestClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) GetNamespace(
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

func updateNamespaceAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- UpdateNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2LogRestClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go updateNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) UpdateNamespace(
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

func deleteNamespaceAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DeleteNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2LogRestClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go deleteNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) DeleteNamespace(
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

func getServiceVersionAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- GetServiceVersionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetServiceVersionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetServiceVersionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetServiceVersionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) GetServiceVersionAsync(
	request *GetServiceVersionRequest,
	callback chan<- GetServiceVersionAsyncResult,
) {
	path := "/system/version"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getServiceVersionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) GetServiceVersion(
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

func queryAccessLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- QueryAccessLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- QueryAccessLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryAccessLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- QueryAccessLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) QueryAccessLogAsync(
	request *QueryAccessLogRequest,
	callback chan<- QueryAccessLogAsyncResult,
) {
	path := "/{namespaceName}/log/access"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Service != nil {
		queryStrings["service"] = core.ToString(*request.Service)
	}
	if request.Method != nil {
		queryStrings["method"] = core.ToString(*request.Method)
	}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}
	if request.Begin != nil {
		queryStrings["begin"] = core.ToString(*request.Begin)
	}
	if request.End != nil {
		queryStrings["end"] = core.ToString(*request.End)
	}
	if request.LongTerm != nil {
		queryStrings["longTerm"] = core.ToString(*request.LongTerm)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go queryAccessLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) QueryAccessLog(
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

func countAccessLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- CountAccessLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CountAccessLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountAccessLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CountAccessLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) CountAccessLogAsync(
	request *CountAccessLogRequest,
	callback chan<- CountAccessLogAsyncResult,
) {
	path := "/{namespaceName}/log/access/count"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Service != nil {
		queryStrings["service"] = core.ToString(*request.Service)
	}
	if request.Method != nil {
		queryStrings["method"] = core.ToString(*request.Method)
	}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}
	if request.Begin != nil {
		queryStrings["begin"] = core.ToString(*request.Begin)
	}
	if request.End != nil {
		queryStrings["end"] = core.ToString(*request.End)
	}
	if request.LongTerm != nil {
		queryStrings["longTerm"] = core.ToString(*request.LongTerm)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go countAccessLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) CountAccessLog(
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

func queryIssueStampSheetLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- QueryIssueStampSheetLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- QueryIssueStampSheetLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryIssueStampSheetLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- QueryIssueStampSheetLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) QueryIssueStampSheetLogAsync(
	request *QueryIssueStampSheetLogRequest,
	callback chan<- QueryIssueStampSheetLogAsyncResult,
) {
	path := "/{namespaceName}/log/issue/stamp/sheet"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Service != nil {
		queryStrings["service"] = core.ToString(*request.Service)
	}
	if request.Method != nil {
		queryStrings["method"] = core.ToString(*request.Method)
	}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}
	if request.Action != nil {
		queryStrings["action"] = core.ToString(*request.Action)
	}
	if request.Begin != nil {
		queryStrings["begin"] = core.ToString(*request.Begin)
	}
	if request.End != nil {
		queryStrings["end"] = core.ToString(*request.End)
	}
	if request.LongTerm != nil {
		queryStrings["longTerm"] = core.ToString(*request.LongTerm)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go queryIssueStampSheetLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) QueryIssueStampSheetLog(
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

func countIssueStampSheetLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- CountIssueStampSheetLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CountIssueStampSheetLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountIssueStampSheetLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CountIssueStampSheetLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) CountIssueStampSheetLogAsync(
	request *CountIssueStampSheetLogRequest,
	callback chan<- CountIssueStampSheetLogAsyncResult,
) {
	path := "/{namespaceName}/log/issue/stamp/sheet/count"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Service != nil {
		queryStrings["service"] = core.ToString(*request.Service)
	}
	if request.Method != nil {
		queryStrings["method"] = core.ToString(*request.Method)
	}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}
	if request.Action != nil {
		queryStrings["action"] = core.ToString(*request.Action)
	}
	if request.Begin != nil {
		queryStrings["begin"] = core.ToString(*request.Begin)
	}
	if request.End != nil {
		queryStrings["end"] = core.ToString(*request.End)
	}
	if request.LongTerm != nil {
		queryStrings["longTerm"] = core.ToString(*request.LongTerm)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go countIssueStampSheetLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) CountIssueStampSheetLog(
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

func queryExecuteStampSheetLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- QueryExecuteStampSheetLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- QueryExecuteStampSheetLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryExecuteStampSheetLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- QueryExecuteStampSheetLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) QueryExecuteStampSheetLogAsync(
	request *QueryExecuteStampSheetLogRequest,
	callback chan<- QueryExecuteStampSheetLogAsyncResult,
) {
	path := "/{namespaceName}/log/execute/stamp/sheet"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Service != nil {
		queryStrings["service"] = core.ToString(*request.Service)
	}
	if request.Method != nil {
		queryStrings["method"] = core.ToString(*request.Method)
	}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}
	if request.Action != nil {
		queryStrings["action"] = core.ToString(*request.Action)
	}
	if request.Begin != nil {
		queryStrings["begin"] = core.ToString(*request.Begin)
	}
	if request.End != nil {
		queryStrings["end"] = core.ToString(*request.End)
	}
	if request.LongTerm != nil {
		queryStrings["longTerm"] = core.ToString(*request.LongTerm)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go queryExecuteStampSheetLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) QueryExecuteStampSheetLog(
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

func countExecuteStampSheetLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- CountExecuteStampSheetLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CountExecuteStampSheetLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountExecuteStampSheetLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CountExecuteStampSheetLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) CountExecuteStampSheetLogAsync(
	request *CountExecuteStampSheetLogRequest,
	callback chan<- CountExecuteStampSheetLogAsyncResult,
) {
	path := "/{namespaceName}/log/execute/stamp/sheet/count"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Service != nil {
		queryStrings["service"] = core.ToString(*request.Service)
	}
	if request.Method != nil {
		queryStrings["method"] = core.ToString(*request.Method)
	}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}
	if request.Action != nil {
		queryStrings["action"] = core.ToString(*request.Action)
	}
	if request.Begin != nil {
		queryStrings["begin"] = core.ToString(*request.Begin)
	}
	if request.End != nil {
		queryStrings["end"] = core.ToString(*request.End)
	}
	if request.LongTerm != nil {
		queryStrings["longTerm"] = core.ToString(*request.LongTerm)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go countExecuteStampSheetLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) CountExecuteStampSheetLog(
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

func queryExecuteStampTaskLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- QueryExecuteStampTaskLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- QueryExecuteStampTaskLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryExecuteStampTaskLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- QueryExecuteStampTaskLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) QueryExecuteStampTaskLogAsync(
	request *QueryExecuteStampTaskLogRequest,
	callback chan<- QueryExecuteStampTaskLogAsyncResult,
) {
	path := "/{namespaceName}/log/execute/stamp/task"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Service != nil {
		queryStrings["service"] = core.ToString(*request.Service)
	}
	if request.Method != nil {
		queryStrings["method"] = core.ToString(*request.Method)
	}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}
	if request.Action != nil {
		queryStrings["action"] = core.ToString(*request.Action)
	}
	if request.Begin != nil {
		queryStrings["begin"] = core.ToString(*request.Begin)
	}
	if request.End != nil {
		queryStrings["end"] = core.ToString(*request.End)
	}
	if request.LongTerm != nil {
		queryStrings["longTerm"] = core.ToString(*request.LongTerm)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go queryExecuteStampTaskLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) QueryExecuteStampTaskLog(
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

func countExecuteStampTaskLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- CountExecuteStampTaskLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CountExecuteStampTaskLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountExecuteStampTaskLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CountExecuteStampTaskLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) CountExecuteStampTaskLogAsync(
	request *CountExecuteStampTaskLogRequest,
	callback chan<- CountExecuteStampTaskLogAsyncResult,
) {
	path := "/{namespaceName}/log/execute/stamp/task/count"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Service != nil {
		queryStrings["service"] = core.ToString(*request.Service)
	}
	if request.Method != nil {
		queryStrings["method"] = core.ToString(*request.Method)
	}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}
	if request.Action != nil {
		queryStrings["action"] = core.ToString(*request.Action)
	}
	if request.Begin != nil {
		queryStrings["begin"] = core.ToString(*request.Begin)
	}
	if request.End != nil {
		queryStrings["end"] = core.ToString(*request.End)
	}
	if request.LongTerm != nil {
		queryStrings["longTerm"] = core.ToString(*request.LongTerm)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go countExecuteStampTaskLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) CountExecuteStampTaskLog(
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

func queryInGameLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- QueryInGameLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- QueryInGameLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryInGameLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- QueryInGameLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) QueryInGameLogAsync(
	request *QueryInGameLogRequest,
	callback chan<- QueryInGameLogAsyncResult,
) {
	path := "/{namespaceName}/ingame/log"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go queryInGameLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) QueryInGameLog(
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

func sendInGameLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- SendInGameLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- SendInGameLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SendInGameLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SendInGameLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) SendInGameLogAsync(
	request *SendInGameLogRequest,
	callback chan<- SendInGameLogAsyncResult,
) {
	path := "/{namespaceName}/ingame/log/user/me/send"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go sendInGameLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) SendInGameLog(
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

func sendInGameLogByUserIdAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- SendInGameLogByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- SendInGameLogByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SendInGameLogByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SendInGameLogByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) SendInGameLogByUserIdAsync(
	request *SendInGameLogByUserIdRequest,
	callback chan<- SendInGameLogByUserIdAsyncResult,
) {
	path := "/{namespaceName}/ingame/log/user/{userId}/send"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go sendInGameLogByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) SendInGameLogByUserId(
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

func queryAccessLogWithTelemetryAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- QueryAccessLogWithTelemetryAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- QueryAccessLogWithTelemetryAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryAccessLogWithTelemetryAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- QueryAccessLogWithTelemetryAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) QueryAccessLogWithTelemetryAsync(
	request *QueryAccessLogWithTelemetryRequest,
	callback chan<- QueryAccessLogWithTelemetryAsyncResult,
) {
	path := "/{namespaceName}/log/access/telemetry"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}
	if request.Begin != nil {
		queryStrings["begin"] = core.ToString(*request.Begin)
	}
	if request.End != nil {
		queryStrings["end"] = core.ToString(*request.End)
	}
	if request.LongTerm != nil {
		queryStrings["longTerm"] = core.ToString(*request.LongTerm)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go queryAccessLogWithTelemetryAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) QueryAccessLogWithTelemetry(
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

func describeInsightsAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInsightsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeInsightsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInsightsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeInsightsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) DescribeInsightsAsync(
	request *DescribeInsightsRequest,
	callback chan<- DescribeInsightsAsyncResult,
) {
	path := "/{namespaceName}/insight"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeInsightsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) DescribeInsights(
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

func createInsightAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- CreateInsightAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CreateInsightAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateInsightAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateInsightAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) CreateInsightAsync(
	request *CreateInsightRequest,
	callback chan<- CreateInsightAsyncResult,
) {
	path := "/{namespaceName}/insight"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go createInsightAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) CreateInsight(
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

func getInsightAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- GetInsightAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetInsightAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetInsightAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetInsightAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) GetInsightAsync(
	request *GetInsightRequest,
	callback chan<- GetInsightAsyncResult,
) {
	path := "/{namespaceName}/insight/{insightName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InsightName != nil && *request.InsightName != "" {
		path = strings.ReplaceAll(path, "{insightName}", core.ToString(*request.InsightName))
	} else {
		path = strings.ReplaceAll(path, "{insightName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getInsightAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) GetInsight(
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

func deleteInsightAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteInsightAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DeleteInsightAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteInsightAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteInsightAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) DeleteInsightAsync(
	request *DeleteInsightRequest,
	callback chan<- DeleteInsightAsyncResult,
) {
	path := "/{namespaceName}/insight/{insightName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InsightName != nil && *request.InsightName != "" {
		path = strings.ReplaceAll(path, "{insightName}", core.ToString(*request.InsightName))
	} else {
		path = strings.ReplaceAll(path, "{insightName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go deleteInsightAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) DeleteInsight(
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

func describeFacetModelsAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeFacetModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeFacetModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFacetModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeFacetModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) DescribeFacetModelsAsync(
	request *DescribeFacetModelsRequest,
	callback chan<- DescribeFacetModelsAsyncResult,
) {
	path := "/{namespaceName}/model/facet"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.NamePrefix != nil {
		queryStrings["namePrefix"] = core.ToString(*request.NamePrefix)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeFacetModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) DescribeFacetModels(
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

func createFacetModelAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- CreateFacetModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CreateFacetModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateFacetModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateFacetModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) CreateFacetModelAsync(
	request *CreateFacetModelRequest,
	callback chan<- CreateFacetModelAsyncResult,
) {
	path := "/{namespaceName}/model/facet"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go createFacetModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) CreateFacetModel(
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

func getFacetModelAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- GetFacetModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetFacetModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFacetModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetFacetModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) GetFacetModelAsync(
	request *GetFacetModelRequest,
	callback chan<- GetFacetModelAsyncResult,
) {
	path := "/{namespaceName}/model/facet/{field}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Field != nil && *request.Field != "" {
		path = strings.ReplaceAll(path, "{field}", core.ToString(*request.Field))
	} else {
		path = strings.ReplaceAll(path, "{field}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getFacetModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) GetFacetModel(
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

func updateFacetModelAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateFacetModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- UpdateFacetModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateFacetModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateFacetModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) UpdateFacetModelAsync(
	request *UpdateFacetModelRequest,
	callback chan<- UpdateFacetModelAsyncResult,
) {
	path := "/{namespaceName}/model/facet/{field}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Field != nil && *request.Field != "" {
		path = strings.ReplaceAll(path, "{field}", core.ToString(*request.Field))
	} else {
		path = strings.ReplaceAll(path, "{field}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go updateFacetModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) UpdateFacetModel(
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

func deleteFacetModelAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteFacetModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DeleteFacetModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteFacetModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteFacetModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) DeleteFacetModelAsync(
	request *DeleteFacetModelRequest,
	callback chan<- DeleteFacetModelAsyncResult,
) {
	path := "/{namespaceName}/model/facet/{field}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Field != nil && *request.Field != "" {
		path = strings.ReplaceAll(path, "{field}", core.ToString(*request.Field))
	} else {
		path = strings.ReplaceAll(path, "{field}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go deleteFacetModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) DeleteFacetModel(
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

func describeDashboardsAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDashboardsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeDashboardsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDashboardsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeDashboardsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) DescribeDashboardsAsync(
	request *DescribeDashboardsRequest,
	callback chan<- DescribeDashboardsAsyncResult,
) {
	path := "/{namespaceName}/dashboard"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.NamePrefix != nil {
		queryStrings["namePrefix"] = core.ToString(*request.NamePrefix)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeDashboardsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) DescribeDashboards(
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

func createDashboardAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- CreateDashboardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CreateDashboardAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateDashboardAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateDashboardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) CreateDashboardAsync(
	request *CreateDashboardRequest,
	callback chan<- CreateDashboardAsyncResult,
) {
	path := "/{namespaceName}/dashboard"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go createDashboardAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) CreateDashboard(
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

func getDashboardAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- GetDashboardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetDashboardAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDashboardAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetDashboardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) GetDashboardAsync(
	request *GetDashboardRequest,
	callback chan<- GetDashboardAsyncResult,
) {
	path := "/{namespaceName}/dashboard/{dashboardName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.DashboardName != nil && *request.DashboardName != "" {
		path = strings.ReplaceAll(path, "{dashboardName}", core.ToString(*request.DashboardName))
	} else {
		path = strings.ReplaceAll(path, "{dashboardName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getDashboardAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) GetDashboard(
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

func updateDashboardAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateDashboardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- UpdateDashboardAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateDashboardAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateDashboardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) UpdateDashboardAsync(
	request *UpdateDashboardRequest,
	callback chan<- UpdateDashboardAsyncResult,
) {
	path := "/{namespaceName}/dashboard/{dashboardName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.DashboardName != nil && *request.DashboardName != "" {
		path = strings.ReplaceAll(path, "{dashboardName}", core.ToString(*request.DashboardName))
	} else {
		path = strings.ReplaceAll(path, "{dashboardName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go updateDashboardAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) UpdateDashboard(
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

func duplicateDashboardAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- DuplicateDashboardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DuplicateDashboardAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DuplicateDashboardAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DuplicateDashboardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) DuplicateDashboardAsync(
	request *DuplicateDashboardRequest,
	callback chan<- DuplicateDashboardAsyncResult,
) {
	path := "/{namespaceName}/dashboard/{dashboardName}/copy"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.DashboardName != nil && *request.DashboardName != "" {
		path = strings.ReplaceAll(path, "{dashboardName}", core.ToString(*request.DashboardName))
	} else {
		path = strings.ReplaceAll(path, "{dashboardName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go duplicateDashboardAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) DuplicateDashboard(
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

func deleteDashboardAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteDashboardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DeleteDashboardAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteDashboardAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteDashboardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) DeleteDashboardAsync(
	request *DeleteDashboardRequest,
	callback chan<- DeleteDashboardAsyncResult,
) {
	path := "/{namespaceName}/dashboard/{dashboardName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.DashboardName != nil && *request.DashboardName != "" {
		path = strings.ReplaceAll(path, "{dashboardName}", core.ToString(*request.DashboardName))
	} else {
		path = strings.ReplaceAll(path, "{dashboardName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go deleteDashboardAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) DeleteDashboard(
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

func queryLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- QueryLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- QueryLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- QueryLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) QueryLogAsync(
	request *QueryLogRequest,
	callback chan<- QueryLogAsyncResult,
) {
	path := "/{namespaceName}/log/v2/query"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go queryLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) QueryLog(
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

func getLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- GetLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) GetLogAsync(
	request *GetLogRequest,
	callback chan<- GetLogAsyncResult,
) {
	path := "/{namespaceName}/log/v2/query/{logRequestId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LogRequestId != nil && *request.LogRequestId != "" {
		path = strings.ReplaceAll(path, "{logRequestId}", core.ToString(*request.LogRequestId))
	} else {
		path = strings.ReplaceAll(path, "{logRequestId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Begin != nil {
		queryStrings["begin"] = core.ToString(*request.Begin)
	}
	if request.End != nil {
		queryStrings["end"] = core.ToString(*request.End)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) GetLog(
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

func queryFacetsAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- QueryFacetsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- QueryFacetsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryFacetsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- QueryFacetsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) QueryFacetsAsync(
	request *QueryFacetsRequest,
	callback chan<- QueryFacetsAsyncResult,
) {
	path := "/{namespaceName}/log/v2/query/facet"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go queryFacetsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) QueryFacets(
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

func queryTimeseriesAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- QueryTimeseriesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- QueryTimeseriesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryTimeseriesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- QueryTimeseriesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) QueryTimeseriesAsync(
	request *QueryTimeseriesRequest,
	callback chan<- QueryTimeseriesAsyncResult,
) {
	path := "/{namespaceName}/log/v2/timeseries"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go queryTimeseriesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) QueryTimeseries(
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

func getTraceAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- GetTraceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetTraceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetTraceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetTraceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) GetTraceAsync(
	request *GetTraceRequest,
	callback chan<- GetTraceAsyncResult,
) {
	path := "/{namespaceName}/log/v2/trace/{traceId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.TraceId != nil && *request.TraceId != "" {
		path = strings.ReplaceAll(path, "{traceId}", core.ToString(*request.TraceId))
	} else {
		path = strings.ReplaceAll(path, "{traceId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Begin != nil {
		queryStrings["begin"] = core.ToString(*request.Begin)
	}
	if request.End != nil {
		queryStrings["end"] = core.ToString(*request.End)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getTraceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) GetTrace(
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

func queryMetricsTimeseriesAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- QueryMetricsTimeseriesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- QueryMetricsTimeseriesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- QueryMetricsTimeseriesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- QueryMetricsTimeseriesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) QueryMetricsTimeseriesAsync(
	request *QueryMetricsTimeseriesRequest,
	callback chan<- QueryMetricsTimeseriesAsyncResult,
) {
	path := "/{namespaceName}/metrics/timeseries"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go queryMetricsTimeseriesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) QueryMetricsTimeseries(
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

func describeMetricsAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) DescribeMetricsAsync(
	request *DescribeMetricsRequest,
	callback chan<- DescribeMetricsAsyncResult,
) {
	path := "/{namespaceName}/model/metrics"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.NamePrefix != nil {
		queryStrings["namePrefix"] = core.ToString(*request.NamePrefix)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) DescribeMetrics(
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

func describeLabelValuesAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLabelValuesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeLabelValuesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLabelValuesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLabelValuesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) DescribeLabelValuesAsync(
	request *DescribeLabelValuesRequest,
	callback chan<- DescribeLabelValuesAsyncResult,
) {
	path := "/{namespaceName}/model/metrics/{metricName}/label"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MetricName != nil && *request.MetricName != "" {
		path = strings.ReplaceAll(path, "{metricName}", core.ToString(*request.MetricName))
	} else {
		path = strings.ReplaceAll(path, "{metricName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.LabelNamePrefix != nil {
		queryStrings["labelNamePrefix"] = core.ToString(*request.LabelNamePrefix)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeLabelValuesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LogRestClient) DescribeLabelValues(
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
