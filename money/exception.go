package money

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

type ReceiptInvalid struct {
}

func (p ReceiptInvalid) Type() string {
	return "MoneyReceiptInvalid"
}

func (p ReceiptInvalid) Code() string {
	return "receipt.payload.invalid"
}

type Conflict struct {
}

func (p Conflict) Type() string {
	return "MoneyConflict"
}

func (p Conflict) Code() string {
	return "wallet.operation.conflict"
}

type Insufficient struct {
}

func (p Insufficient) Type() string {
	return "MoneyInsufficient"
}

func (p Insufficient) Code() string {
	return "wallet.balance.insufficient"
}
