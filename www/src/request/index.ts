import axios from "axios"

const instance = axios.create({
    // baseURL:"http://127.0.0.1:8081",
    baseURL:"http://www.smono.cn:8081",
    timeout:60 * 1000 // 超时设置为1分钟
})


// 请求拦截器
instance.interceptors.request.use(config=>{
    return config
}, err=>{
    return Promise.reject(err)
})

// 响应拦截器
instance.interceptors.response.use(res=>{
    return res.data
},err=>{
    return Promise.reject(err)
})


export default instance