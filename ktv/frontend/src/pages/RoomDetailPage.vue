<template>
  <div class="room-detail-container">
    <h1 class="text-lg font-bold mb-4">包房详情</h1>
    <van-list>
      <van-cell
        v-for="device in devices"
        :key="device.id"
        :title="device.name"
        :label="device.type"
        is-link
        @click="goToDeviceDetail(device.id)"
      >
        <template #right-icon>
          <van-tag :type="device.healthStatus === 'good' ? 'success' : 'danger'">
            {{ device.healthStatus === 'good' ? '健康' : '故障' }}
          </van-tag>
        </template>
      </van-cell>
    </van-list>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { List, Cell, Tag } from 'vant';
import { getRoomDevices } from '@/api/rooms';

const route = useRoute();
const router = useRouter();
const roomId = route.params.roomId as string;

const devices = ref<{ id: string; name: string; type: string; healthStatus: string }[]>([]);

const fetchDevices = async () => {
  try {
    const response = await getRoomDevices(roomId);
    devices.value = response.data;
  } catch (error) {
    console.error('获取设备列表失败：', error);
  }
};

const goToDeviceDetail = (deviceId: string) => {
  router.push(`/devices/${deviceId}`);
};

onMounted(() => {
  fetchDevices();
});
</script>

<style scoped>
.room-detail-container {
  padding: 16px;
}
</style>