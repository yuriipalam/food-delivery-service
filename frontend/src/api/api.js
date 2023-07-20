import {ACCESS_TOKEN_EXPIRED, REFRESH_TOKEN_EXPIRED, SignInError, SignUpError} from "./errors";
import {useAuthStore} from "../store";
import router from "../router";

const root = "http://localhost:8080"

const errors = {
    somethingWentWrong: "Something went wrong"
}

async function apiFetch(url, init) {
    return await fetch(root + url, init).catch((err) => {
        console.log(err.message)
    })
}

async function GET(url, isProtected) {
    return await apiFetch(url, {
        method: 'GET', headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().accessTokenRef,
        } : {},
    })
        .then(async (response) => {
            if (!response.ok) {
                let err = await response.json()
                switch (err.message) {
                    case ACCESS_TOKEN_EXPIRED:
                        return GET_REFRESH().then(async (response) => {
                            if (response.ok) {
                                const data = await response.json()
                                await useAuthStore().setTokens(data.access_token, data.refresh_token)
                                return await GET(url, isProtected)
                            }
                        })
                }
                throw Error(errors.somethingWentWrong)
            }

            return response
        })
}

async function GET_REFRESH() {
    return await apiFetch('/refresh', {
        method: 'GET', headers: {
            Authorization: 'Bearer ' + useAuthStore().refreshTokenRef
        }
    }).then(async (response) => {
        if (!response.ok) {
            let err = await response.json()
            switch (err.message) {
                case REFRESH_TOKEN_EXPIRED:
                    await useAuthStore().signOut()
                    await router.push({name: 'SignIn'})
                    break
                default:
                    throw Error('Something went wrong')
            }
        }
        return response
    })
}

async function POST(url, isProtected, data) {
    return await apiFetch(url, {
        method: 'POST', body: JSON.stringify(data)
    })
}

export async function getSuppliers() {
    return GET("/suppliers", false).then(async (response) => {
        if (!response.ok) {
            throw Error(errors.somethingWentWrong)
        }
        return response.json()
    })
}

export async function getSupplierByID(id) {
    return GET("/supplier/" + id, false).then(async (response) => {
        if (!response.ok) {
            throw Error(errors.somethingWentWrong)
        }
        const data = response.json()
        console.log(data)
        return data
    })
}

export async function getSupplierCategoriesByID(id) {
    return GET("/categories?supplier_id=" + id, false).then(async (response) => {
        if (!response.ok) {
            throw Error(errors.somethingWentWrong)
        }
        return response.json()
    })
}

export async function getSupplierProductsByID(id) {
    return GET("/products?supplier_id=" + id, false).then(async (response) => {
        if (!response.ok) {
            throw Error(errors.somethingWentWrong)
        }
        return response.json()
    })
}

export async function getSuppliersByCategoryID(id) {
    return GET("/suppliers?category_id=" + id, false).then(async (response) => {
        if (!response.ok) {
            throw Error(errors.somethingWentWrong)
        }
        return response.json()
    })
}

export async function getCategories() {
    return GET("/categories", false).then(async (response) => {
        if (!response.ok) {
            throw Error(errors.somethingWentWrong)
        }
        return response.json()
    })
}

export async function signUp(email, phone, firstName, lastName, password, repeatPassword) {
    return POST("/register", false, {
        'email': email, 'phone': phone, 'first_name': firstName, 'last_name': lastName, 'password': password, 'repeat_password': repeatPassword,
    }).then(async (response) => {
        if (!response.ok) {
            let err = await response.json()
            switch (err.message) {
                case SignUpError.emailExists:
                    throw Error("Email already exists!")
                case SignUpError.phoneExists:
                    throw Error("Phone already exists!")
                case SignUpError.passwordMismatch:
                    throw Error("Passwords don't match!")
                default:
                    throw Error(errors.somethingWentWrong)
            }
        }
    })
}

export async function signIn(email, password) {
    return POST("/login", false, {
        'email': email, 'password': password
    }).then(async (response) => {
        if (!response.ok) {
            let err = await response.json()
            switch (err.message) {
                case SignInError.emailDoesntExist:
                    throw Error("Invalid credentials!")
                case SignInError.invalidCredentials:
                    throw Error("Invalid credentials!")
                default:
                    throw Error(errors.somethingWentWrong)
            }
        }
        return response.json()
    })
}

export async function getCustomer() {
    return await GET("/customer", true).then(response => {
        return response.json()
    }).catch(err => {
        if (err === REFRESH_TOKEN_EXPIRED) {

        }
    })
}
