<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="用户名" prop="accountName" width="120" />
        <el-table-column align="left" label="银行账号、钱包、其他账户标识" prop="accountNumber" width="120" />
        <el-table-column align="left" label="变动的积分" prop="addPoints" width="120" />
        <el-table-column align="left" label="兑换金额" prop="amount" width="120" />
        <el-table-column align="left" label="国家" prop="countryName" width="120" />
        <el-table-column align="left" label="币种" prop="currency" width="120" />
        <el-table-column align="left" label="描述" prop="desc" width="120" />
        <el-table-column align="left" label="额外信息汇总；" prop="extInfo" width="120" />
        <el-table-column align="left" label="主渠道" prop="mainChannel" width="120" />
        <el-table-column align="left" label="商户id" prop="merchantId" width="120" />
        <el-table-column align="left" label="审核人id" prop="operatorId" width="120" />
        <el-table-column align="left" label="商户订单号" prop="orderNo" width="120" />
        <el-table-column align="left" label="修改后积分" prop="pointsAfter" width="120" />
        <el-table-column align="left" label="修改前积分余额" prop="pointsBefore" width="120" />
        <el-table-column align="left" label="产品id" prop="productId" width="120" />
        <el-table-column align="left" label="对应订单状态 提现状态 1:待处理 2:待打款 10:打款失败 11:后台拒绝 21:成功" prop="status" width="120" />
        <el-table-column align="left" label="子渠道; bank_code" prop="subChannel" width="120" />
        <el-table-column align="left" label="用户uid" prop="uid" width="120" />
        <el-table-column align="left" label="操作" min-width="120">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateExchangePointsOrderFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="用户名:"  prop="accountName" >
              <el-input v-model="formData.accountName" :clearable="true"  placeholder="请输入用户名" />
            </el-form-item>
            <el-form-item label="银行账号、钱包、其他账户标识:"  prop="accountNumber" >
              <el-input v-model="formData.accountNumber" :clearable="true"  placeholder="请输入银行账号、钱包、其他账户标识" />
            </el-form-item>
            <el-form-item label="变动的积分:"  prop="addPoints" >
              <el-input v-model.number="formData.addPoints" :clearable="true" placeholder="请输入变动的积分" />
            </el-form-item>
            <el-form-item label="兑换金额:"  prop="amount" >
              <el-input v-model="formData.amount" :clearable="true"  placeholder="请输入兑换金额" />
            </el-form-item>
            <el-form-item label="国家:"  prop="countryName" >
              <el-input v-model="formData.countryName" :clearable="true"  placeholder="请输入国家" />
            </el-form-item>
            <el-form-item label="币种:"  prop="currency" >
              <el-input v-model="formData.currency" :clearable="true"  placeholder="请输入币种" />
            </el-form-item>
            <el-form-item label="描述:"  prop="desc" >
              <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入描述" />
            </el-form-item>
            <el-form-item label="额外信息汇总；:"  prop="extInfo" >
            </el-form-item>
            <el-form-item label="主渠道:"  prop="mainChannel" >
              <el-input v-model="formData.mainChannel" :clearable="true"  placeholder="请输入主渠道" />
            </el-form-item>
            <el-form-item label="商户id:"  prop="merchantId" >
            </el-form-item>
            <el-form-item label="审核人id:"  prop="operatorId" >
              <el-input v-model.number="formData.operatorId" :clearable="true" placeholder="请输入审核人id" />
            </el-form-item>
            <el-form-item label="商户订单号:"  prop="orderNo" >
              <el-input v-model="formData.orderNo" :clearable="true"  placeholder="请输入商户订单号" />
            </el-form-item>
            <el-form-item label="修改后积分:"  prop="pointsAfter" >
              <el-input v-model.number="formData.pointsAfter" :clearable="true" placeholder="请输入修改后积分" />
            </el-form-item>
            <el-form-item label="修改前积分余额:"  prop="pointsBefore" >
              <el-input v-model.number="formData.pointsBefore" :clearable="true" placeholder="请输入修改前积分余额" />
            </el-form-item>
            <el-form-item label="产品id:"  prop="productId" >
              <el-input v-model.number="formData.productId" :clearable="true" placeholder="请输入产品id" />
            </el-form-item>
            <el-form-item label="对应订单状态 提现状态 1:待处理 2:待打款 10:打款失败 11:后台拒绝 21:成功:"  prop="status" >
              <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入对应订单状态 提现状态 1:待处理 2:待打款 10:打款失败 11:后台拒绝 21:成功" />
            </el-form-item>
            <el-form-item label="子渠道; bank_code:"  prop="subChannel" >
              <el-input v-model="formData.subChannel" :clearable="true"  placeholder="请输入子渠道; bank_code" />
            </el-form-item>
            <el-form-item label="用户uid:"  prop="uid" >
              <el-input v-model="formData.uid" :clearable="true"  placeholder="请输入用户uid" />
            </el-form-item>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="用户名">
                        {{ formData.accountName }}
                </el-descriptions-item>
                <el-descriptions-item label="银行账号、钱包、其他账户标识">
                        {{ formData.accountNumber }}
                </el-descriptions-item>
                <el-descriptions-item label="变动的积分">
                        {{ formData.addPoints }}
                </el-descriptions-item>
                <el-descriptions-item label="兑换金额">
                        {{ formData.amount }}
                </el-descriptions-item>
                <el-descriptions-item label="国家">
                        {{ formData.countryName }}
                </el-descriptions-item>
                <el-descriptions-item label="币种">
                        {{ formData.currency }}
                </el-descriptions-item>
                <el-descriptions-item label="描述">
                        {{ formData.desc }}
                </el-descriptions-item>
                <el-descriptions-item label="额外信息汇总；">
                        {{ formData.extInfo }}
                </el-descriptions-item>
                <el-descriptions-item label="主渠道">
                        {{ formData.mainChannel }}
                </el-descriptions-item>
                <el-descriptions-item label="商户id">
                        {{ formData.merchantId }}
                </el-descriptions-item>
                <el-descriptions-item label="审核人id">
                        {{ formData.operatorId }}
                </el-descriptions-item>
                <el-descriptions-item label="商户订单号">
                        {{ formData.orderNo }}
                </el-descriptions-item>
                <el-descriptions-item label="修改后积分">
                        {{ formData.pointsAfter }}
                </el-descriptions-item>
                <el-descriptions-item label="修改前积分余额">
                        {{ formData.pointsBefore }}
                </el-descriptions-item>
                <el-descriptions-item label="产品id">
                        {{ formData.productId }}
                </el-descriptions-item>
                <el-descriptions-item label="对应订单状态 提现状态 1:待处理 2:待打款 10:打款失败 11:后台拒绝 21:成功">
                        {{ formData.status }}
                </el-descriptions-item>
                <el-descriptions-item label="子渠道; bank_code">
                        {{ formData.subChannel }}
                </el-descriptions-item>
                <el-descriptions-item label="用户uid">
                        {{ formData.uid }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createExchangePointsOrder,
  deleteExchangePointsOrder,
  deleteExchangePointsOrderByIds,
  updateExchangePointsOrder,
  findExchangePointsOrder,
  getExchangePointsOrderList
} from '@/api/exchangePointsOrder'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'ExchangePointsOrder'
})

// 自动化生成的字典（可能为空）以及字段
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

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getExchangePointsOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteExchangePointsOrderFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteExchangePointsOrderByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateExchangePointsOrderFunc = async(row) => {
    const res = await findExchangePointsOrder({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reexchangePointsOrder
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteExchangePointsOrderFunc = async (row) => {
    const res = await deleteExchangePointsOrder({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findExchangePointsOrder({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reexchangePointsOrder
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
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
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
