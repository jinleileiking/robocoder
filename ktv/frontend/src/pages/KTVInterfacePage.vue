<template>
  <div class="ktv-interface-container">
    <van-list v-model="loading" :finished="finished" @load="onLoad">
      <van-cell
        v-for="room in rooms"
        :key="room.id"
        :title="room.name"
        :label="room.status"
        is-link
        @click="goToRoomDetail(room.id)"
      />
    </van-list>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { List, Cell } from 'vant';
import { getRooms } from '@/api/ktv';

const router = useRouter();
const loading = ref(false);
const finished = ref(false);
const rooms = ref([]);

const onLoad = async () => {
  loading.value = true;
  try {
    const response = await getRooms();
    if (response.success) {
      rooms.value = response.data;
      finished.value = true;
    } else {
      // Handle error
    }
  } catch (error) {
    // Handle error
  } finally {
    loading.value = false;
  }
};

const goToRoomDetail = (roomId: number) => {
  router.push(`/ktv/room/${roomId}`);
};
</script>

<style scoped>
.ktv-interface-container {
  padding: 16px;
}
</style>