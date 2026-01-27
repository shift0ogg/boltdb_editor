<template>
  <div class="min-h-screen bg-gray-100">
    <header class="bg-white shadow">
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold text-gray-900">BoltDB Editor</h1>
      </div>
    </header>
    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div class="px-4 py-6 sm:px-0">
        <DatabaseSelector @database-opened="onDatabaseOpened" />
        <div v-if="currentDb" class="mt-8">
          <BucketManager :db-path="currentDb" />
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import DatabaseSelector from './components/DatabaseSelector.vue'
import BucketManager from './components/BucketManager.vue'

export default {
  name: 'App',
  components: {
    DatabaseSelector,
    BucketManager
  },
  data() {
    return {
      currentDb: null
    }
  },
  methods: {
    onDatabaseOpened(path) {
      this.currentDb = path
    }
  }
}
</script>