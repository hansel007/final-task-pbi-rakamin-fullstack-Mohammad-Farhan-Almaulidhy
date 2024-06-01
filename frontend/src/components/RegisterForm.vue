<template>
  <div>
    <h2>Register</h2>
    <form @submit.prevent="register">
      <div>
        <label for="username">Username</label>
        <input type="text" v-model="username" required />
      </div>
      <div>
        <label for="email">Email</label>
        <input type="email" v-model="email" required />
      </div>
      <div>
        <label for="password">Password</label>
        <input type="password" v-model="password" required minlength="6" />
      </div>
      <button type="submit">Register</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      email: '',
      password: '',
    };
  },
  methods: {
    async register() {
      try {
        const user = {
          username: this.username,
          email: this.email,
          password: this.password,
        };
        await axios.post('http://localhost:127.0.0.1/dbfinaltaskbtpn/users/register', user);
        this.$router.push('/login');
      } catch (error) {
        console.error(error);
      }
    },
  },
};
</script>
