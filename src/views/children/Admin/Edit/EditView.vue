<template>
    <div>
        {{ content }}
        <Editor ref="myEditor" :options="editorOptions" height="500px" initialEditType="markdown" previewStyle="vertical" />
        <Viewer ref="myViewer" :initialValue="viewerText" />
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
            viewerText: '# This is Viewer.\n Hello World.'
        }
    },
    components: {
        Editor,
        Viewer
    },
    methods: {
        postArticle() {
            // 获取输入的内容
            let markdown = this.$refs.myEditor.invoke('getMarkdown')
            this.$refs.myViewer.invoke('setMarkdown', markdown)
        }
    }
}

</script>
  