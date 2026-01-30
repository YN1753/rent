import request from '@/utils/request'

// 登录
export function loginAPI(data) {
  return request({ url: '/user/login', method: 'post', data })
}

// 注册
export function registerAPI(data) {
  return request({ url: '/user/register', method: 'post', data })
}

//生成验证码
export function sendCodeAPI(data) {
  return request({ url: '/user/gencode', method: 'post', data })
}

//校验验证码
export function verifyCodeAPI(data) {
  return request({ url: '/user/authcode', method: 'post', data })
}

// 上传用户位置
export function updateUserLocation(data) {
  return request({
    url: '/location',
    method: 'get',
    params: data
  })
}
