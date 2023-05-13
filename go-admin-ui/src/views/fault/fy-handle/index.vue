
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
              allow-create
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in equipmentOptionsForm"
                :key="item.id"
                :label="item.equipname"
                :value="item.equipname"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="设备编号" prop="equipcode">
            <!-- <el-input
              v-model="queryParams.equipcode"
              placeholder="请输入设备编号"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <el-select
              v-model="queryParams.equipcode"
              placeholder="请输入设备编号"
              clearable
              filterable
              allow-create
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in equipmentOptionsForm"
                :key="item.id"
                :label="item.equipcode"
                :value="item.equipcode"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="故障编号" prop="faultcode">
            <!-- <el-input
              v-model="queryParams.faultcode"
              placeholder="请输入故障编号"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <el-select
              v-model="queryParams.faultcode"
              placeholder="请输入故障编号"
              clearable
              filterable
              allow-create
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in faultOptionsForm"
                :key="item.id"
                :label="item.faultcode"
                :value="item.faultcode"
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
              allow-create
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in typeOptionsForm"
                :key="item.id"
                :label="item.type"
                :value="item.type"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="处理方法" prop="handlemethod"
            ><el-input
              v-model="queryParams.handlemethod"
              placeholder="请输入处理方法"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item>
          <el-form-item label="处理结果" prop="handlestatus">
            <!-- <el-input
              v-model="queryParams.handlestatus"
              placeholder="请输入处理结果"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <div
              style="display: flex; justify-content: space-between; width: 100%"
            >
              <el-radio-group
                v-model="queryParams.handlestatus"
                @keyup.enter.native="handleQuery"
              >
                <el-radio-button label="未解决"></el-radio-button>
                <el-radio-button label="已处理"></el-radio-button>
              </el-radio-group>
              <el-button @click="cancelhandlestatus">重置</el-button>
            </div>
          </el-form-item>
          <el-form-item label="处理时间" prop="handletime"
            ><el-input
              v-model="queryParams.handletime"
              placeholder="请输入处理时间"
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
              v-permisaction="['admin:fyHandle:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              disabled
              @click="handleAdd"
              >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:fyHandle:edit']"
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
              v-permisaction="['admin:fyHandle:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
              >删除
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button size="mini" style="background: #e5e9f2"
              >共 {{ total }} 条</el-button
            >
          </el-col>
        </el-row>

        <el-table
          v-loading="loading"
          :data="fyHandleList"
          height="500"
          @selection-change="handleSelectionChange"
          :row-class-name="tableRowClassName"
          @row-click="handleRowClick"
        >
          <!-- v-clickOutSide="clickNoDom" -->
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="序号" align="center">
            <template slot-scope="scope">
              {{
                (queryParams.pageIndex - 1) * queryParams.pageSize +
                scope.$index +
                1
              }}
            </template>
          </el-table-column>
          <el-table-column
            label="设备名"
            align="center"
            prop="equipname"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="设备编号"
            align="center"
            prop="equipcode"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="故障编号"
            align="center"
            prop="faultcode"
            :show-overflow-tooltip="true"
          />
          <!-- <el-table-column label="故障编号" align="center" prop="faultcode">
            <template slot-scope="scope">
              <span :style="{ color: getFaultColor(scope.row.faultcode) }">{{
                scope.row.faultcode
              }}</span>
            </template>
          </el-table-column> -->

          <el-table-column
            label="型号"
            align="center"
            prop="type"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="处理方法"
            align="center"
            prop="handlemethod"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="处理结果"
            align="center"
            prop="handlestatus"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="处理时间"
            align="center"
            prop="handletime"
            :show-overflow-tooltip="true"
          >
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.handletime) }}</span>
            </template> </el-table-column
          ><el-table-column
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
                >修改处理记录
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
                disabled
                placeholder="设备名"
              />
            </el-form-item>
            <el-form-item label="设备编号" prop="equipcode">
              <el-input
                v-model="form.equipcode"
                disabled
                placeholder="设备编号"
              />
            </el-form-item>
            <el-form-item label="故障编号" prop="faultcode">
              <el-input
                v-model="form.faultcode"
                disabled
                placeholder="故障编号"
              />
            </el-form-item>
            <el-form-item label="型号" prop="type">
              <el-input v-model="form.type" disabled placeholder="型号" />
            </el-form-item>
            <el-form-item label="处理方法" prop="handlemethod">
              <el-input v-model="form.handlemethod" placeholder="处理方法" />
            </el-form-item>
            <el-form-item label="处理结果" prop="handlestatus">
              <!-- <el-input v-model="form.handlestatus" placeholder="处理结果" /> -->
              <div
                style="
                  display: flex;
                  justify-content: space-between;
                  width: 100%;
                "
              >
                <el-radio-group v-model="form.handlestatus">
                  <el-radio-button label="未解决"></el-radio-button>
                  <el-radio-button label="已解决"></el-radio-button>
                </el-radio-group>
                <el-button @click="cancelhandlestatusForm">重置 </el-button>
              </div>
            </el-form-item>
            <el-form-item label="处理时间" prop="handletime">
              <el-date-picker
                v-model="form.handletime"
                type="datetime"
                placeholder="选择日期"
              >
              </el-date-picker>
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
  addFyHandle,
  delFyHandle,
  getFyHandle,
  listFyHandle,
  updateFyHandle,
} from "@/api/admin/fy-handle";
import { listFyFault } from "@/api/admin/fy-fault";
import { listFyEquipment } from "@/api/admin/fy-equipment";
import { listFyType } from "@/api/admin/fy-type";
import "@/api/equipment/status-color.css";
import { convertHexToRgba } from "@/utils/costum";

export default {
  name: "FyHandle",
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
      fyHandleList: [],

      // 关系表类型
      typeOptionsForm: [],
      equipmentOptionsForm: [],
      faultOptionsForm: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        equipname: undefined,
        equipcode: undefined,
        faultcode: undefined,
        type: undefined,
        handlemethod: undefined,
        handlestatus: undefined,
        handletime: undefined,
        remark: undefined,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        equipname: [
          { required: true, message: "设备名不能为空", trigger: "blur" },
        ],
        equipcode: [
          { required: false, message: "设备编号不能为空", trigger: "blur" },
        ],
        faultcode: [
          { required: false, message: "故障编号不能为空", trigger: "blur" },
        ],
        type: [{ required: false, message: "型号不能为空", trigger: "blur" }],
        handlemethod: [
          { required: false, message: "处理方法不能为空", trigger: "blur" },
        ],
        handlestatus: [
          { required: false, message: "处理结果不能为空", trigger: "blur" },
        ],
        handletime: [
          { required: false, message: "处理时间不能为空", trigger: "blur" },
        ],
        remark: [{ required: false, message: "备注不能为空", trigger: "blur" }],
      },

      linkfaultcode: undefined, //故障编码跳转查询
      faultColors: {}, // 保存故障编号对应的颜色
      selectedRow: undefined, //改变选中背景实验
    };
  },
  created() {
    // console.log("实验重置是否created");
    this.setfaultcode();
    this.getList();
  },
  methods: {
    //接受传递错误编码参数
    setfaultcode() {
      // this.queryParams.faultcode = undefined;
      this.linkfaultcode = this.$route.query.queryfaultcode;
      if (this.linkfaultcode !== undefined) {
        this.queryParams.faultcode = this.linkfaultcode;
        // console.log("跳转信息", this.queryParams.faultcode);
      }
    },
    /** 查询参数列表 */
    getList() {
      this.loading = true;
      //表单获取型号数据
      listFyType(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
      ).then((response) => {
        this.typeOptionsForm = response.data.list;
      });
      //表单获取设备数据
      listFyEquipment(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
      ).then((response) => {
        this.equipmentOptionsForm = response.data.list;
      });
      //获取故障数据
      listFyFault(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
      ).then((response) => {
        this.faultOptionsForm = response.data.list;
      });
      //获取处理台账数据
      listFyHandle(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.fyHandleList = response.data.list;
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
        equipname: undefined,
        equipcode: undefined,
        faultcode: undefined,
        type: undefined,
        handlemethod: undefined,
        handlestatus: undefined,
        handletime: undefined,
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
    //删除cancelhandlestatus
    cancelhandlestatusForm() {
      this.form.handlestatus = undefined;
    },
    cancelhandlestatus() {
      this.queryParams.handlestatus = undefined;
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = [];
      this.resetForm("queryForm");
      this.queryParams.faultcode = undefined;
      this.handleQuery();
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset();
      this.open = true;
      this.title = "添加处理台账";
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
      getFyHandle(id).then((response) => {
        this.form = response.data;
        this.open = true;
        this.title = "修改处理台账";
        this.isEdit = true;
      });
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateFyHandle(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addFyHandle(this.form).then((response) => {
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
          return delFyHandle({ ids: Ids });
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
    //日期格式化
    dateFormatter(row, column) {
      let datetime = row.createdAt;
      if (datetime) {
        datetime = new Date(datetime);
        let y = datetime.getFullYear() + "-";
        let mon = datetime.getMonth() + 1 + "-";
        let d = datetime.getDate() + " ";
        let h = datetime.getHours() + ":";
        let min = datetime.getMinutes() + ":";
        let s = datetime.getSeconds();
        return y + mon + d + h + min + s;
      }
    },
    //根据faultcode随机颜色
    getFaultColor(faultcode) {
      if (!this.faultColors[faultcode]) {
        // 如果没有为故障编号设置颜色，生成一个新的随机颜色并保存
        this.faultColors[faultcode] = convertHexToRgba(
          `#${Math.floor(Math.random() * 16777216)
            .toString(16)
            .padStart(6, "0")}`,
          0.7
        );
      }

      return this.faultColors[faultcode];

      // return this.faultColors[faultcode];
    },

    // //表单按状态颜色显示和警告
    // tableRowClassStatus({ row }) {
    //   let color  = this.getFaultColor(row.faultcode);
    //   return {
    //     background: color ,
    //   };
    // },

    // //sleep函数
    // sleep(time) {
    //   var timeStamp = new Date().getTime();
    //   var endTime = timeStamp + time;
    //   while (true) {
    //     if (new Date().getTime() > endTime) {
    //       return;
    //     }
    //   }
    // },

    tableRowClassName({ row }) {
      if (this.selectedRow && row.id === this.selectedRow.id) {
        return "selected-row";
      } else if (
        this.selectedRow &&
        row.faultcode === this.selectedRow.faultcode
      ) {
        return "selected-row";
      } else {
        return "";
      }
    },
    handleRowClick(row) {
      if (this.selectedRow === row) {
        this.selectedRow = null;
      } else {
        this.selectedRow = row;
      }
    },
  },
};
</script>

<style>
.selected-row td {
  background-color: #efefef;
}

/* .selected-row:hover td {
  background-color: #297eff;
} */

/* .el-table .el-table__row:hover {
  background-color: #ff7373;
} */

/* .el-table .el-table__body-wrapper tbody tr:hover {
  background-color: #ff7373;
} */
</style>

