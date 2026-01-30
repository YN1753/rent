<template>
  <div class="header">
    <!-- 左侧Logo和城市选择 -->
    <div class="left">
      <img src="@/assets/vue.svg" alt="logo" class="logo" />
      <span class="city">• {{ currentCity }}</span>
    </div>

    <!-- 中间导航菜单 -->
    <nav class="nav">
      <router-link to="/secondhand">二手房</router-link>
      <router-link to="/newhouse">新房</router-link>
      <router-link to="/rent">租房</router-link>
      <router-link to="/decorate">装修</router-link>
      <router-link to="/commercial">商业办公</router-link>
      <router-link to="/community">小区</router-link>
      <router-link to="/tools">工具</router-link>
      <router-link to="/publish">发布房源</router-link>
    </nav>

    <!-- 右侧用户区 - 核心功能 -->
    <div class="user-area">
      <template v-if="isLogin">
        <router-link to="/user" class="user-link">
          <van-icon name="user-o" size="18" />
          <span>我的</span>
        </router-link>
      </template>
      <template v-else>
        <router-link to="/login" class="user-link">
          登录
        </router-link>
        <router-link to="/register" class="user-link">
          立即注册
        </router-link>
      </template>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue';
import { useUserStore } from '@/store/modules/user';

const userStore = useUserStore();
const isLogin = computed(() => !!userStore.token);

// 模拟当前城市
const currentCity = ref('宜宾');
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  position: sticky;
  top: 0;
  z-index: 100;
}

.left {
  display: flex;
  align-items: center;
  gap: 10px;
}
.logo {
  width: 60px;
  height: 30px;
}
.city {
  font-size: 12px;
  color: #999;
}

.nav {
  display: flex;
  gap: 20px;
}
.nav a {
  text-decoration: none;
  color: #333;
  font-size: 14px;
}
.nav a:hover {
  color: #1989fa;
}

.user-area {
  display: flex;
  gap: 10px;
}
.user-link {
  text-decoration: none;
  color: #333;
  font-size: 14px;
}
.user-link:hover {
  color: #1989fa;
}
</style>
