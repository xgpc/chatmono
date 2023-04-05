import { Avatar, Button, Modal } from 'antd';
import { useState } from 'react';

import img0 from '@/assets/images/0.jpg'
import img1 from '@/assets/images/1.webp'
import img2 from '@/assets/images/2.webp'
import img3 from '@/assets/images/3.webp'
import img4 from '@/assets/images/4.webp'
import img5 from '@/assets/images/5.webp'
import img6 from '@/assets/images/6.webp'
import img7 from '@/assets/images/7.webp'
import img8 from '@/assets/images/8.webp'
import img9 from '@/assets/images/9.webp'
import img10 from '@/assets/images/10.webp'
import img11 from '@/assets/images/11.webp'
import img12 from '@/assets/images/12.webp'
import img13 from '@/assets/images/13.webp'
import img14 from '@/assets/images/14.webp'
import img15 from '@/assets/images/15.webp'
import img16 from '@/assets/images/16.webp'


export default function AvatarSelector({ value, onChange }:{value:string, onChange:Function}) {
  const [modalVisible, setModalVisible] = useState(false);
  const [selectedAvatar, setSelectedAvatar] = useState(value != '' ? value : img0);
  const avatars1 = [
    img0,
    img1,
    img2,
    img3,
    img4,
    img5,
    img6,
    img7,
    img8,
    img9,
    img10,
    img11,
    img12,
    img13,
    img14,
    img15,
    img16,
  ];


  const handleSelectAvatar = (avatar: string) => {
    setSelectedAvatar(avatar);
  };

  const handleOk = () => {
    onChange(selectedAvatar);
    setModalVisible(false);
  };

  const handleCancel = () => {
    setSelectedAvatar(value);
    setModalVisible(false);
  };

  return (
    <>
      <Avatar size={64} src={selectedAvatar} onClick={() => setModalVisible(true)} />
      <Modal title="选择头像"
        open={modalVisible}
        // visible={modalVisible}
        onOk={handleOk}
        onCancel={handleCancel}
      >
        <div style={{ display: 'flex', flexWrap: 'wrap' }}>
          {avatars1.map((avatar) => (
            <Avatar key={avatar} size={64} src={avatar} style={{ margin: '12px' }}
              onClick={() => handleSelectAvatar(avatar)} className={selectedAvatar === avatar ? 'selected' : ''} />
          ))}

        </div>
      </Modal>
    </>
  );
}