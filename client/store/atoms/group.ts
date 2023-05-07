import { atom } from "recoil";
import { Group } from "../../models/group";

export const groupsAtom = atom<Group[]>({
	key: "Group",
	default: [],
});
