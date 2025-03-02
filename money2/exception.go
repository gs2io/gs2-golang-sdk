package money2

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
	return "Money2ReceiptInvalid"
}

func (p ReceiptInvalid) Code() string {
	return "receipt.payload.invalid"
}

type Conflict struct {
}

func (p Conflict) Type() string {
	return "Money2Conflict"
}

func (p Conflict) Code() string {
	return "wallet.operation.conflict"
}

type Insufficient struct {
}

func (p Insufficient) Type() string {
	return "Money2Insufficient"
}

func (p Insufficient) Code() string {
	return "wallet.balance.insufficient"
}

type AlreadyUsed struct {
}

func (p AlreadyUsed) Type() string {
	return "Money2AlreadyUsed"
}

func (p AlreadyUsed) Code() string {
	return "subscription.transaction.used"
}

type MyselfSubscription struct {
}

func (p MyselfSubscription) Type() string {
	return "Money2MyselfSubscription"
}

func (p MyselfSubscription) Code() string {
	return "subscription.transaction.myself"
}

type LockPeriodNotElapsed struct {
}

func (p LockPeriodNotElapsed) Type() string {
	return "Money2LockPeriodNotElapsed"
}

func (p LockPeriodNotElapsed) Code() string {
	return "subscription.transaction.used"
}
