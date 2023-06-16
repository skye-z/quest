const api = {
    host: 'http://192.168.1.160:8080/opo-shenzhen',
    // host: 'http://192.168.1.2:20005/opo-shenzhen',
    get: (path) => {
        return new Promise((resolve, reject) => {
            $.ajax({
                url: api.host + path,
                method: 'GET',
                timeout: 5000,
                dataType: 'json',
                success: result => {
                    resolve(result)
                },
                error: err => {
                    reject(err)
                }
            })
        })
    },
    checkLogin: () => {
        return api.get('/v2/api/exp/check.ajax');
    }
}