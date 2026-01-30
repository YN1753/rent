import request from '@/utils/request'

export async function getAddressByLocation({ latitude, longitude }) {
  try {
    const res = await request.post('/geocode/reverse'//待改后端接口
      , {
        lat: latitude,
        lng: longitude
      })
    return res
  } catch (e) {
    return { data: { formattedAddress: `${latitude},${longitude}` } }
  }
}
