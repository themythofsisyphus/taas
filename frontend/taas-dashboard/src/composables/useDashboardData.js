import { ref, onMounted, onUnmounted } from 'vue'

export function useDashboardData() {
  // Reactive data
  const stats = ref({
    totalTags: 0,
    totalEntities: 0,
    totalMappings: 0,
    apiRequests: 0
  })

  const recentOperations = ref([])
  const topTags = ref([])
  const tags = ref([])
  const entities = ref([])
  const entityTypes = ref([])
  const mappingStats = ref({
    total: 0,
    usersTagged: 0,
    productsTagged: 0,
    avgTagsPerEntity: 0
  })
  const recentMappings = ref([])

  let updateInterval = null

  // Mock data generation
  const generateMockData = () => {
    // Generate stats
    stats.value = {
      totalTags: Math.floor(Math.random() * 500) + 100,
      totalEntities: Math.floor(Math.random() * 2000) + 500,
      totalMappings: Math.floor(Math.random() * 5000) + 1000,
      apiRequests: Math.floor(Math.random() * 10000) + 5000
    }

    // Generate recent operations
    const operationTypes = ['CREATE', 'UPDATE', 'DELETE', 'MAP', 'UNMAP']
    const operationDescriptions = [
      'Created new tag',
      'Updated tag name',
      'Deleted unused tag',
      'Mapped tags to entity',
      'Removed tag mapping'
    ]
    
    recentOperations.value = Array.from({ length: 8 }, (_, i) => ({
      type: operationTypes[Math.floor(Math.random() * operationTypes.length)],
      description: operationDescriptions[Math.floor(Math.random() * operationDescriptions.length)],
      details: `Operation completed successfully`,
      timestamp: new Date(Date.now() - i * 60000 * Math.random() * 10)
    }))

    // Generate top tags
    topTags.value = [
      { name: 'premium', entityType: 'users', usageCount: Math.floor(Math.random() * 200) + 100 },
      { name: 'featured', entityType: 'products', usageCount: Math.floor(Math.random() * 150) + 80 },
      { name: 'verified', entityType: 'users', usageCount: Math.floor(Math.random() * 120) + 60 },
      { name: 'trending', entityType: 'products', usageCount: Math.floor(Math.random() * 100) + 50 },
      { name: 'archived', entityType: 'content', usageCount: Math.floor(Math.random() * 80) + 30 }
    ].sort((a, b) => b.usageCount - a.usageCount)

    // Generate tags for table
    tags.value = Array.from({ length: 15 }, (_, i) => ({
      id: i + 1,
      name: `tag-${i + 1}`,
      usageCount: Math.floor(Math.random() * 100) + 10,
      createdAt: new Date(Date.now() - Math.random() * 30 * 24 * 60 * 60 * 1000)
    }))

    // Generate entities
    entities.value = Array.from({ length: 20 }, (_, i) => ({
      id: i + 1,
      name: `Entity ${i + 1}`,
      type: ['user', 'product', 'content'][Math.floor(Math.random() * 3)],
      tags: Array.from({ length: Math.floor(Math.random() * 5) + 1 }, (_, j) => `tag-${j + 1}`)
    }))

    // Generate entity types
    entityTypes.value = [
      { name: 'Users', count: Math.floor(Math.random() * 1000) + 500, taggedCount: Math.floor(Math.random() * 800) + 400 },
      { name: 'Products', count: Math.floor(Math.random() * 800) + 300, taggedCount: Math.floor(Math.random() * 600) + 200 },
      { name: 'Content', count: Math.floor(Math.random() * 600) + 200, taggedCount: Math.floor(Math.random() * 400) + 100 }
    ]

    // Generate mapping stats
    mappingStats.value = {
      total: Math.floor(Math.random() * 5000) + 1000,
      usersTagged: Math.floor(Math.random() * 800) + 400,
      productsTagged: Math.floor(Math.random() * 600) + 200,
      avgTagsPerEntity: (Math.random() * 3 + 2)
    }

    // Generate recent mappings
    recentMappings.value = Array.from({ length: 10 }, (_, i) => ({
      id: i + 1,
      entityName: `Entity ${i + 1}`,
      entityType: ['user', 'product', 'content'][Math.floor(Math.random() * 3)],
      tags: Array.from({ length: Math.floor(Math.random() * 4) + 1 }, (_, j) => `tag-${j + 1}`),
      mappedAt: new Date(Date.now() - i * 60000 * Math.random() * 10)
    }))
  }

  // Lifecycle hooks
  onMounted(() => {
    generateMockData()
    
    // Update data every 30 seconds
    updateInterval = setInterval(() => {
      generateMockData()
    }, 30000)
  })

  onUnmounted(() => {
    if (updateInterval) {
      clearInterval(updateInterval)
    }
  })

  return {
    stats,
    recentOperations,
    topTags,
    tags,
    entities,
    entityTypes,
    mappingStats,
    recentMappings
  }
}
