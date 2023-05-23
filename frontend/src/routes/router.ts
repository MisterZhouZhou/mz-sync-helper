import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import Home from '../views/home/home'
import UploadText from '../views/home/uploadText.vue'
import UploadFile from '../views/home/uploadFile.vue'
import UploadScreenShot from '../views/home/uploadScreenShot.vue'
import UploadList from '../views/home/uploadList.vue'
import Downloads from '../views/downloads.vue'

const routes = [
  {
    path: '/',
    component: Home,
    redirect: '/message', // 一级默认路由
    children: [
      {
        path: 'message',
        name: 'uploadText',
        component: UploadText,
      },
      {
        path: 'file',
        name: 'uploadFile',
        component: UploadFile,
      },
      {
        path: 'screenshot',
        name: 'uploadScreenShot',
        component: UploadScreenShot,
      },
      {
        path: 'uploadList',
        name: 'uploadList',
        component: UploadList,
      }
    ]
  },
  {
    path: '/downloads',
    name: 'downloads',
    component: Downloads,
  }
] as RouteRecordRaw[];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;