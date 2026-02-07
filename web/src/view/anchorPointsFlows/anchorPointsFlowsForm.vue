<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="anchorUid字段:" prop="anchorUid">
          <el-input v-model="formData.anchorUid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="channelId字段:" prop="channelId">
          <el-input v-model="formData.channelId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户消费金币:" prop="coin">
          <el-input v-model.number="formData.coin" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="eventId字段:" prop="eventId">
          <el-input v-model="formData.eventId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="额外字段:" prop="ext">
       </el-form-item>
        <el-form-item label="guildId字段:" prop="guildId">
          <el-input v-model="formData.guildId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="num字段:" prop="num">
          <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="payUid字段:" prop="payUid">
          <el-input v-model="formData.payUid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="积分收益数:" prop="pointsAmount">
          <el-input v-model.number="formData.pointsAmount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="pointsChg字段:" prop="pointsChg">
          <el-input v-model.number="formData.pointsChg" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="pointsChgAfter字段:" prop="pointsChgAfter">
          <el-input v-model.number="formData.pointsChgAfter" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="pointsChgBefore字段:" prop="pointsChgBefore">
          <el-input v-model.number="formData.pointsChgBefore" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="积分流水id:" prop="pointsFlowId">
          <el-input v-model.number="formData.pointsFlowId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="兑换比例 10=10%:" prop="ratio">
          <el-input v-model.number="formData.ratio" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="source字段:" prop="source">
          <el-input v-model="formData.source" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="sourceType字段:" prop="sourceType">
          <el-input v-model="formData.sourceType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="status字段:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="subSource字段:" prop="subSource">
          <el-input v-model="formData.subSource" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createAnchorPointsFlows,
  updateAnchorPointsFlows,
  findAnchorPointsFlows
} from '@/api/anchorPointsFlows'

defineOptions({
    name: 'AnchorPointsFlowsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            anchorUid: '',
            channelId: '',
            coin: 0,
            eventId: '',
            guildId: '',
            num: 0,
            payUid: '',
            pointsAmount: 0,
            pointsChg: 0,
            pointsChgAfter: 0,
            pointsChgBefore: 0,
            pointsFlowId: 0,
            ratio: 0,
            source: '',
            sourceType: '',
            status: false,
            subSource: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAnchorPointsFlows({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reanchorPointsFlows
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAnchorPointsFlows(formData.value)
               break
             case 'update':
               res = await updateAnchorPointsFlows(formData.value)
               break
             default:
               res = await createAnchorPointsFlows(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
