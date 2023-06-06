<template>
    <div class="all">
        <Header />
        <div class="info">
            <div style="margin:3%"></div>
            <el-form>
                <el-descriptions class="margin-top" title="个人账号信息" :column="3" :size="size" border>
                    <el-descriptions-item>
                        <template slot="label">
                            <i class="el-icon-user"></i>
                            用户名
                        </template>
                        <el-form-item prop="username">
                            <input :value="username">
                        </el-form-item>
                    </el-descriptions-item>
                    <el-descriptions-item>
                        <template slot="label">
                            <i class="el-icon-mobile-phone"></i>
                            手机号
                        </template>
                        18100000000
                    </el-descriptions-item>
                    <el-descriptions-item>
                        <template slot="label">
                            <i class="el-icon-location-outline"></i>
                            居住地
                        </template>
                        苏州市
                    </el-descriptions-item>
                    <el-descriptions-item>
                        <template slot="label">
                            <i class="el-icon-tickets"></i>
                            备注
                        </template>
                        <el-tag size="small">学校</el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item>
                        <template slot="label">
                            <i class="el-icon-office-building"></i>
                            联系地址
                        </template>
                        江苏省苏州市吴中区吴中大道 1188 号
                    </el-descriptions-item>
                </el-descriptions>
                <div style="margin:3%"></div>
                <el-button style="position:absolute;left:40%;" type="primary" size="small">保存账号信息</el-button>
                <div style="margin:9%"></div>
                <el-descriptions class="margin-top" title="身体状况" :column="3" :size="size">
                    <template slot="extra">
                    </template>
                    <el-descriptions-item label="身高">100cm</el-descriptions-item>
                    <el-descriptions-item label="体重">100kg</el-descriptions-item>
                    <el-descriptions-item label="未知">xxxx</el-descriptions-item>
                    <el-descriptions-item label="未知">
                        <el-tag size="small">xxxxx</el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item label="联系地址">江苏省苏州市吴中区吴中大道 1188 号</el-descriptions-item>
                </el-descriptions>
                <el-button style="margin:3% 40%" type="primary" size="small">保存个人资料</el-button>
            </el-form>
            <el-button style="margin:5% 85% 0" type="danger" @click="des">返回</el-button>
        </div>
    </div>
</template>

<script>
import Header from '../components/Header.vue'
import axios from 'axios'

export default {
    name: "p-ersonal",
    data() {
        return {
            size: '',
            username: ''
        }
    },
    created() {
        let id = localStorage.getItem('id')
        axios.get('user', { params: { id } })
            .then(res => this.username = res.data.data.username)
            .catch(error => console.error(error));
    },
    methods: {
        des() {
            this.$router.push('index')
        }
    },
    components: {
        Header,
    }
}
</script>

<style>
.all {
    background-color: #409EFF;
    height: 100%;
    width: 100%;
}

.info {
    border-radius: 1%;
    position: absolute;
    height: 90%;
    width: 90%;
    top: 10%;
    left: 5%;
    background-color: #fff;
}
</style>