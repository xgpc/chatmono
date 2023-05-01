import { Button, Input, Layout, Space } from "antd";
import TextArea from "antd/es/input/TextArea";
import Sider from "antd/es/layout/Sider";
import { Content } from "antd/es/layout/layout";
import { log } from "console";
import { SetStateAction, useEffect, useState } from "react";

const App = ({ value, onclick}: { value: string, onclick: Function }) => {

    const [inputData, setinputData] = useState(value)
    // 根据输入的数据返回

    const InputOnclick = () => {
        onclick(inputData)
    };


    const handleKeyDown = (e: any) => {
        if (e.key === 'Enter' && e.altKey) {
            setinputData(inputData + '\n');
        } else if (e.key === 'Enter') {
            InputOnclick();
        }
    };

    const TextAreaOnChange = (e: { target: { value: SetStateAction<string>; }; }) => {
        setinputData(e.target.value);
    };

    const CleanText = () => {
        setinputData('')
    }

    return (
        
            <Layout >
                <Content>
                    <TextArea 
                    value={inputData}
                    onKeyDown={handleKeyDown} 
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
                    <Button 
                    style={{height:'70px',}}
                    type="primary" 
                    
                    block onClick={CleanText} >   Clean </Button>
                </Sider>

            </Layout>
        
    )

}

export default App;