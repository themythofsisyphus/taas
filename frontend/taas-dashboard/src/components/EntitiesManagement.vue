<template>
  <div class="px-4 py-6 sm:px-0">
    <div class="bg-white shadow rounded-lg">
      <div class="px-4 py-5 sm:p-6">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg leading-6 font-medium text-gray-900">
            Entity Management
          </h3>
          <button class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700">
            Create New Entity
          </button>
        </div>

        <!-- Entity Types Overview -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
          <div v-for="entityType in entityTypes" :key="entityType.name" class="bg-gray-50 p-4 rounded-lg">
            <div class="flex items-center justify-between">
              <div>
                <h4 class="text-sm font-medium text-gray-900">{{ entityType.name }}</h4>
                <p class="text-2xl font-bold text-gray-900">{{ formatNumber(entityType.count) }}</p>
              </div>
              <div class="text-sm text-gray-500">
                {{ entityType.taggedCount }} tagged
              </div>
            </div>
          </div>
        </div>

        <!-- Entities Table -->
        <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg">
          <table class="min-w-full divide-y divide-gray-300">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Entity Name
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Type
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Tags
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="entity in entities" :key="entity.id">
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {{ entity.name }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-gray-100 text-gray-800">
                    {{ entity.type }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <div class="flex flex-wrap gap-1">
                    <span v-for="tag in entity.tags.slice(0, 3)" :key="tag" 
                          class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-blue-100 text-blue-800">
                      {{ tag }}
                    </span>
                    <span v-if="entity.tags.length > 3" class="text-xs text-gray-400">
                      +{{ entity.tags.length - 3 }} more
                    </span>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <button class="text-blue-600 hover:text-blue-900 mr-3">Manage Tags</button>
                  <button class="text-red-600 hover:text-red-900">Delete</button>
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
  entities: {
    type: Array,
    required: true
  },
  entityTypes: {
    type: Array,
    required: true
  }
})

// Utility functions
const formatNumber = (num) => {
  return new Intl.NumberFormat().format(num)
}
</script>
