<template>
  <div class="space-y-6">
    <div class="bg-white shadow overflow-hidden sm:rounded-md">
      <div class="px-4 py-5 sm:px-6">
        <h3 class="text-lg leading-6 font-medium text-gray-900">Buckets</h3>
        <div class="mt-2 max-w-xl text-sm text-gray-500">
          <p>Manage buckets in the database.</p>
        </div>
        <div class="mt-5">
          <button
            @click="showCreateBucket = true"
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-green-600 hover:bg-green-700"
          >
            Create Bucket
          </button>
        </div>
        <ul class="mt-5 divide-y divide-gray-200">
          <li v-for="bucket in buckets" :key="bucket.name" class="px-4 py-4">
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <p class="text-sm font-medium text-gray-900">{{ bucket.name }}</p>
              </div>
              <div class="flex space-x-2">
                <button
                  @click="selectBucket(bucket.name)"
                  class="text-indigo-600 hover:text-indigo-900"
                >
                  View
                </button>
                <button
                  @click="deleteBucket(bucket.name)"
                  class="text-red-600 hover:text-red-900"
                >
                  Delete
                </button>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>

    <BucketViewer v-if="selectedBucket" :bucket="selectedBucket" :db-path="dbPath" />

    <CreateBucketModal v-if="showCreateBucket" @close="showCreateBucket = false" @created="loadBuckets" />
  </div>
</template>

<script>
import axios from 'axios'
import BucketViewer from './BucketViewer.vue'
import CreateBucketModal from './CreateBucketModal.vue'

export default {
  name: 'BucketManager',
  components: {
    BucketViewer,
    CreateBucketModal
  },
  props: {
    dbPath: String
  },
  data() {
    return {
      buckets: [],
      selectedBucket: null,
      showCreateBucket: false
    }
  },
  mounted() {
    this.loadBuckets()
  },
  methods: {
    async loadBuckets() {
      try {
        const response = await axios.get('/api/buckets')
        this.buckets = response.data
      } catch (error) {
        alert('Error loading buckets: ' + error.response.data.error)
      }
    },
    selectBucket(name) {
      this.selectedBucket = name
    },
    async deleteBucket(name) {
      if (confirm(`Delete bucket ${name}?`)) {
        try {
          await axios.delete(`/api/bucket/${name}`)
          this.loadBuckets()
          if (this.selectedBucket === name) {
            this.selectedBucket = null
          }
        } catch (error) {
          alert('Error deleting bucket: ' + error.response.data.error)
        }
      }
    }
  }
}
</script>