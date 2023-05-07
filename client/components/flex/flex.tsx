import { Box } from "@mui/system";
import { ComponentProps } from "react";

type FlexType = Omit<ComponentProps<typeof Box>, "flex">;

export const Flex = (props: FlexType) => (
	<Box {...props} display={"flex"}></Box>
);
