
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form
          ref="queryForm"
          :model="queryParams"
          :inline="true"
          label-width="68px"
        >
          <el-form-item>
            <el-button
              type="primary"
              icon="el-icon-search"
              size="mini"
              @click="handleQuery"
              >搜索</el-button
            >
            <el-button icon="el-icon-refresh" size="mini" @click="resetQuery"
              >重置</el-button
            >
          </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:articleA:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
              >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:articleA:edit']"
              type="success"
              icon="el-icon-edit"
              size="mini"
              :disabled="single"
              @click="handleUpdate"
              >修改
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:articleA:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
              >删除
            </el-button>
          </el-col>
        </el-row>

        <el-table
          v-loading="loading"
          :data="articleAList"
          @selection-change="handleSelectionChange"
        >
          <el-table-column
            type="selection"
            width="55"
            align="center"
          /><el-table-column
            label="用户名"
            align="center"
            prop="username"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="地区"
            align="center"
            prop="area"
            :formatter="areaFormat"
            width="100"
          >
            <template slot-scope="scope">
              {{ areaFormat(scope.row) }}
            </template> </el-table-column
          ><el-table-column
            label="型号"
            align="center"
            prop="type"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="状态"
            align="center"
            prop="status"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="备注"
            align="center"
            prop="remark"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="操作"
            align="center"
            class-name="small-padding fixed-width"
          >
            <template slot-scope="scope">
              <el-popconfirm
                class="delete-popconfirm"
                title="确认要修改吗?"
                confirm-button-text="修改"
                @onConfirm="handleUpdate(scope.row)"
              >
                <el-button
                  slot="reference"
                  v-permisaction="['admin:articleA:edit']"
                  size="mini"
                  type="text"
                  icon="el-icon-edit"
                  >修改
                </el-button>
              </el-popconfirm>
              <el-popconfirm
                class="delete-popconfirm"
                title="确认要删除吗?"
                confirm-button-text="删除"
                @onConfirm="handleDelete(scope.row)"
              >
                <el-button
                  slot="reference"
                  v-permisaction="['admin:articleA:remove']"
                  size="mini"
                  type="text"
                  icon="el-icon-delete"
                  >删除
                </el-button>
              </el-popconfirm>
            </template>
          </el-table-column>
        </el-table>

        <pagination
          v-show="total > 0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="500px">
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-form-item label="用户名" prop="username">
              <el-input v-model="form.username" placeholder="用户名" />
            </el-form-item>
            <el-form-item label="地区" prop="area">
              <el-select
                v-model="form.area"
                placeholder="请选择"
                @change="$forceUpdate()"
              >
                <el-option
                  v-for="item in areaOptionstemp"
                  :key="item.id"
                  :label="item.area"
                  :value="item.id"
                />
              </el-select>
              <!-- <el-input v-model="form.area" placeholder="地区" /> -->
            </el-form-item>
            <el-form-item label="型号" prop="type">
              <el-input v-model="form.type" placeholder="型号" />
            </el-form-item>
            <el-form-item label="状态" prop="status">
              <el-input v-model="form.status" placeholder="状态" />
            </el-form-item>
            <el-form-item label="备注" prop="remark">
              <el-input v-model="form.remark" placeholder="备注" />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import {
  addArticleA,
  delArticleA,
  getArticleA,
  listArticleA,
  updateArticleA,
} from "@/api/admin/article-a";

import { listArticle } from "@/api/admin/article";
export default {
  name: "ArticleA",
  components: {},
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,
      isEdit: false,
      // 类型数据字典
      typeOptions: [],
      articleAList: [],

      // 关系表类型
      areaOptions: [],
      areaOptionstemp: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {},
    };
  },
  created() {
    this.getList();
    this.getArticleItems();
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true;
      listArticleA(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.articleAList = response.data.list;
          this.total = response.data.count;
          this.loading = false;
        }
      );
    },
    // 取消按钮
    cancel() {
      this.open = false;
      this.reset();
    },
    // 表单重置
    reset() {
      this.form = {
        id: undefined,
        username: undefined,
        area: undefined,
        type: undefined,
        status: undefined,
        remark: undefined,
      };
      this.resetForm("form");
    },
    getImgList: function () {
      this.form[this.fileIndex] =
        this.$refs["fileChoose"].resultList[0].fullUrl;
    },
    fileClose: function () {
      this.fileOpen = false;
    },
    areaFormat(row) {
      return this.selectItemsLabel(this.areaOptions, row.area);
    },
    // 关系
    getArticleItems() {
      this.getItems(listArticle, undefined).then((res) => {
        this.areaOptions = this.setItems(res, "area", "area");
      });
    },
    // 文件
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1;
      this.getList();
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = [];
      this.resetForm("queryForm");
      this.handleQuery();
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset();
      this.open = true;
      this.title = "添加ArticleA";
      this.isEdit = false;
      this.areaOptionstemp = [];

      listArticle(
        this.addDateRange(
          {
            pageIndex: 1,
            pageSize: 10000,
          },
          this.dateRange
        )
      ).then((response) => {
        this.areaOptions = response.data.list;
        let datatemp = response.data.list;
        let countt = response.data.count;
        let tempnum = [];
        //通过循环去除area相同的对象
        for (let index = 0; index < countt; index++) {
          if (tempnum.indexOf(datatemp[index].area) == -1) {
            // console.log(tempnum.indexOf(datatemp[index].area));
            tempnum.push(datatemp[index].area);
            this.areaOptionstemp.push(response.data.list[index]);
          }
        }
      });
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map((item) => item.id);
      this.single = selection.length !== 1;
      this.multiple = !selection.length;
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset();
      const id = row.id || this.ids;
      getArticleA(id).then((response) => {
        this.form = response.data;
        this.open = true;
        this.title = "修改ArticleA";
        this.isEdit = true;
      });
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateArticleA(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addArticleA(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          }
        }
      });
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      var Ids = (row.id && [row.id]) || this.ids;

      this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(function () {
          return delArticleA({ ids: Ids });
        })
        .then((response) => {
          if (response.code === 200) {
            this.msgSuccess(response.msg);
            this.open = false;
            this.getList();
          } else {
            this.msgError(response.msg);
          }
        })
        .catch(function () {});
    },
  },
};
</script>
