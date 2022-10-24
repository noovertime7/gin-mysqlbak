<template>
  <div>
    <div class="table-page-search-wrapper">
      <a-form layout="inline">
        <a-row :gutter="48">
          <a-col :md="8" :sm="24">
            <a-form-item label="数据库">
              <a-input v-model="queryParam.info" placeholder=""/>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <a-form-item label="使用状态">
              <a-select v-model="queryParam.status" placeholder="请选择" default-value="0">
                <a-select-option value="0">全部</a-select-option>
                <a-select-option value="1">停止</a-select-option>
                <a-select-option value="2">运行中</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <span class="table-page-search-submitButtons" :style="advanced && { float: 'right', overflow: 'hidden' } || {} ">
              <a-button type="primary" @click="$refs.table.refresh(true)">查询</a-button>
              <a-button style="margin-left: 8px" @click="() => queryParam = {}">重置</a-button>
            </span>
          </a-col>
        </a-row>
      </a-form>
    </div>

    <div class="table-operator">
      <a-button type="primary" icon="plus" @click="handleEdit()">新建</a-button>
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
              <a href="javascript:;">详情</a>
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
import { taskDelete, taskList } from '@/api/task'
import { startAllBakByHost, startBak, stopAllBakByHost, stopBak } from '@/api/bak'

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
  components: {
    STable
  },
  data () {
    return {
      ghost: true,
      mdl: {},
      // 高级搜索 展开/关闭
      advanced: false,
      // 查询参数
      queryParam: {},
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
          title: '数据库',
          dataIndex: 'db_name'
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
        this.queryParam['host_id'] = this.hostID
        return taskList(Object.assign(parameter, this.queryParam))
          .then(res => {
            return res.data
          })
      },
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
      // 查询相关
      hostID: 0
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
    this.hostID = this.$route.params && this.$route.params.hostID
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
    handleEdit (record) {
      this.$emit('onEdit', record)
    },
    handleOk () {

    },
    startTask (record) {
      const query = {
        'id': record.id,
        'host_id': this.hostID
      }
      startBak(query).then((res) => {
        this.$message.success(res.data)
        this.$refs.table.refresh()
      })
    },
    startBakByHost () {
      const query = {
        'host_id': this.hostID
      }
      startAllBakByHost(query).then((res) => {
        if (res) {
          this.$message.success(res.data)
          this.$refs.table.refresh()
        }
      })
    },
    stopBakByHost () {
      const query = {
        'host_id': this.hostID
      }
      stopAllBakByHost(query).then((res) => {
        if (res) {
          this.$message.success(res.data)
          this.$refs.table.refresh()
        }
      })
    },
    stopBak (record) {
      const query = {
        'id': record.id,
        'host_id': this.hostID
      }
      stopBak(query).then((res) => {
        this.$message.success(res.data)
        this.$refs.table.refresh()
      })
    },
    delete (record) {
      const query = {
        'id': record.id
      }
      taskDelete(query).then((res) => {
        this.$success(res.data)
        this.$refs.table.refresh()
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
                'id': record.id
              }
              taskDelete(query).then((res) => {
                self.$message.success(res.data)
                self.$refs.table.refresh(true)
                resolve()
              })
            })
          },
          onCancel () {}
        })
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    }
  }
}
</script>
