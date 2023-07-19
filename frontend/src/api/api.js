import {SignInError, SignUpError} from "./errors";

const root = "http://localhost:8080"

async function apiFetch(url, init) {
    return await fetch(root + url, init)
}

async function GET(url) {
    return await apiFetch(url, {
        method: 'GET',
    })
}

async function POST(url, isProtected, data) {
    return await apiFetch(url, {
        method: 'POST',
        body: JSON.stringify(data)
    })
}

export async function signUp(email, phone, firstName, lastName, password, repeatPassword) {
    return POST("/register", false, {
        'email': email,
        'phone': phone,
        'last_name': lastName,
        'password': password,
        'repeat_password': repeatPassword,
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
                    throw Error("Something went wrong")
            }
        }
    }).catch(
        (error) => {
            throw Error("Something went wrong")
        }
    )
}

export async function signIn(email, password) {
    return POST("/login", false, {
        'email': email,
        'password': password
    }).then(async (response) => {
        if (!response.ok) {
            let err = await response.json()
            switch (err.message) {
                case SignInError.emailDoesntExist:
                    throw Error("Invalid credentials!")
                case SignInError.invalidCredentials:
                    throw Error("Invalid credentials!")
                default:
                    throw Error("Something went wrong")
            }
        }
    }).catch(
        (error) => {
            throw Error("Something went wrong")
        }
    )
}
