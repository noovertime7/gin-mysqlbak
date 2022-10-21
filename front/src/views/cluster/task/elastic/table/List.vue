<template>
  <div>
    <div class="table-page-search-wrapper">
      <a-form layout="inline">
        <a-row :gutter="48">
          <a-col :md="8" :sm="24">
            <a-form-item label="应用地址">
              <a-select v-model="select_host" style="width: 250px" @change="handleSelectChange">
                <a-select-option :value="item.host" v-for="(item,key) in host_list" :key="key">
                  {{ item.host }}
                </a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <a-form-item label="运行状态">
              <a-select v-model="queryParam.status" placeholder="请选择" default-value="0">
                <a-select-option value="0">全部</a-select-option>
                <a-select-option value="1">停止</a-select-option>
                <a-select-option value="2">运行中</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <span
              class="table-page-search-submitButtons"
              :style="advanced && { float: 'right', overflow: 'hidden' } || {} ">
              <a-button type="primary" @click="$refs.table.refresh(true)">查询</a-button>
              <a-button style="margin-left: 8px" @click="() => queryParam = {}">重置</a-button>
            </span>
          </a-col>
        </a-row>
      </a-form>
    </div>

    <div class="table-operator">
      <a-button type="primary" icon="plus" @click="handleEdit">新建</a-button>
      <a-button type="primary" ghost="ghost" icon="rocket" @click="startBakByHost()">start all</a-button>
      <a-button type="primary" ghost="ghost" icon="poweroff" @click="stopBakByHost()">stop all</a-button>
    </div>

    <s-table
      ref="table"
      size="default"
      rowKey="key"
      :columns="columns"
      :data="loadData"
      :alert="false"
    >
      <span slot="status" slot-scope="text">
        <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
      </span>
      <span slot="serial" slot-scope="text, record, index">
        {{ index + 1 }}
      </span>
      <span slot="action" slot-scope="text, record">
        <template>
          <a @click="startTask(record)">启动</a>
          <a-divider type="vertical" />
        </template>
        <template>
          <a @click="stopBak(record)">停止</a>
          <a-divider type="vertical" />
        </template>
        <template>
          <a @click="handleEdit(record)">编辑</a>
          <a-divider type="vertical" />
        </template>
        <a-dropdown>
          <a class="ant-dropdown-link">
            更多 <a-icon type="down" />
          </a>
          <a-menu slot="overlay">
            <a-menu-item>
              <a @click="handleDetail">详情</a>
            </a-menu-item>
            <a-menu-item>
              <a @click="handleDelete(record)">删除</a>
            </a-menu-item>
          </a-menu>
        </a-dropdown>
      </span>
    </s-table>
  </div>
</template>

<script>
import { STable } from '@/components'
import {
  DeleteAgentTask,
  GetAgentTaskList,
  StartAgentHostTask,
  StopAgentHostTask
} from '@/api/agent-task'
import { GetHostNames } from '@/api/agent-host'
import { startEsTask, stopEsTask } from '@/api/elastic'

const statusMap = {
  0: {
    status: 'error',
    text: '停止'
  },
  1: {
    status: 'success',
    text: '运行中'
  }
}

export default {
  name: 'TableList',
  props: {
    // 获取edit组件的主机名
    hostByEdit: {
      type: String,
      default: ''
    }
  },
  components: {
    STable
  },
  data () {
    return {
      // 主机相关
      host_list: [],
      select_host: '',
      ghost: true,
      mdl: {},
      // 当前服务名
      service_name: '',
      // 查询参数
      queryParam: {},
      // 存储主机ID与name的对应关系
      TestMap: {},
      // 表头
      columns: [
        {
          title: '任务ID',
          dataIndex: 'id'
        },
        {
          title: '主机',
          dataIndex: 'host'
        },
        {
          title: '下次备份时间',
          dataIndex: 'backup_cycle'
        },
        {
          title: '保留天数',
          dataIndex: 'keep_number',
          customRender: (text) => text + ' 天'
        },
        {
          title: '状态',
          dataIndex: 'status',
          scopedSlots: { customRender: 'status' }
        },
        {
          title: '创建时间',
          dataIndex: 'create_at',
          sorter: true
        },
        {
          title: '操作',
          dataIndex: 'action',
          width: '220px',
          scopedSlots: { customRender: 'action' }
        }
      ],
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        if (this.select_host === '') {
          return
        }
        // 构建查询参数
        this.queryParam = { 'service_name': this.service_name, 'host_id': this.getvalue(this.select_host), 'type': 2 }
        return GetAgentTaskList(Object.assign(parameter, this.queryParam))
          .then(res => {
            return res.data
          })
      },
      selectedRowKeys: [],
      selectedRows: [],
      // custom table alert & rowSelection
      options: {
        alert: {
          show: true,
clear: () => {
            this.selectedRowKeys = []
          }
        },
        rowSelection: {
          selectedRowKeys: this.selectedRowKeys,
          onChange: this.onSelectChange
        }
      },
      optionAlertShow: false
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
    this.tableOption()
    this.service_name = this.$route.params && this.$route.params.service_name
    // 获取主机列表
    this.getHostList()
  },
  methods: {
    // 获取主机ID
    getvalue (inhost) {
      for (let i = 0; i < this.TestMap.length; i++) {
        if (this.TestMap[i].Host === inhost) {
          return this.TestMap[i].id
        }
      }
    },
    getHostList () {
      // 设置服务名与type
      const query = {
        'service_name': this.service_name,
        'type': 2
      }
      GetHostNames(query).then((res) => {
        if (res.data.list === null) {
          this.$message.warn('当前应用列表为空，请先添加Mysql或ES应用')
          this.$router.push('/cluster/app')
        }
        this.host_list = res.data.list
        if (this.hostByEdit !== '' || undefined) {
          this.select_host = this.hostByEdit
        } else {
          this.select_host = this.host_list[0].host
        }
        this.TestMap = this.host_list.map(item => ({
          id: item.host_id,
          Host: item.host
        }))
        this.$refs.table.refresh(true)
      })
    },
    tableOption () {
      if (!this.optionAlertShow) {
        this.options = {
          alert: {
            show: true,
clear: () => {
              this.selectedRowKeys = []
            }
          },
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
    handleEdit (record) {
      record.host_id_by_list = this.getvalue(this.select_host)
      this.$emit('onEdit', record)
    },
    handleOk () {

    },
    handleDetail (value) {
      this.$message.warn('正在开发中...')
    },
    handleSelectChange (value) {
      this.$refs.table.refresh(true)
    },
    startTask (record) {
      const query = {
        'task_id': record.id,
        'service_name': this.service_name
      }
      startEsTask(query).then((res) => {
        this.$message.success(res.data)
        this.$refs.table.refresh(true)
      })
    },
    startBakByHost () {
      const query = {
        'service_name': this.service_name,
        'host_id': this.getvalue(this.select_host)
      }
      StartAgentHostTask(query).then((res) => {
        if (res) {
          this.$message.success(res.data)
          this.$refs.table.refresh(true)
        }
      })
    },
    stopBakByHost () {
      const query = {
        'service_name': this.service_name,
        'host_id': this.getvalue(this.select_host)
      }
      StopAgentHostTask(query).then((res) => {
        if (res) {
          this.$message.success(res.data)
          this.$refs.table.refresh(true)
        }
      })
    },
    stopBak (record) {
      const query = {
        'task_id': record.id,
        'service_name': this.service_name
      }
      stopEsTask(query).then((res) => {
        this.$message.success(res.data)
        this.$refs.table.refresh(true)
      })
    },
    handleDelete (record) {
      const self = this
      this.$confirm({
        title: '您确认要删除此任务吗?',
        content: '删除后，任务无法通过页面管理，已启动的任务仍在后台运行',
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
              'service_name': self.service_name
            }
            DeleteAgentTask(query).then((res) => {
              self.$message.success(res.data)
              self.$refs.table.refresh()
              resolve()
            })
          })
        },
        onCancel () {
        }
      })
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    }
  }
}
</script>
