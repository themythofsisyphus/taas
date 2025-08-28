<template>
  <div id="app" class="min-h-screen bg-gray-50">
    <!-- Navigation -->
    <AppNavigation
      :active-tab="activeTab"
      :current-tenant="currentTenant"
      :is-connected="isConnected"
      @tab-change="activeTab = $event"
    />

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <!-- Overview Tab -->
      <DashboardOverview
        v-if="activeTab === 'overview'"
        :stats="stats"
        :recent-operations="recentOperations"
        :top-tags="topTags"
      />

      <!-- Tags Tab -->
      <TagsManagement
        v-if="activeTab === 'tags'"
        :tags="tags"
      />

      <!-- Entities Tab -->
      <EntitiesManagement
        v-if="activeTab === 'entities'"
        :entities="entities"
        :entity-types="entityTypes"
      />

      <!-- Tag Mappings Tab -->
      <TagMappings
        v-if="activeTab === 'mappings'"
        :mapping-stats="mappingStats"
        :recent-mappings="recentMappings"
      />
    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import AppNavigation from './components/AppNavigation.vue'
import DashboardOverview from './components/DashboardOverview.vue'
import TagsManagement from './components/TagsManagement.vue'
import EntitiesManagement from './components/EntitiesManagement.vue'
import TagMappings from './components/TagMappings.vue'
import { useDashboardData } from './composables/useDashboardData.js'

// Tab management
const activeTab = ref('overview')

// Connection status
const currentTenant = ref(1001)
const isConnected = ref(true)

// Dashboard data
const {
  stats,
  recentOperations,
  topTags,
  tags,
  entities,
  entityTypes,
  mappingStats,
  recentMappings
} = useDashboardData()
</script>

<style>
#app {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}
</style>
