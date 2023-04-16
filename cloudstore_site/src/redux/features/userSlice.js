import { createSlice } from '@reduxjs/toolkit'

export const userSlice = createSlice({
    name: 'user',
    initialState: {
        loggedin: null,
        value: null,
        cart: null
    },
    reducers: {
        login: (state, action) => {
            state.value = action.payload
            state.loggedin = true
        },
        logout: (state) => {
            state.value = null
            state.loggedin = false
        },
        clearCart: (state, action) => {
            state.cart = null
        },
        updateCart: (state, action) => {
            state.cart = action.payload
        },
    },
})

export const { login, logout, updateCart, clearCart } = userSlice.actions

export default userSlice.reducer
