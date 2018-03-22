// @flow
export default class Auth {
    static setToken = (token: string) => localStorage.setItem('token', token);

    static getToken = (): void => localStorage.getItem('token');

    static removeToken = (): void => localStorage.removeItem('token');

    static doesTokenExist(): boolean {
        return this.getToken() !== null &&
        this.getToken() !== undefined;
    }
}
