
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
          <el-form-item label="光纤编号" prop="fabric"
            >
            <!-- <el-input
              v-model="queryParams.fabric"
              placeholder="请输入光纤编号"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <el-select
              v-model="queryParams.fabric"
              placeholder="请输入光纤编号"
              allow-create
              clearable
              filterable
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in fyFabricList"
                :key="item.id"
                :label="item.fabric"
                :value="item.fabric"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="光栅数量" prop="raster"
            >
            <el-input
              v-model="queryParams.raster"
              placeholder="请输入光栅数量"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item>
          <el-form-item label="说明" prop="instruction"
            ><el-input
              v-model="queryParams.instruction"
              placeholder="请输入说明"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item>
          <el-form-item label="备注" prop="remark"
            ><el-input
              v-model="queryParams.remark"
              placeholder="请输入备注"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item>

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
              v-permisaction="['admin:fyFabric:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
              >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:fyFabric:edit']"
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
              v-permisaction="['admin:fyFabric:remove']"
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
          :data="fyFabricList"
          @selection-change="handleSelectionChange"
        >
          <el-table-column
            type="selection"
            width="55"
            align="center"
          /><el-table-column
            label="光纤编号"
            align="center"
            prop="fabric"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="光栅数量"
            align="center"
            prop="raster"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="说明"
            align="center"
            prop="instruction"
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
              <el-button
                slot="reference"
                v-permisaction="['admin:fyFabric:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
                >修改
              </el-button>

              <el-button
                slot="reference"
                v-permisaction="['admin:fyFabric:remove']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
                >删除
              </el-button>
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
            <el-form-item label="光纤编号" prop="fabric">
              <el-input v-model="form.fabric" placeholder="光纤编号" />
            </el-form-item>
            <el-form-item label="光栅数量" prop="raster">
              <el-input v-model="form.raster" placeholder="光栅数量" />
            </el-form-item>
            <el-form-item label="说明" prop="instruction">
              <el-input v-model="form.instruction" placeholder="说明" />
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
  addFyFabric,
  delFyFabric,
  getFyFabric,
  listFyFabric,
  updateFyFabric,
} from "@/api/admin/fy-fabric";

export default {
  name: "FyFabric",
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
      fyFabricList: [],

      // 关系表类型

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        fabric: undefined,
        raster: undefined,
        instruction: undefined,
        remark: undefined,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        fabric: [
          { required: true, message: "光纤编号不能为空", trigger: "blur" },
        ],
        raster: [
          { required: true, message: "光栅数量不能为空", trigger: "blur" },
        ],
        instruction: [
          { required: false, message: "说明不能为空", trigger: "blur" },
        ],
        remark: [{ required: false, message: "备注不能为空", trigger: "blur" }],
      },
    };
  },
  created() {
    this.getList();
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true;
      listFyFabric(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.fyFabricList = response.data.list;
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
        fabric: undefined,
        raster: undefined,
        instruction: undefined,
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
    // 关系
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
      this.title = "添加光纤";
      this.isEdit = false;
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map((item) => item.id);
      this.single = selection.length !== 1;
      this.multiple = !selection.length;
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      console.log("execute handleupdate");
      this.reset();
      const id = row.id || this.ids;
      getFyFabric(id).then((response) => {
        this.form = response.data;
        this.open = true;
        this.title = "修改光纤";
        this.isEdit = true;
      });
    },
    /** 提交按钮 */
    submitForm: function () {
      this.form.raster = parseInt(this.form.raster);
      this.$refs["form"].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateFyFabric(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addFyFabric(this.form).then((response) => {
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
          return delFyFabric({ ids: Ids });
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
    ceshishuju() {
      console.log("测试数据");
    },
  },
};
</script>
