import request from '@/utils/request'

export const orderApi = {
  // 创建订单
  createOrder(data) {
    return request.post('/order/create', data)
  },

  // 获取房东 ID (用于支付或联系)
  getLandlordId(houseId) {
    return request.post('/ownerId', { "o": houseId })
  },

  // 获取我的订单列表 - 待改
  getOrderList(params) {
    console.warn('接口 /order/list 尚未实现，返回模拟数据')
    // 临时返回空数组或模拟数据
    return Promise.resolve({
      data: {
        list: [],
        total: 0
      }
    })
  },

  // 获取订单详情 - 待改
  getOrderDetail(id) {
    console.warn('接口 /order/detail 尚未实现，返回模拟数据')
    return Promise.resolve({
      data: null
    })
  }
}