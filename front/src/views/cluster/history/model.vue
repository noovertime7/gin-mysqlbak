<template>
  <page-header-wrapper
    :tab-list="tabList"
    :tab-active-key="tabActiveKey"
    :tab-change="handleTabChange">
    <component :is="currentComponet"></component>
  </page-header-wrapper>
</template>

<script>
import ClusterMysqlHistory from './mysql/ClusterMysqlHistory'
import ClusterElasticHistory from './elastic/ClusterElasticHistory'
import AInput from 'ant-design-vue/lib/input/Input'
import ATextarea from 'ant-design-vue/lib/input/TextArea'

export default {
  name: 'Model',
  components: {
    AInput,
    ATextarea,
    ClusterMysqlHistory,
    ClusterElasticHistory
  },
  data () {
    return {
      currentComponet: 'ClusterMysqlHistory',
      // 表头
      tabList: [
        { key: 'ClusterMysqlHistory', tab: 'Mysql' },
        { key: 'ClusterElasticHistory', tab: 'Elastic' }
      ],
      tabActiveKey: 'ClusterMysqlHistory'
    }
  },
  methods: {
    handleTabChange (key) {
      this.tabActiveKey = key
      this.currentComponet = key
    }
  }
}
</script>

<style scoped>

</style>
