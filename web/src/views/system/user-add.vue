<template>
  <div class="app-container">
    <el-form ref="userRule" :model="form" :rules="rules" label-width="120px">
      <el-form-item label="用户名" prop="name">
        <el-input v-model="form.name" placeholder="请填写用户名" />
      </el-form-item>
      <el-form-item label="用户密码" prop="pwd">
        <el-input v-model="form.pwd" :disabled="(this.$route.query.id) != null ? true : false" placeholder="请填写用户密码" show-password />
      </el-form-item>
      <el-form-item label="角色">
        <el-col :span="10">
          <el-select v-model="form.role" placeholder="请选择角色" @change="getChange">
            <el-option
              v-for="item in roleList"
              :key="item.label"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-col>
<!--        <el-col :span="10" >-->
<!--          <el-form-item label="年龄" prop="age">-->
<!--            <el-input v-model="form.age" placeholder="请填写年龄" />-->
<!--          </el-form-item>-->
<!--        </el-col>-->
      </el-form-item>
<!--      <el-form-item label="角色">-->
<!--        <el-radio-group v-model="form.role">-->
<!--          <el-radio :label="1">admin</el-radio>-->
<!--          <el-radio :label="2">user</el-radio>-->
<!--        </el-radio-group>-->
<!--      </el-form-item>-->
      <el-form-item label="用户介绍">
        <el-input v-model="form.introduction"  type="textarea" :rows="10" placeholder="请填写用户描述"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
        <el-button @click="onClear">清除</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getDetail, save } from '@/api/auth'

export default {
  data() {
    return {
      roleList: [
        { value: 1, label: '管理员' },
        { value: 2, label: '开发者' }
      ],
      form: {
        id: 0,
        name: '',
        role: 1,
        pwd: '',
        introduction: '',
      },
      rules: {
        name: [
          { required: true, message: '请输入用户名称', trigger: 'blur' },
          { min: 3, max: 12, message: '长度在 3 到 12 个字符', trigger: 'blur' }
        ],
        // pwd: [
        //   { required: true, message: '请输入用户密码', trigger: 'blur' },
        //   { min: 6, max: 256, message: '长度在 6 到 128 个字符', trigger: 'blur' }
        // ]
        // age: [
        //   { required: true, message: '请填写年龄', trigger: 'blur' },
        //   { type: 'number', message: '请填写数字', trigger: 'blur', transform: (value) => Number(value)},
        // ]
      }
    }
  },
  created() {
    if (this.$route.query.id != null) {
      getDetail(this.$route.query.id).then(response => {
        const res = response.data.result
        this.form.id = res.id
        this.form.name = res.username
        this.form.role = res.role
        this.form.pwd = res.password
        this.form.introduction = res.introduction
      })
    }
  },
  methods: {
    onSubmit() {
      this.$refs['userRule'].validate((valid) => {
        if (valid) {
          save(this.form).then(response => {
            this.$message({
              message: '提交成功',
              type: 'success'
            })
            this.$router.push('/system/user-list')
          })
        } else {
          this.$message({
            message: '参数有误!',
            type: 'warning'
          })
          return false
        }
      })
    },
    onClear() {
      this.form.name = ''
      this.form.role = 1
      this.form.age = ''
      this.form.introduction = ''
      this.$message({
        message: '清除成功!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

