
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
                v-for="item in fyFaultFullList"
                :key="item.id"
                :label="item.faultcode"
                :value="item.faultcode"
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
              allow-create
              @keyup.enter.native="handleQuery"
            >
              <el-option
                v-for="item in areaOptionsForm"
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
              placeholder="请输入负责人"
              clearable
              filterable
              allow-create
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in userOptionsForm"
                :key="item.id"
                :label="item.username"
                :value="item.username"
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
              allow-create
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in statusOptionsForm"
                :key="item.id"
                :label="item.status"
                :value="item.status"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="处理进程" prop="handlestatus">
            <!-- <el-input
              v-model="queryParams.handlestatus"
              placeholder="请输入处理进程"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <div
              style="display: flex; justify-content: space-between; width: 100%"
            >
              <el-radio-group
                v-model="queryParams.handlestatus"
                :vertical="true"
                @keyup.enter.native="handleQuery"
              >
                <el-radio-button label="未处理"></el-radio-button>
                <el-radio-button label="处理中"></el-radio-button>
                <el-radio-button label="已处理"></el-radio-button>
              </el-radio-group>
              <el-button @click="cancelhandlestatus">重置</el-button>
            </div>
          </el-form-item>
          <el-form-item label="故障原因" prop="reason"
            ><el-input
              v-model="queryParams.reason"
              placeholder="请输入故障原因"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item>
          <el-form-item label="故障时间" prop="faulttime">
            <!-- <el-input
              v-model="queryParams.faulttime"
              placeholder="请输入故障时间"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            /> -->
            <el-select
              v-model="queryParams.faulttime"
              placeholder="请输入故障时间"
              clearable
              filterable
              allow-create
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in fyFaultFullList"
                :key="item.id"
                :label="item.faulttime"
                :value="item.faulttime"
              />
            </el-select>
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
              v-permisaction="['admin:fyFault:add']"
              type="danger"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
              >新增故障
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:fyFault:edit']"
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
              v-permisaction="['admin:fyFault:remove']"
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
              v-permisaction="['admin:fyFault:refresh']"
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
          <el-col :span="1.5">
            <el-button size="mini" style="background: #e5e9f2"
              >共 {{ total }} 条</el-button
            >
          </el-col>
        </el-row>

        <el-table
          height="500"
          stripe
          v-loading="loading"
          :data="fyFaultList"
          :cell-style="tableCellClassStatus"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="序号" type="index" align="center" />

          <el-table-column type="expand" label="详细">
            <template slot-scope="props">
              <el-form label-position="left" inline class="demo-table-expand">
                <!-- <el-form-item label="设备编码">
                  <span>{{ props.equipcode }}</span>
                </el-form-item> -->
                <el-form-item label="负责人：">
                  <span>{{ props.row.username }}</span>
                </el-form-item>
                <el-form-item label="地区：">
                  <span>{{ props.row.area }}</span>
                </el-form-item>
                <el-form-item label="型号：">
                  <span>{{ props.row.type }}</span>
                </el-form-item>
              </el-form>
            </template>
          </el-table-column>

          <el-table-column
            label="设备名"
            align="center"
            prop="equipname"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="设备编号"
            align="center"
            prop="equipcode"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="故障编号"
            align="center"
            prop="id"
            :show-overflow-tooltip="true"
          />
          <!-- <el-table-column
            label="地区"
            align="center"
            prop="area"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="负责人"
            align="center"
            prop="username"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="型号"
            align="center"
            prop="type"
            :show-overflow-tooltip="true"
          /> -->
          <!-- <el-table-column
            label="状态记录"
            align="center"
            prop="status"
            :show-overflow-tooltip="true"
          /> -->
          <el-table-column
            label="查看进程"
            align="center"
            prop="handlestatus"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="故障原因"
            align="center"
            prop="reason"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="故障时间"
            align="center"
            prop="faulttime"
            :show-overflow-tooltip="true"
          >
            <template slot-scope="scope">
              <!-- <span>{{ parseTime(scope.row.faulttime) }}</span> -->
              <span>{{ scope.row.faulttime }}</span>
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
                v-permisaction="['admin:fyFault:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
                >修改
              </el-button>

              <el-button
                slot="reference"
                v-permisaction="['admin:fyFault:remove']"
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
                icon="el-icon-chat-dot-round"
                @click="handleEditTable(scope.row)"
                >处理记录
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
              <!-- <el-input v-model="form.equipname" placeholder="设备名" /> -->
              <el-select
                v-model="form.equipname"
                placeholder="请选择设备名"
                clearable
                filterable
                :disabled="!editable"
                @change="equipmessage"
                ><el-option
                  v-for="item in equipmentOptionsForm"
                  :key="item.id"
                  :label="item.equipname"
                  :value="item.equipname"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="故障编号" prop="faultcode">
              <el-input
                disabled
                v-model="form.faultcode"
                placeholder="故障编号"
              />
            </el-form-item>
            <!-- <el-form-item label="状态" prop="status">
              <el-input v-model="form.status" placeholder="状态" />
            </el-form-item> -->
            <el-form-item label="处理进程" prop="handlestatus">
              <!-- <el-input v-model="form.handlestatus" placeholder="处理进程" /> -->
              <div
                style="
                  display: flex;
                  justify-content: space-between;
                  width: 100%;
                "
              >
                <el-radio-group v-model="form.handlestatus">
                  <el-radio-button label="未处理"></el-radio-button>
                  <el-radio-button label="处理中"></el-radio-button>
                  <el-radio-button label="已处理"></el-radio-button>
                </el-radio-group>
                <el-button @click="cancelhandlestatusForm">重置 </el-button>
              </div>
            </el-form-item>
            <el-form-item label="故障原因" prop="reason">
              <el-input v-model="form.reason" placeholder="故障原因" />
            </el-form-item>
            <el-form-item label="故障时间" prop="faulttime">
              <el-date-picker
                v-model="form.faulttime"
                type="datetime"
                placeholder="选择日期"
              >
                <!-- default-value="this.valueOpenForm" -->
              </el-date-picker>
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
  addFyFault,
  delFyFault,
  getFyFault,
  listFyFault,
  updateFyFault,
} from "@/api/admin/fy-fault";
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
import "@/api/equipment/status-color.css";
import { parseTime } from "@/utils";

export default {
  name: "FyFault",
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
      fyFaultList: [],
      fyFaultFullList: [],

      // 关系表类型
      areaOptionsForm: [],
      userOptionsForm: [],
      statusOptionsForm: [],
      typeOptionsForm: [],
      equipmentOptionsForm: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        equipname: undefined,
        equipcode: undefined,
        area: undefined,
        username: undefined,
        type: undefined,
        status: undefined,
        handlestatus: undefined,
        reason: undefined,
        faulttime: undefined,
        faultcode: undefined,
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
        area: [{ required: false, message: "地区不能为空", trigger: "blur" }],
        username: [
          { required: false, message: "负责人不能为空", trigger: "blur" },
        ],
        type: [{ required: false, message: "型号不能为空", trigger: "blur" }],
        status: [{ required: false, message: "状态不能为空", trigger: "blur" }],
        handlestatus: [
          { required: false, message: "处理进程不能为空", trigger: "blur" },
        ],
        reason: [
          { required: false, message: "故障原因不能为空", trigger: "blur" },
        ],
        faulttime: [
          { required: false, message: "故障时间不能为空", trigger: "blur" },
        ],
        faultcode: [
          { required: false, message: "故障编号不能为空", trigger: "blur" },
        ],
        remark: [{ required: false, message: "备注不能为空", trigger: "blur" }],
      },
      //其他数据
      equipmessageid: undefined,
      equipmessageform: {
        id: undefined,
        equipname: undefined,
        username: undefined,
        type: undefined,
        status: undefined,
        reason: undefined,
        remark: undefined,
      },
      valueOpenForm: undefined,
      newfaultid: undefined,
      //控制可编辑表名称
      editable: true,
    };
  },
  created() {
    this.getList();
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true;
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
        this.fyFaultFullList = response.data.list;
        this.newfaultid = response.data.list[0].id;
        // //格式化全表时间数据
        // for (let index = 0; index < response.data.count; index++) {
        //     this.fyFaultFullList[index].faulttime = this.dateListFormatter(
        //       response.data.list[index].faulttime
        //     );
        //   }
      });
      //获取故障数据
      listFyFault(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.fyFaultList = response.data.list;
          this.total = response.data.count;
          //   //格式化时间数据
          //   for (let index = 0; index < response.data.count; index++) {
          //     this.fyFaultList[index].faulttime = this.dateListFormatter(
          //       response.data.list[index].faulttime
          //     );
          //   }
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
        area: undefined,
        username: undefined,
        type: undefined,
        status: undefined,
        handlestatus: undefined,
        reason: undefined,
        faulttime: undefined,
        faultcode: undefined,
        remark: undefined,
      };
      this.resetForm("form");
      this.equipmessageform = {
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
      this.resetForm("equipmessageform");
      this.equipmessageid = undefined;
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
      this.editable = true; // 设置为设备名可编辑
      this.open = true;
      this.form.status = "发生故障";
      this.form.handlestatus = "未处理";
      this.form.faultcode = this.newfaultid + 1;
      this.form.faulttime = new Date();
      this.title = "添加故障台账";
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
      this.editable = false; // 设置为设备名不可编辑
      this.reset();
      const id = row.id || this.ids;
      this.equipmessage(row);
      getFyFault(id).then((response) => {
        this.form = response.data;
        this.open = true;
        this.form.faultcode = id;
        this.title = "修改故障台账";
        this.isEdit = true;
      });
    },
    /** 提交按钮 */
    submitForm: function () {
      //故障判断
      this.equipmessage();
      // console.log("this.equipmessageform.status:",this.equipmessageform.status,"this.form.id:",this.form.id);
      if (
        this.equipmessageform.status !== "正常运行" &&
        this.form.id == undefined
      ) {
        this.$message.warning("仅可添加未故障设备");
        return;
      }

      //故障编号改为字符串
      this.form.faultcode = this.form.faultcode.toString();
      //先提交equipment的status数据
      //   this.$refs["equipmessageform"].validate((valid) => {
      if (this.form.handlestatus === "未处理") {
        //故障编号转化为字符串,并存储故障编号
        this.equipmessageform.faultcount = this.form.faultcode;
        //数据处理
        this.equipmessageform.status = "发生故障";
        this.equipmessageform.createdAt = undefined;
        // this.equipmessageform.updatedAt = undefined;
        //提交设备表数据
        updateFyEquipment(this.equipmessageform).then((response) => {
          if (response.code === 200) {
            this.$message.success("成功更新设备表");
          } else {
            this.msgError(response.msg);
          }
        });
      }
      //   });

      this.$refs["form"].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateFyFault(this.form).then((response) => {
              if (response.code === 200) {
                this.$message.success("成功更新本表");
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addFyFault(this.form).then((response) => {
              if (response.code === 200) {
                this.$message.success("成功添加本表");
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
          return delFyFault({ ids: Ids });
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
    equipmessage(row) {
      //根据选择设备名返回对应id
      let name = this.form.equipname || row.equipname;
      for (let index = 0; index < this.equipmentOptionsForm.length; index++) {
        if (this.equipmentOptionsForm[index].equipname === name) {
          this.equipmessageid = this.equipmentOptionsForm[index].id;
        }
      }
      //根据对应id返回设编号等数据
      getFyEquipment(this.equipmessageid).then((response) => {
        this.equipmessageform = response.data;
        //插入数据
        this.form.equipcode = this.equipmessageform.equipcode;
        this.form.area = this.equipmessageform.area;
        this.form.username = this.equipmessageform.username;
        this.form.type = this.equipmessageform.type;
      });
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
    //日期格式化,改原数据
    dateListFormatter(datetime) {
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
    //删除cancelhandlestatus
    cancelhandlestatusForm() {
      this.form.handlestatus = undefined;
    },
    cancelhandlestatus() {
      this.queryParams.handlestatus = undefined;
    },
    //打开对应处理记录
    handleEditTable(row) {
      this.$store.dispatch("tagsView/delView", { path: "/fault/fy-handle" });
      this.$store.dispatch("tagsView/delView", { path: "/fault/fy-handle/" });
      // this.$store.dispatch("tagsView/delAllViews");

      // console.log("route............." ,this.$route.path,"type", typeof currentRoute);

      const tfaultcode = row.faultcode || this.ids[0];
      // console.log(tablename);
      this.$router.push({
        path: "/fault/fy-handle/",
        query: { queryfaultcode: tfaultcode },
      });
    },
    //查看进程渲染
    tableCellClassStatus({ row, column, rowIndex, columnIndex }) {
      let cellStyle;
      switch (row.handlestatus) {
        case "已处理":
          cellStyle = "color: rgb(0, 196, 0);font-weight: 700";
          break;
        case "未处理":
          cellStyle = "background:rgba(255, 0, 0, 0.15);font-weight: 700";
          break;
        case "处理中":
          cellStyle = "background:rgba(255, 255, 0, 0.2);font-weight: 700";
          break;
        default:
          cellStyle = "";
      }
      if (column.label == "查看进程") {
        return cellStyle;
      }
    },

    ceshishuju() {
      console.log("表单form", this.form);
      console.log("表单equipmessag", this.equipmessageform);
      console.log("queryParams", this.queryParams);
    },
  },
};
</script>

<style>
.demo-table-expand label {
  width: 90px;
  color: #255ba7;
}
.demo-table-expand .el-form-item {
  text-align: center;
  margin-right: 0;
  margin-bottom: 0;
  width: 33%;
  font-weight: 600;
}
</style>