import {$post} from "@/utils/request";
import {Decrypt, Encrypt} from "@/utils/security";
// import md5 from 'js-md5'
// 登陆api
export const $Login =  async (params) => {
    params.password =  Encrypt(params.password)
    console.log(params.password)
   let  depasswd  = Decrypt(params.password)
    console.log(depasswd)
   let data =  $post('/admin_login/login',params)
    console.log(data)
}
