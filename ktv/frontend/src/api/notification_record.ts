import axios from 'axios';

const BASE_URL = '/api/';

const apiClient = axios.create({
  baseURL: BASE_URL,
});

// 获取通知记录列表
export const getNotificationRecords = async (): Promise<any> => {
  try {
    const response = await apiClient.get('v1/notification_records');
    return response.data;
  } catch (error) {
    console.error('Error fetching notification records:', error);
    throw error;
  }
};

// 获取单个通知记录
export const getNotificationRecord = async (id: string): Promise<any> => {
  try {
    const response = await apiClient.get(`v1/notification_records/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching notification record:', error);
    throw error;
  }
};

// 创建通知记录
export const createNotificationRecord = async (data: {
  recipient: string;
  content_template: string;
  room_name: string;
  device_name: string;
  link?: string;
  sent_at: string;
  status: string;
}): Promise<any> => {
  try {
    const response = await apiClient.post('v1/notification_records', data);
    return response.data;
  } catch (error) {
    console.error('Error creating notification record:', error);
    throw error;
  }
};

// 更新通知记录
export const updateNotificationRecord = async (id: string, data: {
  recipient: string;
  content_template: string;
  room_name: string;
  device_name: string;
  link?: string;
  sent_at: string;
  status: string;
}): Promise<any> => {
  try {
    const response = await apiClient.put(`v1/notification_records/${id}`, data);
    return response.data;
  } catch (error) {
    console.error('Error updating notification record:', error);
    throw error;
  }
};

// 删除通知记录
export const deleteNotificationRecord = async (id: string): Promise<void> => {
  try {
    await apiClient.delete(`v1/notification_records/${id}`);
  } catch (error) {
    console.error('Error deleting notification record:', error);
    throw error;
  }
};