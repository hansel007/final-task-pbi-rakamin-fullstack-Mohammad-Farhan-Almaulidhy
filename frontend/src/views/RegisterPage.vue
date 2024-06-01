<template>
  <div class="register">
    <h1>Register</h1>
    <RegisterComponent />
  </div>
</template>

<script>
import RegisterComponent from '@/components/RegisterForm.vue';

export default {
  name: 'RegisterView',
  components: {
    RegisterComponent,
  },
};
</script>

<style scoped>
.register {
  text-align: center;
  margin-top: 20px;
}
</style>
