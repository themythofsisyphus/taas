<template>
  <div id="app" class="min-h-screen bg-gray-50">
    <!-- Navigation -->
    <nav class="bg-white shadow-sm border-b">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <h1 class="text-xl font-bold text-gray-900">TaaS Dashboard</h1>
            </div>
            <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
              <button
                @click="activeTab = 'overview'"
                :class="[
                  'inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium',
                  activeTab === 'overview'
                    ? 'border-blue-500 text-gray-900'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                ]"
              >
                Overview
              </button>
              <button
                @click="activeTab = 'analytics'"
                :class="[
                  'inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium',
                  activeTab === 'analytics'
                    ? 'border-blue-500 text-gray-900'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                ]"
              >
                Analytics
              </button>
              <button
                @click="activeTab = 'endpoints'"
                :class="[
                  'inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium',
                  activeTab === 'endpoints'
                    ? 'border-blue-500 text-gray-900'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                ]"
              >
                Endpoints
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

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <!-- Overview Tab -->
      <div v-if="activeTab === 'overview'" class="px-4 py-6 sm:px-0">
        <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
          <!-- Stats Cards -->
          <div class="bg-white overflow-hidden shadow rounded-lg">
            <div class="p-5">
              <div class="flex items-center">
                <div class="flex-shrink-0">
                  <Activity class="h-6 w-6 text-gray-400" />
                </div>
                <div class="ml-5 w-0 flex-1">
                  <dl>
                    <dt class="text-sm font-medium text-gray-500 truncate">
                      Total Requests
                    </dt>
                    <dd class="text-lg font-medium text-gray-900">
                      {{ formatNumber(stats.totalRequests) }}
                    </dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-white overflow-hidden shadow rounded-lg">
            <div class="p-5">
              <div class="flex items-center">
                <div class="flex-shrink-0">
                  <Clock class="h-6 w-6 text-gray-400" />
                </div>
                <div class="ml-5 w-0 flex-1">
                  <dl>
                    <dt class="text-sm font-medium text-gray-500 truncate">
                      Avg Response Time
                    </dt>
                    <dd class="text-lg font-medium text-gray-900">
                      {{ stats.avgResponseTime }}ms
                    </dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-white overflow-hidden shadow rounded-lg">
            <div class="p-5">
              <div class="flex items-center">
                <div class="flex-shrink-0">
                  <CheckCircle class="h-6 w-6 text-green-400" />
                </div>
                <div class="ml-5 w-0 flex-1">
                  <dl>
                    <dt class="text-sm font-medium text-gray-500 truncate">
                      Success Rate
                    </dt>
                    <dd class="text-lg font-medium text-gray-900">
                      {{ stats.successRate }}%
                    </dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-white overflow-hidden shadow rounded-lg">
            <div class="p-5">
              <div class="flex items-center">
                <div class="flex-shrink-0">
                  <AlertCircle class="h-6 w-6 text-red-400" />
                </div>
                <div class="ml-5 w-0 flex-1">
                  <dl>
                    <dt class="text-sm font-medium text-gray-500 truncate">
                      Error Rate
                    </dt>
                    <dd class="text-lg font-medium text-gray-900">
                      {{ stats.errorRate }}%
                    </dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Recent Activity -->
        <div class="mt-8">
          <div class="bg-white shadow rounded-lg">
            <div class="px-4 py-5 sm:p-6">
              <h3 class="text-lg leading-6 font-medium text-gray-900 mb-4">
                Recent API Activity
              </h3>
              <div class="flow-root">
                <ul class="-mb-8">
                  <li v-for="(activity, index) in recentActivity" :key="index" class="relative pb-8">
                    <div v-if="index !== recentActivity.length - 1" class="absolute top-4 left-4 -ml-px h-full w-0.5 bg-gray-200"></div>
                    <div class="relative flex space-x-3">
                      <div>
                        <span :class="[
                          'h-8 w-8 rounded-full flex items-center justify-center ring-8 ring-white',
                          activity.status === 'success' ? 'bg-green-500' : 'bg-red-500'
                        ]">
                          <CheckCircle v-if="activity.status === 'success'" class="h-5 w-5 text-white" />
                          <XCircle v-else class="h-5 w-5 text-white" />
                        </span>
                      </div>
                      <div class="min-w-0 flex-1 pt-1.5 flex justify-between space-x-4">
                        <div>
                          <p class="text-sm text-gray-500">
                            <span class="font-medium text-gray-900">{{ activity.method }}</span>
                            {{ activity.endpoint }}
                            <span :class="[
                              'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ml-2',
                              activity.status === 'success' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                            ]">
                              {{ activity.statusCode }}
                            </span>
                          </p>
                          <p class="text-xs text-gray-400">
                            Response time: {{ activity.responseTime }}ms
                          </p>
                        </div>
                        <div class="text-right text-sm whitespace-nowrap text-gray-500">
                          {{ formatTime(activity.timestamp) }}
                        </div>
                      </div>
                    </div>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Analytics Tab -->
      <div v-if="activeTab === 'analytics'" class="px-4 py-6 sm:px-0">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- Traffic Chart -->
          <div class="bg-white shadow rounded-lg p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-4">API Traffic (Last 24 Hours)</h3>
            <div class="h-64">
              <canvas ref="trafficChart"></canvas>
            </div>
          </div>

          <!-- Response Time Chart -->
          <div class="bg-white shadow rounded-lg p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-4">Response Time Trends</h3>
            <div class="h-64">
              <canvas ref="responseTimeChart"></canvas>
            </div>
          </div>

          <!-- Status Code Distribution -->
          <div class="bg-white shadow rounded-lg p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-4">Status Code Distribution</h3>
            <div class="h-64">
              <canvas ref="statusChart"></canvas>
            </div>
          </div>

          <!-- Top Endpoints -->
          <div class="bg-white shadow rounded-lg p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-4">Most Used Endpoints</h3>
            <div class="space-y-4">
              <div v-for="endpoint in topEndpoints" :key="endpoint.path" class="flex items-center justify-between">
                <div class="flex items-center space-x-3">
                  <span :class="[
                    'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                    getMethodColor(endpoint.method)
                  ]">
                    {{ endpoint.method }}
                  </span>
                  <span class="text-sm font-medium text-gray-900">{{ endpoint.path }}</span>
                </div>
                <div class="flex items-center space-x-2">
                  <span class="text-sm text-gray-500">{{ formatNumber(endpoint.count) }}</span>
                  <div class="w-16 bg-gray-200 rounded-full h-2">
                    <div 
                      class="bg-blue-600 h-2 rounded-full" 
                      :style="{ width: (endpoint.count / Math.max(...topEndpoints.map(e => e.count)) * 100) + '%' }"
                    ></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Endpoints Tab -->
      <div v-if="activeTab === 'endpoints'" class="px-4 py-6 sm:px-0">
        <div class="bg-white shadow rounded-lg">
          <div class="px-4 py-5 sm:p-6">
            <h3 class="text-lg leading-6 font-medium text-gray-900 mb-4">
              API Endpoints
            </h3>
            <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg">
              <table class="min-w-full divide-y divide-gray-300">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Method
                    </th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Endpoint
                    </th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Requests
                    </th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Avg Response Time
                    </th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Success Rate
                    </th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Status
                    </th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="endpoint in endpoints" :key="endpoint.id">
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span :class="[
                        'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                        getMethodColor(endpoint.method)
                      ]">
                        {{ endpoint.method }}
                      </span>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                      {{ endpoint.path }}
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      {{ formatNumber(endpoint.requests) }}
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      {{ endpoint.avgResponseTime }}ms
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      {{ endpoint.successRate }}%
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap">
                      <span :class="[
                        'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                        endpoint.status === 'healthy' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                      ]">
                        {{ endpoint.status }}
                      </span>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { Activity, Clock, CheckCircle, AlertCircle, XCircle } from 'lucide-vue-next'

// Reactive data
const activeTab = ref('overview')
const currentTenant = ref(1001)
const isConnected = ref(true)

const stats = ref({
  totalRequests: 0,
  avgResponseTime: 0,
  successRate: 0,
  errorRate: 0
})

const recentActivity = ref([])
const topEndpoints = ref([])
const endpoints = ref([])

// Chart references
const trafficChart = ref(null)
const responseTimeChart = ref(null)
const statusChart = ref(null)

let trafficChartInstance = null
let responseTimeChartInstance = null
let statusChartInstance = null
let updateInterval = null

// Mock data generation
const generateMockData = () => {
  // Generate stats
  stats.value = {
    totalRequests: Math.floor(Math.random() * 100000) + 50000,
    avgResponseTime: Math.floor(Math.random() * 200) + 50,
    successRate: (Math.random() * 10 + 90).toFixed(1),
    errorRate: (Math.random() * 5 + 1).toFixed(1)
  }

  // Generate recent activity
  const methods = ['GET', 'POST', 'PUT', 'DELETE']
  const endpointPaths = [
    '/api/tags',
    '/api/tags/search',
    '/api/entities',
    '/api/user/tag_mappings/123',
    '/api/product/tag_mappings/456',
    '/api/tenants'
  ]
  
  recentActivity.value = Array.from({ length: 10 }, (_, i) => ({
    method: methods[Math.floor(Math.random() * methods.length)],
    endpoint: endpointPaths[Math.floor(Math.random() * endpointPaths.length)],
    status: Math.random() > 0.1 ? 'success' : 'error',
    statusCode: Math.random() > 0.1 ? 200 : (Math.random() > 0.5 ? 404 : 500),
    responseTime: Math.floor(Math.random() * 300) + 20,
    timestamp: new Date(Date.now() - i * 60000 * Math.random() * 10)
  }))

  // Generate top endpoints
  topEndpoints.value = [
    { method: 'GET', path: '/api/tags', count: Math.floor(Math.random() * 5000) + 2000 },
    { method: 'POST', path: '/api/tags', count: Math.floor(Math.random() * 3000) + 1000 },
    { method: 'GET', path: '/api/tags/search', count: Math.floor(Math.random() * 2000) + 800 },
    { method: 'GET', path: '/api/entities', count: Math.floor(Math.random() * 1500) + 500 },
    { method: 'POST', path: '/api/user/tag_mappings/:id', count: Math.floor(Math.random() * 1000) + 300 }
  ].sort((a, b) => b.count - a.count)

  // Generate endpoints table data
  endpoints.value = [
    {
      id: 1,
      method: 'GET',
      path: '/api/tags',
      requests: Math.floor(Math.random() * 5000) + 2000,
      avgResponseTime: Math.floor(Math.random() * 100) + 50,
      successRate: (Math.random() * 5 + 95).toFixed(1),
      status: 'healthy'
    },
    {
      id: 2,
      method: 'POST',
      path: '/api/tags',
      requests: Math.floor(Math.random() * 3000) + 1000,
      avgResponseTime: Math.floor(Math.random() * 150) + 80,
      successRate: (Math.random() * 5 + 95).toFixed(1),
      status: 'healthy'
    },
    {
      id: 3,
      method: 'GET',
      path: '/api/tags/search',
      requests: Math.floor(Math.random() * 2000) + 800,
      avgResponseTime: Math.floor(Math.random() * 200) + 100,
      successRate: (Math.random() * 5 + 95).toFixed(1),
      status: 'healthy'
    },
    {
      id: 4,
      method: 'GET',
      path: '/api/entities',
      requests: Math.floor(Math.random() * 1500) + 500,
      avgResponseTime: Math.floor(Math.random() * 80) + 40,
      successRate: (Math.random() * 5 + 95).toFixed(1),
      status: 'healthy'
    },
    {
      id: 5,
      method: 'POST',
      path: '/api/entities',
      requests: Math.floor(Math.random() * 800) + 200,
      avgResponseTime: Math.floor(Math.random() * 120) + 60,
      successRate: (Math.random() * 5 + 95).toFixed(1),
      status: 'healthy'
    },
    {
      id: 6,
      method: 'GET',
      path: '/api/:entity_type/tag_mappings/:id',
      requests: Math.floor(Math.random() * 1200) + 400,
      avgResponseTime: Math.floor(Math.random() * 90) + 45,
      successRate: (Math.random() * 5 + 95).toFixed(1),
      status: 'healthy'
    },
    {
      id: 7,
      method: 'POST',
      path: '/api/:entity_type/tag_mappings/:id',
      requests: Math.floor(Math.random() * 600) + 150,
      avgResponseTime: Math.floor(Math.random() * 140) + 70,
      successRate: (Math.random() * 5 + 95).toFixed(1),
      status: 'healthy'
    },
    {
      id: 8,
      method: 'DELETE',
      path: '/api/:entity_type/tag_mappings/:id',
      requests: Math.floor(Math.random() * 300) + 50,
      avgResponseTime: Math.floor(Math.random() * 60) + 30,
      successRate: (Math.random() * 5 + 95).toFixed(1),
      status: 'healthy'
    }
  ]
}

// Chart initialization
const initCharts = async () => {
  // Dynamically import Chart.js
  const { Chart, registerables } = await import('chart.js')
  Chart.register(...registerables)

  // Traffic Chart
  if (trafficChart.value) {
    const ctx = trafficChart.value.getContext('2d')
    const hours = Array.from({ length: 24 }, (_, i) => {
      const hour = new Date()
      hour.setHours(hour.getHours() - (23 - i))
      return hour.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' })
    })
    
    trafficChartInstance = new Chart(ctx, {
      type: 'line',
      data: {
        labels: hours,
        datasets: [{
          label: 'Requests',
          data: Array.from({ length: 24 }, () => Math.floor(Math.random() * 500) + 100),
          borderColor: 'rgb(59, 130, 246)',
          backgroundColor: 'rgba(59, 130, 246, 0.1)',
          tension: 0.4,
          fill: true
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            display: false
          }
        },
        scales: {
          y: {
            beginAtZero: true
          }
        }
      }
    })
  }

  // Response Time Chart
  if (responseTimeChart.value) {
    const ctx = responseTimeChart.value.getContext('2d')
    responseTimeChartInstance = new Chart(ctx, {
      type: 'line',
      data: {
        labels: Array.from({ length: 24 }, (_, i) => {
          const hour = new Date()
          hour.setHours(hour.getHours() - (23 - i))
          return hour.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' })
        }),
        datasets: [{
          label: 'Response Time (ms)',
          data: Array.from({ length: 24 }, () => Math.floor(Math.random() * 200) + 50),
          borderColor: 'rgb(16, 185, 129)',
          backgroundColor: 'rgba(16, 185, 129, 0.1)',
          tension: 0.4,
          fill: true
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            display: false
          }
        },
        scales: {
          y: {
            beginAtZero: true
          }
        }
      }
    })
  }

  // Status Chart
  if (statusChart.value) {
    const ctx = statusChart.value.getContext('2d')
    statusChartInstance = new Chart(ctx, {
      type: 'doughnut',
      data: {
        labels: ['2xx Success', '4xx Client Error', '5xx Server Error'],
        datasets: [{
          data: [85, 12, 3],
          backgroundColor: [
            'rgb(16, 185, 129)',
            'rgb(245, 158, 11)',
            'rgb(239, 68, 68)'
          ]
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            position: 'bottom'
          }
        }
      }
    })
  }
}

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

const getMethodColor = (method) => {
  const colors = {
    GET: 'bg-blue-100 text-blue-800',
    POST: 'bg-green-100 text-green-800',
    PUT: 'bg-yellow-100 text-yellow-800',
    DELETE: 'bg-red-100 text-red-800'
  }
  return colors[method] || 'bg-gray-100 text-gray-800'
}

// Lifecycle hooks
onMounted(async () => {
  generateMockData()
  await initCharts()
  
  // Update data every 30 seconds
  updateInterval = setInterval(() => {
    generateMockData()
    
    // Update charts with new data
    if (trafficChartInstance) {
      trafficChartInstance.data.datasets[0].data = Array.from({ length: 24 }, () => 
        Math.floor(Math.random() * 500) + 100
      )
      trafficChartInstance.update('none')
    }
    
    if (responseTimeChartInstance) {
      responseTimeChartInstance.data.datasets[0].data = Array.from({ length: 24 }, () => 
        Math.floor(Math.random() * 200) + 50
      )
      responseTimeChartInstance.update('none')
    }
  }, 30000)
})

onUnmounted(() => {
  if (updateInterval) {
    clearInterval(updateInterval)
  }
  if (trafficChartInstance) {
    trafficChartInstance.destroy()
  }
  if (responseTimeChartInstance) {
    responseTimeChartInstance.destroy()
  }
  if (statusChartInstance) {
    statusChartInstance.destroy()
  }
})
</script>

<style>
#app {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}
</style>
