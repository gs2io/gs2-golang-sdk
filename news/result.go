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

package news

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
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type PrepareUpdateCurrentNewsMasterResult struct {
    /** アップロード後に結果を反映する際に使用するトークン */
	UploadToken         *core.String	`json:"uploadToken"`
    /** テンプレートアップロード処理の実行に使用するURL */
	TemplateUploadUrl         *core.String	`json:"templateUploadUrl"`
}

func (p *PrepareUpdateCurrentNewsMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.UploadToken != nil {
        data["uploadToken"] = p.UploadToken
    }
    if p.TemplateUploadUrl != nil {
        data["templateUploadUrl"] = p.TemplateUploadUrl
    }
    return &data
}

type PrepareUpdateCurrentNewsMasterAsyncResult struct {
	result *PrepareUpdateCurrentNewsMasterResult
	err    error
}

type UpdateCurrentNewsMasterResult struct {
}

func (p *UpdateCurrentNewsMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type UpdateCurrentNewsMasterAsyncResult struct {
	result *UpdateCurrentNewsMasterResult
	err    error
}

type PrepareUpdateCurrentNewsMasterFromGitHubResult struct {
    /** アップロード後に結果を反映する際に使用するトークン */
	UploadToken         *core.String	`json:"uploadToken"`
}

func (p *PrepareUpdateCurrentNewsMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.UploadToken != nil {
        data["uploadToken"] = p.UploadToken
    }
    return &data
}

type PrepareUpdateCurrentNewsMasterFromGitHubAsyncResult struct {
	result *PrepareUpdateCurrentNewsMasterFromGitHubResult
	err    error
}

type DescribeNewsResult struct {
    /** お知らせ記事のリスト */
	Items         []News	`json:"items"`
    /** お知らせ記事データのハッシュ値 */
	ContentHash         *core.String	`json:"contentHash"`
    /** テンプレートデータのハッシュ値 */
	TemplateHash         *core.String	`json:"templateHash"`
}

func (p *DescribeNewsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]News, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.ContentHash != nil {
        data["contentHash"] = p.ContentHash
    }
    if p.TemplateHash != nil {
        data["templateHash"] = p.TemplateHash
    }
    return &data
}

type DescribeNewsAsyncResult struct {
	result *DescribeNewsResult
	err    error
}

type DescribeNewsByUserIdResult struct {
    /** お知らせ記事のリスト */
	Items         []News	`json:"items"`
    /** お知らせ記事データのハッシュ値 */
	ContentHash         *core.String	`json:"contentHash"`
    /** テンプレートデータのハッシュ値 */
	TemplateHash         *core.String	`json:"templateHash"`
}

func (p *DescribeNewsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]News, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.ContentHash != nil {
        data["contentHash"] = p.ContentHash
    }
    if p.TemplateHash != nil {
        data["templateHash"] = p.TemplateHash
    }
    return &data
}

type DescribeNewsByUserIdAsyncResult struct {
	result *DescribeNewsByUserIdResult
	err    error
}

type WantGrantResult struct {
    /** お知らせコンテンツにアクセスするために設定の必要なクッキー のリスト */
	Items         []SetCookieRequestEntry	`json:"items"`
    /** お知らせコンテンツにアクセスするためのURL */
	BrowserUrl         *core.String	`json:"browserUrl"`
    /** ZIP形式のお知らせコンテンツにアクセスするためのURL Cookieの設定は不要 */
	ZipUrl         *core.String	`json:"zipUrl"`
}

func (p *WantGrantResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]SetCookieRequestEntry, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.BrowserUrl != nil {
        data["browserUrl"] = p.BrowserUrl
    }
    if p.ZipUrl != nil {
        data["zipUrl"] = p.ZipUrl
    }
    return &data
}

type WantGrantAsyncResult struct {
	result *WantGrantResult
	err    error
}

type WantGrantByUserIdResult struct {
    /** お知らせコンテンツにアクセスするために設定の必要なクッキー のリスト */
	Items         []SetCookieRequestEntry	`json:"items"`
    /** お知らせコンテンツにアクセスするためのURL */
	BrowserUrl         *core.String	`json:"browserUrl"`
    /** ZIP形式のお知らせコンテンツにアクセスするためのURL Cookieの設定は不要 */
	ZipUrl         *core.String	`json:"zipUrl"`
}

func (p *WantGrantByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]SetCookieRequestEntry, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.BrowserUrl != nil {
        data["browserUrl"] = p.BrowserUrl
    }
    if p.ZipUrl != nil {
        data["zipUrl"] = p.ZipUrl
    }
    return &data
}

type WantGrantByUserIdAsyncResult struct {
	result *WantGrantByUserIdResult
	err    error
}
