<template>
  <a-modal
    title="新建应用"
    :width="640"
    :visible="visible"
    :confirmLoading="loading"
    @ok="() => { $emit('ok') }"
    @cancel="() => { $emit('cancel') }"
  >
    <a-spin :spinning="loading">
      <a-form :form="form">
        <!-- 检查是否有 id 并且大于0，大于0是修改。其他是新增，新增不显示主键ID -->
        <a-form-item
          v-show="model && model.id > 0"
          label="应用ID"
          :labelCol="labelCol"
          :wrapperCol="wrapperCol">
          <a-input v-decorator="['id', { initialValue: 0 }]" disabled />
        </a-form-item>
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
    </a-spin>
  </a-modal>
</template>

<script>
import pick from 'lodash.pick'

// 表单字段
const fields = ['id', 'host', 'username', 'password', 'content', 'type']

export default {
  props: {
    visible: {
      type: Boolean,
      required: true
    },
    loading: {
      type: Boolean,
      default: () => false
    },
    model: {
      type: Object,
      default: () => null
    }
  },
  data () {
    return {
      form: this.$form.createForm(this),
      labelCol: {
        xs: { span: 24 },
        sm: { span: 7 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 13 }
      }
    }
  },
  created () {
    // 防止表单未注册
    fields.forEach(v => this.form.getFieldDecorator(v))
    // 当 model 发生改变时，为表单设置值
    this.$watch('model', () => {
      this.model && this.form.setFieldsValue(pick(this.model, fields))
    })
  }
}
</script>
