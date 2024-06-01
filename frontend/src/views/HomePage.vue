<template>
  <div class="home">
    <h1>Welcome to the Home Page</h1>
    <PhotoList />
  </div>
</template>

<script>
import PhotoList from '@/components/PhotoList.vue';

export default {
  name: 'HomePage',
  components: {
    PhotoList,
  },
};
</script>

<style scoped>
.home {
  text-align: center;
  margin-top: 20px;
}
</style>
