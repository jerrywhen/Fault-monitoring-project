<template>
    <div>
        上123123
        <keep-alive>
        <router-view v-if="$route.meta.keepAlive" />
        </keep-alive>
        <router-view v-if="$route.meta.keepAlive" />
        下123123
    </div>
</template>

<script>
export default{
    name:'TestParent'
}
</script>