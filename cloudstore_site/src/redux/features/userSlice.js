import { createSlice } from '@reduxjs/toolkit'

export const userSlice = createSlice({
    name: 'user',
    initialState: {
        loggedin: null,
        value: null,
    },
    reducers: {
        login: (state, action) => {

            state.value = action.payload
            state.loggedin = true
        },
        cart: (state, action) => {
            state.value = action.payload
            state.loggedin = true
        },
        logout: (state) => {
            state.value = null
            state.loggedin = false
        },

    },
})

export const { login, logout } = userSlice.actions

export default userSlice.reducer