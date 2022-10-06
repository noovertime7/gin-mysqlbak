<template>
  <page-header-wrapper>
    <a-card title="集群应用列表">
      <div slot="extra">
        <span style="font-weight: bold">服务名:</span>&ensp;
        <a-select v-model="select_service" style="width: 230px" @change="handleSelectChange">
          <a-select-option :value="item.service_name" v-for="(item,key) in service_list" :key="key">
            {{ item.service_name }}
          </a-select-option>
        </a-select>
         &ensp;&ensp;
        <a-input-search
          @search="handlerSearch"
          v-model.trim="searchData"
          style="margin-left: 16px; width: 272px;"
          placeholder="主机名/库名" />
      </div>
      <div class="table-operator">
        <a-button type="primary" icon="plus" @click="handleAdd()">新建</a-button>
        <a-button type="primary" ghost="ghost" icon="lock" @click="startBakByHost()">锁定</a-button>
        <a-button type="primary" ghost="ghost" icon="unlock" @click="stopBakByHost()">解锁</a-button>
      </div>
      <s-table
        ref="table"
        size="default"
        :columns="columns"
        :data="loadData"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
      >
        <span slot="host_status" slot-scope="text">
          <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
        </span>

        <span slot="host_type" slot-scope="text">
          <a-tag :color="text | colorType">
            {{ text | textTypeFilter }}
          </a-tag>
        </span>
        <span slot="action" slot-scope="text, record">
          <a @click="handleEdit(record)">编辑</a>
          <a-divider type="vertical" />
          <a-dropdown>
            <a class="ant-dropdown-link">
              更多 <a-icon type="down" />
            </a>
            <a-menu slot="overlay">
              <a-menu-item>
                <a style="color: red" @click="handleDelete(record)">删除</a>
              </a-menu-item>
              <a-menu-item>
                <a @click="handleTest(record)">测试</a>
              </a-menu-item>
            </a-menu>
          </a-dropdown>
        </span>
      </s-table>
      <create-form
        ref="createModal"
        :visible="visible"
        :loading="confirmLoading"
        :model="mdl"
        @cancel="handleCancel"
        @ok="handleOk"
      />
    </a-card>
  </page-header-wrapper>
</template>

<script>
import CreateForm from './modules/CreateForm'
import { STable } from '@/components'
import { CreateAgentHost, DeleteAgentHost, GetAgentHostList, TestAgentHost, UpdateAgentHost } from '@/api/agent-host'
import { GetServiceList } from '@/api/agent'

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
    STable,
    CreateForm
  },
  filters: {
    statusFilter (type) {
      return statusMap[type].text
    },
    statusTypeFilter (type) {
      return statusMap[type].status
    },
    textTypeFilter (type) {
      return typeMap[type].text
    },
    colorType (type) {
      return typeMap[type].color
    }
  },
  computed: {
    hasSelected () {
      return this.selectedRowKeys.length > 0
    }
  },
  data () {
    return {
      // model相关
      // create model
      visible: false,
      confirmLoading: false,
      mdl: null,
      // 选择框相关
      selectedRowKeys: [],
      columns: [
        {
          title: '主机ID',
          dataIndex: 'id',
          align: 'center'
        },
        {
          title: '主机名',
          dataIndex: 'host',
          align: 'center'
        },
        {
          title: '状态',
          dataIndex: 'host_status',
          scopedSlots: { customRender: 'host_status' },
          align: 'center'
        },
        {
          title: '任务数',
          dataIndex: 'task_num',
          align: 'center'
        },
        {
          title: '类型',
          dataIndex: 'type',
          scopedSlots: { customRender: 'host_type' },
          align: 'center'
        },
        {
          title: '创建时间',
          dataIndex: 'create_at',
          align: 'center'
        },
        {
          title: '备注',
          dataIndex: 'content',
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
      // 服务相关
      service_list: [],
      select_service: '',
      // 搜索框绑定参数
      searchData: '',
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        if (this.select_service === '') {
          return
        }
        this.queryParam = { 'info': this.searchData, 'service_name': this.select_service }
        return GetAgentHostList(Object.assign(parameter, this.queryParam)).then(res => {
          return res.data
        })
      }
    }
  },
  created () {
    this.getServiceList()
  },
  methods: {
    handleAdd () {
      this.mdl = null
      this.visible = true
    },
    handleDelete (record) {
      const self = this
      this.$confirm({
        title: '您确认要删除此主机吗?',
        content: '删除前请手动删除主机下所有任务',
        destroyOnClose: true,
        onOk () {
          return new Promise((resolve, reject) => {
            if (record.task_num !== 0) {
              self.$message.warn('当前主机有任务未删除，请删除后重试')
              resolve()
              return
            }
            const query = {
              'service_name': self.select_service,
              'id': record.id
            }
            DeleteAgentHost(query).then((res) => {
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
    handleTest (record) {
     const query = {
       'service_name': this.select_service,
       'host_id': record.id
     }
      TestAgentHost(query).then((res) => {
        if (res) {
          this.$message.success(res.data)
        }
      })
    },
    handleEdit (record) {
      this.visible = true
      this.mdl = { ...record }
    },
    getServiceList () {
      GetServiceList().then((res) => {
        if (res.data.list.length === 0) {
          this.$message.error('当前没有服务注册，页面加载失败!')
        }
        this.service_list = res.data.list
        this.select_service = this.service_list[0].service_name
        this.$refs.table.refresh(true)
      })
    },
    handlerSearch () {
      this.$refs.table.refresh(true)
      this.searchData = ''
    },
    handleSelectChange (value) {
      this.$refs.table.refresh(true)
    },
    start () {
      this.loading = true
      // ajax request after empty completing
      setTimeout(() => {
        this.loading = false
        this.selectedRowKeys = []
      }, 1000)
    },
    onSelectChange (selectedRowKeys) {
      console.log('selectedRowKeys changed: ', selectedRowKeys)
      this.selectedRowKeys = selectedRowKeys
    },
    handleCancel () {
      this.visible = false
      const form = this.$refs.createModal.form
      form.resetFields() // 清理表单数据（可不做）
    },
    handleOk () {
      const form = this.$refs.createModal.form
      this.confirmLoading = true
      form.validateFields((errors, values) => {
        if (!errors) {
          console.log('values', values)
          // 获取serviceName
          values['service_name'] = this.select_service
          if (values.id > 0) {
            // 修改 e.g.
            new Promise((resolve, reject) => {
              UpdateAgentHost(values).then((res) => {
                if (res) {
                  resolve(res)
                } else {
                  reject(Error('修改失败'))
                }
              })
            }).then(res => {
              // 刷新表格
              this.$refs.table.refresh()
              this.$message.success('修改成功')
            }).finally(() => {
              this.visible = false
              this.confirmLoading = false
              // 重置表单数据
              form.resetFields()
              }
            )
          } else {
            // 新增
            new Promise((resolve, reject) => {
              CreateAgentHost(values).then((res) => {
                if (res) {
                  resolve(res.data)
                } else {
                  reject(Error('添加失败'))
                }
              })
            }).then(res => {
              this.visible = false
              this.confirmLoading = false
              // 重置表单数据
              form.resetFields()
              // 刷新表格
              this.$refs.table.refresh()
              this.$message.success(res)
            }).finally(() => {
              this.visible = false
              this.confirmLoading = false
              // 重置表单数据
              form.resetFields()
            })
          }
        } else {
          this.confirmLoading = false
        }
      })
    }
  }
}
</script>
