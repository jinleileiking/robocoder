import axios from 'axios';

// 定义基础URL
const BASE_URL = '/api/';

// 创建axios实例
const apiClient = axios.create({
  baseURL: BASE_URL,
});

// 获取设备列表
export const getDevices = async (): Promise<any> => {
  try {
    const response = await apiClient.get('v1/devices');
    return response.data;
  } catch (error) {
    console.error('Error fetching devices:', error);
    throw error;
  }
};

// 获取单个设备详情
export const getDeviceById = async (id: string): Promise<any> => {
  try {
    const response = await apiClient.get(`v1/devices/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching device detail:', error);
    throw error;
  }
};

// 创建设备
export const createDevice = async (device: any): Promise<any> => {
  try {
    const response = await apiClient.post('v1/devices', device);
    return response.data;
  } catch (error) {
    console.error('Error creating device:', error);
    throw error;
  }
};

// 更新设备
export const updateDevice = async (id: string, device: any): Promise<any> => {
  try {
    const response = await apiClient.put(`v1/devices/${id}`, device);
    return response.data;
  } catch (error) {
    console.error('Error updating device:', error);
    throw error;
  }
};

// 删除设备
export const deleteDevice = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`v1/devices/${id}`);
  } catch (error) {
    console.error('Error deleting device:', error);
    throw error;
  }
};