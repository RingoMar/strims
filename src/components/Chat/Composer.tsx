import clsx from "clsx";
import filterObj from "filter-obj";
import Prism from "prismjs";
import React, { KeyboardEvent, useCallback, useEffect, useMemo, useState } from "react";
import {
  Descendant,
  Editor,
  Element,
  Node,
  NodeEntry,
  Path,
  Range,
  Text,
  Transforms,
  createEditor,
} from "slate";
import { withHistory } from "slate-history";
import { Editable, RenderLeafProps, Slate, withReact } from "slate-react";
import urlRegex from "url-regex-safe";

import Emote from "./Emote";

const tags = ["nsfw", "weeb", "nsfl", "loud"];
const commands = [
  "help",
  "emotes",
  "me",
  "message",
  "msg",
  "ignore",
  "unignore",
  "highlight",
  "unhighlight",
  "maxlines",
  "mute",
  "unmute",
  "subonly",
  "ban",
  "unban",
  "timestampformat",
  "tag",
  "untag",
  "exit",
  "hideemote",
  "unhideemote",
];

const initialValue: Descendant[] = [
  {
    type: "paragraph",
    children: [{ text: "" }],
  },
];

interface SelectedMatch {
  index: number;
  entry: SearchSourceEntry;
}

type Match = [SearchSource[], SearchSourceEntry[]];

const defaultSelectedMatch: SelectedMatch = { index: 0, entry: null };
const defaultMatch: Match = [[], []];

interface ComposerProps {
  onMessage: (message: string) => void;
  emotes: string[];
  modifiers: string[];
  nicks: string[];
  tags: string[];
}

const Composer: React.FC<ComposerProps> = ({ onMessage, emotes, modifiers, nicks, tags }) => {
  const [[matchSources, matchEntries], setMatch] = useState<Match>(defaultMatch);
  const [selectedMatch, setSelectedMatch] = useState<SelectedMatch>(defaultSelectedMatch);
  const [currentSearch, setSearch] = useState<SearchState | null>(null);
  const [lastSearch, setLastSearch] = useState<SearchState | null>(null);
  const search = lastSearch || currentSearch;

  const editor = useMemo(() => withReact(withNoLineBreaks(withHistory(createEditor()))), []);
  const [value, setValue] = useState<Descendant[]>(initialValue);

  const searchSources = useSearchSources(nicks, tags, commands, emotes, modifiers);

  useEffect(() => {
    if (!search) {
      setMatch(defaultMatch);
      setSelectedMatch(defaultSelectedMatch);
      return;
    }

    const query = search.query.toLowerCase();
    const test = {
      "prefix": ({ index }: SearchSourceEntry) => index.startsWith(query),
      "substring": ({ index }: SearchSourceEntry) => index.indexOf(query) !== -1,
    }[search.queryMode];

    const matches = search.sources
      .map((s) => ({
        ...s,
        entries: s.entries.filter(test),
      }))
      .filter(({ entries }) => entries.length > 0);
    const entries = matches.map(({ entries }) => entries).flat();

    setMatch([matches, entries]);
    setSelectedMatch({ index: 0, entry: entries[0] });
  }, [search]);

  const renderLeaf = useCallback((props: RenderLeafProps) => <Leaf {...props} />, []);

  const grammar = useMemo(
    () => getGrammar(emotes, modifiers, nicks, tags),
    [emotes, modifiers, nicks, tags]
  );

  const decorate = useCallback(
    ([node, path]: NodeEntry<Node>) =>
      Text.isText(node) ? getRanges(node.text, path, grammar) : [],
    [grammar]
  );

  const onChange = useCallback(
    (newValue) => {
      setValue(newValue);
      setSearch(getSearchState(editor, grammar, searchSources));
    },
    [editor, grammar, searchSources]
  );

  const onKeyDown = useCallback(
    (event: KeyboardEvent<HTMLDivElement>) => {
      if (event.key === "Enter") {
        event.preventDefault();
        onMessage(ComposerEditor.text(editor).trim());
        ComposerEditor.clear(editor);
        return;
      }

      if (!search) {
        return;
      }

      const getSelectedMatch = (i: number) => ({
        index: i,
        entry: matchEntries[i % matchEntries.length],
      });

      switch (event.key) {
        case "ArrowDown": {
          event.preventDefault();
          setSelectedMatch(({ index }) => getSelectedMatch(index + 1));
          return;
        }
        case "ArrowUp": {
          event.preventDefault();
          setSelectedMatch(({ index }) => getSelectedMatch(index - 1));
          return;
        }
        case "Tab": {
          event.preventDefault();

          if (!selectedMatch.entry) {
            return;
          }
          const entry = selectedMatch.entry;

          const target =
            entry.type === "modifier" ? currentSearch.modifierTarget : currentSearch.target;
          Transforms.select(editor, target);
          const whitespace = currentSearch.suffixSpace ? " " : "";
          Transforms.insertText(editor, (entry.substitution || entry.value) + whitespace);
          Transforms.move(editor);
          setLastSearch(search);
          setSelectedMatch(({ index }) => getSelectedMatch(index + 1));
          return;
        }
        case "Escape": {
          event.preventDefault();
          setSearch(null);
          return;
        }
      }

      setLastSearch(null);
      // setSelectedMatch(defaultSelectedMatch);
    },
    [matchEntries, selectedMatch, search, currentSearch]
  );

  const showSuggestions = search && matchSources.length > (lastSearch ? 1 : 0);

  // console.log(selectedMatch);
  // console.log({
  //   matchSources,
  //   matchEntries,
  //   selectedMatch,
  //   search,
  // });

  return (
    <div className="chat__composer">
      {showSuggestions && (
        <div className="chat__composer__autocomplete">
          <div className="chat__composer__autocomplete__list">
            {matchSources.map((m, i) => (
              <AutocompleteGroup {...m} selectedMatch={selectedMatch} key={i} />
            ))}
          </div>
        </div>
      )}
      <Slate editor={editor} value={value} onChange={onChange}>
        <Editable
          decorate={decorate}
          onKeyDown={onKeyDown}
          placeholder="Write a message..."
          renderLeaf={renderLeaf}
        />
      </Slate>
    </div>
  );
};

interface AutocompleteGroupProps {
  label: string;
  entries: SearchSourceEntry[];
  selectedMatch: SelectedMatch;
}

const AutocompleteGroup: React.FC<AutocompleteGroupProps> = ({ label, entries, selectedMatch }) => (
  <>
    <div className="chat__composer__autocomplete__label">{label}</div>
    {entries.map((e, i) => (
      <AutocompleteGroupItem entry={e} selectedMatch={selectedMatch} key={i} />
    ))}
  </>
);

interface AutocompleteGroupItemProps {
  entry: SearchSourceEntry;
  selectedMatch: SelectedMatch;
}

const AutocompleteGroupItem: React.FC<AutocompleteGroupItemProps> = ({ entry, selectedMatch }) => {
  let content: React.ReactNode;
  switch (entry.type) {
    case "emote":
      content = (
        <>
          <span className="chat__composer__autocomplete__item__emote">
            <Emote name={entry.value} />
          </span>
          <span className="chat__composer__autocomplete__item__label">{entry.value}</span>
        </>
      );
      break;
    default:
      content = <span className="chat__composer__autocomplete__item__label">{entry.value}</span>;
  }

  return (
    <div
      className={clsx(
        "chat__composer__autocomplete__item",
        `chat__composer__autocomplete__item--${entry.type}`,
        {
          "chat__composer__autocomplete__item--selected": entry === selectedMatch.entry,
        }
      )}
    >
      {content}
    </div>
  );
};

const Leaf: React.FC<RenderLeafProps> = ({ attributes, children, leaf }) => {
  if (leaf.emote) {
    const [name, ...modifiers] = leaf.text.split(":");

    return (
      <Emote {...attributes} name={name} modifiers={modifiers}>
        {children}
      </Emote>
    );
  }

  return (
    <span
      {...attributes}
      className={clsx({
        "chat__composer__span--code": leaf.code,
        "chat__composer__span--spoiler": leaf.spoiler,
        "chat__composer__span--url": leaf.url,
        "chat__composer__span--emote": leaf.emote,
        "chat__composer__span--tag": leaf.tag,
        "chat__composer__span--nick": leaf.nick,
        "chat__composer__span--self": leaf.self,
        "chat__composer__span--greentext": leaf.greentext,
      })}
      spellCheck={!leaf.emote && !leaf.tag && !leaf.nick}
    >
      {children}
    </span>
  );
};

export default Composer;

const noopPattern = /_^/;

const getGrammar = (emotes: string[], modifiers: string[], nicks: string[], tags: string[]) => {
  const nestableEntities = {
    code: {
      pattern: /`(\\`|[^`])*(`|$)/,
      greedy: true,
    },
    emote: {
      pattern: noopPattern,
      lookbehind: true,
    },
    nick: {
      pattern: noopPattern,
      lookbehind: true,
    },
    tag: {
      pattern: noopPattern,
      lookbehind: true,
    },
    url: urlRegex(),
  };

  if (emotes.length !== 0 && modifiers.length !== 0) {
    nestableEntities.emote.pattern = new RegExp(
      `(\\W|^)((${emotes.join("|")})(:(${modifiers.join("|")}))*)(?=\\W|$)`
    );
  } else if (emotes.length !== 0) {
    nestableEntities.emote.pattern = new RegExp(`(\\W|^)(${emotes.join("|")})(?=\\W|$)`);
  }
  if (nicks.length !== 0) {
    nestableEntities.nick.pattern = new RegExp(`(\\W|^)(${nicks.join("|")})(?=\\W|$)`);
  }
  if (tags.length !== 0) {
    nestableEntities.tag.pattern = new RegExp(`(\\W|^)(${tags.join("|")})(?=\\W|$)`);
  }

  const entities = {
    spoiler: {
      pattern: /\|\|(\\\||\|(?!\|)|[^|])*(\|\||$)/,
      inside: nestableEntities,
    },
    ...nestableEntities,
  };
  return {
    greentext: {
      pattern: /^>.*/,
      greedy: true,
      inside: entities,
    },
    self: {
      pattern: /^\/me .*/,
      greedy: true,
      inside: entities,
    },
    ...entities,
  };
};

const getRanges = (text: string, path: Path, grammar: Prism.Grammar) => {
  const ranges: Range[] = [];

  const appendRanges = (tokens: (string | Prism.Token)[], start: number = 0) =>
    tokens.reduce((offset, token) => {
      if (typeof token !== "string") {
        ranges.push({
          [token.type]: true,
          anchor: { path, offset: offset },
          focus: { path, offset: offset + token.length },
        });
        appendRanges(Array.isArray(token.content) ? token.content : [token.content], offset);
      }
      return offset + token.length;
    }, start);
  appendRanges(Prism.tokenize(text, grammar));

  return ranges;
};

interface SearchSourceEntry {
  type: "nick" | "tag" | "command" | "emote" | "modifier";
  value: string;
  substitution?: string;
  index: string;
}

interface SearchSource {
  label: string;
  entries: SearchSourceEntry[];
}

interface SearchSources {
  nicks: SearchSource;
  tags: SearchSource;
  commands: SearchSource;
  emotes: SearchSource;
  modifiers: SearchSource;
}

const useSearchSources = (
  nicks: string[],
  tags: string[],
  commands: string[],
  emotes: string[],
  modifiers: string[]
): SearchSources => {
  const sources: SearchSources = {
    nicks: useMemo(
      () => ({
        label: "members",
        entries: nicks.map((v) => ({
          type: "nick",
          value: v,
          index: v.toLowerCase(),
        })),
      }),
      [nicks]
    ),
    tags: useMemo(
      () => ({
        label: "tags",
        entries: tags.map((v) => ({
          type: "tag",
          value: v,
          index: v.toLowerCase(),
        })),
      }),
      [tags]
    ),
    commands: useMemo(
      () => ({
        label: "commands",
        entries: commands.map((v) => ({
          type: "command",
          value: v,
          substitution: "/" + v,
          index: v.toLowerCase(),
        })),
      }),
      [commands]
    ),
    emotes: useMemo(
      () => ({
        label: "emotes",
        entries: emotes.map((v) => ({
          type: "emote",
          value: v,
          index: v.toLowerCase(),
        })),
      }),
      [emotes]
    ),
    modifiers: useMemo(
      () => ({
        label: "modifiers",
        entries: modifiers.map((v) => ({
          type: "modifier",
          value: v,
          substitution: ":" + v,
          index: v.toLowerCase(),
        })),
      }),
      [modifiers]
    ),
  };
  return useMemo(() => sources, [nicks, tags, commands, emotes, modifiers]);
};

interface SearchState {
  debounceDelay: number;
  queryMode: "substring" | "prefix";
  sources: SearchSource[];
  query: string;
  target: Range;
  suffixSpace: boolean;
  modifierContext: string | undefined;
  modifierTarget: Range;
}

const getSearchState = (
  editor: Editor,
  grammar: Prism.Grammar,
  searchSources: SearchSources
): SearchState => {
  const { selection } = editor;
  if (!selection || !Range.isCollapsed(selection)) {
    return null;
  }

  const [node, path] = Editor.node(editor, selection);
  if (!Text.isText(node)) {
    return null;
  }

  const { text } = node;
  const { offset } = Range.start(selection);
  const [, contiguousContext, delta, prefix, punct, queryStart, suffixStart] =
    /(\w(?=:|@))?((\s*)(:|@|^\/)?(\w*)(\s*))$/.exec(text.substring(0, offset)) || [];
  const [, suffixEnd, queryEnd] = /^(\s*)(\w*)/.exec(text.substring(offset)) || [""];
  const hasSuffix = suffixStart || suffixEnd;
  const query = queryStart + (hasSuffix ? "" : queryEnd);

  const targetStart = { path, offset: offset - (delta.length - prefix.length) };
  const targetEnd = { path, offset: offset + (hasSuffix ? -suffixStart.length : queryEnd.length) };
  const target = Editor.range(editor, targetStart, targetEnd);

  const entityRanges = getRanges(text, path, filterObj(grammar, ["code", "emote", "url"]));

  const contextEnd = (prefix || punct) && { path, offset: offset - delta.length };
  const modifierContextRange =
    contextEnd && entityRanges.find((r) => r.emote && Range.includes(r, contextEnd));
  const modifierContext = modifierContextRange && Editor.string(editor, modifierContextRange);
  const modifierTarget = modifierContextRange && Editor.range(editor, contextEnd, targetEnd);

  const invalidContext =
    (contiguousContext && !modifierContext) ||
    entityRanges.some((r) => (r.code || r.url) && Range.includes(r, selection));

  const sources: SearchSource[] = [];
  if (punct === ":") {
    if (modifierContext) {
      sources.push(searchSources.modifiers);
    }
    if (!contiguousContext) {
      sources.push(searchSources.emotes);
    }
  } else if (punct === "@") {
    sources.push(searchSources.nicks);
  } else if (punct === "/") {
    sources.push(searchSources.commands);
  } else {
    if (modifierContext) {
      sources.push(searchSources.modifiers);
    }
    sources.push(searchSources.emotes, searchSources.nicks, searchSources.tags);
  }

  if (invalidContext || !(punct || query) || sources.length === 0) {
    return null;
  }

  return {
    debounceDelay: punct ? 0 : 100,
    queryMode: punct ? "substring" : "prefix",
    sources,
    query,
    target,
    suffixSpace: !hasSuffix,
    modifierContext,
    modifierTarget,
  };
};

const ComposerEditor = {
  clear: (editor: Editor) => {
    const [[node, path]] = Editor.nodes(editor, { match: (n) => Text.isText(n) });
    if (!Text.isText(node)) {
      return;
    }

    Transforms.select(
      editor,
      Editor.range(
        editor,
        {
          offset: 0,
          path,
        },
        {
          offset: node.text.length,
          path,
        }
      )
    );
    Transforms.delete(editor);
    Transforms.move(editor);
  },

  text: (editor: Editor) => {
    const [[node]] = Editor.nodes(editor, { match: (n) => Text.isText(n) });
    return Text.isText(node) ? node.text : "";
  },
};

const withNoLineBreaks = (editor: Editor) => {
  const { normalizeNode } = editor;

  editor.normalizeNode = (entry) => {
    const [node] = entry;

    if (Editor.isEditor(node) && node.children.length > 1) {
      Transforms.mergeNodes(editor, {
        match: (node) => Element.isElement(node) && node.type === "paragraph",
      });
    }

    normalizeNode(entry);
  };

  return editor;
};
