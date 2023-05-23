<template>
  <div>
    <Center class="row">
      <textarea v-model="inputValue" cols="30" rows="10" placeholder="请输入要保存的内容" class="textarea"></textarea>
    </Center>
    <div class="center row">
      <button type="submit" class="button" @click="onSubmit">上传</button>
    </div>
    <UploadSuccessDialog v-model:visible="dialogVisible" :sourceUrl="fileUrl" :sourceType="fileType" />
  </div>
</template>

<script>
import httpT from '@/service/http.js'
import { uploadText } from '@/service/homeService.js';
import Center from '@/components/Center.vue'
import UploadSuccessDialog from '@/components/UploadSuccessDialog.vue'
import { showLoading, hideLoading } from '@/utils/tips.js'

export default {
  name: 'UploadText',
  components: {
    Center,
    UploadSuccessDialog,
  },
  data() {
    return {
      inputValue: '',
      dialogVisible: false,
      fileUrl: '', // 上传后文件的url
      fileType: '', // 上传的文件类型
      httpT,
      http2: window.location.host
    };
  },
  methods: {
    // 提交，保存文件
    async onSubmit() {
      showLoading();
      const res = await uploadText(this.inputValue);
      if (res.status === 200) {
        this.fileUrl = res.data.url;
        this.fileType = 'text';
        this.dialogVisible = true;
      }
      hideLoading();
    },
  },
}
</script>

<style lang="scss" scoped>
.row {
  margin: 16px 0;
  .textarea {
    width: 100%;
    min-height: 160px;
    line-height: 20px;
    padding: 10px;
    margin: auto;
  }
}
.center {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  .button {
    min-height: 40px;
    padding: 0 60px;
    border: 2px solid #666;
    background-color: #f5b70d;
    border-radius: 8px;
    cursor: pointer;
  }
}
</style>