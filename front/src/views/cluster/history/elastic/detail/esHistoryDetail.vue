<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <a-descriptions title="主机信息">
        <a-descriptions-item label="主机ID">{{ es_host_detail.host_id }}</a-descriptions-item>
        <a-descriptions-item label="主机名">{{ es_host_detail.host }}</a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ es_host_detail.create_at }}</a-descriptions-item>
        <a-descriptions-item label="主机状态">
          <a-tag :color="es_host_detail.status | statusFilter">
            {{ es_host_detail.status | statusTextFilter }}
          </a-tag>
        </a-descriptions-item>
        <a-descriptions-item label="更新时间">{{ es_host_detail.update_at }}</a-descriptions-item>
      </a-descriptions>
      <a-divider style="margin-bottom: 32px" />
      <a-descriptions title="任务信息">
        <a-descriptions-item label="任务ID">{{ es_history_detail.task_id }}</a-descriptions-item>
        <a-descriptions-item label="备份周期">{{ es_task_detail.backup_cycle }}</a-descriptions-item>
        <a-descriptions-item label="保留周期">{{ es_task_detail.keep_number }} 天</a-descriptions-item>
        <a-descriptions-item label="任务状态">
          <a-tag :color="es_task_detail.status | statusFilter">
            {{ es_task_detail.status| statusTextFilter }}
          </a-tag>
        </a-descriptions-item>
        <a-descriptions-item label="创建时间"> {{ es_task_detail.create_at }}</a-descriptions-item>
      </a-descriptions>
      <a-divider style="margin-bottom: 32px" />
      <a-descriptions title="快照信息" bordered >
        <a-descriptions-item label="ID">{{ es_history_detail.id }}</a-descriptions-item>
        <a-descriptions-item label="UUID">{{ es_history_detail.uuid }}</a-descriptions-item>
        <a-descriptions-item label="仓库" > {{ es_history_detail.repository }}</a-descriptions-item>
        <a-descriptions-item label="状态" >
          <a-tag :color="es_history_detail.status| statusFilter">
            {{ es_history_detail.status| statusHistoryFilter }}
          </a-tag>
        </a-descriptions-item>
        <a-descriptions-item label="消耗时间">{{ es_history_detail.duration_in_millis }} ms</a-descriptions-item>
        <a-descriptions-item label="快照名">{{ es_history_detail.snapshot }}</a-descriptions-item>
        <a-descriptions-item label="开始时间" :span="2" > {{ es_history_detail.start_time }}</a-descriptions-item>
        <a-descriptions-item label="结束时间" :span="2" > {{ es_history_detail.end_time }}</a-descriptions-item>
        <a-descriptions-item label="备份索引" :span="1">
          <a-tooltip>
            <template slot="title">
              {{ es_history_detail.indices }}
            </template>
            <a-tag color="blue">
              查看备份索引
            </a-tag>
          </a-tooltip>
        </a-descriptions-item>
      </a-descriptions>
    </a-card>
  </page-header-wrapper>
</template>

<script>

import { getEsHistoryDetail } from '@/api/elastic'

export default {
  filters: {
    statusFilter (status) {
      const statusMap = {
        '1': 'green',
        '0': 'red'
      }
      return statusMap[status]
    },
    statusTextFilter (status) {
      const statusTextMap = {
        '1': '运行中',
        '0': '停止'
      }
      return statusTextMap[status]
    },
    statusHistoryFilter (status) {
      const statusTextMap = {
        '1': '备份成功',
        '0': '备份失败'
      }
      return statusTextMap[status]
    }
  },
  data () {
    return {
      computed: {
        title () {
          return this.$route.meta.title
        }
      },
      service_name: '',
      id: 0,
      // 数据相关
      es_host_detail: {},
      es_task_detail: {},
      es_history_detail: {}
    }
  },
  created () {
    this.id = this.$route.params.id
    this.service_name = this.$route.params.service_name
    this.GetData()
  },
  methods: {
    GetData () {
      const query = { 'service_name': this.service_name, 'id': this.id }
      getEsHistoryDetail(query).then((res) => {
        this.es_history_detail = res.data.es_history_detail
        this.es_task_detail = res.data.es_task_detail
        this.es_host_detail = res.data.es_host_detail
      })
    }
  }
}
</script>

<style lang='less' scoped>
.title {
  color: rgba(0, 0, 0, .85);
  font-size: 16px;
  font-weight: 500;
  margin-bottom: 16px;
}
</style>
