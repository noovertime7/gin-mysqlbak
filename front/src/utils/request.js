import axios from "axios";

var instance = axios.create({
    baseURL: 'http://localhost:8080/api/',
    timeout: 1000,
});

// get方法
export const $get = async (url,params) => {
   let {data} = await instance.get(url,{
        params
    })
    return data
}

// post方法
export const $post = async (url,params) => {
    return await instance.post(url, params)
}

// put方法
export const $put = async (url,data) => {
    let {resp} = await  instance.put(url,data)
    return resp
}

// delete方法
export const $delete = async (url,data) => {
    let {resp} = await  instance.delete(url,{data})
    return resp
}