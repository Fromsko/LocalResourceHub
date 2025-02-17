<template>
  <div v-if="visible" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
    <div class="bg-white rounded-lg shadow-lg w-full max-w-md p-6 relative">
      <button @click="closeModal" class="absolute top-2 right-2 text-gray-500 hover:text-gray-700 transition-colors">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
      <h2 class="text-xl font-semibold text-gray-800 mb-4">文件信息</h2>
      <div class="space-y-2">
        <template v-if="file">
          <p><strong>文件名：</strong>{{ file.file_name }}</p>
          <p><strong>大小：</strong>{{ formatFileSize(file.file_size) }}</p>
          <p><strong>类型：</strong>{{ file.content_type }}</p>
          <p><strong>上传时间：</strong>{{ formatDate(file.CreatedAt) }}</p>
        </template>
        <p v-else class="text-gray-500">没有可用的文件信息</p>
      </div>
      <div class="mt-6 flex justify-end space-x-4">
        <button
          v-if="file"
          @click="copyFileUrl"
          class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition-colors"
        >
          复制地址
        </button>
        <button 
          v-if="file"
          @click="deleteFile" 
          class="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600 transition-colors"
        >
          删除
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import axios from 'axios'

import type { FileItem } from '../types'

export default defineComponent({
  name: 'FileModal',
  props: {
    file: {
      type: Object as () => FileItem | null,
      required: false,
      default: null
    },
    visible: {
      type: Boolean,
      required: true
    }
  },
  emits: ['close', 'delete'],
  setup(props, { emit }) {
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

    const isLoading = ref(false)

    const handleOperation = async (operation: () => Promise<void>, successMessage: string, errorMessage: string) => {
      try {
        isLoading.value = true
        await operation()
        console.log(successMessage)
      } catch (error) {
        console.error(errorMessage, error)
      } finally {
        isLoading.value = false
      }
    }

    const copyFileUrl = async () => {
      if (props.file) {
        const baseUrl = window.location.origin
        const fileUrl = `http://127.0.0.1:3670/api/v1/upload/${props.file.hash}`
        await navigator.clipboard.writeText(fileUrl)
        handleOperation(async () => {}, '文件地址已复制', '文件地址复制失败')
        emit('close')
      }
    }

    const deleteFile = async () => {
      if (props.file) {
        handleOperation(
          async () => {
            await axios.delete(`http://127.0.0.1:3670/api/v1/upload/${props.file?.hash}`)
            emit('delete')
          },
          '文件已删除',
          '文件删除失败'
        )
      }
    }

    const closeModal = () => {
      emit('close')
    }

    return {
      formatFileSize,
      formatDate,
      copyFileUrl,
      deleteFile,
      closeModal
    }
  }
})
</script>
