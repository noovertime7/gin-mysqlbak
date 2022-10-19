<template>
  <a-card :bordered="false">
    <component
      @onEdit="handleEdit"
      @onGoBack="handleGoBack"
      :hostByEdit="hostByEdit"
      :host="hostID"
      :auto="auto"
      :record="record"
      :is="currentComponet"></component>
  </a-card>
</template>

<script>

import ATextarea from 'ant-design-vue/es/input/TextArea'
import AInput from 'ant-design-vue/es/input/Input'
// 动态切换组件
import List from './table/List'
import Edit from './table/Edit'

export default {
  name: 'ClusterMysqlTask',
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
      hostID: 0,
      auto: false,
      hostByEdit: ''
    }
  },
  methods: {
    handleEdit (record, isAuto) {
      this.auto = isAuto
      this.record = record || ''
      this.hostByEdit = record.host
      this.hostID = record.host_id_by_list
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
