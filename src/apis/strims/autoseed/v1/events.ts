import Reader from "@memelabs/protobuf/lib/pb/reader";
import Writer from "@memelabs/protobuf/lib/pb/writer";

import {
  Config as strims_autoseed_v1_Config,
  IConfig as strims_autoseed_v1_IConfig,
  Rule as strims_autoseed_v1_Rule,
  IRule as strims_autoseed_v1_IRule,
} from "./autoseed";

export type IConfigChangeEvent = {
  config?: strims_autoseed_v1_IConfig;
}

export class ConfigChangeEvent {
  config: strims_autoseed_v1_Config | undefined;

  constructor(v?: IConfigChangeEvent) {
    this.config = v?.config && new strims_autoseed_v1_Config(v.config);
  }

  static encode(m: ConfigChangeEvent, w?: Writer): Writer {
    if (!w) w = new Writer();
    if (m.config) strims_autoseed_v1_Config.encode(m.config, w.uint32(10).fork()).ldelim();
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): ConfigChangeEvent {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new ConfigChangeEvent();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.config = strims_autoseed_v1_Config.decode(r, r.uint32());
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export type IRuleChangeEvent = {
  rule?: strims_autoseed_v1_IRule;
}

export class RuleChangeEvent {
  rule: strims_autoseed_v1_Rule | undefined;

  constructor(v?: IRuleChangeEvent) {
    this.rule = v?.rule && new strims_autoseed_v1_Rule(v.rule);
  }

  static encode(m: RuleChangeEvent, w?: Writer): Writer {
    if (!w) w = new Writer();
    if (m.rule) strims_autoseed_v1_Rule.encode(m.rule, w.uint32(10).fork()).ldelim();
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): RuleChangeEvent {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new RuleChangeEvent();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.rule = strims_autoseed_v1_Rule.decode(r, r.uint32());
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export type IRuleDeleteEvent = {
  rule?: strims_autoseed_v1_IRule;
}

export class RuleDeleteEvent {
  rule: strims_autoseed_v1_Rule | undefined;

  constructor(v?: IRuleDeleteEvent) {
    this.rule = v?.rule && new strims_autoseed_v1_Rule(v.rule);
  }

  static encode(m: RuleDeleteEvent, w?: Writer): Writer {
    if (!w) w = new Writer();
    if (m.rule) strims_autoseed_v1_Rule.encode(m.rule, w.uint32(10).fork()).ldelim();
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): RuleDeleteEvent {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new RuleDeleteEvent();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.rule = strims_autoseed_v1_Rule.decode(r, r.uint32());
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}
