
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <!-- 使用ref属性指定表单的引用名称为queryForm，以便在后续代码中可以通过this.$refs.queryForm访问到该表单实例。
使用:model属性绑定表单的数据模型为queryParams，以便在表单项中可以实时同步用户输入的内容。
使用:inline="true"属性指定表单项为行内布局，以便在页面上占用更小的空间。 -->
        <el-form
          ref="queryForm"
          :model="queryParams"
          :inline="true"
          label-width="68px"
        >
          <el-form-item label="设备名" prop="equipname">
            <!-- <el-input
              v-model="queryParams.equipname"
              placeholder="请输入设备名"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <el-select
              v-model="queryParams.equipname"
              placeholder="请输入设备名"
              clearable
              filterable
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in fyEquipmentList"
                :key="item.id"
                :label="item.equipname"
                :value="item.equipname"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="设备编码" prop="equipcode">
            <el-select
              v-model="queryParams.equipcode"
              placeholder="请输入设备编码"
              clearable
              filterable
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in fyEquipmentList"
                :key="item.id"
                :label="item.equipcode"
                :value="item.equipcode"
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
              placeholder="请输入负责人"
              clearable
              filterable
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in userOptionsForm"
                :key="item.id"
                :label="item.username"
                :value="item.username"
              />
            </el-select>
          </el-form-item>
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
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in areaOptionsForm"
                :key="item.id"
                :label="item.area"
                :value="item.area"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="型号" prop="type">
            <!-- <el-input
              v-model="queryParams.type"
              placeholder="请输入型号"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <el-select
              v-model="queryParams.type"
              placeholder="请输入型号"
              clearable
              filterable
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in typeOptionsForm"
                :key="item.id"
                :label="item.type"
                :value="item.type"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="状态" prop="status">
            <!-- <el-input
              v-model="queryParams.status"
              placeholder="请输入状态"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <el-select
              v-model="queryParams.status"
              placeholder="请输入状态"
              clearable
              filterable
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in statusOptionsForm"
                :key="item.id"
                :label="item.status"
                :value="item.status"
              />
            </el-select>
          </el-form-item>
          <!-- <el-form-item label="故障次数" prop="faultcount"
            ><el-input
              v-model="queryParams.faultcount"
              placeholder="请输入故障次数(大于搜索)"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item> -->
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
              v-permisaction="['admin:fyEquipment:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
              >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:fyEquipment:edit']"
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
              v-permisaction="['admin:fyEquipment:remove']"
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
              v-permisaction="['admin:fyEquipment:refresh']"
              type="info"
              icon="el-icon-refresh"
              size="mini"
              @click="getList"
              >刷新
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button type="primary" size="mini" @click="handleopen2"
              >修改每页数据量</el-button
            >
          </el-col>
        </el-row>

        <el-table
          v-loading="loading"
          :data="fyEquipmentList"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="序号" type="index" align="center" />
          <el-table-column
            label="设备名"
            align="center"
            prop="equipname"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="设备编码"
            align="center"
            prop="equipcode"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="负责人"
            align="center"
            prop="username"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="地区"
            align="center"
            prop="area"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="型号"
            align="center"
            prop="type"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="状态"
            align="center"
            prop="status"
            :show-overflow-tooltip="true"
          />
          <!-- <el-table-column
            label="故障次数"
            align="center"
            prop="faultcount"
            :show-overflow-tooltip="true"
          /> -->
          <el-table-column
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
                v-permisaction="['admin:fyEquipment:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
                >修改
              </el-button>

              <el-button
                slot="reference"
                v-permisaction="['admin:fyEquipment:remove']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
                >删除
              </el-button>

              <el-button
                slot="reference"
                v-permisaction="['admin:fyEquipment:edit']"
                size="mini"
                type="text"
                icon="el-icon-s-marketing"
                @click="handleEditTable(scope.row)"
                >实时数据
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
            <el-form-item label="设备名" prop="equipname">
              <el-input
                v-model="form.equipname"
                placeholder="设备名(不能重复)"
              />
            </el-form-item>
            <el-form-item label="设备编码" prop="equipcode">
              <el-input v-model="form.equipcode" placeholder="设备编码" />
            </el-form-item>
            <el-form-item label="负责人" prop="username">
              <!-- <el-input v-model="form.username" placeholder="负责人" /> -->
              <el-select
                v-model="form.username"
                placeholder="请选择负责人"
                clearable
                filterable
                @change="$forceUpdate()"
                ><el-option
                  v-for="item in userOptionsForm"
                  :key="item.id"
                  :label="item.username"
                  :value="item.username"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="地区" prop="area">
              <!-- <el-input v-model="form.area" placeholder="地区" /> -->
              <el-select
                v-model="form.area"
                placeholder="请选择地区"
                clearable
                filterable
                @change="$forceUpdate()"
                ><el-option
                  v-for="item in areaOptionsForm"
                  :key="item.id"
                  :label="item.area"
                  :value="item.area"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="型号" prop="type">
              <!-- <el-input v-model="form.type" placeholder="型号" /> -->
              <el-select
                v-model="form.type"
                placeholder="请选择型号"
                clearable
                filterable
                @change="$forceUpdate()"
                ><el-option
                  v-for="item in typeOptionsForm"
                  :key="item.id"
                  :label="item.type"
                  :value="item.type"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="状态" prop="status">
              <!-- <el-input v-model="form.status" placeholder="状态" /> -->
              <el-select
                v-model="form.status"
                placeholder="请选择状态"
                clearable
                filterable
                @change="$forceUpdate()"
                ><el-option
                  v-for="item in statusOptionsForm"
                  :key="item.id"
                  :label="item.status"
                  :value="item.status"
                />
              </el-select>
            </el-form-item>
            <!-- <el-form-item label="故障次数" prop="faultcount">
              <el-input v-model="form.faultcount" placeholder="故障次数" />
            </el-form-item> -->
            <el-form-item label="备注" prop="remark">
              <el-input v-model="form.remark" placeholder="备注" />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
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
            <el-button type="primary" @click="submitForm2">确 定</el-button>
            <el-button @click="cancel2">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import {
  addFyEquipment,
  delFyEquipment,
  getFyEquipment,
  listFyEquipment,
  updateFyEquipment,
} from "@/api/admin/fy-equipment";
import { listFyUser } from "@/api/admin/fy-user";
import { listFyArea } from "@/api/admin/fy-area";
import { listFyType } from "@/api/admin/fy-type";
import { listFyStatus } from "@/api/admin/fy-status";

export default {
  name: "FyEquipment",
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
      fyEquipmentList: [],

      // 关系表类型
      areaOptionsForm: [],
      userOptionsForm: [],
      statusOptionsForm: [],
      typeOptionsForm: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        equipname: undefined,
        equipcode: undefined,
        username: undefined,
        area: undefined,
        type: undefined,
        status: undefined,
        faultcount: undefined,
        remark: undefined,
      },
      // 表单参数
      form: {},
      form2: {},
      // 表单校验
      rules: {
        equipname: [
          { required: true, message: "设备名不能为空", trigger: "blur" },
        ],
        equipcode: [
          { required: true, message: "设备编码不能为空", trigger: "blur" },
        ],
        username: [
          { required: false, message: "负责人不能为空", trigger: "blur" },
        ],

        area: [{ required: false, message: "地区不能为空", trigger: "blur" }],
        type: [{ required: false, message: "型号不能为空", trigger: "blur" }],
        status: [{ required: false, message: "状态不能为空", trigger: "blur" }],
        faultcount: [
          { required: false, message: "故障次数不能为空", trigger: "blur" },
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
      listFyEquipment(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.fyEquipmentList = response.data.list;
          this.total = response.data.count;
          this.loading = false;
        }
      );
      //表单获取地区数据
      listFyArea(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
      ).then((response) => {
        this.areaOptionsForm = response.data.list;
      });
      //表单获取状态数据
      listFyStatus(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
      ).then((response) => {
        this.statusOptionsForm = response.data.list;
      });
      //表单获取型号数据
      listFyType(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
      ).then((response) => {
        this.typeOptionsForm = response.data.list;
      });
      //表单获取负责人数据
      listFyUser(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
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
    cancel2() {
      this.open2 = false;
      this.reset();
    },

    // 表单重置
    reset() {
      this.form = {
        id: undefined,
        equipname: undefined,
        equipcode: undefined,
        username: undefined,
        area: undefined,
        type: undefined,
        status: undefined,
        faultcount: undefined,
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
      this.form.faultcount = "0";
      this.title = "添加设备台账";
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
      getFyEquipment(id).then((response) => {
        this.form = response.data;
        this.open = true;
        this.title = "修改设备台账";
        this.isEdit = true;
      });
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateFyEquipment(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addFyEquipment(this.form).then((response) => {
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
          return delFyEquipment({ ids: Ids });
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
    //打开对应图像表
    handleEditTable(row) {
      const tablename = row.equipname || this.ids[0];
      // console.log(tablename);
      this.$router.push({
        path: "/equipment/fy-data/",
        query: { equipname: tablename },
      });
    },
  },
};
</script>
