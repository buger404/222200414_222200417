import axios from 'axios';

const instance = axios.create({
    baseURL: 'http://localhost:4523/m1/5195327-4861190-default',
    responseType: 'json',
    timeout: 10000,
});

instance.interceptors.request.use(
    (config) => {
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

instance.interceptors.response.use(
    (response) => {
        return response.data; // 直接返回数据
    },
    (error) => {
        return Promise.reject(error);
    }
);

export default instance;
