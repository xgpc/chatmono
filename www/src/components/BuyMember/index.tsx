import { Col, Modal } from "antd";
import { useState } from "react";

const View = () => {
    const [isModalOpen, setIsModalOpen] = useState(false);

    const showModal = () => {
      setIsModalOpen(true);
    };
  
    const handleOk = () => {
      setIsModalOpen(false);
    };
  
    const handleCancel = () => {
      setIsModalOpen(false);
    };

    return(

        <>
        <Col span={4} onClick={showModal} style={{background:'rgba(0, 191, 255, 1)' }}> 会员购买 </Col>
        <Modal title="Basic Modal" open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
            <p>普通会员</p>
            <p>高级会员</p>
        </Modal>
        </>
    )


}


export default View;