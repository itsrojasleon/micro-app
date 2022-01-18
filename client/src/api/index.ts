import axios from 'axios';

export const api = axios.create({
  // auth service
  baseURL: '/auth'
});

// For every single request try to get stored token
api.interceptors.request.use(async (config) => {
  try {
    const token = await localStorage.getItem('access_token');

    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`;
    }

    return config;
  } catch (err) {
    return Promise.reject(err);
  }
});
