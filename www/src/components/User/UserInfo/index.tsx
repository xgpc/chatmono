import { Drawer } from "antd"

const View = ({Draweropen, DraweronClose}:{Draweropen:boolean, DraweronClose:(e: React.MouseEvent | React.KeyboardEvent) => void}) => {



    return (
        <Drawer title="用户信息" placement="right" onClose={DraweronClose} open={Draweropen}>
        <p>用户名:</p>
        <p>手机号</p>
        <p>头像</p>
        <p> 级别 </p>
        <p> 剩余时间 </p>
        {/* 重新登录, 退出 */}
      </Drawer>
    )
}

export default View