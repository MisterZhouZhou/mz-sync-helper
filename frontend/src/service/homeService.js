import http from './http';
import { getWsClient } from './wsClient';
import { clientId } from '../initializers/clientId';

// 获取上传的资源列表
export const getSourceList = () => {
  return http.get("/api/v1/uploads");
}


// 获取服务ip
export const getIpList = () => {
  return http.get("/api/v1/address");
}

// 上传文本
export const uploadText = (text) => {
  return http.post("/api/v1/texts", {
    raw: text
  }).then(r => notifyPc(r, 'text'))
}

// 上传文件
export const uploadFile = (blob) => {
  const formData = new FormData();
  formData.append("raw", blob);
  return http.post('/api/v1/files', formData, {
    "Content-Type": "multipart/form-data",
  }).then(r => notifyPc(r, 'file'))
};

// socket同步状态
const notifyPc = (response, type) => {
  // 连接socket
  getWsClient().then(c => {
    c.send({ clientId, type, url: response.data.url })
  })
  return response
}