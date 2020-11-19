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

package realtime

type Namespace struct {
	/** ネームスペース */
	NamespaceId *string `json:"namespaceId"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** ネームスペース名 */
	Name *string `json:"name"`
	/** ネームスペースの説明 */
	Description *string `json:"description"`
	/** サーバの種類 */
	ServerType *string `json:"serverType"`
	/** サーバのスペック */
	ServerSpec *string `json:"serverSpec"`
	/** ルームの作成が終わったときのプッシュ通知 */
	CreateNotification *NotificationSetting `json:"createNotification"`
	/** ログの出力設定 */
	LogSetting *LogSetting `json:"logSetting"`
	/** None */
	Status *string `json:"status"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
}

func (p *Namespace) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["namespaceId"] = p.NamespaceId
	data["ownerId"] = p.OwnerId
	data["name"] = p.Name
	data["description"] = p.Description
	data["serverType"] = p.ServerType
	data["serverSpec"] = p.ServerSpec
	if p.CreateNotification != nil {
		data["createNotification"] = *p.CreateNotification.ToDict()
	}
	if p.LogSetting != nil {
		data["logSetting"] = *p.LogSetting.ToDict()
	}
	data["status"] = p.Status
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type Room struct {
	/** ルーム */
	RoomId *string `json:"roomId"`
	/** ルーム名 */
	Name *string `json:"name"`
	/** IPアドレス */
	IpAddress *string `json:"ipAddress"`
	/** 待受ポート */
	Port *int32 `json:"port"`
	/** 暗号鍵 */
	EncryptionKey *string `json:"encryptionKey"`
	/** ルームの作成が終わったときに通知を受けるユーザIDリスト */
	NotificationUserIds *[]string `json:"notificationUserIds"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
}

func (p *Room) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["roomId"] = p.RoomId
	data["name"] = p.Name
	data["ipAddress"] = p.IpAddress
	data["port"] = p.Port
	data["encryptionKey"] = p.EncryptionKey
	if p.NotificationUserIds != nil {
		var _notificationUserIds []string
		for _, item := range *p.NotificationUserIds {
			_notificationUserIds = append(_notificationUserIds, item)
		}
		data["notificationUserIds"] = &_notificationUserIds
	}
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type NotificationSetting struct {
	/** プッシュ通知に使用する GS2-Gateway のネームスペース のGRN */
	GatewayNamespaceId *string `json:"gatewayNamespaceId"`
	/** モバイルプッシュ通知へ転送するか */
	EnableTransferMobileNotification *bool `json:"enableTransferMobileNotification"`
	/** モバイルプッシュ通知で使用するサウンドファイル名 */
	Sound *string `json:"sound"`
}

func (p *NotificationSetting) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["gatewayNamespaceId"] = p.GatewayNamespaceId
	data["enableTransferMobileNotification"] = p.EnableTransferMobileNotification
	data["sound"] = p.Sound
	return &data
}

type LogSetting struct {
	/** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["loggingNamespaceId"] = p.LoggingNamespaceId
	return &data
}
