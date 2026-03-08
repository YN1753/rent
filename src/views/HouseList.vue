<template>
  <div class="home-container">
    <!-- 顶部导航 -->
    <van-nav-bar>
      <template #title>
        <div @click="goToCitySelect">
          <span>{{ currentCity }}</span>
          <van-icon name="arrow-down" size="14" />
        </div>
      </template>
      <template #right>
        <van-icon name="qr" size="20" />
      </template>
    </van-nav-bar>

    <!-- 搜索框 -->
    <van-search
      v-model="keyword"
      placeholder="请输入小区名或地址"
      @search="onSearch"
    />

    <!-- 功能图标 (二手房, 新房等) -->
    <van-grid :column-num="5" icon-size="24">
      <van-grid-item icon="home-o" text="二手房" />
      <van-grid-item icon="building-o" text="新房" />
      <van-grid-item icon="shop-o" text="租房" />
      <van-grid-item icon="edit-data" text="卖房" />
      <van-grid-item icon="more-o" text="更多" />
    </van-grid>

    <!-- 房源列表 -->
    <van-list
      v-model:loading="loading"
      :finished="finished"
      finished-text="没有更多了"
      @load="onLoad"
    >
      <HouseItem
        v-for="(item, index) in houseList"
        :key="index"
        :house="item"
      />
    </van-list>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
// import { useStore } from 'pinia';
import { useRouter } from 'vue-router';
import HouseItem from '@/components/home/HouseItem.vue'; 
import { getCurrentLocation } from '@/utils/geolocation'
import { getAddressByLocation } from '@/api/location'

// const store = useStore();
const router = useRouter();

const currentCity = ref('宜宾');
const keyword = ref('');
const loading = ref(false);
const finished = ref(false);
const houseList = ref([]);
const userAddress = ref('')

// 模拟数据
const mockData = [
  { id: 1, title: '龙观嘉园 南向一居室', price: '98万', desc: '52.99m² 南' },
  { id: 2, title: '电梯2居', price: '86万', desc: '72m² 地铁900米' },
];

const onLoad = () => {
  // 异步加载数据
  setTimeout(() => {
    houseList.value.push(...mockData);
    loading.value = false;
    finished.value = true; // 演示只加载一次
  }, 1000);
};

const onSearch = () => {
  // 执行搜索逻辑
  console.log('搜索:', keyword.value);
};

const goToCitySelect = () => {
  router.push('/city');
};

onMounted(async () => {
  // 页面加载时自动获取地址（静默）
  try {
    const location = await getCurrentLocation()
    const res = await getAddressByLocation(location)
    userAddress.value = res.data.formattedAddress
  } catch (err) {
    // 静默失败，不弹窗（避免打扰用户）
    console.warn('地址获取失败:', err.message)
  }
})
</script>
