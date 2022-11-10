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
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go describeNamespacesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getNamespaceStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go queryAccessLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go countAccessLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go queryIssueStampSheetLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go countIssueStampSheetLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go queryExecuteStampSheetLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go countExecuteStampSheetLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go queryExecuteStampTaskLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go countExecuteStampTaskLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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

func putLogAsyncHandler(
	client Gs2LogRestClient,
	job *core.NetworkJob,
	callback chan<- PutLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutLogResult
	if asyncResult.Err != nil {
		callback <- PutLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PutLogAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PutLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LogRestClient) PutLogAsync(
	request *PutLogRequest,
	callback chan<- PutLogAsyncResult,
) {
	path := "/log/put"

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.LoggingNamespaceId != nil && *request.LoggingNamespaceId != "" {
        bodies["loggingNamespaceId"] = *request.LoggingNamespaceId
    }
    if request.LogCategory != nil && *request.LogCategory != "" {
        bodies["logCategory"] = *request.LogCategory
    }
    if request.Payload != nil && *request.Payload != "" {
        bodies["payload"] = *request.Payload
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go putLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2LogRestClient) PutLog(
	request *PutLogRequest,
) (*PutLogResult, error) {
	callback := make(chan PutLogAsyncResult, 1)
	go p.PutLogAsync(
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go describeInsightsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createInsightAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InsightName != nil && *request.InsightName != ""  {
        path = strings.ReplaceAll(path, "{insightName}", core.ToString(*request.InsightName))
    } else {
        path = strings.ReplaceAll(path, "{insightName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getInsightAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.InsightName != nil && *request.InsightName != ""  {
        path = strings.ReplaceAll(path, "{insightName}", core.ToString(*request.InsightName))
    } else {
        path = strings.ReplaceAll(path, "{insightName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteInsightAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("log").AppendPath(path, replacer),
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
