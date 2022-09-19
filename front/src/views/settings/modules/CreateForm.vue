<template>
  <a-modal
    title="编辑用户"
    :width="640"
    :visible="visible"
    :confirmLoading="loading"
    @ok="() => { $emit('ok') }"
    @cancel="() => { $emit('cancel') }"
  >
    <a-spin :spinning="loading">
      <a-form :form="form" v-bind="formLayout">
        <a-form-item label="用户ID">
          <a-input v-decorator="['id']" disabled/>
        </a-form-item>
        <a-form-item label="用户名">
          <a-input v-decorator="['name']" />
        </a-form-item>
        <a-form-item
          label="用户组"
        >
          <a-select v-decorator="['group_name', {rules:[{required: true, message: '请选择用户组'}]}]">
            <a-select-option :value="1">管理员</a-select-option>
            <a-select-option :value="2">用户</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="描述">
          <a-input v-decorator="['introduction', {rules: [{required: true, min: 3, message: '请输入至少三个字符的规则描述！'}]}]" />
        </a-form-item>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>
import pick from 'lodash.pick'

// 表单字段
const fields = ['introduction', 'id', 'name', 'group_name', 'role_name', 'avatar', 'creatorId', 'login_time', 'status']

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
    this.formLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 7 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 13 }
      }
    }
    return {
      form: this.$form.createForm(this)
    }
  },
  created () {
    console.log('custom modal created')

    // 防止表单未注册
    fields.forEach(v => this.form.getFieldDecorator(v))

    // 当 model 发生改变时，为表单设置值
    this.$watch('model', () => {
      this.model && this.form.setFieldsValue(pick(this.model, fields))
    })
  }
}
</script>
