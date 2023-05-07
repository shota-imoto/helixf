type LocalStorageKey = "authentication" | "refreshToken";
type SetLocalStorageProps = {
	key: LocalStorageKey;
	value: string;
};

export const getLocalStorage = (key: LocalStorageKey) => {
	return localStorage.getItem(key);
};

export const setLocalStorage = ({ key, value }: SetLocalStorageProps) => {
	return localStorage.setItem(key, value);
};
