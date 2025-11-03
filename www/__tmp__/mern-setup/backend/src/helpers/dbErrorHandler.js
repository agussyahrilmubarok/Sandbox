'use strict';

/** 🔹 Handle unique field errors */
const getUniqueErrorMessage = (err) => {
    try {
        const fieldName = Object.keys(err.keyValue)[0];
        return `${fieldName.charAt(0).toUpperCase() + fieldName.slice(1)} already exists`;
    } catch (ex) {
        return 'Unique field already exists';
    }
};

/** 🔹 Get error message from error object */
const getErrorMessage = (err) => {
    if (!err) return 'An unknown error occurred';

    if (err.code) {
        switch (err.code) {
            case 11000:
            case 11001:
                return getUniqueErrorMessage(err);
            default:
                return 'Something went wrong';
        }
    }

    if (err.errors) {
        for (let field in err.errors) {
            if (err.errors[field].message) return err.errors[field].message;
        }
    }

    if (err.name === 'CastError') {
        return `Invalid ${err.path}: ${err.value}`;
    }

    return err.message || 'An unexpected error occurred';
};

export default { getErrorMessage };
