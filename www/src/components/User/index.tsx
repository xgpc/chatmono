import React, { useState } from 'react';
import { DownOutlined } from '@ant-design/icons';
import { Avatar, Col, Drawer, MenuProps, Modal, Row } from 'antd';
import { Dropdown, Space } from 'antd';
import AvatarSelector from "@/components/User/UserImage"
import BuyMember from "@/components/BuyMember"
import UserInfo from './UserInfo';
import UserLogin from './UserLogin';


const App: React.FC = () => {

  const [avatar, setAvatar] = useState('点击登录');

  const handleAvatarChange = (newAvatar: React.SetStateAction<string>) => {
    setAvatar(newAvatar);
  };

    // 用户信息
    const [Draweropen, setDraweropen] = useState(false);

    // 用户信息
    const [LoginOpen, setLoginOpen] = useState(false);

  // 点击后判断是否登录, 没有登录打开登录窗口, 登录则打开用户信息页面
  function AvataronClick(): void {
      const token = localStorage.getItem("token");
      console.log(token);
      
      if (token == "" || token == null || token == undefined) {
        // 打开登录界面
        setLoginOpen(true)
      }else {
        // 打开用户信息页面
        setDraweropen(true)
      }
  }

  

  function DraweronClose(): void {
    setDraweropen(false)
  }

  function LoginClose(): void {
    setLoginOpen(false)

    // 每次关闭都判断一下 是否已经登录了
  }

  return (
    <>

      {/* 用户信息 */}
      <UserInfo Draweropen={Draweropen} DraweronClose={DraweronClose}></UserInfo>


      {/* 用户登录框 */}

      <UserLogin Draweropen={LoginOpen} DraweronClose={LoginClose} ></UserLogin>

      <Row justify="end">
      {/* <BuyMember></BuyMember> */}
        {/* <AvatarSelector value={avatar} onChange={handleAvatarChange} /> */}
       
        <Col span={4} onClick={AvataronClick} style={{color:"red", background:'rgba(	0, 191, 255, 0.5)' }}>
        <Row justify="end">
        <p style={{ marginTop: '12px' }}> {avatar} </p>
          <Avatar size={64} src={''} />    
          </Row>
        </Col>
      </Row>
    </>
  )

}

export default App;