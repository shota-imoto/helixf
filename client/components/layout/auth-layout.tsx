import { ReactElement } from "react";
import { Authentication } from "./authenticaton";
import { Layout } from "./layout";
type Props = {
	children: ReactElement;
};
export const AuthLayout = ({ children }: Props) => {
	return (
		<Authentication>
			<Layout>{children}</Layout>
		</Authentication>
	);
};
