export const init = (_this) => {
    _this.app = {
        name: localStorage.getItem('app:config:name'),
        version: localStorage.getItem('app:config:version')
    }
    let user = localStorage.getItem('user:info')
    if (user) _this.user = JSON.parse(user)
    setTimeout(() => {
        _this.$refs.loading.hide()
    }, 500);
}

export const getSubjectList = (_this,api) => {
    api.getList().then(res => {
        if (res.list) {
            let list = []
            for (let i in res.list) {
                let item = res.list[i]
                list[i] = {
                    label: item.name,
                    value: item.id
                }
            }
            _this.subjects = list
        } else {
            window.$message.warning(res.message ? res.message : '发生意料之外的错误')
        }
    }).catch(() => {

    })
}