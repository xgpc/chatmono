
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'

import { Button, Space } from 'antd';
import { FastBackwardOutlined } from '@ant-design/icons';
import { useRoutes, Link } from "react-router-dom"
import router from "./router"
import Home from './views/Home';

function App() {
  const outlet = useRoutes(router)



  return (
    <div className="App">
     {/* {outlet}  */}
     <Home></Home>
    </div>
  )
}

export default App
