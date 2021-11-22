/* eslint-disable no-console */

// import { Base64 } from "js-base64";
import React, { useContext, useEffect, useRef, useState } from "react";

import { EgressOpenStreamResponse } from "../apis/strims/video/v1/egress";
import { useClient, useLazyCall } from "../contexts/FrontendApi";
import { useLayout } from "../contexts/Layout";
import { PlayerContext, PlayerMode } from "../contexts/Player";
import { useProfile } from "../contexts/Profile";
import { useTheme } from "../contexts/Theme";
import { certificateRoot } from "../lib/certificate";
import * as fmp4 from "../lib/media/fmp4";
import * as mpegts from "../lib/media/mpegts";
import * as webm from "../lib/media/webm";

// import { CallChatClientRequest, OpenChatClientRequest, OpenChatServerRequest } from "../lib/pb";

interface PProfDownload {
  name: string;
  url: string;
}

const HomePage: React.FC = () => {
  const { colorScheme, setColorScheme } = useTheme();
  // const [{ profile }, { clearProfile }] = useProfile();

  const client = useClient();

  const [pprofData, pprof] = useLazyCall("debug", "pProf");
  const [pprofDownloads, setPProfDownloads] = useState([] as PProfDownload[]);

  const videoRef = useRef<HTMLVideoElement>();

  const addDownload = (name: string, data: Uint8Array) => {
    const f = new File([data], name, {
      type: "application/binary",
    });
    setPProfDownloads((prev) => [...prev, { name, url: URL.createObjectURL(f) }]);
  };

  useEffect(() => {
    if (!pprofData.value) {
      return;
    }
    addDownload(`${pprofData.value.name}.profile`, pprofData.value.data);
  }, [pprofData.value]);

  const handleColorToggle = () => setColorScheme(colorScheme === "dark" ? "light" : "dark");

  const [broadcastState, setBroadcastState] =
    useState<{ id: Uint8Array; mediaStream: MediaStream }>(null);
  const handleStartBroadcastClick = async () => {
    const { networks } = await client.network.list();
    const [{ id }, mediaStream] = await Promise.all([
      client.videoCapture.open({
        mimeType: webm.MIME_TYPE,
        networkKeys: networks.map((n) => certificateRoot(n.certificate).key),
        directorySnippet: {
          title: "test broadcast",
          description: "broadcast from getDisplayMedia",
        },
      }),
      // https://github.com/microsoft/TypeScript/issues/33232
      // eslint-disable-next-line
      (navigator.mediaDevices as any).getDisplayMedia({
        video: true,
        audio: {
          autoGainControl: false,
          echoCancellation: false,
          googAutoGainControl: false,
          noiseSuppression: false,
        },
        frameRate: 30,
      }) as Promise<MediaStream>,
    ]);

    setBroadcastState({ id, mediaStream });
  };

  useEffect(() => {
    if (!videoRef.current || !broadcastState) {
      return;
    }

    const { id, mediaStream } = broadcastState;
    const encoder = new webm.Encoder(mediaStream);

    videoRef.current.srcObject = mediaStream;

    encoder.ondata = (data) => client.videoCapture.append({ id, data });
    encoder.onend = (data) => client.videoCapture.append({ id, data, segmentEnd: true });

    return () => {
      encoder.stop();
      void client.videoCapture.close({ id });
      mediaStream.getTracks().forEach((t) => t.stop());
    };
  }, [videoRef.current, broadcastState]);

  // const handleTestClick = async () => {
  //   console.log("starting vpn");
  //   const networkEvents = client.network.startVPN();

  //   networkEvents.on("data", (e) => {
  //     console.log(e);
  //   });

  //   console.log("vpn started");

  //   console.log("waiting for networks...");
  //   await new Promise((resolve) => setTimeout(resolve, 2000));

  //   // console.log("starting swarm");
  //   // const swarm = await client.openVideoServer();
  //   // console.log("started swarm", swarm);

  //   // console.log("publishing swarm");
  //   // await client.publishSwarm({
  //   //   id: swarm.id,
  //   // });
  //   // console.log("swarm published");
  // };

  // const handleChatClientClick = async () => {
  //   const listServersRes = await client.chat.listServers();
  //   const server = listServersRes.chatServers[0];
  //   const chatEvents = client.chat.openClient(
  //     new OpenChatClientRequest({
  //       networkKey: server.networkKey,
  //       serverKey: server.key.public,
  //     })
  //   );
  //   chatEvents.on("data", async (v) => {
  //     let n = 0;

  //     switch (v.body) {
  //       case "open":
  //         console.log("chat client >>>", v.open.clientId);

  //         await new Promise((resolve) => setTimeout(resolve, 2000));

  //         setInterval(() => {
  //           client.chat.callClient(
  //             new CallChatClientRequest({
  //               clientId: v.open.clientId,
  //               message: new CallChatClientRequest.Message({
  //                 time: Date.now(),
  //                 body: `${n++}`,
  //               }),
  //             })
  //           );
  //         }, 500);
  //         return;
  //       case "close":
  //         console.log("chat client closed");
  //         return;
  //       case "message":
  //         console.log("chat message", v.message);
  //         console.log(
  //           "delay - e2e:",
  //           Date.now() - v.message.sentTime,
  //           "client>server:",
  //           v.message.serverTime - v.message.sentTime,
  //           "server>ui:",
  //           Date.now() - v.message.serverTime,
  //           "value:",
  //           v.message.body
  //         );
  //         return;
  //       default:
  //         console.log(v);
  //     }
  //   });
  // };

  // const handleChatServerClick = async () => {
  //   const servers = await client.chat.listServers();
  //   if (servers.chatServers.length === 0) {
  //     return;
  //   }

  //   const server = servers.chatServers[0];
  //   const chatEvents = client.chat.openServer(new OpenChatServerRequest({ server }));
  //   chatEvents.on("data", (v) => {
  //     switch (v.body) {
  //       case "open":
  //         console.log("chat server >>>", v.open.serverId);
  //         return;
  //       case "close":
  //         console.log("chat server closed");
  //         return;
  //       default:
  //         console.log(v);
  //     }
  //   });
  // };

  const handleReadMetricsClick = async () => {
    const { data } = await client.debug.readMetrics({ format: 0 });
    console.log(new TextDecoder().decode(data));
  };

  // useEffect(() => {
  //   handleTestClick();
  // }, []);

  const { setMode } = useContext(PlayerContext);

  return (
    <>
      <header className="home_page__subheader"></header>
      <section className="home_page__main__video">
        <div>
          <button className="input input_button" onClick={handleColorToggle}>
            toggle theme
          </button>
          <button className="input input_button" onClick={handleStartBroadcastClick}>
            start broadcast
          </button>
          {/* <button
              className="input input_button"
              onClick={() => handleViewBroadcastClick(new webm.Decoder())}
            >
              view broadcast
            </button> */}
          {/* <button
              className="input input_button"
              onClick={() => handleViewBroadcastClick(new fmp4.Decoder())}
            >
              view rtmp broadcast
            </button> */}
          <button className="input input_button" onClick={() => pprof({ name: "allocs" })}>
            allocs profile
          </button>
          <button className="input input_button" onClick={() => pprof({ name: "goroutine" })}>
            goroutine profile
          </button>
          <button className="input input_button" onClick={() => pprof({ name: "heap" })}>
            heap profile
          </button>
          {/* <button className="input input_button" onClick={handleTestClick}>
              test
            </button> */}
          {/* <button className="input input_button" onClick={handleChatClientClick}>
              chat client
            </button>
            <button className="input input_button" onClick={handleChatServerClick}>
              chat server
            </button> */}
          <button className="input input_button" onClick={handleReadMetricsClick}>
            read metrics
          </button>
          <button className="input input_button" onClick={() => setMode(PlayerMode.LARGE)}>
            large
          </button>
          <button className="input input_button" onClick={() => setMode(PlayerMode.PIP)}>
            pip
          </button>
          <button className="input input_button" onClick={() => setMode(PlayerMode.CLOSED)}>
            closed
          </button>
        </div>
        <div>
          {pprofDownloads.map(({ name, url }, i) => (
            <a href={url} download={name} key={i}>
              {name}
            </a>
          ))}
        </div>
        <video className="home_page__main__video__video" controls autoPlay ref={videoRef} />
        <div style={{ height: "1000px", width: "50px" }} />
      </section>
    </>
  );
};

export default HomePage;
