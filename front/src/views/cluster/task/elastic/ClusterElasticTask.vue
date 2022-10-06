<template>
  <a-card :bordered="false">
    <component @onEdit="handleEdit" @onGoBack="handleGoBack" :record="record" :host="HostID" :is="currentComponet"></component>
  </a-card>
</template>

<script>

import ATextarea from 'ant-design-vue/es/input/TextArea'
import AInput from 'ant-design-vue/es/input/Input'
// 动态切换组件
import List from './table/List'
import Edit from './table/Edit'

export default {
  name: 'ClusterElasticTask',
  components: {
    AInput,
    ATextarea,
    List,
    Edit
  },
  data () {
    return {
      currentComponet: 'List',
      record: '',
      HostID: 0
    }
  },
  mounted () {
    const temp = this.$route.params && this.$route.params.hostID
    this.HostID = Number(temp)
  },
  methods: {
    handleEdit (record) {
      this.record = record || ''
      this.currentComponet = 'Edit'
    },
    handleGoBack () {
      this.record = ''
      this.currentComponet = 'List'
    }
  },
  watch: {
    '$route.path' () {
      this.record = ''
      this.currentComponet = 'List'
    }
  }
}
</script>
