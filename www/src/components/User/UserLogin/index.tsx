import { Button, Modal } from "antd";
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
        <Modal title="Basic Modal" open={Draweropen} onOk={handleOk} onCancel={handleCancel} okText='登录' cancelText='取消'>
          
        </Modal>
      </>
    );
  };

export default View