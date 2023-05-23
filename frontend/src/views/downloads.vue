<template>
  <div>
    <header-view msg="同步传" />
    <!-- text -->
    <Center v-if="type === 'text'" verical>
      <textarea :value="inputValue" readOnly placeholder="请输入要保存的内容" class="textarea"></textarea>
      <Space />
      <Center virtical>
        <button class="button">请手动复制上面文本</button>
      </Center>
    </Center>
    <!-- file -->
    <Center v-if="type === 'file'" verical>
      <a :href="url">
        <Center>
          <button class="button">点击下载文件</button>
        </Center>
      </a>
    </Center>
    <!-- image -->
    <Center  v-if="type === 'image'">
      <a :href="url" class="img-wrap">
        <img :src="url" class="image" />
        <Center>
          <button class="button">长按或点击，即可下载图片</button>
        </Center>
      </a>
    </Center>
    <Space />
    <Center>
      <button class="button" @click="onClickUpload">我也要上传</button>
    </Center>
  </div>
</template>

<script>
import HeaderView from '@/components/Header.vue'
import Center from '@/components/Center.vue'
import Space from '@/components/Space.vue'
import http from 'axios'

export default {
  name: 'DownloadPage',
  components: {
    HeaderView,
    Center,
    Space,
  },
  data() {
    return {
      inputValue: '',
      type: '', // 文件类型
      url: '', // 文件url
    };
  },
  async mounted() {
    // 从query中获取
    const { type, url } = this.$route.query;
    this.type = this.normalizeType(type);
    this.url = decodeURIComponent(url);
    // 获取文本内容
    if (this.type === "text") {
      const res = await http.get(this.url)
      if (res.status === 200) {
        this.inputValue = res.data;
      }
    }
  },
  methods: {
    // 匹配文件类型
    normalizeType(type) {
      if (/^image\/.*$/.test(type)) {
        return 'image'
      } else if (/^text(\/.*)?$/.test(type)) {
        return 'text'
      } else {
        return 'file'
      }
    },
    onClickUpload() {
      this.$router.push("/")
    }
  }
};
</script>

<style lang="scss" scoped>
.textarea {
  width: 96%;
  min-height: 160px;
  line-height: 20px;
  padding: 10px;
}
.button {
  min-height: 40px;
  padding: 0 60px;
  border: 2px solid #666;
  background: #f5b70d;
  border-radius: 8px;
  cursor: pointer;
}
.img-wrap {
  text-align: center;
  .image {
    border: 2px solid #333;
    margin: 16px;
    width: 200px;
    height: 300px;
  }
}
</style>
