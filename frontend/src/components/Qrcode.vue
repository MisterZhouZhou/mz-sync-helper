<template>
  <div>
    <img v-if="imgUrl" :src="imgUrl" width="256" height="256" />
  </div>
</template>

<script>
import http from '@/service/http'

export default {
  name: 'QrcodeComponent',
  props: {
    content: String, // 二维码内容
  },
  computed: {
    imgUrl() {
      return `${http.defaults.baseURL}/api/v1/qrcode?content=${encodeURIComponent(this.content)}`
    }
  }, 
};
</script>
