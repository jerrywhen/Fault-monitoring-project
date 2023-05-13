<template>
  <div class="dashboard-editor-container">
    <el-row :gutter="12">
      <el-col
        :sm="24"
        :xs="24"
        :md="6"
        :xl="6"
        :lg="6"
        :style="{ marginBottom: '12px' }"
      >
        <chart-card flag="false" title="设备总数" :total="equipmentcount">
          <el-tooltip
            slot="action"
            class="item"
            effect="dark"
            content="在用设备包括异常设备"
            placement="top-start"
          >
            <i class="el-icon-warning-outline" />
          </el-tooltip>
          <div>
            <div
              style="display: inline-block; margin-right: 16px; font-size: 20px"
            >
              <div style="font-size: 16px; color: #333">
                在用设备
                <span style="color: green; font-weight: 700">{{
                  normalequip
                }}</span>
              </div>
            </div>

            <div
              style="display: inline-block; margin-right: 16px; font-size: 20px"
            >
              <div style="font-size: 16px; color: #333">
                停用设备
                <span style="color: gray; font-weight: 700">{{
                  abnormalequip
                }}</span>
              </div>
            </div>
          </div>
          <template slot="footer">设备状况</template>
        </chart-card>
      </el-col>
      <el-col
        :sm="24"
        :xs="24"
        :md="6"
        :xl="6"
        :lg="6"
        :style="{ marginBottom: '12px' }"
      >
        <chart-card
          flag="false"
          title="异常设备"
          :total="faultequip + warnequip"
        >
          <el-tooltip
            slot="action"
            class="item"
            effect="dark"
            content="异常状态设备数量"
            placement="top-start"
          >
            <i class="el-icon-warning-outline" />
          </el-tooltip>
          <div>
            <div
              style="display: inline-block; margin-right: 16px; font-size: 16px"
            >
              <div style="font-size: 16px; color: #333">
                故障设备
                <span style="color: rgba(255, 0, 0, 0.7); font-weight: 700">{{
                  faultequip
                }}</span>
              </div>
            </div>
            <div
              style="display: inline-block; margin-right: 16px; font-size: 16px"
            >
              <div style="font-size: 16px; color: #333">
                警告设备
                <span style="color: rgba(255, 166, 0, 0.7); font-weight: 700">{{
                  warnequip
                }}</span>
              </div>
            </div>
          </div>
          <template slot="footer">故障状况</template>
        </chart-card>
      </el-col>
      <el-col
        :sm="24"
        :xs="24"
        :md="6"
        :xl="6"
        :lg="6"
        :style="{ marginBottom: '12px' }"
      >
        <chart-card flag="false" title="故障总数" :total="faultcount">
          <el-tooltip
            slot="action"
            class="item"
            effect="dark"
            content="故障记录总览"
            placement="top-start"
          >
            <i class="el-icon-warning-outline" />
          </el-tooltip>
          <div>
            <div
              style="display: inline-block; margin-right: 16px; font-size: 16px"
            >
              <div style="font-size: 16px; color: #333">
                未处理
                <span style="color: rgba(255, 0, 0, 0.7); font-weight: 700">{{
                  untreatedcount
                }}</span>
              </div>
            </div>
            <div
              style="display: inline-block; margin-right: 16px; font-size: 16px"
            >
              <div style="font-size: 16px; color: #333">
                处理中
                <span style="color: rgba(255, 166, 0, 0.7); font-weight: 700">{{
                  treatingcount
                }}</span>
              </div>
            </div>
          </div>
          <template slot="footer">处理状况</template>
        </chart-card>
      </el-col>
      <el-col
        :sm="24"
        :xs="24"
        :md="6"
        :xl="6"
        :lg="6"
        :style="{ marginBottom: '12px' }"
      >
        <chart-card title="处理总数" :total="handlecount">
          <el-tooltip
            slot="action"
            class="item"
            effect="dark"
            content="处理台账数目"
            placement="top-start"
          >
            <i class="el-icon-warning-outline" />
          </el-tooltip>
          <div>
            <mini-progress
              color="rgb(19, 194, 194)"
              :target="100"
              :percentage="targetcount"
              height="8px"
            />
          </div>
          <template slot="footer">
            <!-- <trend flag="top" style="margin-right: 16px" rate="12">
              <span slot="term">设备4 </span>
            </trend> -->
            处理进度
          </template>
        </chart-card>
      </el-col>
    </el-row>

    <el-card :bordered="false" :body-style="{ padding: '0' }">
      <div class="salesCard">
        <el-tabs>
          <el-tab-pane label="设备信息">
            <el-row>
              <el-col :xl="16" :lg="12" :md="12" :sm="24" :xs="24">
                <bar :list="barData" title="故障数量" />
              </el-col>
              <el-col :xl="8" :lg="12" :md="12" :sm="24" :xs="24">
                <rank-list title="设备状况" :list="rankList" />
              </el-col>
            </el-row>
          </el-tab-pane>
          <!-- <el-tab-pane label="警报信息">
            <el-row>
              <el-col :xl="16" :lg="12" :md="12" :sm="24" :xs="24">
                <bar :list="barData2" title="警报数据" />
              </el-col>
              <el-col :xl="8" :lg="12" :md="12" :sm="24" :xs="24">
                <rank-list title="警报记录" :list="rankList" />
              </el-col>
            </el-row>
          </el-tab-pane> -->
        </el-tabs>
      </div>
    </el-card>
  </div>
</template>

<script>
import ChartCard from "@/components/ChartCard";
import Trend from "@/components/Trend";
import MiniArea from "@/components/MiniArea";
import MiniBar from "@/components/MiniBar";
import MiniProgress from "@/components/MiniProgress";
import RankList from "@/components/RankList/index";
import Bar from "@/components/Bar.vue";

import { listFyEquipment } from "@/api/admin/fy-equipment";
import { listFyFault } from "@/api/admin/fy-fault";
import { listFyHandle } from "@/api/admin/fy-handle";

export default {
  name: "DashboardAdmin",
  components: {
    ChartCard,
    Trend,
    MiniArea,
    MiniBar,
    MiniProgress,
    RankList,
    Bar,
  },
  data() {
    return {
      barData: [],
      barData2: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10000,
      },
      // 关系表类型
      equipmentOptionsForm: [],
      faultOptionsForm: [],
      handleOptionsForm: [],
      //相关表数量数据
      equipmentcount: undefined,
      faultcount: undefined,
      handlecount: undefined,
      //统计数据
      normalequip: 0,
      abnormalequip: 0,
      faultequip: 0,
      warnequip: 0,
      untreatedcount: 0,
      treatingcount: 0,
      targetcount: 0,
      //时间特殊处理
      createtlist: [],
      bardata1: [],
      outputtime: [],
      rankList: [],
    };
  },
  created() {
    this.getList();
    // setTimeout(() => {
    //   this.initdata();
    // }, 200);
  },
  methods: {
    getList() {
      this.loading = true;
      //获取相关数据
      listFyEquipment(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.equipmentOptionsForm = response.data.list;
          this.equipmentcount = response.data.count;

          this.abnormalequip = this.equipmentOptionsForm.reduce((acc, item) => {
            return item.status === "设备报废" ? acc + 1 : acc;
          }, 0);
          this.normalequip = response.data.count - this.abnormalequip;
          this.faultequip = this.equipmentOptionsForm.reduce((acc, item) => {
            return item.status === "发生故障" ? acc + 1 : acc;
          }, 0);
          this.warnequip = this.equipmentOptionsForm.reduce((acc, item) => {
            return item.status === "存在风险" ? acc + 1 : acc;
          }, 0);
        }
      );

      listFyFault(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.faultOptionsForm = response.data.list;
          this.faultcount = response.data.count;
          this.treatingcount = this.faultOptionsForm.reduce((acc, item) => {
            return item.handlestatus === "未处理" ? acc + 1 : acc;
          }, 0);
          this.untreatedcount = this.faultOptionsForm.reduce((acc, item) => {
            return item.handlestatus === "处理中" ? acc + 1 : acc;
          }, 0);

          if (this.treatingcount + this.untreatedcount == 0) {
            this.targetcount = 100;
          } else {
            this.targetcount =
              (this.treatingcount * 100) /
              (this.treatingcount + this.untreatedcount);
          }
          // console.log(this.targetcount * 100);

          //时间处理和排序
          this.createtlist = response.data.list.map((item) => item.faulttime);
          this.createtlist.sort((a, b) => new Date(a) - new Date(b));
          let handletime = this.createtlist;
          // console.log(this.createtlist);

          //筛选出最近的12个月的数据
          const twelveMonthsAgo = new Date();
          twelveMonthsAgo.setMonth(twelveMonthsAgo.getMonth() - 12);
          const filteredhandletime = handletime.filter((time) => {
            const temptime = new Date(time);
            const r = temptime.setMonth(temptime.getMonth()) > twelveMonthsAgo;
            if (r) {
              return time;
            }
          });
          // console.log(filteredhandletime);

          // //初始化对象x和y值
          // const nowtime = new Date().getMonth();
          // 初始化每个月份的数据为0
          const monthlyTotalArray = new Array(12).fill(0);
          // 计算每个月份的数据总数2023-01-01T00:00:00+08:00
          filteredhandletime.forEach((time) => {
            const nowtime = new Date();
            const nowyear = nowtime.getFullYear();
            const date = new Date(time);
            const month = date.getMonth();
            const beginmonth = nowtime.getMonth(); //去年起始月份

            if (date.getFullYear() < nowyear) {
              monthlyTotalArray[month - beginmonth]++;
            } else {
              monthlyTotalArray[month + 12 - beginmonth - 1]++;
            }
          });
          // console.log(monthlyTotalArray);
          this.outputtime = monthlyTotalArray;
          this.initdata();
        }
      );
      listFyHandle(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.handleOptionsForm = response.data.list;
          this.handlecount = response.data.count;
        }
      );
      this.loading = false;
    },
    initdata() {
      for (let i = 0; i < 12; i++) {
        var beginmonth = new Date().getMonth();
        var y = new Date().getFullYear() - 1;
        if (beginmonth + i > 10) {
          beginmonth = beginmonth - 12;
          y++;
        }
        this.barData.push({
          x: `${y}-${beginmonth + i + 2}`,
          y: this.outputtime[i],
        });
        this.rankList.push({
          name: `${y}-${beginmonth + i + 2}`,
          total: this.outputtime[i],
        });
        this.rankList.sort((a, b) => b.total - a.total);
      }
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
  },
};
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 12px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .github-corner {
    position: absolute;
    top: 0;
    border: 0;
    right: 0;
  }

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}

::v-deep .el-tabs__item {
  padding-left: 16px !important;
  height: 50px;
  line-height: 50px;
}

@media (max-width: 1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
