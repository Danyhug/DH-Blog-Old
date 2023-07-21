<template>
    <div>
        <div class="box">
            <!-- 标题处 -->
            <el-divider content-position="center">
                <p class="tip">文章标题</p>
            </el-divider>
            <el-input placeholder="输入文章标题" v-model="title" clearable class="title-input"></el-input>
        </div>
        <div class="box">
            <!-- 文章内容 -->
            <el-divider content-position="center">
                <p class="tip">文章内容</p>
            </el-divider>
            <div ref="myEditor"></div>
        </div>
        <div class="func-btn">
            <el-button type="primary" @click="postArticle">发布</el-button>
            <el-button type="danger" @click="clearArticle">清空</el-button>
        </div>
    </div>
</template>
  
<script>
import '@toast-ui/editor/dist/toastui-editor.css';
import '@toast-ui/editor/dist/i18n/zh-cn';
import Editor from '@toast-ui/editor';

import 'prismjs/themes/prism.css';
import '@toast-ui/editor-plugin-code-syntax-highlight/dist/toastui-editor-plugin-code-syntax-highlight.css';
import codeSyntaxHighlight from '@toast-ui/editor-plugin-code-syntax-highlight/dist/toastui-editor-plugin-code-syntax-highlight-all.js';

export default {
    name: 'Edit',
    data() {
        return {
            title: '',
            content: '',
            editorOptions: {
                minHeight: '200px',
                language: 'zh-CN',
                height: '500px',
                initialEditType: 'markdown',
                previewStyle: 'vertical',
                plugins: [codeSyntaxHighlight]
            },
            editor: {},
            viewerOptions: {
                plugins: [codeSyntaxHighlight]
            },
        }
    },
    components: {
    },
    methods: {
        InitEditer() {
            this.editor = new Editor({
                el: this.$refs.myEditor,
                ...this.editorOptions
            })
            // 移除默认事件
            this.editor.removeHook('addImageBlobHook')
            this.editor.addHook('addImageBlobHook', (blob, callback) => {
                const formData = new FormData()
                formData.append('file', blob)

                this.$http.post('article/uploadImg', formData, {
                    headers: {
                        //文件上传的类型
                        'Content-Type': 'multipart/form-data',
                    }
                }).then(res => {
                    let path = res.data
                    path = path.replace('.', this.$serveUrl + '/article')
                    // 获取内容
                    let markdown = this.editor.getMarkdown()
                    // 将图片添加
                    markdown += `![image](${path})`
                    this.editor.setMarkdown(markdown)
                    console.log(path)
                })
            })
            console.log(this.editor)
        },
        async postArticle() {
            // 获取输入的内容
            let markdown = this.editor.getMarkdown()
            // 将内容发送到数据库
            const data = {
                title: this.title,
                content: markdown
            }
            let res = await this.$http.post('article/add', data)
            if (res.data.code == 1) {
                this.$message({
                    message: '文章发布成功',
                    type: 'success'
                });
                this.clearArticle(false)
            } else {
                this.$message.error('文章发布失败')
                console.error(res.data)
            }
        },
        clearArticle(manual = true) {
            if (!manual) this.editor.setMarkdown('')
            else {
                this.$confirm('是否确定清空内容?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.editor.setMarkdown('')
                    this.$message({
                        type: 'success',
                        message: '已清空'
                    });
                }).catch(() => { })
            }
        }
    },
    mounted() {
        this.InitEditer()
    },
}

</script>

<style lang="less" scoped>
.box {
    margin: 30px 0;

    .el-divider {
        background-color: #1E9FFF;
    }
}

.tip {
    font-size: 24px;
    margin: 10px 0;
}

.title-input {
    font-size: 24px;
}
</style>