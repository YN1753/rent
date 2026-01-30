import axios from 'axios'

const request = axios.create({
  baseURL: 'https://your-api-domain.com/api',
  timeout: 5000
})

// 请求拦截器
request.interceptors.request.use(config => {
  // 在请求头中添加 token
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
}, error => Promise.reject(error));

// 响应拦截器
request.interceptors.response.use(response => {
  return response.data;
}, error => {
  console.error('API Error:', error);
  return Promise.reject(error);
});

export default request;