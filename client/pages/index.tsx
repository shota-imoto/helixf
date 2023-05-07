import * as React from "react";
import Box from "@mui/material/Box";
import { AuthLayout } from "../components/layout";

export default function Home() {
	return (
		<AuthLayout>
			<Box></Box>
		</AuthLayout>
	);
}

// export const getServerSideProps: GetServerSideProps = async (context) =>
// 	authentication(context);
