<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true" >
        <el-row>
          <el-form-item label="名称">
            <el-input v-model.trim="searchParams.name"></el-input>
          </el-form-item>
          <el-form-item label="性别">
            <el-select v-model.trim="searchParams.sex">
              <el-option label="全部" value=""></el-option>
              <el-option
                v-for="item in sexList"
                :key="item.value"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="fetchData()">搜索</el-button>
          </el-form-item>
        </el-row>
      </el-form>
    </div>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="ID" width="95">
        <template slot-scope="scope">
          {{ scope.$index }}
        </template>
      </el-table-column>
      <el-table-column label="名称" >
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="年龄"  align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.age }}</span>
        </template>
      </el-table-column>
      <el-table-column label="角色"  align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.role=='1'">管理员</span>
          <span v-else>普通用户</span>
        </template>
      </el-table-column>
      <el-table-column label="描述"  align="center">
        <template slot-scope="scope">
          {{ scope.row.introduction }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="性别" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.sex | statusFilter">
            <span v-if = "scope.row.sex == 1">男</span>
            <span v-else>女</span>
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="创建日期" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.create_at | formatTime}}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" width="220">
        <template slot-scope="scope">
          <el-row>
            <el-button type="info" @click="toEdit(scope.row.id)">编辑</el-button>
            <el-button type="danger" @click="remove(scope.row.id)">删除</el-button>
          </el-row>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
  import { getList,Delete } from '@/api/demo'

  export default {
    filters: {
      statusFilter(status) {
        var res = 'success'
        if (status == "2"){
          res = 'danger'
        }
        // deleted: 'gray'
        return res
      }
    },
    data() {
      return {
        searchParams: {
          page_size: 20,
          page: 1,
          name: '',
          sex: ""
        },
        sexList: [
          {
            value: '1',
            label: '男'
          },
          {
            value: '2',
            label: '女'
          }
        ],
        list: null,
        listLoading: true
      }
    },
    created() {
      this.fetchData()
    },
    methods: {
      fetchData() {
        this.listLoading = true
        getList(this.searchParams).then(response => {
          this.list = response.data.result
          this.listLoading = false
        })
      },
      remove (id) {
        this.$appConfirm(() => {
          Delete(id).then( response => {
            this.refresh()
          })
        })
      },
      refresh () {
        this.fetchData(() => {
          this.$message.success('刷新成功')
        })
      },
      toEdit (id) {
        let path = ''
        path = '/user/add?id='+id
        // if (item === null) {
        //   path = '/user/add'
        // } else {
        //   path = `/user/add/${item.id}`
        // }
        this.$router.push(path)
      }
    }
  }
</script>
