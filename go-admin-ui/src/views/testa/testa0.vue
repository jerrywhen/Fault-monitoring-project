<template>
    <div>
        testabbbbb
        <button @click="counter+=1">点击</button>
        <span>counter={{ counter }}</span>
        <button @click="clickHandle()">按钮2</button>
        <p>{{message}}</p>
        <button @click="say('hi')">say hi</button>
        <button @click="say('what')">say what</button>

        <ul>
            <li @click="clickItemHandle(item)"  v-for="(item,index) in names" :key="index">{{ item }}</li>
        </ul>
    </div>
</template>

<script>

export default{
name:'testa0',
components: {

},

mounted(){

},
data(){
    return{
        counter:1,
        message:"消息通知",
        names:["qqq","www","eeee"]
    }
},
methods:{
    clickHandle(event){
        console.log("hahaha")
        this.message="消息被撤回了"
        console.log(event);
        event.target.innerTHML="点击后"
    },
    say(data){
        console.log("11");
    },
    clickItemHandle(item){
        console.log(item);
    }

},

}
</script>