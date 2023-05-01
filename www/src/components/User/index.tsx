import React, { useState, useEffect } from 'react';
import { DownOutlined } from '@ant-design/icons';
import { Avatar, Col, Drawer, MenuProps, Modal, Row } from 'antd';
import { Dropdown, Space } from 'antd';
import AvatarSelector from "@/components/User/UserImage"
import BuyMember from "@/components/BuyMember"
import UserInfo, {Iuser} from './UserInfo';
import UserLogin from './UserLogin';


const App: React.FC = () => {

  const temp:Iuser = {
    id: 0,
    user_name: '',
    user_img: '',
    city: '',
    mobile: ''
  }
  // 用户信息
  const [userData, setuserData] = useState(temp);

  // 用户信息是否打开
  const [Draweropen, setDraweropen] = useState(false);

  // 是否登录
  const [LoginOpen, setLoginOpen] = useState(false);



  useEffect(()=>{
    const token = localStorage.getItem("token");
    if (token == "" || token == null){
      return
    }

    // token不为空的情况下获取用户信息
    getUserInfo(token)
 
  })


  const getUserInfo = (token:string)=>{
    if (token == ""){
      return
    }


  }



  // 点击后判断是否登录, 没有登录打开登录窗口, 登录则打开用户信息页面
  function AvataronClick(): void {
    const token = localStorage.getItem("token");
    console.log(token);

    if (token == "" || token == null || token == undefined) {
      // 打开登录界面
      setLoginOpen(true)
    } else {
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

  function UserInfoOnCancel(params: any) {
    setuserData(params)
  }

  return (
    <>

      {/* 用户信息 */}
      <UserInfo Draweropen={Draweropen} 
        DraweronClose={DraweronClose}
        onCancel={UserInfoOnCancel} userInfo={userData} />

      {/* 用户登录框 */}
      <UserLogin
        Draweropen={LoginOpen}
        DraweronClose={LoginClose}
        onChange={(e: React.SetStateAction<Iuser>) => {
          
          console.log(e);          
          setuserData(e)
        }} />

      <Row justify="end">
        
          <Row justify="end">
            <p style={{}}> { userData.id == 0 ?'点击登录':userData.user_name} </p>
            <Avatar onClick={AvataronClick} size={64} src={userData.user_img}>点击登录</Avatar>
          </Row>
        
      </Row>
    </>
  )

}

export default App;