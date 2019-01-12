<template>
    <nav class="breadcrumbs transparent">
        <div class="nav-wrapper">
            <div class="col s12">
                <a href="/" class="breadcrumbs_item breadcrumb grey-text text-darken-4">На главную</a>
                <a href="/litra/" class="breadcrumbs_item breadcrumb grey-text text-darken-4">Литература</a>
                <!--<a href="#!" class="breadcrumb grey-text text-darken-4">Third</a>-->
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
    .breadcrumbs_item{
        color: #1976d2 !important;
    }
    .breadcrumbs_item:hover{
        color: #0d47a1 !important;
    }
</style>