import React from "react";
import { Link, Navigate, useParams } from "react-router-dom";

import { Tag } from "../../../apis/strims/chat/v1/chat";
import {
  MenuCell,
  MenuItem,
  MenuLink,
  Table,
  TableMenu,
  TableTitleBar,
} from "../../../components/Settings/Table";
import { useCall, useLazyCall } from "../../../contexts/FrontendApi";

export interface ChatTagTableProps {
  serverId: bigint;
  tags: Tag[];
  onDelete: () => void;
}

const ChatTagTable: React.FC<ChatTagTableProps> = ({ serverId, tags, onDelete }) => {
  const [, deleteChatTag] = useLazyCall("chatServer", "deleteTag", { onComplete: onDelete });

  if (!tags) {
    return null;
  }

  const rows = tags.map((modifier) => {
    const handleDelete = () => deleteChatTag({ serverId, id: modifier.id });

    return (
      <tr key={modifier.id.toString()}>
        <td>
          <Link to={`/settings/chat-servers/${serverId}/tags/${modifier.id}`}>{modifier.name}</Link>
        </td>
        <MenuCell>
          <MenuItem label="Delete" onClick={handleDelete} />
        </MenuCell>
      </tr>
    );
  });
  return (
    <Table>
      <thead>
        <tr>
          <th>Title</th>
          <th></th>
        </tr>
      </thead>
      <tbody>{rows}</tbody>
    </Table>
  );
};

const ChatTagList: React.FC = () => {
  const { serverId } = useParams<"serverId">();
  const [{ loading, value }, getTags] = useCall("chatServer", "listTags", {
    args: [{ serverId: BigInt(serverId) }],
  });

  if (loading) {
    return null;
  }
  if (!value?.tags.length) {
    return <Navigate to={`/settings/chat-servers/${serverId}/tags/new`} />;
  }
  return (
    <>
      <TableTitleBar label="Emote Tags" backLink={`/settings/chat-servers/${serverId}`}>
        <TableMenu label="Create">
          <MenuLink label="Create Tag" to={`/settings/chat-servers/${serverId}/tags/new`} />
        </TableMenu>
      </TableTitleBar>
      <ChatTagTable serverId={BigInt(serverId)} tags={value.tags} onDelete={() => getTags()} />
    </>
  );
};

export default ChatTagList;
