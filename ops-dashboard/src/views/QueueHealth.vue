<script setup>
import { onMounted, onUnmounted, ref } from 'vue'
import { queues } from '../api'

const rows = ref([])
let timer
const load = async () => { rows.value = await queues() }
onMounted(() => { load(); timer = setInterval(load, 5000) })
onUnmounted(() => clearInterval(timer))
</script>

<template>
  <h1>RabbitMQ queues</h1>
  <table>
    <tr><th>Queue</th><th>Depth</th><th>Consumers</th></tr>
    <tr v-for="q in rows" :key="q.name" :class="{ warn: q.depth > 1000 || q.consumers === 0 }">
      <td>{{ q.name }}</td><td>{{ q.depth }}</td><td>{{ q.consumers }}</td>
    </tr>
  </table>
</template>
