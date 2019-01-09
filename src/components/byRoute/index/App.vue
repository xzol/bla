<template>
    <div class="app container">
        <app-navigation></app-navigation>
        <div class="row">
            <div class="col s12">
                <app-breadcrumbs></app-breadcrumbs>
            </div>
        </div>
        <div class="row">
            <div class="col s12">
                <app-body></app-body>
            </div>
        </div>
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