import axios from "axios"

export const signIn = (username: string, password: string) => {
    return axios.post('/api/auth/login', { username, password })
}

export const signUp = (data: RegisterInput) => {
    return axios.post('/api/auth/signup', data)
}