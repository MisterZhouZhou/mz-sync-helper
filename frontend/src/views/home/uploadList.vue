<template>
  <el-table
    :data="fileList"
    style="width: 100%"
    stripe
    max-height="700"
    >
    <el-table-column
      fixed
      label="文件名"
      >
      <template #default="scope">
        <el-icon v-show="normalizeType(scope?.row?.type) === 'image'" color="#409EFC" class="no-inherit">
          <PictureFilled />
        </el-icon>
        <el-icon v-show="normalizeType(scope?.row?.type) === 'file'" color="#409EFC" class="no-inherit">
          <Paperclip />
        </el-icon>
        <el-icon v-show="normalizeType(scope?.row?.type) === 'text'" color="#409EFC" class="no-inherit">
          <Tickets />
        </el-icon>
        <span style="margin-left: 10px">{{ scope.row.name }}</span>
      </template>
    </el-table-column>
    <el-table-column
      fixed="right"
      label="操作"
      width="120">
      <template #default="scope">
        <el-button
          @click.native.prevent="lookForDetail(scope.$index)"
          type="primary"
          size="small">
          查看
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import { getIpList, getSourceList } from '@/service/homeService'
import { PictureFilled, Paperclip, Tickets } from '@element-plus/icons-vue'

export default {
  name: 'UploadList',
  components: {
    [PictureFilled.name]: PictureFilled,
    [Paperclip.name]: Paperclip,
    [Tickets.name]: Tickets
  },
  data() {
    return {
      fileList: [], // 文件列表
      address: '', // ip地址
    };
  },
  computed: {
    
  },
  async mounted() {
    // 请求ip列表
    const ipRes = await getIpList();
    if (ipRes.status === 200) {
      const ipData = ipRes.data.addresses;
      if (ipData.length > 0) {
        this.address = ipData[0];
      }
    }
    // 获取文件列表
    const res = await getSourceList()
    if (res.status === 200) {
      this.fileList = res.data
      console.log('$---', this.fileList);
    }
  },
  methods: {
    lookForDetail(index) {
      const data = this.fileList[index];
      const { type, path } = data;
      this.$router.push({
        path:'/downloads',
        query: {
          type,
          url: `${encodeURIComponent(`http://${this.address}:27149${path}`)}`
        }
      });
    },
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
  }
};
</script>

<style lang="scss" scoped>

</style>
