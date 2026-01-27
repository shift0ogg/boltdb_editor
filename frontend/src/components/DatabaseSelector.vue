<template>
  <div class="bg-white overflow-hidden shadow rounded-lg">
    <div class="px-4 py-5 sm:p-6">
      <h3 class="text-lg leading-6 font-medium text-gray-900">Open Database</h3>
      <div class="mt-2 max-w-xl text-sm text-gray-500">
        <p>Enter the path to your BoltDB database file.</p>
      </div>
      <form @submit.prevent="openDatabase" class="mt-5">
        <div class="flex">
          <input
            v-model="dbPath"
            type="text"
            placeholder="e.g., /path/to/database.db"
            class="flex-1 rounded-l-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
          />
          <button
            type="submit"
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-r-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Open
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'DatabaseSelector',
  data() {
    return {
      dbPath: ''
    }
  },
  methods: {
    async openDatabase() {
      try {
        await axios.post('/api/open', { path: this.dbPath })
        this.$emit('database-opened', this.dbPath)
        alert('Database opened successfully')
      } catch (error) {
        alert('Error opening database: ' + error.response.data.error)
      }
    }
  }
}
</script>