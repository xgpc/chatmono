import { Button, Layout, Row, Space, Spin } from "antd";
import TextArea from "antd/es/input/TextArea";
import Sider from "antd/es/layout/Sider";
import UserMenu from '@/components/User'
import MessageList from '@/components/Message/MessageList'
import { SetStateAction, useEffect, useState } from "react";
import { MenuFoldOutlined } from "@ant-design/icons";
import { sendMessage } from "@/request/api";


const { Header, Content, Footer } = Layout;

const App = () => {
  let sessionDataStr: any[] | (() => any[]) = []
  
  const [sessionData, setSessionData] = useState(sessionDataStr)
  const [inputData, setinputData] = useState('')

  const [isSpinning, setisSpinning] = useState(false)
  // 根据输入的数据返回

  // useEffect(() => {
  //   setinputData(value)
  // }, [value])

  const InputOnclick = () => {
    getInput(inputData)
  };

  const TextAreaOnChange = (e: { target: { value: SetStateAction<string>; }; }) => {
    setinputData(e.target.value);
  };


  const getInput = async (e: string) => {

    let sessionList = [...sessionData]

    sessionList.push({
      role: 'user',
      content: e
    })


    setisSpinning(true);

    // 测试发送数据
    await sendMessage({ messages: sessionList }).then((res: any) => {

      if (res.code == 0) {
        sessionList.push(res.data.choices[0].message)
        setSessionData([...sessionList])
        setinputData('')
      }
    }).catch((e) => {
      console.log(e);

    }).finally(() => {
      setisSpinning(false);
    })

  }

  const cleanonClick = () => {
    setSessionData([])
  }

  return (

    // <Layout >

    //   <Header style={{ display: 'flex', justifyContent: 'flex-end', alignItems: 'center' }}>
    //     <Button icon={<MenuFoldOutlined />} />
    //   </Header>

    <Layout style={{ display: 'flex', flexDirection: 'column', height: '100vh', width: '100vw' }}>
      <Header style={{ display: 'flex', justifyContent: 'flex-end', alignItems: 'center' }}>
      
        <Button icon={<MenuFoldOutlined />} />
        <Button icon={<MenuFoldOutlined />} />
        
      </Header>


      <Content style={{ flex: 1, display: 'flex', flexDirection: 'column', justifyContent: 'center' }}>
      {/* <MessageList messages={sessionData} AvatarPath={''} style={{ flexGrow: 1 }} /> */}
      <MessageList messages={sessionData} AvatarPath={''}  />

      <Space align="center" style={{ width: '100vw' }}>
      <Button onClick={cleanonClick} style={{}}> 清空会话 </Button>
      <Button onClick={()=>{setisSpinning(false)}} style={{}}> 提前结束 </Button>
    </Space>
        <Spin spinning={isSpinning}>
          <Space align="center" style={{ width: '100vw' }}>
          
            <TextArea
              // bordered={false}
              
              maxLength={100}
              autoSize={{ minRows: 2, maxRows: 3 }}
              showCount
              onChange={TextAreaOnChange}
              placeholder="输入内容"
              style={{ display: 'flex', width: '80vw' }} />
            <Button
              type="primary"
              block onClick={InputOnclick} > 发送</Button>
          </Space>
        </Spin>

      </Content>


    </Layout>

  )

}

export default App;