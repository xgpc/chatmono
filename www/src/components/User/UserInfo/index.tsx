import BuyMember from "@/components/BuyMember"
import { Avatar, Button, Drawer, Space } from "antd"
import React from "react"
import { useEffect, useState } from "react"


export interface Iuser {
  "id": number,
  "user_name": string,
  "user_img": string,
  "city": "",
  "mobile": string,
}

const View = ({ userInfo, Draweropen, onCancel, DraweronClose }: {
  Draweropen: boolean,
  DraweronClose: (e: React.MouseEvent | React.KeyboardEvent | undefined) => void ,
  onCancel: Function,
  userInfo: Iuser
}) => {


  const tmp:Iuser = {
    id: 0,
    user_name: "",
    user_img: "",
    city: "",
    mobile: ""
  }
  const [user,setuser] = useState(tmp)

  useEffect(() => {
    // 获取token
    // 获取用户信息
    console.log('用户信息变更');
    console.log(userInfo);
    setuser(userInfo)



  }, [userInfo])



  // 获取用户信息


  // 退出
  const onclickQuit = () => {
    localStorage.setItem('token', '') 
    setuser(tmp)
    onCancel(tmp)
    DraweronClose(undefined)
    
  }


  // button style
  const ButtonStyle = {minWidth:'300px', minHeight:'50px'}

  return (
    <Drawer title="用户信息" placement="right"
      onClose={DraweronClose}
      open={Draweropen}>

      <Space size='large' direction="vertical">
        <br />
        <Space size='large'>
          <Avatar size={120} src={user.user_img} />
          <Space direction="vertical" size="middle" style={{ display: 'flex' }}>
            <p>用户名: {user.user_name}</p>
            <p>手机号: {user.mobile}</p>
            {/* <p> 级别 </p> */}
            <p> 到期时间: { } </p>
          </Space>
          <br />
        </Space>
        <Button style={ButtonStyle}>修改信息</Button>
        <Button style={ButtonStyle} onClick={onclickQuit} >退出登录</Button>
        {/* <Button style={ButtonStyle}>会员购买</Button> */}
        <BuyMember/>
      </Space>




    </Drawer>
  )
}

export default View
