<template>
    <div>
        
            <el-row>
                <el-row class="tool-bar">
                    <el-row class="left">
                        <el-input v-model="keyword" placeholder="关键字" style="width:200px;"></el-input>
                        <el-button @click="list" type="primary" icon="el-icon-search" style="margin-left:5px;">搜索</el-button>
                    </el-row>
                    <el-row class="right">
                        <el-button type="success" @click="create">添加合同</el-button>
                    </el-row>
                </el-row>
            </el-row>
            <el-table :data="tableData" stripe style="width:100%">
                <el-table-column prop="id" label="ID" min-width="100" header-align="center" align="center" />
                <el-table-column prop="title" label="标题" min-width="100" header-align="center" align="center" />
                <el-table-column prop="startDate" label="开始时间" min-width="100" header-align="center" align="center" />
                <el-table-column prop="endDate" label="结束时间" min-width="100" header-align="center" align="center" />
                <el-table-column prop="desc" label="描述" min-width="100" header-align="center" align="center" />
                <el-table-column  prop="state" label="状态" min-width="100" header-align="center" align="center" >
                    <template slot-scope="scope">
                        <span v-if="scope.row.state=='0'">未签约</span>
                        <span v-if="scope.row.state=='1'">签约中</span>
                        <span v-if="scope.row.state=='2'">已生效</span>
                        <span v-if="scope.row.state=='3'">已终止</span>
                        <span v-if="scope.row.state=='4'">已过期</span>
                    </template>
                </el-table-column>
                <el-table-column  label="操作" min-width="150" header-align="center" align="center" >
                    <template slot-scope="scope">
                        <el-button type="primary" title="合同文件"  @click="fileList(scope.row)">合同文件</el-button>
                        <el-button type="warning" icon="el-icon-edit" v-permission="['editor','admin']" circle @click="update(scope.row)"></el-button>
                        <el-button type="danger" icon="el-icon-delete" v-permission="['admin']" circle @click="remove(scope.row.id)"></el-button>
                    </template>
                </el-table-column>
            </el-table>

            <el-row>
                <el-pagination
                    @size-change="handleSizeChange"
                    @current-change="handleCurrentChange"
                    :current-page.sync="page"
                    :page-sizes="[1, 10, 20, 50, 100]"
                    :page-size="size"
                    layout="total, sizes, prev, pager, next, jumper"
                    :total="total">
                </el-pagination>
            </el-row>
<!-- 新增、编辑信息 -->
        <el-dialog
            :title="dialogTitle"
            :visible.sync="dialogVisible"
            width="40%">
            
            <el-form ref="form" :model="contract" label-width="80px">
                <el-form-item label="标题">
                    <el-input v-model="contract.title" style="width:300px"></el-input>
                    <span style="margin-left: 10px">标题不能为空</span>
                </el-form-item>
                <el-form-item label="开始时间">
                    <el-input v-model="contract.startDate" style="width:300px;"></el-input>
                </el-form-item>
                <el-form-item label="结束时间">
                    <el-input v-model="contract.endDate"  style="width:300px;"></el-input>
                </el-form-item>
                <el-form-item label="状态">
                    <template>
                        <el-select v-model="contract.state" placeholder="请选择">
                            <el-option label="未签约" value="0">未签约</el-option>
                            <el-option label="签约中" value="1">签约中</el-option>
                            <el-option label="已生效" value="2">已生效</el-option>
                            <el-option label="已终止" value="3">已终止</el-option>
                            <el-option label="已过期" value="4">已过期</el-option>
                        </el-select>
                    </template>
                </el-form-item>
                <el-form-item label="合同简介">
                    <el-input v-model="contract.desc" type="textarea" rows="2" style="width:300px;"></el-input>
                </el-form-item>

            </el-form>
            
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button v-if="this.dialogTitle=='新增合同'" type="primary" @click="add">提 交</el-button>
                <el-button v-if="this.dialogTitle=='编辑'" type="primary" @click="edit">保 存</el-button>
            </span>
        </el-dialog>

<!-- 文件上传 -->
        <el-dialog
            title="合同文件管理"
            :visible.sync="fileListDialogVisible"
            :close-on-click-modal="false"
            :show-close="false"
            append-to-body
            width="40%"
        >
            <contract-file-list :contract="contract" v-if="fileListDialogVisible"></contract-file-list>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="reset">返回</el-button>
            </span>
        </el-dialog>
        
    </div>
    
</template>



<script>

// import http from '@/utils/http'
import {listContract,createContract,updateContract,deleteContract} from '@/api/contract'

import contractFileList from '@/views/contract/fileList.vue'

export default {
    components:{
        contractFileList
    },
    data(){
        return {
            keyword: "",
            page: 1,
            size: 10,
            total: 0,
            tableData: [],
            dialogVisible: false,
            dialogTitle:"",
            fileListDialogVisible:false,
            contract : {},
        
        }
    },
    created(){
        this.list();
    },
    methods:{
        // dialog“返回”按钮点击事件
        reset(){
            this.fileListDialogVisible = false;
            this.list();
        },
        // "合同文件"按钮点击事件
        fileList(contract){
            this.fileListDialogVisible = true;
           
            this.contract = contract;

        },
        create(){
            this.dialogVisible = true;
            this.dialogTitle = "新增合同";
            this.contract = { }
        },
        // 添加提交
        add(){
            createContract(this.contract).then(res=>{
                this.$message({
                    message:'添加成功',
                    type:'success'    
                });
                this.list();
                this.dialogVisible = false;
            })
        },
        // 数据回显
        update(contract){
            this.dialogVisible = true;
            this.dialogTitle = "编辑";
            this.contract = contract;
        },
        // 修改提交
        edit(){
            updateContract(this.contract).then(res=>{
                this.$message({
                    message:'修改成功',
                    type:'success'    
                });
                this.list();
                this.dialogVisible = false;
            })
        },
        // 删除提交
        remove(contractId){
            this.contract = {
                id :contractId
            }
            // 删除确认
            if (window.confirm ("你确定真的要删除吗")){
                deleteContract(this.contract).then(res=>{
                    this.$message({
                    message:'删除成功',
                    type:'success'    
                });
                    this.list();
                })
            }
        },
        list(){
            let data = {"keyword":this.keyword, "page":this.page,"size":this.size}
            listContract(data).then(res=>{
                this.tableData = res.data.records;
                this.total = res.data.total;
            })
        },
        handleSizeChange(size){
            this.page = 1;
            this.size = size;
            this.list();
        },
        handleCurrentChange(page){
            this.page = page;
            this.list();
        },
    }
}
</script>