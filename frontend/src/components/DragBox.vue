<template>
  <div class="box" :class="{ dragging: dragging }" @drop.prevent="drop" @dragover.prevent="dragOver" @dragleave.prevent="dragLeave">
    <input :type="type" :accept="accept" capture="camera" class="input" @change="onChange">
    <p>{{desc}}</p>
  </div>
</template>

<script>
export default {
  name: 'BoxComponent',
  props: {
    type: { // 类型 'file'
      type: String,
      default: 'file'
    },
    accept: String, // 限制可用文件类型
    desc: String, // 描述
    drag: { // 是否可以拖拽
      type: Boolean,
      default: true,
    },
    onFileChange: { // 拖拽、文件发生改变触发
      type: Function,
      default: function () { }
    }
  },
  data() {
    return {
      dragging: false
    };
  },
  methods: {
    drop(e) {
      if (!this.drag) return;
      const file = e.dataTransfer?.items?.[0]?.getAsFile();
      if (!file) return;
      const type = file.type || "unknown";
      // 重置高亮状态
      this.dragging = false;
      this.$emit("onFileChange", { type, file });
    },
    dragOver() {
      if (!this.drag) return;
      this.dragging = true;
    },
    dragLeave() {
      if (!this.drag) return;
      this.dragging = false;
    },
    async onChange(e) {
      const file = e.target?.files?.[0];
      if (!file) return;
      const type = file.type || "unknown";
      this.$emit("onFileChange", { type, file });
    }
  },
};
</script>

<style lang="scss" scoped>
  .box {
    min-height: 160px; 
    border: 2px dashed $borderColor; 
    position: relative;
    overflow: hidden;
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 8px;

    &.dragging {
      border-color: $highlightColor;
      color: $highlightColor;
    }

    .input {
      position: absolute;
      right: 0;
      top: 0;
      width: 100%;
      height: 100%;
      opacity: 0;
      cursor: pointer;
    }
  }
</style>
