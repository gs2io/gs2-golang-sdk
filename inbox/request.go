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

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	PageToken    *string            `json:"pageToken"`
	Limit        *int64             `json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId                  *core.RequestId      `json:"requestId"`
	ContextStack               *core.ContextStack   `json:"contextStack"`
	Name                       *string              `json:"name"`
	Description                *string              `json:"description"`
	IsAutomaticDeletingEnabled *bool                `json:"isAutomaticDeletingEnabled"`
	ReceiveMessageScript       *ScriptSetting       `json:"receiveMessageScript"`
	ReadMessageScript          *ScriptSetting       `json:"readMessageScript"`
	DeleteMessageScript        *ScriptSetting       `json:"deleteMessageScript"`
	QueueNamespaceId           *string              `json:"queueNamespaceId"`
	KeyId                      *string              `json:"keyId"`
	ReceiveNotification        *NotificationSetting `json:"receiveNotification"`
	LogSetting                 *LogSetting          `json:"logSetting"`
}

type GetNamespaceStatusRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateNamespaceRequest struct {
	RequestId                  *core.RequestId      `json:"requestId"`
	ContextStack               *core.ContextStack   `json:"contextStack"`
	NamespaceName              *string              `json:"namespaceName"`
	Description                *string              `json:"description"`
	IsAutomaticDeletingEnabled *bool                `json:"isAutomaticDeletingEnabled"`
	ReceiveMessageScript       *ScriptSetting       `json:"receiveMessageScript"`
	ReadMessageScript          *ScriptSetting       `json:"readMessageScript"`
	DeleteMessageScript        *ScriptSetting       `json:"deleteMessageScript"`
	QueueNamespaceId           *string              `json:"queueNamespaceId"`
	KeyId                      *string              `json:"keyId"`
	ReceiveNotification        *NotificationSetting `json:"receiveNotification"`
	LogSetting                 *LogSetting          `json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeMessagesRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeMessagesByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type SendMessageByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Metadata               *string            `json:"metadata"`
	ReadAcquireActions     *[]*AcquireAction  `json:"readAcquireActions"`
	ExpiresAt              *int64             `json:"expiresAt"`
	ExpiresTimeSpan        *TimeSpan          `json:"expiresTimeSpan"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetMessageRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	MessageName            *string            `json:"messageName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetMessageByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	MessageName            *string            `json:"messageName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type ReceiveGlobalMessageRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type ReceiveGlobalMessageByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type OpenMessageRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	MessageName            *string            `json:"messageName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type OpenMessageByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	MessageName            *string            `json:"messageName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type ReadMessageRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	MessageName            *string            `json:"messageName"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type ReadMessageByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	MessageName            *string            `json:"messageName"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteMessageRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	MessageName            *string            `json:"messageName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DeleteMessageByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	MessageName            *string            `json:"messageName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type OpenByStampTaskRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampTask              *string            `json:"stampTask"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type ExportMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetCurrentMessageMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateCurrentMessageMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Settings      *string            `json:"settings"`
}

type UpdateCurrentMessageMasterFromGitHubRequest struct {
	RequestId       *core.RequestId        `json:"requestId"`
	ContextStack    *core.ContextStack     `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

type DescribeGlobalMessageMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateGlobalMessageMasterRequest struct {
	RequestId          *core.RequestId    `json:"requestId"`
	ContextStack       *core.ContextStack `json:"contextStack"`
	NamespaceName      *string            `json:"namespaceName"`
	Name               *string            `json:"name"`
	Metadata           *string            `json:"metadata"`
	ReadAcquireActions *[]*AcquireAction  `json:"readAcquireActions"`
	ExpiresTimeSpan    *TimeSpan          `json:"expiresTimeSpan"`
	ExpiresAt          *int64             `json:"expiresAt"`
}

type GetGlobalMessageMasterRequest struct {
	RequestId         *core.RequestId    `json:"requestId"`
	ContextStack      *core.ContextStack `json:"contextStack"`
	NamespaceName     *string            `json:"namespaceName"`
	GlobalMessageName *string            `json:"globalMessageName"`
}

type UpdateGlobalMessageMasterRequest struct {
	RequestId          *core.RequestId    `json:"requestId"`
	ContextStack       *core.ContextStack `json:"contextStack"`
	NamespaceName      *string            `json:"namespaceName"`
	GlobalMessageName  *string            `json:"globalMessageName"`
	Metadata           *string            `json:"metadata"`
	ReadAcquireActions *[]*AcquireAction  `json:"readAcquireActions"`
	ExpiresTimeSpan    *TimeSpan          `json:"expiresTimeSpan"`
	ExpiresAt          *int64             `json:"expiresAt"`
}

type DeleteGlobalMessageMasterRequest struct {
	RequestId         *core.RequestId    `json:"requestId"`
	ContextStack      *core.ContextStack `json:"contextStack"`
	NamespaceName     *string            `json:"namespaceName"`
	GlobalMessageName *string            `json:"globalMessageName"`
}

type DescribeGlobalMessagesRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetGlobalMessageRequest struct {
	RequestId         *core.RequestId    `json:"requestId"`
	ContextStack      *core.ContextStack `json:"contextStack"`
	NamespaceName     *string            `json:"namespaceName"`
	GlobalMessageName *string            `json:"globalMessageName"`
}

type GetReceivedByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type UpdateReceivedByUserIdRequest struct {
	RequestId                  *core.RequestId    `json:"requestId"`
	ContextStack               *core.ContextStack `json:"contextStack"`
	NamespaceName              *string            `json:"namespaceName"`
	UserId                     *string            `json:"userId"`
	ReceivedGlobalMessageNames *[]string          `json:"receivedGlobalMessageNames"`
	XGs2DuplicationAvoider     *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteReceivedByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}
