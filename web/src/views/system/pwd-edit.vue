<template>
  <div class="app-container">
    <el-form ref="pwdRule" :model="form" :rules="rules" label-width="120px">
      <el-form-item label="旧密码" prop="old_pwd">
        <el-input v-model="form.old_pwd" placeholder="请填写旧密码" show-password />
      </el-form-item>
      <el-form-item label="新密码" prop="new_pwd_1">
        <el-input v-model="form.new_pwd_1" placeholder="请填写新密码" show-password />
      </el-form-item>
      <el-form-item label="确认新密码" prop="new_pwd_2">
        <el-input v-model="form.new_pwd_2" placeholder="请确认新密码" show-password />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">确认修改</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { changePwd } from '@/api/auth'

export default {
  data() {
    return {
      form: {
        old_pwd: '',
        new_pwd_1: '',
        new_pwd_2: ''
      },
      rules: {
        old_pwd: [
          { required: true, message: '请输入旧密码', trigger: 'blur' },
          { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' }
        ],
        new_pwd_1: [
          { required: true, message: '请输入新密码', trigger: 'blur' },
          { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' }
        ],
        new_pwd_2: [
          { required: true, message: '请再次输入新密码', trigger: 'blur' },
          { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' }
        ]
      }
    }
  },
  created() {

  },
  methods: {
    onSubmit() {
      this.$refs['pwdRule'].validate((valid) => {
        if (valid) {
          changePwd(this.form).then(response => {
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

