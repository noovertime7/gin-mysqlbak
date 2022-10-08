<template>
  <page-header-wrapper
    content="任务总览: 定时同步集群内所有服务任务，可以方便快捷的启动停止任务，也可以还原被删除任务。"
  >
    <template v-slot:extraContent>
      <div style="width: 120px; margin-top: -20px;"><img style="width: 100%" :src="extraImage" /></div>
    </template>
    <a-card :bordered="false" title="任务总览">
      <div slot="extra">
        <span style="font-weight: bold">应用:</span>&ensp;
        <a-select v-model="select_type" style="width: 230px" @change="handleSelectChange">
          <a-select-option :value="0">全部</a-select-option>
          <a-select-option :value="1">mysql</a-select-option>
          <a-select-option :value="2">elasticSearch</a-select-option>
        </a-select>
        &ensp;&ensp;
        <a-radio-group v-model="radioStatus" @change="handleRadioClick">
          <a-radio-button :value="0">全部</a-radio-button>
          <a-radio-button :value="2">运行</a-radio-button>
          <a-radio-button :value="1">停止</a-radio-button>
        </a-radio-group>
        <a-input-search
          @search="handlerSearch"
          v-model.trim="searchData"
          style="margin-left: 16px; width: 272px;"
          placeholder="主机名/数据库/服务名" />
      </div>
      <div class="table-operator">
        <a-button type="primary" ghost="ghost" icon="sync" @click="handleSync">手动同步</a-button>
        <a-button type="primary" ghost="ghost" icon="sync" :disabled="getSelectStatus" @click="handleSync">批量启动</a-button>
        <a-button type="primary" ghost="ghost" icon="sync" :disabled="getSelectStatus" @click="handleSync">批量停止</a-button>
      </div>
      <s-table
        ref="table"
        size="default"
        :columns="columns"
        :data="loadData"
        :alert="options.alert"
        :rowSelection="options.rowSelection"
      >
        <span slot="host_status" slot-scope="text">
          <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
        </span>

        <span slot="deleted_status" slot-scope="text">
          <a-badge :status="text | deletedStatusTypeFilter" :text="text | deletedStatusFilter" />
        </span>

        <span slot="host_type" slot-scope="text">
          <a-tag :color="text | colorType">
            {{ text | textTypeFilter }}
          </a-tag>
        </span>
        <span slot="action" slot-scope="text, record">
          <a @click="handleStart(record)">启动</a>
          <a-divider type="vertical"/>
          <a @click="handleStop(record)">停止</a>
          <a-divider type="vertical"/>
          <a-dropdown>
            <a class="ant-dropdown-link">
              更多 <a-icon type="down"/>
            </a>
            <a-menu slot="overlay">
              <a-menu-item>
                <a @click="handleRestore(record)">还原</a>
              </a-menu-item>
              <a-menu-item>
                <a @click="handleDelete(record)">删除</a>
              </a-menu-item>
            </a-menu>
          </a-dropdown>
        </span>
      </s-table>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import { STable } from '@/components'
import overViewImg from '@/assets/overview.png'
import {
  deleteOverviewBak,
  GetAgentTaskOverViewList, restoreOverviewBak,
  startOverviewBak,
  stopOverviewBak, syncOverviewBak
} from '@/api/agent-task_overview'

const statusMap = {
  0: {
    status: 'default',
    text: '停止'
  },
  1: {
    status: 'success',
    text: '运行'
  }
}

const DeletedStatusMap = {
  1: {
    status: 'error',
    text: '已删除'
  },
  0: {
    status: 'success',
    text: '正常'
  }
}

const typeMap = {
  1: {
    color: '#4248ff',
    text: 'Mysql'
  },
  2: {
    color: '#1dcdf0',
    text: 'ElasticSearch'
  }
}
export default {
  components: {
    STable
  },
  filters: {
    statusFilter (type) {
      return statusMap[type].text
    },
    statusTypeFilter (type) {
      return statusMap[type].status
    },
    deletedStatusFilter (type) {
      return DeletedStatusMap[type].text
    },
    deletedStatusTypeFilter (type) {
      return DeletedStatusMap[type].status
    },
    textTypeFilter (type) {
      return typeMap[type].text
    },
    colorType (type) {
      return typeMap[type].color
    }
  },
  computed: {
    // 控制开关状态，只有选中行才会返回false
    getSelectStatus () {
      return this.selectedRows.length < 1
    }
  },
  data () {
    return {
      selectedRowKeys: [],
      selectedRows: [],

      // custom table alert & rowSelection
      options: {
        alert: { show: true, clear: () => { this.selectedRowKeys = [] } },
        rowSelection: {
          selectedRowKeys: this.selectedRowKeys,
          onChange: this.onSelectChange
        }
      },
      optionAlertShow: false,
      extraImage: overViewImg,
      select_type: 0,
      searchData: '',
      radioStatus: 0,
      columns: [
        {
          title: 'ID',
          dataIndex: 'id',
          align: 'center',
          sorter: true
        },
        {
          title: '服务名',
          dataIndex: 'service_name',
          align: 'center'
        },
        {
          title: '主机',
          dataIndex: 'host',
          align: 'center'
        },
        {
          title: '数据库',
          dataIndex: 'db_name',
          align: 'center'
        },
        {
          title: '下次备份时间',
          dataIndex: 'backup_cycle',
          align: 'center',
          sorter: true
        },
        {
          title: '完成数',
          dataIndex: 'finish_num',
          align: 'center',
          sorter: true,
          needTotal: true
        },
        {
          title: '保留周期',
          dataIndex: 'keep_number',
          customRender: (text) => text + ' 天',
          align: 'center',
          sorter: true
        },
        {
          title: '任务状态',
          dataIndex: 'is_deleted',
          scopedSlots: { customRender: 'deleted_status' },
          align: 'center',
          sorter: true
        },
        {
          title: '运行状态',
          dataIndex: 'status',
          scopedSlots: { customRender: 'host_status' },
          align: 'center'
        },
        {
          title: '类型',
          dataIndex: 'type',
          scopedSlots: { customRender: 'host_type' },
          align: 'center'
        },
        {
          table: '操作',
          dataIndex: 'action',
          scopedSlots: { customRender: 'action' },
          align: 'center'
        }
      ],
      // 查询条件参数
      queryParam: {},
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        this.queryParam = { 'status': this.radioStatus, 'info': this.searchData, 'type': this.select_type }
        return GetAgentTaskOverViewList(Object.assign(parameter, this.queryParam)).then(res => {
          return res.data
        })
      }
    }
  },
  methods: {
    tableOption () {
      if (!this.optionAlertShow) {
        this.options = {
          alert: { show: true, clear: () => { this.selectedRowKeys = [] } },
          rowSelection: {
            selectedRowKeys: this.selectedRowKeys,
            onChange: this.onSelectChange
          }
        }
        this.optionAlertShow = true
      } else {
        this.options = {
          alert: false,
          rowSelection: null
        }
        this.optionAlertShow = false
      }
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },
    handlerSearch () {
      this.$refs.table.refresh(true)
    },
    handleSelectChange (value) {
      this.$refs.table.refresh(true)
    },
    handleRadioClick () {
      this.$refs.table.refresh(true)
    },
    handleRestore (record) {
      const query = {
        'id': record.id,
        'service_name': record.service_name,
        'task_id': record.task_id,
        'type': record.type
      }
      restoreOverviewBak(query).then((res) => {
        if (res) {
          this.$message.success(res.data)
          this.$refs.table.refresh(true)
        }
      })
    },
    handleStart (record) {
      const query = {
        'id': record.id,
        'service_name': record.service_name,
        'task_id': record.task_id,
        'type': record.type
      }
      startOverviewBak(query).then((res) => {
        if (res) {
          this.$message.success(res.data)
          this.$refs.table.refresh(true)
        }
      })
    },
    handleStop (record) {
      const query = {
        'id': record.id,
        'service_name': record.service_name,
        'task_id': record.task_id,
        'type': record.type
      }
      stopOverviewBak(query).then((res) => {
        if (res) {
          this.$message.success(res.data)
          this.$refs.table.refresh(true)
        }
      })
    },
    handleDelete (record) {
      const self = this
      this.$confirm({
        title: '您确认要删除此任务吗?',
        content: '删除前请手动停止任务',
        destroyOnClose: true,
        onOk () {
          return new Promise((resolve, reject) => {
            if (record.status === 1) {
              self.$message.warn('任务运行中，请停止后重试')
              resolve()
              return
            }
            const query = {
              'id': record.id,
              'service_name': record.service_name,
              'task_id': record.task_id,
              'type': record.type
            }
            deleteOverviewBak(query).then((res) => {
              if (res) {
                self.$message.success(res.data)
                self.$refs.table.refresh(true)
                resolve()
              }
            })
          })
        },
        onCancel () {}
      })
    },
    handleSync () {
      syncOverviewBak().then((res) => {
        if (res) {
          this.$message.success(res.data)
          this.$refs.table.refresh(true)
        }
      })
    }
  }
}
</script>
