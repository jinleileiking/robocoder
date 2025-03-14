<template>
  <div class="room-list-container">
    <van-list v-model:loading="loading" :finished="finished" finished-text="没有更多了" @load="onLoad">
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
import { getRooms } from '@/api/room';

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
      if (response.data.length < 10) {
        finished.value = true;
      }
    } else {
      finished.value = true;
    }
  } catch (error) {
    finished.value = true;
  } finally {
    loading.value = false;
  }
};

const goToRoomDetail = (roomId: number) => {
  router.push({ path: `/rooms/${roomId}` });
};
</script>

<style scoped>
.room-list-container {
  padding: 16px;
}
</style>