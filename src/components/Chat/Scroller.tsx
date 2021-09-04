import clsx from "clsx";
import React from "react";
import { Scrollbars } from "react-custom-scrollbars-2";
import {
  AutoSizer,
  CellMeasurer,
  CellMeasurerCache,
  Dimensions,
  Grid,
  List,
  ListRowRenderer,
  OnScrollParams,
} from "react-virtualized";

export interface MessageProps {
  style: React.CSSProperties;
  index: number;
}

interface ScrollerProps {
  renderMessage: (MessageProps) => React.ReactNode;
  messageCount: number;
  messageSizeCache: CellMeasurerCache;
  autoScrollThreshold?: number;
  resizeDebounceTimeout?: number;
}

const Scroller: React.FC<ScrollerProps> = (props) => {
  return (
    <AutoSizer>
      {(dimensions: Dimensions) => <ScrollerContent {...dimensions} {...props} />}
    </AutoSizer>
  );
};

interface ListInternal {
  Grid: Grid & {
    _scrollingContainer: HTMLElement;
    _onScroll: (e: React.UIEvent) => void;
  };
}

interface ScrollbarsInternal {
  view: HTMLElement;
}

const ScrollerContent: React.FC<ScrollerProps & Dimensions> = ({
  height,
  width,
  messageCount,
  renderMessage,
  messageSizeCache,
  autoScrollThreshold = 20,
  resizeDebounceTimeout = 100,
}) => {
  const list = React.useRef<List & ListInternal>();
  const scrollbars = React.useRef<Scrollbars & ScrollbarsInternal>();
  const [autoScroll, setAutoScroll] = React.useState(true);
  const [scrolling, setScrolling] = React.useState(true);
  const [resizing, setResizing] = React.useState(false);

  React.useEffect(() => {
    messageSizeCache.clearAll();
    list.current?.recomputeRowHeights();

    setResizing(true);
    const id = setTimeout(() => setResizing(false), resizeDebounceTimeout);
    return () => clearTimeout(id);
  }, [list, height, width]);

  React.useEffect(() => {
    if (autoScroll) {
      list.current?.scrollToRow(messageCount - 1);
    }
  }, [autoScroll, list, messageCount]);

  React.useEffect(() => {
    if (list.current && scrollbars.current) {
      list.current.Grid._scrollingContainer = scrollbars.current.view;
    }
  }, [list, scrollbars]);

  const handleScroll = React.useCallback((e) => list.current?.Grid?._onScroll(e), [list]);
  const handleScrollStart = React.useCallback(() => setScrolling(true), []);
  const handleScrollStop = React.useCallback(() => setScrolling(false), []);

  const handleListScroll = React.useCallback(
    ({ scrollHeight, scrollTop, clientHeight }: OnScrollParams) => {
      const thresholdExceeded = scrollHeight - scrollTop - clientHeight < autoScrollThreshold;
      const enabled = resizing ? autoScroll : thresholdExceeded;

      if (resizing && enabled) {
        list.current?.scrollToRow(messageCount - 1);
      }
      setAutoScroll(enabled);
    },
    [autoScroll, list, messageCount, resizing]
  );

  const renderRow: ListRowRenderer = React.useCallback(
    ({ index, key, style, parent }) => (
      <CellMeasurer
        cache={messageSizeCache}
        columnIndex={0}
        key={key}
        parent={parent}
        rowIndex={index}
        width={width}
      >
        {renderMessage({ index, style })}
      </CellMeasurer>
    ),
    [renderMessage, messageSizeCache, width]
  );

  const renderScrollThumb = React.useCallback(
    (props) => (
      <div
        {...props}
        className={clsx({
          "chat__scroller__scrollbar_handle": true,
          "chat__scroller__scrollbar_handle--scrolling": scrolling && !autoScroll,
        })}
      />
    ),
    [scrolling, autoScroll]
  );

  return (
    <Scrollbars
      ref={scrollbars}
      onScroll={handleScroll}
      onScrollStart={handleScrollStart}
      onScrollStop={handleScrollStop}
      style={{ height, width }}
      renderThumbVertical={renderScrollThumb}
    >
      <List
        ref={list}
        height={height}
        width={width}
        style={{ overflowX: "visible", overflowY: "visible" }}
        deferredMeasurementCache={messageSizeCache}
        rowHeight={messageSizeCache.rowHeight}
        rowCount={messageCount}
        rowRenderer={renderRow}
        onScroll={handleListScroll}
      />
    </Scrollbars>
  );
};

export default Scroller;
