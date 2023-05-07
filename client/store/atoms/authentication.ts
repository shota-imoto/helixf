import { atom } from "recoil";

type User = {
	name: string;
};

export type AuthenticationStateType = {
	authorization: string;
	refleshToken: string;
	user: User;
};

export const authenticationAtom = atom<AuthenticationStateType>({
	key: "AuthenticationState",
	default: { authorization: "", refleshToken: "", user: { name: "" } },
});
