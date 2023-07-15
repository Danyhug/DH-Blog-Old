<template>
    <div>
        <router-link to="/">首页</router-link>
        <div class="blog-container"></div>
        <div class="info">
            <span>
                更新于 2023-07-13
            </span>
            <span>
                阅读次数 9 次
            </span>
        </div>
    </div>
</template>
<script>
export default {
    name: 'ArticleChild'
}
</script>
<style lang="less" scoped>
.info {
    padding: .625rem 0;
    font-size: .75rem;
    color: var(--gery-6);
    text-align: right;
    border-top: .0625rem solid var(--grey-4);

    span {
        margin: 0 .625rem;
    }
}
</style>