
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
          <el-form-item label="地区" prop="area">
            <!-- <el-input
              v-model="queryParams.area"
              placeholder="请输入地区"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <el-select
              v-model="queryParams.area"
              placeholder="请输入地区"
              clearable
              filterable
              @change="$forceUpdate()"
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in fyAreaList"
                :key="item.id"
                :label="item.area"
                :value="item.area"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="负责人" prop="username">
            <!-- <el-input
              v-model="queryParams.username"
              placeholder="请输入负责人"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <el-select
              v-model="queryParams.username"
              placeholder="负责人"
              clearable
              allow-create
              filterable
              size="small"
              @change="$forceUpdate()"
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in userOptionsForm"
                :key="item.id"
                :label="item.username"
                :value="item.username"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="设备数量" prop="equipcount"
            ><el-input
              v-model="queryParams.equipcount"
              placeholder="请输入设备数量"
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
              v-permisaction="['admin:fyArea:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
              >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:fyArea:edit']"
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
              v-permisaction="['admin:fyArea:remove']"
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
              v-permisaction="['admin:fyArea:refresh']"
              type="info"
              icon="el-icon-refresh"
              size="mini"
              @click="getList"
              >刷新
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              type="primary"
              icon="el-icon-edit"
              size="mini"
              @click="ceshishuju"
              >测试数据
            </el-button>
          </el-col>
        </el-row>

        <el-table
          v-loading="loading"
          :data="fyAreaList"
          stripe
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="序号" type="index" align="center" />
          <el-table-column
            label="地区"
            align="center"
            prop="area"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="负责人"
            align="center"
            prop="username"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="设备数量"
            align="center"
            prop="equipcount"
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
                v-permisaction="['admin:fyArea:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
                >修改
              </el-button>

              <el-button
                slot="reference"
                v-permisaction="['admin:fyArea:remove']"
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
            <el-form-item label="地区" prop="area">
              <el-input v-model="form.area" placeholder="地区(不能重复)" />
            </el-form-item>
            <el-form-item label="负责人" prop="username">
              <!-- <el-input v-model="form.username" placeholder="用户" /> -->
              <el-select
                v-model="form.username"
                placeholder="请选择负责人"
                clearable
                filterable
                multiple
                @change="$forceUpdate()"
              >
                <el-option
                  v-for="item in userOptionsForm"
                  :key="item.id"
                  :label="item.username"
                  :value="item.username"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="设备数量" prop="equipcount">
              <el-input
                v-model="form.equipcount"
                placeholder="设备数量(不能为空)"
              />
            </el-form-item>
            <el-form-item label="备注" prop="remark">
              <el-input v-model="form.remark" placeholder="备注" />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="ceshishuju">测试数据</el-button>
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
  addFyArea,
  delFyArea,
  getFyArea,
  listFyArea,
  updateFyArea,
} from "@/api/admin/fy-area";
import { listFyUser } from "@/api/admin/fy-user";

export default {
  name: "FyArea",
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
      fyAreaList: [],

      // 关系表类型
      userOptionsForm: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        area: undefined,
        username: undefined,
        equipcount: undefined,
        remark: undefined,
      },
      // 表单参数
      form: {},
      //表单唯一验证
      formunique: [],
      formunique2: [],
      // 表单校验
      rules: {
        area: [
          { required: true, message: "地区不能为空", trigger: "blur" },
          //表单唯一性校验
          {
            validator: (rule, value, callback) => {
              // console.log("rule:", rule, "value:", value, "唯一表:", this.formunique);
              if (this.formunique.includes(value)) {
                callback(new Error("字段重复"));
              } else {
                callback();
              }
            },
            trigger: "blur",
          },
        ],
        username: [
          { required: false, message: "负责人不能为空", trigger: "blur" },
        ],
        equipcount: [
          { required: false, message: "设备数量不能为空", trigger: "blur" },
        ],
        remark: [{ required: false, message: "备注不能为空", trigger: "blur" }],
      },
      formrulestemp: {},
    };
  },
  created() {
    this.getList();
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true;
      listFyArea(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.fyAreaList = response.data.list;
          this.total = response.data.count;
          //表单唯一验证对比数据存储
          // this.formunique = [];
          // this.formunique2 = [];
          // for (let index = 0; index < this.total; index++) {
          //   this.formunique.push(this.fyAreaList[index].area);
          // }
          this.formunique = this.fyAreaList.map((item) => item.area);
          this.formunique2 = this.formunique.slice();

          // console.log(
          //   "长度1",
          //   this.formunique.length,
          //   "长度2",
          //   this.formunique2.length
          // );
        }
      );

      //表单获取负责人数据
      listFyUser(
        this.addDateRange(
          {
            pageIndex: 1,
            pageSize: 10000,
          },
          this.dateRange
        )
      ).then((response) => {
        this.userOptionsForm = response.data.list;
        this.loading = false;
      });
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
        area: undefined,
        username: undefined,
        equipcount: undefined,
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
      this.form.equipcount = "0";
      this.title = "添加地区";
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
      this.reset();
      const id = row.id || this.ids;
      getFyArea(id).then((response) => {
        this.form = response.data;
        //表单唯一性验证对比数据更新
        this.formunique = this.formunique2.slice(); //复制模板数组
        const tempareaunique = response.data.area;
        const indexareaunique = this.formunique.indexOf(tempareaunique);
        this.formunique.splice(indexareaunique, 1);
        // this.formunique2.splice(indexareaunique, 0,tempareaunique)

        // console.log("username", this.form.username);
        //以tempusername作为媒介，对多选下拉框进行数据回显
        if (this.form.username !== "" && this.form.username !== undefined) {
          const tempusername = this.form.username.split(",");
          this.form.username = tempusername;
        } else {
          this.form.username = undefined;
        }

        this.open = true;
        this.title = "修改地区";
        this.isEdit = true;
      });
    },

    /** 提交按钮 */
    submitForm: function () {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          // console.log(this.form);
          //传递equipcount需要使用字符串类型
          this.form.equipcount = this.form.equipcount.toString();
          //如果username不为空则返回拼接字符串
          // console.log("username", this.form.username);
          if (this.form.username !== ("" || undefined)) {
            this.form.username = this.form.username.join(",");
          }
          // console.log(this.form.username);
          if (this.form.id !== undefined) {
            updateFyArea(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addFyArea(this.form).then((response) => {
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
          return delFyArea({ ids: Ids });
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
      console.log("表单", this.form);
      console.log("表单唯一1", this.formunique.length);
      console.log("表单唯一2", this.formunique2.length);
    },
  },
};
</script>
