<template>
  <div :style="{ height: `${height}px`}" />
</template>

<script>
export default {
  name: 'SpaceComponent',
  props: {
    height: {
      type: Number, // 高度
      default: 32
    }
  }
};
</script>