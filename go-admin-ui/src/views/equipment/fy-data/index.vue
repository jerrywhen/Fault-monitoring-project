
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <div id="fycharts"></div>
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
                v-for="item in equipmentOptionsForm"
                :key="item.id"
                :label="item.equipname"
                :value="item.equipname"
              />
            </el-select>
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
              v-permisaction="['admin:fyData:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
              >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:fyData:edit']"
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
              v-permisaction="['admin:fyData:remove']"
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
              v-permisaction="['admin:fyData:refresh']"
              type="danger"
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
          :data="fyDataList"
          height="250"
          @click="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="序号" type="index" align="center" />
          <!-- <el-table-column
            label="ID"
            align="center"
            prop="id"
            :show-overflow-tooltip="true"
          /> -->
          <el-table-column
            label="设备名"
            align="center"
            prop="equipname"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="x数据"
            align="center"
            prop="datax"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="y数据"
            align="center"
            prop="datay"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="创建时间"
            align="center"
            prop="createdAt"
            :formatter="dateFormatter"
            :show-overflow-tooltip="true"
          />
          <!-- <el-table-column
            -HH-MM-SS
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
                  v-permisaction="['admin:fyData:edit']"
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
                  v-permisaction="['admin:fyData:remove']"
                  size="mini"
                  type="text"
                  icon="el-icon-delete"
                  >删除
                </el-button>
              </el-popconfirm>
            </template>
          </el-table-column> -->
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
              <el-input v-model="form.equipname" placeholder="设备名" />
            </el-form-item>
            <el-form-item label="x数据" prop="datax">
              <el-input v-model="form.datax" placeholder="x数据" />
            </el-form-item>
            <el-form-item label="y数据" prop="datay">
              <el-input v-model="form.datay" placeholder="y数据" />
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
  addFyData,
  delFyData,
  getFyData,
  listFyData,
  updateFyData,
} from "@/api/admin/fy-data";
import { listFyEquipment } from "@/api/admin/fy-equipment";

import echarts from "echarts";

export default {
  name: "fydata",
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
      fyDataList: [],

      //新建数据
      fydata: {},
      data1: [],
      data2: [],
      count: 0,
      timetest: [],
      linkequipname: undefined,

      // 关系表类型
      equipmentOptionsForm: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        equipname: undefined,
      },
      queryParamstemp: {
        pageIndex: 1,
        pageSize: 100000,
        equipname: undefined,
      },
      // 表单参数
      form: {},
      form2: {},
      // 表单校验
      rules: {
        equipname: [
          { required: true, message: "设备名不能为空", trigger: "blur" },
        ],
      },
      //other
      IntervalTimer: undefined,
    };
  },
  beforeCreate() {},
  created() {
    this.setequipname();
    this.getList();
  },
  mounted() {
    var chartDom = document.getElementById("fycharts"); //初始化
    var myChart = echarts.init(chartDom);
    var option;

    option = {
      title: {
        text: "模拟数据",
      },
      tooltip: {
        trigger: "axis",
        formatter: function (params) {
          params = params[0];
          return "x=" + params.axisValue + " , y=" + params.data;
        },

        axisPointer: {
          animation: false, //是否开启动画
        },
      },
      dataZoom: [
        //加上dataZoom可实现缩放
        {
          type: "slider", //x轴
          xAxisIndex: 0,
        },
        {
          type: "inside", //x轴
          xAxisIndex: 0,
          zoomOnMouseWheel: true, //"alt", //如何触发缩放。可选值为：true：表示不按任何功能键，鼠标滚轮能触发缩放。false：表示鼠标滚轮不能触发缩放。'shift'：表示按住 shift 和鼠标滚轮能触发缩放。'ctrl'：表示按住 ctrl 和鼠标滚轮能触发缩放。'alt'：表示按住 alt 和鼠标滚轮能触发缩放。
        },
        // {
        //     type: 'slider',//y轴
        //     yAxisIndex: 0,
        // },
        // {
        //     type: 'inside',//y轴
        //     yAxisIndex: 0,
        // }
      ],

      //aplitLine：网格线
      xAxis: {
        // type: "time",
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
          name: "Fake Data",
          type: "line",
          showSymbol: false,
          data: [],
        },
      ],
    };
    option && myChart.setOption(option);

    this.IntervalTimer = setInterval(() => {
      // console.time("testtime");
      let fydata = [];
      let data1 = [];
      let data2 = [];
      let count = [];
      // console.log("进入循环");
      this.queryParamstemp.equipname = this.queryParams.equipname;
      listFyData(this.addDateRange(this.queryParamstemp, this.dateRange)).then(
        (response) => {
          // console.log("response.data", response.data);
          fydata = response.data.list;
          count = response.data.count;
          //   console.log("in getfydatalist", fydata);
          for (let i = 0; i < count; i++) {
            // if (this.queryParams.equipname === fydata[i].equipname) {
            data1.push(fydata[i].datax);
            data2.push(fydata[i].datay);
            // }
          }
          // let data1,data2=fydata.map((val)=>{
          //   return val.datax,val.datay
          // })
          // console.log("data1,data2",data1,data2);
          myChart.setOption({
            xAxis: {
              data: data1,
            },
            series: [
              {
                data: data2,
              },
            ],
          });
        }
      );
      //   this.getFyDataList();

      // console.log("结束循环");
      // console.timeEnd("testtime");
    }, 100);

    // console.log(a);
    // clearInterval(a)
  },
  beforeDestroy() {
    if (this.IntervalTimer) {
      // 如果定时器还在运行 或者直接关闭，不用判断
      clearInterval(this.IntervalTimer); //关闭
    }
  },

  methods: {
    // getFyDataList() {
    //   listFyData().then(
    //     (response) => {
    //       this.fydata = response.data.list;
    //       this.count = response.data.count;
    //       console.log("in getfydatalist",this.fydata);
    //     }
    //   );
    // },
    setequipname() {
      this.linkequipname = this.$route.query.equipname;
      if (this.linkequipname !== undefined) {
        this.queryParams.equipname = this.linkequipname;
        // console.log("this.$route.query.equipname", this.$route.query.equipname);
      }
    },
    /** 查询参数列表 */
    getList() {
      this.loading = true;
      //表单获取设备数据
      listFyEquipment(
        this.addDateRange({ pageIndex: 1, pageSize: 10000 }, this.dateRange)
      ).then((response) => {
        this.equipmentOptionsForm = response.data.list;
      });
      listFyData(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.fyDataList = response.data.list;
          this.total = response.data.count;
          this.loading = false;
        }
      );
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
        equipname: undefined,
        datax: undefined,
        datay: undefined,
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
      this.title = "添加FyData";
      this.isEdit = false;
    },
    /** 修改sizepage按钮操作 */
    handleopen2() {
      this.open2 = true;
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
      getFyData(id).then((response) => {
        this.form = response.data;
        this.open = true;
        this.title = "修改FyData";
        this.isEdit = true;
      });
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateFyData(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addFyData(this.form).then((response) => {
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
          return delFyData({ ids: Ids });
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
    //测试数据
    ceshishuju() {
      console.log("表单form", this.form);
    },
  },
};
</script>


<style>
#fycharts {
  width: 500px;
  height: 500px;
}
</style>