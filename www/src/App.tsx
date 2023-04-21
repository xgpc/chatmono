import { Button, Space } from 'antd';
import { FastBackwardOutlined } from '@ant-design/icons';
import { useRoutes, Link } from "react-router-dom"
import router from "./router"
import Home from './views/Home';
import Notice from './components/Notice';
import LayoutDemo from '@/components/LayoutDemo';

function App() {
  const outlet = useRoutes(router)



  return (
    <div className="App" >
     {/* {outlet}  */}

{/* <LayoutDemo></LayoutDemo> */}

     <Notice></Notice>
     <Home/>
    </div>
  )
}

export default App
