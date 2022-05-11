<template>
    <div>
        <el-row>
            <el-row class="tool-bar">
                <el-row class="right">
                    <el-button type="success" v-permission="['editor','admin']" @click="addFile">添加文件</el-button>
                </el-row>
            </el-row>
        </el-row>

        <el-row v-for="item,index in contract.files" :key="index">
            <el-col :span="18" style="margin-top:20px; font-size:15px;">{{index}} - {{item}}</el-col>
            <el-col :span="6" style="margin-top:15px;">
                <el-button size="mini" type="primary" @click="showImage(item)">查看</el-button>
                <el-button size="mini" v-permission="['admin']" @click="deleteFile(item)">删除</el-button>
            </el-col>
        </el-row>

        <el-dialog
            :title="dialogTitle"
            append-to-body
            :visible.sync="dialogImage"
            width="40%"
        >
            <img :src="path" style="width:400px;height:400px">
            
        </el-dialog>

        <el-dialog
            :title="dialogTitle"
            append-to-body
            :visible.sync="dialogVisible"
            width="40%"
        >
            
            <el-form ref="form" :model="contract" label-width="80px">

                <el-form-item label="文件上传">
                    <el-upload
                    class="upload-demo"
                    drag
                    ref="upload"
                    action="xxx"
                    :limit="1"
                    :http-request="uploadFile"
                    :auto-upload="false"
                    multiple>
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                    <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
                    </el-upload>
                </el-form-item>
            </el-form>

            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="uploadSubmit">上 传</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>

// import http from '@/utils/http'
import {uploadContractFile,addUploadFile,deleteFile} from '@/api/contract.js'

export default {
    props: [
        'contract',
    ],
    data(){
        return {
            tableData: [],
            dialogVisible: false,
            dialogImage:false,
            dialogTitle: "添加合同文件",

            file:{},
            contractNew:{},

            path:""

            // path:require('@/assets/path/61ab1d5d72bb7232f81f6d1e/dog.jpg')

        }
    },
    created(){
        console.log("created----------");
        console.log(this.contract);  
        
        this.contractNew = this.contract;

    },
    methods:{
        // 图片预览
        showImage(fileName){
            this.dialogImage = true

            this.dialogTitle = '图片预览'

            console.log(fileName)

            this.path = require('@/assets/path/'+this.contractNew.id+'/'+fileName)
        },

        // 提交按钮
        uploadSubmit(){
            this.$refs.upload.submit();
            this.dialogVisible = false;
        },
        // 上传文件
        uploadFile(param){
            let formData = new FormData();
            formData.append("contract", this.contractNew.id)
            formData.append("file", param.file)
            formData.append("name","namename123123")
            
            uploadContractFile(formData).then(res=>{

                //file--文件名
                this.file = res.data.file

                console.log("-----res.date")
                console.log(res.data)


                // 添加文件名到数据库
                this.contractNew.files.push(this.file)

                addUploadFile(this.contractNew).then(res=>{
                    this.$message({
                        message:'文件上传成功',
                        type:'success'    
                    });
                    console.log(res)  
                })
                // 调用父级reset方法
                this.$parent.$parent.reset();       
            })
        },
        addFile(){
            this.dialogVisible = true;
        },
        deleteFile(item){
            this.contractNew.name = item

            // 删除确认
            if (window.confirm ("你确定真的要删除吗")){
                deleteFile(this.contractNew).then(res=>{
                    this.$message({
                        message:'删除成功',
                        type:'success'    
                    });
                    console.log(res)  
                })
                // 调用父级reset方法
                this.$parent.$parent.reset();
            }
            
        },


    }
}
</script>