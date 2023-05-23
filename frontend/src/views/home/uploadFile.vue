<template>
  <div>
    <div className="row">
      <DragBox type="file" desc="点击打开文件 或 拖拽文件至此" @onFileChange="onFileChange" />
    </div>
    <UploadSuccessDialog v-model:visible="dialogVisible" :sourceUrl="fileUrl" :sourceType="fileType" />
  </div>
</template>

<script>
import DragBox from '@/components/DragBox.vue'
import { uploadFile } from '@/service/homeService'
import { showLoading, hideLoading } from '@/utils/tips'
import UploadSuccessDialog from '@/components/UploadSuccessDialog.vue'

export default {
  name: 'UploadText',
  components: {
    DragBox,
    UploadSuccessDialog,
  },
  data() {
    return {
      dialogVisible: false,
      fileUrl: '', // 上传后文件的url
      fileType: '', // 上传的文件类型
    };
  },
  methods: {
    async onFileChange(data) {
      const { file, type } = data;
      showLoading();
      const res = await uploadFile(file);
      if (res.status === 200) {
        this.fileUrl = res.data.url;
        this.fileType = type;
        this.dialogVisible = true;
      }
      hideLoading();
    },
  },
};
</script>

<style lang="scss" scoped>
  .row {
    margin: 16px 0;
  }
</style>
