<template>
  <div :class="prefixCls">
    <div class="chart-wrapper" :style="{ height: 46 }">
      <v-chart
        :force-fit="true"
        :height="100"
        :data="dataSource"
        :scale="clusterTaskDataByDateScale"
        :padding="[36, 0, 18, 0]">
        <v-tooltip />
        <v-smooth-line position="date*task_num" :size="2" />
        <v-smooth-area position="date*task_num" />
      </v-chart>
    </div>
  </div>
</template>

<script>
// import { clusterDataByDate } from '@/api/dashboard'

export default {
  name: 'ClusterTaskNumChart',
  props: {
    prefixCls: {
      type: String,
      default: 'ant-pro-smooth-area'
    },
    dataSource: {
      type: Array,
      required: true
    }
  },
  data () {
    return {
      height: 100,
      // dataSource: [],
      clusterTaskDataByDateScale: [
        {
          dataKey: 'date',
          alias: '日期'
        },
        {
          dataKey: 'task_num',
          alias: '任务数'
        }]
    }
  }
  // methods: {
  //   async  GetClusterDataByDate () {
  //     const query = { 'day': 8 }
  //     const res = await clusterDataByDate(query)
  //     if (res) {
  //       this.dataSource = res.data
  //       console.log('dataSource == ', this.dataSource)
  //     }
  //   }
  // },
  // created () {
  //   this.GetClusterDataByDate()
  // }
}
</script>

<style lang="less" scoped>
  @import "smooth.area.less";
</style>
