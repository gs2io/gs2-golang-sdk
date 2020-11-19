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

package key

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

type DescribeKeysResult struct {
	/** 暗号鍵のリスト */
	Items *[]*Key `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeKeysResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeKeysAsyncResult struct {
	result *DescribeKeysResult
	err    error
}

type CreateKeyResult struct {
	/** 作成した暗号鍵 */
	Item *Key `json:"item"`
}

func (p *CreateKeyResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type CreateKeyAsyncResult struct {
	result *CreateKeyResult
	err    error
}

type UpdateKeyResult struct {
	/** 更新した暗号鍵 */
	Item *Key `json:"item"`
}

func (p *UpdateKeyResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UpdateKeyAsyncResult struct {
	result *UpdateKeyResult
	err    error
}

type GetKeyResult struct {
	/** 暗号鍵 */
	Item *Key `json:"item"`
}

func (p *GetKeyResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetKeyAsyncResult struct {
	result *GetKeyResult
	err    error
}

type DeleteKeyResult struct {
}

func (p *DeleteKeyResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	return &data
}

type DeleteKeyAsyncResult struct {
	result *DeleteKeyResult
	err    error
}

type EncryptResult struct {
	/** 暗号化済みデータ */
	Data *string `json:"data"`
}

func (p *EncryptResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Data != nil {
		data["Data"] = p.Data
	}
	return &data
}

type EncryptAsyncResult struct {
	result *EncryptResult
	err    error
}

type DecryptResult struct {
	/** 復号済みデータ */
	Data *string `json:"data"`
}

func (p *DecryptResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Data != nil {
		data["Data"] = p.Data
	}
	return &data
}

type DecryptAsyncResult struct {
	result *DecryptResult
	err    error
}

type DescribeGitHubApiKeysResult struct {
	/** GitHub のAPIキーのリスト */
	Items *[]*GitHubApiKey `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeGitHubApiKeysResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeGitHubApiKeysAsyncResult struct {
	result *DescribeGitHubApiKeysResult
	err    error
}

type CreateGitHubApiKeyResult struct {
	/** 作成したGitHub のAPIキー */
	Item *GitHubApiKey `json:"item"`
}

func (p *CreateGitHubApiKeyResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type CreateGitHubApiKeyAsyncResult struct {
	result *CreateGitHubApiKeyResult
	err    error
}

type UpdateGitHubApiKeyResult struct {
	/** 更新したGitHub のAPIキー */
	Item *GitHubApiKey `json:"item"`
}

func (p *UpdateGitHubApiKeyResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UpdateGitHubApiKeyAsyncResult struct {
	result *UpdateGitHubApiKeyResult
	err    error
}

type GetGitHubApiKeyResult struct {
	/** GitHub のAPIキー */
	Item *GitHubApiKey `json:"item"`
}

func (p *GetGitHubApiKeyResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetGitHubApiKeyAsyncResult struct {
	result *GetGitHubApiKeyResult
	err    error
}

type DeleteGitHubApiKeyResult struct {
}

func (p *DeleteGitHubApiKeyResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	return &data
}

type DeleteGitHubApiKeyAsyncResult struct {
	result *DeleteGitHubApiKeyResult
	err    error
}
