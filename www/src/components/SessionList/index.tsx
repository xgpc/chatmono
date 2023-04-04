import React, { useState } from 'react';
import {
    AppstoreAddOutlined,
  AppstoreOutlined,
  CalendarOutlined,
  CommentOutlined,
  LinkOutlined,
  MailOutlined,
  SettingOutlined,
  SolutionOutlined,
} from '@ant-design/icons';
import { Button, Menu, notification, Switch } from 'antd';
import type { MenuProps, MenuTheme } from 'antd/es/menu';
import { time } from 'console';
import debug from 'debug';

type MenuItem = Required<MenuProps>['items'][number];

function getItem(
  label: React.ReactNode,
  key?: React.Key | null,
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

const items: MenuItem[] = [
  getItem('默认会话', 'default', <SolutionOutlined />), 
];



const App: React.FC = () => {
    
    const [num, setNum] = useState(1);
    const [menuKeys, setMenuKeys] = useState(items);
    const [currentKey, setcurrentKey] = useState(['default']);

    const handleClickAdd = () => {
        
        // 获取
        // TODO:从后端获取会话ID

        setMenuKeys([...menuKeys, getItem('New' + num , num, <CommentOutlined />)]);
        setNum(num + 1)
    };
    


    const [api, contextHolder] = notification.useNotification();


    const handleClickDel = () => {
        // setMenuKeys([...menuKeys, getItem('New', '3', <AppstoreOutlined/>)]);
        console.log('del onclick', currentKey);
        for (let index = 0; index < menuKeys.length; index++) {
            const element = menuKeys[index];
    
            if (currentKey[0] == 'default') {
                api['error']({
                    message: '删除失败',
                    description:
                      '当前只能删除自己的创建的对话',
                  });

                  break;
            }
            setMenuKeys([...menuKeys])
            if (element?.key == currentKey[0]) {
                menuKeys.splice(index, 1)
                setMenuKeys([...menuKeys])
                setcurrentKey(['default'])
            }
            
        }
        
        
    };

    const menuClick = (e:{key:string}) =>{
        // 记录当前所选会话
        console.log(e);
        setcurrentKey([e.key])
    }

    

  return (    
    <>
    {contextHolder}
     <Button onClick={handleClickAdd}>Add Menu</Button>
     <Button onClick={handleClickDel}>Del Menu</Button>
      <Menu
        // defaultSelectedKeys={[currentRoute]}
        mode="inline"
        onClick={menuClick}
        selectedKeys={ currentKey }
        // onOpenChange={handleOpenChange}
        //当前展开项的数租
        // onChange={onChangeMenu}
        items={menuKeys}
      />
    </>
    
      
  );
};

export default App;