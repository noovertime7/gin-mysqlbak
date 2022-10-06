<template>
  <global-footer class="footer custom-render">
    <template v-slot:links>
      <a href="https://github.com/noovertime7/gin-mysqlbak" target="_blank">Github</a>
      <a href="https://github.com/noovertime7" target="_blank">@noovertime7</a>
      <a href="https://github.com/noovertime7/gin-mysqlbak-agent" target="_blank">agent</a>
    </template>
    <!--    <template v-slot:copyright>-->
    <!--      <a href="https://github.com/vueComponent" target="_blank">vueComponent</a>-->
    <!--    </template>-->
  </global-footer>
</template>

<script>
import { GlobalFooter } from '@ant-design-vue/pro-layout'

export default {
  name: 'ProGlobalFooter',
  components: {
    GlobalFooter
  }
}
</script>
