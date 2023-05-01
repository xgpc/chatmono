// 路由懒加载
import Login from '@/views/Login'
import React, { Children, lazy } from 'react'
import { Navigate, RouteObject } from "react-router-dom"


// 进行懒加载
// import Home from "../views/Home"
// import About from "../views/About"
const Home = lazy(() => import("../views/Home"))
const About = lazy(() => import("../views/About"))

const withLoadingComponent = (comp: JSX.Element) => (
    <React.Suspense fallback={<div> Loding...</div>}>
        {comp}
    </React.Suspense>
)


const routes = [

    // 进入 更目录 跳转 /page1
    {
        path: "/",
        element: <Navigate to="/"></Navigate>

    },
    {
        path: "/",
        element: <Home></Home>,
        children: [
   
        ]
    },
    {
        path:"/login",
        element:withLoadingComponent(<Login></Login>)
    } ,
    {
        path: "*",
        element: <Navigate to="/page1"></Navigate>
    }
    ,
    {
        path: "/about",
        element: withLoadingComponent(<About></About>)
    }
    //  ,
    // {
    //     path:"/page1",
    //     element:withLoadingComponent(<Page1></Page1>)
    // },
    // {
    //     path:"/page2",
    //     element:withLoadingComponent(<Page2></Page2>)
    // }

    // {
    //     path:"/",
    //     element:<Home></Home>,
    //     childre:[
    //         {
    //             path:"/page1",
    //             element:withLoadingComponent(<Page1></Page1>)
    //         },
    //         {
    //             path:"/page2",
    //             element:withLoadingComponent(<Page2></Page2>)
    //         }
    //     ]
    // }

];


export default routes