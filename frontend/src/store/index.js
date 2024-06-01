import { createStore } from 'vuex';
import axios from 'axios';

export default createStore({
  state: {
    token: localStorage.getItem('token') || '',
    user: {},
  },
  mutations: {
    auth_success(state, { token, user }) {
      state.token = token;
      state.user = user;
    },
    logout(state) {
      state.token = '';
      state.user = {};
    },
  },
  actions: {
    async login({ commit }, user) {
      try {
        const response = await axios.post('http://localhost:127.0.0.1/dbfinaltaskbtpn/users/login', user);
        const token = response.data.token;
        // const user = response.data.user
        localStorage.setItem('token', token);
        axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
        commit('auth_success', { token, user });
      } catch (error) {
        console.error(error);
      }
    },
    logout({ commit }) {
      localStorage.removeItem('token');
      delete axios.defaults.headers.common['Authorization'];
      commit('logout');
    },
  },
  modules: {},
});
