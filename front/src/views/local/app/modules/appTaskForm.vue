<template>
  <a-form :form="form">
    <a-form-item
      label="应用地址"
      :labelCol="labelCol"
      :wrapperCol="wrapperCol"
    >
      <a-input v-decorator="['host', {rules:[{required: true, message: '请输入应用地址'}]}]" placeholder="应用地址 如: 127.0.0.1:3306"/>
    </a-form-item>
    <a-form-item
      label="用户名"
      :labelCol="labelCol"
      :wrapperCol="wrapperCol"
    >
      <a-input v-decorator="['username', {rules:[{required: true, message: '请输入用户名'}]}]" placeholder="用户名 如: root"/>
    </a-form-item>
    <a-form-item
      label="密码"
      :labelCol="labelCol"
      :wrapperCol="wrapperCol"
    >
      <a-input-password v-decorator="['password', {rules:[{required: true, message: '请输入密码'}]}]" placeholder="密码 如: password"/>
    </a-form-item>
    <a-form-item
      label="主机类型"
      :labelCol="labelCol"
      :wrapperCol="wrapperCol"
    >
      <a-select v-decorator="['type', {rules:[{required: true, message: '请选择主机类型'}]}]" placeholder="请选择主机类型">
        <a-select-option :value="1">mysql</a-select-option>
        <a-select-option :value="2">elasticSearch</a-select-option>
      </a-select>
    </a-form-item>
    <a-form-item
      label="备注"
      :labelCol="labelCol"
      :wrapperCol="wrapperCol"
    >
      <a-textarea v-decorator="['content']" placeholder="备注 任意内容"></a-textarea>
    </a-form-item>
  </a-form>
</template>

<script>
import pick from 'lodash.pick'
import { hostAdd, hostUpdate } from '@/api/host'

const fields = ['host', 'username', 'password', 'content', 'type']

export default {
  name: 'TaskForm',
  props: {
    record: {
      type: Object,
      default: null
    }
  },
  data () {
    return {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 7 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 13 }
      },
      form: this.$form.createForm(this)
    }
  },
  mounted () {
    this.record && this.form.setFieldsValue(pick(this.record, fields))
  },
  methods: {
    onOk () {
      console.log('监听了 modal ok 事件')
      if (Object.keys(this.record).length === 0) {
        this.appHandleADD()
        return new Promise(resolve => {
          resolve(true)
        })
      }
      this.appHandleEdit()
      return new Promise(resolve => {
        resolve(true)
      })
    },
    onCancel () {
      console.log('监听了 modal cancel 事件')
      return new Promise(resolve => {
        resolve(true)
      })
    },
    // 不知道为什么不触发，等后面在研究，现在先使用v-model来做了
    appHandleADD (e) {
      const { form: { validateFields } } = this
      this.visible = true
      validateFields((errors, values) => {
        if (!errors) {
          const query = {
            'host': values.host,
            'username': values.username,
            'password': values.password,
            'content': values.content,
            'type': values.type
          }
          hostAdd(query).then((res) => {
            if (res.errno === 0) {
              this.$message.success(res.data)
            }
          })
        }
      })
    },
    appHandleEdit (e) {
      const { form: { validateFields } } = this
      this.visible = true
      validateFields((errors, values) => {
        if (!errors) {
          const query = {
            'id': this.record.id,
            'host': values.host,
            'username': values.username,
            'password': values.password,
            'content': values.content,
            'type': values.type
          }
          hostUpdate(query).then((res) => {
            if (res.errno === 0) {
              this.$message.success(res.data)
            }
          })
        }
      })
    }
  }
}
</script>
