<template>
    <div>
        <router-link to="/">首页</router-link>
        <div class="blog-container">
            <Viewer ref="myViewer" :options="viewerOptions" height="6.25rem" />
        </div>
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
import { Viewer } from '@toast-ui/vue-editor';
import 'prismjs/themes/prism.css';
import '@toast-ui/editor-plugin-code-syntax-highlight/dist/toastui-editor-plugin-code-syntax-highlight.css';
import codeSyntaxHighlight from '@toast-ui/editor-plugin-code-syntax-highlight/dist/toastui-editor-plugin-code-syntax-highlight-all.js';

export default {
    name: 'ArticleChild',
    props: {
        content: {
            type: String,
            require: true
        }
    },
    data() {
        return {
            viewerOptions: {
                plugins: [codeSyntaxHighlight]
            },
        }
    },
    methods: {
        setContent() {
            console.log('content', this.content)
            this.$refs.myViewer.invoke('setMarkdown', this.content)
        }
    },
    watch: {
        content: {
            handler() {
                this.setContent()
            }
        }
    },
    components: {
        Viewer
    }
}
</script>
<style lang="less" scoped>
.blog-container {
    padding: 1.25rem 1.5rem;
}

.info {
    padding: 10px 0;
    font-size: 12px;
    color: var(--gery-6);
    text-align: right;
    border-top: 1px solid var(--grey-4);

    span {
        margin: 0 10px;
    }
}
</style>