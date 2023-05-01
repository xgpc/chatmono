import React from 'react';
import { Layout, Space } from 'antd';

const { Header, Footer, Sider, Content } = Layout;

const headerStyle: React.CSSProperties = {
  textAlign: 'center',
  color: '#fff',
  height: 64,
  paddingInline: 50,
  lineHeight: '64px',
  backgroundColor: '#7dbcea',
};

const contentStyle: React.CSSProperties = {
  textAlign: 'center',
  minHeight: '100%',
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

const footerStyle: React.CSSProperties = {
  textAlign: 'center',
  color: '#fff',
  backgroundColor: '#7dbcea',
};

const App: React.FC = () => (


    <Layout style={{minHeight:'100vh'}}>
      <Header style={headerStyle}>Header</Header>
      <Layout>
        <Sider style={siderStyle}>Sider</Sider>
        <Content style={contentStyle}>
        

  
        </Content>
      </Layout>
      <Footer style={footerStyle}>Footer</Footer>
    </Layout>


);

export default App;