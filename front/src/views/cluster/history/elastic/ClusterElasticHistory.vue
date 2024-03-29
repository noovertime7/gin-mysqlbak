<template>
  <div>
    <a-card :bordered="false">
      <a-row>
        <a-col :sm="8" :xs="24">
          <info title="完成总任务数" :value="all_nums" :bordered="true" />
        </a-col>
        <a-col :sm="8" :xs="24">
          <info title="本周完成任务数" :value="week_nums" :bordered="true" />
        </a-col>
        <a-col :sm="8" :xs="24">
          <info title="失败任务数" :value="fail_nums" :redColor="fail_nums !== '0'" />
        </a-col>
      </a-row>
    </a-card>
    <a-card
      style="margin-top: 24px"
      :bordered="false"
      title="Es快照历史">
      <div slot="extra">
        服务名：
        <a-select v-model="select_service" style="width: 230px" @change="handleSelectChange">
          <a-select-option :value="item.service_name" v-for="(item,key) in service_list" :key="key">
            {{ item.service_name }}
          </a-select-option>
        </a-select>
        &ensp;&ensp;
        <a-radio-group v-model="radioStatus" @change="handleRadioClick">
          <a-radio-button value="all">全部</a-radio-button>
          <a-radio-button value="success">成功</a-radio-button>
          <a-radio-button value="fail">失败</a-radio-button>
        </a-radio-group>
        <a-input-search @search="handlerSearch" v-model="searchData" style="margin-left: 16px; width: 272px;" placeholder="主机名/库名"/>
      </div>
      <s-table
        ref="table"
        size="default"
        :columns="columns"
        :data="loadData"
      >
        <a-input-search style="margin-left: 16px; width: 272px;" />
        <span slot="status" slot-scope="text">
          <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
        </span>
        <span slot="action" slot-scope="text, record">
          <a @click="deleteHistory(record)">删除</a>
          <a-divider type="vertical"/>
          <a @click="handleDetail(record)">详情</a>
        </span>
        <p slot="expandedRowRender" slot-scope="record" style="margin: 0">
          UUID: {{ record.uuid }}
        </p>
      </s-table>
    </a-card>
  </div>
</template>

<script>
import { STable } from '@/components'
import Info from './components/Info'
import { GetServiceList } from '@/api/agent'
import { deleteEsHistory, getEsHistoryInfo, getEsHistoryList } from '@/api/elastic'
const statusMap = {
  0: {
    status: 'default',
    text: '失败'
  },
  1: {
    status: 'success',
    text: '成功'
  },
  2: {
    status: 'processing',
    text: '未启用'
  }
}

export default {
  name: 'ClusterElasticHistory',
  components: {
    STable,
    Info
  },
  data () {
    return {
      radioStatus: 'all',
      searchData: '',
      columns: [
        {
          title: 'ID',
          dataIndex: 'id',
          sorter: true,
          width: '80px',
          align: 'center'
        },
        {
          title: '应用主机',
          dataIndex: 'host',
          width: '200px',
          align: 'center'
        },
        {
          title: '备份仓库',
          dataIndex: 'repository',
          align: 'center',
          width: '100px'
        },
        {
          title: '快照名',
          dataIndex: 'snapshot',
          width: '180px',
          align: 'center'
        },
        {
          title: '备份状态',
          dataIndex: 'status',
          scopedSlots: { customRender: 'status' },
          width: '100px',
          align: 'center'
        },
        {
          title: '备注',
          dataIndex: 'message',
          align: 'center'
        },
        {
          title: '备份时间',
          dataIndex: 'start_time',
          width: '250px',
          align: 'center'
        },
        {
          title: '操作',
          width: '150px',
          dataIndex: 'action',
          scopedSlots: { customRender: 'action' },
          align: 'center'
        }
      ],
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        if (this.select_service === '') {
          return
        }
        this.queryParam = { 'status': this.radioStatus, 'info': this.searchData, 'service_name': this.select_service }
        return getEsHistoryList(Object.assign(parameter, this.queryParam))
          .then(res => {
            return res.data
          })
      },
      // 数量信息
      week_nums: '',
      all_nums: '',
      fail_nums: '',
      // 服务相关
      service_list: [],
      select_service: ''
    }
  },
  filters: {
    statusFilter (type) {
      return statusMap[type].text
    },
    statusTypeFilter (type) {
      return statusMap[type].status
    }
  },
  created () {
    this.getServiceList()
  },
  methods: {
    getMysqlHistoryNunInfo () {
      const query = { 'service_name': this.select_service }
      getEsHistoryInfo(query).then((res) => {
        this.week_nums = res.data.week_nums.toString()
        this.all_nums = res.data.all_nums.toString()
        this.fail_nums = res.data.fail_nums.toString()
        this.$refs.table.refresh()
      })
    },
    getServiceList () {
      GetServiceList().then((res) => {
        if (res.data.list.length === 0) {
          this.$message.error('当前没有服务注册，页面加载失败!')
        }
        this.service_list = res.data.list
        this.select_service = this.service_list[0].service_name
        this.getMysqlHistoryNunInfo()
      })
    },
    handleSelectChange (value) {
      this.getMysqlHistoryNunInfo()
      this.$refs.table.refresh()
    },
    handleDetail (value) {
      this.$router.push('/cluster/history/detail/' + value.id + '/' + this.select_service)
    },
    deleteHistory (record) {
      const deleteQuery = {
        'id': record.id,
        'service_name': this.select_service
      }
      deleteEsHistory(deleteQuery).then((res) => {
        this.$message.success(res.data)
        this.$refs.table.refresh()
      })
    },
    handlerSearch () {
      this.$refs.table.refresh()
      this.searchData = ''
    },
    handleRadioClick () {
      this.queryParam = { 'status': this.radioStatus }
      this.$refs.table.refresh()
    },
    handleDownLoad () {
      this.$message.warn('功能正在开发中...')
    },
    handleRestore () {
      this.$message.warn('功能正在开发中...')
    }
  }
}
</script>
