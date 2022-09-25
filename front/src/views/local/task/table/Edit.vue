<template>
  <div>
    <a-form :form="form" @submit="handleSubmit">

      <a-form-item
        v-show="isEdit"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="任务ID"
        hasFeedback
      >
        <a-input
          placeholder="任务ID"
          v-decorator="[
            'id',
            {rules: [{ required: true, message: '请输入规则编号' }]}
          ]"
          :disabled="true"
        ></a-input>
      </a-form-item>

      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="数据库"
        hasFeedback
      >
        <a-input :min="1" style="width: 100%" v-decorator="['db_name', {rules: [{ required: true }]}]" />
      </a-form-item>

      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="备份周期"
        hasFeedback
      >
        <a-input style="width: 100%" v-decorator="['backup_cycle', {rules: [{ required: true }]}]" />
      </a-form-item>

      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="数据保留周期"
      >
        <a-input style="width: 100%" v-decorator="['keep_number', {rules: [{ required: true }]}]" />
      </a-form-item>

      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="功能"
        hasFeedback
      >
        <template>
          钉钉通知:&ensp;<a-switch v-model="dingStatus" checked-children="开" un-checked-children="关" default-checked />
        </template> &ensp;&ensp;&ensp;&ensp;&ensp;
        <template>
          存储设置:&ensp;<a-switch v-model="ossStatus" checked-children="开" un-checked-children="关" default-checked />
        </template>
      </a-form-item>
      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="钉钉AccessToken"
        v-show="dingStatus"
      >
        <a-input style="width: 100%" v-decorator="['ding_access_token', {rules: [{}]}]" />
      </a-form-item>
      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="钉钉Secret"
        v-show="dingStatus"
      >
        <a-input style="width: 100%" v-decorator="['ding_secret', {rules: [{}]}]" />
      </a-form-item>
      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="存储类型"
        v-show="ossStatus"
      >
        <a-select v-model="ossType" style="width: 120px">
          <a-select-option :value="0">
            AliOSS
          </a-select-option>
          <a-select-option :value="1">
            minio
          </a-select-option>
          <a-select-option :value="3">
            other
          </a-select-option>
        </a-select>
      </a-form-item>

      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="oss_access"
        v-show="ossStatus"
      >
        <a-input style="width: 100%" v-decorator="['oss_access', {rules: [{}]}]" />
      </a-form-item>
      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="oss_secret"
        v-show="ossStatus"
      >
        <a-input style="width: 100%" v-decorator="['oss_secret', {rules: [{}]}]" />
      </a-form-item>
      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="bucket_name"
        v-show="ossStatus"
      >
        <a-input style="width: 100%" v-decorator="['bucket_name', {rules: [{}]}]" />
      </a-form-item>
      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="directory"
        v-show="ossStatus"
      >
        <a-input style="width: 100%" v-decorator="['directory', {rules: [{}]}]" />
      </a-form-item>
      <a-form-item
        v-bind="buttonCol"
      >
        <a-row>
          <a-col span="6">
            <a-button type="primary" html-type="submit">提交</a-button>
          </a-col>
          <a-col span="10">
            <a-button @click="handleGoBack">返回</a-button>
          </a-col>
          <a-col span="8"></a-col>
        </a-row>
      </a-form-item>
    </a-form>
  </div>
</template>

<script>

import pick from 'lodash.pick'
import { taskDetail } from '@/api/task'

export default {
  name: 'TableEdit',
  props: {
    record: {
      type: [Object, String],
      default: ''
    }
  },
  data () {
    return {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 5 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 12 }
      },
      buttonCol: {
        wrapperCol: {
          xs: { span: 24 },
          sm: { span: 12, offset: 5 }
        }
      },
      form: this.$form.createForm(this),
      id: 0,
      isEdit: false,
      // 开关状态相关
      dingStatus: false,
      ossStatus: false,
      // 存储类型相关
      ossType: 1
    }
  },
  // beforeCreate () {
  //   this.form = this.$form.createForm(this)
  // },
  mounted () {
    this.$nextTick(() => {
      this.loadEditInfo(this.record)
    })
  },
  methods: {
    handleGoBack () {
      this.$emit('onGoBack')
    },
    handleSubmit () {
      const { form: { validateFields } } = this
      // eslint-disable-next-line handle-callback-err
      validateFields((err, values) => {
         if (this.isEdit) {
           alert(values.id)
           this.handleGoBack()
           // taskUpdate(values).then((res) => {
           //   if (res) {
           //     this.$message.success(res.data)
           //   }
           // })
         }
      })
    },
    handleGetInfo () {

    },
    loadEditInfo (data) {
      const { form } = this
      console.log('data = ', data)
      console.log(`将加载 ${this.id} 信息到表单`)
      if (data) {
        this.isEdit = true
        new Promise((resolve) => {
          resolve()
        }).then(() => {
          const query = { 'id': data.id }
          taskDetail(query).then((res) => {
            const formData = pick(res.data, ['host', 'backup_cycle', 'db_name', 'status', 'keep_number', 'id',
              'is_ding_send', 'ding_access_token', 'ding_secret',
            'is_oss_save', 'oss_type', 'endpoint', 'oss_access', 'oss_secret', 'bucket_name', 'directory'
            ])
            console.log('ding_secret = ', formData.ding_secret)
            this.dingStatus = Boolean(formData.is_ding_send)
            this.ossStatus = Boolean(formData.is_oss_save)
            // 如果开关打开,修改存储类型
            if (formData.is_oss_save !== 0) {
              this.ossStatus = formData.oss_type
            }
            form.setFieldsValue(formData)
          })
        })
      }
    }
  }
}
</script>
