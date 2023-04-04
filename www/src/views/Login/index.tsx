import styles from './login.module.scss'
import { Button, Input, Space } from 'antd';
import { useState } from 'react';

const View = () => {


    const [ usernameVal, setusernameVal] = useState("");
    const [ passwordVal, setpasswordVal] = useState("");

    // 获取username
    const usernameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setusernameVal(e.target.value);
    }

    const passwordValChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setpasswordVal(e.target.value);
    }

    // 点击登录按钮
    const goToLogin = () => {
        
        console.log('usernameVal is' ,usernameVal );
        console.log('passwordVal is' ,passwordVal );
        
    }

    return (
        <div className={styles.loginPage}>
            <p>this is login</p>
            <Space direction='vertical'>
                <Input placeholder="手机号"  onChange={ usernameChange}/>
                <Input.Password placeholder="密码" onChange={passwordValChange}/>


                <div className="capchaBox">



                </div>

                <Space size='large'>
                    <Button type="primary">注册</Button>
                    <Button type='primary' onClick={ goToLogin}>登录</Button>
                </Space >

            </Space>
        </div>
    )
}

export default View