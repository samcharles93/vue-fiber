import axios from 'axios'
import { useMainStore } from '@/stores/main';

const api = axios.create({
    baseURL: "http://localhost:3000/api/v1"
})

api.interceptors.request.use(function (config) {
    const store = useMainStore()
    if (store.token) {
        config.headers.Authorization = `Bearer ${store.token}`;
    }
    return config
}, function (error) {
    return Promise.reject(error)
});

export default api
