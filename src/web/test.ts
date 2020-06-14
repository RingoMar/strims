import Worker from "worker-loader!./svc.worker";

import { Bus, WindowBridge } from "../lib/bridge";
import { WSReadWriter } from "../lib/ws";

class Success {}

(async () => {
  const bridge = new WindowBridge(Worker as any);
  const bus = await new Promise<Bus>((resolve) => {
    bridge.once("busopen:default", (b: any) => resolve(b));
  });

  const ws = new WSReadWriter(`wss://${location.host}/manage`);

  bus.on("data", (d: Uint8Array) => ws.write(d));
  ws.on("data", (d: Uint8Array) => bus.write(d));

  ws.on("close", () => {
    throw new Success();
  });
})();