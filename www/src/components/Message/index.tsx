import SessionView from './SessionView'
import SessionInput from './SessionInput'
import { useState } from 'react'
import { sendMessage } from '@/request/api'
import { Alert, Button, FloatButton, Input, Layout, Spin } from 'antd'

import MessageList from './MessageList'
import SessionList from '../SessionList'
import Sider from 'antd/es/layout/Sider'
import { Content } from 'antd/es/layout/layout'

import Loading from '@/components/Loading'

const App: React.FC = () => {
  let sessionDataStr: any[] | (() => any[]) = []

  const [sessionData, setsessionData] = useState(sessionDataStr)
  const [data, setData] = useState('')
  const [isSpinning, setisSpinning] = useState(false)

  // 获取拿到的SessionKey
  // 获取拿到的sessionData
  // 渲染对应的视图

  const cleanonClick = () => {
    setsessionData([])
  }

  const getInput = async (e: string) => {
    let sessionList = sessionData

    sessionList.push({
      role: 'user',
      content: e
    })

    setisSpinning(true);

    // 测试发送数据
    await sendMessage({ messages: sessionData }).then((res: any) => {

      if (res.code == 0) {
        sessionList.push(res.data.choices[0].message)
        setsessionData([...sessionList])
        setData('')
      }
    }).catch((e) => {
      console.log(e);

    }).finally(() => {
      setisSpinning(false);
    })

  }

  return (

    <Layout style={{ height: '100%', overflow: 'hidden' }}>
      <Sider style={{ backgroundColor: '#FFFF', height: '100%' }}>
        <SessionList></SessionList>
      </Sider>
      
      <Content style={
        { minHeight: '', overflow: 'hidden' }
      }>


          <MessageList messages={sessionData} AvatarPath={''} ></MessageList>
          <Button onClick={cleanonClick} style={{}}> 清空会话 </Button>
          <Button onClick={()=>{setisSpinning(true)}} style={{}}> 提前结束 </Button>
          <Spin spinning={isSpinning}>
          <SessionInput onChange={getInput} value={data}></SessionInput>
          </Spin>
      </Content>
      
    </Layout>
    
  )

}
export default App;
