import request from '@/utils/request'

export const payApi = {
  // 获取支付二维码
  getPayQRCode(data) {
    return request.post('/pay/qrcode', data)
  }
}