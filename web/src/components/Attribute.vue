<template>
    <div class="Attributes">
      <el-form :model="attributes" status-icon :rules="rules" ref="attributesForm" label-width="100px" class="demo-ruleForm">
        <div class="inputInfo">
          <el-form-item prop="height">
            <el-input type="number" placeholder="身高" v-model.number="attributes.height">
                <template slot="append">cm</template>
            </el-input>
          </el-form-item>
          <el-form-item prop="weight">
            <el-input type="number" placeholder="体重" v-model.number="attributes.weight">
                <template slot="append">kg</template>
            </el-input>
          </el-form-item>
        </div>
        <span class="twoButton">
            <el-button type="danger" @click="submit">提交资料</el-button>
            <el-button type="danger" @click="Submit">暂不填写</el-button>
        </span>
      </el-form>
    </div>
</template>

<script>
export default {
    name:"att-ribute",
    data() {
    return {
      attributes: {
        height: null,
        weight: null 
      },
      rules: {
        height: [
          { required: true, message: "请输入身高", trigger: "blur" },
        ],
        weight: [
          { required: true, message: "请输入体重", trigger: "blur" },
        ],
      },
    }
  },
  methods: {
    submit() {
      this.$refs.attributesForm.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('提交失败，输入字段不符合')
        }
        const { data:res } = await this.$http.post('att/add', this.attributes)
        if (res.status != 202) return this.$message.error(res.message)
        this.$emit("next", 2)
        return this.$message.success('提交成功，点击返回登陆')
      })
    },
    Submit() {
        const { data:res } = this.$http.post('att/add', this.attributes)
        if (res.status != 202) return this.$message.error(res.message)
        this.$emit("next", 2)
        return this.$message.success('点击返回登陆')
    }
  },
  props:{
    active: Number,
  },
}
</script>

<style>
.Attributes {
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
.twoButton {
    position: absolute;
    bottom: 13%;
    left: 35%;
}
</style>