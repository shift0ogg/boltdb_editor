<template>
  <div class="bg-white shadow overflow-hidden sm:rounded-md">
    <div class="px-4 py-5 sm:px-6">
      <h3 class="text-lg leading-6 font-medium text-gray-900">Bucket: {{ bucket }}</h3>
      <div class="mt-2 max-w-xl text-sm text-gray-500">
        <p>Manage key-value pairs in this bucket.</p>
      </div>
      <div class="mt-5">
        <button
          @click="showCreateKey = true"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-green-600 hover:bg-green-700"
        >
          Add Key-Value
        </button>
      </div>
      <ul class="mt-5 divide-y divide-gray-200">
        <li v-for="kv in keyValues" :key="kv.key" class="px-4 py-4">
          <div class="flex items-center justify-between">
            <div class="flex-1">
              <p class="text-sm font-medium text-gray-900">{{ kv.key }}</p>
              <p class="text-sm text-gray-500">{{ kv.value }}</p>
            </div>
            <div class="flex space-x-2">
              <button
                @click="editKeyValue(kv)"
                class="text-indigo-600 hover:text-indigo-900"
              >
                Edit
              </button>
              <button
                @click="deleteKey(kv.key)"
                class="text-red-600 hover:text-red-900"
              >
                Delete
              </button>
            </div>
          </div>
        </li>
      </ul>
    </div>

    <KeyValueModal
      v-if="showCreateKey || editingKeyValue"
      :bucket="bucket"
      :key-value="editingKeyValue"
      @close="closeModal"
      @saved="loadKeyValues"
    />
  </div>
</template>

<script>
import axios from 'axios'
import KeyValueModal from './KeyValueModal.vue'

export default {
  name: 'BucketViewer',
  components: {
    KeyValueModal
  },
  props: {
    bucket: String,
    dbPath: String
  },
  data() {
    return {
      keyValues: [],
      showCreateKey: false,
      editingKeyValue: null
    }
  },
  mounted() {
    this.loadKeyValues()
  },
  methods: {
    async loadKeyValues() {
      try {
        const response = await axios.get(`/api/bucket/${this.bucket}/keys`)
        this.keyValues = response.data
      } catch (error) {
        alert('Error loading key-values: ' + error.response.data.error)
      }
    },
    editKeyValue(kv) {
      this.editingKeyValue = kv
    },
    closeModal() {
      this.showCreateKey = false
      this.editingKeyValue = null
    },
    async deleteKey(key) {
      if (confirm(`Delete key ${key}?`)) {
        try {
          await axios.delete(`/api/bucket/${this.bucket}/key/${key}`)
          this.loadKeyValues()
        } catch (error) {
          alert('Error deleting key: ' + error.response.data.error)
        }
      }
    }
  }
}
</script>