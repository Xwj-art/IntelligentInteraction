<template>
    <div class="register">
        <div class="process">
            <el-steps :space="400" :active="active" finish-status="success">
                <el-step title="注册"></el-step>
                <el-step title="完善信息"></el-step>
                <el-step title="返回登陆"></el-step>
            </el-steps>
        </div>
        <div class="main">
            <div v-if="active == 0">
                <UserInfo :active="active" @next="next" />
            </div>
            <div v-else-if="active == 1">
                <Attribute :active="active" @next="next" />
            </div>
            <div v-else-if="active == 2">
                <Finish :active="active" @next="next" />
            </div>
        </div>
    </div>
</template>

<script>
import Attribute from '@/components/Attribute.vue';
import Finish from '@/components/Finish.vue';
import UserInfo from '@/components/UserInfo.vue';

export default {
    data() {
        return {
            active: 0
        };
    },
    methods: {
        next(i) {
            this.active = i
        }
    },

    components: {
        Attribute,
        Finish,
        UserInfo
    }
}
</script>

<style scoped>
.register {
    height: 100%;
    background-color: #fff;
}

.process {
    position: relative;
    top: 5%;
    left: 18%;
}

.main {
    position: relative;
    left: 10%;
    top: 10%;
    height: 80%;
    width: 80%;
    background-color: white;
}
</style>
