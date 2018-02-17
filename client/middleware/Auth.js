export default class Auth {
    /**
     * Authenticate a user. Save a token string in Local Storage
     *
     * @param {string} token
     */
    static setToken = token => localStorage.setItem('token', token);

    /**
     * Get a token value.
     *
     * @returns {string} token
     */

    static getToken = () => localStorage.getItem('token');

    /**
     * Check if a user is authenticated - check if a token is saved in Local Storage
     *
     * @returns {boolean}
     */
    static isUserAuthenticated() {
        return this.getToken() !== null &&
               this.getToken() !== undefined;
    }

    /**
     * Deauthenticate a user. Remove a token from Local Storage.
     *
     */
    static removeToken = () => localStorage.removeItem('token');
}
