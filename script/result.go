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

package script

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
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type DescribeScriptsResult struct {
	/** スクリプトのリスト */
	Items *[]*Script `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeScriptsResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeScriptsAsyncResult struct {
	result *DescribeScriptsResult
	err    error
}

type CreateScriptResult struct {
	/** 作成したスクリプト */
	Item *Script `json:"item"`
}

func (p *CreateScriptResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type CreateScriptAsyncResult struct {
	result *CreateScriptResult
	err    error
}

type CreateScriptFromGitHubResult struct {
	/** 作成したスクリプト */
	Item *Script `json:"item"`
}

func (p *CreateScriptFromGitHubResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type CreateScriptFromGitHubAsyncResult struct {
	result *CreateScriptFromGitHubResult
	err    error
}

type GetScriptResult struct {
	/** スクリプト */
	Item *Script `json:"item"`
}

func (p *GetScriptResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetScriptAsyncResult struct {
	result *GetScriptResult
	err    error
}

type UpdateScriptResult struct {
	/** 更新したスクリプト */
	Item *Script `json:"item"`
}

func (p *UpdateScriptResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UpdateScriptAsyncResult struct {
	result *UpdateScriptResult
	err    error
}

type UpdateScriptFromGitHubResult struct {
	/** 更新したスクリプト */
	Item *Script `json:"item"`
}

func (p *UpdateScriptFromGitHubResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UpdateScriptFromGitHubAsyncResult struct {
	result *UpdateScriptFromGitHubResult
	err    error
}

type DeleteScriptResult struct {
}

func (p *DeleteScriptResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	return &data
}

type DeleteScriptAsyncResult struct {
	result *DeleteScriptResult
	err    error
}

type InvokeScriptResult struct {
	/** ステータスコード */
	Code *int32 `json:"code"`
	/** 戻り値 */
	Result *string `json:"result"`
	/** スクリプトの実行時間(ミリ秒) */
	ExecuteTime *int32 `json:"executeTime"`
	/** 費用の計算対象となった時間(秒) */
	Charged *int32 `json:"charged"`
	/** 標準出力の内容のリスト */
	Output *[]string `json:"output"`
}

func (p *InvokeScriptResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Code != nil {
		data["Code"] = p.Code
	}
	if p.Result != nil {
		data["Result"] = p.Result
	}
	if p.ExecuteTime != nil {
		data["ExecuteTime"] = p.ExecuteTime
	}
	if p.Charged != nil {
		data["Charged"] = p.Charged
	}
	if p.Output != nil {
	}
	return &data
}

type InvokeScriptAsyncResult struct {
	result *InvokeScriptResult
	err    error
}

type DebugInvokeResult struct {
	/** ステータスコード */
	Code *int32 `json:"code"`
	/** 戻り値 */
	Result *string `json:"result"`
	/** スクリプトの実行時間(ミリ秒) */
	ExecuteTime *int32 `json:"executeTime"`
	/** 費用の計算対象となった時間(秒) */
	Charged *int32 `json:"charged"`
	/** 標準出力の内容のリスト */
	Output *[]string `json:"output"`
}

func (p *DebugInvokeResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Code != nil {
		data["Code"] = p.Code
	}
	if p.Result != nil {
		data["Result"] = p.Result
	}
	if p.ExecuteTime != nil {
		data["ExecuteTime"] = p.ExecuteTime
	}
	if p.Charged != nil {
		data["Charged"] = p.Charged
	}
	if p.Output != nil {
	}
	return &data
}

type DebugInvokeAsyncResult struct {
	result *DebugInvokeResult
	err    error
}
