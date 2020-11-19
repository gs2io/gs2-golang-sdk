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

package showcase

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
	RequestId        *core.RequestId    `json:"requestId"`
	ContextStack     *core.ContextStack `json:"contextStack"`
	Name             *string            `json:"name"`
	Description      *string            `json:"description"`
	QueueNamespaceId *string            `json:"queueNamespaceId"`
	KeyId            *string            `json:"keyId"`
	LogSetting       *LogSetting        `json:"logSetting"`
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
	RequestId        *core.RequestId    `json:"requestId"`
	ContextStack     *core.ContextStack `json:"contextStack"`
	NamespaceName    *string            `json:"namespaceName"`
	Description      *string            `json:"description"`
	QueueNamespaceId *string            `json:"queueNamespaceId"`
	KeyId            *string            `json:"keyId"`
	LogSetting       *LogSetting        `json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeSalesItemMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateSalesItemMasterRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	Name           *string            `json:"name"`
	Description    *string            `json:"description"`
	Metadata       *string            `json:"metadata"`
	ConsumeActions *[]*ConsumeAction  `json:"consumeActions"`
	AcquireActions *[]*AcquireAction  `json:"acquireActions"`
}

type GetSalesItemMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	SalesItemName *string            `json:"salesItemName"`
}

type UpdateSalesItemMasterRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	SalesItemName  *string            `json:"salesItemName"`
	Description    *string            `json:"description"`
	Metadata       *string            `json:"metadata"`
	ConsumeActions *[]*ConsumeAction  `json:"consumeActions"`
	AcquireActions *[]*AcquireAction  `json:"acquireActions"`
}

type DeleteSalesItemMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	SalesItemName *string            `json:"salesItemName"`
}

type DescribeSalesItemGroupMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateSalesItemGroupMasterRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	Name           *string            `json:"name"`
	Description    *string            `json:"description"`
	Metadata       *string            `json:"metadata"`
	SalesItemNames *[]string          `json:"salesItemNames"`
}

type GetSalesItemGroupMasterRequest struct {
	RequestId          *core.RequestId    `json:"requestId"`
	ContextStack       *core.ContextStack `json:"contextStack"`
	NamespaceName      *string            `json:"namespaceName"`
	SalesItemGroupName *string            `json:"salesItemGroupName"`
}

type UpdateSalesItemGroupMasterRequest struct {
	RequestId          *core.RequestId    `json:"requestId"`
	ContextStack       *core.ContextStack `json:"contextStack"`
	NamespaceName      *string            `json:"namespaceName"`
	SalesItemGroupName *string            `json:"salesItemGroupName"`
	Description        *string            `json:"description"`
	Metadata           *string            `json:"metadata"`
	SalesItemNames     *[]string          `json:"salesItemNames"`
}

type DeleteSalesItemGroupMasterRequest struct {
	RequestId          *core.RequestId    `json:"requestId"`
	ContextStack       *core.ContextStack `json:"contextStack"`
	NamespaceName      *string            `json:"namespaceName"`
	SalesItemGroupName *string            `json:"salesItemGroupName"`
}

type DescribeShowcaseMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateShowcaseMasterRequest struct {
	RequestId          *core.RequestId       `json:"requestId"`
	ContextStack       *core.ContextStack    `json:"contextStack"`
	NamespaceName      *string               `json:"namespaceName"`
	Name               *string               `json:"name"`
	Description        *string               `json:"description"`
	Metadata           *string               `json:"metadata"`
	DisplayItems       *[]*DisplayItemMaster `json:"displayItems"`
	SalesPeriodEventId *string               `json:"salesPeriodEventId"`
}

type GetShowcaseMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	ShowcaseName  *string            `json:"showcaseName"`
}

type UpdateShowcaseMasterRequest struct {
	RequestId          *core.RequestId       `json:"requestId"`
	ContextStack       *core.ContextStack    `json:"contextStack"`
	NamespaceName      *string               `json:"namespaceName"`
	ShowcaseName       *string               `json:"showcaseName"`
	Description        *string               `json:"description"`
	Metadata           *string               `json:"metadata"`
	DisplayItems       *[]*DisplayItemMaster `json:"displayItems"`
	SalesPeriodEventId *string               `json:"salesPeriodEventId"`
}

type DeleteShowcaseMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	ShowcaseName  *string            `json:"showcaseName"`
}

type ExportMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetCurrentShowcaseMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateCurrentShowcaseMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Settings      *string            `json:"settings"`
}

type UpdateCurrentShowcaseMasterFromGitHubRequest struct {
	RequestId       *core.RequestId        `json:"requestId"`
	ContextStack    *core.ContextStack     `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

type DescribeShowcasesRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeShowcasesByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetShowcaseRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	ShowcaseName           *string            `json:"showcaseName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetShowcaseByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	ShowcaseName           *string            `json:"showcaseName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type BuyRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	ShowcaseName           *string            `json:"showcaseName"`
	DisplayItemId          *string            `json:"displayItemId"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type BuyByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	ShowcaseName           *string            `json:"showcaseName"`
	DisplayItemId          *string            `json:"displayItemId"`
	UserId                 *string            `json:"userId"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}
