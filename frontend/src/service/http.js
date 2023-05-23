import axios from "axios";
import 'element-plus/es/components/message/style/css'
import { ElMessage } from 'element-plus';

const http = axios.create({
  timeout: 5000,
});

// 响应拦截
http.interceptors.response.use((res) => {
  // console.log('$---', res);
  if (res.data.code || res.data.status) {
      // code
      const resCode = String(res.data.code || res.data.status);
      // 错误提示
      if (resCode !== '200') {
        ElMessage.error(res.data.msg);
        return;
      }
      // 请求成功
      return res.data || {};
    } else {
      ElMessage.error(res.data.msg || '接口错误');
    }
  },
  (err) => {
    // status 非 200 时
    if (err.response && err.response.data) {
      ElMessage.error(err.response.data.msg);
    }
  }
);

let hostUrl = window.location.hostname.includes('wails') ? '127.0.0.1' : window.location.hostname
http.defaults.baseURL = `http://${hostUrl}:27149`

export default http;