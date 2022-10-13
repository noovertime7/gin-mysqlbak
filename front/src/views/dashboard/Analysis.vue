<template>
  <div>
    <a-row :gutter="24">
      <a-col :sm="24" :md="12" :xl="6" :style="{ marginBottom: '24px' }">
        <chart-card :loading="loading" title="服务数量" :total="serviceTotal">
          <a-tooltip title="集群内服务数量及在线情况" slot="action">
            <a-icon type="info-circle-o" />
          </a-tooltip>
          <div>
            <trend flag="down" style="margin-right: 16px;">
              <span slot="term">在线服务</span>
              {{ onlineService }}
            </trend>
            <trend flag="up">
              <span slot="term">离线服务</span>
              {{ offlineService }}
            </trend>
          </div>
          <template slot="footer">服务在线率<span> {{ onlineServicePercent }} %</span></template>
        </chart-card>
      </a-col>
      <a-col :sm="24" :md="12" :xl="6" :style="{ marginBottom: '24px' }">
        <chart-card :loading="loading" title="完成数量" :total="finish_total">
          <a-tooltip title="备份完成总数" slot="action">
            <a-icon type="info-circle-o" />
          </a-tooltip>
          <div>
            <trend flag="down" style="margin-right: 16px;">
              <span slot="term">mysql</span>
              {{ mysql_finish }}
            </trend>
            <trend flag="up">
              <span slot="term">elastic</span>
              {{ elastic_finish }}
            </trend>
          </div>
          <template slot="footer">完成占比<span> {{ finish_by_service }} </span></template>
        </chart-card>
      </a-col>
      <a-col :sm="24" :md="12" :xl="6" :style="{ marginBottom: '24px' }">
        <chart-card :loading="loading" title="任务数量" :total="task_total">
          <a-tooltip title="任务数量详情" slot="action">
            <a-icon type="info-circle-o" />
          </a-tooltip>
          <div>
            <trend flag="down" style="margin-right: 16px;">
              <span slot="term">mysql</span>
              {{ mysql_task }}
            </trend>
            <trend flag="up">
              <span slot="term">elastic</span>
              {{ elastic_task }}
            </trend>
          </div>
          <template slot="footer">任务占比<span> {{ task_by_service }} </span></template>
        </chart-card>
      </a-col>
      <a-col :sm="24" :md="12" :xl="6" :style="{ marginBottom: '24px' }">
        <chart-card :loading="loading" title="成功占比" :total="success_persent_str">
          <a-tooltip title="成功失败占比" slot="action">
            <a-icon type="info-circle-o" />
          </a-tooltip>
          <div>
            <mini-progress color="rgb(19, 194, 194)" :target="100" :percentage="success_persent" height="8px" />
          </div>
          <template slot="footer">
            <trend flag="down" style="margin-right: 16px;">
              <span slot="term">成功数量</span>
              {{ success_total }}
            </trend>
            <trend flag="up">
              <span slot="term">失败数量</span>
              {{ fail_total }}
            </trend>
          </template>
        </chart-card>
      </a-col>
    </a-row>

    <a-card :loading="loading" :bordered="false" :body-style="{padding: '0'}">
      <div class="salesCard">
        <a-tabs default-active-key="1" size="large" :tab-bar-style="{marginBottom: '24px', paddingLeft: '16px'}">
          <div class="extra-wrapper" slot="tabBarExtraContent">
            <div class="extra-item">
              <a>{{ $t('dashboard.analysis.all-day') }}</a>
              <a>{{ $t('dashboard.analysis.all-week') }}</a>
              <a>{{ $t('dashboard.analysis.all-month') }}</a>
              <a>{{ $t('dashboard.analysis.all-year') }}</a>
            </div>
            <a-range-picker :style="{width: '256px'}" />
          </div>
          <a-tab-pane loading="true" :tab="$t('dashboard.analysis.sales')" key="1">
            <a-row>
              <a-col :xl="16" :lg="12" :md="12" :sm="24" :xs="24">
                <bar :data="barData" :title="$t('dashboard.analysis.sales-trend')" />
              </a-col>
              <a-col :xl="8" :lg="12" :md="12" :sm="24" :xs="24">
                <rank-list :title="$t('dashboard.analysis.sales-ranking')" :list="rankList" />
              </a-col>
            </a-row>
          </a-tab-pane>
          <a-tab-pane :tab="$t('dashboard.analysis.visits')" key="2">
            <a-row>
              <a-col :xl="16" :lg="12" :md="12" :sm="24" :xs="24">
                <bar :data="barData2" :title="$t('dashboard.analysis.visits-trend')" />
              </a-col>
              <a-col :xl="8" :lg="12" :md="12" :sm="24" :xs="24">
                <rank-list :title="$t('dashboard.analysis.visits-ranking')" :list="rankList" />
              </a-col>
            </a-row>
          </a-tab-pane>
        </a-tabs>
      </div>
    </a-card>

    <div class="antd-pro-pages-dashboard-analysis-twoColLayout" :class="!isMobile && 'desktop'">
      <a-row :gutter="24" type="flex" :style="{ marginTop: '24px' }">
        <a-col :xl="12" :lg="24" :md="24" :sm="24" :xs="24">
          <a-card :loading="loading" :bordered="false" title="集群服务统计" :style="{ height: '100%' }">
            <a-row :gutter="68">
              <a-col :xs="24" :sm="12" :style="{ marginBottom: ' 24px'}">
                <number-info
                  :total="serviceNums.task_total"
                  message="七天前: "
                  :sub-total="serviceNums.task_increase_num"
                  status="down"
                  :decreaseTotal="serviceNums.task_decrease_num">
                  <span slot="subtitle">
                    <span>任务完成数</span>
                    <a-tooltip title="每天的任务完成数" slot="action">
                      <a-icon type="info-circle-o" :style="{ marginLeft: '8px' }" />
                    </a-tooltip>
                  </span>
                </number-info>
                <!-- miniChart -->
                <div>
                  <cluster-task-num-chart :dataSource="dataSource" :style="{ height: '45px' }" />
                </div>
              </a-col>
              <a-col :xs="24" :sm="12" :style="{ marginBottom: ' 24px'}">
                <number-info
                  :total="serviceNums.finish_total"
                  :sub-total="serviceNums.finish_increase_num"
                  status="down"
                  message="七天前: "
                  :decreaseTotal="serviceNums.finish_decrease_num">>
                  <span slot="subtitle">
                    <span>完成数量</span>
                    <a-tooltip :title="$t('dashboard.analysis.introduce')" slot="action">
                      <a-icon type="info-circle-o" :style="{ marginLeft: '8px' }" />
                    </a-tooltip>
                  </span>
                </number-info>
                <!-- miniChart -->
                <div>
                  <cluster-finish-num-chart :dataSource="dataSource" :style="{ height: '45px' }"/>
                </div>
              </a-col>
            </a-row>
            <div class="ant-table-wrapper">
              <a-table
                row-key="index"
                size="small"
                :columns="searchTableColumns"
                :dataSource="serviceListData"
                :pagination="{ pageSize: 5 }"
              >
                <span slot="status" slot-scope="text">
                  <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
                </span>
              </a-table>
            </div>
          </a-card>
        </a-col>
        <a-col :xl="12" :lg="24" :md="24" :sm="24" :xs="24">
          <a-card
            class="antd-pro-pages-dashboard-analysis-salesCard"
            :loading="loading"
            :bordered="false"
            title="集群服务任务占比"
            :style="{ height: '100%' }">
            <div slot="extra" style="height: inherit;">
              <!-- style="bottom: 12px;display: inline-block;" -->
              <span class="dashboard-analysis-iconGroup">
              </span>
            </div>
            <h4>任务数量</h4>
            <div>
              <!-- style="width: calc(100% - 240px);" -->
              <div>
                <PieChart :chart-data="pieData"></PieChart>
              </div>

            </div>
          </a-card>
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script>
import moment from 'moment'
import {
  ChartCard,
  MiniArea,
  MiniBar,
  MiniProgress,
  RankList,
  Bar,
  Trend,
  NumberInfo,
  MiniSmoothArea
} from '@/components'
import { baseMixin } from '@/store/app-mixin'
import { clusterDataByDate, getSvcFinishNum, getSvcTNum } from '@/api/dashboard'
import PieChart from '@/views/dashboard/components/Piechart'
import { GetServiceList } from '@/api/agent'
import ClusterTaskNumChart from '@/components/Charts/ClusterTaskNumChart'
import ClusterFinishNumChart from '@/components/Charts/ClusterFinishNumChart'
import { GetAgentTaskOverViewList } from '@/api/agent-task_overview'

const barData = []
const barData2 = []
for (let i = 0; i < 12; i += 1) {
  barData.push({
    x: `${i + 1}月`,
    y: Math.floor(Math.random() * 1000) + 200
  })
  barData2.push({
    x: `${i + 1}月`,
    y: Math.floor(Math.random() * 1000) + 200
  })
}

const rankList = []
for (let i = 0; i < 7; i++) {
  rankList.push({
    name: '白鹭岛 ' + (i + 1) + ' 号店',
    total: 1234.56 - i * 100
  })
}

const searchUserData = []
for (let i = 0; i < 7; i++) {
  searchUserData.push({
    x: moment().add(i, 'days').format('YYYY-MM-DD'),
    y: Math.ceil(Math.random() * 10)
  })
}
const statusMap = {
  0: {
    status: 'error',
    text: '离线'
  },
  1: {
    status: 'success',
    text: '在线'
  }
}
export default {
  name: 'Analysis',
  mixins: [baseMixin],
  components: {
    ClusterTaskNumChart,
    ClusterFinishNumChart,
    PieChart,
    ChartCard,
    MiniArea,
    MiniBar,
    MiniProgress,
    RankList,
    Bar,
    Trend,
    NumberInfo,
    MiniSmoothArea
  },
  filters: {
    statusFilter (type) {
      return statusMap[type].text
    },
    statusTypeFilter (type) {
      return statusMap[type].status
    }
  },
  data () {
    return {
      loading: true,
      rankList,
      // 搜索用户数
      searchUserData,
      serviceListData: [],
      // 左下角集群服务统计相关
      dataSource: [],
      serviceNums: {},
      // 左上角服务数量相关
      serviceTotal: 0,
      onlineService: 0,
      offlineService: 0,
      onlineServicePercent: 0,
      // 顶部中左完成数量数据
      taskList: [],
      finish_total: 0,
      mysql_finish: 0,
      elastic_finish: 0,
      finish_by_service: 0,
      // 顶部中右任务数量数据
      task_total: 0,
      mysql_task: 0,
      elastic_task: 0,
      task_by_service: 0,
      // 顶部最右侧成功失败占比数据
      all_total: 0,
      fail_total: 0,
      success_total: 0,
      success_persent: 0,
      success_persent_str: '',
      barData,
      barData2,
      // 饼状图数据
      pieData: [],
      // 集群服务统计
      clusterTaskDataByDate: [],
      clusterTaskDataByDateScale: [
        {
          dataKey: 'date',
          alias: '时间'
        },
        {
          dataKey: 'task_num',
          alias: '任务数',
          min: 0,
          max: 10
        }],
      searchUserScale: [
        {
          dataKey: 'x',
          alias: '时间'
        },
        {
          dataKey: 'y',
          alias: '任务数',
          min: 0,
          max: 10
        }]
    }
  },
  computed: {
    searchTableColumns () {
      return [
        {
          dataIndex: 'service_name',
          title: '服务名',
          align: 'center'
        },
        {
          dataIndex: 'content',
          title: '备注',
          align: 'center'
        },
        {
          dataIndex: 'task_num',
          title: '任务数',
          sorter: (a, b) => a.task_num - b.task_num,
          align: 'center'
        },
        {
          dataIndex: 'finish_num',
          title: '完成数',
          sorter: (a, b) => a.finish_num - b.finish_num,
          align: 'center'
        },
        {
          dataIndex: 'agent_status',
          title: '状态',
          scopedSlots: { customRender: 'status' },
          align: 'center'
        }
      ]
    }
  },
  created () {
    this.GetClusterDataByDate()
    this.GetSvcTNum()
    this.getAgentServiceList()
    this.getTaskOverview()
    this.GetSvcFinishNum()
    this.loading = !this.loading
  },
  methods: {
      async  GetClusterDataByDate () {
        const query = { 'day': 7 }
        const res = await clusterDataByDate(query)
        if (res) {
          this.dataSource = res.data.list
          this.serviceNums.task_total = res.data.task_total
          this.serviceNums.finish_total = res.data.finish_total
          this.serviceNums.task_increase_num = res.data.task_increase_num
          this.serviceNums.task_decrease_num = res.data.task_decrease_num
          this.serviceNums.finish_increase_num = res.data.finish_increase_num
          this.serviceNums.finish_decrease_num = res.data.finish_decrease_num
        }
      },
    GetSvcTNum () {
      getSvcTNum().then((res) => {
        if (res) {
          this.pieData = res.data
        }
      })
    },
    getAgentServiceList () {
      GetServiceList().then((res) => {
        if (res) {
          this.serviceListData = res.data.list
          this.serviceTotal = res.data.total
          for (let i = 0; i < this.serviceListData.length; i++) {
             if (this.serviceListData[i].agent_status === 1) {
                  this.onlineService++
             } else {
               this.offlineService++
             }
          }
          this.onlineServicePercent = Math.round(this.onlineService / this.serviceTotal * 100)
        }
      })
    },
    getTaskOverview () {
      const queryParam = { 'type': 0 }
      GetAgentTaskOverViewList(queryParam).then((res) => {
          if (res) {
            this.task_total = res.data.total
            this.taskList = res.data.list
            for (let i = 0; i < res.data.list.length; i++) {
              this.finish_total += this.taskList[i].finish_num
              if (res.data.list[i].type === 1) {
                this.mysql_finish += this.taskList[i].finish_num
                this.mysql_task++
              } else {
                this.elastic_finish += this.taskList[i].finish_num
                this.elastic_task++
              }
            }
            this.finish_by_service = Math.round(this.finish_total / this.serviceTotal)
            this.task_by_service = Math.round(this.task_total / this.serviceTotal)
          }
      })
    },
    GetSvcFinishNum () {
      getSvcFinishNum().then((res) => {
        if (res) {
          this.all_total = res.data.all_finish_total
          this.fail_total = res.data.all_fail_total
          this.success_total = this.all_total - this.fail_total
          this.success_persent = Math.round(this.success_total / this.all_total * 100)
          this.success_persent_str = this.success_persent + '%'
        }
      })
    }
  }
}
</script>

<style lang='less' scoped>
.extra-wrapper {
  line-height: 55px;
  padding-right: 24px;

  .extra-item {
    display: inline-block;
    margin-right: 24px;

    a {
      margin-left: 24px;
    }
  }
}

.antd-pro-pages-dashboard-analysis-twoColLayout {
  position: relative;
  display: flex;
  display: block;
  flex-flow: row wrap;
}

.antd-pro-pages-dashboard-analysis-salesCard {
  height: calc(100% - 24px);

  :deep(.ant-card-head) {
    position: relative;
  }
}

.dashboard-analysis-iconGroup {
  i {
    margin-left: 16px;
    color: rgba(0, 0, 0, .45);
    cursor: pointer;
    transition: color .32s;
    color: black;
  }
}

.analysis-salesTypeRadio {
  position: absolute;
  right: 54px;
  bottom: 12px;
}
</style>
