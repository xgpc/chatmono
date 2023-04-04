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

const { Header, Content, Footer, Sider } = Layout;

const contentStyle: React.CSSProperties = {
  textAlign: 'center',
  minHeight: 120,
  lineHeight: '120px',
  color: '#fff',
  backgroundColor: '#108ee9',
};

const siderStyle: React.CSSProperties = {
  textAlign: 'center',
  lineHeight: '120px',
  color: '#fff',
  backgroundColor: '#3ba0e9',
};

const View: React.FC = () => {
  //antd 数据
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer },
  } = theme.useToken();

  return (
    <Layout style={{ minHeight: '100vh' }}>
      {/* 左边侧边栏 */}
      <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
        <SessionList></SessionList>
      </Sider>


      {/* 右边内容 */}
      <Layout className="site-layout">
        {/* 右边头部 */}
        <Header style={{background: colorBgContainer }} >

          <UserMenu></UserMenu>


        </Header>

        {/* 右边内容 白色底部 */}
        <Content style={{ background: colorBgContainer }} className="">


        <MessageView></MessageView>
        
          {/* 窗口部分 */}
          {/* <Outlet /> */}
          {/* 右边底部 */}
        </Content>
        <Footer style={{ textAlign: 'center', padding: 0, lineHeight: "32px" }}>管理员: 82471454@qq.com</Footer>
      </Layout>

    </Layout>
  );
};

export default View;