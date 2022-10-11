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
        <v-smooth-line position="date*finish_num" :size="2" />
        <v-smooth-area position="date*finish_num" />
      </v-chart>
    </div>
  </div>
</template>

<script>

export default {
  name: 'ClusterFinishNumChart',
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
      clusterTaskDataByDateScale: [
        {
          dataKey: 'date',
          alias: '日期'
        },
        {
          dataKey: 'finish_num',
          alias: '完成数'
        }]
    }
  }
}
</script>

<style lang="less" scoped>
  @import "smooth.area.less";
</style>
