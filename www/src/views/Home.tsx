import React, { lazy,useState, useEffect } from 'react';
import {
  AppstoreAddOutlined,
  DesktopOutlined,
  FileOutlined,
  PieChartOutlined,
  TeamOutlined,
  UserOutlined,
} from '@ant-design/icons';
import { Avatar, Button, Col, DatePicker, Input, MenuProps, Row, Space } from 'antd';
import { Breadcrumb, Layout, Menu, theme } from 'antd';
import { Outlet, useNavigate } from "react-router-dom";
import IsMobile from '@/components/Ismobile';


const MessageView = lazy(()=>import('@/components/Message')) 
const UserMenu  = lazy(()=>import('@/components/User')) 

const MobileView  = lazy(()=>import('@/components/MobileView')) 


const { Header, Content, Footer } = Layout;

const View: React.FC = () => {

  const [mobile, Setobile] = useState(false)


  useEffect(()=>{
    Setobile(IsMobile())
    return (()=>{
      console.log('刷新');
      
    })
  })


  return (

    <div style={{ height: '100vh' }}>

      {mobile == true && (<> <MobileView></MobileView> </>)}



      {mobile == false && (<Layout style={{ height: '100vh' }}>
        <Header >
          <UserMenu></UserMenu>
        </Header>

        <Content >
          <MessageView></MessageView>
        </Content>

        <Footer style={{ textAlign: 'center', padding: 0, lineHeight: "32px" }}>管理员: 82471454@qq.com</Footer>
      </Layout>)}

    </div>

  )


};

export default View;