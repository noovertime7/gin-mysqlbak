//引入加密模块
import JSEncrypt from 'jsencrypt';



export const Encrypt = (data) =>{
        const  publicKey = `-----BEGIN PUBLIC KEY-----
        MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDAJRUhzFZ9P64cic8slOpn82Vl
        YJUusLKWTKqugn7lgNVUpWdVCagfhtkViTUg5KRvpGrESmrQPlRiImm/iX/rOKtQ
        SyMq6UroBpafjL3t6sPyHHFohdDVakR7b6S6UG6ZOaTMOC/avPWtInXgzU05sHjR
        qEIGFapVejRtgfPtbwIDAQAB
        -----END PUBLIC KEY-----`
        const jsEncrypt = new JSEncrypt();
        jsEncrypt.setPublicKey(publicKey);
        return jsEncrypt.encrypt(data);
    }


export const Decrypt = (data) => {
    const privateKey = `-----BEGIN RSA PRIVATE KEY-----
    MIICXQIBAAKBgQDAJRUhzFZ9P64cic8slOpn82VlYJUusLKWTKqugn7lgNVUpWdV
    CagfhtkViTUg5KRvpGrESmrQPlRiImm/iX/rOKtQSyMq6UroBpafjL3t6sPyHHFo
    hdDVakR7b6S6UG6ZOaTMOC/avPWtInXgzU05sHjRqEIGFapVejRtgfPtbwIDAQAB
    AoGBAJ+grxOrHNdlBhLzckhJVwwRK1WzjXyCk3tGKi5cf2vPQmvWFiiRozi94K+B
    k7/F885EO+bjJCXpAlWc3Vmgs8GYPIqUPJSTVkiDMf8nV5qicxzOH/pqye6KE1DD
    C/m3gPJTwOk/oT1KWsA4AHn8wmup01mDkMX82U9WFtJ1ZXXBAkEA5vTvXsFhaizK
    sAm8rkKjavxulCtuuNgyN+w68AUcdX92Cb4Hw6cWnAlUyqYncWLu8+3/TVWrtJWT
    vP0zcm1k8QJBANT6xzIRdLZQYPHODZDD0p575rfJgR4wuZi0tzhPmJKdndjfaxaW
    fco9GjL2yZzH4aNpF+ReN21RmT1ewgKy+F8CQDrEcHRH+KWvqBOLJrugsTxz5x9E
    vfPC72RTc9vHMSqkuEBaXldmmNYzeaPnC3pKlkrzcFcZSYu109XvB7xCIcECQDiE
    P73SkgUbOU6RXlovDMIPoP7eUwwe4/FY61HfFV66wrtdNj6tOr4jDsO9Z2zaQc8q
    QTPRqKWyxJZbgeJTecMCQQDFnb5C0dHqpDE1PBHklzo9TnycUG9T1gBIVC7oZDex
    ImeHIUC/olM27UPRzf4ku+ZtMb+bTZpjcUcRBzs5JnAK
    -----END RSA PRIVATE KEY-----`
    const decrypt = new JSEncrypt()
    decrypt.setPrivateKey(privateKey)
    return decrypt.decrypt(data)
}