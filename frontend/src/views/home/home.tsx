import { defineComponent } from "vue"
import HeaderView from '@/components/Header.vue'
import NavView from '@/components/Nav.vue'

export default defineComponent({
  name: 'HomePage',
  render() {
    return (
      <>
        <HeaderView msg="同步传" />
        <NavView />
        <router-view></router-view>
      </>
    );
  }
})