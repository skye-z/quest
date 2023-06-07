import axios from 'axios'

const request = axios.create({
    baseURL: import.meta.env.VITE_APP_BASE_URL,
    timeout: 10000
})

request.interceptors.request.use(
    config => {
        let token = localStorage.getItem("user:access:token");
        if (token) config.headers['Authorization'] = 'Bearer '+token;
        return config
    },
    error => {
        return Promise.reject(error)
    }
)

// request.interceptors.response.use(
//     response => {
//         let code = parseInt(response.data.code);
//         if (code === 1000003 || code === 1000004) {
            
//         } else return response.data
//     }, () => {
//         throw "网络异常";
//     }
// )

export default request
