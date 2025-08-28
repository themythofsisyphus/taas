<template>
  <div class="px-4 py-6 sm:px-0">
    <!-- Stats Cards -->
    <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
      <!-- Total Tags -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <Tag class="h-6 w-6 text-blue-400" />
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">
                  Total Tags
                </dt>
                <dd class="text-lg font-medium text-gray-900">
                  {{ formatNumber(stats.totalTags) }}
                </dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <!-- Total Entities -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <Database class="h-6 w-6 text-green-400" />
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">
                  Total Entities
                </dt>
                <dd class="text-lg font-medium text-gray-900">
                  {{ formatNumber(stats.totalEntities) }}
                </dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <!-- Total Mappings -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <Link class="h-6 w-6 text-purple-400" />
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">
                  Tag Mappings
                </dt>
                <dd class="text-lg font-medium text-gray-900">
                  {{ formatNumber(stats.totalMappings) }}
                </dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <!-- API Requests -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <Activity class="h-6 w-6 text-orange-400" />
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">
                  API Requests (24h)
                </dt>
                <dd class="text-lg font-medium text-gray-900">
                  {{ formatNumber(stats.apiRequests) }}
                </dd>
              </dl>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Recent Tag Operations -->
    <div class="mt-8">
      <div class="bg-white shadow rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg leading-6 font-medium text-gray-900 mb-4">
            Recent Tag Operations
          </h3>
          <div class="flow-root">
            <ul class="-mb-8">
              <li v-for="(operation, index) in recentOperations" :key="index" class="relative pb-8">
                <div v-if="index !== recentOperations.length - 1" class="absolute top-4 left-4 -ml-px h-full w-0.5 bg-gray-200"></div>
                <div class="relative flex space-x-3">
                  <div>
                    <span :class="[
                      'h-8 w-8 rounded-full flex items-center justify-center ring-8 ring-white',
                      getOperationColor(operation.type)
                    ]">
                      <component :is="getOperationIcon(operation.type)" class="h-5 w-5 text-white" />
                    </span>
                  </div>
                  <div class="min-w-0 flex-1 pt-1.5 flex justify-between space-x-4">
                    <div>
                      <p class="text-sm text-gray-500">
                        <span class="font-medium text-gray-900">{{ operation.type }}</span>
                        {{ operation.description }}
                      </p>
                      <p class="text-xs text-gray-400">
                        {{ operation.details }}
                      </p>
                    </div>
                    <div class="text-right text-sm whitespace-nowrap text-gray-500">
                      {{ formatTime(operation.timestamp) }}
                    </div>
                  </div>
                </div>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- Tag Usage Chart -->
    <div class="mt-8">
      <div class="bg-white shadow rounded-lg p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">Most Used Tags</h3>
        <div class="space-y-4">
          <div v-for="tag in topTags" :key="tag.name" class="flex items-center justify-between">
            <div class="flex items-center space-x-3">
              <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                {{ tag.name }}
              </span>
              <span class="text-sm text-gray-500">{{ tag.entityType }}</span>
            </div>
            <div class="flex items-center space-x-2">
              <span class="text-sm text-gray-500">{{ formatNumber(tag.usageCount) }} mappings</span>
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div 
                  class="bg-blue-600 h-2 rounded-full" 
                  :style="{ width: (tag.usageCount / Math.max(...topTags.map(t => t.usageCount)) * 100) + '%' }"
                ></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Tag, Database, Link, Activity, Plus, Edit, Trash2 } from 'lucide-vue-next'

// Props
const props = defineProps({
  stats: {
    type: Object,
    required: true
  },
  recentOperations: {
    type: Array,
    required: true
  },
  topTags: {
    type: Array,
    required: true
  }
})

// Utility functions
const formatNumber = (num) => {
  return new Intl.NumberFormat().format(num)
}

const formatTime = (timestamp) => {
  return new Date(timestamp).toLocaleTimeString('en-US', {
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getOperationColor = (type) => {
  const colors = {
    CREATE: 'bg-green-500',
    UPDATE: 'bg-blue-500',
    DELETE: 'bg-red-500',
    MAP: 'bg-purple-500',
    UNMAP: 'bg-orange-500'
  }
  return colors[type] || 'bg-gray-500'
}

const getOperationIcon = (type) => {
  const icons = {
    CREATE: Plus,
    UPDATE: Edit,
    DELETE: Trash2,
    MAP: Link,
    UNMAP: Link
  }
  return icons[type] || Activity
}
</script>
