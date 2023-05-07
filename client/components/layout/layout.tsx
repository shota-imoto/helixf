import { Box } from "@mui/material";
import { ReactElement } from "react";
import { Authentication } from "./authenticaton";
import Header from "./header";

type Props = {
	children: ReactElement;
};

export const Layout = ({ children }: Props) => {
	return (
		<Box>
			<>
				<Header />
				<Box>{children}</Box>
			</>
		</Box>
	);
};
