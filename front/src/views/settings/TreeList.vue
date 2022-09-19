<template>
  <a-card :bordered="false">
    <a-row :gutter="8">
      <a-col :span="5">
        <s-tree
          :dataSource="orgTree"
          :openKeys.sync="openKeys"
          :search="true"
          @click="handleClick"
          @titleClick="handleTitleClick">
        </s-tree>
      </a-col>
      <a-col :span="19">
        <s-table
          ref="table"
          size="default"
          :columns="columns"
          :data="loadData"
          :alert="false"
          :rowSelection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
        >
          <span slot="status" slot-scope="text">
            <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
          </span>
          <span slot="action" slot-scope="text, record">
            <template v-if="$auth('table.update')">
              <a @click="handleEdit(record)">编辑</a>
              <a-divider type="vertical" />
            </template>
            <a-dropdown>
              <a class="ant-dropdown-link">
                更多 <a-icon type="down" />
              </a>
              <a-menu slot="overlay">
                <a-menu-item>
                  <a @click="changePasswd(record)">修改密码</a>
                </a-menu-item>
                <a-menu-item >
                  <a @click="resetPasswd(record)">重置密码</a>
                </a-menu-item>
                <a-menu-item >
                  <a-popconfirm
                    title="您确定要删除此用户吗?"
                    ok-text="Yes"
                    cancel-text="No"
                    @confirm="deleteUser(record)"
                    placement="leftTop"
                  >
                    <a>删除用户</a>
                  </a-popconfirm>
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
      </a-col>
    </a-row>
    <a-modal
      title="修改密码"
      :visible="changePwdVisible"
      :confirm-loading="changePwdConfirmLoading"
      @ok="changePwdHandleOk"
      @cancel="handleChangePwdChancel"
    >
      <template>
        <div class="components-input-demo-presuffix">
          <a-input ref="oldPwdInput" v-model="oldPwd" placeholder="原密码">
            <a-icon slot="prefix" type="lock" />
            <a-tooltip slot="suffix" title="请输入原密码">
              <a-icon type="info-circle" style="color: rgba(0,0,0,.45)" />
            </a-tooltip>
          </a-input>
          <br />
          <br />
          <a-input ref="newPwdInput" v-model="newPwd" placeholder="新密码">
            <a-icon slot="prefix" type="safety" />
            <a-tooltip slot="suffix" title="请输入新密码">
              <a-icon type="info-circle" style="color: rgba(0,0,0,.45)" />
            </a-tooltip>
          </a-input>
        </div>
      </template>
    </a-modal>
    <org-modal ref="modal" @ok="handleSaveOk" @close="handleSaveClose" />
  </a-card>
</template>

<script>
import STree from '@/components/Tree/Tree'
import { STable } from '@/components'
import OrgModal from './modules/OrgModal'
import { ChangePwd, DeleteUser, getgroupList, getUserByGroup, ResetUserPassword, UpdateUserInfo } from '@/api/user'
import CreateForm from './modules/CreateForm'

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

const groupMap = {
  '管理员': {
     id: 1
  },
  '用户': {
    id: 2
  }
}

export default {
  name: 'TreeList',
  components: {
    STable,
    STree,
    OrgModal,
    CreateForm
  },
  data () {
    return {
      openKeys: ['key-01'],

      // 查询参数
      queryParam: {},
      // 表头
      columns: [
        {
          title: '用户名',
          dataIndex: 'name',
          sorter: false
        },
        {
          title: '角色',
          dataIndex: 'role_name',
          sorter: false
        },
        {
          title: '用户组',
          dataIndex: 'group_name',
          sorter: false
        },
        {
          title: '状态',
          dataIndex: 'status',
          scopedSlots: { customRender: 'status' }
        },
        {
          title: '修改用户',
          dataIndex: 'creatorId'
        },
        {
          title: '备注',
          dataIndex: 'introduction',
          sorter: false
        },
        {
          title: '登录时间',
          dataIndex: 'login_time',
          sorter: true
        },
        {
          title: '操作',
          dataIndex: 'action',
          width: '150px',
          scopedSlots: { customRender: 'action' }
        }
      ],
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        return getUserByGroup(Object.assign(parameter, this.queryParam))
          .then(res => {
            return res.data
          })
      },
      orgTree: [],
      selectedRowKeys: [],
      selectedRows: [],
      // create model
      visible: false,
      confirmLoading: false,
      mdl: null,
      // 修改密码弹窗相关
      changePwdVisible: false,
      changePwdConfirmLoading: false,
      oldPwd: '',
      newPwd: '',
      userID: 0,
      recordName: ''
    }
  },
  created () {
    getgroupList().then(res => {
      this.orgTree.push(res.data)
    })
  },
  filters: {
    statusFilter (type) {
      return statusMap[type].text
    },
    statusTypeFilter (type) {
      return statusMap[type].status
    }
  },
  methods: {
    changePwdHandleOk (e) {
      const changeQuery = {
        'id': this.userID,
        'old_password': this.oldPwd,
        'password': this.newPwd
      }
      this.changePwdConfirmLoading = true
      ChangePwd(changeQuery).then((res) => {
        if (res) {
          this.$message.success(res.data)
          this.changePwdVisible = false
          this.changePwdConfirmLoading = false
          if (this.recordName === this.userInfo().name) {
            this.$store.dispatch('Logout')
          }
        }
        this.changePwdConfirmLoading = false
      })
    },
    handleGroupName (type) {
      if (typeof type === 'number') {
        return type
      }
      return groupMap[type].id
    },
    handleClick (e) {
      this.queryParam = {
        key: e.key
      }
      this.$refs.table.refresh(true)
    },
    // handleAdd (item) {
    //   console.log('add button, item', item)
    //   this.$message.info(`提示：你点了 ${item.key} - ${item.title} `)
    //   this.$refs.modal.add(item.key)
    // },
    handleTitleClick (item) {
      alert('title')
      console.log('handleTitleClick', item)
    },
    titleClick (e) {
      console.log('titleClick', e)
    },
    handleSaveOk () {

    },
    handleSaveClose () {

    },
    handleChangePwdChancel () {
      this.changePwdVisible = false
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
            values.group_name = this.handleGroupName(values.group_name)
            UpdateUserInfo(values).then((res) => {
              this.$message.success(res.data)
              this.$refs.table.refresh(true)
              this.visible = false
              this.confirmLoading = false
            })
          }
      })
    },
    handleEdit (record) {
      this.visible = true
      this.mdl = { ...record }
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },
    changePasswd (record) {
      this.userID = record.id
      this.recordName = record.name
      this.changePwdVisible = true
    },
    resetPasswd (record) {
      const query = {
        'id': record.id
      }
      ResetUserPassword(query).then((res) => {
        this.$message.success(res.data)
        // 如果操作的是自己，马上退出登录
       if (record.name === this.userInfo().name) {
         this.$store.dispatch('Logout')
       }
      })
    },
    deleteUser (record) {
      const query = {
        'id': record.id
      }
      DeleteUser(query).then((res) => {
        this.$message.success(res.data)
        this.$refs.table.refresh(true)
        if (record.name === this.userInfo().name) {
          this.$store.dispatch('Logout')
        }
      })
    },
    userInfo () {
      return this.$store.getters.userInfo
    }
  }
}
</script>

<style lang="less">
  .custom-tree {

    :deep(.ant-menu-item-group-title) {
      position: relative;
      &:hover {
        .btn {
          display: block;
        }
      }
    }

    :deep(.ant-menu-item) {
      &:hover {
        .btn {
          display: block;
        }
      }
    }

    :deep(.btn) {
      display: none;
      position: absolute;
      top: 0;
      right: 10px;
      width: 20px;
      height: 40px;
      line-height: 40px;
      z-index: 1050;

      &:hover {
        transform: scale(1.2);
        transition: 0.5s all;
      }
    }
  }
</style>
