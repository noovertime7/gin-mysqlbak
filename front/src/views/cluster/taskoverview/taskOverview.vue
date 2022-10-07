<template>
  <page-header-wrapper
    content="任务总览: 定时同步集群内所有服务任务，可以方便快捷的启动停止任务。"
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
      <s-table
        ref="table"
        size="default"
        :columns="columns"
        :data="loadData"
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
          <a @click="edit(record)">启动</a>
          <a-divider type="vertical"/>
          <a @click="edit(record)">停止</a>
          <a-divider type="vertical"/>
          <a-dropdown>
            <a class="ant-dropdown-link">
              更多 <a-icon type="down"/>
            </a>
            <a-menu slot="overlay">
              <a-menu-item>
                <a href="javascript:;">还原</a>
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
import { GetAgentTaskOverViewList } from '@/api/agent-task_overview'

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
  data () {
    return {
      extraImage: 'https://gw.alipayobjects.com/zos/rmsportal/RzwpdLnhmvDJToTdfDPe.png',
      select_type: 0,
      searchData: '',
      radioStatus: 0,
      columns: [
        {
          title: 'ID',
          dataIndex: 'id',
          align: 'center',
          sorter: true,
          needTotal: true
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
          align: 'center'
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
    edit (row) {
      // axios 发送数据到后端 修改数据成功后
      // 调用 refresh() 重新加载列表数据
      // 这里 setTimeout 模拟发起请求的网络延迟..
      setTimeout(() => {
        this.$refs.table.refresh() // refresh() 不传参默认值 false 不刷新到分页第一页
      }, 1500)
    },
    handlerSearch () {
      this.$refs.table.refresh(true)
    },
    handleSelectChange (value) {
      this.$refs.table.refresh(true)
    },
    handleRadioClick () {
      this.$refs.table.refresh(true)
    }
  }
}
</script>
