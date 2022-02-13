/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
export type RegisterRequest = {
  email?: string
  password?: string
}

export type LoginRequest = {
  email?: string
  password?: string
}

export type AccessTokenResponse = {
  tokenType?: string
  accessToken?: string
  expiresIn?: number
  refreshToken?: string
  scope?: string
}

export class Auth {
  static Register(req: RegisterRequest, initReq?: fm.InitReq): Promise<AccessTokenResponse> {
    return fm.fetchReq<RegisterRequest, AccessTokenResponse>(`/v1/auth/register`, {...initReq, method: "POST"})
  }
  static Login(req: LoginRequest, initReq?: fm.InitReq): Promise<AccessTokenResponse> {
    return fm.fetchReq<LoginRequest, AccessTokenResponse>(`/v1/auth/login`, {...initReq, method: "POST"})
  }
}