<template>
  <h1 class="header">{{ msg }}</h1>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'HeaderView',
  props: {
    msg: String
  },
  setup() {
    
  }
})
</script>

<style lang="scss" scoped>
.header {
  margin-top: 48px;
  margin-bottom: 32px;
  text-align: center;
}
</style>
