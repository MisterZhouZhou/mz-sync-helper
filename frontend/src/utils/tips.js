import 'element-plus/es/components/loading/style/css'
import { ElLoading } from 'element-plus'

let loading;

export const showLoading = () => {
  loading = ElLoading.service({
    lock: true,
    text: 'Loading',
    background: 'rgba(0, 0, 0, 0.7)'
  });
}

export const hideLoading = () => {
  loading.close()
}

export default {}