<template>
  <!--整个页面-->
  <div class="login">
    <!--登陆框-->
    <div class="loginBox">
      <!--1-12-->
      <!--6-40-->
      <el-form :model="userInfo" status-icon :rules="rules" ref="loginForm" label-width="100px" class="demo-ruleForm">
        <div class="inputInfo">
          <el-form-item prop="username">
            <el-input @keyup.enter.native="login" type="text" placeholder="用户名" maxlength="12" show-word-limit v-model="userInfo.username"><i
                class="el-icon-view"></i></el-input>
          </el-form-item>
          <div style="margin: 20px 0 0 0;"></div>
          <el-form-item prop="password">
            <el-input @keyup.enter.native="login" type="password" placeholder="密码" maxlength="20" show-word-limit suffix-icon="el-icon-question" v-model="userInfo.password"
             ></el-input>
          </el-form-item>
        </div>
        <span class="twoButton">
            <el-button type="primary" @click="login">登陆</el-button>
            <el-button type="danger" @click="register">注册</el-button>
        </span>
      </el-form>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'LoginPanel',
  data() {
    return {
      userInfo: {
        username: "",
        password: "",
        role: 2,
      },
      rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          { min: 1, max: 12, message: "用户名小于12个字符", trigger: "blur" },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { min: 6, max: 20, message: "密码在6-20个字符之间", trigger: "blur" },
        ],
      },
    }
  },
  methods: {
    login() {
      this.$refs.loginForm.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('登陆失败，登陆字段不符合')
        }
        const { data:res } = await this.$http.post('login', this.userInfo)
        if (res.status != 202) return this.$message.error(res.message)
        const id = res.id
        localStorage.setItem('token', res.token)
        localStorage.setItem('id', id)

        let tep
        axios.get('att/:id', { params: { id } })
          .then(({ data:res }) => tep = res.status )
          .catch(error => console.error(error));

        tep = res.status
        if (tep != 202) return this.$message.error('登陆失败，请先完成注册全部步骤')
        this.$router.push('index')
        return this.$message.success('登陆成功')
      })
    },
    register() {
      this.$router.push('register')
      /*
      this.$refs.loginForm.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('注册失败，输入字段不符合')
        }
        const { data:res } = await this.$http.post('user/add', this.userInfo)
        if (res.status != 202) return this.$message.error(res.message)
        return this.$message.success('注册成功，返回登陆')
      })
      */
    },
  },
}

</script>

<style scoped>
.login {
  height: 100%;
  background-color: #409EFF;
}

.loginBox {
  position: relative;
  left: 20%;
  top: 20%;
  height: 60%;
  width: 60%;
  border-radius: 5%;
  background-color: white;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.9)
}

.twoButton {
  position: absolute;
  bottom: 10%;
  left: 0;
  width: 100%;
  text-align: center;
}

.inputInfo {
  position: absolute;
  width: 60%;
  left: 20%;
  top: 30%;
  border-radius: 3%;
}
</style>