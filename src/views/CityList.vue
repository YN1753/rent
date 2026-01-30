<template>
  <div class="city-container">
    <van-search placeholder="请输入城市名称" v-model="searchKey" />

    <div class="location">
      <span>定位：</span>
      <van-button type="primary" size="small" @click="getLocation">开启定位</van-button>
    </div>

    <van-index-bar :index-list="indexList">
      <!-- 热门城市 -->
      <van-index-anchor index="hot">热门城市</van-index-anchor>
      <van-cell
        v-for="city in hotCities"
        :key="city"
        :title="city"
        @click="selectCity(city)"
      />

      <!-- 城市列表 -->
      <template v-for="(item, key) in cities" :key="key">
        <van-index-anchor :index="key">{{ key }}</van-index-anchor>
        <van-cell
          v-for="city in item"
          :key="city"
          :title="city"
          @click="selectCity(city)"
        />
      </template>
    </van-index-bar>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { getPreciseLocation } from '@/utils/geolocation';
import { updateUserLocation } from '@/api/user';

const router = useRouter();

// 模拟数据结构
const citiesData = {
  A: ['安庆', '安阳'],
  B: ['宝鸡', '包头', '保定', '北海'],
  // ... 更多数据
};

const hotCities = ['北京', '上海', '广州', '深圳', '成都'];

const searchKey = ref('');

// 计算属性：根据首字母生成索引列表
const indexList = computed(() => {
  return Object.keys(citiesData);
});

const selectCity = (cityName) => {
  // 保存选中的城市到Pinia或LocalStoage
  console.log('选中城市:', cityName);
  router.push('/home'); // 跳转回首页
};

const getLocation = async () => {
  try {
    const loc = await getPreciseLocation();
    await updateUserLocation({
      longitude: loc.longitude,
      latitude: loc.latitude
    });
  } catch (e) {
    selectCity('宜宾');
  }
};
</script>
