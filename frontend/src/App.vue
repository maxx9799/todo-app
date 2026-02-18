<script setup>
import { ref } from "vue"

const todos = ref([])
const text = ref("")

async function createTodo() {
  const res = await fetch("http://localhost:8080/health")
  const data = await res.text()
  todos.value.push(data)
}
</script>

<template>
  <div style="padding:40px">
    <h1>Todo App</h1>

    <button @click="createTodo">Call Backend</button>

    <ul>
      <li v-for="t in todos">{{ t }}</li>
    </ul>
  </div>
</template>
