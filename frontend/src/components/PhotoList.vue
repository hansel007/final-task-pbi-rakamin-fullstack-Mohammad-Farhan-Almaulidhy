<template>
  <div>
    <h2>Photos</h2>
    <ul>
      <li v-for="photo in photos" :key="photo.id">
        <h3>{{ photo.title }}</h3>
        <p>{{ photo.caption }}</p>
        <img :src="photo.photoUrl" alt="Photo" />
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      photos: [],
    };
  },
  async created() {
    try {
      const response = await axios.get('http://localhost:127.0.0.1/dbfinaltaskbtpn/photos');
      this.photos = response.data.data;
    } catch (error) {
      console.error(error);
    }
  },
};
</script>
