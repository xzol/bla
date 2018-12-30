import Vue from 'vue'
import App from './App.vue'

import 'materialize-css/dist/css/materialize.css';
import 'materialize-css/dist/js/materialize.js';

import AppNavigation from "../../components/layout/menu/AppNavigation.vue"
Vue.component('app-navigation', AppNavigation);

new Vue({
    el: '#app',
    render: h => h(App)
})