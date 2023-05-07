export const postHeader = () => {
	return new Headers({
		"Access-Control-Request-Method": "POST",
		"Access-Control-Request-Headers":
			"origin, content-type, accept, access-control-request-method, authorization",
		Origin: process.env.NEXT_PUBLIC_FRONTEND_HOST || "",
		"Content-Type": "application/json",
		Accept: "application/json",
	});
};
