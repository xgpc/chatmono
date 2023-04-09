import React, { useState } from 'react';
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

import MainMenu from '@/components/MainMenu'
import MessageView from '@/components/Message'
import SessionList from '@/components/SessionList'
import UserMenu from '@/components/User'
import Sider from 'antd/es/layout/Sider';

const { Header, Content, Footer } = Layout;

const View: React.FC = () => {




  return (

    <Layout style={{ height: '100vh' }}>
      {/* 右边头部 */}
      <Header >
        <UserMenu></UserMenu>
      </Header>

      <Content >


        <MessageView></MessageView>

      </Content>

      <Footer style={{ textAlign: 'center', padding: 0, lineHeight: "32px" }}>管理员: 82471454@qq.com</Footer>

    </Layout>



  );
};

export default View;