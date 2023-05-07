import { CircularProgress } from "@mui/material";
import { Flex } from "../flex/flex";

export const Loading = () => {
	return (
		<Flex sx={{ justifyContent: "center", width: "100%" }}>
			<CircularProgress />
		</Flex>
	);
};
