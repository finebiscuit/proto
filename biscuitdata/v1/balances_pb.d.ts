// package: biscuitdata.v1
// file: v1/balances.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class BalanceData extends jspb.Message { 
    getCreatedAt(): number;
    setCreatedAt(value: number): BalanceData;
    getUpdatedAt(): number;
    setUpdatedAt(value: number): BalanceData;
    getDisplayName(): string;
    setDisplayName(value: string): BalanceData;
    getAccountName(): string;
    setAccountName(value: string): BalanceData;
    getInstitutionName(): string;
    setInstitutionName(value: string): BalanceData;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BalanceData.AsObject;
    static toObject(includeInstance: boolean, msg: BalanceData): BalanceData.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BalanceData, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BalanceData;
    static deserializeBinaryFromReader(message: BalanceData, reader: jspb.BinaryReader): BalanceData;
}

export namespace BalanceData {
    export type AsObject = {
        createdAt: number,
        updatedAt: number,
        displayName: string,
        accountName: string,
        institutionName: string,
    }
}

export class EntryData extends jspb.Message { 
    getCreatedAt(): number;
    setCreatedAt(value: number): EntryData;
    getUpdatedAt(): number;
    setUpdatedAt(value: number): EntryData;

    getValuesMap(): jspb.Map<string, string>;
    clearValuesMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EntryData.AsObject;
    static toObject(includeInstance: boolean, msg: EntryData): EntryData.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EntryData, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EntryData;
    static deserializeBinaryFromReader(message: EntryData, reader: jspb.BinaryReader): EntryData;
}

export namespace EntryData {
    export type AsObject = {
        createdAt: number,
        updatedAt: number,

        valuesMap: Array<[string, string]>,
    }
}
