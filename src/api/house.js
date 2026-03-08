import request from '@/utils/request'

export const houseApi = {
  // 获取房源列表 (支持搜索 keyword, 位置 location, 价格范围等)
  getHouseList(params) {
    return request.get('/house/gethouse', { params })
  },

  // 获取房源详情
  getHouseDetail(id) {
    return request.get('/house/gethouse', {
      params: { id }
    })
  },

  // 获取随机房源 (首页用)
  getRandomHouses(params) {
    return request.get('/house/getrandomhouse', { params })
  },

  // POI 搜索 (地图搜房)
  searchPOI(params) {
    return request.get('/house/location/poi', { params })
  },

  // 输入提示 (在填写房子的信息输入地址的时候，返回最相近的地址)
  getInputSuggestion(params) {
    return request.get('/house/location/tips', { params })
  },

  // 获取电子合同模板（待改）
  getContractTemplate(houseId) {
    return request.get(`/house/contract?houseId=${houseId}`)// 待改
  },

  // 创建房源 (房东用)
  createHouse(data) {
    return request.post('/house/create', data)
  }
}