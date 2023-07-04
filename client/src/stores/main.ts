import { defineStore } from 'pinia'
import api from '../plugins/api'

export const useMainStore = defineStore('main', {
	state: () => ({
		token: "",
		user: {}
	}),
	actions: {
		async login(email: string, password: string) {
			try {
				const res = await api.post('/auth/login', { email: email, password: password })
				this.token = res.data.access_token
				if (this.token) {
					localStorage.setItem('token', this.token)
				}
			} catch (err) {
				console.log(err)
			}
		},
		async signup(email: string, password: string) {
			try {
				const res = await api.post("/auth/signup", { email: email, password: password })
				this.user = res.data
			} catch (err) {
				console.log(err)
			}
		},
		setToken(newToken: string) {
			this.token = newToken
			localStorage.setItem('token', newToken)
		},
		clearToken() {
			this.token = ""
			localStorage.removeItem('token')
		},
	}
})
