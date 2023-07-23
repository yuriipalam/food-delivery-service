export const ResponseError = {
    somethingWentWrong: "Something went wrong",
    notFound: "Not found",
    sessionExpired: "Session expired"
}

export const ACCESS_TOKEN_EXPIRED = "access token expired"
export const REFRESH_TOKEN_EXPIRED = "refresh token expired"

const EMAIL_EXISTS = "email already exist"
const INVALID_CREDENTIALS = "invalid credentials"
const ID_MUST_BE_INT = "id must be integer"

export const SupplierError = {
    idMustBeInt: ID_MUST_BE_INT,
    supplierNotFound: "supplier not found",
    suppliersNotFound: "no suppliers found",
}

export const CategoryError = {
    idMustBeInt: ID_MUST_BE_INT,
    categoryNotFound: "category not found",
    categoriesNotFound: "no categories found",
}

export const ProductError = {
    idMustBeInt: ID_MUST_BE_INT,
    productNotFound: "product not found",
    productsNotFound: "no products found",
}

export const SignInError = {
    emailDoesntExist: "email doesn't exist",
    invalidCredentials: INVALID_CREDENTIALS,
}

export const SignUpError = {
    emailExists: EMAIL_EXISTS,
    phoneExists: "phone already exist",
    passwordMismatch: "password mismatch"
}

export const CustomerError = {
    customerNotFound: "customer not found",
}

export const OrderError = {
    ordersNotFound: "no orders found",
    atMostTwoSuppliers: "at most two suppliers in one order",
}
