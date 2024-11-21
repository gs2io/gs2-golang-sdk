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
)

type DescribeNamespacesResult struct {
	Items         []Namespace          `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

func NewDescribeNamespacesResultFromJson(data string) DescribeNamespacesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNamespacesResultFromDict(dict)
}

func NewDescribeNamespacesResultFromDict(data map[string]interface{}) DescribeNamespacesResult {
	return DescribeNamespacesResult{
		Items: func() []Namespace {
			if data["items"] == nil {
				return nil
			}
			return CastNamespaces(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

func NewCreateNamespaceResultFromJson(data string) CreateNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceResultFromDict(dict)
}

func NewCreateNamespaceResultFromDict(data map[string]interface{}) CreateNamespaceResult {
	return CreateNamespaceResult{
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateNamespaceResult) Pointer() *CreateNamespaceResult {
	return &p
}

type GetNamespaceStatusResult struct {
	Status   *string              `json:"status"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

func NewGetNamespaceStatusResultFromJson(data string) GetNamespaceStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceStatusResultFromDict(dict)
}

func NewGetNamespaceStatusResultFromDict(data map[string]interface{}) GetNamespaceStatusResult {
	return GetNamespaceStatusResult{
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
		}(),
	}
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"status": p.Status,
	}
}

func (p GetNamespaceStatusResult) Pointer() *GetNamespaceStatusResult {
	return &p
}

type GetNamespaceResult struct {
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

func NewGetNamespaceResultFromJson(data string) GetNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceResultFromDict(dict)
}

func NewGetNamespaceResultFromDict(data map[string]interface{}) GetNamespaceResult {
	return GetNamespaceResult{
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetNamespaceResult) Pointer() *GetNamespaceResult {
	return &p
}

type UpdateNamespaceResult struct {
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

func NewUpdateNamespaceResultFromJson(data string) UpdateNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceResultFromDict(dict)
}

func NewUpdateNamespaceResultFromDict(data map[string]interface{}) UpdateNamespaceResult {
	return UpdateNamespaceResult{
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateNamespaceResult) Pointer() *UpdateNamespaceResult {
	return &p
}

type DeleteNamespaceResult struct {
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

func NewDeleteNamespaceResultFromJson(data string) DeleteNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteNamespaceResultFromDict(dict)
}

func NewDeleteNamespaceResultFromDict(data map[string]interface{}) DeleteNamespaceResult {
	return DeleteNamespaceResult{
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
	return &p
}

type QueryAccessLogResult struct {
	Items         []AccessLog          `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	TotalCount    *int64               `json:"totalCount"`
	ScanSize      *int64               `json:"scanSize"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type QueryAccessLogAsyncResult struct {
	result *QueryAccessLogResult
	err    error
}

func NewQueryAccessLogResultFromJson(data string) QueryAccessLogResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQueryAccessLogResultFromDict(dict)
}

func NewQueryAccessLogResultFromDict(data map[string]interface{}) QueryAccessLogResult {
	return QueryAccessLogResult{
		Items: func() []AccessLog {
			if data["items"] == nil {
				return nil
			}
			return CastAccessLogs(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		TotalCount: func() *int64 {
			v, ok := data["totalCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["totalCount"])
		}(),
		ScanSize: func() *int64 {
			v, ok := data["scanSize"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["scanSize"])
		}(),
	}
}

func (p QueryAccessLogResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAccessLogsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"totalCount":    p.TotalCount,
		"scanSize":      p.ScanSize,
	}
}

func (p QueryAccessLogResult) Pointer() *QueryAccessLogResult {
	return &p
}

type CountAccessLogResult struct {
	Items         []AccessLogCount     `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	TotalCount    *int64               `json:"totalCount"`
	ScanSize      *int64               `json:"scanSize"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type CountAccessLogAsyncResult struct {
	result *CountAccessLogResult
	err    error
}

func NewCountAccessLogResultFromJson(data string) CountAccessLogResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountAccessLogResultFromDict(dict)
}

func NewCountAccessLogResultFromDict(data map[string]interface{}) CountAccessLogResult {
	return CountAccessLogResult{
		Items: func() []AccessLogCount {
			if data["items"] == nil {
				return nil
			}
			return CastAccessLogCounts(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		TotalCount: func() *int64 {
			v, ok := data["totalCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["totalCount"])
		}(),
		ScanSize: func() *int64 {
			v, ok := data["scanSize"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["scanSize"])
		}(),
	}
}

func (p CountAccessLogResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAccessLogCountsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"totalCount":    p.TotalCount,
		"scanSize":      p.ScanSize,
	}
}

func (p CountAccessLogResult) Pointer() *CountAccessLogResult {
	return &p
}

type QueryIssueStampSheetLogResult struct {
	Items         []IssueStampSheetLog `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	TotalCount    *int64               `json:"totalCount"`
	ScanSize      *int64               `json:"scanSize"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type QueryIssueStampSheetLogAsyncResult struct {
	result *QueryIssueStampSheetLogResult
	err    error
}

func NewQueryIssueStampSheetLogResultFromJson(data string) QueryIssueStampSheetLogResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQueryIssueStampSheetLogResultFromDict(dict)
}

func NewQueryIssueStampSheetLogResultFromDict(data map[string]interface{}) QueryIssueStampSheetLogResult {
	return QueryIssueStampSheetLogResult{
		Items: func() []IssueStampSheetLog {
			if data["items"] == nil {
				return nil
			}
			return CastIssueStampSheetLogs(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		TotalCount: func() *int64 {
			v, ok := data["totalCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["totalCount"])
		}(),
		ScanSize: func() *int64 {
			v, ok := data["scanSize"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["scanSize"])
		}(),
	}
}

func (p QueryIssueStampSheetLogResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastIssueStampSheetLogsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"totalCount":    p.TotalCount,
		"scanSize":      p.ScanSize,
	}
}

func (p QueryIssueStampSheetLogResult) Pointer() *QueryIssueStampSheetLogResult {
	return &p
}

type CountIssueStampSheetLogResult struct {
	Items         []IssueStampSheetLogCount `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
	TotalCount    *int64                    `json:"totalCount"`
	ScanSize      *int64                    `json:"scanSize"`
	Metadata      *core.ResultMetadata      `json:"metadata"`
}

type CountIssueStampSheetLogAsyncResult struct {
	result *CountIssueStampSheetLogResult
	err    error
}

func NewCountIssueStampSheetLogResultFromJson(data string) CountIssueStampSheetLogResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountIssueStampSheetLogResultFromDict(dict)
}

func NewCountIssueStampSheetLogResultFromDict(data map[string]interface{}) CountIssueStampSheetLogResult {
	return CountIssueStampSheetLogResult{
		Items: func() []IssueStampSheetLogCount {
			if data["items"] == nil {
				return nil
			}
			return CastIssueStampSheetLogCounts(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		TotalCount: func() *int64 {
			v, ok := data["totalCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["totalCount"])
		}(),
		ScanSize: func() *int64 {
			v, ok := data["scanSize"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["scanSize"])
		}(),
	}
}

func (p CountIssueStampSheetLogResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastIssueStampSheetLogCountsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"totalCount":    p.TotalCount,
		"scanSize":      p.ScanSize,
	}
}

func (p CountIssueStampSheetLogResult) Pointer() *CountIssueStampSheetLogResult {
	return &p
}

type QueryExecuteStampSheetLogResult struct {
	Items         []ExecuteStampSheetLog `json:"items"`
	NextPageToken *string                `json:"nextPageToken"`
	TotalCount    *int64                 `json:"totalCount"`
	ScanSize      *int64                 `json:"scanSize"`
	Metadata      *core.ResultMetadata   `json:"metadata"`
}

type QueryExecuteStampSheetLogAsyncResult struct {
	result *QueryExecuteStampSheetLogResult
	err    error
}

func NewQueryExecuteStampSheetLogResultFromJson(data string) QueryExecuteStampSheetLogResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQueryExecuteStampSheetLogResultFromDict(dict)
}

func NewQueryExecuteStampSheetLogResultFromDict(data map[string]interface{}) QueryExecuteStampSheetLogResult {
	return QueryExecuteStampSheetLogResult{
		Items: func() []ExecuteStampSheetLog {
			if data["items"] == nil {
				return nil
			}
			return CastExecuteStampSheetLogs(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		TotalCount: func() *int64 {
			v, ok := data["totalCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["totalCount"])
		}(),
		ScanSize: func() *int64 {
			v, ok := data["scanSize"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["scanSize"])
		}(),
	}
}

func (p QueryExecuteStampSheetLogResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastExecuteStampSheetLogsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"totalCount":    p.TotalCount,
		"scanSize":      p.ScanSize,
	}
}

func (p QueryExecuteStampSheetLogResult) Pointer() *QueryExecuteStampSheetLogResult {
	return &p
}

type CountExecuteStampSheetLogResult struct {
	Items         []ExecuteStampSheetLogCount `json:"items"`
	NextPageToken *string                     `json:"nextPageToken"`
	TotalCount    *int64                      `json:"totalCount"`
	ScanSize      *int64                      `json:"scanSize"`
	Metadata      *core.ResultMetadata        `json:"metadata"`
}

type CountExecuteStampSheetLogAsyncResult struct {
	result *CountExecuteStampSheetLogResult
	err    error
}

func NewCountExecuteStampSheetLogResultFromJson(data string) CountExecuteStampSheetLogResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountExecuteStampSheetLogResultFromDict(dict)
}

func NewCountExecuteStampSheetLogResultFromDict(data map[string]interface{}) CountExecuteStampSheetLogResult {
	return CountExecuteStampSheetLogResult{
		Items: func() []ExecuteStampSheetLogCount {
			if data["items"] == nil {
				return nil
			}
			return CastExecuteStampSheetLogCounts(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		TotalCount: func() *int64 {
			v, ok := data["totalCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["totalCount"])
		}(),
		ScanSize: func() *int64 {
			v, ok := data["scanSize"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["scanSize"])
		}(),
	}
}

func (p CountExecuteStampSheetLogResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastExecuteStampSheetLogCountsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"totalCount":    p.TotalCount,
		"scanSize":      p.ScanSize,
	}
}

func (p CountExecuteStampSheetLogResult) Pointer() *CountExecuteStampSheetLogResult {
	return &p
}

type QueryExecuteStampTaskLogResult struct {
	Items         []ExecuteStampTaskLog `json:"items"`
	NextPageToken *string               `json:"nextPageToken"`
	TotalCount    *int64                `json:"totalCount"`
	ScanSize      *int64                `json:"scanSize"`
	Metadata      *core.ResultMetadata  `json:"metadata"`
}

type QueryExecuteStampTaskLogAsyncResult struct {
	result *QueryExecuteStampTaskLogResult
	err    error
}

func NewQueryExecuteStampTaskLogResultFromJson(data string) QueryExecuteStampTaskLogResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQueryExecuteStampTaskLogResultFromDict(dict)
}

func NewQueryExecuteStampTaskLogResultFromDict(data map[string]interface{}) QueryExecuteStampTaskLogResult {
	return QueryExecuteStampTaskLogResult{
		Items: func() []ExecuteStampTaskLog {
			if data["items"] == nil {
				return nil
			}
			return CastExecuteStampTaskLogs(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		TotalCount: func() *int64 {
			v, ok := data["totalCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["totalCount"])
		}(),
		ScanSize: func() *int64 {
			v, ok := data["scanSize"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["scanSize"])
		}(),
	}
}

func (p QueryExecuteStampTaskLogResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastExecuteStampTaskLogsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"totalCount":    p.TotalCount,
		"scanSize":      p.ScanSize,
	}
}

func (p QueryExecuteStampTaskLogResult) Pointer() *QueryExecuteStampTaskLogResult {
	return &p
}

type CountExecuteStampTaskLogResult struct {
	Items         []ExecuteStampTaskLogCount `json:"items"`
	NextPageToken *string                    `json:"nextPageToken"`
	TotalCount    *int64                     `json:"totalCount"`
	ScanSize      *int64                     `json:"scanSize"`
	Metadata      *core.ResultMetadata       `json:"metadata"`
}

type CountExecuteStampTaskLogAsyncResult struct {
	result *CountExecuteStampTaskLogResult
	err    error
}

func NewCountExecuteStampTaskLogResultFromJson(data string) CountExecuteStampTaskLogResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountExecuteStampTaskLogResultFromDict(dict)
}

func NewCountExecuteStampTaskLogResultFromDict(data map[string]interface{}) CountExecuteStampTaskLogResult {
	return CountExecuteStampTaskLogResult{
		Items: func() []ExecuteStampTaskLogCount {
			if data["items"] == nil {
				return nil
			}
			return CastExecuteStampTaskLogCounts(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		TotalCount: func() *int64 {
			v, ok := data["totalCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["totalCount"])
		}(),
		ScanSize: func() *int64 {
			v, ok := data["scanSize"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["scanSize"])
		}(),
	}
}

func (p CountExecuteStampTaskLogResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastExecuteStampTaskLogCountsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"totalCount":    p.TotalCount,
		"scanSize":      p.ScanSize,
	}
}

func (p CountExecuteStampTaskLogResult) Pointer() *CountExecuteStampTaskLogResult {
	return &p
}

type QueryInGameLogResult struct {
	Items         []InGameLog          `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	TotalCount    *int64               `json:"totalCount"`
	ScanSize      *int64               `json:"scanSize"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type QueryInGameLogAsyncResult struct {
	result *QueryInGameLogResult
	err    error
}

func NewQueryInGameLogResultFromJson(data string) QueryInGameLogResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQueryInGameLogResultFromDict(dict)
}

func NewQueryInGameLogResultFromDict(data map[string]interface{}) QueryInGameLogResult {
	return QueryInGameLogResult{
		Items: func() []InGameLog {
			if data["items"] == nil {
				return nil
			}
			return CastInGameLogs(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		TotalCount: func() *int64 {
			v, ok := data["totalCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["totalCount"])
		}(),
		ScanSize: func() *int64 {
			v, ok := data["scanSize"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["scanSize"])
		}(),
	}
}

func (p QueryInGameLogResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInGameLogsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"totalCount":    p.TotalCount,
		"scanSize":      p.ScanSize,
	}
}

func (p QueryInGameLogResult) Pointer() *QueryInGameLogResult {
	return &p
}

type SendInGameLogResult struct {
	Item     *InGameLog           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SendInGameLogAsyncResult struct {
	result *SendInGameLogResult
	err    error
}

func NewSendInGameLogResultFromJson(data string) SendInGameLogResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendInGameLogResultFromDict(dict)
}

func NewSendInGameLogResultFromDict(data map[string]interface{}) SendInGameLogResult {
	return SendInGameLogResult{
		Item: func() *InGameLog {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInGameLogFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p SendInGameLogResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p SendInGameLogResult) Pointer() *SendInGameLogResult {
	return &p
}

type SendInGameLogByUserIdResult struct {
	Item     *InGameLog           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SendInGameLogByUserIdAsyncResult struct {
	result *SendInGameLogByUserIdResult
	err    error
}

func NewSendInGameLogByUserIdResultFromJson(data string) SendInGameLogByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendInGameLogByUserIdResultFromDict(dict)
}

func NewSendInGameLogByUserIdResultFromDict(data map[string]interface{}) SendInGameLogByUserIdResult {
	return SendInGameLogByUserIdResult{
		Item: func() *InGameLog {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInGameLogFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p SendInGameLogByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p SendInGameLogByUserIdResult) Pointer() *SendInGameLogByUserIdResult {
	return &p
}

type QueryAccessLogWithTelemetryResult struct {
	Items         []AccessLogWithTelemetry `json:"items"`
	NextPageToken *string                  `json:"nextPageToken"`
	TotalCount    *int64                   `json:"totalCount"`
	ScanSize      *int64                   `json:"scanSize"`
	Metadata      *core.ResultMetadata     `json:"metadata"`
}

type QueryAccessLogWithTelemetryAsyncResult struct {
	result *QueryAccessLogWithTelemetryResult
	err    error
}

func NewQueryAccessLogWithTelemetryResultFromJson(data string) QueryAccessLogWithTelemetryResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQueryAccessLogWithTelemetryResultFromDict(dict)
}

func NewQueryAccessLogWithTelemetryResultFromDict(data map[string]interface{}) QueryAccessLogWithTelemetryResult {
	return QueryAccessLogWithTelemetryResult{
		Items: func() []AccessLogWithTelemetry {
			if data["items"] == nil {
				return nil
			}
			return CastAccessLogWithTelemetries(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		TotalCount: func() *int64 {
			v, ok := data["totalCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["totalCount"])
		}(),
		ScanSize: func() *int64 {
			v, ok := data["scanSize"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["scanSize"])
		}(),
	}
}

func (p QueryAccessLogWithTelemetryResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAccessLogWithTelemetriesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"totalCount":    p.TotalCount,
		"scanSize":      p.ScanSize,
	}
}

func (p QueryAccessLogWithTelemetryResult) Pointer() *QueryAccessLogWithTelemetryResult {
	return &p
}

type DescribeInsightsResult struct {
	Items         []Insight            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeInsightsAsyncResult struct {
	result *DescribeInsightsResult
	err    error
}

func NewDescribeInsightsResultFromJson(data string) DescribeInsightsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInsightsResultFromDict(dict)
}

func NewDescribeInsightsResultFromDict(data map[string]interface{}) DescribeInsightsResult {
	return DescribeInsightsResult{
		Items: func() []Insight {
			if data["items"] == nil {
				return nil
			}
			return CastInsights(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeInsightsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInsightsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeInsightsResult) Pointer() *DescribeInsightsResult {
	return &p
}

type CreateInsightResult struct {
	Item     *Insight             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateInsightAsyncResult struct {
	result *CreateInsightResult
	err    error
}

func NewCreateInsightResultFromJson(data string) CreateInsightResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateInsightResultFromDict(dict)
}

func NewCreateInsightResultFromDict(data map[string]interface{}) CreateInsightResult {
	return CreateInsightResult{
		Item: func() *Insight {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInsightFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateInsightResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateInsightResult) Pointer() *CreateInsightResult {
	return &p
}

type GetInsightResult struct {
	Item     *Insight             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetInsightAsyncResult struct {
	result *GetInsightResult
	err    error
}

func NewGetInsightResultFromJson(data string) GetInsightResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetInsightResultFromDict(dict)
}

func NewGetInsightResultFromDict(data map[string]interface{}) GetInsightResult {
	return GetInsightResult{
		Item: func() *Insight {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInsightFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetInsightResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetInsightResult) Pointer() *GetInsightResult {
	return &p
}

type DeleteInsightResult struct {
	Item     *Insight             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteInsightAsyncResult struct {
	result *DeleteInsightResult
	err    error
}

func NewDeleteInsightResultFromJson(data string) DeleteInsightResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteInsightResultFromDict(dict)
}

func NewDeleteInsightResultFromDict(data map[string]interface{}) DeleteInsightResult {
	return DeleteInsightResult{
		Item: func() *Insight {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInsightFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteInsightResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteInsightResult) Pointer() *DeleteInsightResult {
	return &p
}
