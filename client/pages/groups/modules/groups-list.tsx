import { Box } from "@mui/material";
import { useEffect, useState } from "react";
import { Loading } from "../../../components/loading/loading";
import { Group } from "../../../models/group";
import { useGroupsState } from "../../../store/hooks/group";
import { getLocalStorage } from "../../../utils/local-storage";
import { fetchGroups } from "../api/groups";
import { useRegisterGroup } from "./hooks/use-register-group";

export const GroupsList = () => {
	const [groups, setGroups] = useGroupsState();
	useRegisterGroup();

	useEffect(() => {
		const fetchGroupsAsync = async () => {
			const authentication = getLocalStorage("authentication") || "";
			const res = await fetchGroups(authentication);
			setGroups(res.groups);
		};
		fetchGroupsAsync();
	}, [setGroups]);

	if (!groups) return <Loading />;

	return (
		<Box>
			{groups.map((g) => (
				<Box key={g.id}>{g.groupName}</Box>
			))}
		</Box>
	);
};
