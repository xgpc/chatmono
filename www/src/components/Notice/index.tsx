import { Modal } from "antd"
import { useState } from "react";



const View = () => {

    const [isModalOpen, setIsModalOpen] = useState(true);

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
        <Modal title="Basic Modal" open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
            <h1>chatmono持续更新中</h1>
            <h1> 使用中有问题可以联系下方作者邮箱,或者添加QQ </h1>
            <p>  当前页面只提供5次上下文对话, 既你可以提问5次 + openAI回答5次</p>
            <p>  对话达到最大次数后请点击清空会话重新发起提问!</p>

      </Modal>
    )
}

export default View