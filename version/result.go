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

package version

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         []Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeNamespacesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Namespace, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

type CreateNamespaceResult struct {
    /** 作成したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *CreateNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

type GetNamespaceStatusResult struct {
    /** None */
	Status         *core.String	`json:"status"`
}

func (p *GetNamespaceStatusResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Status != nil {
        data["status"] = p.Status
    }
    return &data
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

type GetNamespaceResult struct {
    /** ネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *GetNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

type UpdateNamespaceResult struct {
    /** 更新したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *UpdateNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

type DeleteNamespaceResult struct {
    /** 削除したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type DescribeVersionModelMastersResult struct {
    /** バージョンマスターのリスト */
	Items         []VersionModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeVersionModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]VersionModelMaster, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeVersionModelMastersAsyncResult struct {
	result *DescribeVersionModelMastersResult
	err    error
}

type CreateVersionModelMasterResult struct {
    /** 作成したバージョンマスター */
	Item         *VersionModelMaster	`json:"item"`
}

func (p *CreateVersionModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateVersionModelMasterAsyncResult struct {
	result *CreateVersionModelMasterResult
	err    error
}

type GetVersionModelMasterResult struct {
    /** バージョンマスター */
	Item         *VersionModelMaster	`json:"item"`
}

func (p *GetVersionModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetVersionModelMasterAsyncResult struct {
	result *GetVersionModelMasterResult
	err    error
}

type UpdateVersionModelMasterResult struct {
    /** 更新したバージョンマスター */
	Item         *VersionModelMaster	`json:"item"`
}

func (p *UpdateVersionModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateVersionModelMasterAsyncResult struct {
	result *UpdateVersionModelMasterResult
	err    error
}

type DeleteVersionModelMasterResult struct {
    /** 削除したバージョンマスター */
	Item         *VersionModelMaster	`json:"item"`
}

func (p *DeleteVersionModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteVersionModelMasterAsyncResult struct {
	result *DeleteVersionModelMasterResult
	err    error
}

type DescribeVersionModelsResult struct {
    /** バージョン設定のリスト */
	Items         []VersionModel	`json:"items"`
}

func (p *DescribeVersionModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]VersionModel, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeVersionModelsAsyncResult struct {
	result *DescribeVersionModelsResult
	err    error
}

type GetVersionModelResult struct {
    /** バージョン設定 */
	Item         *VersionModel	`json:"item"`
}

func (p *GetVersionModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetVersionModelAsyncResult struct {
	result *GetVersionModelResult
	err    error
}

type DescribeAcceptVersionsResult struct {
    /** 承認したバージョンのリスト */
	Items         []AcceptVersion	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeAcceptVersionsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]AcceptVersion, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeAcceptVersionsAsyncResult struct {
	result *DescribeAcceptVersionsResult
	err    error
}

type DescribeAcceptVersionsByUserIdResult struct {
    /** 承認したバージョンのリスト */
	Items         []AcceptVersion	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeAcceptVersionsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]AcceptVersion, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeAcceptVersionsByUserIdAsyncResult struct {
	result *DescribeAcceptVersionsByUserIdResult
	err    error
}

type AcceptResult struct {
    /** 承認したバージョン */
	Item         *AcceptVersion	`json:"item"`
}

func (p *AcceptResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type AcceptAsyncResult struct {
	result *AcceptResult
	err    error
}

type AcceptByUserIdResult struct {
    /** 承認したバージョン */
	Item         *AcceptVersion	`json:"item"`
}

func (p *AcceptByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type AcceptByUserIdAsyncResult struct {
	result *AcceptByUserIdResult
	err    error
}

type GetAcceptVersionResult struct {
    /** 承認したバージョン */
	Item         *AcceptVersion	`json:"item"`
}

func (p *GetAcceptVersionResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetAcceptVersionAsyncResult struct {
	result *GetAcceptVersionResult
	err    error
}

type GetAcceptVersionByUserIdResult struct {
    /** 承認したバージョン */
	Item         *AcceptVersion	`json:"item"`
}

func (p *GetAcceptVersionByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetAcceptVersionByUserIdAsyncResult struct {
	result *GetAcceptVersionByUserIdResult
	err    error
}

type DeleteAcceptVersionResult struct {
}

func (p *DeleteAcceptVersionResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteAcceptVersionAsyncResult struct {
	result *DeleteAcceptVersionResult
	err    error
}

type DeleteAcceptVersionByUserIdResult struct {
}

func (p *DeleteAcceptVersionByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteAcceptVersionByUserIdAsyncResult struct {
	result *DeleteAcceptVersionByUserIdResult
	err    error
}

type CheckVersionResult struct {
    /** プロジェクトトークン */
	ProjectToken         *core.String	`json:"projectToken"`
    /** バージョンの検証結果のリスト */
	Warnings         []Status	`json:"warnings"`
    /** バージョンの検証結果のリスト */
	Errors         []Status	`json:"errors"`
}

func (p *CheckVersionResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.ProjectToken != nil {
        data["projectToken"] = p.ProjectToken
    }
    if p.Warnings != nil {
    	items := make([]Status, 0)
    	for _, item := range p.Warnings {
			items = append(items, item)
		}
		data["warnings"] = items
    }
    if p.Errors != nil {
    	items := make([]Status, 0)
    	for _, item := range p.Errors {
			items = append(items, item)
		}
		data["errors"] = items
    }
    return &data
}

type CheckVersionAsyncResult struct {
	result *CheckVersionResult
	err    error
}

type CheckVersionByUserIdResult struct {
    /** プロジェクトトークン */
	ProjectToken         *core.String	`json:"projectToken"`
    /** バージョンの検証結果のリスト */
	Warnings         []Status	`json:"warnings"`
    /** バージョンの検証結果のリスト */
	Errors         []Status	`json:"errors"`
}

func (p *CheckVersionByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.ProjectToken != nil {
        data["projectToken"] = p.ProjectToken
    }
    if p.Warnings != nil {
    	items := make([]Status, 0)
    	for _, item := range p.Warnings {
			items = append(items, item)
		}
		data["warnings"] = items
    }
    if p.Errors != nil {
    	items := make([]Status, 0)
    	for _, item := range p.Errors {
			items = append(items, item)
		}
		data["errors"] = items
    }
    return &data
}

type CheckVersionByUserIdAsyncResult struct {
	result *CheckVersionByUserIdResult
	err    error
}

type CalculateSignatureResult struct {
    /** ボディ */
	Body         *core.String	`json:"body"`
    /** 署名 */
	Signature         *core.String	`json:"signature"`
}

func (p *CalculateSignatureResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Body != nil {
        data["body"] = p.Body
    }
    if p.Signature != nil {
        data["signature"] = p.Signature
    }
    return &data
}

type CalculateSignatureAsyncResult struct {
	result *CalculateSignatureResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効なバージョン */
	Item         *CurrentVersionMaster	`json:"item"`
}

func (p *ExportMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

type GetCurrentVersionMasterResult struct {
    /** 現在有効なバージョン */
	Item         *CurrentVersionMaster	`json:"item"`
}

func (p *GetCurrentVersionMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentVersionMasterAsyncResult struct {
	result *GetCurrentVersionMasterResult
	err    error
}

type UpdateCurrentVersionMasterResult struct {
    /** 更新した現在有効なバージョン */
	Item         *CurrentVersionMaster	`json:"item"`
}

func (p *UpdateCurrentVersionMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentVersionMasterAsyncResult struct {
	result *UpdateCurrentVersionMasterResult
	err    error
}

type UpdateCurrentVersionMasterFromGitHubResult struct {
    /** 更新した現在有効なバージョン */
	Item         *CurrentVersionMaster	`json:"item"`
}

func (p *UpdateCurrentVersionMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentVersionMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentVersionMasterFromGitHubResult
	err    error
}
