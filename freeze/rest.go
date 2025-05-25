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

package freeze

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2FreezeRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2FreezeRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeStagesAsyncHandler(
	client Gs2FreezeRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStagesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStagesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStagesResult
	if asyncResult.Err != nil {
		callback <- DescribeStagesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStagesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStagesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FreezeRestClient) DescribeStagesAsync(
	request *DescribeStagesRequest,
	callback chan<- DescribeStagesAsyncResult,
) {
	path := "/"

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

	go describeStagesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("freeze").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FreezeRestClient) DescribeStages(
	request *DescribeStagesRequest,
) (*DescribeStagesResult, error) {
	callback := make(chan DescribeStagesAsyncResult, 1)
	go p.DescribeStagesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStageAsyncHandler(
	client Gs2FreezeRestClient,
	job *core.NetworkJob,
	callback chan<- GetStageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStageResult
	if asyncResult.Err != nil {
		callback <- GetStageAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStageAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FreezeRestClient) GetStageAsync(
	request *GetStageRequest,
	callback chan<- GetStageAsyncResult,
) {
	path := "/{stageName}"
	if request.StageName != nil && *request.StageName != "" {
		path = strings.ReplaceAll(path, "{stageName}", core.ToString(*request.StageName))
	} else {
		path = strings.ReplaceAll(path, "{stageName}", "null")
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

	go getStageAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("freeze").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FreezeRestClient) GetStage(
	request *GetStageRequest,
) (*GetStageResult, error) {
	callback := make(chan GetStageAsyncResult, 1)
	go p.GetStageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func promoteStageAsyncHandler(
	client Gs2FreezeRestClient,
	job *core.NetworkJob,
	callback chan<- PromoteStageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PromoteStageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PromoteStageResult
	if asyncResult.Err != nil {
		callback <- PromoteStageAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PromoteStageAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PromoteStageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FreezeRestClient) PromoteStageAsync(
	request *PromoteStageRequest,
	callback chan<- PromoteStageAsyncResult,
) {
	path := "/{stageName}/promote"
	if request.StageName != nil && *request.StageName != "" {
		path = strings.ReplaceAll(path, "{stageName}", core.ToString(*request.StageName))
	} else {
		path = strings.ReplaceAll(path, "{stageName}", "null")
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

	go promoteStageAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("freeze").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FreezeRestClient) PromoteStage(
	request *PromoteStageRequest,
) (*PromoteStageResult, error) {
	callback := make(chan PromoteStageAsyncResult, 1)
	go p.PromoteStageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func rollbackStageAsyncHandler(
	client Gs2FreezeRestClient,
	job *core.NetworkJob,
	callback chan<- RollbackStageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RollbackStageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RollbackStageResult
	if asyncResult.Err != nil {
		callback <- RollbackStageAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RollbackStageAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RollbackStageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FreezeRestClient) RollbackStageAsync(
	request *RollbackStageRequest,
	callback chan<- RollbackStageAsyncResult,
) {
	path := "/{stageName}/rollback"
	if request.StageName != nil && *request.StageName != "" {
		path = strings.ReplaceAll(path, "{stageName}", core.ToString(*request.StageName))
	} else {
		path = strings.ReplaceAll(path, "{stageName}", "null")
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

	go rollbackStageAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("freeze").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FreezeRestClient) RollbackStage(
	request *RollbackStageRequest,
) (*RollbackStageResult, error) {
	callback := make(chan RollbackStageAsyncResult, 1)
	go p.RollbackStageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeOutputsAsyncHandler(
	client Gs2FreezeRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeOutputsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeOutputsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeOutputsResult
	if asyncResult.Err != nil {
		callback <- DescribeOutputsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeOutputsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeOutputsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FreezeRestClient) DescribeOutputsAsync(
	request *DescribeOutputsRequest,
	callback chan<- DescribeOutputsAsyncResult,
) {
	path := "/{stageName}/progress/output"
	if request.StageName != nil && *request.StageName != "" {
		path = strings.ReplaceAll(path, "{stageName}", core.ToString(*request.StageName))
	} else {
		path = strings.ReplaceAll(path, "{stageName}", "null")
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

	go describeOutputsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("freeze").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FreezeRestClient) DescribeOutputs(
	request *DescribeOutputsRequest,
) (*DescribeOutputsResult, error) {
	callback := make(chan DescribeOutputsAsyncResult, 1)
	go p.DescribeOutputsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getOutputAsyncHandler(
	client Gs2FreezeRestClient,
	job *core.NetworkJob,
	callback chan<- GetOutputAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetOutputAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetOutputResult
	if asyncResult.Err != nil {
		callback <- GetOutputAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetOutputAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetOutputAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FreezeRestClient) GetOutputAsync(
	request *GetOutputRequest,
	callback chan<- GetOutputAsyncResult,
) {
	path := "/{stageName}/progress/output/{outputName}"
	if request.StageName != nil && *request.StageName != "" {
		path = strings.ReplaceAll(path, "{stageName}", core.ToString(*request.StageName))
	} else {
		path = strings.ReplaceAll(path, "{stageName}", "null")
	}
	if request.OutputName != nil && *request.OutputName != "" {
		path = strings.ReplaceAll(path, "{outputName}", core.ToString(*request.OutputName))
	} else {
		path = strings.ReplaceAll(path, "{outputName}", "null")
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

	go getOutputAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("freeze").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FreezeRestClient) GetOutput(
	request *GetOutputRequest,
) (*GetOutputResult, error) {
	callback := make(chan GetOutputAsyncResult, 1)
	go p.GetOutputAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
