// @flow
export default class Auth {
    static setToken = (token: string): void => localStorage.setItem('token', token);

    static getToken = (): ?string => localStorage.getItem('token');

    static removeToken = (): void => localStorage.removeItem('token');

    static doesTokenExist(): boolean {
        return this.getToken() !== null &&
        this.getToken() !== undefined;
    }
}
