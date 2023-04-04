import { Children, useState } from 'react';
import {
    DesktopOutlined,
    FileOutlined,
    PieChartOutlined,
    TeamOutlined,
    UserOutlined,
} from '@ant-design/icons';
import type { MenuProps } from 'antd';
import { Menu } from 'antd';
import { useNavigate, useLocation } from "react-router-dom";


type MenuItem = Required<MenuProps>['items'][number];





function getItem(
    label: React.ReactNode,
    key: React.Key,
    icon?: React.ReactNode,
    children?: MenuItem[],
): MenuItem {
    return {
        key,
        icon,
        children,
        label,
    } as MenuItem;
}


// 此处只是菜单栏, 还要同步页面地址
const items: MenuItem[] = [
    getItem('新增会话', 'session', <PieChartOutlined />, [
        getItem('default', 'default', <TeamOutlined />),
    ]),
];

const Comp: React.FC = () => {

    const navigateTo = useNavigate()


    let firstopenKey:string = ''
    const findItem = (item:MenuItem) => {
        return item?.key == useLocation().pathname
    }


    for (let index = 0; index < items.length; index++) {
        const item:MenuItem = items[index]
        //TODO:后续修正
        // if (item['children']  && (item['children'] as MenuItem[]).find(findItem) == true){
        //     firstopenKey = item!.key as string;
        //     break;
        // }

    }


    const [openKeys, setopenKeys] = useState([firstopenKey]);
    const currentRoute = useLocation().pathname


    

    // 因为openkeys ,默认是空, 所以需要更具currentRoute比较路由来判断他的父级 key

    const menuClick = (e: { key: string }) => {
        console.log("dddd", e.key);
        // 编程式导航  点击以后跳转
        navigateTo(e.key)

    }
    const handleOpenChange = (keys: string[]) => {

        setopenKeys(keys.slice(-1))
    }

    return (
        <Menu theme="dark"
            defaultSelectedKeys={[currentRoute]}
            mode="inline"
            items={items}
            onClick={menuClick}
            onOpenChange={handleOpenChange}
            //当前展开项的数租
            openKeys={openKeys}
        />


    );
};

export default Comp;