import { group } from "console";
import { useEffect } from "react";
import { GetQuery } from "../../../../hooks/query";
import { useGroupsState } from "../../../../store/hooks/group";
import { getLocalStorage } from "../../../../utils/local-storage";
import { registerGroup } from "../../api/groups";

export const useRegisterGroup = async () => {
	const groupId = GetQuery("group_id");

	const [groups, setGroups] = useGroupsState();

	useEffect(() => {
		const handleRegisterGroup = async () => {
			const authorization = getLocalStorage("authentication") || "";

			const response = await registerGroup({ authorization, groupId });
			if (response.group && !groups.includes(response.group)) {
				setGroups([...groups, response.group]);
			}
		};
		if (groupId) {
			handleRegisterGroup();
		}
		// lintに従うと無限にregisterGroupリクエストしてしまうので無効化
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, []);
};
