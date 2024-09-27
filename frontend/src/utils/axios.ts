import axios from 'axios';

const instance = axios.create({
    responseType: 'json',
    timeout: 10000
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
