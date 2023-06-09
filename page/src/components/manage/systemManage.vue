<template>
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
</template>
<script>
import { sys } from '../../plugins/api'
export default {
    data: () => ({
        timer: 0,
        system: {},
        use: {},
    }),
    methods: {
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
        stop() {
            clearInterval(this.timer)
        }
    },
    mounted() {
        setTimeout(() => {
            this.getSysInfo()
        }, 500);
    }
};
</script>
<style scoped>
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
</style>