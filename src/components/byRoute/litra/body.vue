<template>
    <div class="body">
        <div class="col s12">
            <button @click="ls">asdasd</button>
        </div>
    </div>
</template>

<script>
    export default {
        name: "body",
        watch: {},
        mounted(){
                this.ls()

        },
        methods: {
            ls:function () {
                console.log(fetch("https://jsonplaceholder.typicode.com/users/1"));
                // return
                // axios.get('/litra/ls')
                //     .then(function (response) {
                //         if (response.data.status === true) {
                //             for (let field in response.data.records){
                //                 self.fields.push(response.data.records[field])
                //             }
                //         }
                //     })
                //     .catch(function (error) {
                //
                //     });
            }
        },
        props:{

        }
    }
</script>

<style scoped>
    .body{
        margin-top: -38px;
    }
</style>