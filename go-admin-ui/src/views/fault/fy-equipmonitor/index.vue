
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
              allow-create
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in fyEquipmentFullList"
                :key="item.id"
                :label="item.equipname"
                :value="item.equipname"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="设备编号" prop="equipcode">
            <el-select
              v-model="queryParams.equipcode"
              placeholder="请输入设备编号"
              clearable
              filterable
              allow-create
              @keyup.enter.native="handleQuery"
              ><el-option
                v-for="item in fyEquipmentFullList"
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
              disabled
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
              >修改每页数据量
            </el-button>
          </el-col>

          <el-col :span="1.5">
            <el-badge :value="faultsumcount">
              <el-button
                v-permisaction="['admin:fyFault:remove']"
                type="danger"
                icon="el-icon-warning"
                size="mini"
                @click="handlefaultlist"
                >查看异常
              </el-button>
            </el-badge>
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
          style="width: 100%"
          v-loading="loading"
          :data="fyEquipmentList"
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
          />
          <el-table-column
            label="设备编码"
            align="center"
            prop="equipcode"
            :show-overflow-tooltip="true"
          />
          <!-- <el-table-column
            label="负责人"
            align="center"
            prop="username"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="地区"
            align="center"
            prop="area"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="型号"
            align="center"
            prop="type"
            :show-overflow-tooltip="true"
          /> -->
          <el-table-column
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
            min-width="200%"
          >
            <template slot-scope="scope">
              <el-button
                slot="reference"
                v-permisaction="['admin:fyEquipment:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
                >修改状态
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
                v-permisaction="['admin:fyEquipment:remove']"
                size="mini"
                type="text"
                icon="el-icon-tickets"
                @click="handlefault(scope.row)"
                >故障处理
              </el-button>

              <el-button
                slot="reference"
                v-permisaction="['admin:fyEquipment:edit']"
                size="mini"
                type="text"
                icon="el-icon-s-marketing"
                @click="handleEditTable(scope.row)"
                >图像
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
                disabled
                v-model="form.equipname"
                placeholder="设备名(不能重复)"
              />
            </el-form-item>
            <el-form-item label="设备编码" prop="equipcode">
              <el-input
                disabled
                v-model="form.equipcode"
                placeholder="设备编码"
              />
            </el-form-item>
            <el-form-item label="负责人" prop="username">
              <!-- <el-input v-model="form.username" placeholder="负责人" /> -->
              <el-select
                v-model="form.username"
                placeholder="请选择负责人"
                disabled
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
                disabled
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
                disabled
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

        <!-- 处理表单 -->
        <el-dialog :title="title" :visible.sync="openhandle" width="500px">
          <el-form
            ref="formhandle"
            :model="formhandle"
            :rules="handlerules"
            label-width="80px"
          >
            <el-form-item label="设备名" prop="equipname">
              <el-input
                v-model="formhandle.equipname"
                disabled
                placeholder="设备名"
              />
            </el-form-item>
            <el-form-item label="设备编号" prop="equipcode">
              <el-input
                v-model="formhandle.equipcode"
                disabled
                placeholder="设备编号"
              />
            </el-form-item>
            <el-form-item label="故障编号" prop="faultcode">
              <el-input
                v-model="formhandle.faultcode"
                disabled
                placeholder="故障编号"
              />
            </el-form-item>
            <!-- <el-form-item label="型号" prop="type">
              <el-input v-model="formhandle.type" placeholder="型号" />
            </el-form-item> -->
            <el-form-item label="处理方法" prop="handlemethod">
              <el-input
                v-model="formhandle.handlemethod"
                placeholder="处理方法"
              />
            </el-form-item>
            <el-form-item label="处理进展" prop="handlestatus">
              <!-- <el-input
                v-model="formhandle.handlestatus"
                placeholder="处理进展"
              /> -->
              <el-radio-group v-model="formhandle.handlestatus">
                <el-radio-button label="未解决"></el-radio-button>
                <el-radio-button label="已解决"></el-radio-button>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="处理时间" prop="handletime">
              <el-date-picker
                v-model="formhandle.handletime"
                type="datetime"
                placeholder="选择日期"
              >
              </el-date-picker>
            </el-form-item>
            <el-form-item label="备注" prop="remark">
              <el-input v-model="formhandle.remark" placeholder="备注" />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="ceshishuju">测试数据</el-button>
            <el-button type="primary" @click="submitFormHandle"
              >确 定</el-button
            >
            <el-button @click="cancelHandle">取 消</el-button>
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
import {
  addFyFault,
  delFyFault,
  getFyFault,
  listFyFault,
  updateFyFault,
} from "@/api/admin/fy-fault";
import { addFyHandle } from "@/api/admin/fy-handle";
import { listFyUser } from "@/api/admin/fy-user";
import { listFyArea } from "@/api/admin/fy-area";
import { listFyType } from "@/api/admin/fy-type";
import { listFyStatus } from "@/api/admin/fy-status";
import "@/api/equipment/status-color.css";

export default {
  name: "FyEquipmentmonitor",
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
      openhandle: false,
      isEdit: false,
      // 类型数据字典
      typeOptions: [],
      fyEquipmentList: [],
      fyEquipmentFullList: [],

      // 关系表类型
      areaOptionsForm: [],
      userOptionsForm: [],
      statusOptionsForm: [],
      typeOptionsForm: [],
      faultOptionsForm: [],
      handleOptionForm: [],

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
      // 查询复数参数
      queryMultiParams: {
        // pageIndex: 1,
        // pageSize: 10000,
        // equipname: undefined,
        // equipcode: undefined,
        // username: undefined,
        // area: undefined,
        // type: undefined,
        status: undefined,
        // faultcount: undefined,
        // remark: undefined,
      },
      // 表单参数
      form: {},
      form2: {},
      formhandle: {
        id: undefined,
        equipname: undefined,
        equipcode: undefined,
        faultcode: undefined,
        type: undefined,
        handlemethod: undefined,
        handlestatus: undefined,
        handletime: undefined,
        remark: undefined,
      },
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
      //提交表规则
      handlerules: {},
      //handletime
      handletime: undefined,
      //其他数据
      faultmessageId: undefined,
      faultmessageForm: {
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
      },
      faultsumcount: undefined,
    };
  },
  created() {
    this.getList();
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true;
      //表单获取故障表数据
      listFyFault(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
      ).then((response) => {
        this.faultOptionsForm = response.data.list;
      });
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
      //获取角标
      listFyEquipment(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
      ).then((response) => {
        // console.log(response.data.list);
        this.faultsumcount = 0;
        for (let index = 0; index < response.data.list.length; index++) {
          if (
            response.data.list[index].status === "存在风险" ||
            response.data.list[index].status === "发生故障" ||
            response.data.list[index].status === "设备维修"
          ) {
            // console.log(response.data.list[index].status);
            this.faultsumcount++;
          }
        }

        // this.fyEquipmentList = response.data.list;
        // this.faultsumcount = response.data.list.length;
        // console.log(this.faultsumcount);
      });
      //获取本表数据
      listFyEquipment(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.fyEquipmentList = response.data.list;
          this.total = response.data.count;
        }
      );

      //获取全列表设备数据
      listFyEquipment(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
      ).then((response) => {
        this.fyEquipmentFullList = response.data.list;

        // //复数查询状态返回数据
        // if (queryMultiParams.status !== undefined) {
        //   console.log("begin");
        //   let templist = [];
        //   for (
        //     let index = 0;
        //     index < this.fyEquipmentFullList.length;
        //     index++
        //   ) {
        //     if (
        //       this.queryMultiParams.status.includes(
        //         this.fyEquipmentFullList[index].status
        //       )
        //     ) {
        //       templist.push(this.fyEquipmentFullList[index]);
        //     }
        //     this.fyEquipmentList = templist;
        //     this.total = templist.length;
        //     this.queryParams.pageIndex = 1;
        //   }
        // }
      });
      this.loading = false;
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
    cancelHandle() {
      this.openhandle = false;
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
      this.formhandle = {
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
      this.resetForm("formhandle");
      this.faultmessageForm = {
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
      this.resetForm("faultmessageForm");
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
        this.title = "修改设备状态";
        this.isEdit = true;
      });
    },
    /** 处理故障按钮操作 */
    handlefault(row) {
      if (row.status !== "发生故障" && row.status !== "设备维修") {
        this.$message({
          message: "处理仅登记“发生故障”与“设备维修”设备",
          type: "warning",
        });
      } else {
        this.reset();
        const id = row.id || this.ids;
        getFyEquipment(id).then((response) => {
          this.form = response.data;
        });
        getFyEquipment(id).then((response) => {
          this.formhandle.equipname = response.data.equipname;
          this.formhandle.equipcode = response.data.equipcode;
          this.formhandle.faultcode = response.data.faultcount;
          this.formhandle.type = response.data.type;
          this.formhandle.handletime = new Date();
          this.openhandle = true;
          this.title = "添加处理台账";

          //处理故障台账对应故障编号id
          this.faultmessage();
          //得到对应id所属对象
          getFyFault(this.faultmessageId).then((response) => {
            this.faultmessageForm = response.data;
          });
        });
        this.isEdit = false;
      }
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
          const pageSize = parseInt(this.form2.pageSize);
          if (isNaN(pageSize) || pageSize <= 0) {
            this.queryParams.pageSize = 10; // 设置默认值
          } else {
            this.queryParams.pageSize = pageSize;
          }
          this.open2 = false;
          this.getList();
        }
      });
    },

    /** 故障处理提交按钮 */
    submitFormHandle: function () {
      //提交fault表,修改设备表处理状态
      if (this.formhandle.handlestatus == "已解决") {
        this.faultmessageForm.handlestatus = "已处理";
        this.form.status = "正常运行";
      } else {
        this.faultmessageForm.handlestatus = "处理中";
        this.form.status = "设备维修";
        // this.form.status = "设备维修";
      }
      //改变监控表设备状态
      updateFyEquipment(this.form).then((response) => {
        if (response.code === 200) {
          this.$message.success("已修改本表");
        } else {
          this.msgError(response.msg);
        }
      });

      // console.log(this.faultmessageForm);
      //提交fault表对应id行（修改）
      updateFyFault(this.faultmessageForm).then((response) => {
        if (response.code === 200) {
          this.$message.success("已修改故障表");
        } else {
          this.msgError(response.msg);
        }
      });

      //提交处理记录
      this.$refs["formhandle"].validate((valid) => {
        this.formhandle.faultcode = this.formhandle.faultcode.toString();
        if (valid) {
          addFyHandle(this.formhandle).then((response) => {
            if (response.code === 200) {
              this.$message.success("已增加处理记录");
              this.openhandle = false;
              this.getList();
            } else {
              this.msgError(response.msg);
            }
          });
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
      const equipname = row.equipname || this.ids[0];
      this.$store.dispatch("tagsView/delView", {
        path: "/fault/fy-monitordata",
      });
      this.$router.push({
        path: "/fault/fy-monitordata/",
        query: { equipname: equipname },
      });
    },
    //查看故障条目
    handlefaultlist() {
      this.queryMultiParams.status = ["发生故障", "存在风险", "设备维修"];
      this.queryParams.status = undefined;
      // this.getList();
      let templist = [];
      for (let index = 0; index < this.fyEquipmentFullList.length; index++) {
        // if (
        //   this.queryMultiParams.status.includes(
        //     this.fyEquipmentFullList[index].status
        //   )
        // )
        if (this.fyEquipmentFullList[index].status == "发生故障") {
          templist.push(this.fyEquipmentFullList[index]);
        }
      }
      for (let index = 0; index < this.fyEquipmentFullList.length; index++) {
        if (this.fyEquipmentFullList[index].status == "设备维修") {
          templist.push(this.fyEquipmentFullList[index]);
        }
      }
      for (let index = 0; index < this.fyEquipmentFullList.length; index++) {
        if (this.fyEquipmentFullList[index].status == "存在风险") {
          templist.push(this.fyEquipmentFullList[index]);
        }
      }

      this.fyEquipmentList = templist;
      this.total = templist.length;
      this.queryParams.pageIndex = 1;
    },
    //查看fault信息
    faultmessage() {
      //根据故障编号返回对应id
      let fcode = this.formhandle.faultcode;
      for (let index = 0; index < this.faultOptionsForm.length; index++) {
        if (this.faultOptionsForm[index].faultcode == fcode) {
          this.faultmessageId = this.faultOptionsForm[index].id;
        }
      }
    },
    //测试数据
    ceshishuju() {
      console.log("表单form", this.form);
      console.log("表单formhandle", this.formhandle);
      console.log("表单faultmassageForm", this.faultmessageForm);
    },
    // //表单按状态颜色显示和警告
    // tableRowClassStatus({ row }) {
    //   // console.log(row);
    //   if (row.status == "正常运行") {
    //     // console.log("正常运行");
    //     return "normal-row";
    //   } else if (row.status == "发生故障") {
    //     // console.log("发生故障");
    //     return "fault-row";
    //   } else if (row.status == "存在风险") {
    //     // console.log("存在风险");
    //     return "warning-row";
    //   } else if (row.status == "设备报废" || "设备检查") {
    //     // console.log("设备报废", "设备检查");
    //     return "notuse-row";
    //   } else if (row.status == "设备维修") {
    //     // console.log("设备维修");
    //     return "repair-row";
    //   }
    //   return "";
    // },
    //表单按状态颜色显示和警告
    tableCellClassStatus({ row, column, rowIndex, columnIndex }) {
      let cellStyle;
      switch (row.status) {
        case "正常运行":
          cellStyle = "color: rgb(0, 196, 0);font-weight: 700";
          break;
        case "发生故障":
          cellStyle = "background:rgba(255, 0, 0, 0.15);font-weight: 700";
          break;
        case "存在风险":
          cellStyle = "background:rgba(255, 166, 0, 0.15);font-weight: 700";
          break;
        case "设备报废":
        case "设备检查":
          cellStyle = "color:rgba(128, 128, 128, 0.5);font-weight: 700";
          break;
        case "设备维修":
          cellStyle = "background:rgba(255, 255, 0, 0.2);font-weight: 700";
          break;
        default:
          cellStyle = "";
      }
      if (column.label == "状态") {
        return cellStyle;
      }
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
}
</style>


