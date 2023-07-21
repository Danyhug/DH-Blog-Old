<template>
    <div>
        <el-dialog title="文章管理" :visible.sync="dialogFormVisible">
            <el-form>
                <el-form-item label="标题" label-width="100px">
                    <el-input v-model="title" autocomplete="on"></el-input>
                </el-form-item>
                <el-form-item label="内容" label-width="100px">
                    <div ref="myEditor"></div>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="updateArticle">确 定</el-button>
            </div>
        </el-dialog>

        <el-table :data="articleList" stripe style="width: 100%; font-size: 18px;">
            <el-table-column prop="ID" label="ID号" width="180">
            </el-table-column>
            <el-table-column prop="Title" label="文章标题" width="180">
            </el-table-column>
            <el-table-column prop="Content" label="文章内容">
            </el-table-column>
            <el-table-column fixed="right" label="操作" width="120">
                <template slot-scope="scope">
                    <el-button @click.native.prevent="showEditWindow(scope.$index, articleList)" type="text" size="small">
                        修改
                    </el-button>
                    <el-button @click.native.prevent="deleteRow(scope.$index, tableData)" type="text" size="small">
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
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
    data() {
        return {
            articleList: [],
            start: 0,
            limit: 6,
            pageSize: 0,
            dialogFormVisible: false,
            article: {},
            // 编辑器内容
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
    methods: {
        async getArticleList() {
            let page_id = this.$route.params.id
            // 获取起始页
            this.start = page_id * this.limit - this.limit

            let data = {
                start: this.start,
                limit: this.limit
            }
            let res = await this.$http.post('article/query', data)
            this.articleList = res.data.data
            console.log(this.articleList)
        },
        async getPageSize() {
            let res = await this.$http.post('article/pageSize')
            this.pageSize = res.data.size
        },
        showEditWindow(index, rows) {
            // 文章信息
            this.article = rows[index];
            // 将数据保存
            ({ Title: this.title, Content: this.content } = this.article)
            this.dialogFormVisible = true
            setTimeout(() => {
                this.InitEditer()
            })
        },
        // 编辑器
        InitEditer() {
            this.editor = new Editor({
                el: this.$refs.myEditor,
                ...this.editorOptions
            })
            this.editor.setMarkdown(this.content)
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
        async updateArticle() {
            const data = {
                id: this.article.ID,
                title: this.title,
                content: this.editor.getMarkdown()
            }
            let res = await this.$http.post('article/update', data)
            if (res.data.code == 1) {
                this.$message({
                    message: '文章修改成功',
                    type: 'success'
                });
            } else {
                this.$message.error('文章修改失败')
                console.error(res.data)
            }

            this.dialogFormVisible = false
            // 更新列表
            this.getArticleList()
        },
        deleteRow() {
            const h = this.$createElement;

            this.$notify({
                title: '警告',
                message: h('i', { style: 'color: teal' }, '不让删')
            });
        }
    },
    mounted() {
        this.getArticleList()
        this.getPageSize()
    }
}
</script>
<style lang="less" scoped>
.el-button {
    font-size: 16px;
}
</style>