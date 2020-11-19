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

type QueryAccessLogResult struct {
	/** アクセスログのリスト */
	Items *[]*AccessLog `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
	/** クエリ結果の総件数 */
	TotalCount *int64 `json:"totalCount"`
	/** 検索時にスキャンした総容量 */
	ScanSize *int64 `json:"scanSize"`
}

func (p *QueryAccessLogResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	if p.TotalCount != nil {
		data["TotalCount"] = p.TotalCount
	}
	if p.ScanSize != nil {
		data["ScanSize"] = p.ScanSize
	}
	return &data
}

type QueryAccessLogAsyncResult struct {
	result *QueryAccessLogResult
	err    error
}

type CountAccessLogResult struct {
	/** アクセスログ集計のリスト */
	Items *[]*AccessLogCount `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
	/** クエリ結果の総件数 */
	TotalCount *int64 `json:"totalCount"`
	/** 検索時にスキャンした総容量 */
	ScanSize *int64 `json:"scanSize"`
}

func (p *CountAccessLogResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	if p.TotalCount != nil {
		data["TotalCount"] = p.TotalCount
	}
	if p.ScanSize != nil {
		data["ScanSize"] = p.ScanSize
	}
	return &data
}

type CountAccessLogAsyncResult struct {
	result *CountAccessLogResult
	err    error
}

type QueryIssueStampSheetLogResult struct {
	/** スタンプシート発行ログのリスト */
	Items *[]*IssueStampSheetLog `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
	/** クエリ結果の総件数 */
	TotalCount *int64 `json:"totalCount"`
	/** 検索時にスキャンした総容量 */
	ScanSize *int64 `json:"scanSize"`
}

func (p *QueryIssueStampSheetLogResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	if p.TotalCount != nil {
		data["TotalCount"] = p.TotalCount
	}
	if p.ScanSize != nil {
		data["ScanSize"] = p.ScanSize
	}
	return &data
}

type QueryIssueStampSheetLogAsyncResult struct {
	result *QueryIssueStampSheetLogResult
	err    error
}

type CountIssueStampSheetLogResult struct {
	/** スタンプシート発行ログ集計のリスト */
	Items *[]*IssueStampSheetLogCount `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
	/** クエリ結果の総件数 */
	TotalCount *int64 `json:"totalCount"`
	/** 検索時にスキャンした総容量 */
	ScanSize *int64 `json:"scanSize"`
}

func (p *CountIssueStampSheetLogResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	if p.TotalCount != nil {
		data["TotalCount"] = p.TotalCount
	}
	if p.ScanSize != nil {
		data["ScanSize"] = p.ScanSize
	}
	return &data
}

type CountIssueStampSheetLogAsyncResult struct {
	result *CountIssueStampSheetLogResult
	err    error
}

type QueryExecuteStampSheetLogResult struct {
	/** スタンプシート実行ログのリスト */
	Items *[]*ExecuteStampSheetLog `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
	/** クエリ結果の総件数 */
	TotalCount *int64 `json:"totalCount"`
	/** 検索時にスキャンした総容量 */
	ScanSize *int64 `json:"scanSize"`
}

func (p *QueryExecuteStampSheetLogResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	if p.TotalCount != nil {
		data["TotalCount"] = p.TotalCount
	}
	if p.ScanSize != nil {
		data["ScanSize"] = p.ScanSize
	}
	return &data
}

type QueryExecuteStampSheetLogAsyncResult struct {
	result *QueryExecuteStampSheetLogResult
	err    error
}

type CountExecuteStampSheetLogResult struct {
	/** スタンプシート実行ログ集計のリスト */
	Items *[]*ExecuteStampSheetLogCount `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
	/** クエリ結果の総件数 */
	TotalCount *int64 `json:"totalCount"`
	/** 検索時にスキャンした総容量 */
	ScanSize *int64 `json:"scanSize"`
}

func (p *CountExecuteStampSheetLogResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	if p.TotalCount != nil {
		data["TotalCount"] = p.TotalCount
	}
	if p.ScanSize != nil {
		data["ScanSize"] = p.ScanSize
	}
	return &data
}

type CountExecuteStampSheetLogAsyncResult struct {
	result *CountExecuteStampSheetLogResult
	err    error
}

type QueryExecuteStampTaskLogResult struct {
	/** スタンプタスク実行ログのリスト */
	Items *[]*ExecuteStampTaskLog `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
	/** クエリ結果の総件数 */
	TotalCount *int64 `json:"totalCount"`
	/** 検索時にスキャンした総容量 */
	ScanSize *int64 `json:"scanSize"`
}

func (p *QueryExecuteStampTaskLogResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	if p.TotalCount != nil {
		data["TotalCount"] = p.TotalCount
	}
	if p.ScanSize != nil {
		data["ScanSize"] = p.ScanSize
	}
	return &data
}

type QueryExecuteStampTaskLogAsyncResult struct {
	result *QueryExecuteStampTaskLogResult
	err    error
}

type CountExecuteStampTaskLogResult struct {
	/** スタンプタスク実行ログ集計のリスト */
	Items *[]*ExecuteStampTaskLogCount `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
	/** クエリ結果の総件数 */
	TotalCount *int64 `json:"totalCount"`
	/** 検索時にスキャンした総容量 */
	ScanSize *int64 `json:"scanSize"`
}

func (p *CountExecuteStampTaskLogResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	if p.TotalCount != nil {
		data["TotalCount"] = p.TotalCount
	}
	if p.ScanSize != nil {
		data["ScanSize"] = p.ScanSize
	}
	return &data
}

type CountExecuteStampTaskLogAsyncResult struct {
	result *CountExecuteStampTaskLogResult
	err    error
}
