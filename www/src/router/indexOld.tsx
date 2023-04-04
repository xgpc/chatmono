import {BrowserRouter, Routes, Route, Navigate} from 'react-router-dom'
import App from "@/App"
import Home from "@/views/Home"
import About from "../views/About"


const baseRouter= () => (
<BrowserRouter>
        <Routes>
            <Route path='/' element={<App/>}>

                {/* 重定向 */}
                <Route path="/" element={<Navigate to="/home"/>}></Route>

                
                <Route path='/home' element={<Home/>}></Route>
                <Route path='/about' element={<About/>}></Route>
            </Route>
        </Routes>
    </BrowserRouter>
)

export default baseRouter