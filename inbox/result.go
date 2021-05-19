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

package inbox

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         []Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
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
	Status         *string	`json:"status"`
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

type DescribeMessagesResult struct {
    /** メッセージのリスト */
	Items         []Message	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeMessagesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Message, 0)
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

type DescribeMessagesAsyncResult struct {
	result *DescribeMessagesResult
	err    error
}

type DescribeMessagesByUserIdResult struct {
    /** メッセージのリスト */
	Items         []Message	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeMessagesByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Message, 0)
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

type DescribeMessagesByUserIdAsyncResult struct {
	result *DescribeMessagesByUserIdResult
	err    error
}

type SendMessageByUserIdResult struct {
    /** 作成したメッセージ */
	Item         *Message	`json:"item"`
}

func (p *SendMessageByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SendMessageByUserIdAsyncResult struct {
	result *SendMessageByUserIdResult
	err    error
}

type GetMessageResult struct {
    /** メッセージ */
	Item         *Message	`json:"item"`
}

func (p *GetMessageResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetMessageAsyncResult struct {
	result *GetMessageResult
	err    error
}

type GetMessageByUserIdResult struct {
    /** メッセージ */
	Item         *Message	`json:"item"`
}

func (p *GetMessageByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetMessageByUserIdAsyncResult struct {
	result *GetMessageByUserIdResult
	err    error
}

type ReceiveGlobalMessageResult struct {
    /** 受信したメッセージ一覧 */
	Item         []Message	`json:"item"`
}

func (p *ReceiveGlobalMessageResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    	items := make([]Message, 0)
    	for _, item := range p.Item {
			items = append(items, item)
		}
		data["item"] = items
    }
    return &data
}

type ReceiveGlobalMessageAsyncResult struct {
	result *ReceiveGlobalMessageResult
	err    error
}

type ReceiveGlobalMessageByUserIdResult struct {
    /** 受信したメッセージ一覧 */
	Item         []Message	`json:"item"`
}

func (p *ReceiveGlobalMessageByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    	items := make([]Message, 0)
    	for _, item := range p.Item {
			items = append(items, item)
		}
		data["item"] = items
    }
    return &data
}

type ReceiveGlobalMessageByUserIdAsyncResult struct {
	result *ReceiveGlobalMessageByUserIdResult
	err    error
}

type OpenMessageResult struct {
    /** メッセージ */
	Item         *Message	`json:"item"`
}

func (p *OpenMessageResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type OpenMessageAsyncResult struct {
	result *OpenMessageResult
	err    error
}

type OpenMessageByUserIdResult struct {
    /** メッセージ */
	Item         *Message	`json:"item"`
}

func (p *OpenMessageByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type OpenMessageByUserIdAsyncResult struct {
	result *OpenMessageByUserIdResult
	err    error
}

type ReadMessageResult struct {
    /** メッセージ */
	Item         *Message	`json:"item"`
    /** スタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
}

func (p *ReadMessageResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type ReadMessageAsyncResult struct {
	result *ReadMessageResult
	err    error
}

type ReadMessageByUserIdResult struct {
    /** メッセージ */
	Item         *Message	`json:"item"`
    /** スタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
}

func (p *ReadMessageByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type ReadMessageByUserIdAsyncResult struct {
	result *ReadMessageByUserIdResult
	err    error
}

type DeleteMessageResult struct {
}

func (p *DeleteMessageResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteMessageAsyncResult struct {
	result *DeleteMessageResult
	err    error
}

type DeleteMessageByUserIdResult struct {
}

func (p *DeleteMessageByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteMessageByUserIdAsyncResult struct {
	result *DeleteMessageByUserIdResult
	err    error
}

type OpenByStampTaskResult struct {
    /** メッセージ */
	Item         *Message	`json:"item"`
    /** スタンプタスクの実行結果を記録したコンテキスト */
	NewContextStack         *string	`json:"newContextStack"`
}

func (p *OpenByStampTaskResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.NewContextStack != nil {
        data["newContextStack"] = p.NewContextStack
    }
    return &data
}

type OpenByStampTaskAsyncResult struct {
	result *OpenByStampTaskResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効なグローバルメッセージ設定 */
	Item         *CurrentMessageMaster	`json:"item"`
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

type GetCurrentMessageMasterResult struct {
    /** 現在有効なグローバルメッセージ設定 */
	Item         *CurrentMessageMaster	`json:"item"`
}

func (p *GetCurrentMessageMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentMessageMasterAsyncResult struct {
	result *GetCurrentMessageMasterResult
	err    error
}

type UpdateCurrentMessageMasterResult struct {
    /** 更新した現在有効なグローバルメッセージ設定 */
	Item         *CurrentMessageMaster	`json:"item"`
}

func (p *UpdateCurrentMessageMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentMessageMasterAsyncResult struct {
	result *UpdateCurrentMessageMasterResult
	err    error
}

type UpdateCurrentMessageMasterFromGitHubResult struct {
    /** 更新した現在有効なグローバルメッセージ設定 */
	Item         *CurrentMessageMaster	`json:"item"`
}

func (p *UpdateCurrentMessageMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentMessageMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentMessageMasterFromGitHubResult
	err    error
}

type DescribeGlobalMessageMastersResult struct {
    /** 全ユーザに向けたメッセージのリスト */
	Items         []GlobalMessageMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeGlobalMessageMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]GlobalMessageMaster, 0)
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

type DescribeGlobalMessageMastersAsyncResult struct {
	result *DescribeGlobalMessageMastersResult
	err    error
}

type CreateGlobalMessageMasterResult struct {
    /** 作成した全ユーザに向けたメッセージ */
	Item         *GlobalMessageMaster	`json:"item"`
}

func (p *CreateGlobalMessageMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateGlobalMessageMasterAsyncResult struct {
	result *CreateGlobalMessageMasterResult
	err    error
}

type GetGlobalMessageMasterResult struct {
    /** 全ユーザに向けたメッセージ */
	Item         *GlobalMessageMaster	`json:"item"`
}

func (p *GetGlobalMessageMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetGlobalMessageMasterAsyncResult struct {
	result *GetGlobalMessageMasterResult
	err    error
}

type UpdateGlobalMessageMasterResult struct {
    /** 全ユーザに向けたメッセージ */
	Item         *GlobalMessageMaster	`json:"item"`
}

func (p *UpdateGlobalMessageMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateGlobalMessageMasterAsyncResult struct {
	result *UpdateGlobalMessageMasterResult
	err    error
}

type DeleteGlobalMessageMasterResult struct {
}

func (p *DeleteGlobalMessageMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteGlobalMessageMasterAsyncResult struct {
	result *DeleteGlobalMessageMasterResult
	err    error
}

type DescribeGlobalMessagesResult struct {
    /** 全ユーザに向けたメッセージのリスト */
	Items         []GlobalMessage	`json:"items"`
}

func (p *DescribeGlobalMessagesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]GlobalMessage, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeGlobalMessagesAsyncResult struct {
	result *DescribeGlobalMessagesResult
	err    error
}

type GetGlobalMessageResult struct {
    /** 全ユーザに向けたメッセージ */
	Item         *GlobalMessage	`json:"item"`
}

func (p *GetGlobalMessageResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetGlobalMessageAsyncResult struct {
	result *GetGlobalMessageResult
	err    error
}

type GetReceivedByUserIdResult struct {
    /** 受信済みグローバルメッセージ名 */
	Item         *Received	`json:"item"`
}

func (p *GetReceivedByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetReceivedByUserIdAsyncResult struct {
	result *GetReceivedByUserIdResult
	err    error
}

type UpdateReceivedByUserIdResult struct {
}

func (p *UpdateReceivedByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type UpdateReceivedByUserIdAsyncResult struct {
	result *UpdateReceivedByUserIdResult
	err    error
}

type DeleteReceivedByUserIdResult struct {
}

func (p *DeleteReceivedByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteReceivedByUserIdAsyncResult struct {
	result *DeleteReceivedByUserIdResult
	err    error
}
