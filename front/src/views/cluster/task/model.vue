<template>
  <page-header-wrapper
    :tab-list="tabList"
    :tab-active-key="tabActiveKey"
    :tab-change="handleTabChange">
    <component :is="currentComponet"></component>
  </page-header-wrapper>
</template>

<script>
import ClusterMysqlTask from './mysql/ClusterMysqlTask'
import ClusterElasticTask from './elastic/ClusterElasticTask'
import AInput from 'ant-design-vue/lib/input/Input'
import ATextarea from 'ant-design-vue/lib/input/TextArea'

export default {
  name: 'Model',
  components: {
    AInput,
    ATextarea,
    ClusterMysqlTask,
    ClusterElasticTask
  },
  data () {
    return {
      currentComponet: 'ClusterMysqlTask',
      // 表头
      tabList: [
        { key: 'ClusterMysqlTask', tab: 'Mysql' },
        { key: 'ClusterElasticTask', tab: 'Elastic' }
      ],
      tabActiveKey: 'ClusterMysqlTask'
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
