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

package distributor

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2DistributorRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2DistributorRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2DistributorRestClient,
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

func (p Gs2DistributorRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) DescribeNamespaces(
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
	client Gs2DistributorRestClient,
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

func (p Gs2DistributorRestClient) CreateNamespaceAsync(
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
	if request.AssumeUserId != nil && *request.AssumeUserId != "" {
		bodies["assumeUserId"] = *request.AssumeUserId
	}
	if request.AutoRunStampSheetNotification != nil {
		bodies["autoRunStampSheetNotification"] = request.AutoRunStampSheetNotification.ToDict()
	}
	if request.AutoRunTransactionNotification != nil {
		bodies["autoRunTransactionNotification"] = request.AutoRunTransactionNotification.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
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
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) CreateNamespace(
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
	client Gs2DistributorRestClient,
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

func (p Gs2DistributorRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) GetNamespaceStatus(
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
	client Gs2DistributorRestClient,
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

func (p Gs2DistributorRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) GetNamespace(
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
	client Gs2DistributorRestClient,
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

func (p Gs2DistributorRestClient) UpdateNamespaceAsync(
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
	if request.AssumeUserId != nil && *request.AssumeUserId != "" {
		bodies["assumeUserId"] = *request.AssumeUserId
	}
	if request.AutoRunStampSheetNotification != nil {
		bodies["autoRunStampSheetNotification"] = request.AutoRunStampSheetNotification.ToDict()
	}
	if request.AutoRunTransactionNotification != nil {
		bodies["autoRunTransactionNotification"] = request.AutoRunTransactionNotification.ToDict()
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
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
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) UpdateNamespace(
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
	client Gs2DistributorRestClient,
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

func (p Gs2DistributorRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) DeleteNamespace(
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

func describeDistributorModelMastersAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDistributorModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDistributorModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDistributorModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeDistributorModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDistributorModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeDistributorModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) DescribeDistributorModelMastersAsync(
	request *DescribeDistributorModelMastersRequest,
	callback chan<- DescribeDistributorModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/distributor"
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

	go describeDistributorModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) DescribeDistributorModelMasters(
	request *DescribeDistributorModelMastersRequest,
) (*DescribeDistributorModelMastersResult, error) {
	callback := make(chan DescribeDistributorModelMastersAsyncResult, 1)
	go p.DescribeDistributorModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createDistributorModelMasterAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- CreateDistributorModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateDistributorModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateDistributorModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateDistributorModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateDistributorModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateDistributorModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) CreateDistributorModelMasterAsync(
	request *CreateDistributorModelMasterRequest,
	callback chan<- CreateDistributorModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/distributor"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.InboxNamespaceId != nil && *request.InboxNamespaceId != "" {
		bodies["inboxNamespaceId"] = *request.InboxNamespaceId
	}
	if request.WhiteListTargetIds != nil {
		var _whiteListTargetIds []interface{}
		for _, item := range request.WhiteListTargetIds {
			_whiteListTargetIds = append(_whiteListTargetIds, item)
		}
		bodies["whiteListTargetIds"] = _whiteListTargetIds
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

	go createDistributorModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) CreateDistributorModelMaster(
	request *CreateDistributorModelMasterRequest,
) (*CreateDistributorModelMasterResult, error) {
	callback := make(chan CreateDistributorModelMasterAsyncResult, 1)
	go p.CreateDistributorModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getDistributorModelMasterAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- GetDistributorModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDistributorModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDistributorModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetDistributorModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDistributorModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetDistributorModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) GetDistributorModelMasterAsync(
	request *GetDistributorModelMasterRequest,
	callback chan<- GetDistributorModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/distributor/{distributorName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.DistributorName != nil && *request.DistributorName != "" {
		path = strings.ReplaceAll(path, "{distributorName}", core.ToString(*request.DistributorName))
	} else {
		path = strings.ReplaceAll(path, "{distributorName}", "null")
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

	go getDistributorModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) GetDistributorModelMaster(
	request *GetDistributorModelMasterRequest,
) (*GetDistributorModelMasterResult, error) {
	callback := make(chan GetDistributorModelMasterAsyncResult, 1)
	go p.GetDistributorModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateDistributorModelMasterAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateDistributorModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateDistributorModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateDistributorModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateDistributorModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateDistributorModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateDistributorModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) UpdateDistributorModelMasterAsync(
	request *UpdateDistributorModelMasterRequest,
	callback chan<- UpdateDistributorModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/distributor/{distributorName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.DistributorName != nil && *request.DistributorName != "" {
		path = strings.ReplaceAll(path, "{distributorName}", core.ToString(*request.DistributorName))
	} else {
		path = strings.ReplaceAll(path, "{distributorName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.InboxNamespaceId != nil && *request.InboxNamespaceId != "" {
		bodies["inboxNamespaceId"] = *request.InboxNamespaceId
	}
	if request.WhiteListTargetIds != nil {
		var _whiteListTargetIds []interface{}
		for _, item := range request.WhiteListTargetIds {
			_whiteListTargetIds = append(_whiteListTargetIds, item)
		}
		bodies["whiteListTargetIds"] = _whiteListTargetIds
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

	go updateDistributorModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) UpdateDistributorModelMaster(
	request *UpdateDistributorModelMasterRequest,
) (*UpdateDistributorModelMasterResult, error) {
	callback := make(chan UpdateDistributorModelMasterAsyncResult, 1)
	go p.UpdateDistributorModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteDistributorModelMasterAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteDistributorModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteDistributorModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteDistributorModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteDistributorModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteDistributorModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteDistributorModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) DeleteDistributorModelMasterAsync(
	request *DeleteDistributorModelMasterRequest,
	callback chan<- DeleteDistributorModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/distributor/{distributorName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.DistributorName != nil && *request.DistributorName != "" {
		path = strings.ReplaceAll(path, "{distributorName}", core.ToString(*request.DistributorName))
	} else {
		path = strings.ReplaceAll(path, "{distributorName}", "null")
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

	go deleteDistributorModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) DeleteDistributorModelMaster(
	request *DeleteDistributorModelMasterRequest,
) (*DeleteDistributorModelMasterResult, error) {
	callback := make(chan DeleteDistributorModelMasterAsyncResult, 1)
	go p.DeleteDistributorModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeDistributorModelsAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDistributorModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDistributorModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDistributorModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeDistributorModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDistributorModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeDistributorModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) DescribeDistributorModelsAsync(
	request *DescribeDistributorModelsRequest,
	callback chan<- DescribeDistributorModelsAsyncResult,
) {
	path := "/{namespaceName}/distributor"
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

	go describeDistributorModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) DescribeDistributorModels(
	request *DescribeDistributorModelsRequest,
) (*DescribeDistributorModelsResult, error) {
	callback := make(chan DescribeDistributorModelsAsyncResult, 1)
	go p.DescribeDistributorModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getDistributorModelAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- GetDistributorModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDistributorModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDistributorModelResult
	if asyncResult.Err != nil {
		callback <- GetDistributorModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDistributorModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetDistributorModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) GetDistributorModelAsync(
	request *GetDistributorModelRequest,
	callback chan<- GetDistributorModelAsyncResult,
) {
	path := "/{namespaceName}/distributor/{distributorName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.DistributorName != nil && *request.DistributorName != "" {
		path = strings.ReplaceAll(path, "{distributorName}", core.ToString(*request.DistributorName))
	} else {
		path = strings.ReplaceAll(path, "{distributorName}", "null")
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

	go getDistributorModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) GetDistributorModel(
	request *GetDistributorModelRequest,
) (*GetDistributorModelResult, error) {
	callback := make(chan GetDistributorModelAsyncResult, 1)
	go p.GetDistributorModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- ExportMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- ExportMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2DistributorRestClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	path := "/{namespaceName}/master/export"
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

	go exportMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) ExportMaster(
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

func getCurrentDistributorMasterAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentDistributorMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentDistributorMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentDistributorMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentDistributorMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentDistributorMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentDistributorMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) GetCurrentDistributorMasterAsync(
	request *GetCurrentDistributorMasterRequest,
	callback chan<- GetCurrentDistributorMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
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

	go getCurrentDistributorMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) GetCurrentDistributorMaster(
	request *GetCurrentDistributorMasterRequest,
) (*GetCurrentDistributorMasterResult, error) {
	callback := make(chan GetCurrentDistributorMasterAsyncResult, 1)
	go p.GetCurrentDistributorMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentDistributorMasterAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentDistributorMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentDistributorMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentDistributorMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentDistributorMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentDistributorMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentDistributorMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) UpdateCurrentDistributorMasterAsync(
	request *UpdateCurrentDistributorMasterRequest,
	callback chan<- UpdateCurrentDistributorMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Settings != nil && *request.Settings != "" {
		bodies["settings"] = *request.Settings
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

	go updateCurrentDistributorMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) UpdateCurrentDistributorMaster(
	request *UpdateCurrentDistributorMasterRequest,
) (*UpdateCurrentDistributorMasterResult, error) {
	callback := make(chan UpdateCurrentDistributorMasterAsyncResult, 1)
	go p.UpdateCurrentDistributorMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentDistributorMasterFromGitHubAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentDistributorMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentDistributorMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentDistributorMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentDistributorMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentDistributorMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentDistributorMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) UpdateCurrentDistributorMasterFromGitHubAsync(
	request *UpdateCurrentDistributorMasterFromGitHubRequest,
	callback chan<- UpdateCurrentDistributorMasterFromGitHubAsyncResult,
) {
	path := "/{namespaceName}/master/from_git_hub"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.CheckoutSetting != nil {
		bodies["checkoutSetting"] = request.CheckoutSetting.ToDict()
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

	go updateCurrentDistributorMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) UpdateCurrentDistributorMasterFromGitHub(
	request *UpdateCurrentDistributorMasterFromGitHubRequest,
) (*UpdateCurrentDistributorMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentDistributorMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentDistributorMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func distributeAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- DistributeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DistributeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DistributeResult
	if asyncResult.Err != nil {
		callback <- DistributeAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DistributeAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DistributeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) DistributeAsync(
	request *DistributeRequest,
	callback chan<- DistributeAsyncResult,
) {
	path := "/{namespaceName}/distribute/{distributorName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.DistributorName != nil && *request.DistributorName != "" {
		path = strings.ReplaceAll(path, "{distributorName}", core.ToString(*request.DistributorName))
	} else {
		path = strings.ReplaceAll(path, "{distributorName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.DistributeResource != nil {
		bodies["distributeResource"] = request.DistributeResource.ToDict()
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

	go distributeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) Distribute(
	request *DistributeRequest,
) (*DistributeResult, error) {
	callback := make(chan DistributeAsyncResult, 1)
	go p.DistributeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func distributeWithoutOverflowProcessAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- DistributeWithoutOverflowProcessAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DistributeWithoutOverflowProcessAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DistributeWithoutOverflowProcessResult
	if asyncResult.Err != nil {
		callback <- DistributeWithoutOverflowProcessAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DistributeWithoutOverflowProcessAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DistributeWithoutOverflowProcessAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) DistributeWithoutOverflowProcessAsync(
	request *DistributeWithoutOverflowProcessRequest,
	callback chan<- DistributeWithoutOverflowProcessAsyncResult,
) {
	path := "/distribute"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.DistributeResource != nil {
		bodies["distributeResource"] = request.DistributeResource.ToDict()
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

	go distributeWithoutOverflowProcessAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) DistributeWithoutOverflowProcess(
	request *DistributeWithoutOverflowProcessRequest,
) (*DistributeWithoutOverflowProcessResult, error) {
	callback := make(chan DistributeWithoutOverflowProcessAsyncResult, 1)
	go p.DistributeWithoutOverflowProcessAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func runVerifyTaskAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- RunVerifyTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RunVerifyTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RunVerifyTaskResult
	if asyncResult.Err != nil {
		callback <- RunVerifyTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RunVerifyTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RunVerifyTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) RunVerifyTaskAsync(
	request *RunVerifyTaskRequest,
	callback chan<- RunVerifyTaskAsyncResult,
) {
	path := "/{namespaceName}/distribute/stamp/verifyTask/run"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.VerifyTask != nil && *request.VerifyTask != "" {
		bodies["verifyTask"] = *request.VerifyTask
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go runVerifyTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) RunVerifyTask(
	request *RunVerifyTaskRequest,
) (*RunVerifyTaskResult, error) {
	callback := make(chan RunVerifyTaskAsyncResult, 1)
	go p.RunVerifyTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func runStampTaskAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- RunStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RunStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RunStampTaskResult
	if asyncResult.Err != nil {
		callback <- RunStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RunStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RunStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) RunStampTaskAsync(
	request *RunStampTaskRequest,
	callback chan<- RunStampTaskAsyncResult,
) {
	path := "/{namespaceName}/distribute/stamp/task/run"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampTask != nil && *request.StampTask != "" {
		bodies["stampTask"] = *request.StampTask
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go runStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) RunStampTask(
	request *RunStampTaskRequest,
) (*RunStampTaskResult, error) {
	callback := make(chan RunStampTaskAsyncResult, 1)
	go p.RunStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func runStampSheetAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- RunStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RunStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RunStampSheetResult
	if asyncResult.Err != nil {
		callback <- RunStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RunStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RunStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) RunStampSheetAsync(
	request *RunStampSheetRequest,
	callback chan<- RunStampSheetAsyncResult,
) {
	path := "/{namespaceName}/distribute/stamp/sheet/run"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampSheet != nil && *request.StampSheet != "" {
		bodies["stampSheet"] = *request.StampSheet
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go runStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) RunStampSheet(
	request *RunStampSheetRequest,
) (*RunStampSheetResult, error) {
	callback := make(chan RunStampSheetAsyncResult, 1)
	go p.RunStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func runStampSheetExpressAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- RunStampSheetExpressAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RunStampSheetExpressAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RunStampSheetExpressResult
	if asyncResult.Err != nil {
		callback <- RunStampSheetExpressAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RunStampSheetExpressAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RunStampSheetExpressAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) RunStampSheetExpressAsync(
	request *RunStampSheetExpressRequest,
	callback chan<- RunStampSheetExpressAsyncResult,
) {
	path := "/{namespaceName}/distribute/stamp/run"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampSheet != nil && *request.StampSheet != "" {
		bodies["stampSheet"] = *request.StampSheet
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go runStampSheetExpressAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) RunStampSheetExpress(
	request *RunStampSheetExpressRequest,
) (*RunStampSheetExpressResult, error) {
	callback := make(chan RunStampSheetExpressAsyncResult, 1)
	go p.RunStampSheetExpressAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func runVerifyTaskWithoutNamespaceAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- RunVerifyTaskWithoutNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RunVerifyTaskWithoutNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RunVerifyTaskWithoutNamespaceResult
	if asyncResult.Err != nil {
		callback <- RunVerifyTaskWithoutNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RunVerifyTaskWithoutNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RunVerifyTaskWithoutNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) RunVerifyTaskWithoutNamespaceAsync(
	request *RunVerifyTaskWithoutNamespaceRequest,
	callback chan<- RunVerifyTaskWithoutNamespaceAsyncResult,
) {
	path := "/stamp/verifyTask/run"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.VerifyTask != nil && *request.VerifyTask != "" {
		bodies["verifyTask"] = *request.VerifyTask
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go runVerifyTaskWithoutNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) RunVerifyTaskWithoutNamespace(
	request *RunVerifyTaskWithoutNamespaceRequest,
) (*RunVerifyTaskWithoutNamespaceResult, error) {
	callback := make(chan RunVerifyTaskWithoutNamespaceAsyncResult, 1)
	go p.RunVerifyTaskWithoutNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func runStampTaskWithoutNamespaceAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- RunStampTaskWithoutNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RunStampTaskWithoutNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RunStampTaskWithoutNamespaceResult
	if asyncResult.Err != nil {
		callback <- RunStampTaskWithoutNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RunStampTaskWithoutNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RunStampTaskWithoutNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) RunStampTaskWithoutNamespaceAsync(
	request *RunStampTaskWithoutNamespaceRequest,
	callback chan<- RunStampTaskWithoutNamespaceAsyncResult,
) {
	path := "/stamp/task/run"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampTask != nil && *request.StampTask != "" {
		bodies["stampTask"] = *request.StampTask
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go runStampTaskWithoutNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) RunStampTaskWithoutNamespace(
	request *RunStampTaskWithoutNamespaceRequest,
) (*RunStampTaskWithoutNamespaceResult, error) {
	callback := make(chan RunStampTaskWithoutNamespaceAsyncResult, 1)
	go p.RunStampTaskWithoutNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func runStampSheetWithoutNamespaceAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- RunStampSheetWithoutNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RunStampSheetWithoutNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RunStampSheetWithoutNamespaceResult
	if asyncResult.Err != nil {
		callback <- RunStampSheetWithoutNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RunStampSheetWithoutNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RunStampSheetWithoutNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) RunStampSheetWithoutNamespaceAsync(
	request *RunStampSheetWithoutNamespaceRequest,
	callback chan<- RunStampSheetWithoutNamespaceAsyncResult,
) {
	path := "/stamp/sheet/run"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampSheet != nil && *request.StampSheet != "" {
		bodies["stampSheet"] = *request.StampSheet
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go runStampSheetWithoutNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) RunStampSheetWithoutNamespace(
	request *RunStampSheetWithoutNamespaceRequest,
) (*RunStampSheetWithoutNamespaceResult, error) {
	callback := make(chan RunStampSheetWithoutNamespaceAsyncResult, 1)
	go p.RunStampSheetWithoutNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func runStampSheetExpressWithoutNamespaceAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- RunStampSheetExpressWithoutNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RunStampSheetExpressWithoutNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RunStampSheetExpressWithoutNamespaceResult
	if asyncResult.Err != nil {
		callback <- RunStampSheetExpressWithoutNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RunStampSheetExpressWithoutNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RunStampSheetExpressWithoutNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) RunStampSheetExpressWithoutNamespaceAsync(
	request *RunStampSheetExpressWithoutNamespaceRequest,
	callback chan<- RunStampSheetExpressWithoutNamespaceAsyncResult,
) {
	path := "/stamp/run"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampSheet != nil && *request.StampSheet != "" {
		bodies["stampSheet"] = *request.StampSheet
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go runStampSheetExpressWithoutNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) RunStampSheetExpressWithoutNamespace(
	request *RunStampSheetExpressWithoutNamespaceRequest,
) (*RunStampSheetExpressWithoutNamespaceResult, error) {
	callback := make(chan RunStampSheetExpressWithoutNamespaceAsyncResult, 1)
	go p.RunStampSheetExpressWithoutNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setTransactionDefaultConfigAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- SetTransactionDefaultConfigAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetTransactionDefaultConfigAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetTransactionDefaultConfigResult
	if asyncResult.Err != nil {
		callback <- SetTransactionDefaultConfigAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetTransactionDefaultConfigAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetTransactionDefaultConfigAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) SetTransactionDefaultConfigAsync(
	request *SetTransactionDefaultConfigRequest,
	callback chan<- SetTransactionDefaultConfigAsyncResult,
) {
	path := "/transaction/user/me/config"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	go setTransactionDefaultConfigAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) SetTransactionDefaultConfig(
	request *SetTransactionDefaultConfigRequest,
) (*SetTransactionDefaultConfigResult, error) {
	callback := make(chan SetTransactionDefaultConfigAsyncResult, 1)
	go p.SetTransactionDefaultConfigAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setTransactionDefaultConfigByUserIdAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- SetTransactionDefaultConfigByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetTransactionDefaultConfigByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetTransactionDefaultConfigByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetTransactionDefaultConfigByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetTransactionDefaultConfigByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetTransactionDefaultConfigByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) SetTransactionDefaultConfigByUserIdAsync(
	request *SetTransactionDefaultConfigByUserIdRequest,
	callback chan<- SetTransactionDefaultConfigByUserIdAsyncResult,
) {
	path := "/transaction/user/{userId}/config"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	go setTransactionDefaultConfigByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) SetTransactionDefaultConfigByUserId(
	request *SetTransactionDefaultConfigByUserIdRequest,
) (*SetTransactionDefaultConfigByUserIdResult, error) {
	callback := make(chan SetTransactionDefaultConfigByUserIdAsyncResult, 1)
	go p.SetTransactionDefaultConfigByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func freezeMasterDataAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- FreezeMasterDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- FreezeMasterDataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result FreezeMasterDataResult
	if asyncResult.Err != nil {
		callback <- FreezeMasterDataAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- FreezeMasterDataAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- FreezeMasterDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) FreezeMasterDataAsync(
	request *FreezeMasterDataRequest,
	callback chan<- FreezeMasterDataAsyncResult,
) {
	path := "/{namespaceName}/user/me/masterdata/freeze"
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

	go freezeMasterDataAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) FreezeMasterData(
	request *FreezeMasterDataRequest,
) (*FreezeMasterDataResult, error) {
	callback := make(chan FreezeMasterDataAsyncResult, 1)
	go p.FreezeMasterDataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func freezeMasterDataByUserIdAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- FreezeMasterDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- FreezeMasterDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result FreezeMasterDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- FreezeMasterDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- FreezeMasterDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- FreezeMasterDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) FreezeMasterDataByUserIdAsync(
	request *FreezeMasterDataByUserIdRequest,
	callback chan<- FreezeMasterDataByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/masterdata/freeze"
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

	go freezeMasterDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) FreezeMasterDataByUserId(
	request *FreezeMasterDataByUserIdRequest,
) (*FreezeMasterDataByUserIdResult, error) {
	callback := make(chan FreezeMasterDataByUserIdAsyncResult, 1)
	go p.FreezeMasterDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func signFreezeMasterDataTimestampAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- SignFreezeMasterDataTimestampAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SignFreezeMasterDataTimestampAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SignFreezeMasterDataTimestampResult
	if asyncResult.Err != nil {
		callback <- SignFreezeMasterDataTimestampAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SignFreezeMasterDataTimestampAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SignFreezeMasterDataTimestampAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) SignFreezeMasterDataTimestampAsync(
	request *SignFreezeMasterDataTimestampRequest,
	callback chan<- SignFreezeMasterDataTimestampAsyncResult,
) {
	path := "/{namespaceName}/masterdata/freeze/timestamp"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Timestamp != nil {
		bodies["timestamp"] = *request.Timestamp
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go signFreezeMasterDataTimestampAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) SignFreezeMasterDataTimestamp(
	request *SignFreezeMasterDataTimestampRequest,
) (*SignFreezeMasterDataTimestampResult, error) {
	callback := make(chan SignFreezeMasterDataTimestampAsyncResult, 1)
	go p.SignFreezeMasterDataTimestampAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func freezeMasterDataBySignedTimestampAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- FreezeMasterDataBySignedTimestampAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- FreezeMasterDataBySignedTimestampAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result FreezeMasterDataBySignedTimestampResult
	if asyncResult.Err != nil {
		callback <- FreezeMasterDataBySignedTimestampAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- FreezeMasterDataBySignedTimestampAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- FreezeMasterDataBySignedTimestampAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) FreezeMasterDataBySignedTimestampAsync(
	request *FreezeMasterDataBySignedTimestampRequest,
	callback chan<- FreezeMasterDataBySignedTimestampAsyncResult,
) {
	path := "/{namespaceName}/user/me/masterdata/freeze/timestamp"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Body != nil && *request.Body != "" {
		bodies["body"] = *request.Body
	}
	if request.Signature != nil && *request.Signature != "" {
		bodies["signature"] = *request.Signature
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go freezeMasterDataBySignedTimestampAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) FreezeMasterDataBySignedTimestamp(
	request *FreezeMasterDataBySignedTimestampRequest,
) (*FreezeMasterDataBySignedTimestampResult, error) {
	callback := make(chan FreezeMasterDataBySignedTimestampAsyncResult, 1)
	go p.FreezeMasterDataBySignedTimestampAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func freezeMasterDataByTimestampAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- FreezeMasterDataByTimestampAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- FreezeMasterDataByTimestampAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result FreezeMasterDataByTimestampResult
	if asyncResult.Err != nil {
		callback <- FreezeMasterDataByTimestampAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- FreezeMasterDataByTimestampAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- FreezeMasterDataByTimestampAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) FreezeMasterDataByTimestampAsync(
	request *FreezeMasterDataByTimestampRequest,
	callback chan<- FreezeMasterDataByTimestampAsyncResult,
) {
	path := "/{namespaceName}/user/me/masterdata/freeze/timestamp/raw"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Timestamp != nil {
		bodies["timestamp"] = *request.Timestamp
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

	go freezeMasterDataByTimestampAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) FreezeMasterDataByTimestamp(
	request *FreezeMasterDataByTimestampRequest,
) (*FreezeMasterDataByTimestampResult, error) {
	callback := make(chan FreezeMasterDataByTimestampAsyncResult, 1)
	go p.FreezeMasterDataByTimestampAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func batchExecuteApiAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- BatchExecuteApiAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BatchExecuteApiAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BatchExecuteApiResult
	if asyncResult.Err != nil {
		callback <- BatchExecuteApiAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- BatchExecuteApiAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- BatchExecuteApiAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) BatchExecuteApiAsync(
	request *BatchExecuteApiRequest,
	callback chan<- BatchExecuteApiAsyncResult,
) {
	path := "/batch/execute"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RequestPayloads != nil {
		var _requestPayloads []interface{}
		for _, item := range request.RequestPayloads {
			_requestPayloads = append(_requestPayloads, item)
		}
		bodies["requestPayloads"] = _requestPayloads
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

	go batchExecuteApiAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) BatchExecuteApi(
	request *BatchExecuteApiRequest,
) (*BatchExecuteApiResult, error) {
	callback := make(chan BatchExecuteApiAsyncResult, 1)
	go p.BatchExecuteApiAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func ifExpressionByUserIdAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- IfExpressionByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- IfExpressionByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result IfExpressionByUserIdResult
	if asyncResult.Err != nil {
		callback <- IfExpressionByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IfExpressionByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- IfExpressionByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) IfExpressionByUserIdAsync(
	request *IfExpressionByUserIdRequest,
	callback chan<- IfExpressionByUserIdAsyncResult,
) {
	path := "/{namespaceName}/expression/if"
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
	if request.Condition != nil {
		bodies["condition"] = request.Condition.ToDict()
	}
	if request.TrueActions != nil {
		var _trueActions []interface{}
		for _, item := range request.TrueActions {
			_trueActions = append(_trueActions, item)
		}
		bodies["trueActions"] = _trueActions
	}
	if request.FalseActions != nil {
		var _falseActions []interface{}
		for _, item := range request.FalseActions {
			_falseActions = append(_falseActions, item)
		}
		bodies["falseActions"] = _falseActions
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go ifExpressionByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) IfExpressionByUserId(
	request *IfExpressionByUserIdRequest,
) (*IfExpressionByUserIdResult, error) {
	callback := make(chan IfExpressionByUserIdAsyncResult, 1)
	go p.IfExpressionByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func andExpressionByUserIdAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- AndExpressionByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AndExpressionByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AndExpressionByUserIdResult
	if asyncResult.Err != nil {
		callback <- AndExpressionByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AndExpressionByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AndExpressionByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) AndExpressionByUserIdAsync(
	request *AndExpressionByUserIdRequest,
	callback chan<- AndExpressionByUserIdAsyncResult,
) {
	path := "/{namespaceName}/expression/and"
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
	if request.Actions != nil {
		var _actions []interface{}
		for _, item := range request.Actions {
			_actions = append(_actions, item)
		}
		bodies["actions"] = _actions
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

	go andExpressionByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) AndExpressionByUserId(
	request *AndExpressionByUserIdRequest,
) (*AndExpressionByUserIdResult, error) {
	callback := make(chan AndExpressionByUserIdAsyncResult, 1)
	go p.AndExpressionByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func orExpressionByUserIdAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- OrExpressionByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- OrExpressionByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result OrExpressionByUserIdResult
	if asyncResult.Err != nil {
		callback <- OrExpressionByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- OrExpressionByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- OrExpressionByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) OrExpressionByUserIdAsync(
	request *OrExpressionByUserIdRequest,
	callback chan<- OrExpressionByUserIdAsyncResult,
) {
	path := "/{namespaceName}/expression/or"
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
	if request.Actions != nil {
		var _actions []interface{}
		for _, item := range request.Actions {
			_actions = append(_actions, item)
		}
		bodies["actions"] = _actions
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

	go orExpressionByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) OrExpressionByUserId(
	request *OrExpressionByUserIdRequest,
) (*OrExpressionByUserIdResult, error) {
	callback := make(chan OrExpressionByUserIdAsyncResult, 1)
	go p.OrExpressionByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func ifExpressionByStampTaskAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- IfExpressionByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- IfExpressionByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result IfExpressionByStampTaskResult
	if asyncResult.Err != nil {
		callback <- IfExpressionByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IfExpressionByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- IfExpressionByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) IfExpressionByStampTaskAsync(
	request *IfExpressionByStampTaskRequest,
	callback chan<- IfExpressionByStampTaskAsyncResult,
) {
	path := "/stamp/expression/if"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampTask != nil && *request.StampTask != "" {
		bodies["stampTask"] = *request.StampTask
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go ifExpressionByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) IfExpressionByStampTask(
	request *IfExpressionByStampTaskRequest,
) (*IfExpressionByStampTaskResult, error) {
	callback := make(chan IfExpressionByStampTaskAsyncResult, 1)
	go p.IfExpressionByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func andExpressionByStampTaskAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- AndExpressionByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AndExpressionByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AndExpressionByStampTaskResult
	if asyncResult.Err != nil {
		callback <- AndExpressionByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AndExpressionByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AndExpressionByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) AndExpressionByStampTaskAsync(
	request *AndExpressionByStampTaskRequest,
	callback chan<- AndExpressionByStampTaskAsyncResult,
) {
	path := "/stamp/expression/and"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampTask != nil && *request.StampTask != "" {
		bodies["stampTask"] = *request.StampTask
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go andExpressionByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) AndExpressionByStampTask(
	request *AndExpressionByStampTaskRequest,
) (*AndExpressionByStampTaskResult, error) {
	callback := make(chan AndExpressionByStampTaskAsyncResult, 1)
	go p.AndExpressionByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func orExpressionByStampTaskAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- OrExpressionByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- OrExpressionByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result OrExpressionByStampTaskResult
	if asyncResult.Err != nil {
		callback <- OrExpressionByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- OrExpressionByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- OrExpressionByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) OrExpressionByStampTaskAsync(
	request *OrExpressionByStampTaskRequest,
	callback chan<- OrExpressionByStampTaskAsyncResult,
) {
	path := "/stamp/expression/or"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampTask != nil && *request.StampTask != "" {
		bodies["stampTask"] = *request.StampTask
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
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

	go orExpressionByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) OrExpressionByStampTask(
	request *OrExpressionByStampTaskRequest,
) (*OrExpressionByStampTaskResult, error) {
	callback := make(chan OrExpressionByStampTaskAsyncResult, 1)
	go p.OrExpressionByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStampSheetResultAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- GetStampSheetResultAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStampSheetResultAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStampSheetResultResult
	if asyncResult.Err != nil {
		callback <- GetStampSheetResultAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStampSheetResultAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStampSheetResultAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) GetStampSheetResultAsync(
	request *GetStampSheetResultRequest,
	callback chan<- GetStampSheetResultAsyncResult,
) {
	path := "/{namespaceName}/user/me/stampSheet/{transactionId}/result"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		path = strings.ReplaceAll(path, "{transactionId}", core.ToString(*request.TransactionId))
	} else {
		path = strings.ReplaceAll(path, "{transactionId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getStampSheetResultAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) GetStampSheetResult(
	request *GetStampSheetResultRequest,
) (*GetStampSheetResultResult, error) {
	callback := make(chan GetStampSheetResultAsyncResult, 1)
	go p.GetStampSheetResultAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStampSheetResultByUserIdAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- GetStampSheetResultByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStampSheetResultByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStampSheetResultByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetStampSheetResultByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStampSheetResultByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStampSheetResultByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) GetStampSheetResultByUserIdAsync(
	request *GetStampSheetResultByUserIdRequest,
	callback chan<- GetStampSheetResultByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stampSheet/{transactionId}/result"
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
	if request.TransactionId != nil && *request.TransactionId != "" {
		path = strings.ReplaceAll(path, "{transactionId}", core.ToString(*request.TransactionId))
	} else {
		path = strings.ReplaceAll(path, "{transactionId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
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

	go getStampSheetResultByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) GetStampSheetResultByUserId(
	request *GetStampSheetResultByUserIdRequest,
) (*GetStampSheetResultByUserIdResult, error) {
	callback := make(chan GetStampSheetResultByUserIdAsyncResult, 1)
	go p.GetStampSheetResultByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func runTransactionAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- RunTransactionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RunTransactionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RunTransactionResult
	if asyncResult.Err != nil {
		callback <- RunTransactionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RunTransactionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RunTransactionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) RunTransactionAsync(
	request *RunTransactionRequest,
	callback chan<- RunTransactionAsyncResult,
) {
	path := "/system/{ownerId}/{namespaceName}/user/{userId}/transaction/run"
	if request.OwnerId != nil && *request.OwnerId != "" {
		path = strings.ReplaceAll(path, "{ownerId}", core.ToString(*request.OwnerId))
	} else {
		path = strings.ReplaceAll(path, "{ownerId}", "null")
	}
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
	if request.Transaction != nil && *request.Transaction != "" {
		bodies["transaction"] = *request.Transaction
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

	go runTransactionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) RunTransaction(
	request *RunTransactionRequest,
) (*RunTransactionResult, error) {
	callback := make(chan RunTransactionAsyncResult, 1)
	go p.RunTransactionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getTransactionResultAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- GetTransactionResultAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetTransactionResultAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetTransactionResultResult
	if asyncResult.Err != nil {
		callback <- GetTransactionResultAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetTransactionResultAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetTransactionResultAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) GetTransactionResultAsync(
	request *GetTransactionResultRequest,
	callback chan<- GetTransactionResultAsyncResult,
) {
	path := "/{namespaceName}/user/me/transaction/{transactionId}/result"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		path = strings.ReplaceAll(path, "{transactionId}", core.ToString(*request.TransactionId))
	} else {
		path = strings.ReplaceAll(path, "{transactionId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getTransactionResultAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) GetTransactionResult(
	request *GetTransactionResultRequest,
) (*GetTransactionResultResult, error) {
	callback := make(chan GetTransactionResultAsyncResult, 1)
	go p.GetTransactionResultAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getTransactionResultByUserIdAsyncHandler(
	client Gs2DistributorRestClient,
	job *core.NetworkJob,
	callback chan<- GetTransactionResultByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetTransactionResultByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetTransactionResultByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetTransactionResultByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetTransactionResultByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetTransactionResultByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DistributorRestClient) GetTransactionResultByUserIdAsync(
	request *GetTransactionResultByUserIdRequest,
	callback chan<- GetTransactionResultByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/transaction/{transactionId}/result"
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
	if request.TransactionId != nil && *request.TransactionId != "" {
		path = strings.ReplaceAll(path, "{transactionId}", core.ToString(*request.TransactionId))
	} else {
		path = strings.ReplaceAll(path, "{transactionId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
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

	go getTransactionResultByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("distributor").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DistributorRestClient) GetTransactionResultByUserId(
	request *GetTransactionResultByUserIdRequest,
) (*GetTransactionResultByUserIdResult, error) {
	callback := make(chan GetTransactionResultByUserIdAsyncResult, 1)
	go p.GetTransactionResultByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
