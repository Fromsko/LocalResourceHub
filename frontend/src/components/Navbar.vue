<template>
  <div
    class="fixed top-0 left-0 right-0 z-50 bg-gray-200 text-gray-800 flex justify-between items-center px-4 py-2 border-b border-gray-300 select-none"
    :style="{ '--wails-draggable': 'drag' }"
  >
    <div class="flex items-center">
      <router-link
        to="/"
        class="text-lg font-bold tracking-wider text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-cyan-400"
      >
        Local Resource Hub
      </router-link>
    </div>
    <div class="controls flex space-x-2">
      <button
        @click="minimize"
        class="hover:bg-gray-300 p-2 rounded-full transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
        </svg>
      </button>
      <button
        @click="close"
        class="hover:bg-red-600 p-2 rounded-full transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { Quit, WindowMinimise } from '@wails/runtime/runtime'

export default defineComponent({
  name: 'Navbar',
  setup() {
    const minimize = () => {
      WindowMinimise()
    }

    const close = () => {
      Quit()
    }

    return {
      minimize,
      close
    }
  }
})
</script>
