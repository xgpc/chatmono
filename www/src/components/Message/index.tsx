import SessionView from './SessionView'
import SessionInput from './SessionInput'
import { useState } from 'react'
import { sendMessage } from '@/request/api'
import { Alert, Button, FloatButton, Input, Layout, Spin } from 'antd'

import MessageList from './MessageList'
import SessionList from '../SessionList'
import Sider from 'antd/es/layout/Sider'
import { Content } from 'antd/es/layout/layout'
import { encode, decode } from 'gpt-3-encoder'


import Loading from '@/components/Loading'

const App: React.FC = () => {
  let sessionDataStr: any[] | (() => any[]) = []

  const [sessionData, setsessionData] = useState(sessionDataStr)
  const [inputData, setinputData] = useState('')
  const [isSpinning, setisSpinning] = useState(false)
  

  // 获取拿到的SessionKey
  // 获取拿到的sessionData
  // 渲染对应的视图

  const cleanonClick = () => {
    setsessionData([])
  }

  const getInput = async (e: string) => {

    calcToken(e)
    return 
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
        setsessionData([...sessionList])
        setinputData('')
      }
    }).catch((e) => {
      console.log(e);

    }).finally(() => {
      setisSpinning(false);
    })

  }

  const calcToken = (e:string)=>{
    console.log(e);

    

    const str = 'This is an example sentence to try encoding out on!'
    const encoded = encode(str)
    console.log('Encoded this string looks like: ', encoded)

    console.log('We can look at each token and what it represents')
    for(let token of encoded){
      console.log({token, string: decode([token])})
    }

    const decoded = decode(encoded)
    console.log('We can decode it back into:\n', decoded)
    
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
          <Button onClick={()=>{setisSpinning(false)}} style={{}}> 提前结束 </Button>
          
          <Spin spinning={isSpinning}>
          <SessionInput onChange={getInput} value={inputData}></SessionInput>
          </Spin>
      </Content>
      
    </Layout>
    
  )

}
export default App;
