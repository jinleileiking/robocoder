// src/router/index.ts
import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import RoomListPage from '@/pages/RoomListPage.vue';
import RoomDetailPage from '@/pages/RoomDetailPage.vue';
import DeviceDetailPage from '@/pages/DeviceDetailPage.vue';
import KTVInterfacePage from '@/pages/KTVInterfacePage.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/rooms',
    name: 'RoomListPage',
    component: RoomListPage,
  },
  {
    path: '/rooms/:roomId',
    name: 'RoomDetailPage',
    component: RoomDetailPage,
  },
  {
    path: '/devices/:deviceId',
    name: 'DeviceDetailPage',
    component: DeviceDetailPage,
  },
  {
    path: '/ktv',
    name: 'KTVInterfacePage',
    component: KTVInterfacePage,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;