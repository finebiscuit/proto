/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
export type EncryptedData = {
  versionHash?: string
  keyId?: string
  algo?: number
  base64Value?: string
}

export type Balance = {
  id?: string
  encData?: EncryptedData
  ownerId?: string
  typeId?: string
  currencyId?: string
  currentEntry?: Entry
}

export type Entry = {
  id?: string
  encData?: EncryptedData
  balanceId?: string
}

export type ListBalancesRequest = {
}

export type ListBalancesResponse = {
  balances?: Balance[]
}

export type GetBalanceRequest = {
  balanceId?: string
}

export type GetBalanceResponse = {
  balance?: Balance
}

export type CreateBalanceRequest = {
  balance?: Balance
}

export type CreateBalanceResponse = {
  balance?: Balance
}

export class Accounting {
  static ListBalances(req: ListBalancesRequest, initReq?: fm.InitReq): Promise<ListBalancesResponse> {
    return fm.fetchReq<ListBalancesRequest, ListBalancesResponse>(`/v1/balances?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetBalance(req: GetBalanceRequest, initReq?: fm.InitReq): Promise<GetBalanceResponse> {
    return fm.fetchReq<GetBalanceRequest, GetBalanceResponse>(`/v1/balances/${req["balanceId"]}?${fm.renderURLSearchParams(req, ["balanceId"])}`, {...initReq, method: "GET"})
  }
  static CreateBalance(req: CreateBalanceRequest, initReq?: fm.InitReq): Promise<CreateBalanceResponse> {
    return fm.fetchReq<CreateBalanceRequest, CreateBalanceResponse>(`/v1/balances`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}