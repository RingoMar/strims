import React, { Suspense, useEffect } from "react";
import { Scrollbars } from "react-custom-scrollbars-2";
import { Outlet } from "react-router";
import { useToggle } from "react-use";

import NetworkNav from "../../components/Layout/NetworkNav";
import { useLayout } from "../../contexts/Layout";
import LoadingPlaceholder from "../../root/LoadingPlaceholder";
import SwipablePanel, { DragState } from "../SwipablePanel";
import Chat from "./Chat";
import Player from "./Player";

export const LayoutBody: React.FC = ({ children }) => {
  const { showVideo, overlayState, setOverlayState } = useLayout();

  const handleDragStateChange = (state: DragState) => {
    setOverlayState({
      open: !state.closed,
      transitioning: state.transitioning,
    });
  };

  const [open, toggleOpen] = useToggle(false);
  useEffect(() => {
    if (!overlayState.transitioning) {
      toggleOpen(overlayState.open);
    }
  }, [overlayState]);

  return (
    <>
      <div className="layout__nav">
        <NetworkNav />
      </div>
      <main className="foo_1">
        <div className="content_panel">
          <Scrollbars autoHide>
            <div className="scroll_content_test">
              <Suspense fallback={<LoadingPlaceholder />}>
                <Outlet />
                {children}
              </Suspense>
            </div>
          </Scrollbars>
        </div>
        <SwipablePanel
          className="foo_2"
          direction="left"
          open={open}
          onDragStateChange={handleDragStateChange}
        >
          {showVideo && (
            <div className="layout__video">
              <Player />
            </div>
          )}
          <div className="layout__chat">
            <Chat />
          </div>
        </SwipablePanel>
      </main>
    </>
  );
};

export default LayoutBody;
