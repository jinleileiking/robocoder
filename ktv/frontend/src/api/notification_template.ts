import axios from 'axios';

const BASE_URL = '/api/';

const apiClient = axios.create({
  baseURL: BASE_URL,
});

// 获取通知模板列表
export const getNotificationTemplates = async (): Promise<any> => {
  try {
    const response = await apiClient.get('v1/notification_templates');
    return response.data;
  } catch (error) {
    console.error('Error fetching notification templates:', error);
    throw error;
  }
};

// 根据ID获取通知模板详情
export const getNotificationTemplateById = async (id: number): Promise<any> => {
  try {
    const response = await apiClient.get(`v1/notification_templates/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching notification template by ID:', error);
    throw error;
  }
};

// 创建新的通知模板
export const createNotificationTemplate = async (template: { name: string; content: string; receiver_id: number }): Promise<any> => {
  try {
    const response = await apiClient.post('v1/notification_templates', template);
    return response.data;
  } catch (error) {
    console.error('Error creating notification template:', error);
    throw error;
  }
};

// 更新通知模板
export const updateNotificationTemplate = async (id: number, template: { name: string; content: string; receiver_id: number }): Promise<any> => {
  try {
    const response = await apiClient.put(`v1/notification_templates/${id}`, template);
    return response.data;
  } catch (error) {
    console.error('Error updating notification template:', error);
    throw error;
  }
};

// 删除通知模板
export const deleteNotificationTemplate = async (id: number): Promise<void> => {
  try {
    await apiClient.delete(`v1/notification_templates/${id}`);
  } catch (error) {
    console.error('Error deleting notification template:', error);
    throw error;
  }
};