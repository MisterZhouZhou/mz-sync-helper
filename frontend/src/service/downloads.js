import http from './http';

// 获取服务ip
export const getTextContent = (url) => {
  return http.get(url);
}