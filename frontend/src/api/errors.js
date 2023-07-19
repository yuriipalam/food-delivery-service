const EMAIL_EXISTS = "email already exist"
const INVALID_CREDENTIALS = "invalid credentials"

export const SignInError = {
    emailDoesntExist: "email doesn't exist",
    invalidCredentials: INVALID_CREDENTIALS,
}

export const SignUpError = {
    emailExists: EMAIL_EXISTS,
    phoneExists: "phone already exist",
    passwordMismatch: "password mismatch"
}
