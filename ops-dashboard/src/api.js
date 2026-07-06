const RABBIT = import.meta.env.VITE_RABBIT_API ?? 'http://localhost:15672/api'
const CORE = import.meta.env.VITE_CORE_API ?? 'http://localhost:8081/api/v1'

export async function queues() {
  const res = await fetch(`${RABBIT}/queues`, { headers: { Authorization: 'Basic ' + btoa('guest:guest') } })
  return (await res.json()).map((q) => ({ name: q.name, depth: q.messages, consumers: q.consumers }))
}

export async function ipPools() {
  return (await fetch(`${CORE}/ip-pools`)).json()
}
