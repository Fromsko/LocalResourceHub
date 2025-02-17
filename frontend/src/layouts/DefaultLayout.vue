<template>
  <Navbar />
  <div class="flex flex-col h-screen" @contextmenu="handleContextMenu">
    <div class="flex-1 overflow-y-auto mt-12 px-4 py-6">
      <router-view />
    </div>
    <ContextMenu ref="contextMenuRef" />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import Navbar from '@/components/Navbar.vue'
import ContextMenu from '@/components/ContextMenu.vue'

export default defineComponent({
  name: 'DefaultLayout',
  components: {
    Navbar,
    ContextMenu
  },
  setup() {
    const contextMenuRef = ref<typeof ContextMenu | null>(null)

    const handleContextMenu = (event: MouseEvent) => {
      // Ignore right-click on the Navbar area
      // Ignore right-click on the Navbar area by checking target element's class or tag
      const target = event.target as HTMLElement
      if (target.closest('.navbar') || target.tagName === 'NAV') {
        return
      }

      if (contextMenuRef.value && typeof contextMenuRef.value.showMenu === 'function') {
        contextMenuRef.value.showMenu(event)
      }
    }

    return {
      contextMenuRef,
      handleContextMenu
    }
  }
})
</script>
