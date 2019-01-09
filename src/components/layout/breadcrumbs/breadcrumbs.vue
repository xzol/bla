<template>
    <nav class="transparent">
        <div class="nav-wrapper">
            <div class="col s12">
                <a href="#!" class="breadcrumb grey-text text-darken-4">First</a>
                <a href="#!" class="breadcrumb grey-text text-darken-4">Second</a>
                <a href="#!" class="breadcrumb grey-text text-darken-4">Third</a>
            </div>
        </div>
    </nav>
</template>

<script>
    export default {
        name: "breadcrumbs"
    }
</script>

<style scoped>
    nav{
        box-shadow:none;
    }
    .breadcrumb:before {
        color: #B39DDB;
    }
</style>