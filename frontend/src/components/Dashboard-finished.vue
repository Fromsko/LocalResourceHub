<script setup lang="ts">
import axios from "axios"
import { onMounted, ref } from "vue"
import { Quit, WindowMinimise } from "../../wailsjs/runtime/runtime"
import FileModal from "./FileModal.vue"

// ========== 侧边栏状态 ==========
const isSidebarOpen = ref(true)
const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}

// ========== 文件数据 ==========
interface FileItem {
  ID: number
  file_name: string
  file_size: number
  hash: string
  CreatedAt: string
}

const files = ref<FileItem[]>([])
const selectedFile = ref<FileItem | null>(null)
const isModalVisible = ref(false)

const fetchFiles = async () => {
  try {
    const response = await axios.get("http://127.0.0.1:3670/api/v1/files")
    files.value = response.data.data.files
  } catch (error) {
    console.error("Failed to fetch files:", error)
  }
}

onMounted(() => {
  fetchFiles()
})

const formatFileSize = (size: number): string => {
  if (size < 1024) return `${size} B`
  if (size < 1024 * 1024) return `${(size / 1024).toFixed(2)} KB`
  if (size < 1024 * 1024 * 1024)
    return `${(size / (1024 * 1024)).toFixed(2)} MB`
  return `${(size / (1024 * 1024 * 1024)).toFixed(2)} GB`
}

const formatDate = (dateString: string): string => {
  const date = new Date(dateString)
  return date.toLocaleString()
}

const openModal = (file: FileItem) => {
  selectedFile.value = file
  isModalVisible.value = true
}

const closeModal = () => {
  isModalVisible.value = false
  selectedFile.value = null
}

const deleteFile = async () => {
  if (selectedFile.value) {
    try {
      const currentFile = selectedFile.value
      await axios.delete(
        `http://127.0.0.1:3670/api/v1/upload/${currentFile.hash}`
      )
      files.value = files.value.filter((f) => f.hash !== currentFile.hash)
      closeModal()
      fetchFiles()
    } catch (error) {
      console.error("Failed to delete file:", error)
    }
  }
}
</script>

<template>
  <div class="bg-gray-100 text-gray-800 font-sans select-none">
    <div class="flex flex-col h-screen">
      <!-- 顶部导航栏 -->
      <header class="bg-white shadow border-b" style="--wails-draggable: drag">
        <div class="flex items-center justify-between px-6 py-4">
          <div class="flex items-center">
            <!-- 移动端菜单按钮 -->
            <button
              id="menu-toggle"
              class="text-gray-600 hover:text-gray-800"
              @click="toggleSidebar"
            >
              <svg
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4 6h16M4 12h16M4 18h16"
                />
              </svg>
            </button>
            <h2 class="text-xl font-semibold ml-2 lg:ml-0">文件管理</h2>
          </div>
          <div class="flex items-center">
            <button class="btn-min" @click="WindowMinimise">-</button>
            <button class="btn-close" @click="Quit">x</button>
          </div>
        </div>
      </header>

      <div class="flex-1 flex overflow-hidden">
        <!-- 侧边栏 -->
        <transition name="slide">
          <aside
            v-show="isSidebarOpen"
            id="sidebar"
            class="bg-white text-gray-800 w-36 flex-shrink-0 lg:block border-r"
          >
            <nav class="mt-4 space-y-1">
              <!-- 主页：本地路由 -->
              <RouterLink
                to="/"
                class="side-item"
                active-class="font-bold text-blue-600"
              >
                主页
              </RouterLink>

              <!-- 项目：外部链接，跳转到 GitHub 并新窗口打开 -->
              <a
                href="https://github.com/fromsko"
                target="_blank"
                rel="noopener noreferrer"
                class="side-item"
              >
                项目
              </a>
            </nav>
          </aside>
        </transition>

        <!-- 主要内容 -->
        <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-200 p-6">
          <!-- 文件数量 -->
          <div
            class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6"
          >
            <div class="bg-white rounded-lg shadow-md p-6">
              <h3 class="text-lg font-semibold mb-2">文件总数</h3>
              <p class="text-3xl font-bold">{{ files.length }}</p>
            </div>
          </div>

          <!-- 上传日志 -->
          <div class="bg-white rounded-lg shadow-md p-6">
            <h3 class="text-xl font-semibold mb-4">上传日志</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              <div
                v-for="file in files"
                :key="file.ID"
                class="bg-white rounded-lg shadow-md p-4 transition-transform duration-200 hover:scale-105 cursor-pointer"
                @click="openModal(file)"
              >
                <h3 class="text-lg font-semibold text-gray-800 truncate">
                  {{ file.file_name }}
                </h3>
                <p class="text-sm text-gray-500 mt-2">
                  {{ formatFileSize(file.file_size) }}
                </p>
                <p class="text-sm text-gray-500 mt-1">
                  {{ formatDate(file.CreatedAt) }}
                </p>
              </div>
            </div>
          </div>
        </main>
      </div>
    </div>

    <!-- 文件详情弹窗 -->
    <FileModal
      v-if="isModalVisible"
      :file="selectedFile"
      :visible="isModalVisible"
      @close="closeModal"
      @delete="deleteFile"
    />
  </div>
</template>
