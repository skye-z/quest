<template>
    <div class="no-select">
        <loading ref="loading" />
        <head-bar :app="app" :user="user" />
        <div class="app-body">
            <div class="sys-box flex justify-between">
                <div class="sys-item card mr-10 mb-10 full-width">
                    <div class="flex justify-between">
                        <div class="sys-name">CPU负载</div>
                        <div class="text-gray">{{ system.cpuPhysical }}核{{ system.cpuLogical }}线程</div>
                    </div>
                    <div class="sys-sub flex justify-between">
                        <div v-if="use.avg">
                            <div class="flex">
                                <div class="cpu-label">1分钟</div>
                                <div>{{ use.avg.load1 }}%</div>
                            </div>
                            <div class="flex">
                                <div class="cpu-label">5分钟</div>
                                <div>{{ use.avg.load5 }}%</div>
                            </div>
                            <div class="flex">
                                <div class="cpu-label">15分钟</div>
                                <div>{{ use.avg.load15 }}%</div>
                            </div>
                        </div>
                        <div v-else>N/A</div>
                        <div class="sys-value" :class="{ level2: use.level == 2, level3: use.level == 3 }">{{ use.cpu }}%
                        </div>
                    </div>
                </div>
                <div class="sys-item card mr-10 mb-10 full-width">
                    <div class="flex justify-between">
                        <div class="sys-name">内存占用</div>
                        <div class="text-gray">{{ getSize(system.memory) }}</div>
                    </div>
                    <div class="sys-sub flex justify-between" v-if="use.memory">
                        <div class="flex">
                            <div class="mr-5">可用</div>
                            <div>{{ getSize(use.memory.free) }}</div>
                        </div>
                        <div class="sys-value" :class="{ level2: use.memory.level == 2, level3: use.memory.level == 3 }">{{
                            use.memory.percent }}%</div>
                    </div>
                </div>
                <div class="sys-item card mr-10 mb-10 full-width">
                    <div class="flex justify-between">
                        <div class="sys-name">硬盘使用</div>
                        <div class="text-gray">{{ getSize(system.disk) }}</div>
                    </div>
                    <div class="sys-sub flex justify-between" v-if="use.disk">
                        <div class="flex">
                            <div class="mr-5">可用</div>
                            <div>{{ getSize(use.disk.free) }}</div>
                        </div>
                        <div class="sys-value" :class="{ level2: use.disk.level == 2, level3: use.disk.level == 3 }">{{
                            use.disk.percent }}%</div>
                    </div>
                </div>
                <div class="sys-item card mb-10 full-width">
                    <div class="flex justify-between">
                        <div class="sys-name">操作系统</div>
                        <div class="text-gray">{{ system.arch }}</div>
                    </div>
                    <div>
                        <div>系统名称: {{ system.type }}</div>
                        <div>系统版本: {{ system.version }}</div>
                        <div>在线时间: {{ distance(system.bootTime) }}</div>
                    </div>
                </div>
            </div>
            <div class="admin-box flex">
                <div class="admin-left full-width">
                    <div class="card mb-10">
                        <n-button class="float-right mt-5 mr-5" size="small" type="primary" round>
                            <template #icon>
                                <n-icon>
                                    <AddCircle24Regular />
                                </n-icon>
                            </template>
                            添加
                        </n-button>
                        <div class="card-name">科目 <span class="text-gray">{{ subjects.length }}</span></div>
                        <n-scrollbar style="height: 300px">
                            <div class="subject-item flex align-center justify-between" v-for="item in subjects"
                                @click="editSubject(item)">
                                <div>
                                    <div class="subject-name line1">{{ item.name }}</div>
                                    <div v-if="item.tag.length > 0">
                                        <n-tag class="mr-5" v-for="tag in item.tags" size="small" type="success"
                                            :bordered="false">{{ tag }}</n-tag>
                                    </div>
                                </div>
                                <n-button @click.stop="removeSubject(item.id, item.name)" strong secondary circle
                                    type="error">
                                    <template #icon>
                                        <n-icon>
                                            <DeleteForeverRound />
                                        </n-icon>
                                    </template>
                                </n-button>
                            </div>
                        </n-scrollbar>
                    </div>
                    <div class="card">
                        <n-button class="float-right mt-5 mr-5" size="small" type="primary" round>
                            <template #icon>
                                <n-icon>
                                    <AddCircle24Regular />
                                </n-icon>
                            </template>
                            添加
                        </n-button>
                        <div class="card-name">用户</div>
                        <div class="user-list"></div>
                    </div>
                </div>
                <div class="admin-right full-width">
                    <div class="card mb-10">
                        <n-button class="float-right mt-5 mr-5" size="small" type="primary" round>
                            <template #icon>
                                <n-icon>
                                    <AddCircle24Regular />
                                </n-icon>
                            </template>
                            添加
                        </n-button>
                        <div class="card-name">题目</div>
                        <div class="question-list"></div>
                    </div>
                    <div class="card">
                        <div class="card-name">考试</div>
                        <div class="exam-list"></div>
                    </div>
                </div>
            </div>
        </div>
        <foot-bar :app="app" />
    </div>
</template>
  
<script>
import Loading from "../components/loading.vue";
import HeadBar from "../components/headBar.vue";
import FootBar from "../components/footBar.vue";
import { DeleteForeverRound } from '@vicons/material'
import { Warning24Filled, AddCircle24Regular } from '@vicons/fluent'
import { sys, subject } from '../plugins/api'
import { init } from "../plugins/common"

export default {
    name: "Admin",
    components: { Loading, HeadBar, FootBar, Warning24Filled, DeleteForeverRound, AddCircle24Regular },
    data: () => ({
        timer: 0,
        system: {},
        use: {},
        app: {
            name: 'Quest云题库',
            version: '1.0.0'
        },
        user: {},
        subjects: []
    }),
    methods: {
        jump(name) {
            this.$router.push('/' + name)
        },
        getSysInfo() {
            sys.getInfo().then(res => {
                if (res.id) {
                    this.system = res
                    this.getSysUse()
                    this.timer = setInterval(() => this.getSysUse(), 5000)
                } else {
                    window.$message.warning(res.message ? res.message : '发生意料之外的错误')
                }
            }).catch(() => {

            })
        },
        getSysUse() {
            sys.getUse().then(res => {
                if (res) {
                    if (res.avg) {
                        res.avg.load1 = res.avg.load1.toFixed(2)
                        res.avg.load5 = res.avg.load5.toFixed(2)
                        res.avg.load15 = res.avg.load15.toFixed(2)
                    }
                    if (res.cpu) {
                        let all = 0.0
                        for (let i in res.cpu) all += res.cpu[i]
                        all = (all / res.cpu.length)
                        let level = 1;
                        if (all > 90) level = 3
                        else if (all > 65) level = 2
                        res.cpu = all.toFixed(2)
                        res.level = level
                    }
                    if (res.memory) {
                        let level = 1;
                        if (res.memory.usedPercent > 90) level = 3
                        else if (res.memory.usedPercent > 65) level = 2
                        res.memory = {
                            free: res.memory.available / 1024 / 1024,
                            percent: res.memory.usedPercent.toFixed(2),
                            level
                        }
                    }
                    if (res.disk) {
                        let level = 1;
                        if (res.disk.usedPercent > 90) level = 3
                        else if (res.disk.usedPercent > 65) level = 2
                        res.disk = {
                            free: res.disk.free / 1024 / 1024,
                            percent: res.disk.usedPercent.toFixed(2),
                            level
                        }
                    }
                    this.use = res
                } else {
                    window.$message.warning(res.message ? res.message : '发生意料之外的错误')
                }
            }).catch(err => {
                console.log(err)
            })
        },
        getSize(num) {
            if (num > (1024 * 1024)) return parseFloat(num / 1024 / 1024).toFixed(0) + 'TB'
            else if (num > 1024) return parseFloat(num / 1024).toFixed(0) + 'GB'
            else return num + 'MB'
        },
        getSubjectList() {
            subject.getList().then(res => {
                if (res.list) {
                    let list = [];
                    for (let i in res.list) {
                        let item = res.list[i]
                        let tags = [];
                        if (item.tag.indexOf(',') > -1) tags = item.tag.split(',')
                        else if (item.tag.indexOf('，') > -1) tags = item.tag.split(',')
                        else if (item.tag.indexOf(' ') > -1) tags = item.tag.split(' ')
                        else if (item.tag) tags.push(item.tag)
                        item.tags = tags
                        list.push(item)
                    }
                    this.subjects = res.list
                } else {
                    window.$message.warning(res.message ? res.message : '发生意料之外的错误')
                }
            }).catch(() => {

            })
        },
        editSubject(item) {
            console.log(item)
        },
        removeSubject(id, name) {
            window.$dialog.warning({
                title: '操作确认',
                content: '确认要删除《' + name + '》科目?',
                positiveText: '确认',
                negativeText: '取消',
                onPositiveClick: () => {

                }
            })
        },
        distance(timestamp) {
            if (timestamp === undefined || timestamp == null)
                return "无数据";
            timestamp = timestamp * 1000;

            let minute = 1000 * 60;
            let hour = minute * 60;
            let day = hour * 24;
            let now = new Date().getTime();
            let diffValue = now - timestamp;

            if (diffValue < 0) {
                return '不久前';
            }

            let dayC = diffValue / day;
            let hourC = diffValue / hour;
            let minC = diffValue / minute;

            if (dayC >= 1) {
                return parseInt(dayC) + "天";
            } else if (hourC >= 1) {
                return parseInt(hourC) + "小时";
            } else if (minC >= 1) {
                return parseInt(minC) + "分钟";
            }
            return '刚刚';
        },

    },
    mounted() {
        init(this)
        this.getSysInfo()
        this.getSubjectList()
    },
    beforeRouteLeave() {
        console.log('leave')
        clearInterval(this.timer)
    }
};
</script>
  
<style scoped>
.tips {
    padding-top: 30vh;
    font-size: 24px;
    width: 100%;
}

.card {
    background-color: #3c4a5f;
    border-radius: 8px;
}

.card-name {
    border-bottom: #556276 1px solid;
    padding: 10px;
}

.sys-item {
    padding: 10px 10px 8px 10px;
    width: calc(100% - 20px);
    position: relative;
}

.sys-sub {
    align-items: flex-end;
    height: 67px;
}

.sys-name {
    font-weight: 700;
    font-size: 16px;
}

.sys-value {
    font-size: 32px;
    line-height: 36px;
}

.sys-value.level2 {
    color: #f2c97d;
}

.sys-value.level3 {
    color: #e88080;
}

.cpu-label {
    width: 50px;
}

.mem-value {
    position: absolute;
    bottom: 10px;
    right: 10px;
}

.admin-left {
    margin-right: 10px;
    max-width: 292px;
    min-width: 292px;
}

.subject-list {
    height: 300px;
}

.subject-item {
    border-top: #556276 1px solid;
    cursor: pointer;
    padding: 10px;
}

.subject-item:hover {
    background-color: #45546a;
}

.user-list {
    height: 300px;
}

.question-list {
    height: 300px;
}

.exam-list {
    height: 300px;
}

.admin-right {}

@media (max-width:970px) {
    .sys-box {
        justify-content: space-between;
        align-items: center;
        flex-wrap: wrap;
        display: flex;
    }

    .sys-item {
        width: calc(50% - 25px);
        margin-right: 0;
    }

}

@media (max-width:500px) {

    .sys-item {
        width: calc(100% - 20px);
        margin-right: 0;
    }

    .sys-box {
        display: block;
    }

}

@media (max-width:670px) {
    .admin-left {
        max-width: none;
    }

    .admin-box {
        display: block;
    }

    .user-list {
        margin-bottom: 10px;
    }
}
</style>
  
