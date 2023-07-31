import {
    ACCESS_TOKEN_EXPIRED,
    CategoryError,
    CustomerError,
    OrderError,
    ProductError,
    REFRESH_TOKEN_EXPIRED,
    ResponseError,
    SignInError,
    SignUpError,
    SupplierError
} from "./errors";
import {useAuthStore} from "../store";

// for development
//const root = "http://localhost:8080"

// for production
const root = "/api"

async function apiFetch(url, init) {
    return await fetch(root + url, init)
}

// helper functions to check auth (tokens expiration) //

async function handleResponse(response, url, isProtected, callback) {
    if (!response.ok) {
        let err = await response.json();
        if (err.message === ACCESS_TOKEN_EXPIRED) {
            return GET_REFRESH().then(async (refreshResponse) => {
                if (refreshResponse.ok) {
                    const data = await refreshResponse.json();
                    await useAuthStore().setTokens(data.access_token, data.refresh_token);
                    return await callback(url, isProtected);
                }
            })
        }
        return Promise.reject(err);
    }
    return response.json();
}

async function GET_REFRESH() {
    return await apiFetch('/refresh', {
        method: 'GET', headers: {
            Authorization: 'Bearer ' + useAuthStore().refreshTokenRef
        }
    }).then(async (response) => {
        if (!response.ok) {
            let err = await response.json()
            if (err.message === REFRESH_TOKEN_EXPIRED) {
                await useAuthStore().signOut()
            }
            return Promise.reject(err)
        }
        return response
    })
}

// GET, POST functions for requests //

async function GET(url, isProtected) {
    return await apiFetch(url, {
        method: 'GET', headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().accessTokenRef,
        } : {},
    }).then(async response => await handleResponse(response, url, isProtected, GET))
}

async function POST(url, isProtected, data) {
    return await apiFetch(url, {
        method: 'POST', body: JSON.stringify(data), headers: isProtected ? {
            Authorization: 'Bearer ' + useAuthStore().accessTokenRef,
        } : {},
    }).then(async response => await handleResponse(response, url, isProtected, POST))
}

// ACCESSING ENDPOINTS //

// SUPPLIER //

export async function getSuppliers() {
    return GET("/suppliers", false)
        .catch(err => {
            switch (err.message) {
                case SupplierError.suppliersNotFound:
                    throw Error(ResponseError.notFound)
                default:
                    throw Error(ResponseError.somethingWentWrong)
            }
        })
}

export async function getSupplierByID(id) {
    return GET("/supplier/" + id, false)
        .catch(err => {
            switch (err.message) {
                case SupplierError.idMustBeInt:
                case SupplierError.supplierNotFound:
                    throw Error(ResponseError.notFound)
                default:
                    throw Error(ResponseError.somethingWentWrong)
            }
        })
}

export async function getSuppliersByCategoryID(id) {
    return GET("/suppliers?category_id=" + id, false)
        .catch(err => {
            switch (err.message) {
                case SupplierError.idMustBeInt:
                case SupplierError.suppliersNotFound:
                    throw Error(ResponseError.notFound)
                default:
                    throw Error(ResponseError.somethingWentWrong)
            }
        })
}

// CATEGORY //

export async function getCategoriesBySupplierID(id) {
    return GET("/categories?supplier_id=" + id, false).catch(err => {
        switch (err.message) {
            case CategoryError.idMustBeInt:
            case CategoryError.categoriesNotFound:
                throw Error(ResponseError.notFound)
            default:
                throw Error(ResponseError.somethingWentWrong)
        }
    })
}

export async function getCategories() {
    return GET("/categories", false)
        .catch(err => {
            switch (err.message) {
                case CategoryError.categoriesNotFound:
                    throw Error(ResponseError.notFound)
                default:
                    throw Error(ResponseError.somethingWentWrong)
            }
        })
}

// PRODUCT //

export async function getProductsBySupplierID(id) {
    return GET("/products?supplier_id=" + id, false)
        .catch(err => {
            switch (err.message) {
                case ProductError.idMustBeInt:
                case ProductError.productsNotFound:
                    throw Error(ResponseError.notFound)
                default:
                    throw Error(ResponseError.somethingWentWrong)
            }
        })
}

// AUTH //

export async function signUp(email, phone, firstName, lastName, password, repeatPassword) {
    return POST("/register", false, {
        'email': email,
        'phone': phone,
        'first_name': firstName,
        'last_name': lastName,
        'password': password,
        'repeat_password': repeatPassword,
    }).catch(err => {
        switch (err.message) {
            case SignUpError.emailExists:
                throw Error("Email already exists!")
            case SignUpError.phoneExists:
                throw Error("Phone already exists!")
            case SignUpError.passwordMismatch:
                throw Error("Passwords don't match!")
            default:
                throw Error(ResponseError.somethingWentWrong)
        }
    })
}

export async function signIn(email, password) {
    return POST("/login", false, {
        'email': email, 'password': password
    }).then(async data => {
        await useAuthStore().setTokens(data.access_token, data.refresh_token)
        await getCustomer().then(async data => {
            await useAuthStore().setUser(data.id, data.email, data.phone, data.first_name, data.last_name)
        })
    }).catch(err => {
        switch (err.message) {
            case SignInError.emailDoesntExist:
            case SignInError.invalidCredentials:
                throw Error("Invalid credentials!")
            default:
                throw Error(ResponseError.somethingWentWrong)
        }
    })
}

// CUSTOMER //

export async function getCustomer() {
    return await GET("/customer", true)
        .catch(err => {
            switch (err.message) {
                case CustomerError.customerNotFound:
                    throw Error(ResponseError.notFound)
                case REFRESH_TOKEN_EXPIRED:
                    throw Error(ResponseError.sessionExpired)
                default:
                    throw Error(ResponseError.somethingWentWrong)
            }
        })
}

// ORDER //

export async function createOrder(customerID, recipientFullName, address, price, supplierIDs, products) {
    return await POST("/orders", true, {
        'customer_id': customerID,
        'recipient_full_name': recipientFullName,
        'address': address,
        'price': price,
        'supplier_ids': supplierIDs,
        'products': products
    }).catch(err => {
        console.log(err)
        switch (err.message) {
            case OrderError.atMostTwoSuppliers:
                throw Error("At most two suppliers can be chosen!")
            case REFRESH_TOKEN_EXPIRED:
                throw Error(ResponseError.sessionExpired)
            default:
                throw Error(ResponseError.somethingWentWrong)
        }
    })
}

export async function getOrders() {
    return await GET("/orders", true)
        .catch(err => {
            switch (err.message) {
                case OrderError.ordersNotFound:
                    return []
                case REFRESH_TOKEN_EXPIRED:
                    throw Error(ResponseError.sessionExpired)
                default:
                    throw Error(ResponseError.somethingWentWrong)
            }
        })
}
