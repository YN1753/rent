import { defineStore } from 'pinia';
import { loginAPI, registerAPI, updateUserLocation } from '@/api/user'
import { getCurrentLocation } from '@/utils/geolocation'


export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    userInfo: null
  }),

  actions: {
    // 登录
    async login(credentials) {
      try {
        const res = await loginAPI(credentials)
        this.token = res.data.token
        localStorage.setItem('token', res.data.token)

        // 登录成功后，自动上传位置
        this.uploadLocationSilently()

        return true
      } catch (error) {
        return false
      }
    },

    // 注册
    async register(data) {
      try {
        await registerAPI(data)
        return true
      } catch (error) {
        return false
      }
    },

    //静默上传位置
    async uploadLocationSilently() {
      try {
        const location = await getCurrentLocation()
        await updateUserLocation({
          longitude: location.longitude,
          latitude: location.latitude
        })
        console.log('位置已上传:', location)
      } catch (err) {
        // 静默失败，不影响登录
        console.warn('位置上传失败:', err.message)
      }
    },

    logout() {
      this.token = '';
      this.userInfo = null;
      localStorage.removeItem('token');
    }
  }
});
