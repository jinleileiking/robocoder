<template>
  <div class="device-detail-container">
    <div class="device-info">
      <van-cell-group>
        <van-cell title="设备名称" :value="device.name" />
        <van-cell title="种类" :value="device.type" />
        <van-cell title="健康状态" :value="device.healthStatus" />
        <van-cell title="设备型号" :value="device.model" />
        <van-cell title="开机时长" :value="device.uptime" />
      </van-cell-group>
    </div>
    <div class="device-stats">
      <van-progress :percentage="device.cpuUsage" color="blue" />
      <van-progress :percentage="device.memoryUsage" color="green" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { CellGroup, Cell, Progress } from 'vant';
import { getDeviceDetail } from '@/api/device';

interface Device {
  name: string;
  type: string;
  healthStatus: string;
  model: string;
  uptime: string;
  cpuUsage: number;
  memoryUsage: number;
}

const route = useRoute();
const deviceId = route.params.deviceId as string;
const device = ref<Device>({
  name: '',
  type: '',
  healthStatus: '',
  model: '',
  uptime: '',
  cpuUsage: 0,
  memoryUsage: 0,
});

const fetchDeviceDetail = async () => {
  try {
    const response = await getDeviceDetail(deviceId);
    device.value = response.data;
  } catch (error) {
    console.error('Failed to fetch device details:', error);
  }
};

onMounted(() => {
  fetchDeviceDetail();
});
</script>

<style scoped>
.device-detail-container {
  padding: 20px;
}
.device-info {
  margin-bottom: 20px;
}
.device-stats {
  margin-top: 20px;
}
</style>