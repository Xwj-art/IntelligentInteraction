<Finish />
<template>
   <div class="index">
    <Header />
        <div class="main">
            <div class="dialog">
                <span class="btn">
                    <el-button type="success" icon="el-icon-document-add"></el-button>
                    <el-button type="success" icon="el-icon-s-promotion"></el-button>
                </span>
            </div>
        </div>
   </div>  
</template>

<script>
import axios from 'axios'
import Header from '../components/Header.vue'

export default {
    components:{
        Header,
    },
    name:"main-page",
    data() {
        return {
            username:"",
        }
    },
    created() {
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
        upload() {
        }
    }
}
</script>

<style>
.index {
    position: absolute;
    height: 100%;
    width:100%;
    background-color: #409EFF;
}
.header {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    width:100%;
    height: 10%;
    background-color: #B3C0D1;
}
.main {
    position: absolute;
    top:10%;
    left: 0;
    width:100%;
    height: 90%;
    background-color: #409EFF;
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
.el-icon-user-solid {
    position: relative;
    height: 100%;
    width: 55%;
    top: 10%;
    left: 0;
    font-size: 5.5vh;
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
.dialog {
    position: absolute;
    background-color: #fff;
    height: 80%;
    width: 80%;
    right: 10%;
    bottom: 10%;
}
.btn {
    position: absolute;
    right: 10%;
    bottom: 10%;
}
</style>