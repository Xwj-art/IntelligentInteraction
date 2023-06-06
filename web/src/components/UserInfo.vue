<template>
    <div class="userInfo">
      <el-form :model="userInfo" status-icon :rules="rules" ref="registerForm" label-width="100px" class="demo-ruleForm">
        <div class="inputInfo">
          <el-form-item prop="username">
            <el-input type="text" placeholder="用户名" maxlength="12" show-word-limit v-model="userInfo.username"><i
                class="el-icon-view"></i></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input type="password" placeholder="密码" maxlength="20" suffix-icon="el-icon-question" v-model="userInfo.password"></el-input>
          </el-form-item>
          <el-form-item prop="">
            <el-input type="password" placeholder="确认密码" maxlength="20" suffix-icon="el-icon-question" v-model="confirmPassword"></el-input>
          </el-form-item>
        </div>
        <span class="oneButton">
            <el-button type="danger" @click="register">注册</el-button>
        </span>
      </el-form>
    </div>
</template>

<script>
export default {
    name:"user-info",
    data() {
    return {
      userInfo: {
        username: "",
        password: "",
        role: 2,
      },
      confirmPassword: "",
      rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          { min: 1, max: 12, message: "用户名必须小于12个字符", trigger: "blur" },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { min: 6, max: 20, message: "用户名必须在6-20个字符之间", trigger: "blur" },
        ],
        confirmPassword: [
            { required: true, message: "请输入密码", trigger: "blur" },
            { min: 6, max: 20, message: "用户名必须在6-20个字符之间", trigger: "blur" },
        ],
      },
    }
  },
  methods: {
    register() {
      this.$refs.registerForm.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('注册失败，输入字段不符合')
        }
        if (this.confirmPassword != this.userInfo.password) {
            this.$refs.registerForm.resetFields()
            this.confirmPassword = ""
            return this.$message.error('两次密码不同，请重新输入')
        }
        const { data:res } = await this.$http.post('user/add', this.userInfo)
        if (res.status != 202) return this.$message.error(res.message)
        localStorage.setItem('id', res.data.id)
        this.$emit("next", 1)
        return this.$message.success('注册成功，继续完善个人信息')
      })
    },
  },
  props:{
    active: Number,
  },
}
</script>

<style scoped>
.userInfo {
    position: absolute;
    height: 70%;
    width: 60%;
    left: 20%;
    bottom: 15%;
    border:1px solid black;
    background-color: #409EFF;
    border-radius: 3%;
}

.inputInfo {
  position: absolute;
  width: 80%;
  right: 20%;
  top: 20%;
  border-radius: 3%;
}
.oneButton {
    position: absolute;
    bottom: 13%;
    left: 46%;
}
</style>