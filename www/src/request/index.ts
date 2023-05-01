import axios from "axios"

const instance = axios.create({
    // baseURL:"http://127.0.0.1:8081",
    baseURL:"http://www.smono.cn:8081",
    // baseURL:"http://chat.smono.cn:8081",
    timeout: 2* 60 * 1000 // 超时设置为2分钟
})


// 请求拦截器
instance.interceptors.request.use(config=>{
    
    let token = localStorage.getItem("token");
    console.log(config);
    console.log(token);
    
    if ( token ) {
        let index:number = token.indexOf('"')
        let last:number = token.lastIndexOf('"')
        if( index!=-1 && last !=-1 ){
            token = token.substr( 1 , token.length-2 )
        }
        config.headers['token'] = token;
    }
    console.log(config);
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