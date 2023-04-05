import { KeyOutlined, UserOutlined } from "@ant-design/icons";
import { Button, Input, Modal } from "antd";
import { useState } from "react";

const View = ({Draweropen, DraweronClose}:{Draweropen:boolean, DraweronClose:Function}) => {
    

    
    
    const handleOk = () => {
        // 调用登录

      DraweronClose();
    };
  
    const handleCancel = () => {
        DraweronClose();
    };
  
    return (
      <>
        <Modal title="用户名登录" open={Draweropen} onOk={handleOk} onCancel={handleCancel} okText='登录' cancelText='取消'>
           <Input size="large" placeholder="账号" prefix={<UserOutlined />} />
          <br />
          <Input.Password size="large" placeholder="密码" prefix={<KeyOutlined />} />
          <br />
          
        </Modal>
      </>
    );
  };

export default View