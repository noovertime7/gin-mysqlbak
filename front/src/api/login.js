import {$post} from "@/utils/request";
import { Encrypt} from "@/utils/security";
// import md5 from 'js-md5'
// 登陆api
export const $Login =  async (params) => {
    params.password = Encrypt(params.password)
    let data = $post('/admin_login/login', params)
    console.log(data)
}
