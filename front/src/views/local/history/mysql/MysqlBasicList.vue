<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <a-row>
        <a-col :sm="8" :xs="24">
          <info title="完成总任务数" :value="all_nums" :bordered="true" />
        </a-col>
        <a-col :sm="8" :xs="24">
          <info title="本周完成任务数" :value="week_nums" :bordered="true" />
        </a-col>
        <a-col :sm="8" :xs="24">
          <info title="备份文件大小" :value="all_filesize" />
        </a-col>
      </a-row>
    </a-card>
    <a-card
      style="margin-top: 24px"
      :bordered="false"
      title="本地备份历史">
      <div slot="extra">
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
          <a-dropdown>
            <a class="ant-dropdown-link">
              更多 <a-icon type="down"/>
            </a>
            <a-menu slot="overlay">
              <a-menu-item>
                <a @click="handleDownLoad">下载文件</a>
              </a-menu-item>
              <a-menu-item>
                <a @click="handleRestore">还原文件</a>
              </a-menu-item>
            </a-menu>
          </a-dropdown>
        </span>
        <p slot="expandedRowRender" slot-scope="record" style="margin: 0">
          备份文件：{{ record.file_name }}
        </p>
      </s-table>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import { STable } from '@/components'
import Info from './components/Info'
import { deleteLocalHistory, getHistoryNumInfo, getLocalHistoryList } from '@/api/bak'
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
  components: {
    STable,
    Info
  },
  data () {
    return {
      radioStatus: 'all',
      searchData: '',
      // 表头
      columns: [
        {
          title: 'ID',
          dataIndex: 'id',
          sorter: true
        },
        {
          title: '应用主机',
          dataIndex: 'host'
        },
        {
          title: '库名',
          dataIndex: 'db_name'
        },
        {
          title: '文件大小',
          dataIndex: 'file_size',
          sorter: true,
          customRender: (text) => text + ' KB'
        },
        {
          title: '备份状态',
          dataIndex: 'message',
          ellipsis: true
        },
        {
          title: '存储状态',
          dataIndex: 'oss_status',
          scopedSlots: { customRender: 'status' }
        },
        {
          title: '通知状态',
          dataIndex: 'ding_status',
          scopedSlots: { customRender: 'status' }
        },
        {
          title: '备份时间',
          dataIndex: 'bak_time'
        },
        {
          title: '操作',
          width: '150px',
          dataIndex: 'action',
          scopedSlots: { customRender: 'action' }
        }
      ],
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        this.queryParam = { 'status': this.radioStatus, 'info': this.searchData }
        return getLocalHistoryList(Object.assign(parameter, this.queryParam))
          .then(res => {
            return res.data
          })
      },
      // 数量信息
      week_nums: '',
      all_nums: '',
      all_filesize: ''
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
    getHistoryNumInfo().then((res) => {
      this.week_nums = res.data.week_nums.toString()
      this.all_nums = res.data.all_nums.toString()
      this.all_filesize = res.data.all_filesize
    })
  },
  methods: {
    deleteHistory (record) {
      const deleteQuery = {
        'id': record.id
      }
      deleteLocalHistory(deleteQuery).then((res) => {
        this.$message.success(res.data)
        this.$refs.table.refresh(true)
      })
    },
    handlerSearch () {
      this.$refs.table.refresh(true)
      this.searchData = ''
    },
    handleRadioClick () {
      this.queryParam = { 'status': this.radioStatus }
      this.$refs.table.refresh(true)
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
