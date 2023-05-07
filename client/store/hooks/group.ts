import { useRecoilState, useSetRecoilState } from "recoil";
import { groupsAtom } from "../atoms/group";

export const useSetGroups = () => useSetRecoilState(groupsAtom);
export const useGroupsState = () => useRecoilState(groupsAtom);
