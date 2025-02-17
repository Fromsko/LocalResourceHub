<template>
  <div class="flex flex-col h-full">
    <div class="flex-1 overflow-y-auto scrollbar-thin scrollbar-thumb-rounded-full scrollbar-thumb-gray-400 px-4 py-6">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="file in files"
          :key="file.ID"
          class="bg-white rounded-lg shadow-md overflow-hidden transition-transform duration-200 hover:scale-105 cursor-pointer"
          @click="openModal(file)"
        >
          <div class="p-4">
            <h3 class="text-lg font-semibold text-gray-800 truncate">{{ file.file_name }}</h3>
            <p class="text-sm text-gray-500 mt-2">{{ formatFileSize(file.file_size) }}</p>
            <p class="text-sm text-gray-500 mt-1">{{ formatDate(file.CreatedAt) }}</p>
          </div>
        </div>
      </div>
    </div>
    <FileModal
      v-if="isModalVisible"
      :file="selectedFile"
      :visible="isModalVisible"
      @close="closeModal"
      @delete="deleteFile"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import type { FileItem } from '../types'

import axios from 'axios'
import FileModal from './FileModal.vue'

export default defineComponent({
  name: 'FileList',
  components: {
    FileModal
  },
  setup() {
      const files = ref<FileItem[]>([])
      const selectedFile = ref<FileItem | null>(null)
    const isModalVisible = ref(false)

    const fetchFiles = async () => {
      try {
        const response = await axios.get('http://127.0.0.1:3670/api/v1/files')
        files.value = response.data.data.files
      } catch (error) {
        console.error('Failed to fetch files:', error)
      }
    }

    onMounted(() => {
      fetchFiles()
    })

    const formatFileSize = (size: number): string => {
      if (size < 1024) return `${size} B`
      if (size < 1024 * 1024) return `${(size / 1024).toFixed(2)} KB`
      if (size < 1024 * 1024 * 1024) return `${(size / (1024 * 1024)).toFixed(2)} MB`
      return `${(size / (1024 * 1024 * 1024)).toFixed(2)} GB`
    }

    const formatDate = (dateString: string): string => {
      const date = new Date(dateString)
      return date.toLocaleString()
    }

    const openModal = (file: any) => {
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
            const currentFile = selectedFile.value // Create a local reference
            await axios.delete(`http://127.0.0.1:3670/api/v1/upload/${currentFile.hash}`)
            files.value = files.value.filter(f => f.hash !== currentFile.hash)
            closeModal()
            fetchFiles()
          } catch (error) {
            console.error('Failed to delete file:', error)
          }
        }
    }

    return {
      files,
      formatFileSize,
      formatDate,
      selectedFile,
      isModalVisible,
      openModal,
      closeModal,
      deleteFile
    }
  }
})
</script>
