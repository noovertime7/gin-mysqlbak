<template>
  <div>
    <v-chart :forceFit="true" :height="height" :data="data" :scale="scale">
      <v-tooltip :showTitle="false" dataKey="item*percent" />
      <v-axis />
      <v-legend dataKey="item" />
      <v-pie position="percent" color="item" :vStyle="pieStyle" :label="labelConfig" />
      <v-coord type="theta" :radius="0.75" :innerRadius="0.6" />
    </v-chart>
  </div>
</template>

<script>
const DataSet = require('@antv/data-set')
export default {
  props: {
    chartData: {
      type: Array,
      default () {
        return [
          { item: '事例一', count: 40 },
          { item: '事例二', count: 21 },
          { item: '事例三', count: 17 },
          { item: '事例四', count: 13 },
          { item: '事例五', count: 9 }
        ]
      }
    }
  },
  data () {
    return {
      data: {},
      scale: [],
      height: 400,
      pieStyle: {
        stroke: '#fff',
        lineWidth: 1
      },
      labelConfig: ['percent', {
        formatter: (val, item) => {
          return item.point.item + ': ' + val
        }
      }]
    }
  },
  created () {
    this.GetData()
  },
  watch: {
    chartData: {
      deep: true,
      handler () {
        this.GetData()
      }
    }
  },
  methods: {
    GetData () {
      const sourceData = this.chartData
       this.scale = [{
        dataKey: 'percent',
        min: 0,
        formatter: '.0%'
      }]

      const dv = new DataSet.View().source(sourceData)
      dv.transform({
        type: 'percent',
        field: 'count',
        dimension: 'item',
        as: 'percent'
      })
       this.data = dv.rows
    }
  }
}
</script>
