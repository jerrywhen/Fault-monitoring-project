
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
          <el-form-item label="用户名" prop="username">
            <!-- <el-input
              v-model="queryParams.username"
              placeholder="请输入用户名"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <el-select
              v-model="queryParams.username"
              placeholder="请输入用户名"
              clearable
              filterable
              size="small"
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in fyUserList"
                :key="item.id"
                :label="item.username"
                :value="item.username"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="负责地区" prop="area">
            <!-- <el-input
              v-model="queryParams.area"
              placeholder="请输入负责地区"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <el-select
              v-model="queryParams.area"
              placeholder="请输入负责地区"
              clearable
              filterable
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in areaOptionsForm"
                :key="item.id"
                :label="item.area"
                :value="item.area"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="人员状态" prop="userstatus">
            <!-- <el-input
              v-model="queryParams.userstatus"
              placeholder="请输入人员状态"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <el-radio-group v-model="queryParams.userstatus">
              <el-radio-button label="在职">在职</el-radio-button>
              <el-radio-button label="停职">停职</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="职责" prop="duty"
            ><el-input
              v-model="queryParams.duty"
              placeholder="请输入职责"
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
              v-permisaction="['admin:fyUser:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
              >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:fyUser:edit']"
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
              v-permisaction="['admin:fyUser:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
              >删除
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:fyUser:refresh']"
              type="info"
              icon="el-icon-refresh"
              size="mini"
              @click="getList"
              >刷新
            </el-button>
          </el-col>
          <!-- <el-col :span="1.5">
            <el-button type="primary" size="mini" @click="ceshishuju"
              >测试数据
            </el-button>
          </el-col> -->
          <el-col :span="1.5">
            <el-button type="primary" size="mini" @click="handleopen2"
              >修改每页数据量</el-button
            >
          </el-col>
        </el-row>

        <el-table
          v-loading="loading"
          :data="fyUserList"
          :row-class-name="tableRowClassUserStatus"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="序号" type="index" align="center" />
          <el-table-column
            label="用户名"
            align="center"
            prop="username"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="负责地区"
            align="center"
            prop="area"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="人员状态"
            align="center"
            prop="userstatus"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="职责"
            align="center"
            prop="duty"
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
                v-permisaction="['admin:fyUser:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
                >修改
              </el-button>
              <el-button
                slot="reference"
                v-permisaction="['admin:fyUser:remove']"
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
            <el-form-item label="用户名" prop="username">
              <el-input
                v-model="form.username"
                placeholder="用户名(不能重复)"
              />
            </el-form-item>
            <el-form-item label="负责地区" prop="area">
              <!-- <el-input v-model="form.area" placeholder="负责地区" /> -->
              <el-select
                v-model="form.area"
                placeholder="请选择地区"
                clearable
                filterable
                multiple
                @change="$forceUpdate()"
                ><el-option
                  v-for="item in areaOptionsForm"
                  :key="item.id"
                  :label="item.area"
                  :value="item.area"
                />
              </el-select>
              <!-- <el-select
                v-model="form.area"
                placeholder="请选择地区"
                clearable
                filterable
                @focus="cleararea()"
                @blur="selectBlur()"
                @clear="selectClear()"
                @change="selectChange()"
              >
                <el-option
                  v-for="item in areaOptionsForm"
                  :key="item.id"
                  :label="item.area"
                  :value="item.area"
                />
              </el-select> -->
            </el-form-item>
            <el-form-item label="人员状态" prop="userstatus">
              <!-- <el-input v-model="form.userstatus" placeholder="人员状态" /> -->
              <el-radio-group v-model="form.userstatus">
                <el-radio label="在职">在职</el-radio>
                <el-radio label="停职">停职</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="职责" prop="duty">
              <el-input v-model="form.duty" placeholder="职责" />
            </el-form-item>
            <el-form-item label="备注" prop="remark">
              <el-input v-model="form.remark" placeholder="备注" />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <!-- <el-button type="primary" @click="ceshishuju">测试数据</el-button> -->
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>

        <!-- open2表单-->
        <el-dialog :title="title" :visible.sync="open2" width="400px">
          <el-form ref="form2" :model="form2" label-width="100px">
            <el-form-item label="每页数据量" prop="pageSize">
              <el-input v-model="form2.pageSize" placeholder="请输入pageSize" />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <!-- <el-button type="primary" @click="ceshishuju">测试数据</el-button> -->
            <el-button type="primary" @click="submitForm2">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import {
  addFyUser,
  delFyUser,
  getFyUser,
  listFyUser,
  updateFyUser,
} from "@/api/admin/fy-user";
import { listFyArea } from "@/api/admin/fy-area";

export default {
  name: "FyUser",
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
      open2: false,
      isEdit: false,
      // 类型数据字典
      typeOptions: [],
      fyUserList: [],

      // 关系表类型
      areaOptionsForm: [],

      //test
      test1: [],
      test2: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        username: undefined,
        area: undefined,
        userstatus: undefined,
        duty: undefined,
        remark: undefined,
      },
      // 表单参数
      form: {},
      form2: {},
      // 表单校验
      rules: {
        username: [
          { required: true, message: "用户名不能为空", trigger: "blur" },
        ],
        area: [
          { required: false, message: "负责地区不能为空", trigger: "blur" },
        ],
        userstatus: [
          { required: false, message: "人员状态不能为空", trigger: "blur" },
        ],
        duty: [{ required: false, message: "职责不能为空", trigger: "blur" }],
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
      listFyUser(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.fyUserList = response.data.list;
          this.total = response.data.count;
        }
      );
      //表单获取地区数据
      listFyArea(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
      ).then((response) => {
        this.areaOptionsForm = response.data.list;
        this.loading = false;
      });
    },
    // 取消按钮
    cancel() {
      this.open = false;
      this.open2 = false;
      this.reset();
    },
    // 表单重置
    reset() {
      this.form = {
        id: undefined,
        username: undefined,
        area: undefined,
        userstatus: undefined,
        duty: undefined,
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
      this.title = "添加负责人";
      this.isEdit = false;
    },
    /** 修改sizepage按钮操作 */
    handleopen2() {
      this.open2 = true;
      this.title = "修改每页数据量";
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
      getFyUser(id).then((response) => {
        this.form = response.data;
        //以temparea作为媒介，对多选下拉框进行数据回显
        const temparea = this.form.area.split(",");
        this.form.area = temparea;
        // console.log("修改按钮,修饰后的area", temparea);
        // console.log("修改按钮获得的form", this.form);
        // console.log("修改按钮获得的form", this.form);
        // console.log("修改按钮获得的forms.area",this.form.area.split(","));
        this.open = true;
        this.title = "修改负责人";
        this.isEdit = true;
      });
    },
    //表单重构area数据
    cleararea() {
      this.form.area = "";
      this.$forceUpdate();
    },
    //表单测试数据按钮
    ceshishuju() {
      console.log("表单", this.form);
      console.log("datarange", this.dateRange);
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          //如果area不为空则返回拼接字符串
          if (this.form.area !== ("" || undefined)) {
            this.form.area = this.form.area.join(",");
          }
          if (this.form.id !== undefined) {
            updateFyUser(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addFyUser(this.form).then((response) => {
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
    /** 提交按钮2 */
    submitForm2: function () {
      this.$refs["form2"].validate((valid) => {
        if (valid) {
          this.queryParams.pageSize = parseInt(this.form2.pageSize);
          this.open2 = false;
          this.getList();
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
          return delFyUser({ ids: Ids });
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
    //表单按状态颜色停职
    tableRowClassUserStatus({ row }) {
      // console.log(row);
      if (row.userstatus == "停职") {
        console.log("停职");
        return "suspension-row";
      }
      return "";
    },
  },

  // selectBlur(e) {
  //   console.log("eeeeeeeee");
  //   // 意见类型
  //   if (e.target.value !== "") {
  //     this.value = e.target.value + "(其他)";
  //     this.$forceUpdate(); // 强制更新
  //   }
  // },
  // selectClear() {
  //   this.value = "";
  //   this.$forceUpdate();
  // },
  // selectChange(val) {
  //   this.value = val;
  //   this.$forceUpdate();
  // },
};
</script>


<style>
.el-table .suspension-row {
  background: #c0c0c050;
}
/* .el-table .warning-row {
  background: rgba(255, 166, 0, 0.15);
} */
</style>