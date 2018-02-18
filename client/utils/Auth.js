export default class Auth {
    static setToken = token => localStorage.setItem('token', token);

    static getToken = () => localStorage.getItem('token');

    static removeToken = () => localStorage.removeItem('token');

    static doesTokenExist() {
        return this.getToken() !== null &&
        this.getToken() !== undefined;
    }
}
