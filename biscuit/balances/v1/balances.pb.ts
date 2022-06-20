/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
export type Payload = {
  version?: string
  keyId?: string
  scheme?: number
  base64Value?: string
}

export type Balance = {
  id?: string
  payload?: Payload
  typeId?: string
  currencyId?: string
  currentEntry?: Entry
}

export type Entry = {
  ymd?: string
  payload?: Payload
}

export type ListRequest = {
}

export type ListResponse = {
  balances?: Balance[]
}

export type GetRequest = {
  balanceId?: string
}

export type GetResponse = {
  balance?: Balance
}

export type CreateRequest = {
  typeId?: string
  currencyId?: string
  balancePayload?: Payload
  entryYmd?: string
  entryPayload?: Payload
}

export type CreateResponse = {
  balance?: Balance
}

export type UpsertEntryRequest = {
  balanceId?: string
  entryYmd?: string
  entryPayload?: Payload
  versionMatch?: string
}

export type UpsertEntryResponse = {
  entry?: Entry
}

export class Balances {
  static List(req: ListRequest, initReq?: fm.InitReq): Promise<ListResponse> {
    return fm.fetchReq<ListRequest, ListResponse>(`/v1/balances?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static Get(req: GetRequest, initReq?: fm.InitReq): Promise<GetResponse> {
    return fm.fetchReq<GetRequest, GetResponse>(`/v1/balances/${req["balanceId"]}?${fm.renderURLSearchParams(req, ["balanceId"])}`, {...initReq, method: "GET"})
  }
  static Create(req: CreateRequest, initReq?: fm.InitReq): Promise<CreateResponse> {
    return fm.fetchReq<CreateRequest, CreateResponse>(`/v1/balances`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpsertEntry(req: UpsertEntryRequest, initReq?: fm.InitReq): Promise<UpsertEntryResponse> {
    return fm.fetchReq<UpsertEntryRequest, UpsertEntryResponse>(`/v1/balances/${req["balanceId"]}/entries/${req["entryYmd"]}`, {...initReq, method: "PUT", body: JSON.stringify(req)})
  }
}