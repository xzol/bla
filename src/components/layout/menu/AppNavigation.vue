<template>
    <div class="app-navigation">
        <nav>
            <div class="nav-wrapper deep-purple lighten-3">
                <a href="/" class="brand-logo app-navigation_app-navigation">к0д.рф</a>
                <ul id="nav-mobile" class="right hide-on-med-and-down">
                    <li><a href="/litra/">Литература</a></li>
                    <li><a href="">Книги</a></li>
                    <li><a href="">Видео</a></li>
                </ul>
            </div>
        </nav>
    </div>
</template>

<script>
    export default {
        name: "AppNavigation"
    }
</script>

<style scoped>
    .app-navigation_app-navigation{
        padding: 5px;
    }
</style>