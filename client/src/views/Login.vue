<template>
  <div class="flex justify-center items-center mt-12 px-4">
    <div class="w-full max-w-md">
      <form @submit.prevent="login" class="bg-gray-100 shadow-md rounded px-8 pt-6 pb-8 mb-4">
        <h3 class="mb-4 font-bold text-gray-700">Login</h3>
				<div class="mb-4">
					<label class="block text-gray-700 text-sm font-bold mb-2" for="email">
						Email
					</label>
					<input
						class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
						id="email" type="email" placeholder="Email" v-model="emailAddress">
				</div>
				<div class="mb-6">
					<label class="block text-gray-700 text-sm font-bold mb-2" for="password">
						Password
					</label>
					<input
						class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
						id="password" type="password" placeholder="Password" v-model="password">
				</div>
				<div class="flex items-center justify-between">
					<button
						class="bg-blue-500 hover:bg-blue-800 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
						type="submit">
						Sign In
					</button>
				</div>
			</form>
		</div>
	</div>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { useMainStore } from '@/stores/main'
import { useRouter } from "vue-router";

const store = useMainStore()
const router = useRouter()
const emailAddress = ref('')
const password = ref('')

const login = async () => {
  await store.login(emailAddress.value, password.value)
  router.push("/dashboard")
}
</script>
