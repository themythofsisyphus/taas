<template>
  <nav class="bg-white shadow-sm border-b">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <h1 class="text-xl font-bold text-gray-900">TaaS Dashboard</h1>
          </div>
          <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
            <button
              v-for="tab in tabs"
              :key="tab.id"
              @click="$emit('tab-change', tab.id)"
              :class="[
                'inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium',
                activeTab === tab.id
                  ? 'border-blue-500 text-gray-900'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
              ]"
            >
              {{ tab.name }}
            </button>
          </div>
        </div>
        <div class="flex items-center">
          <div class="flex items-center space-x-4">
            <span class="text-sm text-gray-500">Tenant ID: {{ currentTenant }}</span>
            <div class="flex items-center space-x-2">
              <div :class="[
                'w-2 h-2 rounded-full',
                isConnected ? 'bg-green-400' : 'bg-red-400'
              ]"></div>
              <span class="text-sm text-gray-500">
                {{ isConnected ? 'Connected' : 'Disconnected' }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
// Props
const props = defineProps({
  activeTab: {
    type: String,
    required: true
  },
  currentTenant: {
    type: Number,
    required: true
  },
  isConnected: {
    type: Boolean,
    required: true
  }
})

// Emits
const emit = defineEmits(['tab-change'])

// Tab configuration
const tabs = [
  { id: 'overview', name: 'Overview' },
  { id: 'tags', name: 'Tags' },
  { id: 'entities', name: 'Entities' },
  { id: 'mappings', name: 'Tag Mappings' }
]
</script>
