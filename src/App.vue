<template>
  <div id="app">
    <!-- 路由出口，页面在这里显示 -->
    <Header />

    <!-- 页面主题内容 -->
     <main class="main-content">
       <router-view />
     </main>
     <!-- 引入登录弹窗组件 -->
     <LoginAccount 
    v-if="loginType === 'account'" 
    @close="loginType = null"
    @switch="loginType = $event"
    @open-register="showRegisterModal = true"
  />
  <LoginEmail 
    v-if="loginType === 'email'" 
    @close="loginType = null"
    @switch="loginType = $event"
    @open-register="showRegisterModal = true"
  />
     <!-- 引入注册弹窗组件 -->
  <Register 
  v-if="showRegisterModal" 
  @close="showRegisterModal = false" 
/>
  </div>
</template>

<script setup>
import { ref } from 'vue';
// import { ref, watch } from 'vue';
// import { useRoute, useRouter } from 'vue-router';
import Header from '@/views/Header.vue';
import Register from '@/views/Register.vue';
import LoginAccount from '@/views/LoginAccount.vue'
import LoginEmail from '@/views/LoginEmail.vue'
// const route = useRoute();
// const router = useRouter();

// 控制注册弹窗显示
const showRegisterModal = ref(false);
// 控制登录弹窗显示
const loginType = ref(null);
// 打开登录弹窗
const openLogin = () => {
  loginType.value = 'account' // 默认打开用户名登录
}
</script> 

<style scoped>
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.main-content {
  flex: 1;
}
</style>
