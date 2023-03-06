package core

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

import "encoding/json"

type ClientError interface {
    Type() string
    Code() string
}

type RequestError struct {
    Component *string `json:"component"`
    Message   *string `json:"message"`
    Code      *string `json:"code"`
}

type Gs2Exception interface {
    RequestErrors() []RequestError
    SetClientError(err ClientError) Gs2Exception
    ClientError() *ClientError
    Error() string
}

type BadRequestException struct {
    Errors    []RequestError `json:"errors"`
    clientErr *ClientError
}

func (p BadRequestException) RequestErrors() []RequestError {
    return p.Errors
}

func (p BadRequestException) SetClientError(err ClientError) Gs2Exception {
    p.clientErr = &err
    return p
}

func (p BadRequestException) ClientError() *ClientError {
    return p.clientErr
}

func (p BadRequestException) Error() string {
    Message, _ := json.Marshal(p.Errors)
    return string(Message)
}

type BadGatewayException struct {
    Errors    []RequestError
    clientErr *ClientError
}

func (p BadGatewayException) RequestErrors() []RequestError {
    return p.Errors
}

func (p BadGatewayException) Error() string {
    Message, _ := json.Marshal(p.Errors)
    return string(Message)
}

func (p BadGatewayException) SetClientError(err ClientError) Gs2Exception {
    p.clientErr = &err
    return p
}

func (p BadGatewayException) ClientError() *ClientError {
    return p.clientErr
}

type ConflictException struct {
    Errors    []RequestError
    clientErr *ClientError
}

func (p ConflictException) RequestErrors() []RequestError {
    return p.Errors
}

func (p ConflictException) Error() string {
    Message, _ := json.Marshal(p.Errors)
    return string(Message)
}

func (p ConflictException) SetClientError(err ClientError) Gs2Exception {
    p.clientErr = &err
    return p
}

func (p ConflictException) ClientError() *ClientError {
    return p.clientErr
}

type InternalServerErrorException struct {
    Errors    []RequestError
    clientErr *ClientError
}

func (p InternalServerErrorException) RequestErrors() []RequestError {
    return p.Errors
}

func (p InternalServerErrorException) Error() string {
    Message, _ := json.Marshal(p.Errors)
    return string(Message)
}

func (p InternalServerErrorException) SetClientError(err ClientError) Gs2Exception {
    p.clientErr = &err
    return p
}

func (p InternalServerErrorException) ClientError() *ClientError {
    return p.clientErr
}

type NotFoundException struct {
    Errors    []RequestError
    clientErr *ClientError
}

func (p NotFoundException) RequestErrors() []RequestError {
    return p.Errors
}

func (p NotFoundException) Error() string {
    Message, _ := json.Marshal(p.Errors)
    return string(Message)
}

func (p NotFoundException) SetClientError(err ClientError) Gs2Exception {
    p.clientErr = &err
    return p
}

func (p NotFoundException) ClientError() *ClientError {
    return p.clientErr
}

type QuotaExceedException struct {
    Errors    []RequestError
    clientErr *ClientError
}

func (p QuotaExceedException) RequestErrors() []RequestError {
    return p.Errors
}

func (p QuotaExceedException) Error() string {
    Message, _ := json.Marshal(p.Errors)
    return string(Message)
}

func (p QuotaExceedException) SetClientError(err ClientError) Gs2Exception {
    p.clientErr = &err
    return p
}

func (p QuotaExceedException) ClientError() *ClientError {
    return p.clientErr
}

type RequestTimeoutException struct {
    Errors    []RequestError
    clientErr *ClientError
}

func (p RequestTimeoutException) RequestErrors() []RequestError {
    return p.Errors
}

func (p RequestTimeoutException) Error() string {
    Message, _ := json.Marshal(p.Errors)
    return string(Message)
}

func (p RequestTimeoutException) SetClientError(err ClientError) Gs2Exception {
    p.clientErr = &err
    return p
}

func (p RequestTimeoutException) ClientError() *ClientError {
    return p.clientErr
}

type ServiceUnavailableException struct {
    Errors    []RequestError
    clientErr *ClientError
}

func (p ServiceUnavailableException) RequestErrors() []RequestError {
    return p.Errors
}

func (p ServiceUnavailableException) Error() string {
    Message, _ := json.Marshal(p.Errors)
    return string(Message)
}

func (p ServiceUnavailableException) SetClientError(err ClientError) Gs2Exception {
    p.clientErr = &err
    return p
}

func (p ServiceUnavailableException) ClientError() *ClientError {
    return p.clientErr
}

type UnauthorizedException struct {
    Errors    []RequestError
    clientErr *ClientError
}

func (p UnauthorizedException) RequestErrors() []RequestError {
    return p.Errors
}

func (p UnauthorizedException) Error() string {
    Message, _ := json.Marshal(p.Errors)
    return string(Message)
}

func (p UnauthorizedException) SetClientError(err ClientError) Gs2Exception {
    p.clientErr = &err
    return p
}

func (p UnauthorizedException) ClientError() *ClientError {
    return p.clientErr
}

type UnknownException struct {
    Errors    []RequestError
    clientErr *ClientError
}

func (p UnknownException) RequestErrors() []RequestError {
    return p.Errors
}

func (p UnknownException) Error() string {
    Message, _ := json.Marshal(p.Errors)
    return string(Message)
}

func (p UnknownException) SetClientError(err ClientError) Gs2Exception {
    p.clientErr = &err
    return p
}

func (p UnknownException) ClientError() *ClientError {
    return p.clientErr
}

type NoInternetConnectionException struct {
    Errors    []RequestError
    clientErr *ClientError
}

func (p NoInternetConnectionException) RequestErrors() []RequestError {
    return p.Errors
}

func (p NoInternetConnectionException) Error() string {
    Message, _ := json.Marshal(p.Errors)
    return string(Message)
}

func (p NoInternetConnectionException) SetClientError(err ClientError) Gs2Exception {
    p.clientErr = &err
    return p
}

func (p NoInternetConnectionException) ClientError() *ClientError {
    return p.clientErr
}
