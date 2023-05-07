// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { GetServerSideProps, NextApiRequest, NextApiResponse } from "next";

// export const getServerSideProps: GetServerSideProps = async (context) => {
// 	return {
// 		redirect: {
// 			permanent: false,
// 			destination: "https://google.co.jp",
// 		},
// 	};
// 	// APIやDBからのデータ取得処理などを記載
// 	const authorizationHeader = context.req.headers.authorization;
// 	const authorizationQuery = context.query.authorization;

// 	if (!authorizationHeader && !authorizationQuery) {
// 		return {
// 			redirect: {
// 				permanent: false,
// 				destination: "https://google.co.jp",
// 			},
// 		};
// 	}
// 	return { props: {} };
// };
