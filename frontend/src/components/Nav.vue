<template>
  <nav class="nav">
    <ul>
      <li>
        <router-link to="/message" active-class="selected">
          传消息
        </router-link>
      </li>
      <li>
        <router-link to="/file" active-class="selected">
          传文件
        </router-link>
      </li>
      <li>
        <router-link to="/screenshot" active-class="selected">
          传截图
        </router-link>
      </li>
      <li>
        <router-link to="/uploadList" active-class="selected">
          文件列表
        </router-link>
      </li>
    </ul>
  </nav>
</template>

<script>
export default {
  name: 'NavComponent',
  data() {
    return {

    };
  },
};
</script>

<style lang="scss" scoped>
  .nav {
    text-align: center;
    ul {
      display: flex;
      border-top: 1px solid $borderColor;
      border-left: 1px solid $borderColor;
      border-right: 1px solid $borderColor;
      list-style: none;
      padding: 0;
      li {
        flex-grow: 1;
        border-bottom: 1px solid $borderColor;
        a {
          display: block;
          padding: 8px 0;
          text-decoration: none;
          
          &.selected {
            background-color: $highlightColor;
          }
        }
      }
    }
  }
</style>
