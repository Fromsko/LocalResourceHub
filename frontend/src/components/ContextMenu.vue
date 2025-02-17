<template>
  <div
    v-if="visible"
    :style="{ top: position.y + 'px', left: position.x + 'px' }"
    class="fixed z-50 bg-white shadow-lg rounded-lg overflow-hidden border border-gray-200"
  >
    <ul class="divide-y divide-gray-100">
      <li
        v-for="(item, index) in menuItems"
        :key="index"
        @click="handleMenuItemClick(item)"
        class="px-4 py-2 text-gray-700 hover:bg-gray-100 cursor-pointer transition-colors"
      >
        {{ item.label }}
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'ContextMenu',
  setup() {
    const visible = ref(false)
    const position = ref({ x: 0, y: 0 })
    const router = useRouter()

    const menuItems = [
      { label: '主页', route: '/' },
      { label: '文件列表', route: '/files' },
      { label: '后台', route: '/admin' },
      { label: '项目地址', route: 'https://github.com/fromsko' }
    ]

    const showMenu = (event: MouseEvent) => {
      event.preventDefault()
      visible.value = true
      position.value = { x: event.clientX, y: event.clientY }
    }

    const hideMenu = () => {
      visible.value = false
    }

    const handleMenuItemClick = (item: { label: string; route: string }) => {
      if (item.route.startsWith('http')) {
        window.open(item.route, '_blank')
      } else {
        router.push(item.route)
      }
      hideMenu()
    }

    return {
      visible,
      position,
      menuItems,
      showMenu,
      hideMenu,
      handleMenuItemClick
    }
  }
})
</script>
