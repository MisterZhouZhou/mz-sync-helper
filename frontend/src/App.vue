<template>
  <div>
    <!-- <button @click="greet">ddd</button>
    {{ name }} -->
    <router-view></router-view>
  </div>
</template>

<!-- <script lang="ts" setup>
  import { ref } from 'vue';
  // 只能在APP中使用
  import {Greet} from '../wailsjs/go/main/App'

  const name = ref('')
  function greet() {
    Greet(Math.random()+'').then(result => {
      name.value = result
    })
  }
</script> -->