<template>
    <div id="app">
        <Header></Header>
        <Banner>
            <h3>{{ title }}</h3>
            <div class="top">
                <span class="date">发表于 {{ $getDate(created, true) }}</span>
                <span class="num-word">本文字数 {{ $getWordCount(content) }} 字</span>
                <span class="time-consum">阅读时长 {{ (content.length / 300 + 0.5).toFixed(0) }} 分钟</span>
            </div>
        </Banner>
        <Main>
            <ArticleChild :content="content" :update="update" :viewnum="viewnum"></ArticleChild>
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
    data() {
        return {
            // 文章信息
            id: 1,
            title: '',
            content: '',
            created: 0,
            update: 0,
            viewnum: 0
        }
    },
    props: {
    },
    methods: {
        // 获取文章内容
        async getArticle() {
            this.id = this.$route.params.id
            let res = await this.$http.get('article/' + this.id)
            let data = res.data;
            ({ Title: this.title, Content: this.content, Created: this.created, Updated: this.update, Viewnum: this.viewnum } = data.data)

        }
    },
    mounted() {
        this.getArticle()
    },
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