<template>
    <div>
        <!-- 这里是文章列表 -->
        <div class="tip">
            <p>文章列表</p>
        </div>
        <div class="posts">
            <Article></Article>
            <Article></Article>
            <Article></Article>
            <Article></Article>
            <Article></Article>
            <Article></Article>
        </div>
    </div>
</template>
<script>
import Article from '@/components/Article/Article.vue'
export default {
    name: 'HomeChild',
    components: {
        Article
    }
}
</script>
<style lang="less" scoped>
.tip {
    display: block;
    width: 100%;
    text-align: center;
    font-size: 24px;
    font-weight: 700;
    border-top: 1px solid var(--grey-4);

    p {
        display: inline-block;
        transform: translateY(-50%);
        background-color: #fff;
        padding: 0 16px;
        color: var(--grey-4);
    }
}

/* 文章父元素 */
.posts {
    padding: 0 30px;
}

Article {
    margin-bottom: 40px;
}
</style>