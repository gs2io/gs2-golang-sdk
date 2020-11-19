/**
 * Copyright 2016 Game Server Services, Inc. or its affiliates. All Rights
 * Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */
package core

import "encoding/json"

type RequestError struct {
	Component *string
	Message   *string
}

type Gs2Exception interface {
	RequestErrors() *[]RequestError
}

type BadRequestException struct {
	Errors []RequestError
}

func (p BadRequestException) RequestErrors() *[]RequestError {
	return &p.Errors
}

func (p BadRequestException) Error() string {
	Message, _ := json.Marshal(p.Errors)
	return string(Message)
}

type BadGatewayException struct {
	Errors []RequestError
}

func (p BadGatewayException) RequestErrors() *[]RequestError {
	return &p.Errors
}

func (p BadGatewayException) Error() string {
	Message, _ := json.Marshal(p.Errors)
	return string(Message)
}

type ConflictException struct {
	Errors []RequestError
}

func (p ConflictException) RequestErrors() *[]RequestError {
	return &p.Errors
}

func (p ConflictException) Error() string {
	Message, _ := json.Marshal(p.Errors)
	return string(Message)
}

type InternalServerErrorException struct {
	Errors []RequestError
}

func (p InternalServerErrorException) RequestErrors() *[]RequestError {
	return &p.Errors
}

func (p InternalServerErrorException) Error() string {
	Message, _ := json.Marshal(p.Errors)
	return string(Message)
}

type NotFoundException struct {
	Errors []RequestError
}

func (p NotFoundException) RequestErrors() *[]RequestError {
	return &p.Errors
}

func (p NotFoundException) Error() string {
	Message, _ := json.Marshal(p.Errors)
	return string(Message)
}

type QuotaExceedException struct {
	Errors []RequestError
}

func (p QuotaExceedException) RequestErrors() *[]RequestError {
	return &p.Errors
}

func (p QuotaExceedException) Error() string {
	Message, _ := json.Marshal(p.Errors)
	return string(Message)
}

type RequestTimeoutException struct {
	Errors []RequestError
}

func (p RequestTimeoutException) RequestErrors() *[]RequestError {
	return &p.Errors
}

func (p RequestTimeoutException) Error() string {
	Message, _ := json.Marshal(p.Errors)
	return string(Message)
}

type ServiceUnavailableException struct {
	Errors []RequestError
}

func (p ServiceUnavailableException) RequestErrors() *[]RequestError {
	return &p.Errors
}

func (p ServiceUnavailableException) Error() string {
	Message, _ := json.Marshal(p.Errors)
	return string(Message)
}

type UnauthorizedException struct {
	Errors []RequestError
}

func (p UnauthorizedException) RequestErrors() *[]RequestError {
	return &p.Errors
}

func (p UnauthorizedException) Error() string {
	Message, _ := json.Marshal(p.Errors)
	return string(Message)
}

type UnknownException struct {
	Errors []RequestError
}

func (p UnknownException) RequestErrors() *[]RequestError {
	return &p.Errors
}

func (p UnknownException) Error() string {
	Message, _ := json.Marshal(p.Errors)
	return string(Message)
}

type NoInternetConnectionException struct {
	Errors []RequestError
}

func (p NoInternetConnectionException) RequestErrors() *[]RequestError {
	return &p.Errors
}

func (p NoInternetConnectionException) Error() string {
	Message, _ := json.Marshal(p.Errors)
	return string(Message)
}
