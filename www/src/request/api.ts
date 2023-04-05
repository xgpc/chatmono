import axios from "./index"


export function sendMessage( params:object):Promise<void>{

    let url:string = `/api/openAI/send`
    return axios.post(url, params);
}


export function login(params:object):Promise<void>{

    let url:string = `/api/user/login`
    return axios.post(url, params);
}
