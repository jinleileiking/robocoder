import axios from 'axios';

const BASE_URL = '/api/v1';

const apiClient = axios.create({
  baseURL: BASE_URL,
});

export interface Exception {
  ID?: number;
  CreatedAt?: string;
  UpdatedAt?: string;
  DeletedAt?: string | null;
  room_name: string;
  device_name: string;
  exception_time: string;
  description?: string;
  handler?: string;
  handling_time?: string;
  handling_result?: string;
}

// 获取异常列表
export const getExceptions = async (): Promise<Exception[]> => {
  try {
    const response = await apiClient.get('/exceptions');
    return response.data;
  } catch (error) {
    console.error('Error fetching exceptions:', error);
    throw error;
  }
};

// 获取单个异常详情
export const getExceptionById = async (id: number): Promise<Exception> => {
  try {
    const response = await apiClient.get(`/exceptions/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching exception:', error);
    throw error;
  }
};

// 创建异常
export const createException = async (exception: Omit<Exception, 'ID' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt'>): Promise<Exception> => {
  try {
    const response = await apiClient.post('/exceptions', exception);
    return response.data;
  } catch (error) {
    console.error('Error creating exception:', error);
    throw error;
  }
};

// 更新异常
export const updateException = async (id: number, exception: Omit<Exception, 'ID' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt'>): Promise<Exception> => {
  try {
    const response = await apiClient.put(`/exceptions/${id}`, exception);
    return response.data;
  } catch (error) {
    console.error('Error updating exception:', error);
    throw error;
  }
};

// 删除异常
export const deleteException = async (id: number): Promise<void> => {
  try {
    await apiClient.delete(`/exceptions/${id}`);
  } catch (error) {
    console.error('Error deleting exception:', error);
    throw error;
  }
};