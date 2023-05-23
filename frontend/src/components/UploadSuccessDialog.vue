<template>
  <el-dialog v-model="props.visible" :show-close="false" :width="isMobileBo ? '90%' : '50%'">
    <div class="pop">
      <h2 class="title"></h2>
      <div v-if="ipAddressList.length > 0">
        <p>
          请 Windows 用户在防火墙入站规则中开通 27149 端口（<a href="https://jingyan.baidu.com/article/09ea3ede7311dec0afde3977.html" target="_blank" rel="noreferrer">教程</a>）
        </p>
        <p>
          <label class="ip-wrap">
            <span class="ip-title">请选择手机可以访问的局域网IP</span>
            <select :value="address" @change="onChange">
              <option value="" disabled>
                - 请选择 -
              </option>
              <option v-for="(item, index) in ipAddressList" :key="index">{{ item }}</option>
            </select>
          </label>
        </p>
      </div>
      <!-- 二维码 -->
      <div class="item">
        <Qrcode :content="downLoadUrl" />
      </div>
      <div class="item">
        <a :href="encodeURI(downLoadUrl)">请 手机扫码 或 点击下载</a>
      </div>
      <div class="space" />
      <div class="item">
        <button class="close" @click="onClose">关闭</button>
      </div>
    </div>
  </el-dialog>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue';
import { getIpList } from '@/service/homeService'
import Qrcode from '@/components/Qrcode.vue'
import { isMobile } from '@/utils/device'

const emits = defineEmits(['close', 'update:visible'])

const props = defineProps({
  visible: Boolean, // dialog是否显示
  sourceUrl: String, // 下载文件的url
  sourceType: String, // 下载文件的类型
  close: Function, // 弹窗关闭事件
})

const ipAddressList = ref([])
const address = ref(localStorage.getItem("address") || '')
const isMobileBo = isMobile()

const downLoadUrl = computed(() => {
  let downUrl
  if (process.env.NODE_ENV === 'development') { // 开发环境
    downUrl = `http://${address.value}:8000/#/downloads?type=${props.sourceType}&url=${encodeURIComponent(`http://${address.value}:27149${props.sourceUrl}`)}`;
  } else if (process.env.NODE_ENV === 'production') { // 生产环境
    downUrl = `http://${address.value}:27149/static/#/downloads?type=${props.sourceType}&url=${encodeURIComponent(`http://${address.value}:27149${props.sourceUrl}`)}`;
  }
  return downUrl
})

const onChange = (e) => {
  address.value = e.target.value;
  localStorage.setItem("address", e.target.value);
}

const onClose = () => {
  emits("close")
  emits("update:visible", false);
}

onMounted(async () => {
  // 请求ip列表
  const res = await getIpList();
  if (res && res.status === 200) {
    ipAddressList.value = res.data.addresses;
    if (ipAddressList.value.length > 0) {
      address.value = ipAddressList.value[0];
    }
  }
})
</script>

<style lang="scss" scoped>
.pop {
  margin-top: -30px;
  .title {
    font-weight: bold;
    font-size: 24px;
    margin-bottom: 16px;
  }
  .ip-wrap {
    display: flex;
    padding: 4px 0; 
    justify-content: flex-start;
    align-items: center;
    min-height: 40px;
    white-space: nowrap;
    .ip-title {
      margin-right: 8px;
    }
  }
  .item {
    display: flex;
    flex-direction: 'column';
    justify-content: center;
    align-items: center;
  }
  .space {
    height: 10px;
  }
  .close {
    min-height: 40px;
    padding: 0 60px;
    border: 2px solid #666;
    background: #f5b70d;
    border-radius: 8px;
    cursor: pointer;
  }
}
</style>
