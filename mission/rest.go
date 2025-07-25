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
	"net/http"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2MissionRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2MissionRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeCompletesAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCompletesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeCompletesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCompletesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCompletesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DescribeCompletesAsync(
	request *DescribeCompletesRequest,
	callback chan<- DescribeCompletesAsyncResult,
) {
	path := "/{namespaceName}/user/me/complete"
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

	go describeCompletesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DescribeCompletes(
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

func describeCompletesByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCompletesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeCompletesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCompletesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCompletesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DescribeCompletesByUserIdAsync(
	request *DescribeCompletesByUserIdRequest,
	callback chan<- DescribeCompletesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/complete"
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

	go describeCompletesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DescribeCompletesByUserId(
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

func completeAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- CompleteAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CompleteAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CompleteAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CompleteAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) CompleteAsync(
	request *CompleteRequest,
	callback chan<- CompleteAsyncResult,
) {
	path := "/{namespaceName}/user/me/complete/group/{missionGroupName}/task/{missionTaskName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		path = strings.ReplaceAll(path, "{missionTaskName}", core.ToString(*request.MissionTaskName))
	} else {
		path = strings.ReplaceAll(path, "{missionTaskName}", "null")
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

	go completeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) Complete(
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

func completeByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- CompleteByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CompleteByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CompleteByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CompleteByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) CompleteByUserIdAsync(
	request *CompleteByUserIdRequest,
	callback chan<- CompleteByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/complete/group/{missionGroupName}/task/{missionTaskName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		path = strings.ReplaceAll(path, "{missionTaskName}", core.ToString(*request.MissionTaskName))
	} else {
		path = strings.ReplaceAll(path, "{missionTaskName}", "null")
	}
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

	go completeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) CompleteByUserId(
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

func batchCompleteAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- BatchCompleteAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BatchCompleteAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BatchCompleteResult
	if asyncResult.Err != nil {
		callback <- BatchCompleteAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- BatchCompleteAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- BatchCompleteAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) BatchCompleteAsync(
	request *BatchCompleteRequest,
	callback chan<- BatchCompleteAsyncResult,
) {
	path := "/{namespaceName}/user/me/complete/group/{missionGroupName}/task/any/batch"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.MissionTaskNames != nil {
		var _missionTaskNames []interface{}
		for _, item := range request.MissionTaskNames {
			_missionTaskNames = append(_missionTaskNames, item)
		}
		bodies["missionTaskNames"] = _missionTaskNames
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

	go batchCompleteAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) BatchComplete(
	request *BatchCompleteRequest,
) (*BatchCompleteResult, error) {
	callback := make(chan BatchCompleteAsyncResult, 1)
	go p.BatchCompleteAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func batchCompleteByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- BatchCompleteByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BatchCompleteByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BatchCompleteByUserIdResult
	if asyncResult.Err != nil {
		callback <- BatchCompleteByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- BatchCompleteByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- BatchCompleteByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) BatchCompleteByUserIdAsync(
	request *BatchCompleteByUserIdRequest,
	callback chan<- BatchCompleteByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/complete/group/{missionGroupName}/task/any/batch"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.MissionTaskNames != nil {
		var _missionTaskNames []interface{}
		for _, item := range request.MissionTaskNames {
			_missionTaskNames = append(_missionTaskNames, item)
		}
		bodies["missionTaskNames"] = _missionTaskNames
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

	go batchCompleteByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) BatchCompleteByUserId(
	request *BatchCompleteByUserIdRequest,
) (*BatchCompleteByUserIdResult, error) {
	callback := make(chan BatchCompleteByUserIdAsyncResult, 1)
	go p.BatchCompleteByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func receiveByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- ReceiveByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- ReceiveByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReceiveByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) ReceiveByUserIdAsync(
	request *ReceiveByUserIdRequest,
	callback chan<- ReceiveByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/complete/group/{missionGroupName}/task/{missionTaskName}/receive"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		path = strings.ReplaceAll(path, "{missionTaskName}", core.ToString(*request.MissionTaskName))
	} else {
		path = strings.ReplaceAll(path, "{missionTaskName}", "null")
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

	go receiveByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) ReceiveByUserId(
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

func batchReceiveByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- BatchReceiveByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BatchReceiveByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BatchReceiveByUserIdResult
	if asyncResult.Err != nil {
		callback <- BatchReceiveByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- BatchReceiveByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- BatchReceiveByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) BatchReceiveByUserIdAsync(
	request *BatchReceiveByUserIdRequest,
	callback chan<- BatchReceiveByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/complete/group/{missionGroupName}/task/any/receive/batch"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.MissionTaskNames != nil {
		var _missionTaskNames []interface{}
		for _, item := range request.MissionTaskNames {
			_missionTaskNames = append(_missionTaskNames, item)
		}
		bodies["missionTaskNames"] = _missionTaskNames
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

	go batchReceiveByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) BatchReceiveByUserId(
	request *BatchReceiveByUserIdRequest,
) (*BatchReceiveByUserIdResult, error) {
	callback := make(chan BatchReceiveByUserIdAsyncResult, 1)
	go p.BatchReceiveByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func revertReceiveByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- RevertReceiveByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- RevertReceiveByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RevertReceiveByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RevertReceiveByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) RevertReceiveByUserIdAsync(
	request *RevertReceiveByUserIdRequest,
	callback chan<- RevertReceiveByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/complete/group/{missionGroupName}/task/{missionTaskName}/revert"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		path = strings.ReplaceAll(path, "{missionTaskName}", core.ToString(*request.MissionTaskName))
	} else {
		path = strings.ReplaceAll(path, "{missionTaskName}", "null")
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

	go revertReceiveByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) RevertReceiveByUserId(
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

func getCompleteAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- GetCompleteAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetCompleteAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCompleteAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCompleteAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) GetCompleteAsync(
	request *GetCompleteRequest,
	callback chan<- GetCompleteAsyncResult,
) {
	path := "/{namespaceName}/user/me/complete/group/{missionGroupName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
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

	go getCompleteAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetComplete(
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

func getCompleteByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- GetCompleteByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetCompleteByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCompleteByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCompleteByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) GetCompleteByUserIdAsync(
	request *GetCompleteByUserIdRequest,
	callback chan<- GetCompleteByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/complete/group/{missionGroupName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
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

	go getCompleteByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetCompleteByUserId(
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

func evaluateCompleteAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- EvaluateCompleteAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- EvaluateCompleteAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result EvaluateCompleteResult
	if asyncResult.Err != nil {
		callback <- EvaluateCompleteAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- EvaluateCompleteAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- EvaluateCompleteAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) EvaluateCompleteAsync(
	request *EvaluateCompleteRequest,
	callback chan<- EvaluateCompleteAsyncResult,
) {
	path := "/{namespaceName}/user/me/complete/group/{missionGroupName}/eval"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
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

	go evaluateCompleteAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) EvaluateComplete(
	request *EvaluateCompleteRequest,
) (*EvaluateCompleteResult, error) {
	callback := make(chan EvaluateCompleteAsyncResult, 1)
	go p.EvaluateCompleteAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func evaluateCompleteByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- EvaluateCompleteByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- EvaluateCompleteByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result EvaluateCompleteByUserIdResult
	if asyncResult.Err != nil {
		callback <- EvaluateCompleteByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- EvaluateCompleteByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- EvaluateCompleteByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) EvaluateCompleteByUserIdAsync(
	request *EvaluateCompleteByUserIdRequest,
	callback chan<- EvaluateCompleteByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/complete/group/{missionGroupName}/eval"
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
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

	go evaluateCompleteByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) EvaluateCompleteByUserId(
	request *EvaluateCompleteByUserIdRequest,
) (*EvaluateCompleteByUserIdResult, error) {
	callback := make(chan EvaluateCompleteByUserIdAsyncResult, 1)
	go p.EvaluateCompleteByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteCompleteByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteCompleteByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DeleteCompleteByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteCompleteByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteCompleteByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DeleteCompleteByUserIdAsync(
	request *DeleteCompleteByUserIdRequest,
	callback chan<- DeleteCompleteByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/complete/group/{missionGroupName}"
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
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go deleteCompleteByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DeleteCompleteByUserId(
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

func verifyCompleteAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyCompleteAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCompleteAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCompleteResult
	if asyncResult.Err != nil {
		callback <- VerifyCompleteAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCompleteAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyCompleteAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) VerifyCompleteAsync(
	request *VerifyCompleteRequest,
	callback chan<- VerifyCompleteAsyncResult,
) {
	path := "/{namespaceName}/user/me/complete/group/{missionGroupName}/task/{missionTaskName}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		path = strings.ReplaceAll(path, "{missionTaskName}", core.ToString(*request.MissionTaskName))
	} else {
		path = strings.ReplaceAll(path, "{missionTaskName}", "null")
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		path = strings.ReplaceAll(path, "{multiplyValueSpecifyingQuantity}", core.ToString(*request.MultiplyValueSpecifyingQuantity))
	} else {
		path = strings.ReplaceAll(path, "{multiplyValueSpecifyingQuantity}", "null")
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

	go verifyCompleteAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) VerifyComplete(
	request *VerifyCompleteRequest,
) (*VerifyCompleteResult, error) {
	callback := make(chan VerifyCompleteAsyncResult, 1)
	go p.VerifyCompleteAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyCompleteByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyCompleteByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCompleteByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCompleteByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyCompleteByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCompleteByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyCompleteByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) VerifyCompleteByUserIdAsync(
	request *VerifyCompleteByUserIdRequest,
	callback chan<- VerifyCompleteByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/complete/group/{missionGroupName}/task/{missionTaskName}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		path = strings.ReplaceAll(path, "{missionTaskName}", core.ToString(*request.MissionTaskName))
	} else {
		path = strings.ReplaceAll(path, "{missionTaskName}", "null")
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		path = strings.ReplaceAll(path, "{multiplyValueSpecifyingQuantity}", core.ToString(*request.MultiplyValueSpecifyingQuantity))
	} else {
		path = strings.ReplaceAll(path, "{multiplyValueSpecifyingQuantity}", "null")
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

	go verifyCompleteByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) VerifyCompleteByUserId(
	request *VerifyCompleteByUserIdRequest,
) (*VerifyCompleteByUserIdResult, error) {
	callback := make(chan VerifyCompleteByUserIdAsyncResult, 1)
	go p.VerifyCompleteByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func receiveByStampTaskAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- ReceiveByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- ReceiveByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReceiveByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) ReceiveByStampTaskAsync(
	request *ReceiveByStampTaskRequest,
	callback chan<- ReceiveByStampTaskAsyncResult,
) {
	path := "/stamp/receive"

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

	go receiveByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) ReceiveByStampTask(
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

func batchReceiveByStampTaskAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- BatchReceiveByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BatchReceiveByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BatchReceiveByStampTaskResult
	if asyncResult.Err != nil {
		callback <- BatchReceiveByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- BatchReceiveByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- BatchReceiveByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) BatchReceiveByStampTaskAsync(
	request *BatchReceiveByStampTaskRequest,
	callback chan<- BatchReceiveByStampTaskAsyncResult,
) {
	path := "/stamp/receive/batch"

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

	go batchReceiveByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) BatchReceiveByStampTask(
	request *BatchReceiveByStampTaskRequest,
) (*BatchReceiveByStampTaskResult, error) {
	callback := make(chan BatchReceiveByStampTaskAsyncResult, 1)
	go p.BatchReceiveByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func revertReceiveByStampSheetAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- RevertReceiveByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- RevertReceiveByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RevertReceiveByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RevertReceiveByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) RevertReceiveByStampSheetAsync(
	request *RevertReceiveByStampSheetRequest,
	callback chan<- RevertReceiveByStampSheetAsyncResult,
) {
	path := "/stamp/revert"

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

	go revertReceiveByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) RevertReceiveByStampSheet(
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

func verifyCompleteByStampTaskAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyCompleteByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCompleteByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCompleteByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyCompleteByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCompleteByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyCompleteByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) VerifyCompleteByStampTaskAsync(
	request *VerifyCompleteByStampTaskRequest,
	callback chan<- VerifyCompleteByStampTaskAsyncResult,
) {
	path := "/stamp/complete/verify"

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

	go verifyCompleteByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) VerifyCompleteByStampTask(
	request *VerifyCompleteByStampTaskRequest,
) (*VerifyCompleteByStampTaskResult, error) {
	callback := make(chan VerifyCompleteByStampTaskAsyncResult, 1)
	go p.VerifyCompleteByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeCounterModelMastersAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCounterModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeCounterModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCounterModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCounterModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DescribeCounterModelMastersAsync(
	request *DescribeCounterModelMastersRequest,
	callback chan<- DescribeCounterModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/counter"
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

	go describeCounterModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DescribeCounterModelMasters(
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

func createCounterModelMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- CreateCounterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CreateCounterModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateCounterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateCounterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) CreateCounterModelMasterAsync(
	request *CreateCounterModelMasterRequest,
	callback chan<- CreateCounterModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/counter"
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go createCounterModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) CreateCounterModelMaster(
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

func getCounterModelMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- GetCounterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetCounterModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCounterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCounterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) GetCounterModelMasterAsync(
	request *GetCounterModelMasterRequest,
	callback chan<- GetCounterModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/counter/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
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

	go getCounterModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetCounterModelMaster(
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

func updateCounterModelMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCounterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- UpdateCounterModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCounterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCounterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) UpdateCounterModelMasterAsync(
	request *UpdateCounterModelMasterRequest,
	callback chan<- UpdateCounterModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/counter/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go updateCounterModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) UpdateCounterModelMaster(
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

func deleteCounterModelMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteCounterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DeleteCounterModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteCounterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteCounterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DeleteCounterModelMasterAsync(
	request *DeleteCounterModelMasterRequest,
	callback chan<- DeleteCounterModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/counter/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
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

	go deleteCounterModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DeleteCounterModelMaster(
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

func describeMissionGroupModelMastersAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMissionGroupModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeMissionGroupModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionGroupModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMissionGroupModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DescribeMissionGroupModelMastersAsync(
	request *DescribeMissionGroupModelMastersRequest,
	callback chan<- DescribeMissionGroupModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/group"
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

	go describeMissionGroupModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DescribeMissionGroupModelMasters(
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

func createMissionGroupModelMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- CreateMissionGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CreateMissionGroupModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateMissionGroupModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateMissionGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) CreateMissionGroupModelMasterAsync(
	request *CreateMissionGroupModelMasterRequest,
	callback chan<- CreateMissionGroupModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group"
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
	if request.AnchorTimestamp != nil {
		bodies["anchorTimestamp"] = *request.AnchorTimestamp
	}
	if request.Days != nil {
		bodies["days"] = *request.Days
	}
	if request.CompleteNotificationNamespaceId != nil && *request.CompleteNotificationNamespaceId != "" {
		bodies["completeNotificationNamespaceId"] = *request.CompleteNotificationNamespaceId
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

	go createMissionGroupModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) CreateMissionGroupModelMaster(
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

func getMissionGroupModelMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- GetMissionGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetMissionGroupModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMissionGroupModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMissionGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) GetMissionGroupModelMasterAsync(
	request *GetMissionGroupModelMasterRequest,
	callback chan<- GetMissionGroupModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{missionGroupName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
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

	go getMissionGroupModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetMissionGroupModelMaster(
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

func updateMissionGroupModelMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateMissionGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- UpdateMissionGroupModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMissionGroupModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateMissionGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) UpdateMissionGroupModelMasterAsync(
	request *UpdateMissionGroupModelMasterRequest,
	callback chan<- UpdateMissionGroupModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{missionGroupName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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
	if request.AnchorTimestamp != nil {
		bodies["anchorTimestamp"] = *request.AnchorTimestamp
	}
	if request.Days != nil {
		bodies["days"] = *request.Days
	}
	if request.CompleteNotificationNamespaceId != nil && *request.CompleteNotificationNamespaceId != "" {
		bodies["completeNotificationNamespaceId"] = *request.CompleteNotificationNamespaceId
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

	go updateMissionGroupModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) UpdateMissionGroupModelMaster(
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

func deleteMissionGroupModelMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteMissionGroupModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DeleteMissionGroupModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMissionGroupModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteMissionGroupModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DeleteMissionGroupModelMasterAsync(
	request *DeleteMissionGroupModelMasterRequest,
	callback chan<- DeleteMissionGroupModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{missionGroupName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
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

	go deleteMissionGroupModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DeleteMissionGroupModelMaster(
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

func describeNamespacesAsyncHandler(
	client Gs2MissionRestClient,
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

func (p Gs2MissionRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DescribeNamespaces(
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
	client Gs2MissionRestClient,
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

func (p Gs2MissionRestClient) CreateNamespaceAsync(
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
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) CreateNamespace(
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
	client Gs2MissionRestClient,
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

func (p Gs2MissionRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetNamespaceStatus(
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
	client Gs2MissionRestClient,
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

func (p Gs2MissionRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetNamespace(
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
	client Gs2MissionRestClient,
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

func (p Gs2MissionRestClient) UpdateNamespaceAsync(
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
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) UpdateNamespace(
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
	client Gs2MissionRestClient,
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

func (p Gs2MissionRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DeleteNamespace(
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
	client Gs2MissionRestClient,
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

func (p Gs2MissionRestClient) GetServiceVersionAsync(
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
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetServiceVersion(
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

func dumpUserDataByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DumpUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DumpUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DumpUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DumpUserDataByUserIdAsync(
	request *DumpUserDataByUserIdRequest,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	path := "/system/dump/user/{userId}"
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

	go dumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DumpUserDataByUserId(
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

func checkDumpUserDataByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CheckDumpUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckDumpUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CheckDumpUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) CheckDumpUserDataByUserIdAsync(
	request *CheckDumpUserDataByUserIdRequest,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	path := "/system/dump/user/{userId}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
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

	go checkDumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) CheckDumpUserDataByUserId(
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

func cleanUserDataByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CleanUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CleanUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CleanUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) CleanUserDataByUserIdAsync(
	request *CleanUserDataByUserIdRequest,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	path := "/system/clean/user/{userId}"
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

	go cleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) CleanUserDataByUserId(
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

func checkCleanUserDataByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CheckCleanUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckCleanUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CheckCleanUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) CheckCleanUserDataByUserIdAsync(
	request *CheckCleanUserDataByUserIdRequest,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	path := "/system/clean/user/{userId}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
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

	go checkCleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) CheckCleanUserDataByUserId(
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

func prepareImportUserDataByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- PrepareImportUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareImportUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareImportUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) PrepareImportUserDataByUserIdAsync(
	request *PrepareImportUserDataByUserIdRequest,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	path := "/system/import/user/{userId}/prepare"
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

	go prepareImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) PrepareImportUserDataByUserId(
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

func importUserDataByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- ImportUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ImportUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ImportUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) ImportUserDataByUserIdAsync(
	request *ImportUserDataByUserIdRequest,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	path := "/system/import/user/{userId}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
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

	go importUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) ImportUserDataByUserId(
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

func checkImportUserDataByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CheckImportUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckImportUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CheckImportUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) CheckImportUserDataByUserIdAsync(
	request *CheckImportUserDataByUserIdRequest,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	path := "/system/import/user/{userId}/{uploadToken}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		path = strings.ReplaceAll(path, "{uploadToken}", core.ToString(*request.UploadToken))
	} else {
		path = strings.ReplaceAll(path, "{uploadToken}", "null")
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

	go checkImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) CheckImportUserDataByUserId(
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

func describeCountersAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCountersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeCountersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCountersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCountersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DescribeCountersAsync(
	request *DescribeCountersRequest,
	callback chan<- DescribeCountersAsyncResult,
) {
	path := "/{namespaceName}/user/me/counter"
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

	go describeCountersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DescribeCounters(
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

func describeCountersByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCountersByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeCountersByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCountersByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCountersByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DescribeCountersByUserIdAsync(
	request *DescribeCountersByUserIdRequest,
	callback chan<- DescribeCountersByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter"
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

	go describeCountersByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DescribeCountersByUserId(
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

func increaseCounterByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- IncreaseCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "counter.increase.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
		}
		callback <- IncreaseCounterByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IncreaseCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- IncreaseCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) IncreaseCounterByUserIdAsync(
	request *IncreaseCounterByUserIdRequest,
	callback chan<- IncreaseCounterByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
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

	go increaseCounterByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) IncreaseCounterByUserId(
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

func setCounterByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- SetCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetCounterByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetCounterByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "counter.increase.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
		}
		callback <- SetCounterByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) SetCounterByUserIdAsync(
	request *SetCounterByUserIdRequest,
	callback chan<- SetCounterByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Values != nil {
		var _values []interface{}
		for _, item := range request.Values {
			_values = append(_values, item)
		}
		bodies["values"] = _values
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

	go setCounterByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) SetCounterByUserId(
	request *SetCounterByUserIdRequest,
) (*SetCounterByUserIdResult, error) {
	callback := make(chan SetCounterByUserIdAsyncResult, 1)
	go p.SetCounterByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func decreaseCounterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DecreaseCounterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DecreaseCounterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DecreaseCounterResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "counter.increase.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
		}
		callback <- DecreaseCounterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseCounterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DecreaseCounterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DecreaseCounterAsync(
	request *DecreaseCounterRequest,
	callback chan<- DecreaseCounterAsyncResult,
) {
	path := "/{namespaceName}/user/me/counter/{counterName}/decrease"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
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

	go decreaseCounterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DecreaseCounter(
	request *DecreaseCounterRequest,
) (*DecreaseCounterResult, error) {
	callback := make(chan DecreaseCounterAsyncResult, 1)
	go p.DecreaseCounterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func decreaseCounterByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DecreaseCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "counter.increase.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
		}
		callback <- DecreaseCounterByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DecreaseCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DecreaseCounterByUserIdAsync(
	request *DecreaseCounterByUserIdRequest,
	callback chan<- DecreaseCounterByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter/{counterName}/decrease"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
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

	go decreaseCounterByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DecreaseCounterByUserId(
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

func getCounterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- GetCounterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetCounterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCounterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCounterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) GetCounterAsync(
	request *GetCounterRequest,
	callback chan<- GetCounterAsyncResult,
) {
	path := "/{namespaceName}/user/me/counter/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
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

	go getCounterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetCounter(
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

func getCounterByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- GetCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetCounterByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) GetCounterByUserIdAsync(
	request *GetCounterByUserIdRequest,
	callback chan<- GetCounterByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
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

	go getCounterByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetCounterByUserId(
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

func verifyCounterValueAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyCounterValueAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCounterValueAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCounterValueResult
	if asyncResult.Err != nil {
		callback <- VerifyCounterValueAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCounterValueAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyCounterValueAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) VerifyCounterValueAsync(
	request *VerifyCounterValueRequest,
	callback chan<- VerifyCounterValueAsyncResult,
) {
	path := "/{namespaceName}/user/me/counter/{counterName}/verify/counter/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ScopeType != nil && *request.ScopeType != "" {
		bodies["scopeType"] = *request.ScopeType
	}
	if request.ResetType != nil && *request.ResetType != "" {
		bodies["resetType"] = *request.ResetType
	}
	if request.ConditionName != nil && *request.ConditionName != "" {
		bodies["conditionName"] = *request.ConditionName
	}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go verifyCounterValueAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) VerifyCounterValue(
	request *VerifyCounterValueRequest,
) (*VerifyCounterValueResult, error) {
	callback := make(chan VerifyCounterValueAsyncResult, 1)
	go p.VerifyCounterValueAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyCounterValueByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyCounterValueByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCounterValueByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCounterValueByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyCounterValueByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCounterValueByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyCounterValueByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) VerifyCounterValueByUserIdAsync(
	request *VerifyCounterValueByUserIdRequest,
	callback chan<- VerifyCounterValueByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter/{counterName}/verify/counter/{verifyType}"
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
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ScopeType != nil && *request.ScopeType != "" {
		bodies["scopeType"] = *request.ScopeType
	}
	if request.ResetType != nil && *request.ResetType != "" {
		bodies["resetType"] = *request.ResetType
	}
	if request.ConditionName != nil && *request.ConditionName != "" {
		bodies["conditionName"] = *request.ConditionName
	}
	if request.Value != nil {
		bodies["value"] = *request.Value
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

	go verifyCounterValueByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) VerifyCounterValueByUserId(
	request *VerifyCounterValueByUserIdRequest,
) (*VerifyCounterValueByUserIdResult, error) {
	callback := make(chan VerifyCounterValueByUserIdAsyncResult, 1)
	go p.VerifyCounterValueByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func resetCounterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- ResetCounterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetCounterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetCounterResult
	if asyncResult.Err != nil {
		callback <- ResetCounterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetCounterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ResetCounterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) ResetCounterAsync(
	request *ResetCounterRequest,
	callback chan<- ResetCounterAsyncResult,
) {
	path := "/{namespaceName}/user/me/counter/{counterName}/reset"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Scopes != nil {
		var _scopes []interface{}
		for _, item := range request.Scopes {
			_scopes = append(_scopes, item)
		}
		bodies["scopes"] = _scopes
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

	go resetCounterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) ResetCounter(
	request *ResetCounterRequest,
) (*ResetCounterResult, error) {
	callback := make(chan ResetCounterAsyncResult, 1)
	go p.ResetCounterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func resetCounterByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- ResetCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetCounterByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetCounterByUserIdResult
	if asyncResult.Err != nil {
		callback <- ResetCounterByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ResetCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) ResetCounterByUserIdAsync(
	request *ResetCounterByUserIdRequest,
	callback chan<- ResetCounterByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter/{counterName}/reset"
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
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Scopes != nil {
		var _scopes []interface{}
		for _, item := range request.Scopes {
			_scopes = append(_scopes, item)
		}
		bodies["scopes"] = _scopes
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

	go resetCounterByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) ResetCounterByUserId(
	request *ResetCounterByUserIdRequest,
) (*ResetCounterByUserIdResult, error) {
	callback := make(chan ResetCounterByUserIdAsyncResult, 1)
	go p.ResetCounterByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteCounterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteCounterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteCounterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteCounterResult
	if asyncResult.Err != nil {
		callback <- DeleteCounterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteCounterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteCounterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DeleteCounterAsync(
	request *DeleteCounterRequest,
	callback chan<- DeleteCounterAsyncResult,
) {
	path := "/{namespaceName}/user/me/counter/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
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

	go deleteCounterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DeleteCounter(
	request *DeleteCounterRequest,
) (*DeleteCounterResult, error) {
	callback := make(chan DeleteCounterAsyncResult, 1)
	go p.DeleteCounterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteCounterByUserIdAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DeleteCounterByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DeleteCounterByUserIdAsync(
	request *DeleteCounterByUserIdRequest,
	callback chan<- DeleteCounterByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter/{counterName}"
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
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go deleteCounterByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DeleteCounterByUserId(
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

func increaseByStampSheetAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- IncreaseByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- IncreaseByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IncreaseByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- IncreaseByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) IncreaseByStampSheetAsync(
	request *IncreaseByStampSheetRequest,
	callback chan<- IncreaseByStampSheetAsyncResult,
) {
	path := "/stamp/increase"

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

	go increaseByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) IncreaseByStampSheet(
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

func setByStampSheetAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- SetByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetByStampSheetResult
	if asyncResult.Err != nil {
		callback <- SetByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) SetByStampSheetAsync(
	request *SetByStampSheetRequest,
	callback chan<- SetByStampSheetAsyncResult,
) {
	path := "/stamp/set"

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

	go setByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) SetByStampSheet(
	request *SetByStampSheetRequest,
) (*SetByStampSheetResult, error) {
	callback := make(chan SetByStampSheetAsyncResult, 1)
	go p.SetByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func decreaseByStampTaskAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DecreaseByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DecreaseByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DecreaseByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DecreaseByStampTaskAsync(
	request *DecreaseByStampTaskRequest,
	callback chan<- DecreaseByStampTaskAsyncResult,
) {
	path := "/stamp/decrease"

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

	go decreaseByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DecreaseByStampTask(
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

func resetByStampTaskAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- ResetByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetByStampTaskResult
	if asyncResult.Err != nil {
		callback <- ResetByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ResetByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) ResetByStampTaskAsync(
	request *ResetByStampTaskRequest,
	callback chan<- ResetByStampTaskAsyncResult,
) {
	path := "/stamp/reset"

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

	go resetByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) ResetByStampTask(
	request *ResetByStampTaskRequest,
) (*ResetByStampTaskResult, error) {
	callback := make(chan ResetByStampTaskAsyncResult, 1)
	go p.ResetByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyCounterValueByStampTaskAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyCounterValueByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCounterValueByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCounterValueByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyCounterValueByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCounterValueByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyCounterValueByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) VerifyCounterValueByStampTaskAsync(
	request *VerifyCounterValueByStampTaskRequest,
	callback chan<- VerifyCounterValueByStampTaskAsyncResult,
) {
	path := "/stamp/counter/verify"

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

	go verifyCounterValueByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) VerifyCounterValueByStampTask(
	request *VerifyCounterValueByStampTaskRequest,
) (*VerifyCounterValueByStampTaskResult, error) {
	callback := make(chan VerifyCounterValueByStampTaskAsyncResult, 1)
	go p.VerifyCounterValueByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2MissionRestClient,
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

func (p Gs2MissionRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) ExportMaster(
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

func getCurrentMissionMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentMissionMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetCurrentMissionMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentMissionMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentMissionMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) GetCurrentMissionMasterAsync(
	request *GetCurrentMissionMasterRequest,
	callback chan<- GetCurrentMissionMasterAsyncResult,
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

	go getCurrentMissionMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetCurrentMissionMaster(
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

func preUpdateCurrentMissionMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- PreUpdateCurrentMissionMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreUpdateCurrentMissionMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreUpdateCurrentMissionMasterResult
	if asyncResult.Err != nil {
		callback <- PreUpdateCurrentMissionMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreUpdateCurrentMissionMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PreUpdateCurrentMissionMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) PreUpdateCurrentMissionMasterAsync(
	request *PreUpdateCurrentMissionMasterRequest,
	callback chan<- PreUpdateCurrentMissionMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
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

	go preUpdateCurrentMissionMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) PreUpdateCurrentMissionMaster(
	request *PreUpdateCurrentMissionMasterRequest,
) (*PreUpdateCurrentMissionMasterResult, error) {
	callback := make(chan PreUpdateCurrentMissionMasterAsyncResult, 1)
	go p.PreUpdateCurrentMissionMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentMissionMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentMissionMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- UpdateCurrentMissionMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentMissionMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentMissionMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) UpdateCurrentMissionMasterAsync(
	request *UpdateCurrentMissionMasterRequest,
	callback chan<- UpdateCurrentMissionMasterAsyncResult,
) {
	if request.Settings != nil {
		res, err := p.PreUpdateCurrentMissionMaster(
			&PreUpdateCurrentMissionMasterRequest{
				ContextStack:  request.ContextStack,
				NamespaceName: request.NamespaceName,
			},
		)
		if err != nil {
			callback <- UpdateCurrentMissionMasterAsyncResult{
				err: err,
			}
			return
		}
		{
			req, _ := http.NewRequest("PUT", *res.UploadUrl, strings.NewReader(*request.Settings))
			req.Header.Set("Content-Type", "application/json")

			client := new(http.Client)
			_, err = client.Do(req)
			if err != nil {
				callback <- UpdateCurrentMissionMasterAsyncResult{
					err: err,
				}
				return
			}
		}
		v := "preUpload"
		request.Mode = &v
		request.UploadToken = res.UploadToken
		request.Settings = nil
	}
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Settings != nil && *request.Settings != "" {
		bodies["settings"] = *request.Settings
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
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

	go updateCurrentMissionMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) UpdateCurrentMissionMaster(
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

func updateCurrentMissionMasterFromGitHubAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentMissionMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- UpdateCurrentMissionMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentMissionMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentMissionMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) UpdateCurrentMissionMasterFromGitHubAsync(
	request *UpdateCurrentMissionMasterFromGitHubRequest,
	callback chan<- UpdateCurrentMissionMasterFromGitHubAsyncResult,
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

	go updateCurrentMissionMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) UpdateCurrentMissionMasterFromGitHub(
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

func describeCounterModelsAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCounterModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeCounterModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCounterModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCounterModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DescribeCounterModelsAsync(
	request *DescribeCounterModelsRequest,
	callback chan<- DescribeCounterModelsAsyncResult,
) {
	path := "/{namespaceName}/counter"
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

	go describeCounterModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DescribeCounterModels(
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

func getCounterModelAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- GetCounterModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetCounterModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCounterModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCounterModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) GetCounterModelAsync(
	request *GetCounterModelRequest,
	callback chan<- GetCounterModelAsyncResult,
) {
	path := "/{namespaceName}/counter/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
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

	go getCounterModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetCounterModel(
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

func describeMissionGroupModelsAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMissionGroupModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeMissionGroupModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionGroupModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMissionGroupModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DescribeMissionGroupModelsAsync(
	request *DescribeMissionGroupModelsRequest,
	callback chan<- DescribeMissionGroupModelsAsyncResult,
) {
	path := "/{namespaceName}/group"
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

	go describeMissionGroupModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DescribeMissionGroupModels(
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

func getMissionGroupModelAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- GetMissionGroupModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetMissionGroupModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMissionGroupModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMissionGroupModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) GetMissionGroupModelAsync(
	request *GetMissionGroupModelRequest,
	callback chan<- GetMissionGroupModelAsyncResult,
) {
	path := "/{namespaceName}/group/{missionGroupName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
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

	go getMissionGroupModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetMissionGroupModel(
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

func describeMissionTaskModelsAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMissionTaskModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeMissionTaskModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionTaskModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMissionTaskModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DescribeMissionTaskModelsAsync(
	request *DescribeMissionTaskModelsRequest,
	callback chan<- DescribeMissionTaskModelsAsyncResult,
) {
	path := "/{namespaceName}/group/{missionGroupName}/task"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
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

	go describeMissionTaskModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DescribeMissionTaskModels(
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

func getMissionTaskModelAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- GetMissionTaskModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetMissionTaskModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMissionTaskModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMissionTaskModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) GetMissionTaskModelAsync(
	request *GetMissionTaskModelRequest,
	callback chan<- GetMissionTaskModelAsyncResult,
) {
	path := "/{namespaceName}/group/{missionGroupName}/task/{missionTaskName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		path = strings.ReplaceAll(path, "{missionTaskName}", core.ToString(*request.MissionTaskName))
	} else {
		path = strings.ReplaceAll(path, "{missionTaskName}", "null")
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

	go getMissionTaskModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetMissionTaskModel(
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

func describeMissionTaskModelMastersAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMissionTaskModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeMissionTaskModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionTaskModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMissionTaskModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DescribeMissionTaskModelMastersAsync(
	request *DescribeMissionTaskModelMastersRequest,
	callback chan<- DescribeMissionTaskModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/group/{missionGroupName}/task"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
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

	go describeMissionTaskModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DescribeMissionTaskModelMasters(
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

func createMissionTaskModelMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- CreateMissionTaskModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- CreateMissionTaskModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateMissionTaskModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateMissionTaskModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) CreateMissionTaskModelMasterAsync(
	request *CreateMissionTaskModelMasterRequest,
	callback chan<- CreateMissionTaskModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{missionGroupName}/task"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.VerifyCompleteType != nil && *request.VerifyCompleteType != "" {
		bodies["verifyCompleteType"] = *request.VerifyCompleteType
	}
	if request.TargetCounter != nil {
		bodies["targetCounter"] = request.TargetCounter.ToDict()
	}
	if request.VerifyCompleteConsumeActions != nil {
		var _verifyCompleteConsumeActions []interface{}
		for _, item := range request.VerifyCompleteConsumeActions {
			_verifyCompleteConsumeActions = append(_verifyCompleteConsumeActions, item)
		}
		bodies["verifyCompleteConsumeActions"] = _verifyCompleteConsumeActions
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
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
	}
	if request.TargetResetType != nil && *request.TargetResetType != "" {
		bodies["targetResetType"] = *request.TargetResetType
	}
	if request.TargetValue != nil {
		bodies["targetValue"] = *request.TargetValue
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

	go createMissionTaskModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) CreateMissionTaskModelMaster(
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

func getMissionTaskModelMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- GetMissionTaskModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetMissionTaskModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMissionTaskModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMissionTaskModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) GetMissionTaskModelMasterAsync(
	request *GetMissionTaskModelMasterRequest,
	callback chan<- GetMissionTaskModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{missionGroupName}/task/{missionTaskName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		path = strings.ReplaceAll(path, "{missionTaskName}", core.ToString(*request.MissionTaskName))
	} else {
		path = strings.ReplaceAll(path, "{missionTaskName}", "null")
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

	go getMissionTaskModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) GetMissionTaskModelMaster(
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

func updateMissionTaskModelMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateMissionTaskModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- UpdateMissionTaskModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMissionTaskModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateMissionTaskModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) UpdateMissionTaskModelMasterAsync(
	request *UpdateMissionTaskModelMasterRequest,
	callback chan<- UpdateMissionTaskModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{missionGroupName}/task/{missionTaskName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		path = strings.ReplaceAll(path, "{missionTaskName}", core.ToString(*request.MissionTaskName))
	} else {
		path = strings.ReplaceAll(path, "{missionTaskName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.VerifyCompleteType != nil && *request.VerifyCompleteType != "" {
		bodies["verifyCompleteType"] = *request.VerifyCompleteType
	}
	if request.TargetCounter != nil {
		bodies["targetCounter"] = request.TargetCounter.ToDict()
	}
	if request.VerifyCompleteConsumeActions != nil {
		var _verifyCompleteConsumeActions []interface{}
		for _, item := range request.VerifyCompleteConsumeActions {
			_verifyCompleteConsumeActions = append(_verifyCompleteConsumeActions, item)
		}
		bodies["verifyCompleteConsumeActions"] = _verifyCompleteConsumeActions
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
	if request.CounterName != nil && *request.CounterName != "" {
		bodies["counterName"] = *request.CounterName
	}
	if request.TargetResetType != nil && *request.TargetResetType != "" {
		bodies["targetResetType"] = *request.TargetResetType
	}
	if request.TargetValue != nil {
		bodies["targetValue"] = *request.TargetValue
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

	go updateMissionTaskModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) UpdateMissionTaskModelMaster(
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

func deleteMissionTaskModelMasterAsyncHandler(
	client Gs2MissionRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteMissionTaskModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DeleteMissionTaskModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMissionTaskModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteMissionTaskModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MissionRestClient) DeleteMissionTaskModelMasterAsync(
	request *DeleteMissionTaskModelMasterRequest,
	callback chan<- DeleteMissionTaskModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/group/{missionGroupName}/task/{missionTaskName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}
	if request.MissionTaskName != nil && *request.MissionTaskName != "" {
		path = strings.ReplaceAll(path, "{missionTaskName}", core.ToString(*request.MissionTaskName))
	} else {
		path = strings.ReplaceAll(path, "{missionTaskName}", "null")
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

	go deleteMissionTaskModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mission").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MissionRestClient) DeleteMissionTaskModelMaster(
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
