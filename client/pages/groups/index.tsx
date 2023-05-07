import { Box, Paper, Typography } from "@mui/material";
import { AuthLayout } from "../../components/layout";
import { GroupsList } from "./modules/groups-list";

export const Groups = () => {
	return (
		<AuthLayout>
			<Paper>
				<Box>
					<Typography variant={"h1"}>グループ一覧</Typography>
				</Box>
				<GroupsList />
			</Paper>
		</AuthLayout>
	);
};

export default Groups;
