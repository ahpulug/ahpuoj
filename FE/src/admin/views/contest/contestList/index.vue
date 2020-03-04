<template lang="pug">
.admin-content
  table-tools
    template(#tool)
      router-link(:to="{name:'adminAddContest'}")
        el-button(icon="el-icon-plus") 新建
    template(#search)
      el-input(placeholder="搜索竞赛&作业名称", style="max-width:20em", v-model="queryParam", @keyup.enter.native="handleSearchByParam",maxlength="20", clearable)
        el-button(slot="append",icon="el-icon-search",@click="handleSearchByParam")
  .content__table__wrapper
    el-table(:data="tableData" style="width: 100%", v-loading="loading",border)
      el-table-column(label="ID", prop="id", width="180")
      el-table-column(label="名称",min-width="300")
        template(slot-scope="scope")
          a(:href="`/contest/${scope.row.id}`" target="_blank") {{scope.row.name}}
      el-table-column(label="状态", width="240")
        template(slot-scope="scope")
          el-tag(:type="scope.row.defunct == 0 ? 'success':'danger'",effect="dark") {{scope.row.defunct == 0?'启用':'保留'}}
          el-tag(:type="scope.row.private == 0 ? 'success':'danger'",effect="dark") {{scope.row.private == 0?'公开':'私有'}}
          el-tag(:type="scope.row.team_mode == 0 ? 'success':'primary'",effect="dark") {{scope.row.team_mode == 0?'个人':'团队'}}
      el-table-column(label="操作", width="300")
        template(slot-scope="scope")
          el-button(size="mini", type="primary", @click="$router.push({name:'adminEditContest',params:{id:scope.row.id}})") 编辑
          el-button(size="mini", @click="$router.push({name: (scope.row.team_mode == 0)? 'adminContestManage':'adminContestTeamManage' ,params:{id:scope.row.id}})", :disabled="scope.row.private == 0") 人员
          el-button(size="mini", :type="scope.row.defunct == 0?'danger':'success'", @click="handleToggleContestStatus(scope.row)") {{scope.row.defunct == 0?'保留':'启用'}}
          el-button(size="mini", type="danger", @click="handleDeleteContest(scope.row)") 删除
  .content__pagination__wrapper
    el-pagination(@size-change="handleSizeChange",@current-change="fetchDataList",:current-page.sync="currentPage",:page-sizes="[10, 20, 30, 40,50]",:page-size="10",layout="total, sizes, prev, pager, next, jumper",:total="total")
</template>

<script>
import {
  getContestList,
  deleteContest,
  toggleContestStatus,
} from 'admin/api/contest';

export default {
  name: 'adminContestList',
  data() {
    return {
      loading: true,
      currentPage: 1,
      currentRowId: 0,
      perpage: 10,
      total: 0,
      queryParam: '',
      tableData: [],
    };
  },
  activated() {
    this.fetchDataList();
  },
  methods: {
    async fetchDataList() {
      this.loading = true;
      try {
        const res = await getContestList(
          this.currentPage,
          this.perpage,
          this.queryParam,
        );
        const { data } = res;
        setTimeout(() => {
          this.tableData = data.data;
          this.total = data.total;
          this.loading = false;
        }, 200);
      } catch (err) {
        console.log(err);
      }
    },
    handleSearchByParam() {
      this.currentPage = 1;
      this.loading = true;
      this.fetchDataList();
    },
    handleSizeChange(val) {
      this.perpage = val;
      this.fetchDataList();
    },
    async handleToggleContestStatus(row) {
      const msg = `确认要${row.defunct === 0 ? '保留' : '启用'}竞赛${row.name}吗?`;
      try {
        await this.$confirm(msg, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        });
        try {
          const res = await toggleContestStatus(row.id);
          this.$message({
            type: 'success',
            message: res.data.message,
          });

          row.defunct = !row.defunct;
        } catch (err) {
          this.$message({
            type: 'error',
            message: err.response.data.message,
          });
        }
      } catch (err) {
        this.$message({
          type: 'info',
          message: '已取消操作',
        });
      }
    },
    async handleDeleteContest(row) {
      try {
        await this.$confirm(`确认要删除竞赛${row.name}吗?`, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        });
        try {
          const res = await deleteContest(row.id);
          this.$message({
            type: 'success',
            message: res.data.message,
          });
          // 删除最后一页最后一条记录，如果不是第一页，则当前页码-1
          if (this.tableData.length === 1) {
            if (this.currentPage > 1) {
              this.currentPage -= 1;
            }
          }
          this.fetchDataList();
        } catch (err) {
          this.$message({
            type: 'error',
            message: err.response.data.message,
          });
        }
      } catch (err) {
        this.$message({
          type: 'info',
          message: '已取消删除',
        });
      }
    },
  },
};
</script>

<style lang="scss" scoped>
</style>