/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
export type GetRateRequest = {
  from?: string
  to?: string
}

export type GetRateResponse = {
  value?: string
}

export class Forex {
  static GetRate(req: GetRateRequest, initReq?: fm.InitReq): Promise<GetRateResponse> {
    return fm.fetchReq<GetRateRequest, GetRateResponse>(`/v1/forex-rates/${req["from"]}/${req["to"]}?${fm.renderURLSearchParams(req, ["from", "to"])}`, {...initReq, method: "GET"})
  }
}