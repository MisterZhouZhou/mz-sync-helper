<template>
  <div class="container" :class="{verical: verical}">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: 'CenterComponent',
  props: {
    verical: Boolean, // 是否垂直方向, 默认水平方向
  },
};
</script>

<style lang="scss" scoped>
  .container {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;

    &.verical {
      flex-direction: column;
    }
  }
</style>
