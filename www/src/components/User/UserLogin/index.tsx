import { KeyOutlined, UserOutlined } from "@ant-design/icons";
import { Button, Input, Modal } from "antd";
import { useState } from "react";
import {login} from '@/request/api'

const View = ({Draweropen, DraweronClose}:{Draweropen:boolean, DraweronClose:Function}) => {
    
    const [userName, setUserName] = useState('')

    const [password, setPassword] = useState('')

    
    const handleOk = () => {
        // 调用登录
        login({password:password, userName:userName}).then((res)=>{
            console.log(res);
            
    

        }).catch((e)=>{
            console.log(e);
        })




      DraweronClose();
    };
  
    const handleCancel = () => {
        DraweronClose();
    };
  
    return (
      <>
        <Modal title="用户名登录" open={Draweropen} onOk={handleOk} onCancel={handleCancel} okText='登录' cancelText='取消'>
           <Input size="large" placeholder="账号" prefix={<UserOutlined />} onChange={(e)=>{ 
            setUserName(e.target.value)
           }}/>
          <br />
          <Input.Password size="large" placeholder="密码" prefix={<KeyOutlined />}  onChange={(e)=>{ setPassword(e.target.value)}} />
          <br />
          
        </Modal>
      </>
    );
  };

export default View