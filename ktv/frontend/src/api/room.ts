import axios from 'axios';

const BASE_URL = '/api/';

const apiClient = axios.create({
  baseURL: BASE_URL,
});

// 获取房间列表
export const listRooms = async (): Promise<any> => {
  try {
    const response = await apiClient.get('v1/rooms');
    return response.data;
  } catch (error) {
    console.error('Error fetching rooms:', error);
    throw error;
  }
};

// 根据ID获取房间详情
export const getRoomById = async (id: number): Promise<any> => {
  try {
    const response = await apiClient.get(`v1/rooms/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching room detail:', error);
    throw error;
  }
};

// 创建房间
export const createRoom = async (roomData: { name: string; description: string; status: string; device_count: number; }): Promise<any> => {
  try {
    const response = await apiClient.post('v1/rooms', roomData);
    return response.data;
  } catch (error) {
    console.error('Error creating room:', error);
    throw error;
  }
};

// 更新房间
export const updateRoom = async (id: number, roomData: { name: string; description: string; status: string; device_count: number; }): Promise<any> => {
  try {
    const response = await apiClient.put(`v1/rooms/${id}`, roomData);
    return response.data;
  } catch (error) {
    console.error('Error updating room:', error);
    throw error;
  }
};

// 删除房间
export const deleteRoom = async (id: number): Promise<void> => {
  try {
    await apiClient.delete(`v1/rooms/${id}`);
  } catch (error) {
    console.error('Error deleting room:', error);
    throw error;
  }
};