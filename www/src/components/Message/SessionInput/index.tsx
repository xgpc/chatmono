import { Button, Input, Layout, Space } from "antd";
import TextArea from "antd/es/input/TextArea";
import Sider from "antd/es/layout/Sider";
import { Content } from "antd/es/layout/layout";
import { log } from "console";
import { SetStateAction, useEffect, useState } from "react";

const App = ({ value, onChange }: { value: string, onChange: Function }) => {

    const [inputData, setinputData] = useState(value)
    // 根据输入的数据返回

    useEffect(()=>{
        setinputData(value)
    },[value])

    const InputOnclick = () => {
        onChange(inputData)
    };

    const TextAreaOnChange = (e: { target: { value: SetStateAction<string>; }; }) => {
        setinputData(e.target.value);
    };

    return (
        
            <Layout >
                <Content>
                    <TextArea 
                    onKeyDown={InputOnclick}
                    autoSize={{minRows: 6, maxRows: 6 } }
                    showCount={true}
                    onChange={TextAreaOnChange} 
                    placeholder="输入内容"  
                    style={{ display: 'flex' }} />
                </Content>
                <Sider style={{backgroundColor:'white'}}>
                    <Button 
                    style={{height:'70px',}}
                    type="primary" 
                    
                    block onClick={InputOnclick} >   发送</Button>
                </Sider>

            </Layout>
        
    )

}

export default App;