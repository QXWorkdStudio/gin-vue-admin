<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户名:" prop="accountName">
          <el-input v-model="formData.accountName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="银行账号、钱包、其他账户标识:" prop="accountNumber">
          <el-input v-model="formData.accountNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="变动的积分:" prop="addPoints">
          <el-input v-model.number="formData.addPoints" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="兑换金额:" prop="amount">
          <el-input v-model="formData.amount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="国家:" prop="countryName">
          <el-input v-model="formData.countryName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="币种:" prop="currency">
          <el-input v-model="formData.currency" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="描述:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="额外信息汇总；:" prop="extInfo">
       </el-form-item>
        <el-form-item label="主渠道:" prop="mainChannel">
          <el-input v-model="formData.mainChannel" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商户id:" prop="merchantId">
       </el-form-item>
        <el-form-item label="审核人id:" prop="operatorId">
          <el-input v-model.number="formData.operatorId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商户订单号:" prop="orderNo">
          <el-input v-model="formData.orderNo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="修改后积分:" prop="pointsAfter">
          <el-input v-model.number="formData.pointsAfter" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="修改前积分余额:" prop="pointsBefore">
          <el-input v-model.number="formData.pointsBefore" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产品id:" prop="productId">
          <el-input v-model.number="formData.productId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="对应订单状态 提现状态 1:待处理 2:待打款 10:打款失败 11:后台拒绝 21:成功:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="子渠道; bank_code:" prop="subChannel">
          <el-input v-model="formData.subChannel" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户uid:" prop="uid">
          <el-input v-model="formData.uid" :clearable="true" placeholder="请输入" />
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
  createExchangePointsOrder,
  updateExchangePointsOrder,
  findExchangePointsOrder
} from '@/api/exchangePointsOrder'

defineOptions({
    name: 'ExchangePointsOrderForm'
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
            accountName: '',
            accountNumber: '',
            addPoints: 0,
            amount: '',
            countryName: '',
            currency: '',
            desc: '',
            mainChannel: '',
            operatorId: 0,
            orderNo: '',
            pointsAfter: 0,
            pointsBefore: 0,
            productId: 0,
            status: 0,
            subChannel: '',
            uid: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findExchangePointsOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reexchangePointsOrder
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
               res = await createExchangePointsOrder(formData.value)
               break
             case 'update':
               res = await updateExchangePointsOrder(formData.value)
               break
             default:
               res = await createExchangePointsOrder(formData.value)
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
