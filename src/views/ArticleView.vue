<template>
    <div id="app">
        <Header></Header>
        <Banner>
            <h3>安卓大作业——飞鸽传书（高仿微信）</h3>
            <div class="top">
                <span class="date">发表于 2023-07-11</span>
                <span class="num-word">本文字数 6.5k 字</span>
                <span class="time-consum">阅读时长 6 分钟</span>
            </div>
        </Banner>
        <Main>
            <ArticleChild></ArticleChild>
        </Main>
        <Footer></Footer>
    </div>
</template>
  
<script>
import Header from '@/components/Header/Header'
import Banner from '@/components/Banner/Banner'
import Main from '@/components/Main/Main.vue'
import ArticleChild from '@/components/Main/Children/ArticleChild.vue'
import Footer from '@/components/Footer/Footer'

export default {
    name: 'HomeView',
    components: {
        Header,
        Banner,
        Main,
        ArticleChild,
        Footer
    }
}
</script>
<style lang="less" scoped>
.top {
    margin-top: 1.125rem;
    span {
        margin-right: 1.25rem;
    }
}
</style>