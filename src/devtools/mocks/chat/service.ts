import { PassThrough } from "stream";

import { Readable } from "@memelabs/protobuf/lib/rpc/stream";

import * as chatv1 from "../../../apis/strims/chat/v1/chat";
import assetBundle from "../../mocks/chat/assetBundle";
import MessageEmitter from "../../mocks/chat/MessageEmitter";

export default class ChatService {
  messages: Readable<chatv1.Message>;

  constructor(messages = new MessageEmitter(1000)) {
    console.log(messages);
    this.messages = messages;
  }

  destroy(): void {
    this.messages.destroy();
  }

  openClient(): Readable<chatv1.OpenClientResponse> {
    const ch = new PassThrough({ objectMode: true });

    window.setTimeout(() =>
      ch.push(
        new chatv1.OpenClientResponse({
          body: new chatv1.OpenClientResponse.Body({
            open: new chatv1.OpenClientResponse.Open({
              clientId: BigInt(1),
            }),
          }),
        })
      )
    );

    void assetBundle().then((assetBundle) =>
      ch.push(
        new chatv1.OpenClientResponse({
          body: new chatv1.OpenClientResponse.Body({ assetBundle }),
        })
      )
    );

    this.messages.on("data", (message) =>
      ch.push(
        new chatv1.OpenClientResponse({
          body: new chatv1.OpenClientResponse.Body({ message }),
        })
      )
    );

    return ch;
  }
  clientSendMessage(
    req: chatv1.ClientSendMessageRequest
  ): Promise<chatv1.ClientSendMessageResponse> {
    this.messages.push(
      new chatv1.Message({
        nick: "test_user",
        serverTime: BigInt(Date.now()),
        body: req.body,
        entities: new chatv1.Message.Entities(),
      })
    );
    return Promise.resolve(new chatv1.ClientSendMessageResponse());
  }

  setUIConfig(): Promise<chatv1.SetUIConfigResponse> {
    return Promise.resolve(new chatv1.SetUIConfigResponse());
  }

  getUIConfig(): Promise<chatv1.GetUIConfigResponse> {
    return Promise.resolve(new chatv1.GetUIConfigResponse());
  }
}