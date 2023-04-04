
import axios from 'axios'
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import type { AxiosRequestConfig, AxiosError, AxiosResponse } from 'axios'




axios.defaults.baseURL = "http://127.0.0.1:8081";


interface newAxiosRequestConfig extends AxiosRequestConfig {
    urlKey?: string,
    headersType?: string
}
axios.interceptors.request.use((config: newAxiosRequestConfig): any => {
    NProgress.start();
    config.data = JSON.stringify(config.data);
    config.headers = {
        "Content-Type": "application/json",
    };
    return config;
}, (error: AxiosError): Promise<void> => {
    NProgress.done();
    return Promise.reject(error)
}
)

axios.interceptors.response.use(
    (res: AxiosResponse) => {

        NProgress.done();
        let { status, config } = res;
        // if( status == 200 && config?.data?.headersType == "formData"){
        //     return {code:0};
        // }
        let Code: any = null;
        let Msg: string = "";
        let Data: any = ""
        let response: any = res.data
        Msg = response['msg'] ? response['msg'] : response['Msg']
        Data = response['Data'] ? response['Data'] : response['data']
        if (response['code'] >= 0) Code = response['code']
        if (response['Code'] >= 0) Code = response['Code']
        let resObj: any = {
            code: Code,
            msg: Msg,
            data: Data ?? {}
        }
        if (Code == 0) {
            return resObj
        } else {
            return resObj
        }
    },
    (error: AxiosError): Promise<void> => {
        NProgress.done();
        return Promise.reject(error)
    }
)
export default axios