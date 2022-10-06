<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <a-row>
        <a-col :sm="8" :xs="24">
          <info title="服务总数" :value="all_services" :bordered="true" />
        </a-col>
        <a-col :sm="8" :xs="24">
          <info title="任务总数" :value="all_tasks" :bordered="true" />
        </a-col>
        <a-col :sm="8" :xs="24">
          <info title="完成总数" :value="all_finish_tasks" />
        </a-col>
      </a-row>
    </a-card>
    <a-card :bordered="false" style="margin-top: 24px">
      <s-table
        ref="table"
        size="default"
        :columns="columns"
        :data="loadData"
      >
        <span slot="status" slot-scope="text">
          <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
        </span>
        <span slot="action" slot-scope="text, record">
          <a @click="handleTask(record)">任务</a>
          <a-divider type="vertical"/>
          <a-dropdown>
            <a class="ant-dropdown-link">
              更多 <a-icon type="down"/>
            </a>
            <a-menu slot="overlay">
              <a-menu-item>
                <a @click="handleDeleteService">删除</a>
              </a-menu-item>
              <a-menu-item>
                <a @click="handleRestart">重启</a>
              </a-menu-item>
              <a-menu-item>
                <a @click="handleDownload">上传</a>
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
import { GetServiceList, GetServiceNumInfo } from '@/api/agent'
import Info from './components/Info'

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
  components: {
    STable,
    Info
  },
  data () {
    return {
      columns: [
        {
          title: '服务名',
          dataIndex: 'service_name',
          align: 'center'
        },
        {
          title: '地址',
          dataIndex: 'address',
          align: 'center'
        },
        {
          title: '任务数',
          dataIndex: 'task_num',
          // customRender: (text) => text + ' 个',
          align: 'center'
        },
        {
          title: '完成数',
          dataIndex: 'finish_num',
          // customRender: (text) => text + ' 个',
          align: 'center'
        },
        {
          title: '上次注册时间',
          dataIndex: 'last_time',
          align: 'center'
        },
        {
          title: '创建时间',
          dataIndex: 'create_at',
          align: 'center'
        },
        {
          title: '状态',
          dataIndex: 'agent_status',
          scopedSlots: { customRender: 'status' },
          align: 'center'
        },
        {
          title: '备注',
          dataIndex: 'content',
          align: 'center'
        },
        {
          title: '操作',
          dataIndex: 'action',
          scopedSlots: { customRender: 'action' },
          align: 'center'
        }
      ],
      // 任务数信息相关
      all_services: '',
      all_tasks: '',
      all_finish_tasks: '',
      // 查询条件参数
      queryParam: {},
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        return GetServiceList(Object.assign(parameter, this.queryParam)).then(res => {
          return res.data
        })
      }
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
  mounted () {
    this.GetServiceNum()
  },
  methods: {
    taskManage (row) {
      // axios 发送数据到后端 修改数据成功后
      // 调用 refresh() 重新加载列表数据
      // 这里 setTimeout 模拟发起请求的网络延迟..
      setTimeout(() => {
        this.$refs.table.refresh() // refresh() 不传参默认值 false 不刷新到分页第一页
      }, 1500)
    },
    handleTask (record) {
      this.$router.push('/cluster/app/task-list/' + record.service_name)
    },
    handleDownload () {
      this.$message.warn('功能开发中...')
    },
    handleRestart () {
      this.$message.warn('功能开发中...')
    },
    handleDeleteService () {
      this.$message.warn('功能开发中...')
    },
    GetServiceNum () {
      GetServiceNumInfo().then((res) => {
        this.all_finish_tasks = res.data.all_finish_tasks.toString()
        this.all_services = res.data.all_services.toString()
        this.all_tasks = res.data.all_tasks.toString()
      })
    }
  }
}
</script>
