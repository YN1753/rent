<template>
  <div class="house-item">
    <img :src="house.imgUrl" alt="房源图" />
    <div class="info">
      <h3>{{ house.title }}</h3>
      <p>{{ house.desc }}</p>
      <div class="price">
        <span>{{ house.price }}</span>
        <span>{{ house.unitPrice }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  house: {
    type: Object,
    required: true
  }
});
</script>

<style scoped>
.house-item {
  display: flex;
  padding: 12px;
  border-bottom: 1px solid #eee;
}
.info { flex: 1; margin-left: 12px; }
.price { color: #f60; font-weight: bold; }
</style>