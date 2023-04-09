import React from "react";
import { List, Avatar } from "antd";
import MarkdownIt from "markdown-it";

import ChatAvatar from '@/assets/images/chatImage/apple-touch-icon.png'

const md = new MarkdownIt();

interface IsessionData {
    role?: string,
    content: string
  }

const ChatBox = ({ messages, AvatarPath }:{messages:IsessionData[], AvatarPath:string}) => {
  return (
    <List
      style={{ height: "80%", overflow: "auto" }}
      dataSource={messages}
      renderItem={(item) => (
        <List.Item
          style={{
            
            display: "flex",
            flexDirection: "column",
            alignItems: "flex-start",
            
          }}
        >
            {/* // 后续加入 */}
          <Avatar  src={item.role === "user" ? AvatarPath : ChatAvatar}/>
          <div
            style={{
              maxWidth: "70%",
              padding: "10px",
              borderRadius: "8px",
              backgroundColor: item.role === "user" ? "##00bfff" : "#f0ffff",
            }}
            dangerouslySetInnerHTML={{ __html: md.render(item.content) }}
          />
        </List.Item>
      )}
    />
  );
};

export default ChatBox;