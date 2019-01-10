import Vue from 'vue'
import App from './App.vue'

import 'materialize-css/dist/css/materialize.css';
import 'materialize-css/dist/js/materialize.js';

import AppNavigation from "../../layout/menu/AppNavigation.vue"
Vue.component('app-navigation', AppNavigation);

import breadcrumbs from "../../layout/breadcrumbs/breadcrumbs.vue"
Vue.component('app-breadcrumbs', breadcrumbs);

import centralField from "./centralField.vue"
Vue.component('app-centralField', centralField);

new Vue({
    el: '#app',
    render: h => h(App)
})