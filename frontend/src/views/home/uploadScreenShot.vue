<template>
  <div>
    <div className="row">
      <DragBox :drag="false" type="file" accept="image/*" desc="点击选择图片 或 直接粘贴图片" @onFileChange="onFileChange" />
    </div>
    <UploadSuccessDialog v-model:visible="dialogVisible" :sourceUrl="fileUrl" :sourceType="fileType" />
  </div>
</template>

<script>
import DragBox from '@/components/DragBox.vue'
import UploadSuccessDialog from '@/components/UploadSuccessDialog.vue'
import { showLoading, hideLoading } from '@/utils/tips'
import { uploadFile } from '@/service/homeService'

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
  mounted() {
    window.addEventListener("paste", this.onPaste);
  },
  methods: {
    // 上传文件
    async uploadFile(file) {
      if (!file) return;
      const type = file.type || "unknown";
      showLoading()
      const res = await uploadFile(file);
      if (res.status === 200) {
        this.fileUrl = res.data.url;
        this.fileType = type;
        this.dialogVisible = true;
      }
      hideLoading();
    },
    // 选择、拖拽
    onFileChange(data) {
      const { file } = data;
      this.uploadFile(file)
    },
    // 粘贴事件
    onPaste(e) {
      const { items: [item] } = e.clipboardData;
      this.uploadFile(item?.getAsFile())
    },
  },
  unmounted() {
    window.removeEventListener("paste", this.onPaste);
  },
};
</script>

<style lang="scss" scoped>
  .row {
    margin: 16px 0;
  }
</style>
