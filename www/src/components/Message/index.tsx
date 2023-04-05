import SessionView from './SessionView'
import SessionInput from './SessionInput'
import { useState } from 'react'
import { sendMessage } from '@/request/api'
import { Alert, Button, FloatButton, Input, Spin } from 'antd'

import Notice from '@/components/Notice'
import MessageList from './MessageList'


const App: React.FC = () => {
  let sessionDataStr: any[] | (() => any[]) = []

  const [sessionData, setsessionData] = useState(sessionDataStr)
  const [isSpinning, setisSpinning] = useState(false)

  // 获取拿到的SessionKey
  // 获取拿到的sessionData
  // 渲染对应的视图

  const sleep = () => new Promise(resolve => setTimeout(resolve, 2000))

  const cleanonClick = () => {
    setsessionData([])
  }

  const getInput = async (e: string) => {
    let data = sessionData

    data.push({
      role: 'user',
      content: e
    })

    setisSpinning(true);

    await sleep();

    // 测试发送数据
    await sendMessage({messages:sessionData}).then((res:any)=>{

      if (res.code == 0) {
        data.push(res.data.choices[0].message)
      }
    })

    console.log(data);
    setisSpinning(false);
    setsessionData([...data])

  }




  return (
    <>
      {/* 通知暂时放在这 */}

      <Notice></Notice>


      {/* <Demo></Demo> */}


      <Spin tip="Loading..." spinning={isSpinning}>
        {/* <SessionView sessionList={sessionData}></SessionView> */}
        <MessageList messages={sessionData} AvatarPath={''} ></MessageList>
      
        <Button onClick={cleanonClick}> 清空会话 </Button>
        <SessionInput onChange={getInput} value={''}></SessionInput>
      </Spin>


    </>
  )

}

export default App;

function res(value: void): void | PromiseLike<void> {
  throw new Error('Function not implemented.')
}
