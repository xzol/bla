<template>
    <div class="app">
        <h2>{{message}}</h2>
        <app-navigation></app-navigation>
    </div>
</template>

<script>
    export default {
        name: 'App',
        data () {
            return {
                message: 'Welcome to Vue.js'
            }
        }
    }
</script>

<style>
    #hello {
        font-family: Verdana;
        color: navy;
    }
</style>