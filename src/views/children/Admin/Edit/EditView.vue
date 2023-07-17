<template>
    <div>
        {{ content }}
        <Editor ref="myEditor" :options="editorOptions" height="500px" initialEditType="markdown" previewStyle="vertical" />
        <Viewer ref="myViewer" :options="viewerOptions" :initialValue="viewerText" />
        <el-button type="primary" @click="postArticle">确定</el-button>
    </div>
</template>
  
<script>
import '@toast-ui/editor/dist/toastui-editor.css';
import '@toast-ui/editor/dist/i18n/zh-cn';
import { Editor, Viewer } from '@toast-ui/vue-editor';

import 'prismjs/themes/prism.css';
import '@toast-ui/editor-plugin-code-syntax-highlight/dist/toastui-editor-plugin-code-syntax-highlight.css';
import codeSyntaxHighlight from '@toast-ui/editor-plugin-code-syntax-highlight/dist/toastui-editor-plugin-code-syntax-highlight-all.js';

export default {
    name: 'Edit',
    data() {
        return {
            content: '',
            editorOptions: {
                minHeight: '200px',
                language: 'zh-CN',
                plugins: [codeSyntaxHighlight]
            },
            viewerOptions: {
                plugins: [codeSyntaxHighlight]
            },
            viewerText: '# 我是一个测试文本\n## 是的\n```javascript\nlet n = Number;\n```'
        }
    },
    components: {
        Editor,
        Viewer
    },
    methods: {
        async postArticle() {
            // 获取输入的内容
            let markdown = this.$refs.myEditor.invoke('getMarkdown')
            this.$refs.myViewer.invoke('setMarkdown', markdown)
            // 将内容发送到数据库
            const data = {
                title: '测试内容',
                content: markdown
            }
            let res = await this.$http.post('article/add', data)
            console.log(res.data)
            // const data = {
            //     start: 0,
            //     limit: 10
            // }
            // this.$http.post('article/query', data).then(res => {
            //     console.log(res)
            // })
        }
    }
}

</script>
  