<script setup>
import { onMounted, ref } from 'vue'
import { ipPools } from '../api'

const pools = ref([])
onMounted(async () => { pools.value = await ipPools() })
</script>

<template>
  <h1>IP pool health</h1>
  <div v-for="p in pools" :key="p.id" class="pool">
    <h2>{{ p.name }}</h2>
    <p>{{ p.ips }} IPs · avg health {{ p.avgHealth }} · {{ p.unhealthy }} unhealthy</p>
  </div>
</template>
