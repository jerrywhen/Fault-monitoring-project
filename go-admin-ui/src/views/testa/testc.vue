<template>
    <div>
        <input type="text" v-model="username">
        <input type="text" v-model.lazy="password">
        <p>{{ username }}{{ password }}</p>
        <button @click= "clickGetUserName" >获取用户名</button>
    </div>

</template>

<script>

export default{
name:'testc',
components: {

},

mounted(){

},
data(){
    return{
        username:"",
        password:""
    }
},
methods:{
    clickGetUserName(){
        console.log(this.username); 
        console.log(this.password); 
    }
},

}
</script>
