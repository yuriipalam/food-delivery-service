export const ACCESS_TOKEN_EXPIRED = "access token expired"
export const REFRESH_TOKEN_EXPIRED = "refresh token expired"

const EMAIL_EXISTS = "email already exist"
const INVALID_CREDENTIALS = "invalid credentials"
const ID_MUST_BE_INT = "id must be integer"
const IDS_MUST_BE_INT = "ids must be integers"
const KEY_MUST_BE_INT = "key must be integer"
const CANNOT_CREATE_RESPONSE = "cannot create response"
const CANNOT_CREATE_RESPONSES = "cannot create responses"

export const SupplierError = {
    idMustBeInt: ID_MUST_BE_INT,
    idsMustBeInt: IDS_MUST_BE_INT,
    keyMustBeInt: KEY_MUST_BE_INT,
    cannotCreateResponse: CANNOT_CREATE_RESPONSE,
    cannotCreateResponses: CANNOT_CREATE_RESPONSES,
    cannotFetchSupplier: "cannot fetch supplier",
    cannotFetchSuppliers: "cannot fetch suppliers",
    supplierNotFound: "supplier not found",
    suppliersNotFound: "no suppliers found",
}

export const CategoryError = {
    idMustBeInt: ID_MUST_BE_INT,
    keyMustBeInt: KEY_MUST_BE_INT,
    cannotCreateResponse: CANNOT_CREATE_RESPONSE,
    cannotCreateResponses: CANNOT_CREATE_RESPONSES,
    cannotFetchCategory: "cannot fetch category",
    cannotFetchCategories: "cannot fetch categories",
    categoryNotFound: "category not found",
    categoriesNotFound: "no categories found",
}

export const ProductError = {
    idMustBeInt: ID_MUST_BE_INT,
    keyMustBeInt: KEY_MUST_BE_INT,
    cannotCreateResponse: CANNOT_CREATE_RESPONSE,
    cannotCreateResponses: CANNOT_CREATE_RESPONSES,
    cannotFetchProduct: "cannot fetch product",
    cannotFetchProducts: "cannot fetch products",
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
