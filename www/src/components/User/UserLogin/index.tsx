import { KeyOutlined, MobileOutlined, UserOutlined } from "@ant-design/icons";
import { Button, Input, Modal, Space, Switch } from "antd";
import { useState } from "react";
import { login, logon } from '@/request/api'
import { log } from "console";
import AvatarSelector from "../UserImage";

import img0 from '@/assets/images/0.jpg'

interface ResBody {
    code:number
    data: {}|[]
    msg:string
}

const View = ({ Draweropen, DraweronClose }: { Draweropen: boolean, DraweronClose: Function }) => {

    const [isRegisterClicked, setisRegisterClicked] = useState(false)
    const [userName, setUserName] = useState('')
    const [password, setPassword] = useState('')
    const [againPassword, setagainPassword] = useState('')
    const [mobile, setmobile] = useState('')

    const [avatar, setAvatar] = useState(img0);

    const handleAvatarChange = (newAvatar: React.SetStateAction<string>) => {
        setAvatar(newAvatar);
    };


    const handleOk = async () => {


        if (!isRegisterClicked) {
            // 调用登录
            
            await login({ password: password, user_name: userName }).then((res:ResBody|any) => {
                if (res.code == 0){
                    // 设置token
                    localStorage.setItem('token', res.data.token) 
                }
                
            }).catch((e) => {
                console.log(e);
            })
        } else {
            // 调用注册
            // 调用登录
            if (password != againPassword) {
                // 调用message提醒
                return;
            }
            await logon({ password: password, user_name: userName, mobile: mobile, user_img:avatar }).then((res:ResBody|any) => {
                                if (res.code == 0){
                    // 设置token
                    localStorage.setItem('token', res.data.token) 
                }
                
            }).catch((e) => {
                console.log(e);
            })

        }




        DraweronClose();
    };

    const handleCancel = () => {
        DraweronClose();
    };


    return (
        <>
            <Modal title={isRegisterClicked ? "用户注册" : "用户登录"} 
            open={Draweropen} 
            onOk={handleOk} 
            onCancel={handleCancel} 
            okText={isRegisterClicked ? "注册" : "登录"} 
            cancelText='取消'>
                <Space direction="vertical" size="middle" style={{ display: 'flex' }}>
                    {/* 注册选择头像 */}
                    {isRegisterClicked == true && (
                        <>
                            <AvatarSelector value={avatar} onChange={handleAvatarChange} />
                        </>
                    )}

                    <Input size="large" placeholder="账号" prefix={<UserOutlined />} onChange={(e) => {
                        setUserName(e.target.value)
                    }} />
                    <Input.Password size="large" placeholder="密码" prefix={<KeyOutlined />} onChange={(e) => { setPassword(e.target.value) }} />

                    {/* 注册填写资料 */}
                    {isRegisterClicked == true && (
                        <>
                            <Input.Password size="large" placeholder="确认密码" prefix={<KeyOutlined />} onChange={(e) => { setagainPassword(e.target.value) }} />

                            <Input size="large" placeholder="联系电话" prefix={<MobileOutlined />} onChange={(e) => { setmobile(e.target.value) }} />
                        </>
                    )}

                    <Switch size="default" checkedChildren="切换登录" unCheckedChildren="切换注册" onChange={(e) => {
                        console.log(e);
                        setisRegisterClicked(e)
                    }} />

                </Space>
            </Modal>
        </>
    );
};

export default View