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

package auth

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type LoginRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserId *string	`json:"userId"`
	TimeOffset *int32	`json:"timeOffset"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type LoginBySignatureRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserId *string	`json:"userId"`
	KeyId *string	`json:"keyId"`
	Body *string	`json:"body"`
	Signature *string	`json:"signature"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}
