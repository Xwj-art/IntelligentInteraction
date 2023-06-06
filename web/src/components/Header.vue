<template>
    <div class="header">
        <div class="user" @dblclick="personal">
            <div class="el-icon-user-solid">
                <p v-html="username"></p>
            </div>
        </div>
        <span>
            <el-button @click="quit" class="button1" type="danger">退出登陆</el-button>
            <el-button @click="quit" class="button2" type="primary">切换账号</el-button>
        </span>
    </div>
</template>

<script>
import axios from 'axios'
export default {
    name:"header-page",
    data() {
        return {
            username:"",
            num: 0,
        }
    },
    mounted() {
        let id = localStorage.getItem('id')
        axios.get('user', { params: { id } })
            .then(res => this.username = res.data.data.username)
            .catch(error => console.error(error));
    },
    methods: {
        quit() {
            localStorage.clear()
            this.$router.push('login')
        },
        personal() {
            if (this.$route.path=='/personal') return false
            this.$router.push('personal')
        }
    },
}
</script>

<style>
.header {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    width: 100%;
    height: 10%;
    background-color: #B3C0D1;
}

.user {
    position: absolute;
    background-color: #E9EEF3;
    height: 100%;
    width: 7%;
    left: 0;
    font-size: 5vh;
    text-align: center;
}

.el-icon-user-solid p {
    position: relative;
    bottom: 10%;
    left: 25%;
    width: 15%;
    height: 5%;
    font-size: 2vh;
}

.button1 {
    position: absolute;
    right: 15%;
    bottom: 25%;
}

.button2 {
    position: absolute;
    right: 5%;
    bottom: 25%;
}
</style>