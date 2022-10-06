<template>
  <page-header-wrapper>
    <a-card
      style="margin-top: 24px"
      :bordered="false"
      title="应用管理">

      <div slot="extra">
        <a-input-search style="margin-left: 16px; width: 272px;" />
      </div>

      <div class="operate">
        <a-button type="dashed" style="width: 100%" icon="plus" @click="add">添加</a-button>
      </div>

      <a-list size="large" :pagination="paginationProps" :loading="listLoading">
        <a-list-item :key="index" v-for="(item, index) in data">
          <a-list-item-meta :description="item.content">
            <a-avatar slot="avatar" size="large" shape="square" :src="item.avatar"/>
            <a slot="title" @click="handleTaskList(item)">{{ item.host }}</a>
          </a-list-item-meta>
          <div slot="actions">
            <a @click="edit(item)">编辑</a>
          </div>
          <div slot="actions">
            <a @click="handleTaskList(item)">任务</a>
          </div>
          <div slot="actions">
            <a style="color: red" @click="handleDelete(item)">删除</a>
          </div>
          <div class="list-content">
            <div class="list-content-item">
              <span>状态</span>
              <p>
                <a-badge :status="item.host_status | statusTypeFilter" :text="item.host_status | statusFilter"/>
              </p>
            </div>
            <div class="list-content-item">
              <span>任务数</span>
              <p>{{ item.task_num }}</p>
            </div>
          </div>
        </a-list-item>
      </a-list>
    </a-card>
  </page-header-wrapper>
</template>

<script>
// 演示如何使用 this.$dialog 封装 modal 组件
import TaskForm from './modules/appTaskForm'
import Info from './components/Info'
import { hostDelete, hostList } from '@/api/host'

const statusMap = {
  0: {
    status: 'default',
    text: '离线'
  },
  1: {
    status: 'success',
    text: '在线'
  },
  2: {
    status: 'processing',
    text: '运行中'
  },
  3: {
    status: 'error',
    text: '异常'
  }
}

export default {
  name: 'StandardList',
  components: {
    TaskForm,
    Info
  },
  data () {
    return {
      data: [],
      total: 0,
      pageNo: 1,
      PageSize: 10,
      listLoading: false
    }
  },
  mounted () {
    this.getListData()
  },
  filters: {
    statusFilter (type) {
      return statusMap[type].text
    },
    statusTypeFilter (type) {
      return statusMap[type].status
    }
  },
  computed: {
    paginationProps: function () {
      return {
        showSizeChanger: true,
        showQuickJumper: true,
        current: this.pageNo,
        total: this.total,
        onChange: this.paginationChange,
        loading: this.loading
      }
    }
  },
  methods: {
    getListData () {
      console.log('get data')
      this.listLoading = true
      const query = {
        'page_no': this.pageNo,
        'page_size': this.PageSize
      }
      hostList(query).then((res) => {
        this.data = res.data.list
        this.total = res.data.total
       this.listLoading = false
      })
    },
    handleTaskList (record) {
      this.$router.push('/local/app/task-list/' + record.id)
    },
    handleDelete (record) {
        const query = {
          'id': record.id
        }
      hostDelete(query).then((res) => {
        if (res.errno === 0) {
          this.$message.success(res.data)
          this.getListData()
        }
      })
    },
    paginationChange (page, pageSize) {
      this.pageNo = page
      this.PageSize = pageSize
      this.getListData()
    },
    add () {
      const self = this
      this.$dialog(TaskForm,
        {
          record: {},
          on: {
            ok: () => {
              return new Promise((resolve, reject) => {
                self.getListData()
                self.getListData()
                resolve()
              })
            },
            cancel () {
              console.log('cancel 回调')
            },
            close () {
              console.log('modal close 回调')
            },
            get () {
              console.log('get get get ')
            }
          }
        },
        // modal props
        {
          title: '新增',
          width: 700,
          centered: true,
          maskClosable: false
        })
    },
    edit (record) {
      console.log('record', record)
      this.$dialog(TaskForm,
        // component props
        {
          record,
          on: {
            ok: () => {
              return new Promise((resolve, reject) => {
                self.getListData()
                resolve()
              })
            },
            cancel () {
              console.log('cancel 回调')
            },
            close () {
              console.log('modal close 回调')
            }
          }
        },
        // modal props
        {
          title: '编辑',
          width: 700,
          centered: true,
          maskClosable: false
        })
    }
  }
}
</script>

<style lang="less" scoped>
.ant-avatar-lg {
    width: 48px;
    height: 48px;
    line-height: 48px;
}

.list-content-item {
    color: rgba(0, 0, 0, .45);
    display: inline-block;
    vertical-align: middle;
    font-size: 14px;
    margin-left: 40px;
    span {
        line-height: 20px;
    }
    p {
        margin-top: 4px;
        margin-bottom: 0;
        line-height: 22px;
    }
}
</style>
