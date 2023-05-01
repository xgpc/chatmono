import { Modal, message } from "antd"
import { useState } from "react";



export const MessageError = (msg:string) => {
  const [messageApi, contextHolder] = message.useMessage();
  messageApi.open({
    type: 'error',
    content: msg,
  });
};

const View = () => {
  
    const [isModalOpen, setIsModalOpen] = useState(true);
    const [messageApi, contextHolder] = message.useMessage();
    const showModal = () => {
      setIsModalOpen(true);
    };
  
    const handleOk = () => {
      setIsModalOpen(false);
    };
  
    const handleCancel = () => {
      setIsModalOpen(false);
    };

    return (

      <>
        {contextHolder}
        <Modal title="chat mono" open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
            <h1>chatmono持续更新中</h1>
            <p> 服务器进行了更新, 目前已经迁移俄勒冈州, 稳定使用</p>
            <p> 当前修改为账号登录后使用,<br />
             因之前免费使用不做限制被人拿去训练了</p>
            

      </Modal>
      </>
    )
}

export default View