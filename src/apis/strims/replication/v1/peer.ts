import Reader from "@memelabs/protobuf/lib/pb/reader";
import Writer from "@memelabs/protobuf/lib/pb/writer";

import {
  strims_replication_v1_Event,
  strims_replication_v1_IEvent,
} from "./replication";
import {
  strims_dao_v1_VersionVector,
  strims_dao_v1_IVersionVector,
} from "../../dao/v1/dao";

export type IPeerOpenRequest = {
  version?: number;
  minCompatibleVersion?: number;
  timestamp?: bigint;
  key?: Uint8Array;
  signature?: Uint8Array;
}

export class PeerOpenRequest {
  version: number;
  minCompatibleVersion: number;
  timestamp: bigint;
  key: Uint8Array;
  signature: Uint8Array;

  constructor(v?: IPeerOpenRequest) {
    this.version = v?.version || 0;
    this.minCompatibleVersion = v?.minCompatibleVersion || 0;
    this.timestamp = v?.timestamp || BigInt(0);
    this.key = v?.key || new Uint8Array();
    this.signature = v?.signature || new Uint8Array();
  }

  static encode(m: PeerOpenRequest, w?: Writer): Writer {
    if (!w) w = new Writer();
    if (m.version) w.uint32(8).uint32(m.version);
    if (m.minCompatibleVersion) w.uint32(16).uint32(m.minCompatibleVersion);
    if (m.timestamp) w.uint32(24).int64(m.timestamp);
    if (m.key.length) w.uint32(80010).bytes(m.key);
    if (m.signature.length) w.uint32(80018).bytes(m.signature);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): PeerOpenRequest {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new PeerOpenRequest();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.version = r.uint32();
        break;
        case 2:
        m.minCompatibleVersion = r.uint32();
        break;
        case 3:
        m.timestamp = r.int64();
        break;
        case 10001:
        m.key = r.bytes();
        break;
        case 10002:
        m.signature = r.bytes();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export type IPeerOpenResponse = {
  checkpoint?: strims_dao_v1_IVersionVector;
}

export class PeerOpenResponse {
  checkpoint: strims_dao_v1_VersionVector | undefined;

  constructor(v?: IPeerOpenResponse) {
    this.checkpoint = v?.checkpoint && new strims_dao_v1_VersionVector(v.checkpoint);
  }

  static encode(m: PeerOpenResponse, w?: Writer): Writer {
    if (!w) w = new Writer();
    if (m.checkpoint) strims_dao_v1_VersionVector.encode(m.checkpoint, w.uint32(10).fork()).ldelim();
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): PeerOpenResponse {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new PeerOpenResponse();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.checkpoint = strims_dao_v1_VersionVector.decode(r, r.uint32());
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export type IPeerSendEventsRequest = {
  checkpoint?: strims_dao_v1_IVersionVector;
  events?: strims_replication_v1_IEvent[];
}

export class PeerSendEventsRequest {
  checkpoint: strims_dao_v1_VersionVector | undefined;
  events: strims_replication_v1_Event[];

  constructor(v?: IPeerSendEventsRequest) {
    this.checkpoint = v?.checkpoint && new strims_dao_v1_VersionVector(v.checkpoint);
    this.events = v?.events ? v.events.map(v => new strims_replication_v1_Event(v)) : [];
  }

  static encode(m: PeerSendEventsRequest, w?: Writer): Writer {
    if (!w) w = new Writer();
    if (m.checkpoint) strims_dao_v1_VersionVector.encode(m.checkpoint, w.uint32(10).fork()).ldelim();
    for (const v of m.events) strims_replication_v1_Event.encode(v, w.uint32(18).fork()).ldelim();
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): PeerSendEventsRequest {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new PeerSendEventsRequest();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.checkpoint = strims_dao_v1_VersionVector.decode(r, r.uint32());
        break;
        case 2:
        m.events.push(strims_replication_v1_Event.decode(r, r.uint32()));
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export type IPeerSendEventsResponse = {
  checkpoint?: strims_dao_v1_IVersionVector;
}

export class PeerSendEventsResponse {
  checkpoint: strims_dao_v1_VersionVector | undefined;

  constructor(v?: IPeerSendEventsResponse) {
    this.checkpoint = v?.checkpoint && new strims_dao_v1_VersionVector(v.checkpoint);
  }

  static encode(m: PeerSendEventsResponse, w?: Writer): Writer {
    if (!w) w = new Writer();
    if (m.checkpoint) strims_dao_v1_VersionVector.encode(m.checkpoint, w.uint32(10).fork()).ldelim();
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): PeerSendEventsResponse {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new PeerSendEventsResponse();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.checkpoint = strims_dao_v1_VersionVector.decode(r, r.uint32());
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

/* @internal */
export const strims_replication_v1_PeerOpenRequest = PeerOpenRequest;
/* @internal */
export type strims_replication_v1_PeerOpenRequest = PeerOpenRequest;
/* @internal */
export type strims_replication_v1_IPeerOpenRequest = IPeerOpenRequest;
/* @internal */
export const strims_replication_v1_PeerOpenResponse = PeerOpenResponse;
/* @internal */
export type strims_replication_v1_PeerOpenResponse = PeerOpenResponse;
/* @internal */
export type strims_replication_v1_IPeerOpenResponse = IPeerOpenResponse;
/* @internal */
export const strims_replication_v1_PeerSendEventsRequest = PeerSendEventsRequest;
/* @internal */
export type strims_replication_v1_PeerSendEventsRequest = PeerSendEventsRequest;
/* @internal */
export type strims_replication_v1_IPeerSendEventsRequest = IPeerSendEventsRequest;
/* @internal */
export const strims_replication_v1_PeerSendEventsResponse = PeerSendEventsResponse;
/* @internal */
export type strims_replication_v1_PeerSendEventsResponse = PeerSendEventsResponse;
/* @internal */
export type strims_replication_v1_IPeerSendEventsResponse = IPeerSendEventsResponse;
