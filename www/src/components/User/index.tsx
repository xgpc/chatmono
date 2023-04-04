import React, { useState } from 'react';
import { DownOutlined } from '@ant-design/icons';
import { Col, MenuProps, Modal, Row } from 'antd';
import { Dropdown, Space } from 'antd';
import AvatarSelector from "@/components/UserImage"

const items: MenuProps['items'] = [
  {
    label: (
      <a target="_blank" rel="noopener noreferrer" href="https://www.antgroup.com">
        1st menu item
      </a>
    ),
    key: '0',
  },
  {
    label: (
      <a target="_blank" rel="noopener noreferrer" href="https://www.aliyun.com">
        2nd menu item
      </a>
    ),
    key: '1',
  },
  {
    type: 'divider',
  },
  {
    label: '3rd menu item（disabled）',
    key: '3',
    disabled: true,
  },
];

const App: React.FC = () => {


  const [d, setd] = useState('用户姓名 + 头像');
  const onClickCol = ()=>{
    setd(d+1)
  }


  const [isModalOpen, setIsModalOpen] = useState(false);

  const showModal = () => {
    setIsModalOpen(true);
  };

  const handleOk = () => {
    setIsModalOpen(false);

    setd(d+1)

  };

  const handleCancel = () => {
    setIsModalOpen(false);
  };


  const [avatar, setAvatar] = useState('');

  const handleAvatarChange = (newAvatar: React.SetStateAction<string>) => {
    setAvatar(newAvatar);
  };

  return (
    <>
  
    <Row justify="end">
    <AvatarSelector value={avatar} onChange={handleAvatarChange} />
      <p style={{ marginTop: '12px' }}>Selected Avatar: {avatar}</p>

            {/* <Col span={4} onClick={showModal} style={{background:'rgba(0, 191, 255, 1)' }}>   
              TODO: 鼠标放上去以后最好有个变色
              <p> 购买会员</p>
            </Col> */}

              <Col span={4} onClick={onClickCol} style={{color:"red", background:'rgba(	0, 191, 255, 0.5)' }}>{d}</Col>
              {/* 添加购买按钮 */}
              {/* 添加用户头像 */}
            </Row>

            <Modal title="Basic Modal" open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
            <p>普通会员</p>
            <p>高级会员</p>
            
          </Modal>
  
    </>
  )

}

export default App;