<template>
  <div class="login">
    <h1>Login</h1>
    <LoginComponent />
  </div>
</template>

<script>
import LoginComponent from '@/components/LoginForm.vue';

export default {
  name: 'LoginView',
  components: {
    LoginComponent,
  },
};
</script>

<style scoped>
.login {
  text-align: center;
  margin-top: 20px;
}
</style>
