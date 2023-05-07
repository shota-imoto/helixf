import { useRecoilState, useRecoilValue, useSetRecoilState } from "recoil";
import { authenticationAtom } from "../atoms/authentication";

export const useAuthenticationValue = () => useRecoilValue(authenticationAtom);
export const useSetAuthentication = () => useSetRecoilState(authenticationAtom);
export const useAuthenticationState = () => useRecoilState(authenticationAtom);
