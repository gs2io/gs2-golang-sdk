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

import "core"

type DescribeNamespacesResult struct {
    Items []Namespace `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

func NewDescribeNamespacesResultFromDict(data map[string]interface{}) DescribeNamespacesResult {
    return DescribeNamespacesResult {
        Items: CastNamespaces(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastNamespacesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeNamespacesResult) Pointer() *DescribeNamespacesResult {
    return &p
}

type CreateNamespaceResult struct {
    Item *Namespace `json:"item"`
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

func NewCreateNamespaceResultFromDict(data map[string]interface{}) CreateNamespaceResult {
    return CreateNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateNamespaceResult) Pointer() *CreateNamespaceResult {
    return &p
}

type GetNamespaceStatusResult struct {
    Status *string `json:"status"`
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

func NewGetNamespaceStatusResultFromDict(data map[string]interface{}) GetNamespaceStatusResult {
    return GetNamespaceStatusResult {
        Status: core.CastString(data["status"]),
    }
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "status": p.Status,
    }
}

func (p GetNamespaceStatusResult) Pointer() *GetNamespaceStatusResult {
    return &p
}

type GetNamespaceResult struct {
    Item *Namespace `json:"item"`
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

func NewGetNamespaceResultFromDict(data map[string]interface{}) GetNamespaceResult {
    return GetNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetNamespaceResult) Pointer() *GetNamespaceResult {
    return &p
}

type UpdateNamespaceResult struct {
    Item *Namespace `json:"item"`
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

func NewUpdateNamespaceResultFromDict(data map[string]interface{}) UpdateNamespaceResult {
    return UpdateNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateNamespaceResult) Pointer() *UpdateNamespaceResult {
    return &p
}

type DeleteNamespaceResult struct {
    Item *Namespace `json:"item"`
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

func NewDeleteNamespaceResultFromDict(data map[string]interface{}) DeleteNamespaceResult {
    return DeleteNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
    return &p
}

type QueryAccessLogResult struct {
    Items []AccessLog `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
    TotalCount *int64 `json:"totalCount"`
    ScanSize *int64 `json:"scanSize"`
}

type QueryAccessLogAsyncResult struct {
	result *QueryAccessLogResult
	err    error
}

func NewQueryAccessLogResultFromDict(data map[string]interface{}) QueryAccessLogResult {
    return QueryAccessLogResult {
        Items: CastAccessLogs(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
        TotalCount: core.CastInt64(data["totalCount"]),
        ScanSize: core.CastInt64(data["scanSize"]),
    }
}

func (p QueryAccessLogResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastAccessLogsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
        "totalCount": p.TotalCount,
        "scanSize": p.ScanSize,
    }
}

func (p QueryAccessLogResult) Pointer() *QueryAccessLogResult {
    return &p
}

type CountAccessLogResult struct {
    Items []AccessLogCount `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
    TotalCount *int64 `json:"totalCount"`
    ScanSize *int64 `json:"scanSize"`
}

type CountAccessLogAsyncResult struct {
	result *CountAccessLogResult
	err    error
}

func NewCountAccessLogResultFromDict(data map[string]interface{}) CountAccessLogResult {
    return CountAccessLogResult {
        Items: CastAccessLogCounts(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
        TotalCount: core.CastInt64(data["totalCount"]),
        ScanSize: core.CastInt64(data["scanSize"]),
    }
}

func (p CountAccessLogResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastAccessLogCountsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
        "totalCount": p.TotalCount,
        "scanSize": p.ScanSize,
    }
}

func (p CountAccessLogResult) Pointer() *CountAccessLogResult {
    return &p
}

type QueryIssueStampSheetLogResult struct {
    Items []IssueStampSheetLog `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
    TotalCount *int64 `json:"totalCount"`
    ScanSize *int64 `json:"scanSize"`
}

type QueryIssueStampSheetLogAsyncResult struct {
	result *QueryIssueStampSheetLogResult
	err    error
}

func NewQueryIssueStampSheetLogResultFromDict(data map[string]interface{}) QueryIssueStampSheetLogResult {
    return QueryIssueStampSheetLogResult {
        Items: CastIssueStampSheetLogs(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
        TotalCount: core.CastInt64(data["totalCount"]),
        ScanSize: core.CastInt64(data["scanSize"]),
    }
}

func (p QueryIssueStampSheetLogResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastIssueStampSheetLogsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
        "totalCount": p.TotalCount,
        "scanSize": p.ScanSize,
    }
}

func (p QueryIssueStampSheetLogResult) Pointer() *QueryIssueStampSheetLogResult {
    return &p
}

type CountIssueStampSheetLogResult struct {
    Items []IssueStampSheetLogCount `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
    TotalCount *int64 `json:"totalCount"`
    ScanSize *int64 `json:"scanSize"`
}

type CountIssueStampSheetLogAsyncResult struct {
	result *CountIssueStampSheetLogResult
	err    error
}

func NewCountIssueStampSheetLogResultFromDict(data map[string]interface{}) CountIssueStampSheetLogResult {
    return CountIssueStampSheetLogResult {
        Items: CastIssueStampSheetLogCounts(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
        TotalCount: core.CastInt64(data["totalCount"]),
        ScanSize: core.CastInt64(data["scanSize"]),
    }
}

func (p CountIssueStampSheetLogResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastIssueStampSheetLogCountsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
        "totalCount": p.TotalCount,
        "scanSize": p.ScanSize,
    }
}

func (p CountIssueStampSheetLogResult) Pointer() *CountIssueStampSheetLogResult {
    return &p
}

type QueryExecuteStampSheetLogResult struct {
    Items []ExecuteStampSheetLog `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
    TotalCount *int64 `json:"totalCount"`
    ScanSize *int64 `json:"scanSize"`
}

type QueryExecuteStampSheetLogAsyncResult struct {
	result *QueryExecuteStampSheetLogResult
	err    error
}

func NewQueryExecuteStampSheetLogResultFromDict(data map[string]interface{}) QueryExecuteStampSheetLogResult {
    return QueryExecuteStampSheetLogResult {
        Items: CastExecuteStampSheetLogs(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
        TotalCount: core.CastInt64(data["totalCount"]),
        ScanSize: core.CastInt64(data["scanSize"]),
    }
}

func (p QueryExecuteStampSheetLogResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastExecuteStampSheetLogsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
        "totalCount": p.TotalCount,
        "scanSize": p.ScanSize,
    }
}

func (p QueryExecuteStampSheetLogResult) Pointer() *QueryExecuteStampSheetLogResult {
    return &p
}

type CountExecuteStampSheetLogResult struct {
    Items []ExecuteStampSheetLogCount `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
    TotalCount *int64 `json:"totalCount"`
    ScanSize *int64 `json:"scanSize"`
}

type CountExecuteStampSheetLogAsyncResult struct {
	result *CountExecuteStampSheetLogResult
	err    error
}

func NewCountExecuteStampSheetLogResultFromDict(data map[string]interface{}) CountExecuteStampSheetLogResult {
    return CountExecuteStampSheetLogResult {
        Items: CastExecuteStampSheetLogCounts(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
        TotalCount: core.CastInt64(data["totalCount"]),
        ScanSize: core.CastInt64(data["scanSize"]),
    }
}

func (p CountExecuteStampSheetLogResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastExecuteStampSheetLogCountsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
        "totalCount": p.TotalCount,
        "scanSize": p.ScanSize,
    }
}

func (p CountExecuteStampSheetLogResult) Pointer() *CountExecuteStampSheetLogResult {
    return &p
}

type QueryExecuteStampTaskLogResult struct {
    Items []ExecuteStampTaskLog `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
    TotalCount *int64 `json:"totalCount"`
    ScanSize *int64 `json:"scanSize"`
}

type QueryExecuteStampTaskLogAsyncResult struct {
	result *QueryExecuteStampTaskLogResult
	err    error
}

func NewQueryExecuteStampTaskLogResultFromDict(data map[string]interface{}) QueryExecuteStampTaskLogResult {
    return QueryExecuteStampTaskLogResult {
        Items: CastExecuteStampTaskLogs(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
        TotalCount: core.CastInt64(data["totalCount"]),
        ScanSize: core.CastInt64(data["scanSize"]),
    }
}

func (p QueryExecuteStampTaskLogResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastExecuteStampTaskLogsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
        "totalCount": p.TotalCount,
        "scanSize": p.ScanSize,
    }
}

func (p QueryExecuteStampTaskLogResult) Pointer() *QueryExecuteStampTaskLogResult {
    return &p
}

type CountExecuteStampTaskLogResult struct {
    Items []ExecuteStampTaskLogCount `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
    TotalCount *int64 `json:"totalCount"`
    ScanSize *int64 `json:"scanSize"`
}

type CountExecuteStampTaskLogAsyncResult struct {
	result *CountExecuteStampTaskLogResult
	err    error
}

func NewCountExecuteStampTaskLogResultFromDict(data map[string]interface{}) CountExecuteStampTaskLogResult {
    return CountExecuteStampTaskLogResult {
        Items: CastExecuteStampTaskLogCounts(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
        TotalCount: core.CastInt64(data["totalCount"]),
        ScanSize: core.CastInt64(data["scanSize"]),
    }
}

func (p CountExecuteStampTaskLogResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastExecuteStampTaskLogCountsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
        "totalCount": p.TotalCount,
        "scanSize": p.ScanSize,
    }
}

func (p CountExecuteStampTaskLogResult) Pointer() *CountExecuteStampTaskLogResult {
    return &p
}