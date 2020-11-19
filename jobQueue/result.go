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

package jobQueue

type DescribeNamespacesResult struct {
	/** ネームスペースのリスト */
	Items *[]*Namespace `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeNamespacesResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

type CreateNamespaceResult struct {
	/** 作成したネームスペース */
	Item *Namespace `json:"item"`
}

func (p *CreateNamespaceResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

type GetNamespaceStatusResult struct {
	/** None */
	Status *string `json:"status"`
}

func (p *GetNamespaceStatusResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Status != nil {
		data["Status"] = p.Status
	}
	return &data
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

type GetNamespaceResult struct {
	/** ネームスペース */
	Item *Namespace `json:"item"`
}

func (p *GetNamespaceResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

type UpdateNamespaceResult struct {
	/** 更新したネームスペース */
	Item *Namespace `json:"item"`
}

func (p *UpdateNamespaceResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

type DeleteNamespaceResult struct {
	/** 削除したネームスペース */
	Item *Namespace `json:"item"`
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type DescribeJobsByUserIdResult struct {
	/** ジョブのリスト */
	Items *[]*Job `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeJobsByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeJobsByUserIdAsyncResult struct {
	result *DescribeJobsByUserIdResult
	err    error
}

type GetJobByUserIdResult struct {
	/** ジョブ */
	Item *Job `json:"item"`
}

func (p *GetJobByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetJobByUserIdAsyncResult struct {
	result *GetJobByUserIdResult
	err    error
}

type PushByUserIdResult struct {
	/** 追加した{model_name}の一覧 */
	Items *[]*Job `json:"items"`
}

func (p *PushByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	return &data
}

type PushByUserIdAsyncResult struct {
	result *PushByUserIdResult
	err    error
}

type RunResult struct {
	/** ジョブ */
	Item *Job `json:"item"`
	/** ジョブの実行結果 */
	Result *JobResultBody `json:"result"`
	/** None */
	IsLastJob *bool `json:"isLastJob"`
}

func (p *RunResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.Result != nil {
		data["Result"] = p.Result.ToDict()
	}
	if p.IsLastJob != nil {
		data["IsLastJob"] = p.IsLastJob
	}
	return &data
}

type RunAsyncResult struct {
	result *RunResult
	err    error
}

type RunByUserIdResult struct {
	/** ジョブ */
	Item *Job `json:"item"`
	/** ジョブの実行結果 */
	Result *JobResultBody `json:"result"`
	/** None */
	IsLastJob *bool `json:"isLastJob"`
}

func (p *RunByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.Result != nil {
		data["Result"] = p.Result.ToDict()
	}
	if p.IsLastJob != nil {
		data["IsLastJob"] = p.IsLastJob
	}
	return &data
}

type RunByUserIdAsyncResult struct {
	result *RunByUserIdResult
	err    error
}

type DeleteJobByUserIdResult struct {
	/** ジョブ */
	Item *Job `json:"item"`
}

func (p *DeleteJobByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type DeleteJobByUserIdAsyncResult struct {
	result *DeleteJobByUserIdResult
	err    error
}

type PushByStampSheetResult struct {
	/** 追加した{model_name}の一覧 */
	Items *[]*Job `json:"items"`
}

func (p *PushByStampSheetResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	return &data
}

type PushByStampSheetAsyncResult struct {
	result *PushByStampSheetResult
	err    error
}

type DescribeDeadLetterJobsByUserIdResult struct {
	/** デッドレタージョブのリスト */
	Items *[]*DeadLetterJob `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeDeadLetterJobsByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeDeadLetterJobsByUserIdAsyncResult struct {
	result *DescribeDeadLetterJobsByUserIdResult
	err    error
}

type GetDeadLetterJobByUserIdResult struct {
	/** デッドレタージョブ */
	Item *DeadLetterJob `json:"item"`
}

func (p *GetDeadLetterJobByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetDeadLetterJobByUserIdAsyncResult struct {
	result *GetDeadLetterJobByUserIdResult
	err    error
}

type DeleteDeadLetterJobByUserIdResult struct {
	/** デッドレタージョブ */
	Item *DeadLetterJob `json:"item"`
}

func (p *DeleteDeadLetterJobByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type DeleteDeadLetterJobByUserIdAsyncResult struct {
	result *DeleteDeadLetterJobByUserIdResult
	err    error
}
