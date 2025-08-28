<template>
  <div class="px-4 py-6 sm:px-0">
    <div class="bg-white shadow rounded-lg">
      <div class="px-4 py-5 sm:p-6">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg leading-6 font-medium text-gray-900">
            Tag Mappings
          </h3>
          <button class="bg-purple-600 text-white px-4 py-2 rounded-md hover:bg-purple-700">
            Create New Mapping
          </button>
        </div>

        <!-- Mapping Statistics -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
          <div class="bg-gray-50 p-4 rounded-lg">
            <div class="text-sm font-medium text-gray-500">Total Mappings</div>
            <div class="text-2xl font-bold text-gray-900">{{ formatNumber(mappingStats.total) }}</div>
          </div>
          <div class="bg-gray-50 p-4 rounded-lg">
            <div class="text-sm font-medium text-gray-500">Users Tagged</div>
            <div class="text-2xl font-bold text-gray-900">{{ formatNumber(mappingStats.usersTagged) }}</div>
          </div>
          <div class="bg-gray-50 p-4 rounded-lg">
            <div class="text-sm font-medium text-gray-500">Products Tagged</div>
            <div class="text-2xl font-bold text-gray-900">{{ formatNumber(mappingStats.productsTagged) }}</div>
          </div>
          <div class="bg-gray-50 p-4 rounded-lg">
            <div class="text-sm font-medium text-gray-500">Avg Tags/Entity</div>
            <div class="text-2xl font-bold text-gray-900">{{ mappingStats.avgTagsPerEntity.toFixed(1) }}</div>
          </div>
        </div>

        <!-- Recent Mappings -->
        <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg">
          <table class="min-w-full divide-y divide-gray-300">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Entity
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Type
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Tags Applied
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Mapped At
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="mapping in recentMappings" :key="mapping.id">
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {{ mapping.entityName }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-gray-100 text-gray-800">
                    {{ mapping.entityType }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <div class="flex flex-wrap gap-1">
                    <span v-for="tag in mapping.tags" :key="tag" 
                          class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-blue-100 text-blue-800">
                      {{ tag }}
                    </span>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatTime(mapping.mappedAt) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <button class="text-blue-600 hover:text-blue-900 mr-3">Edit</button>
                  <button class="text-red-600 hover:text-red-900">Remove</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Props
const props = defineProps({
  mappingStats: {
    type: Object,
    required: true
  },
  recentMappings: {
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
</script>
