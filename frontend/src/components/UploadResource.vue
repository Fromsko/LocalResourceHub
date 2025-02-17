<template>
  <div class="flex items-center justify-center bg-gray-100 mt-8">
    <div class="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
      <h1 class="text-2xl font-bold mb-6 text-center">Upload Resource</h1>
      <div
        class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center cursor-pointer hover:border-blue-500 transition-colors"
        @dragover.prevent
        @drop="handleDrop"
        @click="triggerFileInput"
      >
        <input type="file" ref="fileInput" class="hidden" @change="handleFileChange" />
        <p v-if="!fileName" class="text-gray-500">Drag & drop files here or click to select files</p>
        <p v-else class="text-gray-700">Selected File: {{ fileName }} ({{ fileSize }})</p>
      </div>
      <button
        @click="uploadFile"
        class="w-full bg-blue-600 text-white py-2 px-4 rounded mt-4 hover:bg-blue-700 transition duration-300 flex items-center justify-center"
        :disabled="!file || isUploading"
      >
        <svg
          v-if="isUploading"
          class="animate-spin h-5 w-5 mr-2"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
        >
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path
            class="opacity-75"
            fill="currentColor"
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
          ></path>
        </svg>
        {{ isUploading ? 'Uploading...' : 'Upload' }}
      </button>
      <div v-if="progress > 0 && progress < 100" class="mt-4 bg-gray-200 rounded-full overflow-hidden">
        <div class="h-4 bg-blue-500" :style="{ width: `${progress}%` }"></div>
      </div>
      <p
        v-if="message"
        class="mt-4 text-center"
        :class="message.includes('success') ? 'text-green-500' : 'text-red-500'"
      >
        {{ message }}
      </p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import axios from 'axios'

export default defineComponent({
  name: 'UploadResource',
  setup() {
    const file = ref<File | null>(null)
    const fileName = ref<string>('')
    const fileSize = ref<string>('')
    const message = ref<string>('')
    const progress = ref<number>(0)
    const isUploading = ref<boolean>(false)

    const handleFileChange = (event: Event) => {
      const target = event.target as HTMLInputElement
      if (target.files) {
        file.value = target.files[0]
        fileName.value = file.value.name
        fileSize.value = formatFileSize(file.value.size)
      }
    }

    const handleDrop = (event: DragEvent) => {
      event.preventDefault()
      if (event.dataTransfer?.files) {
        file.value = event.dataTransfer.files[0]
        fileName.value = file.value.name
        fileSize.value = formatFileSize(file.value.size)
      }
    }

    const triggerFileInput = () => {
      const fileInput = document.querySelector<HTMLInputElement>('.hidden')
      if (fileInput) {
        fileInput.click()
      }
    }

    const uploadFile = async () => {
      if (!file.value) {
        message.value = 'Please select a file to upload.'
        return
      }

      isUploading.value = true
      progress.value = 0

      const formData = new FormData()
      formData.append('file', file.value)

      try {
        const response = await axios.post('http://0.0.0.0:3670/api/v1/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          },
          onUploadProgress: (progressEvent) => {
            if (progressEvent.total) {
              progress.value = Math.round((progressEvent.loaded * 100) / progressEvent.total)
            }
          }
        })
        const responseData = response.data

        // Handle success
        if (responseData.code === 200) {
          message.value = `File uploaded successfully! File Name: ${responseData.data.file_name}`
          resetUploadState()
        }
        // Handle known errors
        else if (responseData.code === 413) {
          message.value = 'Error: File size exceeds the allowed limit.'
        } else if (responseData.code === 415) {
          message.value = 'Error: Unsupported file type.'
        } else if (responseData.code === 4002) {
          message.value = 'Error: File upload failed. Please try again.'
        } else if (responseData.code === 4006) {
          message.value = 'Error: File already exists.'
        }
        // Handle unknown errors
        else {
          message.value = `Error: ${responseData.message || 'An unexpected error occurred.'}`
        }
      } catch (error) {
        console.log(error)

        message.value = 'Error: An unexpected error occurred. Please try again.'
      } finally {
        isUploading.value = false
      }
    }

    const resetUploadState = () => {
      file.value = null
      fileName.value = ''
      fileSize.value = ''
      progress.value = 0
    }

    const formatFileSize = (size: number): string => {
      const units = ['B', 'KB', 'MB', 'GB']
      let unitIndex = 0
      while (size >= 1024 && unitIndex < units.length - 1) {
        size /= 1024
        unitIndex++
      }
      return `${size.toFixed(2)} ${units[unitIndex]}`
    }

    return {
      file,
      fileName,
      fileSize,
      message,
      progress,
      isUploading,
      handleFileChange,
      handleDrop,
      triggerFileInput,
      uploadFile
    }
  }
})
</script>
