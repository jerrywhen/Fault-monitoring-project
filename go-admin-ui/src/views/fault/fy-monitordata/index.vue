
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <div id="fycharts1" class="drawelcard"></div>
        <div id="fycharts2" class="drawelcard"></div>
        <div id="fycharts3" class="drawelcard"></div>
        <div id="fycharts4" class="drawelcard"></div>
      </el-card>
      <el-card class="box-card">
        <!-- // style="width: 600px;height:400px;" -->
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

          <!-- <el-form-item label="设备编号" prop="equipcode"
            ><el-input
              v-model="queryParams.equipcode"
              placeholder="请输入设备编号"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item> -->

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
              v-permisaction="['admin:fyMonitordata:add']"
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
              v-permisaction="['admin:fyMonitordata:edit']"
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
              v-permisaction="['admin:fyMonitordata:remove']"
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
          :data="fyMonitordataList"
          @selection-change="handleSelectionChange"
        >
          <el-table-column
            type="selection"
            width="55"
            align="center"
          /><el-table-column
            label="设备名"
            align="center"
            prop="equipname"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="设备编号"
            align="center"
            prop="equipcode"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="数据1"
            align="center"
            prop="record1"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="数据2"
            align="center"
            prop="record2"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="数据3"
            align="center"
            prop="record3"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="数据4"
            align="center"
            prop="record4"
            :show-overflow-tooltip="true"
          />
          <!-- <el-table-column
            label="数据5"
            align="center"
            prop="record5"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="数据6"
            align="center"
            prop="record6"
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
                v-permisaction="['admin:fyMonitordata:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
                >修改
              </el-button>
              <el-button
                slot="reference"
                v-permisaction="['admin:fyMonitordata:remove']"
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
                placeholder="设备名"
                disabled
              />
            </el-form-item>
            <el-form-item label="设备编号" prop="equipcode">
              <el-input
                v-model="form.equipcode"
                placeholder="设备编号"
                disabled
              />
            </el-form-item>
            <el-form-item label="数据1" prop="record1">
              <el-input v-model="form.record1" placeholder="数据1" />
            </el-form-item>
            <el-form-item label="数据2" prop="record2">
              <el-input v-model="form.record2" placeholder="数据2" />
            </el-form-item>
            <el-form-item label="数据3" prop="record3">
              <el-input v-model="form.record3" placeholder="数据3" />
            </el-form-item>
            <el-form-item label="数据4" prop="record4">
              <el-input v-model="form.record4" placeholder="数据4" />
            </el-form-item>
            <!-- <el-form-item label="数据5" prop="record5">
              <el-input v-model="form.record5" placeholder="数据5" />
            </el-form-item>
            <el-form-item label="数据6" prop="record6">
              <el-input v-model="form.record6" placeholder="数据6" />
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
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import {
  addFyMonitordata,
  delFyMonitordata,
  getFyMonitordata,
  listFyMonitordata,
  updateFyMonitordata,
} from "@/api/admin/fy-monitordata";
import { listFyEquipment } from "@/api/admin/fy-equipment";
import echarts from "echarts";

export default {
  name: "FyMonitordata",
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
      fyMonitordataList: [],

      // 关系表类型
      equipmentOptionsForm: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        equipname: undefined,
        equipcode: undefined,
        remark: undefined,
      },
      queryParamstemp: {
        pageIndex: 1,
        pageSize: 100000,
        equipname: undefined,
        equipcode: undefined,
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
          { required: true, message: "设备编号不能为空", trigger: "blur" },
        ],
      },
      //other
      IntervalTimer: undefined,
    };
  },
  created() {
    this.setfaultcode();
    this.getList();
  },
  mounted() {
    var myChart1 = echarts.init(document.getElementById("fycharts1")); //初始化
    var myChart2 = echarts.init(document.getElementById("fycharts2")); //初始化
    var myChart3 = echarts.init(document.getElementById("fycharts3")); //初始化
    var myChart4 = echarts.init(document.getElementById("fycharts4")); //初始化
    var option;
    option = {
      title: {
        text: "Record",
      },
      tooltip: {
        trigger: "axis",
        formatter: function (params) {
          params = params[0];
          return "时间：" + params.axisValue + "\n" + "记录：" + params.data;
        },
        axisPointer: {
          animation: false, //是否开启动画
        },
        extraCssText: "width:120px; white-space:pre-wrap",
      },
      dataZoom: [
        {
          type: "slider", //x轴
          xAxisIndex: 0,
        },
        {
          type: "inside", //x轴
          xAxisIndex: 0,
          zoomOnMouseWheel: true,
        },
      ],
      xAxis: {
        splitLine: {
          show: false,
        },
        data: [],
      },
      yAxis: {
        type: "value",
        boundaryGap: [0, "100%"],
        splitLine: {
          show: false,
        },
      },
      series: [
        {
          name: "记录1",
          type: "line",
          showSymbol: false,
          data: [],
        },
      ],
    };
    option && myChart1.setOption(option);
    option && myChart2.setOption(option);
    option && myChart3.setOption(option);
    option && myChart4.setOption(option);
    
    this.IntervalTimer = setInterval(() => {
      let fydata = [];
      let xdata = [];
      let record1 = [];
      let record2 = [];
      let record3 = [];
      let record4 = [];
      let count = [];
      this.queryParamstemp.equipname = this.queryParams.equipname;
      listFyMonitordata(
        this.addDateRange(this.queryParamstemp, this.dateRange)
      ).then((response) => {
        fydata = response.data.list;
        count = response.data.count;
        for (let i = 0; i < count; i++) {
          xdata.push(this.datexdraw(fydata[i].createdAt));
          record1.push(fydata[i].record1);
          record2.push(fydata[i].record2);
          record3.push(fydata[i].record3);
          record4.push(fydata[i].record4);
        }
        myChart1.setOption({
          title: { text: "Record1" },
          xAxis: { data: xdata },
          series: [{ data: record1 }],
        }); //动态数据chart1
        myChart2.setOption({
          title: { text: "Record2" },
          xAxis: { data: xdata },
          series: [{ data: record2 }],
        }); //动态数据chart2
        myChart3.setOption({
          title: { text: "Record3" },
          xAxis: { data: xdata },
          series: [{ data: record3 }],
        }); //动态数据chart3
        myChart4.setOption({
          title: { text: "Record4" },
          xAxis: { data: xdata },
          series: [{ data: record4 }],
        }); //动态数据chart4
      });
    }, 2000);
  },
  beforeDestroy() {
    if (this.IntervalTimer) {
      // 如果定时器还在运行 或者直接关闭，不用判断
      clearInterval(this.IntervalTimer); //关闭
    }
  },
  methods: {
    //接受传递设备名参数
    setfaultcode() {
      // this.queryParams.faultcode = undefined;
      this.linkequipname = this.$route.query.equipname;
      if (this.linkequipname !== undefined) {
        this.queryParams.equipname = this.linkequipname;
        // console.log("跳转信息", this.queryParams.faultcode);
      }
    },
    /** 查询参数列表 */
    getList() {
      this.loading = true;
      listFyMonitordata(
        this.addDateRange(this.queryParams, this.dateRange)
      ).then((response) => {
        this.fyMonitordataList = response.data.list;
        this.total = response.data.count;
      });
      listFyEquipment(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
      ).then((response) => {
        this.equipmentOptionsForm = response.data.list;
      });
      this.loading = false;
      console.log(this.equipmentOptionsForm);
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
        record1: undefined,
        record2: undefined,
        record3: undefined,
        record4: undefined,
        record5: undefined,
        record6: undefined,
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
      this.queryParams.equipname = undefined;
      this.handleQuery();
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset();
      this.open = true;
      this.title = "添加监控数据";
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
      getFyMonitordata(id).then((response) => {
        this.form = response.data;
        this.open = true;
        this.title = "修改监控数据";
        this.isEdit = true;
      });
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateFyMonitordata(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addFyMonitordata(this.form).then((response) => {
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
          return delFyMonitordata({ ids: Ids });
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
    //日期绘图需求
    datexdraw(datetime) {
      if (datetime) {
        datetime = new Date(datetime);
        let y = datetime.getFullYear();
        let mon = datetime.getMonth() + 1;
        let d = datetime.getDate();
        let h = datetime.getHours();
        let min = datetime.getMinutes();
        let s = datetime.getSeconds();
        return h + ":" + min + ":" + s;
      }
    },
  },
};
</script>
<style>
.drawelcard {
  width: 450px;
  height: 250px;
  display: inline-block;
}
#fycharts1 {
  background-color: rgba(255, 239, 219, 0.3);
}
#fycharts2 {
  background-color: rgba(255, 239, 219, 0.4);
}
#fycharts3 {
  background-color: rgba(255, 239, 219, 0.5);
}
#fycharts4 {
  background-color: rgba(255, 239, 219, 0.6);
}
</style>
