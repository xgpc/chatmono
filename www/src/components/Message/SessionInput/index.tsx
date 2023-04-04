import { Button, Input, Layout, Space } from "antd";
import TextArea from "antd/es/input/TextArea";
import Sider from "antd/es/layout/Sider";
import { Content } from "antd/es/layout/layout";
import { log } from "console";
import { SetStateAction, useState } from "react";

const App = ({ value, onChange }: { value: string, onChange: Function }) => {

    const [inputData, setinputData] = useState(value)
    // 根据输入的数据返回

    const InputOnclick = () => {
        onChange(inputData)
    };

    const TextAreaOnChange = (e: { target: { value: SetStateAction<string>; }; }) => {
        setinputData(e.target.value);
    };





    return (
        <div>
            <Layout>
                <Content>
                    <TextArea onChange={TextAreaOnChange} rows={4} placeholder="输入内容" maxLength={300} style={{ display: 'flex' }} />
                </Content>
                <Sider style={{backgroundColor:'white'}}>
                    <Button type="primary" block onClick={InputOnclick} style={{ marginLeft: '10px', flexShrink: '0' }}>   发送</Button>
                </Sider>

            </Layout>
        </div>
    )

}

export default App;