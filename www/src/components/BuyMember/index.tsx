import { AlipayCircleOutlined, WechatOutlined } from "@ant-design/icons";
import { Avatar, Button, Col, Modal, Radio, RadioChangeEvent } from "antd";
import { useState } from "react";

const View = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);


  const [product_id, setproduct_id] = useState(0);
  const [payType, setpayType] = useState(0);

  const showModal = () => {
    setIsModalOpen(true);
  };

  const handleOk = () => {
    // 调用支付api




    setIsModalOpen(false);
  };

  const handleCancel = () => {
    setIsModalOpen(false);
  };

  const onChange = (e: RadioChangeEvent) => {
    console.log(`radio checked:${e.target.value}`);
  };

  const optionsWithDisabled = [
    { label: '尝鲜: 2元1天', value: 1 },
    { label: '普通会员: 20元/每月', value: 2 },
    { label: '高级会员: 50元/每月', value: 3, disabled: true },
  ];

  const onChange4 = ({ target: { value } }: RadioChangeEvent) => {
    console.log('radio4 checked', value);
    setproduct_id(value);
  };

  return (

    <>
      <Button style={{ minWidth: '300px', minHeight: '50px' }} onClick={showModal}>
        会员购买


      </Button>
      <Modal title="Basic Modal"
        cancelText={'取消'}
        okText={'获取支付码'}
        open={isModalOpen}
        onOk={handleOk}
        onCancel={handleCancel}>

        <Radio.Group
        options={optionsWithDisabled}
        onChange={onChange4}
        value={product_id}
        optionType="button"
        buttonStyle="solid"
      />



        {product_id != 0 && (
          <>
          <br /><br />
          <p>购买方式:</p>
              <Radio.Group onChange={onChange} defaultValue="" style={{ marginTop: 16 }}>
            
              <Radio value="wechat" >  <WechatOutlined /> 微信支付</Radio>
              <Radio value="alibaba" disabled> <AlipayCircleOutlined />支付宝支付 </Radio>
            </Radio.Group>
          </>
        )}


        {/* 从全局拿去用户信息, 提示要为谁购买, 时长多久 */}



      </Modal>
    </>



  )


}


export default View;