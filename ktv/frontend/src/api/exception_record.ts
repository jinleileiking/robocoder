import axios from 'axios';

// 定义基础URL
const BASE_URL = '/api/';

// 创建axios实例
const apiClient = axios.create({
  baseURL: BASE_URL,
});

// 获取异常记录列表
export const listExceptionRecords = async (): Promise<any> => {
  try {
    const response = await apiClient.get('v1/exception-records');
    return response.data;
  } catch (error) {
    console.error('Error fetching exception records:', error);
    throw error;
  }
};

// 获取单个异常记录详情
export const getExceptionRecordById = async (id: string): Promise<any> => {
  try {
    const response = await apiClient.get(`v1/exception-records/` + id);
    return response.data;
  } catch (error) {
    console.error('Error fetching exception record detail:', error);
    throw error;
  }
};

// 创建新的异常记录
export const createExceptionRecord = async (exceptionRecord: any): Promise<any> => {
  try {
    const response = await apiClient.post('v1/exception-records', exceptionRecord);
    return response.data;
  } catch (error) {
    console.error('Error creating exception record:', error);
    throw error;
  }
};

// 更新异常记录
export const updateExceptionRecord = async (id: string, exceptionRecord: any): Promise<any> => {
  try {
    const response = await apiClient.put(`v1/exception-records/` + id, exceptionRecord);
    return response.data;
  } catch (error) {
    console.error('Error updating exception record:', error);
    throw error;
  }
};

// 删除异常记录
export const deleteExceptionRecord = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`v1/exception-records/` + id);
  } catch (error) {
    console.error('Error deleting exception record:', error);
    throw error;
  }
};