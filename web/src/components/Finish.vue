<template>
    <div class="finish">
        <div class="button">
        <div class="Main">
            <el-button @click="des" type="primary">返回登陆</el-button>
            <p>{{ second }}后自动返回登陆</p>
        </div>
        </div>
    </div>
</template>

<script>
export default {
    name:"finish-page",
    data() {
        return {
            second: 3,
            timerId: null
        }
    },
    methods:{
        des() {
            this.$router.push('login')
        }
    },
    created() {
        this.timerId = setInterval(() => this.second--, 1000);
    },
    watch: {
        second(newValue) {
            if (newValue == 0) {
                clearTimeout(this.timerId)
                this.$router.push('login')
            }
        }
    }
}
</script>

<style scoped>
.finish {
    position: absolute;
    height: 100%;
    width: 100%;
}
.button {
    position: absolute;
    border-radius: 3%;
    border: 1px solid black;
    height: 50%;
    width: 50%;
    top: 25%;
    left: 20%;
    background-color: skyblue;
}
p {
    color: green;
}
.Main {
    position: absolute;
    top: 45%;
    left: 45%;
}
</style>